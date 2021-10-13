package utils

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (string, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Edg/94.0.992.38",
	}
	return Request(url, "GET", nil, headers)
}

func Document(url string) (*goquery.Document, error) {
	resp, err := Get(url)
	if err != nil {
		return nil,err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func Request(url, method string, body io.Reader, headers map[string]string) (res string, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	if headers != nil {
		for _, s := range headers {
			req.Header.Set(s, headers[s])
		}
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res = string(bytes)
	return
}
