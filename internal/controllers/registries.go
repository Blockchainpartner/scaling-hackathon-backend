package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/microgolang/logs"
	"go.mongodb.org/mongo-driver/bson"
)

type registriesController struct{}

type CairoProgramInputRegistry struct {
	OldRegistry []string `json:"oldRegistry"`
	NewRegistry []string `json:"newRegistry"`
}

func execPythonCairoCompileRegistry(input CairoProgramInputRegistry) ([]string, error) {
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	cairoScriptPath, err := filepath.Abs(`../scripts/cairo.py`)
	if err != nil {
		return nil, err
	}
	cairoProgramPath, err := filepath.Abs(`../scripts/registryProgram.json`)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"python3",
		cairoScriptPath,
		`--exception`,
		`-1`,
		`--program`,
		cairoProgramPath,
		`--program_input`,
		string(jsonInput),
	)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outStr := strings.ReplaceAll(string(out), `'`, `"`)

	result := []string{}
	err = json.Unmarshal([]byte(outStr), &result)
	return result, err
}
func execPythonSharpRegistry(input CairoProgramInputRegistry) ([]string, error) {
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), ``)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	if _, err = tmpFile.Write(jsonInput); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}

	cairoScriptPath, err := filepath.Abs(`../scripts/sharp.py`)
	if err != nil {
		return nil, err
	}
	cairoProgramPath, err := filepath.Abs(`../scripts/registryProgram.json`)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"python3",
		cairoScriptPath,
		`submit`,
		`--program`,
		cairoProgramPath,
		`--program_input`,
		tmpFile.Name(),
	)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	outStr := strings.ReplaceAll(string(out), `'`, `"`)
	result := []string{}
	err = json.Unmarshal([]byte(outStr), &result)
	return result, err
}
func execPythonSharpStatusRegistry(registryKey, fact, job string) (bool, error) {
	cairoScriptPath, err := filepath.Abs(`../scripts/sharp.py`)
	if err != nil {
		return false, err
	}

	for {
		cmd := exec.Command(
			"python3",
			cairoScriptPath,
			`status`,
			fact,
		)
		out, err := cmd.Output()
		if err != nil {
			return false, err
		}
		utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
			`registry`: registryKey,
			`step`:     `Still waiting for Sharp to process the program`,
			`type`:     `info`,
		})
		if string(out) == `IN_PROGRESS` {
			time.Sleep(time.Second * 10)
		} else if string(out) == `PROCESSED` {
			break
		} else {
			return false, errors.New(`Invalid fact`)
		}
	}

	logs.Success(job)

	cmd := exec.Command(
		"python3",
		cairoScriptPath,
		`is_verified`,
		job,
		`--node_url`,
		utils.EthNodeURIHttp,
	)
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	logs.Info(string(out))
	return true, nil
}

//ListRegistries is triggered to retrieve the list of all the registries
func (y registriesController) ListRegistries(c *gin.Context) {
	registries, err := models.NewRegistry().List()
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `could not find registries`})
		return
	}
	result := []string{}
	for _, r := range registries {
		result = append(result, *r.Key)
	}
	c.JSON(http.StatusOK, result)
}

//ListIdentities is triggered to retrieve the list of all the identities in a registry
func (y registriesController) ListIdentities(c *gin.Context) {
	var registryKey = c.Param(`registryKey`)
	if registryKey == `` {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `missing registryKey`})
		return
	}

	identities, err := models.NewRegistryMapping().ListBy(bson.M{`registryKey`: registryKey})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `could not find identities`})
		return
	}
	idMapping := []string{}
	for _, id := range identities {
		idMapping = append(idMapping, *id.Identity)
	}
	c.JSON(http.StatusOK, idMapping)
}

//AddIdentities will a new identity to the registry, but only if the proof is valid
func (y registriesController) AddIdentities(c *gin.Context) {
	type bodyData struct {
		Identities []string `json:"identities"`
	}

	var registryKey = c.Param(`registryKey`)
	if registryKey == `` {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `missing registryKey`})
		return
	}
	/**************************************************************************
	** Do we have a valid body ?
	**************************************************************************/
	var body bodyData
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, err, gin.H{`error`: `bad request`})
		return
	}

	/**************************************************************************
	** Let's retrieve the identities in the DB
	**************************************************************************/
	identities, err := models.NewRegistryMapping().ListBy(bson.M{`registryKey`: registryKey})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `could not find identities`})
		return
	}

	/**************************************************************************
	** Let's populate the registries
	**************************************************************************/
	oldRegistry := []string{}
	newRegistry := []string{}
	for _, id := range identities {
		oldRegistry = append(oldRegistry, *id.Identity)
		newRegistry = append(newRegistry, *id.Identity)
	}
	newRegistry = append(newRegistry, body.Identities...)
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`step`: `will-do`})

	/**************************************************************************
	** First python job
	** This one will send the program & the inputs to the sharp prover in order
	** to validate the outputs and it's execution. The prover is async, that's
	** why we are starting with this one.
	**************************************************************************/
	resultSharp, err := execPythonSharpRegistry(CairoProgramInputRegistry{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to compute proof`})
		return
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`step`: `sharp-send`})

	/**************************************************************************
	** Second python job
	** This one will compute the outputs through the cairo program in order to
	** be able to submit the tx.
	**************************************************************************/
	result, err := execPythonCairoCompileRegistry(CairoProgramInputRegistry{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to compute proof`})
		return
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`step`: `computation-ok`})

	/**************************************************************************
	** Third python job -> Waiting for the job to be PROCESSED
	**************************************************************************/
	resultStatus, err := execPythonSharpStatusRegistry(registryKey, resultSharp[0], resultSharp[1])
	if err != nil || !resultStatus {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to check the fact`})
		return
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`step`: `sharp-ok`})

	c.JSON(http.StatusOK, gin.H{
		`jobID`:  resultSharp[0],
		`fact`:   resultSharp[1],
		`output`: result,
	})
}
