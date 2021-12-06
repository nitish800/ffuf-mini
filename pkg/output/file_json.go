package output

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/ffuf/ffuf/pkg/ffuf"
)

type ejsonFileOutput struct {
	CommandLine string        `json:"commandline"`
	Time        string        `json:"time"`
	Results     []ffuf.Result `json:"results"`
	Config      *ffuf.Config  `json:"config"`
}

type JsonResult struct {
	Input            map[string]string `json:"input"`
	Position         int               `json:"position"`
	StatusCode       int64             `json:"status"`
	ContentLength    int64             `json:"length"`
	ContentWords     int64             `json:"words"`
	ContentLines     int64             `json:"lines"`
	ContentType      string            `json:"content-type"`
	RedirectLocation string            `json:"redirectlocation"`
	Duration         time.Duration     `json:"duration"`
	ResultFile       string            `json:"resultfile"`
	Url              string            `json:"url"`
	Host             string            `json:"host"`
}

type jsonFileOutput struct {
	Results []string `json:"results"`
}

func writeEJSON(filename string, config *ffuf.Config, res []ffuf.Result) error {
	t := time.Now()
	outJSON := ejsonFileOutput{
		CommandLine: config.CommandLine,
		Time:        t.Format(time.RFC3339),
		Results:     res,
	}

	outBytes, err := json.Marshal(outJSON)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, outBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func writeJSON(filename string, config *ffuf.Config, res []ffuf.Result) error {
	jsonRes := []string{}
	for _, r := range res {
		strinput := make(map[string]string)
		for k, v := range r.Input {
			strinput[k] = string(v)
		}
		jsonRes = append(jsonRes, r.Url)
		jsonRes = append(jsonRes, "\n")
	}
	/*
		outJSON := jsonFileOutput{
			Results: jsonRes,
		}
		outBytes, err := json.Marshal(outJSON)
		if err != nil {
			return err
		}
	*/
	/*
		err = ioutil.WriteFile(filename, outBytes, 0644)
		if err != nil {
			return err
		}
	*/
	outBytes := strings.Join(jsonRes, "")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write([]byte(outBytes)); err != nil {
		return err
	}

	return nil
}
