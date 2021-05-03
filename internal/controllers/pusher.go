/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Monday May 3rd 2021
**	@Filename:				pusher.go
******************************************************************************/

package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type pusherController struct{}

//AuthPusher will perform the pusher auth for an user
func (y pusherController) AuthPusher(c *gin.Context) {
	params, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.AbortWithLog(c, http.StatusForbidden, err, gin.H{`error`: `impossible to perform auth`})
		return
	}

	response, err := utils.NewPusher().PrivateAuth(params)
	if err != nil {
		utils.AbortWithLog(c, http.StatusForbidden, err, gin.H{`error`: `impossible to perform auth`})
		return
	}

	c.String(http.StatusOK, string(response))
}
