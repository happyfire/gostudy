package main

import (
	"fmt"
	"net/http"
	"os"
	//"io/ioutil"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprint(os.Stderr, "fecth: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)

		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: io.Copy %s:%v\n", url, err)
		} else {
			fmt.Printf("read %d bytes, http status=%s\n", n, resp.Status)
		}
	}
}
