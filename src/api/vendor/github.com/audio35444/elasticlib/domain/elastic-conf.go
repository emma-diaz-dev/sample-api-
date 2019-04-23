package domain

type ElasticConf struct {
	RootEndpoint string `json:"root-endpoint"`
	Port         string `json:"port"`
	IndicesShow  string `json:"indices-show"`
	IndexName    string `json:"index-name"`
	DocIndex     string `json:"doc-index"`
	DocUpdate    string `json:"doc-update"`
	DocSearch    string `json:"doc-search"`
	DocInsert    string `json:"doc-insert"`
}

/*
GetBaseURL return the base root path
*/
func (ec *ElasticConf) GetBaseURL() string {
	return ec.RootEndpoint + ec.Port
}
