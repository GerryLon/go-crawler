package parser

import (
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/model"
	"io/ioutil"
	"log"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "会员107790366", "http://album.zhenai.com/u/107790366", "107790366")

	if len(result.Items) < 1 {
		log.Printf("Items should have at least ONE element, but got %d", len(result.Items))
	}

	expectedProfile := model.Profile{
		Name:          "会员107790366",
		Age:           28,
		Gender:        "女",
		Height:        165,
		Weight:        0,
		Salary:        "3000元以下",
		Marriage:      "未婚",
		Education:     "大专",
		Occupation:    "",
		NativePlace:   "",
		Workplace:     "陕西渭南",
		Constellation: "双鱼座",
		Zodiac:        "马",
		House:         "",
		Car:           "未购车",
		Pic:           "http://photo11.zastatic.com/images/photo/26948/107790366/1506944293939.jpg",
	}

	expectedItem := engine.Item{
		Url:     "http://album.zhenai.com/u/107790366",
		Type:    "zhenai",
		Id:      "107790366",
		Payload: expectedProfile,
	}

	realItem := result.Items[0]

	if expectedItem != realItem {
		log.Printf("Expected profile:\n%+v\nBut got:\n%+v", expectedItem, realItem)
	}
}

//func TestHan(t *testing.T) {
//	re := regexp.MustCompile(`(a|\p{Han})+`)
//	fmt.Println(re.FindAllString("a2你3242我你好", -1))
//}
