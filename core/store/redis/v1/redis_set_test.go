package redis

//func TestNewRedisSet(t *testing.T) {
//	key, _ := rds.NewRedisKey("business", "set_test_key1")
//
//	key1 := "123"
//	key2 := "456"
//	insertCount, _ := rds.SAdd(key, key1, key2)
//	assert.Equal(t, int64(2), insertCount)
//
//	members, _ := rds.SMembers(key)
//	assert.Equal(t, 2, len(members))
//
//	ok, _ := rds.SIsMember(key, key1)
//	assert.Equal(t, true, ok)
//
//	_, _ = rds.Del(key)
//}
