package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type InfluxDbConfig struct {
	Token             string
	HostUrl           string
	Org               string
	Bucket            string
	AccountHolderName string
}

type Data struct {
	Anz  int64   `json:"anz"`
	Cba  int64   `json:"cba"`
	Goal float64 `json:"goal"`
}

type ReadResponse struct {
	Time    time.Time `json:"time"`
	Data    Data      `json:"data"`
	Savings int64     `json:"savings"`
}

func Read(q InfluxDbConfig) []ReadResponse {

	client := influxdb2.NewClient(q.HostUrl, q.Token)
	queryAPI := client.QueryAPI(q.Org)

	fluxQuery := fmt.Sprintf(`from(bucket: "%v")
	|> range(start: 0h, stop: 168h)
	|> filter(fn: (r) => r["_measurement"] == "bucks")
	|> filter(fn: (r) => r["_field"] == "anz" or r["_field"] == "cba" or r["_field"] == "goal")
	|> filter(fn: (r) => r["account_holder_name"] == "%v")
	|> yield(name: "mean")`, q.Bucket, q.AccountHolderName)

	log.Println(fluxQuery)

	result, err := queryAPI.Query(context.Background(), fluxQuery)
	if err != nil {
		panic(err)
	}

	r := []ReadResponse{}
	re := ReadResponse{}
	dataMap := make(map[time.Time]*Data)
	for result.Next() {
		if dataMap[result.Record().Time()] == nil {
			dataMap[result.Record().Time()] = &Data{}
		}
		switch field := result.Record().Field(); field {
		case "anz":
			dataMap[result.Record().Time()].Anz = result.Record().Value().(int64)
		case "cba":
			dataMap[result.Record().Time()].Cba = result.Record().Value().(int64)
		case "goal":
			dataMap[result.Record().Time()].Goal = result.Record().Value().(float64)
		}
	}

	for t, v := range dataMap {
		re.Time = t
		re.Savings = v.Anz + v.Cba
		re.Data.Goal = v.Goal
		re.Data.Anz = v.Anz
		re.Data.Cba = v.Cba
		r = append(r, re)
	}

	return r
}
