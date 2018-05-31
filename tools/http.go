package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"unsafe"
)

// PostJSONBytes PostJSONBytes
func PostJSONBytes(URL string, jsonBytes []byte) (string, error) {

	reader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", URL, reader)
	if err != nil {
		return err.Error(), err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err.Error(), err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error(), err
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	return (*str), nil
}

// HTTPGet HTTP Get
func HTTPGet(URL string) (string, error) {

	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
