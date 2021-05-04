/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Tuesday May 4th 2021
**	@Filename:				users copy.go
******************************************************************************/

package controllers

import (
	"net/http"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/microgolang/logs"
	"go.mongodb.org/mongo-driver/bson"
)

type proofController struct{}

type CairoProgramInputClaim struct {
	Address  string   `json:"address"`
	Secret   string   `json:"secret"`
	Registry []string `json:"registry"`
}

//Prove is triggered in order to prove something
func (y proofController) Prove(c *gin.Context) {
	type proofData struct {
		Address *string `json:"address"`
		Secret  *string `json:"secret"`
	}
	/**************************************************************************
	** Do we have a valid body ?
	**************************************************************************/
	var body proofData
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, err, gin.H{`error`: `bad request`})
		return
	}

	/**************************************************************************
	** Retrieve the user's UUID
	**************************************************************************/
	var registryKey = c.Param(`registryKey`)
	if registryKey == `` {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `missing registryKey`})
		return
	}

	/**************************************************************************
	** Let's retrieve the identities in the DB
	**************************************************************************/
	identities, err := models.NewRegistryMapping().ListBy(bson.M{`registryKey`: registryKey})
	if err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to retrieve identities`})
		return
	}
	/**************************************************************************
	** Let's populate the registries
	**************************************************************************/
	registry := []string{}
	for _, id := range identities {
		registry = append(registry, *id.Identity)
	}

	utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `Proving identity in registry ...`,
		`type`:     `info`,
	})

	/**************************************************************************
	** First python job
	** This one will send the program & the inputs to the sharp prover in order
	** to validate the outputs and it's execution. The prover is async, that's
	** why we are starting with this one.
	**************************************************************************/
	resultSharp, err := execPythonSharpClaim(CairoProgramInputClaim{Address: *body.Address, Secret: *body.Secret, Registry: registry})
	if err != nil {
		logs.Error(err)
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to submit program to sharp`})
		return
	} else if resultSharp[0] == `0` && resultSharp[1] == `0` {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `Invalid claim`})
		return
	}
	utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof has been sent to Sharp`,
		`type`:     `success`,
	})

	/**************************************************************************
	** Second python job
	** This one will compute the outputs through the cairo program in order to
	** be able to submit the tx.
	**************************************************************************/
	result, err := execPythonCairoCompileClaim(CairoProgramInputClaim{Address: *body.Address, Secret: *body.Secret, Registry: registry})
	if err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to compile claim`})
		return
	}
	utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof has been compiled`,
		`type`:     `success`,
	})
	utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `Waiting for Sharp to process the transaction`,
		`type`:     `info`,
	})

	/**************************************************************************
	** Third python job -> Waiting for the job to be PROCESSED
	**************************************************************************/
	resultStatus, err := execPythonSharpStatusClaim(registryKey, resultSharp[0], resultSharp[1])
	if err != nil || !resultStatus {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `error while waiting for sharp`})
		return
	}
	utils.NewPusher().Claims.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof has been processed`,
		`type`:     `success`,
	})

	c.JSON(http.StatusOK, gin.H{`proof`: result[0], `hash`: result[1]})
}
