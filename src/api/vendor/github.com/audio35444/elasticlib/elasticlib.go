package elasticlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/audio35444/elasticlib/config"
	"github.com/audio35444/elasticlib/domain"
)

func SetBaseEndpoint(url, port string) {
	config.ElasticConf.RootEndpoint = "http://" + url
	config.ElasticConf.Port = ":" + port + "/"
}
func SetDefaultBaseEndpoint() {
	config.ElasticConf.RootEndpoint = "http://localhost"
	config.ElasticConf.Port = ":9200/"
}
func GetDocs(indexName string) (*[]byte, error) {
	return genericRequest(http.MethodGet,
		config.ElasticConf.GetBaseURL()+indexName+config.ElasticConf.DocSearch,
		nil,
		false,
	)
}

func GetDoc(indexName string, docId string) (*[]byte, error) {
	return genericRequest(http.MethodGet,
		config.ElasticConf.GetBaseURL()+fmt.Sprintf(config.ElasticConf.DocIndex, indexName, docId),
		nil,
		false)
}

func InsertDoc(element interface{}, indexName string) (*[]byte, error) {
	result, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	return genericRequest(http.MethodPost,
		config.ElasticConf.GetBaseURL()+indexName+config.ElasticConf.DocInsert,
		strings.NewReader(string(result)),
		true)
}
func UpdateDoc(element interface{}, indexName string, docId string) (*[]byte, error) {
	result, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	flag, err := existDoc(docId, indexName)
	if err != nil {
		return nil, err
	}
	if !flag {
		return nil, errors.New("Doc Not Exist")
	}
	return genericRequest(http.MethodPut,
		config.ElasticConf.GetBaseURL()+fmt.Sprintf(config.ElasticConf.DocIndex, indexName, docId),
		strings.NewReader(string(result)),
		true)
}
func DeleteDoc(indexName string, docId string) (*[]byte, error) {
	return genericRequest(http.MethodDelete,
		config.ElasticConf.GetBaseURL()+fmt.Sprintf(config.ElasticConf.DocIndex, indexName, docId),
		nil,
		false)
}

func GetIndices() (*[]byte, error) {
	return genericRequest(http.MethodGet,
		config.ElasticConf.GetBaseURL()+config.ElasticConf.IndicesShow,
		nil,
		false)
}
func NewIndex(indexName string) (*[]byte, error) {
	return genericRequest(http.MethodPut,
		config.ElasticConf.GetBaseURL()+indexName+config.ElasticConf.IndexName,
		nil,
		false)
}
func DeleteIndex(indexName string) (*[]byte, error) {
	return genericRequest(http.MethodDelete,
		config.ElasticConf.GetBaseURL()+indexName+config.ElasticConf.IndexName,
		nil,
		false)

}
func genericRequest(method string, fullPath string, body io.Reader, isJson bool) (dataResult *[]byte, errResult error) {
	err := isElasticOn()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, fullPath, body)
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	if isJson {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := client.Do(req)
	if err != nil {
		errResult = err
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errResult = err
		return
	}
	dataResult = &data
	return

}

func isElasticOn() error {
	_, err := http.Get(config.ElasticConf.GetBaseURL())
	if err != nil {
		return errors.New("elastic-off")
	}
	return nil
}

func existDoc(docId string, indexName string) (bool, error) {
	data, err := GetDoc(indexName, docId)
	if err != nil {
		return false, err
	}
	objStatus := &domain.Status{}
	if err = json.Unmarshal(*data, objStatus); err != nil {
		return false, err
	}
	return objStatus.Found, nil
}
