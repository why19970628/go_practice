package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"pingguoxueyuan/qimi/redis"
	"sync"
	"sync/atomic"
	"time"
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
	case couponLeftRes == 1: // left为0时, 就是存量为0, 那就是没抢到, 也可能原本为1, 抢完变成了0.
		return couponLeftRes, nil
	default:
		{
			log.Fatal("Unexpected return value.")
			return -1, CouponLeftResError{couponLeftRes}
		}

	}
}

type User struct {
	Name string
	Id   int
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
	test_skill2(dbRedis)
	select {

	}
}


func test_skill2(dbRedis *redis.Redis) {
	tag1 := "skill_tag2"
	err1 := dbRedis.LoadScript(tag1, SkillLua)
	tag2 := "release_skill_tag2"
	err2 := dbRedis.LoadScript(tag2, ReleaseLua)

	if err1 != nil {
		fmt.Println("redis ScriptLoad tag1 failed:", tag1, err1)
	}
	if err2 != nil {
		fmt.Println("redis ScriptLoad tag2 failed:", tag2, err2)
	}

	lock_key := "skill_lock_key2"
	dbRedis.Client().Set(lock_key, 105, 0)

	good_key := "good_key2"
	key_slice := []string{lock_key, good_key}
	var wg sync.WaitGroup
	x := int32(0)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int, num *int32) {
			_ = fmt.Sprintf("第 %v 抢购", i)
			defer wg.Done()
			ul := uuid.NewV4()
			exists, _ := dbRedis.Client().Exists(lock_key).Result()
			if exists == int64(1){
				fmt.Println("加锁中。。。。")
				return
			}
			// 设置redis锁 并减少库存
			_, _ = dbRedis.EvalSha(tag1, key_slice, ul, 10)
			// 记录秒杀数量
			atomic.AddInt32(num, 1)
			fmt.Println("7777")
			// 删除锁
			del_res, _ := dbRedis.EvalSha(tag2, key_slice, ul)
			if del_res == 0{
				fmt.Println("锁过期")
			}
			fmt.Println("减库存成功------")
		}(i, &x)
		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()
	fmt.Println("秒杀了库存数量: ", x)
}


func test_skill(dbRedis *redis.Redis) {
	tag1 := "skill_tag"
	err1 := dbRedis.LoadScript(tag1, LockLua)
	tag2 := "release_skill_tag"
	err2 := dbRedis.LoadScript(tag2, ReleaseLua)

	if err1 != nil {
		fmt.Println("redis ScriptLoad tag1 failed:", tag1, err1)
	}
	fmt.Println("----666--")
	if err2 != nil {
		fmt.Println("redis ScriptLoad tag2 failed:", tag2, err2)
		}

	key := "skill_key"
	dbRedis.Client().Set(key, 105, 0)

	key_slice := []string{"lock_key"}
	var wg sync.WaitGroup
	x := int32(0)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int, num *int32) {
			fmt.Sprintf("第 %v 抢购", i)
			defer wg.Done()
			ul := uuid.NewV4()
			exists, _ := dbRedis.Client().Exists("lock_key").Result()
			if exists == int64(1){
				fmt.Println("加锁中。。。。")
				return
			}
			// 设置redis锁
			_, err3 := dbRedis.EvalSha(tag1, key_slice, ul, 10)
			// 减库存
			dbRedis.Client().Decr(key)
			// 记录秒杀数量
			atomic.AddInt32(num, 1)
			// 删除锁
			del_res, er4 := dbRedis.EvalSha(tag2, key_slice, ul)
			if del_res == 0{
				fmt.Println("锁过期")
			}
			if err3 != nil {
				fmt.Println("err3",err3)
			}
			if er4 != nil {
				fmt.Println("er4",er4)
			}
			fmt.Println("减库存成功------")
		}(i, &x)
		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()
	fmt.Println("秒杀了库存数量: ", x)
}

func test2(dbRedis redis.Redis)  {
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
