/*
Package 的作用是解析5118返回的json用的,一般来说对内部用,外部不应该使用
*/
package responsehandler

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
)

type Service interface {
	validateResponse() error
	ExportData(v interface{}, path string) error
}

type service struct {
	responseContent string
	response        *http.Response
	jsonObj         gjson.Result
}

func (s *service) readContent() (err error) {
	s.responseContent, err = readContentFromResponse(s.response, "utf-8")
	if err != nil {
		if err.Error() == "数据为空" || err.Error() == "读取头部文件为空" {
			return nil
		} else {
			return
		}
	}
	return
}

// 到处errInfo 判断请求接口是否正常
func (s *service) exportErrInf() (errCode int64, errMsg string, err error) {
	errCodeResult := s.jsonObj.Get("errcode")
	var errMsgResult gjson.Result
	if !errCodeResult.Exists() {
		errCodeResult = s.jsonObj.Get("data.errcode")
		if !errCodeResult.Exists() {
			err = errors.New("5118接口返回json, 错误信息解析错误")
			return
		} else {
			errMsgResult = s.jsonObj.Get("data.errmsg")
		}
	} else {
		errMsgResult = s.jsonObj.Get("errmsg")
	}
	errCode = errCodeResult.Int()
	errMsg = errMsgResult.String()
	return
}

// 检测5118返回的内容是否成功,或者内容是否合法
func (s *service) validateResponse() (err error) {
	err = s.readContent()
	if err != nil {
		return err
	}
	if s.response.StatusCode != 200 {
		// 状态码异常
		return errors.New(fmt.Sprintf("5118接口请求解析错误 StatusCode: %d , Message: %s", s.response.StatusCode, s.responseContent))
	} else {
		s.jsonObj = gjson.Parse(s.responseContent)
		if !s.jsonObj.Exists() {
			return errors.New("5118接口请求解析错误, 返回json格式错误")
		}
		var errCode int64
		var errMsg string
		errCode, errMsg, err = s.exportErrInf()
		if err != nil {
			return
		}
		if errMsg != "" || errCode != 0 {
			return errors.New(fmt.Sprintf("5118接口请求解析错误, errcode:%d , errmsg:%s", errCode, errMsg))
		}
	}
	return nil
}

// 根据response到处对应的data
func (s *service) ExportData(v interface{}, path string) (err error) {
	err = s.validateResponse()
	if err != nil {
		return
	}
	dataObj := s.jsonObj
	if path != "" {
		dataObj = s.jsonObj.Get(path)
	}
	if !dataObj.Exists() {
		return errors.New("5118接口请求解析错误,数据结果不存在")
	}
	err = json.Unmarshal([]byte(dataObj.Raw), v)
	return
}

func NewService(response *http.Response) Service {
	return &service{response: response}
}

func readContentFromResponse(response *http.Response, charset string) (string, error) {
	defer response.Body.Close()
	var err error
	var htmlbytes []byte
	contentEncoding, ok := response.Header["Content-Encoding"]
	if ok && contentEncoding[0] == "gzip" {
		gzreader, err := gzip.NewReader(response.Body)
		if err != nil {
			return "", err
		}

		for {
			buf := make([]byte, 1024)
			n, err := gzreader.Read(buf)
			if err != nil && err != io.EOF {
				gzreader.Close()
				return "", err
			}
			if n == 0 {
				break
			}
			htmlbytes = append(htmlbytes, buf...)
		}
		gzreader.Close()
		//htmlbytes,err=ioutil.ReadAll(gzreader)
		//println(string(htmlbytes))
	} else {
		htmlbytes, err = ioutil.ReadAll(response.Body)
	}
	//response.Body = reader

	//if response.StatusCode >= 300 || response.StatusCode < 200 {
	//	return "", errors.New(fmt.Sprintf("状态码为: %d", response.StatusCode))
	//}
	hreader := bytes.NewReader(htmlbytes)

	char, data := detectContentCharset(hreader)
	if charset != "" {
		char = charset
	}
	if data == nil {
		return "", errors.New("数据为空")
	}

	dec := mahonia.NewDecoder(char)

	preRd := dec.NewReader(data)
	if preRd == nil {
		return "", errors.New("读取头部文件为空")
	}
	preBytes, err := ioutil.ReadAll(preRd)
	reBytes, err := ioutil.ReadAll(hreader)
	if err != nil {
		return "", err
	}
	return string(append(preBytes, reBytes...)), err
}

// 这里他娘的Peek是针对Reader的 针对不了io.Reader 所以 io.Reader其实是进行了位移的
func detectContentCharset(body io.Reader) (string, *bufio.Reader) {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); name != "" {
			return name, r
		}
	}
	return "utf-8", r
}
