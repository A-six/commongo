package httpd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RestHeader(url, method string, ma map[string]string, by []byte) ([]byte, error) {
	body := bytes.NewBuffer(by)
	client := &http.Client{}
	//解决x509: certificate signed by unknown authority
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr
	//解决x509: certificate signed by unknown authority

	reqest, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

	for k, v := range ma {
		reqest.Header.Set(k, v)
	}

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return []byte{}, err
	}
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("StatusCode %d", response.StatusCode)
	}
	ret, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	err = response.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return ret, nil
}
