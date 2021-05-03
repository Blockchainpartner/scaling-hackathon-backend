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

//FindUser is triggered to retrieve the information about a specific user
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
	exists := models.NewUser().Exists(bson.M{`UUID`: body.UUID})
	if exists {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `this uuid is already used`})
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

//ValidateUser should be called by the KYC organism in order to validate the information provided by the user.
//This will also add the user in all the valid registries.
func (y usersController) ValidateUser(c *gin.Context) {
	/**************************************************************************
	** Retrieve the user's UUID
	**************************************************************************/
	var userUUID = c.Param(`userUUID`)
	if userUUID == `` {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `missing userUUID`})
		return
	}

	/**************************************************************************
	** Update the status of the user from not verified to verified
	**************************************************************************/
	user, err := models.NewUser().FindByUUID(userUUID)
	if err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to find this user`})
		return
	}
	user.IsVerified = utils.BoolToPtr(true)
	if err := user.SelfUpdate(); err != nil {
		utils.AbortWithLog(c, http.StatusBadRequest, nil, gin.H{`error`: `impossible to save this user`})
		return
	}

	/**************************************************************************
	** Check if the user should be registered in a registry & add it in theses
	** registries.
	**************************************************************************/
	for _, r := range user.Registries {
		if *r.Key == `161373187550089867448191830760110801114155294027693593477164529548269146668` && user.KYC.Dob.Age >= 12 && user.KYC.Dob.Age <= 24 {
			AddIdentitiesToRegistry([]string{*r.Secret}, *r.Key)
		}
		if *r.Key == `418791004851046193537070596848530790547129451305514433175127304050849890764` && user.KYC.Dob.Age >= 60 {
			AddIdentitiesToRegistry([]string{*r.Secret}, *r.Key)
		}
		if *r.Key == `374546399808851745807054416014379391823657543778127138954064098322040293325` && *user.KYC.Disabled {
			AddIdentitiesToRegistry([]string{*r.Secret}, *r.Key)
		}
	}

	c.JSON(http.StatusOK, user)
}
