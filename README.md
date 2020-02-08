# COSTCO inventory Check

For loop check `Costco Online` inventory and push notifications via Instant Messaging.

# Costco Online Support:
- [Taiwan](https://www.costco.com.tw/)

# IM Support:
- [LINE Notify](https://notify-bot.line.me/)


# How to use:
## Usage
```
Usage:
  -i int
        Crawler interval(sec): (default 5)
  -t string
        (require) line notify token
  -v    version
```

## Edit variable in `main.go` if you need to change `product-url`:

```
urls := []string{
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-White/p/125295",              //iPhone 11 128G White
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-Green/p/125301",              //iPhone 11 128G Purple
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128G-Purple/p/125302",              //iPhone 11 128G Green
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-Pro-256GB-Midnight-Green/p/125293", //iPhone 11 Pro 256GB Midnight Green
	}
```