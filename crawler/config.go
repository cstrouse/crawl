package crawler

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func FromJSON(in io.Reader) (*Crawler, error) {
	config := Crawler{
		Connections:     1,
		MaxDepth:        0,
		RobotsUserAgent: "Crawler",
		WaitTime:        "100ms",
	}

	configJSON, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
