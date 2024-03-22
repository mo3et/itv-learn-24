package main

import (
	"os"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/mo3et/itv-learn-24/redis/redis-demo/data"
)

type player struct {
	ID       int
	Username string
}

var (
	rdb *redis.Client
	wg  sync.WaitGroup
)

func main() {
	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller)
	// logger.SetFormatter(&nested.Formatter{
	// 	HideKeys:        false,
	// 	NoColors:        false,
	// 	ShowFullLevel:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// 	FieldsOrder:     []string{"service", "caller", "module", "msg"},
	// })
	logs := log.NewHelper(logger)
	logs.Info("logger is success running.")
	rService, err := data.NewRedisService(logger)
	if err != nil {
		rService.Logger.
		logs.Errorf("Failed to initialize Redis service: %v", err)
		return
	}
	defer rService.Close()

	// players := []player{
	// 	{1, "Alice"},
	// 	{2, "Bob"},
	// 	{3, "Charlie"},
	// 	{4, "David"},
	// 	{5, "Eve"},
	// }
	// var playerIds []int
	// for _, player := range players {
	// 	playerIds = append(playerIds, player.ID)
	// 	if err := rService.SetPlayerUsername(player.ID, player.Username); err != nil {
	// 		logs.Error("SetName error", err)
	// 	}
	// }
	// for _, id := range playerIds {
	// 	name, err := rService.GetPlayerUsername(id)
	// 	if err != nil {
	// 		logs.Error("GetName error", err)
	// 	}
	// 	logs.Infof("Get Player is %s", name)
	// }

	wg.Add(1)
	go data.TestRedisBase(logger)
	wg.Wait()
}
