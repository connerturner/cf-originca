package main

type CertificateList struct {
	Success    bool                    `json:"success,omitempty"`
	Errors     []string                `json:"errors,omitempty"`
	Messages   []string                `json:"messages,omitempty"`
	Result     []CertificateListResult `json:"result,omitempty"`
	ResultInfo []string                `json:"-"`
}

type CertificateListResult struct {
	Id              string   `json:"id,omitempty"`
	Certificate     string   `json:"certificate,omitempty"`
	Hostnames       []string `json:"hostnames,omitempty"`
	Expires         string   `json:"expires_on,omitempty"`
	RequestType     string   `json:"request_type,omitempty"`
	RequestValidity int      `json:"requested_validity,omitempty"`
	Csr             string   `json:"csr,omitempty"`
}
