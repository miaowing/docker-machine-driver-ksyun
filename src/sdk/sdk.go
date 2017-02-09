package sdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	loginUrl = "/login"
)

var Host string = "http://localhost"

func getHost() string {
	return Host
}

func Login(username string, password string) (*http.Response, error) {
	client := http.Client{}

	body := map[string]string{username: username, password: password}
	bodyByte, _ := json.Marshal(body)

	request, _ := http.NewRequest("GET", getHost()+loginUrl,
		strings.NewReader(bytes.NewBuffer(bodyByte).String()))
	request.Header.Add("Content-Type", "application/json")

	return client.Do(request)
}

func CreateInstance() {

}

func StopInstance() {

}

func StartInstance() {

}

func RestartInstance() {
	
}