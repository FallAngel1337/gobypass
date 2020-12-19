
![Text Alert](https://vsoch.github.io/assets/images/posts/learning-go/gophercises_jumping.gif)

# GoBypass
[![Build Status](https://travis-ci.org/dwyl/esta.svg?branch=master)](https://travis-ci.org/dwyl/esta)
[![GolangCI](https://golangci.com/badges/github.com/moul/golang-repo-template.svg)](https://golangci.com/r/github.com/moul/golang-repo-template)


# How to install?
**Note: Make sure you have go installed in your environment.**
**[Here's how you can install Go](https://golang.org/doc/install)**

## Using go get:

`$ go get -u -v github.com/FallAngel1337/gobypass`


# Options
| command | description | Ex |
| --- | --- |
| `-urls` | (Optional) The list containing the 403 urls | `gobypass -urls 403_urls.txt` |
| `-threads` | The number of threads (default 1) | `gobypass -urls 403_urls.txt -threads 5` |
| `-timeout` | The timeout for the connection in seconds (default 5) | `gobypass -urls 403_urls.txt -timeout 10` |
| `-v` |  Enable verbose mode | `gobypass -urls 403_urls.txt -v` |
| `cat urls_403.txt \| gobypass -v` | Read input from stdin |


# Found an error or a suggestion?
| contact | Twiiter |
| --- | --- |
| Fall |  @FallAngel10 |
