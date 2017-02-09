package sdk

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLogin(t *testing.T) {
	Host = "http://localhost"
	response, err := Login("admin", "admin")

	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}