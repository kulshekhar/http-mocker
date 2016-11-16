package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(getAddr(), nil)
}

func getAddr() string {
	port := os.Getenv("INTERCEPTOR_PORT")
	if port == "" {
		port = "18000"
	}
	host := os.Getenv("INTERCEPTOR_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	return fmt.Sprintf(`%s:%s`, host, port)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/_script.js" {
		w.Header().Set("Content-Type", "application/javascript")
		script, _ := Asset("script/script.min.js")
		fmt.Fprint(w, string(script))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")

	contentType := strings.ToLower(r.Header.Get("content-type"))

	switch {

	case strings.Index(contentType, "application/json") >= 0:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
			return
		}

		var resp response
		if err := json.Unmarshal(body, &resp); err != nil {
			fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
			return
		}
		if resp.StatusCode == 0 {
			resp.StatusCode = 200
		}

		w.WriteHeader(resp.StatusCode)
		result, _ := json.Marshal(resp.Response)
		w.Write(result)

	case strings.Index(contentType, "text/plain") >= 0 || strings.Index(contentType, "application/x-www-form-urlencoded") >= 0 || contentType == "":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
			return
		}

		if statusCode, err := strconv.Atoi(r.PostFormValue("statusCode")); err != nil {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(statusCode)
		}

		fmt.Fprint(w, r.PostFormValue("response"))
	default:
		fmt.Fprintf(w, `{"error":"invalid content type - %s"}`, contentType)
	}

}

type response struct {
	Response   interface{} `json:"response"`
	StatusCode int         `json:"statusCode"`
}
