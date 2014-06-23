package main

import (
	"github.com/mrjones/oauth"
	"os"
)

var consumerKey string = os.Getenv("CONSUMER_KEY")
var consumerSecret string = os.Getenv("COMSUMER_SECRET")

var Consumer = oauth.NewConsumer(
	consumerKey,
	consumerSecret,
	oauth.ServiceProvider{
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	})
