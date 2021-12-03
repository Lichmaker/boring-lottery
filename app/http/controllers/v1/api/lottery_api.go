package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lichmaker/boring-lottery/app/models/results"
)

type LotteryApi struct {
}

func (action *LotteryApi) Get(cont *gin.Context) {
	cont.Header("Access-Control-Allow-Origin", "*")

	query, _ := results.GetLatest()
	// fmt.Printf("第二个%p\n", query)
	buildRed := fmt.Sprintf("%02d,%02d,%02d,%02d,%02d,%02d", query.Red1, query.Red2, query.Red3, query.Red4, query.Red5, query.Red6)
	cont.JSON(http.StatusOK, gin.H{
		"period": query.Period,
		"blue":   fmt.Sprintf("%02d", query.Blue),
		"red":    buildRed,
	})
}
