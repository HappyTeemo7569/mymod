package utils

import (
	"fmt"
	"github.com/HappyTeemo7569/mymod/base"
	"math/rand"
	"time"
)

//设置定时器
func SetMyTimer(s int, f func(), n int) chan int {
	time1 := time.NewTicker(time.Second * time.Duration(s))
	ch := make(chan int)

	cnt := 0

	go func() {
		for {
			select {
			case <-time1.C:
				//base.Logger.Debugf("定时器响应")
				f()
				if n > 0 {
					cnt++
					if cnt >= n {
						//base.Logger.Debugf("定时器次数到了")
						time1.Stop()
						return
					}
				}
			case <-ch:
				//base.Logger.Debugf("定时器提前停止")
				time1.Stop()
				return
			}
		}
	}()

	return ch
}

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	msg := fmt.Sprintf("%s_%v", string(result), int(time.Now().Unix()))
	//base.Logger.Debugf("生成随机字符串1：", msg)
	return msg
}

func RandString() string {
	len := 10
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	msg := fmt.Sprintf("%s_%v", string(bytes), int(time.Now().Unix()))
	//base.Logger.Debugf("生成随机字符串2：", msg)
	return msg
}

func InArray(val int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

type ParameterItem struct {
	Key   string
	Value string
}

func Lock(userId, second int, name string) bool {
	rcRpc := base.GetRedisRpc()
	defer rcRpc.Close()

	key := fmt.Sprintf("tank_redis_lock_%s_%d", name, userId)
	res, err := rcRpc.Do("SETNX", key, 1)
	if err != nil {
		base.Logger.Errorf("设置限频锁失败", err)
	}
	if res == int64(1) {
		rcRpc.Do("EXPIRE", key, second)
		return true
	}
	return false //存在，已经锁住了
}

func UnLock(userId int, name string) {
	rcRpc := base.GetRedisRpc()
	defer rcRpc.Close()
	key := fmt.Sprintf("tank_redis_lock_%s_%d", name, userId)
	rcRpc.Do("EXPIRE", key, 0)
}
