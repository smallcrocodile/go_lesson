package parser

import (
	"imooc/craw/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "HiSiri")
	if len(result.Items) != 1 {
		t.Error("Result should contain 1 element;but has %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}
	if profile != expected {
		t.Errorf("expected %v;but was %v", expected, profile)
	}

}
