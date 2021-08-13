package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

type TwitterClient struct {
	API *anaconda.TwitterApi
}

func createTwitterClient() *TwitterClient {
	godotenv.Load()
	return &TwitterClient{
		API: anaconda.NewTwitterApiWithCredentials(
			os.Getenv("ACCESS_TOKEN"),
			os.Getenv("ACCESS_TOKEN_SECRET"),
			os.Getenv("API_KEY"),
			os.Getenv("API_SECRET_KEY")),
	}
}

func (t *TwitterClient) PostImage(imagePath string, text string) (err error) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return
	}
	mediaResponse, err := t.API.UploadMedia(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		return
	}
	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))

	_, err = t.API.PostTweet(text, v)
	if err != nil {
		return
	}
	return
}
