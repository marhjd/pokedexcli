package main

import (
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Minute * 1
	StartRepl()
}
