package initialize

import (
	"SilicomAPPv0.3/global"
	"github.com/olivere/elastic/v7"
)

func Elastic() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
	if err != nil {
		panic(err)
	}
	global.Es = client
}
