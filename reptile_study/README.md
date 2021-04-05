爬虫 正则  study

re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面内容，srcStr是页面内容，-1代表取全部
