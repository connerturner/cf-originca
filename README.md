## OriginCA Tool

Small cross-platform Go utility to work with CF OriginCA certificates

### Install

Build the binary, or download from Releases

````
git clone ...
cd cf-originca
go build -o bin/originca
chmod +x bin/originca
./bin/originca
````

### Usage Information
````
Usage:

originca [COMMAND] [OPTIONS]
COMMAND := list,create,get,revoke

OPTIONS:

  -cert-to-file
    	Output each certificate into a file in the current directory, Only used when listing certificates
  -oca-key Origin CA Key
    	Cloudflare Origin CA Key, needed to authorise any CA Certificate operations
  -zone Zone ID
    	Cloudflare Zone ID, Only used when listing certificates
````