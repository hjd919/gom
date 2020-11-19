package es_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hjd919/gom"
	"github.com/olivere/elastic/v7"
)

// curd test

// add test
func TestEsAdd(t *testing.T) {
	esInit()

	row := &Model{
		User: "99949879",
	}
	Add(row)

	TestEsGetList(t)
}

func TestEsAddById(t *testing.T) {
	id := "3"
	esInit()

	row := &Model{
		User: "hjd23",
	}
	AddById(id, row)

	TestEsGetList(t)
}

// upd test
func TestEsUpd(t *testing.T) {
	id := "3"
	esInit()

	updRow := map[string]interface{}{
		"user":     "3333",
		"retweets": 3,
	}
	UpdById(id, updRow)

	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

// incr test
func TestEsIncrId(t *testing.T) {
	id := "1"
	esInit()
	IncrFieldById(id, "retweets")

	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

// del test
func TestEsDelById(t *testing.T) {
	id := "s0Uv4HUB780-B16Wb4F_"
	esInit()

	DelById(id)

	TestEsGetList(t)
}

// get test
func TestEsGetById(t *testing.T) {
	id := "1"
	esInit()
	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

func TestEsGetList(t *testing.T) {
	esInit()

	total, rows, err := GetList(&ListParam{
		PageNum:   1,
		PageSize:  10,
		SortType:  "",
		SortField: "",
		Keyword:   "",
	})
	log.Println(total, gom.JsonEncode(rows), err)
}

func esInit() {
	esConf := &gom.EsConfig{
		Urls:     []string{"http://es-cn-n6w1r3anu0006zb5t.public.elasticsearch.aliyuncs.com:9200"},
		User:     "elastic",
		Password: "XZ527shortvideo",
	}
	gom.EsInit(esConf)
}

// model
type Model struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Message  string `json:"message"`
	Retweets int    `json:"retweets"`
}

func (t *Model) TableName() string {
	return "twitter"
}

// add&upd
func TestEsUpsert(t *testing.T) {
	esInit()
	client := gom.Es()
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"user":{
				"type":"keyword"
			},
			"message":{
				"type":"text",
				"store": true,
				"fielddata": true
			},
			"retweets":{
				"type":"long"
			},
			"tags":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			},
			"suggest_field":{
				"type":"completion"
			}
		}
	}
}
`
		createIndex, err := client.CreateIndex("twitter").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Index a tweet (using JSON serialization)
	tweet1 := Model{User: "olivere", Message: "Take Five", Retweets: 0}
	put1, err := client.Index().
		Index("twitter").
		Id("1").
		BodyJson(tweet1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}

// add
func Add(row *Model) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()
	// save
	addResult, err := client.Index().
		Index(index).
		BodyJson(row).
		Do(context.Background())
	if err != nil {
		return
	}
	row.Id = addResult.Id
	return
}

func AddById(id string, row *Model) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()
	// save
	_, err = client.Index().
		Index(index).
		Id(id).
		BodyJson(row).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

// update
func UpdById(id string, updRow map[string]interface{}) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	// save
	_, err = client.Update().
		Index(index).
		Id(id).
		Doc(updRow).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

func IncrFieldById(id string, field string) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	script := elastic.NewScript(fmt.Sprintf("ctx._source.%s += params.num", field)).Param("num", 1)
	_, err = client.Update().Index(index).Id(id).
		Script(script).
		Upsert(map[string]interface{}{field: 0}). // init field
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

// del
func DelById(id string) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	_, err = client.Delete().Index(index).Id(id).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// get
func GetById(id string) (row *Model, err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	getResult, err := client.Get().
		Index(index).
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	var r Model
	err = json.Unmarshal(getResult.Source, &r)
	if err != nil {
		return
	}
	// append id
	r.Id = getResult.Id
	row = &r
	return
}

type ListParam struct {
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	SortType  string `json:"sort_type"`
	SortField string `json:"sort_field"`

	Keyword string `json:"keyword"`
}

func GetList(param *ListParam) (total int64, rows []*Model, err error) {
	client := gom.Es()
	search := client.Search()

	// filter
	if param.Keyword != "" {
		termQuery := elastic.NewTermQuery("user", param.Keyword)
		search = search.Query(termQuery)
	}

	// sort
	sortField, sortType := "_id", false
	if param.SortField != "" {
		sortField = param.SortField
	}
	if param.SortType != "" {
		sortType = gom.EsSortType(param.SortType)
	}

	// page
	pageNum := gom.PageNum(param.PageNum)
	pageSize := gom.PageSize(param.PageSize)
	from := (pageNum - 1) * pageSize

	index := (&Model{}).TableName()
	searchResult, err := search.
		Index(index).
		Sort(sortField, sortType).
		From(from).Size(pageSize).
		// Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		log.Println(err)
		return
	}

	total = searchResult.TotalHits()
	if total == 0 {
		fmt.Print("Found no tweets\n")
		return
	}

	rows = make([]*Model, 0)
	for _, hit := range searchResult.Hits.Hits {
		var r Model
		err = json.Unmarshal(hit.Source, &r)
		if err != nil {
			return
		}
		// append id
		r.Id = hit.Id
		rows = append(rows, &r)
	}
	return
}
