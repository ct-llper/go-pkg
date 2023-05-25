package curl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func RequestClient(url, code, contentType string, in []byte) (err error, out string) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(in))
	req.Header.Set("code", code)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, out
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return err, string(body)
}

func RequestHttpJson(url, method string, in []byte) (err error, out string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(in)))
	if err != nil {
		return err, out
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err, out
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	return err, string(body)
}

func PostHttpJson(url string, in interface{}) (err error, out string) {
	jsonByte := make([]byte, 0, 0)

	if str, ok := in.(string); ok {
		jsonByte = []byte(str)
	} else {
		jsonByte, err = json.Marshal(in)
		if err != nil {
			return err, out
		}
	}

	req := bytes.NewBuffer(jsonByte)
	resp, err := http.Post(url, "application/json;charset=utf-8", req)
	if err != nil {
		return err, out
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return err, string(body)
}
