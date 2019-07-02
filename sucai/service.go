package sucai

import (
	"5118api-go/request"
	"5118api-go/responsehandler"
	"strconv"
	"time"
)

type Service interface {
	RequestArticleList(keyword string, pageIndex int, pageSize int) (ald *ArticleListData, err error)
	RequestArticleDetail(guid string) (ad *ArticleDetail, err error)
}

const (
	apiUrl                = "http://apis.5118.com/api/sucai"
	timeOut               = 10 * time.Second
	articleListDataPath   = "data"
	articleDetailDataPath = "data"
)

type service struct {
	//responseHandler responsehandler.Service
	requestService request.Service
}

func (s *service) RequestArticleList(keyword string, pageIndex int, pageSize int) (ald *ArticleListData, err error) {
	res, err := s.requestService.SendRequest(map[string][]string{"keyword": {keyword}, "page_index": {strconv.Itoa(pageIndex)}, "page_size": {strconv.Itoa(pageSize)}})
	if err != nil {
		return
	}
	responseHandler := responsehandler.NewService(res)
	ald = &ArticleListData{}
	err = responseHandler.ExportData(ald, articleListDataPath)
	return
}

func (s *service) RequestArticleDetail(guid string) (ad *ArticleDetail, err error) {
	res, err := s.requestService.SendRequest(map[string][]string{"guid": {guid}})
	if err != nil {
		return
	}
	responseHandler := responsehandler.NewService(res)
	ad = &ArticleDetail{}
	err = responseHandler.ExportData(ad, articleDetailDataPath)
	return
}

func NewService(apiToken string) Service {
	requestService := request.NewService(apiToken, apiUrl, timeOut)
	return &service{requestService: requestService}
}
