package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCity(contents)

	//var expectedNicknames = []string{"征婚", "遇见美好", "重天"}
	//
	//for i, nickname := range expectedNicknames {
	//	if result.Items[i] != nickname {
	//		t.Errorf("expect nickname is %s, but got:%s\n", nickname, result.Items[i])
	//	}
	//}

	var expectedUrls = []string{
		"http://album.zhenai.com/u/106280036",
		"http://album.zhenai.com/u/107822680",
		"http://album.zhenai.com/u/1300644903",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expect city url is %s, but got:%s\n", url, result.Requests[i].Url)
		}
	}

}
