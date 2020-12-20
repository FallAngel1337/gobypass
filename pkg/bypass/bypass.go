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

var (
	timeout time.Duration
	green   = color.FgGreen.Render
	red     = color.FgRed.Render
	yellow  = color.FgYellow.Render
	blue    = color.FgBlue.Render

	// BypassHeaders contains the Headers struct
	BypassHeaders Headers
)

// SetTimeout defines the timeout in seconds
func SetTimeout(seconds int) {
	timeout = time.Second * time.Duration(seconds)
}

func createNewRequest(url, method, ua string, client *http.Client, output chan<- string) {
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
	}
}

func methodBypass(url string, client *http.Client, output chan<- string) {
	useragent := GetRandomAgent()

	createNewRequest(url, "GET", useragent, client, output)
	createNewRequest(url, "POST", useragent, client, output)
	createNewRequest(url, "PUT", useragent, client, output)
	createNewRequest(url, "TRACE", useragent, client, output)
	createNewRequest(url, "OPTIONS", useragent, client, output)
	createNewRequest(url, "DELETE", useragent, client, output)
	createNewRequest(url, "HEAD", useragent, client, output)
	createNewRequest(url, "TRACK", useragent, client, output)
	createNewRequest(url, "CONNECT", useragent, client, output)
	createNewRequest(url, "PATCH", useragent, client, output)
}

func headerBypass(url string, client *http.Client, output chan<- string) {
	req, _ := http.NewRequest("GET", url, nil)
	useragent := GetRandomAgent()
	req.Header.Add("User-Agent", useragent)
	for name, value := range BypassHeaders.Header {
		req.Header.Add(name, value)
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 {
			output <- fmt.Sprintf("%v Using %s: %s >> %s", green("[+]"), name, value, resp.Status)
		}
	}
}

//Bypass try bypass the restrictions
func Bypass(urls <-chan string, output chan<- string, wg *sync.WaitGroup) {
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
			// Some custom headers that acts dinnamicly
			fmt.Println(yellow("[!]"), "Found a 403!")
			BypassHeaders.AddHeader("Referer", url)
			BypassHeaders.AddHeader("X-Origial-URL", parsedURL.Path)

			fmt.Printf("%v Trying changing methods...\n\n", blue("[*]"))
			methodBypass(url, client, output)

			time.Sleep(time.Second * 1)

			fmt.Printf("\n%v Trying using headers...\n\n", blue("[*]"))
			headerBypass(url, client, output)
		} else {
			fmt.Println(red("[-]"), "No 403 found!")
		}
	}
}
