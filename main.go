package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/FallAngel1337/gobypass/pkg/bypass"
)

var (
	urls    string
	threads int
	timeout int
	verbose bool

	wg, out sync.WaitGroup
)

func main() {
	flag.StringVar(&urls, "urls", "", "(Optional) The list containing the 403 urls")
	flag.IntVar(&threads, "threads", 1, "The number of threads")
	flag.IntVar(&timeout, "timeout", 5, "The timeout for the connection in seconds")
	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.Parse()

	bypass.BypassHeaders.LoadDefaultHeaders()
	bypass.SetTimeout(timeout)
	var reader *bufio.Scanner
	var bypassFunc func(urls <-chan string, output chan<- string, wg *sync.WaitGroup) // gambiarra

	if urls != "" {
		fn, _ := os.Open(urls)
		reader = bufio.NewScanner(fn)
	} else {
		reader = bufio.NewScanner(os.Stdin)
	}

	switch verbose {
	case true:
		bypassFunc = bypass.Verbose
	default:
		bypassFunc = bypass.Bypass
	}

	input := make(chan string)
	output := make(chan string)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go bypassFunc(input, output, &wg)
	}

	out.Add(1)
	go func() {
		defer out.Done()
		for o := range output {
			fmt.Println(o)
		}
	}()

	for reader.Scan() {
		input <- reader.Text()
	}

	close(input)
	wg.Wait()
	close(output)
	out.Wait()
}
