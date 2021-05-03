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
	"go.mongodb.org/mongo-driver/bson"
)

type registriesController struct{}

type CairoProgramInput struct {
	OldRegistry []string `json:"oldRegistry"`
	NewRegistry []string `json:"newRegistry"`
}

func execPythonCairoCompile(input CairoProgramInput) ([]string, error) {
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	cairoScriptPath, err := filepath.Abs(`../scripts/cairo.py`)
	if err != nil {
		return nil, err
	}
	cairoProgramPath, err := filepath.Abs(`../scripts/registryHash.json`)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"python3",
		cairoScriptPath,
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
func execPythonSharp(input CairoProgramInput) ([]string, error) {
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
	cairoProgramPath, err := filepath.Abs(`../scripts/registryHash.json`)
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
func execPythonSharpStatus(fact string) (bool, error) {
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
		if string(out) == `IN_PROGRESS` {
			time.Sleep(time.Second * 10)
		} else if string(out) == `PROCESSED` {
			return true, err
		} else {
			return false, errors.New(`Invalid fact`)
		}
	}
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
		OldHash    *string  `json:"oldHash"`
		NewHash    *string  `json:"newHash"`
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

	/**************************************************************************
	** First python job
	** This one will send the program & the inputs to the sharp prover in order
	** to validate the outputs and it's execution. The prover is async, that's
	** why we are starting with this one.
	**************************************************************************/
	resultSharp, err := execPythonSharp(CairoProgramInput{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to compute proof`})
		return
	}

	/**************************************************************************
	** Second python job
	** This one will compute the outputs through the cairo program in order to
	** be able to submit the tx.
	**************************************************************************/
	result, err := execPythonCairoCompile(CairoProgramInput{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to compute proof`})
		return
	}

	/**************************************************************************
	** Third python job -> Waiting for the job to be PROCESSED
	**************************************************************************/
	resultStatus, err := execPythonSharpStatus(resultSharp[0])
	if err != nil || !resultStatus {
		utils.AbortWithLog(c, http.StatusNotFound, err, gin.H{`error`: `impossible to check the fact`})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		`jobID`:  resultSharp[0],
		`fact`:   resultSharp[1],
		`output`: result,
	})
}
