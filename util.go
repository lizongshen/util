package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/lizongshen/logger"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// DecodeToGBK 转换成GBK
func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

// Substr 截取字符串
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

var lock sync.RWMutex

// Store 持久化对象
func Store(data interface{}, filename string) {
	lock.Lock()
	defer lock.Unlock()

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		logger.Error(err)
		return
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		logger.Error(err)
		return
	}
}

// Load 加载数据
func Load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error(err)
		return
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		logger.Error(err)
		return
	}
}

// RemoveSlice 删除切片中的某个元素
func RemoveSlice(slice []interface{}, elem interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveSlice(slice, elem)
		}
	}
	return slice
}

// RandomUA 随机获取UA
func RandomUA() string {
	userAgent := [...]string{
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		"Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;Alibaba.Security.Heimdall.5448812)",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0)",
		"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	}

	return userAgent[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(userAgent))]
}

func Dial(ip string, port int) bool {
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// VerifyProxyIP will use given ip and port as proxy, if the proxy is available, return true, otherwise return false.
func VerifyProxyIP(url string) bool {
	ok := true

	resp, _, errs := gorequest.New().
		Proxy(url).
		Get("http://httpbin.org/get").
		Timeout(time.Second * 6).
		End()

	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			logger.Debug(err)
		}
		ok = false
	}

	if resp == nil || resp.StatusCode != 200 {
		ok = false
	}

	return ok
}

// IsIP will match the given parameter is ip address or not.
func IsIP(ip string) bool {
	return IsInputMatchRegex(ip,
		"^((?:(?:25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(?:25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d))))")
}

// IsInputMatchRegex will verify the input string is match the regex or not.
// This function will recover the panic if regex can't be parsed.
func IsInputMatchRegex(input, regex string) bool {
	result := false
	reg := regexp.MustCompile(regex)
	result = reg.MatchString(input)

	defer func() {
		r := recover()
		if r != nil {
			result = false
			fmt.Println(r)
		}
	}()

	return result
}

func NodeHtml(n *html.Node) string {
	var buf = bytes.NewBuffer([]byte{})
	html.Render(buf, n)
	return buf.String()
}

func NodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		// Keep newlines and spaces, like jQuery
		return node.Data
	} else if node.FirstChild != nil {
		var buf bytes.Buffer
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			buf.WriteString(NodeText(c))
		}
		return buf.String()
	}

	return ""
}

// 创建目录
func Mkdir(Path string) {
	p, _ := path.Split(Path)
	if p == "" {
		return
	}
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		if err = os.MkdirAll(p, 0777); err != nil {
			logger.Errorf("创建路径失败[%v]: %v\n", Path, err)
		}
	}
}

// The GetWDPath gets the work directory path.
func GetWDPath() string {
	wd := os.Getenv("GOPATH")
	if wd == "" {
		panic("GOPATH is not setted in env.")
	}
	return wd
}

// The IsDirExists judges path is directory or not.
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

// The IsFileExists judges path is file or not.
func IsFileExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return !fi.IsDir()
	}
}

// 遍历文件，可指定后缀
func WalkFiles(targpath string, suffixes ...string) (filelist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			filelist = append(filelist, retpath)
			return nil
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(retpath, suffix) {
				filelist = append(filelist, retpath)
			}
		}
		return nil
	})

	if err != nil {
		logger.Errorf("util.WalkFiles: %v\n", err)
		return
	}

	return
}

// 遍历目录，可指定后缀
func WalkDir(targpath string, suffixes ...string) (dirlist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			dirlist = append(dirlist, retpath)
			return nil
		}
		for _, suffix := range suffixes {
			if strings.HasSuffix(retpath, suffix) {
				dirlist = append(dirlist, retpath)
			}
		}
		return nil
	})

	if err != nil {
		logger.Errorf("util.WalkDir: %v\n", err)
		return
	}

	return
}

// 遍历文件，可指定后缀，返回相对路径
func WalkRelFiles(targpath string, suffixes ...string) (filelist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			filelist = append(filelist, RelPath(retpath))
			return nil
		}
		_retpath := RelPath(retpath)
		for _, suffix := range suffixes {
			if strings.HasSuffix(_retpath, suffix) {
				filelist = append(filelist, _retpath)
			}
		}
		return nil
	})

	if err != nil {
		logger.Errorf("util.WalkRelFiles: %v\n", err)
		return
	}

	return
}

// 遍历目录，可指定后缀，返回相对路径
func WalkRelDir(targpath string, suffixes ...string) (dirlist []string) {
	if !filepath.IsAbs(targpath) {
		targpath, _ = filepath.Abs(targpath)
	}
	err := filepath.Walk(targpath, func(retpath string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			return nil
		}
		if len(suffixes) == 0 {
			dirlist = append(dirlist, RelPath(retpath))
			return nil
		}
		_retpath := RelPath(retpath)
		for _, suffix := range suffixes {
			if strings.HasSuffix(_retpath, suffix) {
				dirlist = append(dirlist, _retpath)
			}
		}
		return nil
	})

	if err != nil {
		logger.Errorf("util.WalkRelDir: %v\n", err)
		return
	}

	return
}

// 转相对路径
func RelPath(targpath string) string {
	basepath, _ := filepath.Abs("./")
	rel, _ := filepath.Rel(basepath, targpath)
	return strings.Replace(rel, `\`, `/`, -1)
}
