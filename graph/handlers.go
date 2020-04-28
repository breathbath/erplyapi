package graph

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type VisitsByHourProvider interface {
	VisitsByHour(fromTo FromTo, erplyID string) ([]KeyValue, error)
}

type VisitStatsHandler struct {
	VisitsByHourProvider VisitsByHourProvider
}

/**
@api {get} /graphs/visits-by-hour/:from/:to Visits by hour
@apiDescription Gets the visits stats grouped by hour from 2020-01-01T00:00 till 2020-01-02T00:00
@apiName Visit by hour
@apiGroup Visits
@apiUse JsonHeader
@apiUse AuthHeader

@apiParam {string} from Starting period
@apiParam {string} to End period

@apiSuccessExample Success-Response
HTTP/1.1 200 OK

@apiErrorExample Bad request(400)
HTTP/1.1 400 Bad request

@apiPermission registered user
*/
//CreateVisitsByHandler output graphs
func (vsh VisitStatsHandler) VisitsByHourHandler(c *gin.Context) {
	ft, err := ExtractFromTo(c)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": err.Error()})
		return
	}

	data, err := vsh.VisitsByHourProvider.VisitsByHour(ft, "100234")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
		return
	}

	buf := new(bytes.Buffer)
	err = Generate(buf, Input{
		Type: Line,
		Series: []Series{
			{
				Data: data,
				Name: "Total",
			},
		},
		GraphName: "Visits by hour",
		XName:     "Hour",
		YName:     "Count",
	})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "graph.gohtml", gin.H{
		"title":  "Visits by hour",
		"graph": template.HTML(buf.String()),
	})
}
