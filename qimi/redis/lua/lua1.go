package main

import (
	"fmt"
	"log"
	"pingguoxueyuan/qimi/redis"
)


var secKillSHA string


func init() {
	// 让redis加载秒杀的lua脚本
	//secKillSHA = PrepareScript(SecKillScript)

	// 预热
	//preHeatKeys()
}

// 确保redis加载lua脚本，若未加载则加载
func PrepareScript(script string) string {
	// sha := sha1.Sum([]byte(script))
	rd := redis.New().Client()
	scriptsExists, err := rd.ScriptExists(script).Result()
	if err != nil {
		panic("Failed to check if script exists: " + err.Error())
	}
	if !scriptsExists[0] {
		scriptSHA, err := rd.ScriptLoad(script).Result()
		if err != nil {
			panic("Failed to load script " + script + " err: " + err.Error())
		}
		fmt.Println("Script Exists:", scriptSHA)
		return scriptSHA
	}
	return ""
}


// 执行lua脚本
func EvalSHA(sha string, args []string) (interface{}, error) {
	rd := redis.New().Client()
	val, err := rd.EvalSha(sha, args).Result()
	if err != nil {
		print("Error executing evalSHA... " + err.Error())
		return nil, err
	}
	return val, nil
}


// 尝试在redis进行原子性的秒杀操作
func CacheAtomicSecKill(userName string, sellerName string, couponName string) (int64, error) {
	// 根据sha，执行预先加载的秒杀lua脚本
	userHasCouponsKey := "1"
	couponKey := "2"
	res, err := EvalSHA(secKillSHA, []string{userHasCouponsKey, couponName, couponKey})
	if err != nil {
		return -1, RedisEvalError{}
	}
	fmt.Println(res)

	// 该lua脚本应当返回int值
	couponLeftRes, ok := res.(int64)
	if !ok {
		return -1, CouponLeftResError{res}
	}

	// 此处的-1, -2, -3 和 >=0的判断依据, 与secKillSHA变量lua脚本的返回值保持一致
	// 请看secKillSHA
	switch {
	case couponLeftRes == -1:
		return -1, UserHasCouponError{userName, couponName}
	case couponLeftRes == -2:
		return -1, NoSuchCouponError{sellerName, couponName}
	case couponLeftRes == -3:
		return -1, NoCouponLeftError{sellerName, couponName}
	case couponLeftRes == 1:  // left为0时, 就是存量为0, 那就是没抢到, 也可能原本为1, 抢完变成了0.
		return couponLeftRes, nil
	default: {
		log.Fatal("Unexpected return value.")
		return -1, CouponLeftResError{couponLeftRes}
	}

	}
}


type User struct {
	Name string
	Id int
}

func main() {
	//secKillSHA := PrepareScript(secKillScript)
	//_, err := CacheAtomicSecKill("wang", "maijia", "shangpin")
	//if err == nil {
	//	print("666")
	//	//log.Println(fmt.Sprintf("result: %d", secKillRes))
	//	//coupon := redisService.GetCoupon(paramCouponName)
	//	//交给[协程]完成数据库写入操作
	//	//SecKillChannel <- secKillMessage{claims.Username, coupon}
	//	return
	//}else {
	//	fmt.Println("err:", err)
	//}

	dbRedis := redis.New()
	//tag := "seckill"
	//err := 	dbRedis.LoadScript(tag, UserScript)
	//if err != nil {
	//	log.Fatal("redis ScriptLoad '%s' failed: %v", tag, err)
	//}
	//
	////tag = "seckill2"
	////err = 	dbRedis.LoadScript(tag, UserScript)
	////if err != nil {
	////	log.Fatal("redis ScriptLoad '%s' failed: %v", tag, err)
	////}
	//
	//dbRedis.GetScript(tag)
	//key := []string{"uid"}
	//user := &User{Name: "wang", Id: 1}
	//bt, err := json.Marshal(user)
	//res, err := dbRedis.EvalSha(tag, key, bt)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	fmt.Println("-------")

	tag1 := "hash_tag"
	err1 := dbRedis.LoadScript(tag1, SetStringKeyWithExpire)
	if err1 != nil {
		log.Fatal("redis ScriptLoad '%s' failed: %v", tag1, err1)
	}

	key2 := []string{"uid_hash_key"}


	res, err2 := dbRedis.EvalSha(tag1, key2, "argv2", 20)
	fmt.Println(res == nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(res)

}
