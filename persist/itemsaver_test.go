package persist

import (
	"context"
	"encoding/json"
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/model"
	"github.com/olivere/elastic"
	"testing"
)

func TestSaveItem(t *testing.T) {
	expectedItem := engine.Item{
		Url:  "http://album.zhenai.com/u/107790366",
		Type: "zhenai",
		Id:   "107790366",
		Payload: model.Profile{
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
		},
	}

	index := "data_profile"
	client, err := elastic.NewClient(
		elastic.SetSniff(false), elastic.SetURL("http://192.168.31.65:9200"))

	if err != nil {
		panic(err)
	}

	err = SaveItem(client, index, expectedItem)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expectedItem.Type).
		Id(expectedItem.Id).
		Do(context.Background())

	//fmt.Printf("%s", *resp.Source)

	var actualItem engine.Item

	json.Unmarshal(*resp.Source, &actualItem)

	t.Logf("source: %s", *resp.Source)

	// cast actualItem.payload to a Profile
	profile, _ := model.FromJsonObj(actualItem.Payload)
	actualItem.Payload = profile

	if actualItem != expectedItem {
		t.Errorf("expected %+v, actual: %+v", expectedItem, actualItem)
	}
}
