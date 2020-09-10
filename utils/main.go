package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
 生成指定范围的随机数
*/
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/*
强制刷新随机种子保证每次请求的
*/
func RandIntReal(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//MD5加密
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//获取客户端的IP
func GetClientIP(header http.Header) string {
	ip := "unknown"
	//这个提到最前面，作为优先级,nginx代理会获取到用户真实ip,发在这个环境变量上，必须要nginx配置这个环境变量HTTP_X_FORWARDED_FOR
	if header.Get("X-Forwarded-For") != "" {
		ip = header.Get("X-Forwarded-For")
	} else if header.Get("Remote_addr") != "" { //在nginx作为反向代理的架构中，使用REMOTE_ADDR拿到的将会是反向代理的的ip，即拿到是nginx服务器的ip地址。往往表现是一个内网ip。
		ip = header.Get("Remote_addr")
	} else if header.Get("Client_Ip") != "" { //HTTP_CLIENT_IP攻击者可以伪造一个这样的头部信息，导致获取的是攻击者随意设置的ip地址。
		ip = header.Get("Client_Ip")
	}

	if strings.Index(ip, ",") > 0 {
		arr := strings.Split(ip, ",")
		if len(arr) >= 1 {
			ip = arr[0]
		}
	}
	return ip
}

//获取IP库里的地区
func GetIP2Region(ip string, dbPath string) ip2region.IpInfo {
	//dbPath := filepath.Join(beego.AppPath, "conf", "ip2region.db")
	region, err := ip2region.New(dbPath)
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return ip2region.IpInfo{}
	}

	ip2rgn, err := region.MemorySearch(ip)
	fmt.Println(ip, err)
	return ip2rgn
}

//判断空串 返回默认值
func GetStringDefault(str string, defs string) string {
	if str == "" {
		return defs
	}
	return str
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
