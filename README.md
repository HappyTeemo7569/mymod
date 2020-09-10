# mymod

个人的公共模块

## 配置
config/configBase.ini
```
[MySql] 
Auth = admin  
Pwd = 123456  
Addr = 127.0.0.1
Port = 3306
Db = test

[Redis]
Addr = admin
Pwd = 123456
Port = 6379
Rpc = 4  //暂时只有两个redis库
Api = 3

[Log]
Path = "./log" //日志路径
```

## 使用示例
```
import (
	"github.com/HappyTeemo7569/mymod/base"
	"github.com/HappyTeemo7569/mymod/utils"
)

base.Logger.Info("测试")

res := utils.RandInt(1, 10)
println(res)

redisApi := base.GetRedisApi()
defer redisApi.Close()

var vcoin_amount int
userId := 100002
err := base.Db.Get(&vcoin_amount, "select vcoin_amount from t_user_wallet where user_id = ?", userId)
if err != nil {
    base.Logger.Errorf("数据库读取失败：UserID=%d,err=%v", userId, err)
}
base.Logger.Infof("用户的钱:", vcoin_amount)
```
## webScoket
```

```