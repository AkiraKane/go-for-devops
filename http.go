package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	// initialize the HTTP client
	clinet := &http.Client{}

	// define the url of the remote file
	url := "https://example.com/path/to/file"

	// create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// sent the GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// check for a successful response
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received status code", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return err
	}

	fmt.Println(string(body))
}