package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	ztime "github.com/avayayu/micro/time"
	"github.com/avayayu/quant_data/internal/utils"
)

func (crawler *AnnounceCrawler) SH(page uint16) (announce *SHAnnouce, err error) {

	req, jsonp := getSHReq(ztime.Now().Date(), ztime.Now().Date(), page)

	if err != nil {
		fmt.Println(err)
		return
	}
	var res *http.Response
	client := http.Client{}
	if res, err = client.Do(req); err != nil {
		return
	} else {
		var buf bytes.Buffer
		buf.ReadFrom(res.Body)
		data := buf.String()

		realdata := strings.ReplaceAll(data, fmt.Sprintf("%s(", jsonp), "")
		realdata = realdata[0 : len(realdata)-1]

		announce = &SHAnnouce{}

		if err = json.Unmarshal([]byte(realdata), &announce); err != nil {
			return nil, err
		}

		fmt.Println()

	}
	return
}

func getSHReq(startDate, endDate string, page uint16) (*http.Request, string) {
	today := ztime.Now()
	fmt.Println(today)

	param := map[string]interface{}{}
	param["jsonCallBack"] = fmt.Sprintf("jsonpCallback%d", rand.Int31())
	param["isPagination"] = true
	param["productId"] = ""
	param["keyWord"] = ""
	param["productId"] = ""
	param["securityType"] = "0101,120100,020100,020200,120200"
	param["reportType2"] = ""
	param["reportType"] = "ALL"
	param["beginDate"] = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	param["endDate"] = today.Date()
	param["pageHelp.pageSize"] = "2000"
	param["pageHelp.pageCount"] = "1"
	param["pageHelp.pageNo"] = fmt.Sprintf("%d", page)
	param["pageHelp.beginPage"] = "1"
	param["pageHelp.cacheSize"] = "1"
	param["pageHelp.endPage"] = "1000"
	param["_"] = fmt.Sprintf("%d", time.Now().Unix())
	url := "http://query.sse.com.cn/security/stock/queryCompanyBulletin.do"
	url = fmt.Sprintf("%s%s", url, utils.ParseToStr(param))

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Referer", "http://www.sse.com.cn/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6,la;q=0.5")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Host", "query.sse.com.cn")

	return req, param["jsonCallBack"].(string)
}

func (crawler *AnnounceCrawler) crawlSHTotal(group *sync.WaitGroup, data chan map[string]map[string]string) {
	defer group.Done()
	result := []Result{}
	annouceBegin, err := crawler.SH(1)
	if err != nil || len(annouceBegin.Results) == 0 {
		return
	}
	result = append(result, annouceBegin.Results...)
	pageDesc := annouceBegin.PageHelp
	for pageDesc.PageNo*pageDesc.PageSize < pageDesc.Total {
		num := pageDesc.PageNo + 1
		annouceBegin, err = crawler.SH(num)
		if err != nil || len(annouceBegin.Results) == 0 {
			return
		}
		result = append(result, annouceBegin.Results...)
	}
}
