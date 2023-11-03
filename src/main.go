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

func main2() {
	// Configurar as informações de conexão com o Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.64.3:9200"},
	}

	// Criar um cliente Elasticsearch
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Erro ao criar o cliente Elasticsearch: %s", err)
	}

	// Criar um JSON com os dados do log
	logData := map[string]interface{}{
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
		"message":   "Exemplo de log numero 2 ,",
	}
	jsonLog, err := json.Marshal(logData)
	if err != nil {
		log.Fatalf("Erro ao criar o JSON do log: %s", err)
	}

	// Preparar a solicitação de índice
	indexName := "teste"
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: "", // Deixe em branco para gerar um ID automático
		Body:       bytes.NewReader(jsonLog),
		Refresh:    "true",
	}

	// Enviar a solicitação de índice
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Erro ao enviar a solicitação de índice: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Falha ao indexar o log no Elasticsearch: %s", res.Status())
	} else {
		log.Println("Log indexado com sucesso no Elasticsearch.")
	}
}
