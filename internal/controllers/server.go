package controllers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewRouter create the routes and setup the server
func NewRouter() *gin.Engine {
	gin.EnableJsonDecoderDisallowUnknownFields()
	router := gin.New()
	router.Use(gin.Recovery())
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logger, time.RFC3339, true))
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{`Origin`, `Content-Length`, `Content-Type`, `Authorization`},
	}
	router.Use(cors.New(corsConf))

	{
		c := usersController{}
		router.POST(`user/find`, c.FindUser)
		router.POST(`user/add`, c.AddUser)
		router.POST(`user/validate/:userUUID`, c.ValidateUser)
	}

	{
		c := registriesController{}
		router.GET(`registries`, c.ListRegistries)
		router.GET(`registry/:registryKey/identities`, c.ListIdentities)
		router.POST(`registry/:registryKey/identities`, c.AddIdentities)
	}

	{
		c := pusherController{}
		router.POST(`pusher`, c.AuthPusher)
	}

	return router
}
