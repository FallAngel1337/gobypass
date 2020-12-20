package bypass

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
To-Do:

[  ] More Headers
[  ] More Techniques
[  ] ...
*/

// Headers struct contains the header structure used in the bypass
type Headers struct {
	Header map[string]string
}

// LoadDefaultHeaders initialize the Headers struct with some default headers
func (h *Headers) LoadDefaultHeaders() {
	header := make(map[string]string)

	header["Forwarded-For-Ip"] = "127.0.0.1"
	header["X-Forwarded-For"] = "127.0.0.1"
	header["Forwarded-For"] = "127.0.0.1"
	header["Forwarded"] = "127.0.0.1"
	header["X-Forwarded-For-Original"] = "127.0.0.1"
	header["X-Forwarded-By"] = "127.0.0.1"
	header["X-Forwarded"] = "127.0.0.1"
	header["X-Real-IP"] = "127.0.0.1"
	header["X-Custom-IP-Authorization"] = "127.0.0.1"
	header["Host"] = "example.com"

	h.Header = header
}

// AddHeader adds a custom header to the Headers struct
func (h *Headers) AddHeader(name, value string) {
	h.Header[name] = value
}

// GetRandomAgent return a string with a random user-agent
func GetRandomAgent() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	file, err := filepath.Abs(os.Getenv("GOPATH") + "/src/github.com/FallAngel1337/gobypass/user-agents/user-agents.txt")
	if err != nil {
		log.Fatal(err)
	}

	useragents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	useragent := strings.Split(string(useragents), "\n")[random.Intn(150)]
	useragent = useragent[:len(useragent)-1]
	return useragent
}
