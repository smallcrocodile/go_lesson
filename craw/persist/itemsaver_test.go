package persist

import (
	"context"
	"encoding/json"
	"testing"

	"imooc/craw/engine"
	"imooc/craw/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:     "HiSiri",
			Gender:   "女",
			Age:      28,
			Height:   163,
			Weight:   100,
			Income:   "3001-5000元",
			Marriage: "未婚",
			Hukou:    "内蒙古赤峰",
			Xingzuo:  "金牛座",
			Hose:     "自住",
			Car:      "未购车",
		},
	}

	//save expected
	err := save(expected)
	if err != nil {
		panic(err)
	}

	//TODO: Try to start up elastic search
	//here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	//fetch expected
	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v,%s", resp, resp.Source)
	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//verify result
	if actual != expected {
		t.Errorf("got %v;expected %v", actual, expected)
	}
}
