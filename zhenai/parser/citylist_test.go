package parser

import (
	"io/ioutil"
	"testing"
)

const expectedCityCount = 470

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	if len(result.Items) != expectedCityCount {
		t.Errorf("expected item count: %d, but got: %d\n", expectedCityCount, len(result.Items))
	}

	var expectedCities = []string{"阿坝", "阿克苏", "阿拉善盟"}

	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expect city name is %s, but got:%s\n", city, result.Items[i])
		}
	}

	var expectedUrls = []string{
		"http://city.zhenai.com/aba",
		"http://city.zhenai.com/akesu",
		"http://city.zhenai.com/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expect city url is %s, but got:%s\n", url, result.Requests[i].Url)
		}
	}

}
