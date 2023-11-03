package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main(){

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.64.3:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Erro ao gerar o cliente elasticsearch: %s", err)
	}

	logData := map[string]interface{}{
		"timestamp" : fmt.Sprintf("%d", time.Now().Unix()),
		"message"   : "Example de log com data ",
	}
	jsonLog, err := json.Marshal(logData)
	if err != nil {
		log.Fatalf("Erro ao criar json log: %s", err)
	}

	indexName := "logs-go"
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: "", // Deixe em branco para gerar um ID automático
		Body:       bytes.NewReader(jsonLog),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Erro ao enviar a solicitacao de index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Falha ao indexar o log: %s", res.Status())
	} else {
		log.Println("log indexado com sucesso no elastic.")
	}
}