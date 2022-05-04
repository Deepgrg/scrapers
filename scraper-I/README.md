# SCRAPER-I

`SCRAPER-I` is a [crawler](https://en.wikipedia.org/wiki/Crawler) made with [go programming language](https://go.dev/).

##

#### **SCRAPPER-I** crawls through a site, collects all the **anchor** `<a></a>` tags in the resource and returns all the valid link urls

### Example
```
go run main.go http://wikipedia.com http://amazon.com http://apple.com
```

```
Found 10 unique urls
-https://meta.wikimedia.org/wiki/Terms_of_use
-https://www.amazon.com/gp/help/customer/display.html/ref=footer_cou?ie=UTF8&nodeId=508088
-https://donate.wikimedia.org/?utm_medium=portal&utm_campaign=portalFooter&utm_source=portalFooter
-https://itunes.apple.com/app/apple-store/id324715238?pt=208305&ct=portal&mt=8
-https://itunes.apple.com/app/apple-store/id324715238?pt=208305&ct=portal&mt=8
-https://meta.wikimedia.org/wiki/Special:MyLanguage/List_of_Wikipedias
-https://en.wikipedia.org/wiki/List_of_Wikipedia_mobile_applications
-https://www.amazon.com/gp/help/customer/display.html/ref=footer_cou?ie=UTF8&nodeId=508088
-https://www.amazon.com/gp/help/customer/display.html/ref=footer_privacy?ie=UTF8&nodeId=468496
-https://appleid.apple.com/us/
-https://www.icloud.com
-https://investor.apple.com/
-https://meta.wikimedia.org/wiki/Privacy_policy
```

> Thank you for visiting. 💖