package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/itsshashank/tariff-calculator/types"
)

const endpoint = "ws://localhost:30000/ws"

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
	var data types.OBUData
	obuIDs := genobuIDs(20)
	for {
		for _, id := range obuIDs {
			lat, long := genLatLong()
			data = types.OBUData{
				OBUID: id,
				Lat:   lat,
				Long:  long,
			}
			if err = conn.WriteJSON(data); err != nil {
				log.Fatalln(err)
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func genobuIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = 1000 + rand.Intn(n)
	}
	return ids
}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}
