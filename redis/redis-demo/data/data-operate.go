package data

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// rdb *redis.Client
var (
	wg  sync.WaitGroup // 等待队列
	ctb = context.Background()
)

func TestRedisBase(logs klog.Logger) {
	defer wg.Done()

	rdb := Getclient(logs)

	// Example_String(rdb)
	// Example_List(rdb)
	// Example_Hash(rdb)
	// Example_Set(rdb)
	Example_SortedSet(rdb)
	// Example_HyperLogLog(rdb)
	Example_PubSub(rdb)
	// Example_CMD(rdb)
	// Example_Scan(rdb)
	// Example_Tx(rdb)
	// Example_Script(rdb)
}

func Example_String(rdb *redis.Client) {
	log.Println("Client_String Running.")
	defer log.Println("Client_String Done.")
	// kv 读写
	err := rdb.Set(context.Background(), "key", "value", 2*time.Minute).Err()
	if err != nil {
		log.Println(err)
	}

	// 获取过期时间
	outTime, err := rdb.TTL(context.Background(), "10000").Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("outtime is ", outTime)
	}

	// Get Value
	val, err := rdb.Get(context.Background(), "key").Result()
	log.Println(val, err)

	// Get missing key
	val2, err := rdb.Get(ctb, "missing_key").Result()
	if err == redis.Nil {
		log.Println("missing_key does not exist")
	} else if err != nil {
		log.Println("missing_key", val2, err)
	}

	// 设置过期时间 nx ex (如果不存在 才设置)
	value, err := rdb.SetNX(ctb, "counter", 0, 1*time.Second).Result()
	log.Println("setnx", value, err)

	// Incr
	res, err := rdb.Incr(ctb, "counter").Result()
	log.Println("Incr", res, err)
}

func Example_List(rdb *redis.Client) {
	log.Println("Client_List Running.")
	defer log.Println("Client_List Done.")

	// 添加
	err := rdb.RPush(ctb, "list_test", "message1").Err()
	log.Println(err)
	log.Println(rdb.RPush(ctb, "list_test", "message2").Err())

	// 设置
	log.Println(rdb.LSet(ctb, "list_test", 2, "message set").Err())

	// remove
	res, err := rdb.LRem(ctb, "list_test", 3, "message1").Result()
	log.Println(res, err)

	rLen, err := rdb.LLen(ctb, "list_test").Result()
	log.Println(rLen, err)

	// 遍历
	lists, err := rdb.LRange(ctb, "list_test", 0, rLen-1).Result()
	log.Println("LRange", lists, err)

	// pop 没有时阻塞
	result, err := rdb.BLPop(ctb, 1*time.Second, "list_test").Result()
	log.Println("result", result, err, len(result))
}

func Example_Hash(rdb *redis.Client) {
	log.Println("Client_Hash Running.")
	defer log.Println("Client_Hash Done.")
	datas := map[string]interface{}{
		"name": "YANG FAN",
		"sex":  1,
		"age":  30,
		"tel":  337845818,
	}

	// 添加 HSET是一个Field-Value HMSET是多个fv
	if err := rdb.HMSet(ctb, "hash_test", datas); err != nil {
		log.Fatal(err)
	}

	// 获取 HGET根据key 获取单个字段的值 HMGET 根据key 获取多个字段的值
	ress, err := rdb.HMGet(ctb, "hash_test", "name", "sex").Result()
	log.Println("ress:", ress, err)

	// 获取Hash表指定key的所有字段和值
	resAll, err := rdb.HGetAll(ctb, "hash_test").Result()
	log.Println("resAll", resAll, err)

	// key中的字段是否存在值
	hExist, err := rdb.HExists(ctb, "hash_test", "tel").Result()
	log.Println(hExist, err)

	hRes, err := rdb.HSetNX(ctb, "hash_test", "id", 100).Result()
	log.Println(hRes, err)

	// 删除
	log.Println(rdb.HDel(ctb, "hash_test", "age").Result())
}

