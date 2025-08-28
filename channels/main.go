package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	listOfUrls := []string{}
	listOfUrls = append(listOfUrls, "http://google.com")
	listOfUrls = append(listOfUrls, "http://facebook.com")
	listOfUrls = append(listOfUrls, "http://amazon.com")
	listOfUrls = append(listOfUrls, "http://stackoverflow.com")
	// fmt.Println(listOfUrls)

	c := make(chan string)

	for _, url := range listOfUrls {
		go sendHttpRequest(url, c)
		// time.Sleep(1 * time.Second)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			sendHttpRequest(l, c)
		}(l)
	}
}

func sendHttpRequest(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		c <- url
		return
	}

	fmt.Println(url, "is up")
	c <- url
}
