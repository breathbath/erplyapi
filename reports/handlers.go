package reports

import (
	"bytes"
	"github.com/breathbath/erplyapi/graph"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type ReportsProvider interface {
	VisitsByHour(fromTo FromTo, erplyID string) ([]graph.KeyValue, error)
}

type ReportsHandler struct {
	ReportsProvider ReportsProvider
}

/**
@api {get} /reports/:code/visits-by-hour.:format?from=:from&to=:to Reports visits by hour
@apiName Visits by hour
@apiGroup Reports
@apiDescription Get visits by hour

@apiParam {string} code Erply ID has to be provided
@apiParam {string="json","html"} format=json Defines the output format
@apiParam {string} [from] Defines from date e.g. 2020-04-26T00:00:00, if empty now date -10 hours from now is used
@apiParam {string} [to] Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used

@apiUse AuthHeader

@apiExample {String} Json
		/reports/visits-by-hour.json?from=2020-01-01T00:00&to=2020-01-01T23:00
@apiExample {String} Html
		/reports/visits-by-hour.html?from=2020-01-01T00:00&to=2020-01-01T23:00
@apiSuccessExample Success-Response {json}
HTTP/1.1 200 OK
{
	"data": [
		{
			"key": "27-04-2020 09:00",
			"value": 6
		},
		{
			"key": "28-04-2020 06:00",
			"value": 3
		}
	]
}
@apiPermission user
*/
func (vsh ReportsHandler) VisitsByHourHandler(c *gin.Context) {
	format := formatStrToEnum(c.Param("format"))

	ft, err := extractFromTo(c)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	data, err := vsh.ReportsProvider.VisitsByHour(ft, c.Param("code"))
	if err != nil {
		sendError(c, err, format, http.StatusInternalServerError)
		return
	}

	if format == Json {
		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}

	vsh.generateGraph(c, graph.Input{
		Type: graph.Line,
		Series: []graph.Series{
			{
				Data: data,
				Name: "Total",
			},
		},
		GraphName: "Visits by hour",
		XName:     "Hour",
		YName:     "Count",
	})
}

func (vsh ReportsHandler) generateGraph(c *gin.Context, i graph.Input) {
	buf := new(bytes.Buffer)
	err := graph.Generate(buf, i)
	if err != nil {
		sendError(c, err, Html, http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "report.gohtml", gin.H{
		"title":  i.GraphName,
		"graph": template.HTML(buf.String()),
	})
}


