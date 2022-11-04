package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
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
	certToFile := flag.Bool("cert-to-file", false, "Output each certificate into a file in the current directory, Only used when listing certificates")
	flag.Parse()

	switch cmd := flag.Arg(0); cmd {
	case "create":
		create()
	case "revoke":
		revoke()
	case "get":
		get()
	case "list":
		list(*ocaKey, *zone, *certToFile)
	default:
		fmt.Printf("Unknown command: %s\nUsage:\n", cmd)
		fmt.Println(usage)
		flag.PrintDefaults()
	}

}

func create() {
	panic("unimplemented")
}

func get() {
	panic("unimplemented")
}

func revoke() {
	panic("unimplemented")
}

func list(ocaKey string, zone string, certToFile bool) {

	if ocaKey == "" || zone == "" {
		fmt.Printf("Both oca-key and zone are needed for this operation\n")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, baseApi+"certificates", nil)
	if err != nil {
		fmt.Printf("Unable to create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Add("X-Auth-User-Service-Key", ocaKey)
	query := req.URL.Query()
	query.Add("zone_id", zone)
	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in http request: %s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error: %d\n", response.StatusCode)
		responseBytes, _ := io.ReadAll(response.Body)
		fmt.Println(string(responseBytes))
		os.Exit(1)
	}

	responseData, _ := io.ReadAll(response.Body)

	var certList CertificateList

	e := json.Unmarshal(responseData, &certList)

	if e != nil {
		fmt.Printf("Error unmarhsalling json: \n%s\n", e)
	}

	if certToFile {
		CertsToFile(certList)
	} else {
		fmt.Printf("%+v\n", certList)
	}

}

func CertsToFile(cl CertificateList) {
	for _, cert := range cl.Result {
		file, err := os.Create(cert.Id + ".pem")
		if err != nil {
			fmt.Printf("Error Writing Certificates to File: \n %s", err)
		}

		defer file.Close()

		_, error := file.WriteString(cert.Certificate)
		if error != nil {
			fmt.Printf("Error Writing Certificates to File: \n %s", error)
		}
	}
}
