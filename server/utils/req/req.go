package req

import (
	"app/server/model"
	"app/server/utils"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string, headers map[string]string) (string, error) {
	return Request(url, "GET", nil, headers)
}

func Document(url string, source model.Source) (*goquery.Document, error) {
	var headers map[string]string
	if source.Header != "" {
		utils.ToObj(source.Header, &headers)
	}
	resp, err := Get(url, headers)
	if err != nil {
		return nil, err
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
