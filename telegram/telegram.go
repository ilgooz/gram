package telegram

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var getUpdatesURL = "https://api.telegram.org/bot%s/getUpdates"

type Telegram struct {
	key string
}

func New(key string) *Telegram {
	return &Telegram{
		key: key,
	}
}

func (t *Telegram) GetUpdates(offset int) (updates []Update, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	form := url.Values{
		"offset": []string{fmt.Sprintf("%d", offset)},
	}
	response, err := client.PostForm(fmt.Sprintf(getUpdatesURL, t.key), form)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var updateRes UpdateResponse
	err = json.NewDecoder(response.Body).Decode(&updateRes)
	if err != nil {
		return nil, err
	}

	return updateRes.Result, err
}
