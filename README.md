# go-lang-crawler 1.0 beta

## Documentation

## Installation

```sh
go get github.com/PuerkitoBio/goquery
```

## Usage
```go
func main() {
    // Delcare variable
	config := Config{Domain: "http://www.google.com", Path: "/search?q=apple", Method: "GET"}
	header := map[string]string{"Content-Type": "application/json;charset=UTF-8"}
	parameter := map[string]string{}

	SetConfig(config) // Setup your taget which website you want to do the crawl (Domain, Path, Method)
    	SetHeader(header) // Setup the header method
    
    	doc := GetResult(parameter) // Get the response from the request 
    
    	// Query the selector via Go Query
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
        log.Println(i, s.Text())
        // https://godoc.org/github.com/PuerkitoBio/goquery
	})
}
```
