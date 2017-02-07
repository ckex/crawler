package process

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"fmt"
	"bytes"
	"golang.org/x/net/html"
	"strings"
	"regexp"
	"testing"
)

//  go  test -run="Test_Process"
func Test_Process(t *testing.T) {
	file := "./0a30e0e00f41b32e16cd23a1cb522060.html"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	docs, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		fmt.Println(err)
	}
	dom := docs.Find("script")
	for index, val := range dom.Nodes {
		text := Nodetext(val)
		if strings.Contains(text, "carCompareJson") {
			reg := regexp.MustCompile("\\[\\[\\[.+\\]\\]\\]")
			jsonStr := reg.FindStringSubmatch(text)
			fmt.Println(index, "  ------------------------------------ ", jsonStr)

		}
	}
	//goquery.NewDocument("http://sports.sina.com.cn")
}

func Nodetext(node *html.Node) string {
	if node.Type == html.TextNode {
		// Keep newlines and spaces, like jQuery
		return node.Data
	} else if node.FirstChild != nil {
		var buf bytes.Buffer
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			buf.WriteString(Nodetext(c))
		}
		return buf.String()
	}

	return ""
}