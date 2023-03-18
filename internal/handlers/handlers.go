package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	i "github.com/koustubh25/babybucks/internal/influxdb"
)

func FetchBalancesAndGoals(c *gin.Context) {
	conf := i.InfluxDbConfig{
		Org:               os.Getenv("INFLUXDB_ORG"),
		Token:             os.Getenv("INFLUXDB_TOKEN"),
		HostUrl:           os.Getenv("INFLUXDB_URL"),
		Bucket:            os.Getenv("INFLUXDB_BUCKET"),
		AccountHolderName: os.Getenv("INFLUXDB_ACCOUNT_HOLDER_NAME"),
	}
	c.JSON(http.StatusOK, i.Read(conf))
}
