package main

// Sources:
//   https://gist.github.com/jniltinho/9787946
//   https://gowebexamples.com/http-server/

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func getIpv4Addresses() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}

	var addressStrings []string

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addressStrings = append(addressStrings, ipnet.IP.String())
			}
		}
	}
	return strings.Join(addressStrings, ", ")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getIpv4Addresses()+"\n")
	})

	http.ListenAndServe(":80", nil)
}
