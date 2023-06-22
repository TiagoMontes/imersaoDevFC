package main

import (
	"encoding/json"
	"sync"

	"github.com/TiagoMontes/imersaoDevFC/internal/infra/kafka"
	dto "github.com/TiagoMontes/imersaoDevFC/internal/market/DTO"
	"github.com/TiagoMontes/imersaoDevFC/internal/market/entity"
	"github.com/TiagoMontes/imersaoDevFC/internal/market/transformer"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": 	"host.docker.internal:9094",
		"group.id": 			"myGroup",
		"auto.offset.reset": 	"earliest",
	}

	producer := kafka.NewKafkaProducer(&configMap)
	kafka := kafka.NewConsumer(&configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan) // Create a thread 2

	// recebe do canal do kafka, joga no input, processa, joga no output e dps publica no kafka
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // Thread 3

	go func() {
		for msg := range kafkaMsgChan {
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

}