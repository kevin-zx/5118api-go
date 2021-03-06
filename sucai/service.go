package sucai

import (
	"github.com/kevin-zx/5118api-go/request"
	"github.com/kevin-zx/5118api-go/responsehandler"
	"strconv"
	"time"
)

type Service interface {
	RequestArticleList(keyword string, pageIndex int, pageSize int) (ald *ArticleListData, err error)
	RequestArticleDetail(guid string) (ad *ArticleDetail, err error)
}

const (
	apiUrl                = "http://apis.5118.com/api/sucai"
	timeOut               = 30 * time.Second
	articleListDataPath   = "data"
	articleDetailDataPath = "data.data"
)

type service struct {
	requestService request.Service
}

func (s *service) RequestArticleList(keyword string, pageIndex int, pageSize int) (ald *ArticleListData, err error) {
	// todo: 这里验证100是不可以的,还是要验证
	if pageSize > 20 {
		pageSize = 20
	}
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
