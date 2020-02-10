package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemSave() chan interface{}{
	out := make(chan interface{})
	go func(){
		itemcount := 0
		for {
			item := <-out
			log.Printf("Item Saver:Got$ %d,%v",itemcount,item)
			itemcount++
		}
	}()
	return out
}
func save(item interface{}){
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	_, err = client.Index().Index("Book Profile").Type("Douban").BodyJson(item).Do(context.Background())
	if err != nil{
		panic(err)
	}

}
