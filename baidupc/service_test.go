package baidupc

import (
	"fmt"
	"testing"
)

func TestService_RequestBaiduPcResults(t *testing.T) {
	s := NewBaiduPCRankKeywordService("xxxxxxxxxxxxxxxxxxxx")
	brs, err := s.RequestBaiduPcResults("www.sunnat.com.cn", 1)
	if err != nil {
		panic(err)
	}
	for _, br := range brs {
		fmt.Printf("%+v\n", br)
	}
}
