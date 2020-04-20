package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	getDemo("http://google.com")
}

func getDemo(url string) int{
	resp,err:=http.Get(url)
	if err!=nil {
		fmt.Println("There is an error in the response :",err)
		return resp.StatusCode
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error in IO")
			return 0
		}
		bodyString := string(bodyBytes)
		fmt.Println("Response Body : ", bodyString)
	}
	return resp.StatusCode
}
