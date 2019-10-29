package main

import (
	"fmt"

	"log"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beckxie/go-utils/notify"
	_ "github.com/go-sql-driver/mysql"
)

var wg sync.WaitGroup

const lineToken = ""

func main() {

	urls := []string{
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-White/p/125295",              //iPhone 11 128G White
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-Green/p/125301",              //iPhone 11 128G Purple
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128G-Purple/p/125302",              //iPhone 11 128G Green
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-Pro-256GB-Midnight-Green/p/125293", //iPhone 11 Pro 256GB Midnight Green
	}

	for {
		for _, url := range urls {
			wg.Add(1)
			go crawler(url)
		}
		wg.Wait()
		fmt.Println("finish.", time.Now().Format("2006-01-02 15:04:05.99"))
		time.Sleep(5 * time.Second)
	}
}

func crawler(url string) {
	defer wg.Done()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.js-qty-selector").Each(func(i int, s *goquery.Selection) {
		msg := time.Now().Format("2006-01-02 15:04:05.99") + "\n發現有庫存.\nUrl:" + url
		fmt.Println(msg)
		notify.PushToLine(msg, lineToken)
	})
}
