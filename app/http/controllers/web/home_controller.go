package web

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/lichmaker/boring-lottery/app/models/prediction"
	"github.com/lichmaker/boring-lottery/app/models/results"
)

type HomeController struct {
}

func (hc *HomeController) Home(cont *gin.Context) {
	d := make([]struct {
		DateTime string
		Blue     uint64
		Red      string
	}, 0, 10)
	query, _ := prediction.Get(20)
	for _, v := range query {
		redGroup := []int{
			int(v.Red1),
			int(v.Red2),
			int(v.Red3),
			int(v.Red4),
			int(v.Red5),
			int(v.Red6),
		}
		sort.IntSlice(redGroup).Sort()

		buildRed := fmt.Sprintf("%0d,%0d,%0d,%0d,%0d,%0d", redGroup[0], redGroup[1], redGroup[2], redGroup[3], redGroup[4], redGroup[5])
		d = append(d, struct {
			DateTime string
			Blue     uint64
			Red      string
		}{
			v.CreatedAt.Format("2006-01-02 15:04:05"),
			v.Blue,
			buildRed,
		})
	}

	cont.HTML(http.StatusOK, "index.tmpl", gin.H{
		"arr": d,
	})
}

func (hc *HomeController) Results(cont *gin.Context) {
	d := make([]struct {
		Period string
		Blue   uint64
		Red    string
	}, 0, 10)
	query, _ := results.GetLast10Days()
	for _, v := range query {
		buildRed := fmt.Sprintf("%0d,%0d,%0d,%0d,%0d,%0d", v.Red1, v.Red2, v.Red3, v.Red4, v.Red5, v.Red6)
		d = append(d, struct {
			Period string
			Blue   uint64
			Red    string
		}{
			v.Period,
			v.Blue,
			buildRed,
		})
	}

	cont.HTML(http.StatusOK, "results.tmpl", gin.H{
		"arr": d,
	})
}
