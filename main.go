package main

import (
	"bufio"
	"flag"
	"os"
	"sync"

	"github.com/FallAngel1337/gobypass/gobypass"
)

var (
	urls    string
	threads int
	wg      sync.WaitGroup
)

func main() {
	flag.StringVar(&urls, "u", "", "(Optional) The list containing the 403 urls")
	flag.IntVar(&threads, "t", 1, "The number of threads")
	flag.Parse()

	wg.Add(threads)

	if urls != "" {
		fn, err := os.Open(urls)
		gobypass.CheckError(err)

		reader := bufio.NewScanner(fn)
		urls := make([]string, 0)

		for reader.Scan() {
			urls = append(urls, reader.Text())
		}

		for i := 0; i < threads; i++ {
			go gobypass.Bypass(&urls, &wg)
		}
	} else {
		reader := bufio.NewScanner(os.Stdin)
		urls := make([]string, 0)

		for reader.Scan() {
			urls = append(urls, reader.Text())
		}

		for i := 0; i < threads; i++ {
			go gobypass.Bypass(&urls, &wg)
		}
	}

	wg.Wait()
}
