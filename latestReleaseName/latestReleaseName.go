package latestReleaseName

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ERR_NORELEASE = errors.New("This repository has no release.")

func GetLatestReleaseName(user string, repository string) (name string, err error) {
	targetUrl := "https://github.com/" + user + "/" + repository + "/releases/latest"
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	var resp *http.Response
	resp, err = c.Get(targetUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		err = ERR_NORELEASE
	} else if resp.StatusCode == 302 {
		redirectURL := resp.Header["Location"][0]
		urlArray := strings.Split(redirectURL, "/")
		name = urlArray[len(urlArray)-1]
	}
	return
}
