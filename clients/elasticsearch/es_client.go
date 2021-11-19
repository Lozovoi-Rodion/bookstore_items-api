package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(v interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.0.1.1:9200", "http://10.0.1.2:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}
