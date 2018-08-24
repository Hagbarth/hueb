package ifttt

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const triggerURL = "https://maker.ifttt.com/trigger"

type Client struct {
	key string
}

func NewClient(key string) *Client {
	return &Client{key}
}

func (c *Client) TriggerEvent(event string) error {
	url := fmt.Sprintf("%s/%s/with/key/%s", triggerURL, event, c.key)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error when calling the trigger api: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("could not read response body: %v", err)
		}
		return fmt.Errorf("http error when calling the trigger api - code: %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}
