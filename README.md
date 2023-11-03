# Elasticsearch 7 - Udemy

## Description

Elasticserach nada mais é que um mecanismo de busca desenvolvido em java em cima do LUCENE

vantagens:
    - disponibiliza dados em tempo real 
    - pode ser distribuido
    - é orientado a documentos 
    - disponibiliza uma api RESTful

index sao os bancos de dados 
tipos sap as tabelas 
documentos sao os registros 
campos sao as colunas 

![arquitetura](./img/Screenshot%202023-10-29%20at%2015.54.41.png)

## Components

cluster

node

data node

master node

coordinating-only node

indices
    indices sao containers para armazenar dados semelhantes a um bando de dados em um banco relacional . Um indice contem a colecao de documentos que possuem caracteristicas semelhantes ou entao logicamente relacionados 

type
    tipo é um agrupamento logico dos documentos dentro do indice

documento 
    é a parte indexada pelo elasticsearch em um formato JSON,

shard
    é um sub conjunto completo do indice 

![shards](/img/Screenshot%202023-10-29%20at%2016.02.48.png)


replicas
    é a a ideia de replicar os dados em diferentes maquinas 

>Note: Em resumo os shards sao a quantidade de vezes que voce vai quebrar uma informacao e a replica a quantidade de vezes que ela vai ser repetida.

## Install 

````
    $ wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt add -

    $ sudo apt-get install apt-transport-https

    $ echo "deb https://artifacts.elastic.co/package/7.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-7.x.list

    $ sudo apt-get update && apt-get install elasticserach 
````

## Configuration

````
    $ sudo vim /etc/elasticsearch/elasticsearch.yml

    $ sudo /bin/systemctl daemon-reload

    $ sudo /bin/systemctl enable elasticsearch.service

    $ sudo /bin/systemctl start elascticsearch.service
````

## Instalação do Elasticsearch

### Instalando Elasticsearch 7

Depois de logar no Ubuntu Server, execute os seguintes comandos:

    wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
     
    sudo apt-get install apt-transport-https
     
    echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" |

    sudo tee -a /etc/apt/sources.list.d/elastic-7.x.list
     
    sudo apt-get update && sudo apt-get install elasticsearch

Agora, altere as configurações usando o vi:

    sudo vi /etc/elasticsearch/elasticsearch.yml

Remova o comentário (#) da linha node.nome, para isso coloque o vi em modo de inserção teclando i

Altere network.host para 0.0.0.0

Altere discovery.seed.hosts para [“127.0.0.1”]

E cluster.initial_master_nodes para [“node-1”]

Quando terminar, tecle ESC para sair do modo de inserção, então digite :wq para salvar e sair do vi.

Execute os comandos abaixo para iniciar o Elasticsearch e para configurar a inicialização automática.

    sudo /bin/systemctl daemon-reload
    sudo /bin/systemctl enable elasticsearch.service
    sudo /bin/systemctl start elasticsearch.service

O serviço pode levar alguns minutos até estar pronto.


## Index - shakespeare

``````
    $ wget http://media.sundog-soft.com/es7/shakes-mapping.json
     
    $ curl -H 'Content-Type: application/json' -XPUT 127.0.0.1:9200/shakespeare --data-binary @shakes-mapping.json

    $ wget http://media.sundog-soft.com/es7/shakespeare_7.0.json
     
    $ curl -H 'Content-Type: application/json' -XPOST '127.0.0.1:9200/shakespeare/_bulk?pretty' --data-binary
    @shakespeare_7.0.json

    $ curl -H 'Content-Type: application/json' -XGET
    '127.0.0.1:9200/shakespeare/_search?pretty' -d '
    {
    "query" : {
    "match_phrase" : {
    "text_entry" : "to be or not to be"
    }
    }
    }
    '
``````


## Shards and indexs

Shards primarios e replicas 

requisicoes de escrita sao direcionadas para o shard primario e entao replicadas 
requisicoes de leitura sao direcionadas para o shard primario ou qualquer outra replica 

numero impar de nodes 


## Add data informations 

````
    $ curl -XPUT 127.0.0.1:9200/_bulk?pretty --data-binary @movies.json
````

nesse caso ele adiciona o arquivo , o index vai junto a informacao , ex:

````json
{ "create" : { "_index" : "movies", "_id" : "135569" } }
{ "id": "135569", "title" : "Star Trek Beyond", "year":2016 , "genre":["Action", "Adventure", "Sci-Fi"] }
{ "create" : { "_index" : "movies", "_id" : "122886" } }
{ "id": "122886", "title" : "Star Wars: Episode VII - The Force Awakens", "year":2015 , "genre":["Action", "Adventure", "Fantasy", "Sci-Fi", "IMAX"] }
{ "create" : { "_index" : "movies", "_id" : "109487" } }
{ "id": "109487", "title" : "Interstellar", "year":2014 , "genre":["Sci-Fi", "IMAX"] }
{ "create" : { "_index" : "movies", "_id" : "58559" } }
{ "id": "58559", "title" : "Dark Knight, The", "year":2008 , "genre":["Action", "Crime", "Drama", "IMAX"] }
{ "create" : { "_index" : "movies", "_id" : "1924" } }
{ "id": "1924", "title" : "Plan 9 from Outer Space", "year":1959 , "genre":["Horror", "Sci-Fi"] }
````

o documentos sao imutaveis , o que ele faz no put é criar um novo 

*mappings*

### Query lite

uri search

executando consulta com filtro 

````
    $ curl -XGET "http://127.0.0.1:9200/movies/_search?q=title:star&pretty"
````

Request com JSON

``````
curl -XGET http://127.0.0.1:9200/movies/_search?pretty -d '
{
    "query":{
        "match":{
            "title": "star"
        }
    }
}'
``````

Alguns tipos de filters
    - term
    - terms
    - range 
    - exists
    - missing
    - bool

Alguns tipos de queries 
    - macth_all
    - macth
    - multi_match
    - bool 