package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"

	gobypass "github.com/FallAngel1337/gobypass/pkg"
)

var (
	urls    string
	threads int
	wg, out sync.WaitGroup
)

func main() {
	flag.StringVar(&urls, "u", "", "(Optional) The list containing the 403 urls")
	flag.IntVar(&threads, "t", 1, "The number of threads")
	flag.Parse()

	inputchan := make(chan string)
	outputchan := make(chan string)

	if urls != "" {
		fn, err := os.Open(urls)
		gobypass.CheckError(err)

		reader := bufio.NewScanner(fn)
		wg.Add(threads)
		for i := 0; i < threads; i++ {
			go gobypass.Bypass(inputchan, outputchan, &wg)
		}

		out.Add(1)
		go func() {
			defer out.Done()
			for o := range outputchan {
				fmt.Println(o)
			}
		}()

		for reader.Scan() {
			inputchan <- reader.Text()
		}
	} else {
		reader := bufio.NewScanner(os.Stdin)

		wg.Add(threads)
		for i := 0; i < threads; i++ {
			go gobypass.Bypass(inputchan, outputchan, &wg)
		}

		out.Add(1)
		go func() {
			defer out.Done()
			for o := range outputchan {
				fmt.Println(o)
			}
		}()

		for reader.Scan() {
			inputchan <- reader.Text()
		}
	}
	close(inputchan)
	wg.Wait()
	close(outputchan)
	out.Wait()
}
