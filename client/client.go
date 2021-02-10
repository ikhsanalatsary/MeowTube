package client

import "net/http"

var CookieJar = NewJar()

// Request....
var Request = &http.Client{Jar: CookieJar}
