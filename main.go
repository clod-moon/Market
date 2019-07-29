package main

import (
	"fmt"
	"huobiapi/client"
	"huobiapi/market"
	"github.com/bitly/go-simplejson"
)

type JSON = simplejson.Json

type Market = market.Market
type Listener = market.Listener
type Client = client.Client

func main() {
	// 创建客户端实例
	market, err := market.NewMarket()
	if err != nil {
		panic(err)
	}
	// 订阅主题
	market.Subscribe("market.btcusdt.detail", func(topic string, json *JSON) {
		// 收到数据更新时回调
		//ret,_:= json.Encode()
		close,_ := json.Get("tick").Get("close").Float64()
		open,_ := json.Get("tick").Get("open").Float64()
		fmt.Println(topic,"open：",open," close：",close, "涨幅：",(close-open)/open*100)
	})
	// 请求数据
	//json, err := market.Request("market.eosusdt.detail")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("========>",json)
	//}

	market.Subscribe("market.ethusdt.detail", func(topic string, json * JSON) {
		// 收到数据更新时回调
		close,_ := json.Get("tick").Get("close").Float64()
		open,_ := json.Get("tick").Get("open").Float64()
		fmt.Println(topic,"open：",open," close：",close, "涨幅：",(close-open)/open*100)
	})
	// 请求数据
	//json, err = market.Request("market.eosusdt.detail")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("========>",json)
	//}
	// 进入阻塞等待，这样不会导致进程退出
	market.Loop()
}