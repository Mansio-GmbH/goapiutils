package requestdetails

type RequestDetails struct {
	Headers    map[string]string `json:"headers"`
	DomainName string            `json:"domainName"`
}

type WithRequestDetails struct {
	Request *RequestDetails `json:"request"`
}
