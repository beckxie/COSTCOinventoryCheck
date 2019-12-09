# COSTCOinventoryCheck

for loop check costco online inventory and push notifications via Instant Messaging.

# How to use:

## Edit variable in `main.go` if you need to change `product-url`:

```
urls := []string{
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-White/p/125295",              //iPhone 11 128G White
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128GB-Green/p/125301",              //iPhone 11 128G Purple
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-128G-Purple/p/125302",              //iPhone 11 128G Green
		"https://www.costco.com.tw/Electronics/Apple-Devices/iPhone/iPhone-11-Pro-256GB-Midnight-Green/p/125293", //iPhone 11 Pro 256GB Midnight Green
	}
```

## Default value:

crawler interval(sleep): 
  - 5 seconds

urls: 
  - iPhone 11 128G White
  - iPhone 11 128G Purple
  - iPhone 11 128G Green
  - iPhone 11 Pro 256GB Midnight Green

# IM Support:
- [LINE Notify](https://notify-bot.line.me/)