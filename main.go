package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	backoff "github.com/cenkalti/backoff/v4"
)

var target = "http://httpbin.org/post"

func relay(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	op := func() error {
		req, _ := http.NewRequest("POST", target, bytes.NewReader(b))
		req.Header = r.Header.Clone()
		resp, err := http.DefaultClient.Do(req)
		if err != nil { return err }
		defer resp.Body.Close()
		if resp.StatusCode >= 500 { return io.ErrUnexpectedEOF }
		return nil
	}
	bo := backoff.NewExponentialBackOff()
	if err := backoff.Retry(op, bo); err != nil {
		http.Error(w, err.Error(), 502); return
	}
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hooks", relay)
	log.Println("Webhook relay em :8085 -> POST /hooks (reenvia com retry ao destino)")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
