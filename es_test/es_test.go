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

// init index test
func TestEsInitIndex(t *testing.T) {
	esInit()

	(&Model{}).InitIndex()
}

// curd test

// add test
func TestEsBatchAdd(t *testing.T) {
	esInit()

	rows := []*Model{}
	for i := 0; i < 1500; i++ {
		row := &Model{
			User:     "99949879",
			Message:  "hhhh",
			Retweets: 3,
		}
		rows = append(rows, row)
	}
	BatchAdd(rows)

	TestEsGetList(t)
}

func TestEsAdd(t *testing.T) {
	esInit()

	row := &Model{
		User: "999",
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

func TestEsDelIndexById(t *testing.T) {
	esInit()

	DelIndex()

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
		Urls:       []string{"http://es-cn-n6w1r3anu0006zb5t.public.elasticsearch.aliyuncs.com:9200"},
		User:       "elastic",
		Password:   "XZ527shortvideo",
		BulkWorker: 4,
		LogLevel:   1,
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

// init index
func (t *Model) InitIndex() error {
	index := t.TableName()
	client := gom.Es()
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}
	if !exists {
		// Create a new index.
		mapping := `
{
	"settings":{
		"number_of_shards": 2,
		"number_of_replicas": 0
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
		createIndex, err := client.CreateIndex(index).Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			err := fmt.Errorf("IndexInit-!createIndex.Acknowledged")
			return err
		}
	}
	return err
}

// add
func BatchAdd(rows []*Model, ps ...*elastic.BulkProcessor) (err error) {
	index := (&Model{}).TableName()

	// 如果有多次添加，则从外面传processor进来
	var p *elastic.BulkProcessor
	if len(ps) == 0 {
		p = gom.BulkProcessor(index)
		defer p.Flush()
	} else {
		p = ps[0]
	}
	if err != nil {
		log.Println(err)
		return
	}

	for _, row := range rows {
		r := elastic.NewBulkIndexRequest().Index(index).Doc(row)
		// Add the request r to the processor p
		p.Add(r)
	}
	return
}

func Add(row *Model) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()
	// save
	addResult, err := client.Index().
		Index(index).
		BodyJson(row).
		Refresh("wait_for").
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

func DelByQuery(user string) (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewTermQuery("user", user))

	_, err = client.DeleteByQuery().Index(index).Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DelIndex() (err error) {
	client := gom.Es()
	index := (&Model{}).TableName()

	_, err = client.DeleteIndex(index).Do(context.Background())
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

//bool query 条件

type ListParam struct {
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	SortOrder string `json:"sort_order"`
	SortField string `json:"sort_field"`

	Keyword string `json:"keyword"`
}

func (p *ListParam) ToFilter() *gom.EsSearch {
	var search gom.EsSearch
	if len(p.Keyword) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("user", p.Keyword))
	}

	// range
	// if p.LikeCountMax > 0 {
	// 	rangeQuery := elastic.NewRangeQuery("like_count").
	// 		Gte(param.LikeCountMin).Lt(param.LikeCountMax)
	// 	boolQuery.Filter(rangeQuery)
	// }

	if len(p.SortField) != 0 {
		search.Sorters = append(search.Sorters, gom.EsOrder(p.SortField, p.SortOrder))
	}

	pageNum, pageSize := gom.Page(p.PageNum, p.PageSize)
	search.From = (pageNum - 1) * pageSize
	search.Size = pageSize
	return &search
}

func GetList(param *ListParam) (total int64, rows []*Model, err error) {
	client := gom.Es()

	// filter
	filter := param.ToFilter()
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(filter.MustQuery...)
	boolQuery.MustNot(filter.MustNotQuery...)
	boolQuery.Should(filter.ShouldQuery...)
	boolQuery.Filter(filter.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(filter.MustQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}

	// if param.Tag != "" {
	// 	boolQuery.Filter(elastic.NewTermQuery("tag", param.Tag))
	// }
	// if param.MultiName != "" {
	// 	boolQuery.Filter(elastic.NewMultiMatchQuery(param.MultiName, "video_name", "tag", "video_id", "account_id"))
	// }

	index := (&Model{}).TableName()
	searchResult, err := client.Search().
		Index(index).
		Query(boolQuery).
		Sort(filter.Sorters...).
		From(filter.From).Size(filter.Size).
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
