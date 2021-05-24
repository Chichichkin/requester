package internal

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MakeRequest(urlStr string) (md5Result string, err error) {
	url, err := checkUrl(urlStr)
	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Convert the body to type string
	h := md5.New()
	h.Write(body)
	md5Result = hex.EncodeToString(h.Sum(nil))
	return md5Result, nil
}

func checkUrl(urlStr string) (*url.URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u, nil
}
