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

func checkCodeV(url, method string, resp *http.Response, output chan<- string) {
	if resp.StatusCode == 200 {
		output <- fmt.Sprintf("%v Using %s >> %s", green("[+]"), method, resp.Status)
	} else if resp.StatusCode != 200 && resp.StatusCode != 403 {
		output <- fmt.Sprintf("%v Using %s >> %s", yellow("[!]"), method, resp.Status)
	} else {
		output <- fmt.Sprintf("%v Using %s >> %s", red("[-]"), method, resp.Status)
	}
}

func methodBypassV(url string, client http.Client, output chan<- string) {
	useragent := GetRandomAgent()

	req, err := http.NewRequest("GET", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err := client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "GET", resp, output)

	req, err = http.NewRequest("POST", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Content-Length", "0")
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "POST", resp, output)

	req, err = http.NewRequest("PUT", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "PUT", resp, output)

	req, err = http.NewRequest("TRACE", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "TRACE", resp, output)

	req, err = http.NewRequest("OPTIONS", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "OPTIONS", resp, output)

	req, err = http.NewRequest("DELETE", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "DELETE", resp, output)

	req, err = http.NewRequest("HEAD", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "HEAD", resp, output)

	req, err = http.NewRequest("TRACK", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "TRACK", resp, output)

	req, err = http.NewRequest("CONNECT", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "CONNECT", resp, output)

	req, err = http.NewRequest("PATCH", url, nil)
	errors.VerifyReport(err)
	req.Header.Add("User-Agent", useragent)
	resp, err = client.Do(req)
	errors.VerifyReport(err)
	checkCodeV(url, "PATCH", resp, output)
}

func headerBypassV(url string, client http.Client, output chan<- string) {
	req, _ := http.NewRequest("GET", url, nil)
	useragent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36"
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
			methodBypassV(url, client, output)

			time.Sleep(time.Second * 1)

			fmt.Printf("\n%v Trying using headers...\n\n", blue("[*]"))
			headerBypassV(url, client, output)

		}
	}
}
