package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func main() {
	MakeRequest()
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func MakeRequest() {
	name := os.Args[1]

	resp, err := http.Get("https://mobile.facebook.com/gaming/" + name)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	re2, _ := regexp.Compile("id=[0-9]+")
	match3 := re2.FindString(string(body))

	res1 := strings.Split(match3, "=")
	linkURI := "https://strims.gg/facebook/" + string(res1[1])
	fmt.Println("Opening stream on strims: ", linkURI)
	openBrowser(linkURI)

}
