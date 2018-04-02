package parser

import (
	"imooc/craw/engine"
	"regexp"
)

var (
	cityre    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlre = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">`)
)

func ParseCity(contents []byte) engine.ParseResult {
	amtch := cityre.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range amtch {
		userName := string(m[2])
		//result.Items = append(result.Items, "User: "+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParaserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName)
			},
		})
		//fmt.Printf("%s\n %s \n", m[1], m[2])
	}
	matchs := cityUrlre.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:         string(m[1]),
			ParaserFunc: ParseCity,
		})
	}
	return result
}
