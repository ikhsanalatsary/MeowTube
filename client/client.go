package client

import (
	"net/http"
	"net/url"
)

var CookieJar = NewJar()

// Request is an instance of http.Client that include CookieJar
var Request = &http.Client{Jar: CookieJar}

// Fetch method that implement client structure
func Fetch(baseURL string) (*http.Response, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}
	var host, origin string = u.Host, u.Scheme + "://" + u.Hostname()
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0")
	req.Header.Add("Origin", origin)
	req.Header.Add("Host", host)
	req.Header.Add("Accept", "*/*")

	return Request.Do(req)
}
