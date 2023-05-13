package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/danielgatis/go-orderbook"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
	"time"
)

func main() {
	book := orderbook.NewOrderBook("TeamA/TeamB")

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		side := []orderbook.Side{orderbook.Buy, orderbook.Sell}[rand.Intn(2)]

		book.ProcessPostOnlyOrder(uuid.New().String(), uuid.New().String(), side, decimal.NewFromInt(rand.Int63n(1000)), decimal.NewFromInt(rand.Int63n(1000)))
	}

	depth, err := json.Marshal(book.Depth())
	if err != nil {
		log.Println("Error in parsing orderbook depth")
		return
	} else {
		log.Println("Orderbook built successfully")
	}
	var buf bytes.Buffer
	json.Indent(&buf, depth, "", "  ")
	fmt.Println(buf.String())
}
