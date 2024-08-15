package facecloud

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	params map[string]string
)

func init() {
	params = make(map[string]string)
	params["fd_min_size"] = "0"
	params["fd_max_size"] = "0"
	params["fd_threshold"] = "0.8"
	params["rotate_until_faces_found"] = "false"
	params["orientation_classifier"] = "false"
	params["demographics"] = "true"
	//params["attributes"] = "true"
	//params["landmarks"] = "false"
	//params["liveness"] = "false"
	//params["quality"] = "false"
	//params["masks"] = "false"
}

func Login() (response *ResponseLogin, err error) {

	requestBody := RequestLogin{
		Email:    email,
		Password: password,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		err = errors.Wrapf(err, "failed to marshal request body")
		return
	}

	req, err := http.NewRequest("POST", loginEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		err = errors.Wrapf(err, "failed to make http request")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrapf(err, "failed to send http request")
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrapf(err, "failed to read response body")
		return
	}

	if err = json.Unmarshal(responseBody, &response); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal response")
		return
	}

	if response.StatusCode != 200 {
		err = errors.New(response.Message)
		return
	}

	return
}

func Detect(imagePath string) (response *ResponseDetect, err error) {

	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		err = errors.Wrapf(err, "failed to read image")
		return
	}

	var requestBody bytes.Buffer
	requestBody.Write(imageBytes)

	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Add(k, v)
	}
	endpoint, err := url.Parse(detectEndpoint)
	if err != nil {
		err = errors.Wrapf(err, "failed to parse detect endpoint")
		return
	}
	endpoint.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("POST", endpoint.String(), &requestBody)
	if err != nil {
		err = errors.Wrapf(err, "failed to make http request")
		return
	}

	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrapf(err, "failed to send http request")
		return
	}
	defer resp.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, resp.Body)
	if err != nil {
		err = errors.Wrapf(err, "failed to read response body")
		return
	}

	if err = json.Unmarshal(responseBody.Bytes(), &response); err != nil {
		err = errors.Wrapf(err, "failed to unmarshal response")
		return
	}

	if response.StatusCode != 200 {
		err = errors.New(response.Message)
		return
	}

	return
}
