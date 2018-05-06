package persist

import (
	"imooc/crawer/engine"
	"imooc/crawer/persist"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(itm engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, itm)
	if err == nil {
		*result = "ok"
	}
	return err
}