func Example_Set(rdb *redis.Client) {
	log.Println("Client_Set Running.")
	defer log.Println("Client_Set Done.")

	// 添加Set
	res, err := rdb.SAdd(ctb, "set_test", "11", "22", "33", "44").Result()
	log.Println(res, err)

	// 数量
	count, err := rdb.SCard(ctb, "set_test").Result()
	log.Println(count, err)

	// 删除
	res, err = rdb.SRem(ctb, "set_test", "11", "22").Result()
	log.Println(res, err)

	// 成员
	member, err := rdb.SMembers(ctb, "set_test").Result()
	log.Println(member, err)

	bres, err := rdb.SIsMember(ctb, "set_test", "33").Result()
	log.Println(bres, err)

	rdb.SAdd(ctb, "set_test", "11", "22", "33", "44")
	rdb.SAdd(ctb, "set_test", "11", "22", "33", "55", "66", "77")
	// 差集
	diff, err := rdb.SInter(ctb, "set_a", "set_b").Result()
	log.Println(diff, err)

	// 交集
	inter, err := rdb.SInter(ctb, "set_a", "set_b").Result()
	log.Println(inter, err)

	// 并集
	union, err := rdb.SUnion(ctb, "set_a", "set_b").Result()
	log.Println(union, err)

	res, err = rdb.SDiffStore(ctb, "set_diff", "set_a", "set_b").Result()
	log.Println(res, err)

	ress, err := rdb.SMembers(ctb, "set_diff").Result()
	log.Println(ress, err)
}

