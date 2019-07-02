package request

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type Service interface {
	SendRequest(values map[string][]string) (response *http.Response, err error)
}

type service struct {
	apiToken string
	apiURL   string
	timeOut  time.Duration
}

// 发送请求
func (s *service) SendRequest(values map[string][]string) (response *http.Response, err error) {
	val := url.Values(values)
	postDataStr := val.Encode()
	header := map[string]string{
		"Authorization": "APIKEY " + s.apiToken,
	}
	response, err = doRequest(s.apiURL, header, "POST", []byte(postDataStr), s.timeOut, "")
	return
}

func NewService(apiToken string, apiUrl string, timeOut time.Duration) Service {
	return &service{apiToken: apiToken, apiURL: apiUrl, timeOut: timeOut}
}

func doRequest(targetUrl string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration, proxy string) (*http.Response, error) {
	//timeOut = time.Duration(timeOut * time.Millisecond)
	//urli := url.URL{}
	//urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
	//https认证
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		//Proxy:http.ProxyURL(urlproxy),

	}
	if proxy != "" {
		urli := url.URL{}
		urlProxy, _ := urli.Parse(proxy)
		tr.Proxy = http.ProxyURL(urlProxy)
	}
	client := http.Client{
		Timeout:   timeOut,
		Transport: tr,
	}

	client.Jar, _ = cookiejar.New(nil)

	method = strings.ToUpper(method)
	var req *http.Request
	var err error

	if postData != nil && (method == "POST" || method == "PUT") {
		//print(string(postData))

		req, err = http.NewRequest(method, targetUrl, bytes.NewReader(postData))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	} else {
		req, err = http.NewRequest(method, targetUrl, nil)
		if err != nil {
			return nil, err
		}

	}

	for key, value := range headerMap {
		req.Header.Add(key, value)
	}
	res, err := client.Do(req)
	client.CloseIdleConnections()
	return res, err
}
