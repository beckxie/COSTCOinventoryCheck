package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"log"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beckxie/go-utils/notify"
)

var wg sync.WaitGroup

const (
	version               = "1.0.1"
	intervalSecondDefault = 5
	dateTimeLayout        = "2006-01-02 15:04:05.99"
)

var (
	//product-url
	urls = []string{
		// "https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-White/p/125295",                                       //iPhone 11 128G White
		// "https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-Green/p/125301",                                       //iPhone 11 128G Purple
		// "https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128G-Purple/p/125302",                                       //iPhone 11 128G Green
		// "https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-Pro-256GB-Midnight-Green/p/125293",                          //iPhone 11 Pro 256GB Midnight Green
		"https://www.costco.com.tw/Electronics/Apple-Devices/Apple-TV-Audio-Accessories/AirPods-Pro-with-Wireless-Charging-Case/p/125773", //AirPods Pro with Wireless Charging Case
	}
)

func main() {
	var showVersion bool
	var intervalSecond int
	//LINE Notify Token
	var lineToken string

	flag.BoolVar(&showVersion, "v", false, "version")
	flag.StringVar(&lineToken, "t", "", "(require) line notify token")
	flag.IntVar(&intervalSecond, "i", intervalSecondDefault, "Crawler interval(sec)")
	flag.Parse()

	switch {
	case showVersion:
		fmt.Println("version:" + version)
		return
	case len(lineToken) <= 0:
		fmt.Println("error: Line Notify Token is required.")
		return
	case intervalSecond <= 0:
		fmt.Println("warn: scan interval error.")
		intervalSecond = intervalSecondDefault
	}

	fmt.Printf("Crawler interval(sec): %d \n", intervalSecond)

	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				msg := "Costco crawler: interrupt"
				notify.PushToLine(msg, lineToken)
				fmt.Println(msg)
				os.Exit(0)
			}
		}
	}()

	for {
		for _, url := range urls {
			wg.Add(1)
			go crawler(url, lineToken)
		}
		wg.Wait()
		fmt.Println("finish.", time.Now().Format(dateTimeLayout))
		time.Sleep(time.Duration(intervalSecond) * time.Second)
	}
}

func crawler(url string, lineToken string) {
	defer wg.Done()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div.js-qty-selector").Each(func(i int, s *goquery.Selection) {

		const NotifyMsg = `
Hurry while it's still in stock.
Url:`

		msg := time.Now().Format(dateTimeLayout) + NotifyMsg + url
		fmt.Println("msg:", msg)

		notify.PushToLine(msg, lineToken)
	})
}
