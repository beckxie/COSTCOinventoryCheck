package notify

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PushTelegram : push message to telegram chatroom
// describe: https://api.telegram.org/bot<token>/sendMessage
// input1 httpProxy: null (option)
//
// input2 telegramToken: telegram token (require).Reference:https://core.telegram.org/bots/api#authorizing-your-bot
//
// input3 chatID: telegram chat id  (require)
//
// input4 msg: a message. (require)
func PushTelegram(httpProxy string, telegramToken string, chatID string, msg string) {
	urlmsg := url.QueryEscape(msg)
	parseMode := "Markdown"
	telegramAPIURL := "https://api.telegram.org/bot" + telegramToken + "/sendMessage?parse_mode=" + parseMode + "&chat_id=" + chatID + "&text=" + urlmsg

	req, err := http.NewRequest("POST", telegramAPIURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("cache-control", "no-cache")

	client := &http.Client{}
	if len(httpProxy) > 0 {
		proxyURL, _ := url.Parse(httpProxy)

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		client.Transport = &http.Transport{}
	}
	fmt.Println("telegram:", telegramAPIURL)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do(req) error:")
		fmt.Println(err)
		return
	} else {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("push notify result:", res.StatusCode)
	if res.StatusCode != 200 {
		fmt.Println(string(body))
	}
}

func PushToLine(msg string, lineToken string) {
	lineNotifyAPIURL := "https://notify-api.line.me/api/notify"
	authorization := "Bearer " + lineToken

	if len(msg) <= 0 || len(lineToken) <= 0 {
		fmt.Println("msg<=0 sendToLine:" + lineToken)
		logstr := "msg is null or lineToken is null"
		panic(logstr)
	}

	v := url.Values{}
	v.Set("message", msg)
	payload := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, _ := http.NewRequest("POST", lineNotifyAPIURL, payload)

	req.Header.Add("authorization", authorization)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
