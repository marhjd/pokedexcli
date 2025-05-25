package main

import (
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Second * 5
	StartRepl()
}
