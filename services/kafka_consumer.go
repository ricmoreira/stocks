package services

import (
	"encoding/json"
	"log"
	"stocks/config"
	"stocks/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	config      *config.Config
	stockMovServ *StockMovService
}

func NewKafkaConsumer(config *config.Config, sms *StockMovService) *KafkaConsumer {
	return &KafkaConsumer{
		config:      config,
		stockMovServ: sms,
	}
}

func (kc *KafkaConsumer) Run() {

	log.Println("Start receiving from Kafka")

	configConsumer := kafka.ConfigMap{
		"bootstrap.servers":       kc.config.BootstrapServers,
		"group.id":                kc.config.GroupID,
		"auto.offset.reset":       kc.config.AutoOffsetReset,
		"auto.commit.enable":      kc.config.AutoCommitEnable,
		"auto.commit.interval.ms": kc.config.AutoCommitInterval,
	}

	c, err := kafka.NewConsumer(&configConsumer)

	if err != nil {
		panic(err)
	}

	topicsSubs := kc.config.TopicsSubscribed
	err = c.SubscribeTopics(topicsSubs, nil)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {

			topic := *msg.TopicPartition.Topic

			switch topic {
			case "invoices":
				log.Println(`Reading an invoices topic message`)
				invoices, err := kc.parseInvoicesMessage(msg.Value)
				if err != nil {
					log.Printf("Error parsing event message value. Message %v \n Error: %s\n", msg.Value, err.Error())
					break
				}

				// save stock movements to database
				_, e := kc.stockMovServ.CreateStockMovementsFromInvoices(invoices)
				// save stock movement to database
				if e != nil {
					log.Printf("Error saving stock movements to database\n Error: %s\n", e.Response)
					break
				}
			case "invoice":
				log.Println(`Reading an invoice topic message`)
				invoice, err := kc.parseInvoiceMessage(msg.Value)
				if err != nil {
					log.Printf("Error parsing event message value. Message %v \n Error: %s\n", msg.Value, err.Error())
					break
				}

				// save stock movements to database
				_, e := kc.stockMovServ.CreateStockMovementsFromInvoice(invoice)
				// save stock movement to database
				if e != nil {
					log.Printf("Error saving stock movements to database\n Error: %s\n", e.Response)
					break
				}	
			default: //ignore any other topics
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

func (kc *KafkaConsumer) parseInvoicesMessage(messageValue []byte) (*[]*models.Invoice, error) {
	invoices := make([]*models.Invoice, 0)
	err := json.Unmarshal(messageValue, &invoices)

	if err != nil {
		return nil, err
	}

	return &invoices, nil
}

func (kc *KafkaConsumer) parseInvoiceMessage(messageValue []byte) (*models.Invoice, error) {
	invoice := models.Invoice{}
	err := json.Unmarshal(messageValue, &invoice)

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
