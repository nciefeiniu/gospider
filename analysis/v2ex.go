package analysis

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func AnalysisHot(body string) ([] map[string] string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	result := make([]map[string]string, 10)
	doc.Find("span.item_title").Each(func(index int, selection *goquery.Selection) {
		localTitle := selection.Find("a").Text()
		localUrl, hasAttr := selection.Find("a").Attr("href")
		if !hasAttr {
			fmt.Println(err, "Can not fin href")
		} else {
			titleUrlMap := make(map[string]string)
			titleUrlMap["https://www.v2ex.com"+localUrl] = localTitle
			result = append(result, titleUrlMap)
		}
	})

	return result, nil
}