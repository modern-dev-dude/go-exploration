package main

import(
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"io"
	"os"
)

func getPage (res http.ResponseWriter, req *http.Request){
	fmt.Printf("got / request\n")

	requestDump, err := httputil.DumpRequest(req, true)
	fmt.Println(string(requestDump))
	
	if err != nil {
  	fmt.Println(err)
	}

	template := "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\" /><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" /><title>Document</title></head><body><p>Some text</p></body></html>"

	io.WriteString(res, template)
}

func main() {
	http.HandleFunc("/", getPage)


	if errors.Is(err, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
