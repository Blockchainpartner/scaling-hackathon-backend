package controllers

import (
	"net/http"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

type usersController struct{}

//FindUser is triggered to retrieve the informations about a specific user
func (y usersController) FindUser(c *gin.Context) {
	type bodyData struct {
		UUID     string `json:"uuid"`
		Password string `json:"password"`
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
	** Does the users actually exists in the DB ?
	**************************************************************************/
	user, err := models.NewUser().FindBy(bson.M{`UUID`: body.UUID, `password`: body.Password})
	if err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to find this user`})
		return
	}

	c.JSON(http.StatusOK, user)
}

//AddUser is triggered to save a specific user in the database
func (y usersController) AddUser(c *gin.Context) {
	/**************************************************************************
	** Do we have a valid body ?
	**************************************************************************/
	var body models.User
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, err, gin.H{`error`: `bad request`})
		return
	}
	/**************************************************************************
	** Is the address already used (avoid duplicates)
	**************************************************************************/
	exists := models.NewUser().Exists(bson.M{`address`: body.Address})
	if exists {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `this address is already used`})
		return
	}
	/**************************************************************************
	** Let's save this new user
	**************************************************************************/
	body.IsVerified = utils.BoolToPtr(false)
	err := body.Post()
	if err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to save this user`})
		return
	}

	c.JSON(http.StatusOK, body)
}
