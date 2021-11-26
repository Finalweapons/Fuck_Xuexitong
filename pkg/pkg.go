package pkg

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
)

func Get(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return string(body)
}
func GetQrcode() (qrcode string, enid []string) {
	doc, err := htmlquery.LoadURL("https://passport2.chaoxing.com/login?fid=&newversion=true&refer=http://i.chaoxing.com")
	if err != nil {
		fmt.Println(err)
	}
	enc := htmlquery.FindOne(doc, "/html/body/div/div/div[1]/div[3]/div[1]/div/input[1]")
	enid[0] = htmlquery.SelectAttr(enc, "value")
	uuid := htmlquery.FindOne(doc, "/html/body/div/div/div[1]/div[3]/div[1]/div/input[2]")
	enid[1] = htmlquery.SelectAttr(uuid, "value")
	qrcode = "https://passport2.chaoxing.com/createqr?uuid=" + enid[1] + "&fid=-1"
	return
}
