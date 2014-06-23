package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

var settingFile = ".yo.json"

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func jsonFile() string {
	usr, err := user.Current()
	assert(err)
	return strings.Join([]string{usr.HomeDir, settingFile}, "/")
}

func login() {

	requestToken, url, err := Consumer.GetRequestTokenAndUrl("")
	assert(err)

	fmt.Fprintf(os.Stderr, "1) open: %s\n", url)
	open.Run(url)

	fmt.Fprintf(os.Stderr, "2) Enter the PIN: ")
	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := Consumer.AuthorizeToken(requestToken, verificationCode)
	json, err := json.Marshal(accessToken)
	assert(err)

	ioutil.WriteFile(jsonFile(), json, 0600)
}

func main() {

	user := ""
	if len(os.Args) == 2 {
		user = os.Args[1]
		user = strings.TrimPrefix(user, "@")
		user = "@" + user
	}

	if _, err := os.Stat(jsonFile()); os.IsNotExist(err) {
		login()
	}

	file, err := os.Open(jsonFile())
	assert(err)

	accessToken := &oauth.AccessToken{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		j := scanner.Text()
		err := json.Unmarshal([]byte(j), accessToken)
		assert(err)
	}

	status := fmt.Sprintf(user + " Yo")
	response, err := Consumer.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{
			"status": status,
		},
		accessToken)

	assert(err)
	debug(response)
}
