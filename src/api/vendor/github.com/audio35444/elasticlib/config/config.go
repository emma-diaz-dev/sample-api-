package config

import (
	"github.com/audio35444/elasticlib/domain"
)

/*
ElasticConfig has the global elastic config
*/
var ElasticConf = domain.ElasticConf{
	RootEndpoint: "http://localhost",
	DocUpdate:    "{index_name}/_doc/{id}/_update?pretty",
	DocSearch:    "/_search?pretty",
	IndexName:    "?pretty",
	DocIndex:     "%s/_doc/%s?pretty",
	DocInsert:    "/_doc?pretty",
	IndicesShow:  "_cat/indices?v",
	Port:         ":9200/",
}
