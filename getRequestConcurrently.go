package httpresp

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Body       string
	StatusCode int
}

func BasicOP(urls []string) []Response {
	var responseList []Response
	wc := 2
	var wg sync.WaitGroup
	inputChan := make(chan string)
	var collective_resp []<-chan Response
	for i := range wc {
		fmt.Println(i)
		wg.Add(1)
		collective_resp = append(collective_resp, hitURl(inputChan, &wg))
	}

	go func() {
		defer close(inputChan)
		for _, url := range urls {
			inputChan <- url
		}
	}()

	for _, resp := range collective_resp {
		go func(resp <-chan Response) {
			for response := range resp {
				responseList = append(responseList, response)
			}
		}(resp)

	}

	wg.Wait()
	return responseList
}

func hitURl(inputChan chan string, wg *sync.WaitGroup) <-chan Response {
	response := make(chan Response)
	go func() {
		defer close(response)
		defer wg.Done()
		for url := range inputChan {
			client := http.Client{Timeout: 5 * time.Second}
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("eeror received")
				fmt.Println(err.Error())
				response <- Response{Body: err.Error(), StatusCode: resp.StatusCode}
			}
			result, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			response <- Response{Body: string(result), StatusCode: resp.StatusCode}
		}

	}()
	return response
}