func Example_SortedSet(rdb *redis.Client) {
	log.Println("Client_SortedSet Running.")
	defer log.Println("Client_SortedSet Done.")

	addArgs := make([]*redis.Z, 100)
	for i := 1; i < 100; i++ {
		addArgs = append(addArgs, &redis.Z{Score: float64(i), Member: fmt.Sprintf("a_%d", i)})
	}
	// log.Println(addArgs)

	Shuffle := func(slice []*redis.Z) {
		if len(slice) <= 1 {
			return // 长度为0或1时无需洗牌
		}
		r := rand.New(rand.NewSource(time.Now().Unix()))
		// for len(slice) > 0 {
		// 	n := len(slice)
		// 	randIndex := r.Intn(n)
		// 	slice[n-1], slice[randIndex] = slice[n-1], slice[randIndex]
		// 	slice = slice[:n-1]
		// }
		for i := len(slice) - 1; i > 0; i-- {
			j := r.Intn(i + 1)
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	// 随机打乱
	Shuffle(addArgs)

	// 添加Zset
	res, err := rdb.ZAddNX(ctb, "sortset_test", addArgs...).Result()
	log.Println(res, err)

	// 获取指定 Member 的 score
	score, err := rdb.ZScore(ctb, "sortset_test", "a_10").Result()
	log.Println(score, err)

	// 获取指定 member 的index
	index, err := rdb.ZRank(ctb, "sortset_test", "a_50").Result()
	log.Println(index, err)

	count, err := rdb.SCard(ctb, "sortset_test").Result()
	log.Println(count, err)

	// 返回有序集合指定区间内的成员
	ress, err := rdb.ZRange(ctb, "sortset_test", 10, 20).Result()
	log.Println(ress, err)

	// 返回有序集合指定区间内的成员score从高到低
	ress, err = rdb.ZRevRange(ctb, "sortset_test", 10, 20).Result()
	log.Println(ress, err)

	// 指定score区间的member list
	ress, err = rdb.ZRangeByScore(ctb, "sortset_test", &redis.ZRangeBy{Min: "(30", Max: "(50", Offset: 1, Count: 10}).Result()
	log.Println(ress, err)
}

// 用来做基数统计的算法，HyperLogLog 的优点是，在输入元素的数量或者体积非常非常大时，
// 计算基数所需的空间总是*固定*的，并且是很小的。
// 每个 HyperLogLog 键只需要花费 12 KB 内存，就可以计算接近 2^64 个不同元素的基数
// 适合每月的每日签到情况
func Example_HyperLogLog(rdb *redis.Client) {
	log.Println("Client_HyperLogLog Running.")
	defer log.Println("Client_HyperLogLog Done.")

	for i := 0; i < 10000; i++ {
		rdb.PFAdd(ctb, "pf_test_1", fmt.Sprintf("pfkey%d", i))
	}
	res, err := rdb.PFCount(ctb, "pf_test_1").Result()
	log.Println(res, err)

	for i := 0; i < 10000; i++ {
		rdb.PFAdd(ctb, "pf_test_2", fmt.Sprintf("pfkey%d", i))
	}
	res, err = rdb.PFCount(ctb, "pf_test_2").Result()
	log.Println(res, err)

	rdb.PFMerge(ctb, "pf_test", "pf_test_1", "pf_test_2")
	res, err = rdb.PFCount(ctb, "pf_test").Result()
	log.Println(res, err)
}

// 订阅 发布
func Example_PubSub(rdb *redis.Client) {
	log.Println("Client_PubSub Running.")
	defer log.Println("Client_PubSub Done.")

	// 发布订阅
	pubsub := rdb.Subscribe(ctb, "subkey")
	_, err := pubsub.Receive(ctb)
	if err != nil {
		log.Fatal("pubsub.Receive")
	}
	ch := pubsub.Channel()
	time.AfterFunc(1*time.Second, func() {
		log.Println("Publish")
		err = rdb.Publish(ctb, "subkey", "test publish 1").Err()
		if err != nil {
			log.Fatal("rdb.Publish", err)
		}
		rdb.Publish(ctb, "subkey", "test publish 2")
	})
	for msg := range ch {
		log.Println("recv channel:", msg.Channel, msg.Pattern, msg.Payload)
	}
}

func Example_CMD(rdb *redis.Client) {
	log.Println("Client_CMD Running.")
	defer log.Println("Client_CMD Done.")

	// 执行自定义redis命令
	Get := func(ctx context.Context, rdb *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd(ctb, "get", key)
		if err := rdb.Process(ctb, cmd); err != nil {
			log.Fatal(err)
		}
		return cmd
	}

	v, err := Get(ctb, rdb, "NewStringCmd").Result()
	log.Println("NewStringCmd", v, err)

	v = rdb.Do(ctb, "get", "rdb.do").String() // no v,err ?
	log.Println("rdb.Do", v, err)
}

func Example_Scan(rdb *redis.Client) {
	log.Println("Client_Scan Running.")
	defer log.Println("Client_Scan Done.")

	// scan
	for i := 0; i < 1000; i++ {
		rdb.Set(ctb, fmt.Sprintf("skey_%d", i), i, 0)
	}

	cursor := uint64(0)
	for {
		keys, resCursor, err := rdb.Scan(ctb, cursor, "skey_*", int64(100)).Result()
		log.Println(keys, cursor, err)
		cursor = resCursor
		if cursor == 0 {
			break
		}
	}
}

func Example_Tx(rdb *redis.Client) {
	log.Println("Client_Tx Running.")
	defer log.Println("Client_Tx Done.")

	pipe := rdb.TxPipeline()
	incr := pipe.Incr(ctb, "tx_pipeline_counter")
	pipe.Expire(ctb, "tx_pipeline_counter", time.Hour)

	// Execute
	//
	//     MULTI
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//     EXEC
	//
	// using one rdb-server roundtrip.
	_, err := pipe.Exec(ctb)
	fmt.Println(incr.Val(), err)
}

func Example_Script(rdb *redis.Client) {
	log.Println("Client_Script Running.")
	defer log.Println("Client_Script Done.")

	IncrByXX := redis.NewScript(`
        if redis.call("GET", KEYS[1]) ~= false then
            return redis.call("INCRBY", KEYS[1], ARGV[1])
        end
        return false
    `)

	n, err := IncrByXX.Run(ctb, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctb, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(ctb, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)
}
