package spidercore

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRequest(request *http.Request) (string, error) {
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

func Request(reqType string, url string) string {

	if reqType == "GET" {
		request, err := BuildGetHeaders(url)
		if err != nil {
			fmt.Println("GET request error", err)
		}

		content, err := getRequest(request)
		if err != nil {
			fmt.Println("getRequest error", err)
		}

		return content
	}

	return ""

}