package baidupc

type BaiduPCResult struct {
	Keyword             string `json:"keyword"`
	Rank                int    `json:"rank"`
	BaiduIndex          int    `json:"baidu_index"`
	PageTitle           string `json:"page_title"`
	BaiduURL            string `json:"baidu_url"`
	BidwordCompanycount int    `json:"bidword_companycount"` // 竞价公司数量
	BidwordKwc          int    `json:"bidword_kwc"`          // 竞价激烈程度
	BidwordPcpv         int    `json:"bidword_pcpv"`         //	百度PC检索量
	BidwordWisepv       int    `json:"bidword_wisepv"`       // 百度无线检索量
}
