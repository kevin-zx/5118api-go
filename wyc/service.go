/*
	5118 伪原创api
	包含了整篇文章和单句的
*/
package wyc

import (
	"errors"
	"github.com/kevin-zx/5118api-go/request"
	"github.com/kevin-zx/5118api-go/responsehandler"
	"strings"
	"time"
)

const (
	apiUrl           = "http://apis.5118.com/wyc/akey"
	sentenceTimeOut  = 10 * time.Second
	sentenceDataPath = "data"
	//articleDetailDataPath = "data"
)

type Service interface {
	WycSentence(sentence string) (wycSentence *Sentence, err error)
}

type service struct {
	requestService request.Service
}

func (s *service) WycSentence(sentence string) (wycSentence *Sentence, err error) {
	sentenceLen := len(strings.Split(sentence, ""))
	if sentenceLen > 150 {
		return nil, errors.New("整句伪原创,单句子不能超过150个字符")
	}
	res, err := s.requestService.SendRequest(map[string][]string{"txt": {sentence}})
	if err != nil {
		return
	}
	responseHandler := responsehandler.NewService(res)
	wycSentence = &Sentence{}
	err = responseHandler.ExportData(wycSentence, sentenceDataPath)
	return
}

func NewService(apiToken string) Service {
	return &service{requestService: request.NewService(apiToken, apiUrl, sentenceTimeOut)}
}
