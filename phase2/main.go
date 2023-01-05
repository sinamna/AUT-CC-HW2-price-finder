package main

import (
	"fmt"
	"github.com/ccloud/phase2/api"
	"github.com/ccloud/phase2/pkg/config"
	"github.com/ccloud/phase2/pkg/redis"
	"net/http"
)

func main() {
	fmt.Println("loading config...")
	config.LoadConfig()
	fmt.Printf("%+v\n", config.Conf)
	fmt.Println("connecting to redis...")
	redis.SetupRedis()

	server := api.CreateHttpServer()
	server.RegisterHandler("/currency", http.MethodPost, api.PriceHandler)

	if err := server.StartServer(fmt.Sprintf("%d", config.Conf.Server.Port)); err != nil {
		panic(err)
	}
}
