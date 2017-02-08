package sdk

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLogin(t *testing.T) {
	Host = "http://10.160.57.80:8080"
	response, err := Login("admin", "openstack1")

	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
