/*
	包的作用很简单,实现5118素材接口
*/
package sucai

// 文章的一些基本信息
type ArticleInfo struct {
	GUID                 string `json:"guid"`
	Title                string `json:"title"`
	IssueTime            string `json:"issueTime"`
	CharsCount           int    `json:"charsCount"`
	ReadCount            int    `json:"readCount"`
	LikeCount            int    `json:"likeCount"`
	IsKOL                int    `json:"isKOL"`
	IsOriginal           int    `json:"isOriginal"`
	CoverImage           string `json:"coverImage"`
	AddTime              string `json:"addTime"`
	Intro                string `json:"intro"`
	Name                 string `json:"name"`
	HeadImage            string `json:"headImage"`
	ResourcePlatformName string `json:"resourcePlatformName"`
	CatalogName          string `json:"catalogName"`
	URL                  string `json:"url"`
}

// response返回的具体信息
type ArticleListData struct {
	PageCount       int           `json:"page_count"`
	PageIndex       int           `json:"page_index"`
	PageSize        int           `json:"page_size"`
	Total           int           `json:"total"`
	ArticleInfoList []ArticleInfo `json:"data"`
}

// 文章详情
type ArticleDetail struct {
	GUID                 string `json:"guid"`
	Title                string `json:"title"`
	IssueTime            string `json:"issueTime"`
	CharsCount           int    `json:"charsCount"`
	ReadCount            int    `json:"readCount"`
	LikeCount            int    `json:"likeCount"`
	IsKOL                int    `json:"isKOL"`
	IsOriginal           int    `json:"isOriginal"`
	CoverImage           string `json:"coverImage"`
	AddTime              string `json:"addTime"`
	Name                 string `json:"name"`
	HeadImage            string `json:"headImage"`
	ResourcePlatformName string `json:"resourcePlatformName"`
	CatalogName          string `json:"catalogName"`
	ArticleContent       string `json:"articleContent"`
	URL                  string `json:"url"`
	ForwardCount         int    `json:"forwardCount"`
	CommentCount         int    `json:"commentCount"`
	EnTitle              string `json:"enTitle"`
	ChTitle              string `json:"chTitle"`
}
