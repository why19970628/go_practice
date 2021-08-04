package main

/*
redis.replicate_commands()
是消除redis在 master-slave 情况下, 由于在 Lua 脚本中使用了随机参数(下列 Lua 中的 `redis.call('SET', key, '', 'ex', 2*60)`)
导致 master-slave 在执行 Lua 脚本的情况下在 master 和 slave 中生成的数据不一致, 所以 Redis 报错
Redis 在 4.0 版本下引入了 redis.replicate_commands()
是将 Lua 脚本的数据操作封在一个事务中, 然后再同步 slave 的时候仅仅同步数据而不是执行 Lua 脚本
*/

const UserScript = `

    -- Cache User Info
    -- KEYS[1]: hasCouponKey "{username}-has"
    -- 返回值为1代表设置成功

	-- local couponLeft = redis.call("hset", KEYS[1], "left");
	redis.call("hset", KEYS[1], "left", ARGV[1]);
    redis.call("expire", KEYS[1], 5);
	return 1;
`

const SetHashFieldScriptWithExpire = `
    -- Cache hash Info
    -- KEYS[1]: hash key
    -- ARGV[1]: hash field
    -- ARGV[2]: hash value
    -- ARGV[3]: expire time
    -- 返回值有1代表建立key成功，并设置缓存时间, 返回2代表此key field加值成功， 否则返回过期时间
	if (redis.call("exists", KEYS[1]) == 0) then
		redis.call("hset", KEYS[1], ARGV[1], ARGV[2]);
		redis.call("expire", KEYS[1], ARGV[3]);
		return 1;
	end
	if (redis.call("hexists", KEYS[1], ARGV[1] ) == 1) then
		redis.call("hincrby", KEYS[1], ARGV[1], ARGV[2]);
		return 2;
	end
	return redis.call("pttl", KEYS[1]);
`

const SetStringKeyWithExpire = `
    -- Cache String
    -- KEYS[1]: key
    -- ARGV[1]: expire time
    -- ARGV[2]: value
    -- 返回值有1代表 新建key并设置缓存时间, 返回2代表有此key
    redis.replicate_commands()
	if (redis.call("exists", KEYS[1]) == 0) then
		redis.call("set", KEYS[1], ARGV[1]);
		redis.call("expire", KEYS[1], ARGV[2]);
		return 1;
	end
	if (redis.call("exists", KEYS[1]) == 1) then
		return 2;
	end
`

// 加锁的脚本
const LockLua = `
   	-- KEYS[1]: key
    -- ARGV[1]: expire time
	-- return:  设置锁和过期时间成功，返回 1 。 设置失败，返回 0 。
	if redis.call('SETNX', KEYS[1], ARGV[1]) == 1 or redis.call('TTL', KEYS[1]) < 0 then
		redis.call('PEXPIRE', KEYS[1], ARGV[2])
		return 1
	else
		return 0
	end
`

// 释放锁的脚本
const ReleaseLua = `
   	-- KEYS[1]: key
    -- ARGV[1]: random value 防止锁过期而错删
	-- return:  删除成功，返回 1 。 设置失败，返回 0 。
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
`

// 秒杀的脚本
const SkillLua = `
   	-- KEYS[1]: lock key
   	-- KEYS[2]: good key
    -- ARGV[1]: random value
    -- ARGV[2]: PEXPIRE time
	-- return:  设置锁和过期时间成功，返回 1 。 设置失败，返回 0 。
	if redis.call('SETNX', KEYS[1], ARGV[1]) == 1 or redis.call('TTL', KEYS[1]) < 0 then
		redis.call("decr", KEYS[2])
		redis.call('PEXPIRE', KEYS[1], ARGV[2])
		return 1
	else
		return 0
	end
`
