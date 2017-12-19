package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("參數錯誤")
	}

	year := os.Args[1]
	semester := os.Args[2]
	stdID := os.Args[3]

	if semester > "2" || semester < "1" {
		log.Fatalln("學期輸入錯誤")
	}

	url := fmt.Sprintf("http://140.131.110.236/_eportfolio/_portfolio/studentscoreview.jsp?year=%s&str=%s&id=%s", year, semester, stdID)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}

	var result [][]string
	var score []string
	var firstBlueText = true
	doc.Find("td.content_word").Each(func(i int, s *goquery.Selection) {
		value, exits := s.Attr("style")
		if exits && value == "color:blue" {
			if i%2 == 0 {
				if firstBlueText {
					result = append(result, []string{})
					firstBlueText = false
				}
				score = append(score, s.Text())
			} else if i%2 == 1 {
				score = append(score, s.Text())
				result = append(result, score)
				score = []string{}
			}
			return
		}

		if i%4 == 0 {
			score = append(score, s.Text())
		} else if i%4 == 3 {
			score = append(score, s.Text())
			result = append(result, score)
			score = []string{}
		}
	})

	if len(result) == 0 {
		fmt.Println("找不到相關資訊")
		os.Exit(0)
	}

	fmt.Printf("學年：%s\n", year)
	fmt.Printf("學期：%s\n", semester)
	fmt.Printf("學號：%s\n", stdID)
	fmt.Println("")
	for _, r := range result {
		if len(r) == 0 {
			fmt.Println("")
			continue
		}
		fmt.Println(strings.Join(r, "："))
	}
}
