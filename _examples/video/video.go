package main

import "github.com/ilgooz/gram"

type DownloadVideo struct {
	URL string
}

type DownloadVideoHandler struct {
}

// askURLHandler tells telegram user to provide a video url to download.
func (h *DownloadVideoHandler) askURLHandler(c *gram.Context) {
	err := c.Send("Please provide a video url.")
	if err != nil {
		c.Cancel()
	}
}

// processVideoHandler called when the telegram user sends a message.
// if provided message doesn't qualify a valid url we simply send back an error message
// and ask again.
// To make things scalable gram.Context.Get calls must be separated into another step
// only called once per step.
func (h *DownloadVideoHandler) processVideoHandler(c *gram.Context) {
	err := c.Get(&d.URL)
	if err != nil {
		err = c.Send(unvalidURLError.Error())
		if err != nil {
			c.Cancel()
		}
		return
	}

	err = c.Send("Processing your video, this can take several minutes...")
	if err != nil {
		c.Cancel()
		return
	}

	err = d.download(d.URL)
	if err != nil {
		err = c.Send(serverError.Error())
		if err != nil {
			c.Cancel()
		}
	}
}
