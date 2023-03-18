package influxdb

import (
	"context"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api/write"
)

func GenerateMockData() {
	client := influxdb2.NewClient("http://localhost:8086", "RladvYc1VpWXEtsNtacbFgBs-qAFF5Ssae6Uf17HaL5CPTnrFOWeODnpGa8-vy4XMaCplNpTkzS263gL5N9Crw==")

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("babybucks", "babybucks")
	var p *write.Point
	for i := 0; i < 100; i++ {
		a := rand.Intn(10) + 20
		b := rand.Intn(10) + 20
		p = influxdb2.NewPointWithMeasurement("bucks").
			AddField("anz", a).
			AddField("cba", b).
			AddField("goal", 60.0).
			AddTag("account_holder_name", "koustubh").
			SetTime(time.Now().AddDate(0, 0, 1*i))
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
}
