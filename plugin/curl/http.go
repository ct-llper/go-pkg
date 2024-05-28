package curl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func HttpJson(method string, url string, params []byte) ([]byte, error) {
	body := bytes.NewBuffer(params)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode == http.StatusOK {
		msg := fmt.Sprintf("response status code %v", resp.StatusCode)
		return nil, errors.New(msg)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func HttpDo(method string, url string, values url.Values) ([]byte, error) {
	body := strings.NewReader(values.Encode())

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", cookie)
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Add("x-requested-with", "XMLHttpRequest") //AJAX

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode == http.StatusOK {
		msg := fmt.Sprintf("response status code %v", resp.StatusCode)
		return nil, errors.New(msg)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func HttpGet(url string) ([]byte, error) {
	//发送请求获取响应
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//结束网络释放资源
	if resp != nil {
		defer resp.Body.Close()
	}
	//判断响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	//读取响应实体
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func HttpPost(url string, params url.Values, contentType string) ([]byte, error) {

	body := strings.NewReader(params.Encode())

	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func HttpPostForm(url string, values url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func MakeParams(params url.Values) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bb := bytes.Buffer{}
	for _, key := range keys {
		val := params.Get(key)
		bb.WriteString(key)
		bb.WriteString("=")
		bb.WriteString(val)
		bb.WriteString("&")
	}

	return strings.TrimRight(bb.String(), "&")
}

// https://blog.csdn.net/jined/category_10393849.html
