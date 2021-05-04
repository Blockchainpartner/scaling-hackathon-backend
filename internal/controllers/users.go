package controllers

import (
	"errors"
	"math/big"
	"net/http"
	"sync"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/ethereum"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/microgolang/logs"
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
	var wg sync.WaitGroup
	for _, r := range user.Registries {
		if *r.Key == `161373187550089867448191830760110801114155294027693593477164529548269146668` && user.KYC.Dob.Age >= 12 && user.KYC.Dob.Age <= 24 {
			wg.Add(1)
			go AddIdentitiesToRegistry(&wg, []string{*r.Secret}, *r.Key)
		}
		if *r.Key == `418791004851046193537070596848530790547129451305514433175127304050849890764` && user.KYC.Dob.Age >= 60 {
			wg.Add(1)
			go AddIdentitiesToRegistry(&wg, []string{*r.Secret}, *r.Key)
		}
		if *r.Key == `374546399808851745807054416014379391823657543778127138954064098322040293325` && *user.KYC.Disabled {
			wg.Add(1)
			go AddIdentitiesToRegistry(&wg, []string{*r.Secret}, *r.Key)
		}
	}

	wg.Wait()
	c.JSON(http.StatusOK, user)
}

//AddIdentities will a new identity to the registry, but only if the proof is valid
func AddIdentitiesToRegistry(wg *sync.WaitGroup, newIdentities []string, registryKey string) error {
	defer wg.Done()

	/**************************************************************************
	** Let's retrieve the identities in the DB
	**************************************************************************/
	identities, err := models.NewRegistryMapping().ListBy(bson.M{`registryKey`: registryKey})
	if err != nil {
		logs.Error(err)
		return err
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
	newRegistry = append(newRegistry, newIdentities...)
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `Starting proof for registry`,
		`type`:     `info`,
	})

	/**************************************************************************
	** First python job
	** This one will send the program & the inputs to the sharp prover in order
	** to validate the outputs and it's execution. The prover is async, that's
	** why we are starting with this one.
	**************************************************************************/
	resultSharp, err := execPythonSharpRegistry(CairoProgramInputRegistry{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		logs.Error(err)
		return err
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof has been sent to Sharp`,
		`type`:     `success`,
	})

	/**************************************************************************
	** Second python job
	** This one will compute the outputs through the cairo program in order to
	** be able to submit the tx.
	**************************************************************************/
	result, err := execPythonCairoCompileRegistry(CairoProgramInputRegistry{OldRegistry: oldRegistry, NewRegistry: newRegistry})
	if err != nil {
		logs.Error(err)
		return err
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof has been compiled`,
		`type`:     `success`,
	})
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `Waiting for Sharp to process the transaction`,
		`type`:     `info`,
	})

	/**************************************************************************
	** Third python job -> Waiting for the job to be PROCESSED
	**************************************************************************/
	resultStatus, err := execPythonSharpStatusRegistry(registryKey, resultSharp[0], resultSharp[1])
	if err != nil || !resultStatus {
		return err
	}
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `The proof is now processed`,
		`type`:     `success`,
	})

	oldRegistryHash, success0 := big.NewInt(0).SetString(result[0], 10)
	newRegistryHash, success1 := big.NewInt(0).SetString(result[1], 10)
	cairoPrime, success2 := big.NewInt(0).SetString(`800000000000011000000000000000000000000000000000000000000000001`, 16)
	if !success0 || !success1 || !success2 {
		return errors.New(`invalid cairo prime`)
	}
	if result[0][0] == '-' {
		oldRegistryHash = big.NewInt(0).Add(oldRegistryHash, cairoPrime)
	}
	if result[1][0] == '-' {
		newRegistryHash = big.NewInt(0).Add(newRegistryHash, cairoPrime)
	}

	err = ethereum.UpdateRegistry(registryKey, oldRegistryHash, newRegistryHash)
	if err != nil {
		return err
	}

	for _, id := range newIdentities {
		count, _ := models.NewRegistryMapping().Count(bson.M{`registryKey`: registryKey})
		newID := models.NewRegistryMapping().Init()
		newID.RegistryKey = utils.StrToPtr(registryKey)
		newID.Identity = utils.StrToPtr(id)
		newID.IdentityIndex = utils.Uint64ToPtr(uint64(count) + uint64(1))
		newID.Post()
	}
	return nil
}
