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

func checkCode(url, method string, resp *http.Response, output chan<- string) {
	if resp.StatusCode == 200 {
		output <- fmt.Sprintf("%v Using %s >> %s", green("[+]"), method, resp.Status)
	}
}

func methodBypass(url string, client http.Client, output chan<- string) {
	useragent := GetRandomAgent()

	req, err := http.NewRequest("GET", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err := client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "GET", resp, output)

	req, err = http.NewRequest("POST", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Content-Length", "0")
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "POST", resp, output)

	req, err = http.NewRequest("PUT", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "PUT", resp, output)

	req, err = http.NewRequest("TRACE", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "TRACE", resp, output)

	req, err = http.NewRequest("OPTIONS", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "OPTIONS", resp, output)

	req, err = http.NewRequest("DELETE", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "DELETE", resp, output)

	req, err = http.NewRequest("HEAD", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "HEAD", resp, output)

	req, err = http.NewRequest("TRACK", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "TRACK", resp, output)

	req, err = http.NewRequest("CONNECT", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "CONNECT", resp, output)

	req, err = http.NewRequest("PATCH", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCode(url, "PATCH", resp, output)
}

func headerBypass(url string, client http.Client, output chan<- string) {
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

	client := http.Client{
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
