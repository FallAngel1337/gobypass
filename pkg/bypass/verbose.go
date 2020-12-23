package bypass

import (
	"crypto/tls"
	"fmt"
	"net/http"
	parser "net/url"
	"sync"
	"time"

	"github.com/FallAngel1337/gobypass/pkg/errors"
	"github.com/gookit/color"
)

func createNewRequestV(url, method, ua string, client *http.Client, output chan<- string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		errors.VerifyReport(err)
	}
	req.Header.Add("User-Agent", ua)
	resp, err := client.Do(req)
	if err != nil {
		errors.VerifyReport(err)
	}

	if resp.StatusCode == 200 {
		output <- fmt.Sprintf("%v Using %s >> %s", green("[+]"), method, resp.Status)
	} else if resp.StatusCode != 200 && resp.StatusCode != 403 {
		output <- fmt.Sprintf("%v Using %s >> %s", yellow("[!]"), method, resp.Status)
	} else {
		output <- fmt.Sprintf("%v Using %s >> %s", red("[-]"), method, resp.Status)
	}
}

func methodBypassV(url string, client *http.Client, output chan<- string) {
	useragent := GetRandomAgent()

	createNewRequestV(url, "GET", useragent, client, output)
	createNewRequestV(url, "POST", useragent, client, output)
	createNewRequestV(url, "PUT", useragent, client, output)
	createNewRequestV(url, "TRACE", useragent, client, output)
	createNewRequestV(url, "OPTIONS", useragent, client, output)
	createNewRequestV(url, "DELETE", useragent, client, output)
	createNewRequestV(url, "HEAD", useragent, client, output)
	createNewRequestV(url, "TRACK", useragent, client, output)
	createNewRequestV(url, "CONNECT", useragent, client, output)
	createNewRequestV(url, "PATCH", useragent, client, output)
}

func headerBypassV(url string, client *http.Client, output chan<- string) {
	req, _ := http.NewRequest("GET", url, nil)
	useragent := GetRandomAgent()
	req.Header.Add("User-Agent", useragent)
	for name, value := range BypassHeaders.Header {
		req.Header.Add(name, value)
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 {
			output <- fmt.Sprintf("%v Using %s: %s >> %s", green("[+]"), name, value, resp.Status)
		} else if resp.StatusCode != 200 && resp.StatusCode != 403 {
			output <- fmt.Sprintf("%v Using %s: %s >> %s", yellow("[!]"), name, value, resp.Status)
		} else {
			output <- fmt.Sprintf("%v Using %s: %s >> %s", red("[-]"), name, value, resp.Status)
		}
	}
}

// Verbose is the same function as Bypass but with verbose
func Verbose(urls <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * timeout,
	}

	for url := range urls {
		color.Cyan.Printf("[#] Checking %s\n\n", url)
		parsedURL, _ := parser.Parse(url)
		code, _ := http.Get(url) // First check if the code is 403

		if code.StatusCode == 403 {
			fmt.Println(yellow("[!]"), "Found a 403!")
			// Some custom headers that acts dinnamicly
			BypassHeaders.AddHeader("Referer", url)
			BypassHeaders.AddHeader("X-Origial-URL", parsedURL.Path)

			fmt.Printf("%v Trying changing methods...\n\n", blue("[*]"))
			methodBypassV(url, client, output)

			time.Sleep(time.Second * 1)

			fmt.Printf("\n%v Trying using headers...\n\n", blue("[*]"))
			headerBypassV(url, client, output)
		} else {
			fmt.Println(red("[-]"), "No 403 found!")
		}
	}
}
