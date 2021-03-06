package fetcher

import (
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/utils/text"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.QPS)

// fetch contents from giving url
func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// TODO: should determine encoding automatically
	return text.Gbk2Utf8(contents), nil
}
