package persist

import (
	"context"
	"github.com/GerryLon/go-crawler/engine"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
	"strings"
)

func ItemSaver(elasticIndex string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false), elastic.SetURL("http://192.168.31.65:9200"))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			err := SaveItem(client, elasticIndex, item)

			if err != nil {
				log.Printf("save item error: %v", err)
			}
		}

	}()

	return out, nil
}

// save item to elasticsearch
func SaveItem(client *elastic.Client, elasticIndex string, item engine.Item) error {

	if len(strings.TrimSpace(item.Type)) == 0 {
		return errors.New("item must supply type")
	}

	indexService := client.Index().
		Index(elasticIndex).
		Type(item.Type).
		BodyJson(item)

	if len(strings.TrimSpace(item.Id)) != 0 {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	return err
}
