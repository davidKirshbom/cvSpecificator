package httpClient

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type HttpFetchOptions struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}
type HttpClient struct {
	Client *http.Client
}
var client *HttpClient
 
func New() HttpClient {
	client:= HttpClient{Client: &http.Client{}}
	return client
	
}
func (c *HttpClient) Fetch(options HttpFetchOptions,saveTo any){
	var req *http.Request
	var err error 
	if(options.Body!=""){
		req, err = http.NewRequest(options.Method, options.Url,strings.NewReader(options.Body))
	}else{
		req, err = http.NewRequest(options.Method, options.Url,nil)
	}
	if err != nil {
		panic(err)
	}
	for key, value := range options.Headers {
		req.Header.Set(key, value) 
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	json.Unmarshal(body,saveTo)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
}
func (c *HttpClient) Close() {
	c.Client.CloseIdleConnections()
}