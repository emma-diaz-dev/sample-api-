package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/audio35444/sample-api/src/api/dao/rc"
	"github.com/audio35444/sample-api/src/api/domain"
)

var (
	baseRC = rc.CreateBaseRestClient(120 * time.Millisecond)
)

func GetExamples() (*domain.RequestEntity, error) {
	e := &domain.RequestEntity{}

	bytes, err := baseRC.Get("http://127.0.0.1:3000/entities")
	if err != nil {
		fmt.Printf("[Method: GetExamples] [ErrType: request] [ErrMsg: %s]", err.Error())
		return nil, err
	}
	err = json.Unmarshal(bytes, e)
	if err != nil {
		fmt.Printf("[Method: GetExamples] [ErrType: Unmarshal] [ErrMsg: %s]", err.Error())
		return nil, err
	}
	fmt.Println(e)
	return e, nil
}
