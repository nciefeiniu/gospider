package spidercore

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//
func GetRequest(request *http.Request) (string, error) {
	client := &http.Client{}

	res, err := client.Do(request)

	if err != nil {
		fmt.Println("client.Do error")
		return "", err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("ioutil.ReadAll error")
		return "", err
	}

	return string(content), nil
}