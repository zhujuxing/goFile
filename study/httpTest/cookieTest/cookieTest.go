package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GEt", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	cookie := &http.Cookie{Name: "userId", Value: strconv.Itoa(12345)}
	request.AddCookie(cookie)
	request.AddCookie(&http.Cookie{Name: "session", Value: "YWRtaW4="})
	response, err := client.Do(request)
	fmt.Println(response.Request.Cookies())

}
