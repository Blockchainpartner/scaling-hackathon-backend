package controllers

import (
	"net/http"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type registriesController struct{}

type CairoProgramInputRegistry struct {
	OldRegistry []string `json:"oldRegistry"`
	NewRegistry []string `json:"newRegistry"`
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
