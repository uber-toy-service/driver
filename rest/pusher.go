package rest_api

// The sole purpose of this package is to push notifications
// to the Clients that are registered here.

// No response is excpected here. The Clients should use REST
// API to respond

import "github.com/pusher/pusher-http-go"

var pusherClient pusher.Client

func InitPusherClient() {
	pusherClient = pusher.Client{
		AppID:   "1118332",
		Key:     "09d793073c6c74717c29",
		Secret:  "c6c093f89649b0bb5e4e",
		Cluster: "eu",
		Secure:  true,
	}
}
