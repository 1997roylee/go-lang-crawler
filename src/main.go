package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	Domain string
	Path   string
	Method string
}

var _config Config
var _header map[string]string

func SetConfig(config Config) {
	_config = config
}

func SetHeader(header map[string]string) {
	_header = header
}

func LogErrors(errors error) {
	if errors != nil {
		log.Fatalln(errors)
	}
}

func ParseHTML(res *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	LogErrors(err)
	return doc
}

// func CheckMethod(method string) bool {
// 	switch method {
// 	case "GET":
// 		return true
// 	case "POST":
// 		return true
// 	default:
// 		return false
// 	}
// }

func Request(parameter map[string]string) *goquery.Document {
	var req *http.Request
	var _ error
	client := &http.Client{}
	url := _config.Domain + _config.Path
	fmt.Println("Url:", url, "Method", _config.Method)

	switch _config.Method {
	case "GET":
		req, _ = http.NewRequest(_config.Method, url, nil)
	case "POST":
		jsonValue, _ := json.Marshal(parameter)
		req, _ = http.NewRequest(_config.Method, url, bytes.NewBuffer(jsonValue))
	case "PATCH":
		jsonValue, _ := json.Marshal(parameter)
		req, _ = http.NewRequest(_config.Method, url, bytes.NewBuffer(jsonValue))
	default:
	}

	for key, element := range _header {
		req.Header.Set(key, element)
	}

	res, err := client.Do(req)
	LogErrors(err)

	fmt.Println("response Status Code:", res.StatusCode)
	fmt.Println("response Headers:", res.Header)

	// _body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("response Body:", string(_body))

	body := ParseHTML(res)
	defer res.Body.Close()
	return body
}

func GetResult(parameter map[string]string) *goquery.Document {
	body := Request(parameter)
	return body
}

func main() {
	config := Config{Domain: "http://www.google.com", Path: "/search?q=apple", Method: "GET"}
	header := map[string]string{"Content-Type": "application/json;charset=UTF-8"}
	// parameter := map[string]string{"firstname": "asd"}
	parameter := map[string]string{}

	SetConfig(config)
	SetHeader(header)
	doc := GetResult(parameter)

	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		log.Println(i, s.Text())
	})
}
