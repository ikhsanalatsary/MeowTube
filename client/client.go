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
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("Origin", origin)
	req.Header.Set("Host", host)
	req.Header.Set("Accept", "*/*")

	return Request.Do(req)
}
