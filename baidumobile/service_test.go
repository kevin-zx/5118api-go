package baidupc

import (
	"fmt"
	"testing"
)

func TestService_RequestBaiduMobileResults(t *testing.T) {
	s := NewBaiduMobileRankKeywordService("CDC974A0C6EA45CFB9E5AE328C7DCA51")
	brs, err := s.RequestBaiduMobileResults("www.sunnat.com.cn", 1)
	if err != nil {
		panic(err)
	}
	for _, br := range brs {
		fmt.Printf("%+v\n", br)
	}
}
