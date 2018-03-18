package parser

import (
	"imooc/craw/engine"
	"regexp"
)

const citylistre = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylistre)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:         string(m[1]),
			ParaserFunc: ParseCity,
		})
		//fmt.Printf("%s\n %s \n", m[1], m[2])
	}
	return result
}
