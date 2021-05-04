/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Tuesday May 4th 2021
**	@Filename:				pythonCairo.go
******************************************************************************/

package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

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
	return true, nil
}

func execPythonCairoCompileClaim(input CairoProgramInputClaim) ([]string, error) {
	jsonInput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	cairoScriptPath, err := filepath.Abs(`../scripts/cairo.py`)
	if err != nil {
		return nil, err
	}
	cairoProgramPath, err := filepath.Abs(`../scripts/claimProgram.json`)
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
func execPythonSharpClaim(input CairoProgramInputClaim) ([]string, error) {
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
	cairoProgramPath, err := filepath.Abs(`../scripts/claimProgram.json`)
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
func execPythonSharpStatusClaim(registryKey, fact, job string) (bool, error) {
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
		utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
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
	return true, nil
}
