package boot

import (
	"fmt"
	"gf-server/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

// init 隐式初始化Redis
//func init() {
//	if g.Cfg().GetBool("system.UseMultipoint"){
//		global.GFVA_REDIS = g.Redis()
//	}
//	ping()
//}

func InitializeRedis() {
	if g.Cfg().GetBool("system.UseMultipoint") {
		global.GFVA_REDIS = g.Redis()
	}
	Ping()
}

func Ping() {
	conn, err := global.GFVA_REDIS.Do("PING")
	if err != nil {
		fmt.Println(err)
	}
	global.GFVA_LOG.Infof("redis connect ping response:%v", gconv.String(conn))
}
