package latestReleaseName

import (
	"errors"
	"net/http"
	"strings"
)

var ERR_NORELEASE = errors.New("This repository has no release.")
var ERR_NOTFOUND = errors.New("Not Found.")

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

	//	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		err = ERR_NOTFOUND
	} else if resp.StatusCode == 302 {
		redirectURL := resp.Header["Location"][0]
		urlArray := strings.Split(redirectURL, "/")

		if urlArray[len(urlArray)-1] == "releases" &&
			urlArray[len(urlArray)-2] == repository &&
			urlArray[len(urlArray)-3] == user {
			// this repository doesn't have any release
			err = ERR_NORELEASE
		} else {
			name = urlArray[len(urlArray)-1]
		}
	}
	return
}
