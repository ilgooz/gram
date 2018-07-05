package main

import (
	"errors"

	"github.com/ilgooz/gram"
)

var unvalidURLError = errors.New("Please provide a valid url")
var serverError = errors.New("Oops. Something went wrong in our side. We'll ditch into it. Please try again later.")

type App struct {
	data      *Data
	pipelines *Pipelines
	handlers  *Handlers

	videoURL string

	videoDownloadButton *gram.Button
	tweetButton         *gram.Button
}

type Pipelines struct {
	downloadVideo *gram.Pipeline
	tweet         *gram.Pipeline
}

type Handlers struct {
	downloadVideoHandler *downloadVideoHandler
	tweetHandler         *tweetHandler
}

type Data struct {
	downloadVideos map[string]*DownloadVideo
	tweets         map[string]*Tweet
}

func main() {
	app := &App{
		data: &Data{
			downloadVideos: make(map[string]DownloadVideo, 0),
			tweets:         make(map[string]Tweet, 0),
		},
		videoDownloadButton: gram.NewButton("Download"),
		tweetButton:         gram.NewButton("Tweet"),
	}

	app.pipelines.downloadVideo = gram.Cmd("/dv", "Download videos from given link").
		Step(app.handlers.downloadVideoHandler).
		Step(app.handlers.processVideoHandler)

	anyPipeline = gram.Any().
		Step(app.investigateHandler).
		Step(app.responseHandler)

	b := gram.New()
	b.Attach(
		app.pipelines.downloadVideo,
		app.pipelines.tweet,
		anyPipeline,
	)
	b.Start()
}

func (app *App) investigateHandler(c *gram.Context) {
	var message interface{}
	err := c.Get(&message)
	if err != nil {
		c.Cancel()
		return
	}

	switch v := message.(type) {
	case string:
		if !isVideoURL(v) {
			err := c.Send("I don't understand what you want to do. You can check available commands by typing '/'")
			if err != nil {
				c.Cancel()
			}
			return
		}
		app.videoURL = v

		err := c.Promp("What do you want to do with this video", app.downloadButton, app.tweetButton)
		if err != nil {
			c.Cancel()
		}
	}
}

func (app *App) responseHandler(c *gram.Context) {
	var answer *gram.Button
	err := c.Get(&answer)
	if err != nil {
		c.Cancel()
		return
	}

	switch answer {
	case app.downloadButton:
		app.data.downloadVideos[c.id].URL = app.videoURL
		c.Switch(app.pipelines.downloadVideo)
	case app.tweetButton:
		c.Switch(app.pipelines.tweet)
	}
}

func isVideoURL(s string) bool {
	return false
}

func download(url, path string) error {
	return nil
}
