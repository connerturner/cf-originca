package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

var usage = `
originca [COMMAND] [OPTIONS]
COMMAND := list,create,get,revoke

OPTIONS:
`

func main() {
	ocaKey := flag.String("oca-key", "", "Cloudflare `Origin CA Key`, needed to authorise any CA Certificate operations")
	zone := flag.String("zone", "", "Cloudflare `Zone ID`, Only used when listing certificates")
	flag.Parse()

	switch cmd := flag.Arg(0); cmd {
	case "create":
		create()
	case "list":
		list(*ocaKey, *zone)
	default:
		fmt.Printf("Unknown command: %s\nUsage:\n", cmd)
		fmt.Println(usage)
		flag.PrintDefaults()
	}

}

func create() {
}

func list(ocaKey string, zone string) {

}
