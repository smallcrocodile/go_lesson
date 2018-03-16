package parser

import (
	"imooc/craw/engine"
	"regexp"
)

const cityre = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityre)
	amtch := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range amtch {
		userName := string(m[2])
		result.Items = append(result.Items, "User: "+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParaserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName)
			},
		})
		//fmt.Printf("%s\n %s \n", m[1], m[2])
	}
	return result
}
