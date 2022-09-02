package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GEt", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(request)
	fmt.Println(response.StatusCode)
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

}
