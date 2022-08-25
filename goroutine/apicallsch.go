package goroutine

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"sync"
)

func sendUser(user string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("URL/" + user)
	if err != nil {
		fmt.Println("err handle it")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err handle it")
	}
	ch <- string(b)
}

func AsyncHTTP(users []string) ([]string, error) {
	ch := make(chan string)
	var responses []string
	var user string
	var wg sync.WaitGroup
	for _, user = range users {
		wg.Add(1)
		go sendUser(user, ch, &wg)
	}

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()
	// read from channel as they come in until its closed
	for res := range ch {
		responses = append(responses, res)
	}

	return responses, nil
}
