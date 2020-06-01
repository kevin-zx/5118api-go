package baidupc

import (
	"github.com/kevin-zx/5118api-go/request"
	"github.com/kevin-zx/5118api-go/responsehandler"
	"strconv"
	"time"
)

const (
	apiUrl          = "http://apis.5118.com/keyword/baidumobile"
	timeOut         = 30 * time.Second
	baiduPcDataPath = "data.baidumobile"
)

type Service interface {
	RequestBaiduMobileResults(siteDomain string, pageIndex int) ([]*BaiduMobileResult, error)
}

type service struct {
	requestService request.Service
}

func (s *service) RequestBaiduMobileResults(siteDomain string, pageIndex int) ([]*BaiduMobileResult, error) {
	res, err := s.requestService.SendRequest(map[string][]string{"url": {siteDomain}, "page_index": {strconv.Itoa(pageIndex)}})
	if err != nil {
		return nil, err
	}
	responseHandler := responsehandler.NewService(res)

	var bprs []*BaiduMobileResult
	err = responseHandler.ExportData(&bprs, baiduPcDataPath)
	return bprs, err
}

func NewBaiduMobileRankKeywordService(apiToken string) Service {
	requestService := request.NewService(apiToken, apiUrl, timeOut)
	return &service{requestService: requestService}
}
