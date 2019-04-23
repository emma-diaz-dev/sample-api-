package services

import (
	"github.com/audio35444/elasticlib"
)

func NewIndex(indexName string) (string, error) {
	body, err := elasticlib.NewIndex(indexName)
	if err != nil {
		return "", err
	}
	return string(*body), nil
}
