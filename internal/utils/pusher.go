/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Monday May 3rd 2021
**	@Filename:				pusher.go
******************************************************************************/

package utils

import (
	"github.com/pusher/pusher-http-go/v5"
)

var pusherClient = pusher.Client{}

//Channel define a channel from pusher
type Channel struct {
	ID         string
	EventsFunc map[string]func(interface{})
}

//Pusher define the pusher class
type Pusher struct {
	Identities *Channel
	Claims     *Channel
}

//IdentitiesChannel represent the list of channels for the identities
var IdentitiesChannel = &Channel{
	ID: `identity`,
	EventsFunc: map[string]func(data interface{}){
		`PROCESS`: func(data interface{}) { _ = pusherClient.Trigger(`private-identity`, `processIdentity`, data) },
	},
}

//ClaimsChannel represent the list of channels for the claims
var ClaimsChannel = &Channel{
	ID: `claim`,
	EventsFunc: map[string]func(data interface{}){
		`PROCESS`: func(data interface{}) { _ = pusherClient.Trigger(`private-claims`, `processClaim`, data) },
	},
}

//InitPusher will init the pusher configuration based on the env variables
func InitPusher() {
	pusherClient = pusher.Client{
		AppID:   PusherID,
		Key:     PusherKey,
		Secret:  PusherSecret,
		Cluster: `eu`,
	}
}

// NewPusher create a new pusher Object
func NewPusher() (x *Pusher) {
	return &Pusher{
		Identities: IdentitiesChannel,
		Claims:     ClaimsChannel,
	}
}

//Push will sent a message to the pusher channel
func (c *Channel) Push(event string, payload interface{}) {
	c.EventsFunc[event](payload)
}

//PrivateAuth will perform the private auth for an user
func (p *Pusher) PrivateAuth(params []byte) ([]byte, error) {
	return pusherClient.AuthenticatePrivateChannel(params)
}
