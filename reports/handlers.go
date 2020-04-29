package reports

import (
	"bytes"
	"errors"
	"github.com/breathbath/erplyapi/auth"
	"github.com/breathbath/erplyapi/graph"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type ReportsProvider interface {
	VisitsByHour(fromTo FromTo, erplyID string) ([]graph.KeyValue, error)
	VisitsByDay(fromTo FromTo, erplyID string) ([]graph.KeyValue, error)
	VisitsByMonth(fromTo FromTo, erplyID string) ([]graph.KeyValue, error)
	VisitsByLocation(fromTo FromTo, erplyID string) ([]graph.KeyValue, error)
}

type ReportsHandler struct {
	ReportsProvider ReportsProvider
}

/**
@apiDefine ReportInput
@apiParam {string="json","html"} format=json Defines the output format
@apiParam {string} [to] Defines to date e.g. 2020-04-27T00:00:00, if empty now date is used
@apiParam {string} token JWT token to auth access to the reports
*/

/**
@api {get} /reports/visits-by-hour.:format?from=:from&to=:to Reports visits by hour
@apiName Visits by hour
@apiGroup Reports
@apiDescription Get visits by hour

@apiUse ReportInput
@apiParam {string} [from] Defines from date e.g. 2020-04-26T00:00:00, if empty now date -24 hours from now is used

@apiUse AuthHeader

@apiExample {String} Json
		/reports/visits-by-hour.json?from=2020-01-01T00:00&to=2020-01-01T23:00
@apiExample {String} Html
		/reports/visits-by-hour.html?token=c5gJGJefdePhtuzVTC9oySEQpYW2D3p77tloMBR&from=2020-01-01T00:00&to=2020-01-01T23:00
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

	code, err := vsh.extractErplyIDFromSession(c, format)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	data, err := vsh.ReportsProvider.VisitsByHour(ft, code)
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
				Name: "Visits",
			},
		},
		GraphName: "Visits by hour",
		XName:     "Hour",
		YName:     "Count",
	})
}

/**
@api {get} /reports/visits-by-day.:format?from=:from&to=:to Reports visits by day
@apiName Visits by day
@apiGroup Reports
@apiDescription Get visits by day

@apiUse ReportInput
@apiParam {string} [from] Defines from date e.g. 2020-04-26T00:00:00, if empty now date -7 days from now is used

@apiUse AuthHeader

@apiExample {String} Json
		/reports/visits-by-day.json?from=2020-01-01T00:00&to=2020-01-30T23:59
@apiExample {String} Html
		/reports/visits-by-day.html?token=c5gJGJefdePhtuzVTC9oySEQpYW2D3p77tloMBR&from=2020-01-01T00:00&to=2020-01-30T23:59
@apiSuccessExample Success-Response {json}
HTTP/1.1 200 OK
{
	"data": [
		{
			"key": "27-04-2020",
			"value": 6
		},
		{
			"key": "28-04-2020",
			"value": 3
		}
	]
}
@apiPermission user
*/
func (vsh ReportsHandler) VisitsByDayHandler(c *gin.Context) {
	format := formatStrToEnum(c.Param("format"))

	ft, err := extractFromTo(c)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	code, err := vsh.extractErplyIDFromSession(c, format)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	data, err := vsh.ReportsProvider.VisitsByDay(ft, code)
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
				Name: "Visits",
			},
		},
		GraphName: "Visits by day",
		XName:     "Day",
		YName:     "Count",
	})
}

/**
@api {get} /reports/visits-by-month.:format?from=:from&to=:to Reports visits by month
@apiName Visits by month
@apiGroup Reports
@apiDescription Get visits by month

@apiUse ReportInput
@apiParam {string} [from] Defines from date e.g. 2020-04-26T00:00:00, if empty now date -1 month from now is used

@apiUse AuthHeader

@apiExample {String} Json
		/reports/visits-by-month.json?from=2020-01-01T00:00&to=2020-01-30T23:59
@apiExample {String} Html
		/reports/visits-by-month.html?token=c5gJGJefdePhtuzVTC9oySEQpYW2D3p77tloMBR&from=2020-01-01T00:00&to=2020-01-30T23:59
@apiSuccessExample Success-Response {json}
HTTP/1.1 200 OK
{
	"data": [
		{
			"key": "04-2020",
			"value": 6
		},
		{
			"key": "04-2020",
			"value": 3
		}
	]
}
@apiPermission user
*/
func (vsh ReportsHandler) VisitsByMonthHandler(c *gin.Context) {
	format := formatStrToEnum(c.Param("format"))

	ft, err := extractFromTo(c)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	code, err := vsh.extractErplyIDFromSession(c, format)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	data, err := vsh.ReportsProvider.VisitsByMonth(ft, code)
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
				Name: "Visits",
			},
		},
		GraphName: "Visits by month",
		XName:     "Month",
		YName:     "Count",
	})
}

/**
@api {get} /reports/visits-by-location.:format?from=:from&to=:to Reports visits by location
@apiName Visits by location
@apiGroup Reports
@apiDescription Get visits by location

@apiUse ReportInput
@apiParam {string} [from] Defines from date e.g. 2020-04-26T00:00:00, if empty now date -1 day from now is used

@apiUse AuthHeader

@apiExample {String} Json
		/reports/visits-by-location.json?from=2020-01-01T00:00&to=2020-01-01T23:00
@apiExample {String} Html
		/reports/visits-by-location.html?token=c5gJGJefdePhtuzVTC9oySEQpYW2D3p77tloMBR&from=2020-01-01T00:00&to=2020-01-01T23:00
@apiSuccessExample Success-Response {json}
HTTP/1.1 200 OK
{
	"data": [
		{
			"key": "Rome",
			"value": 6
		},
		{
			"key": "Berlin",
			"value": 3
		}
	]
}
@apiPermission user
*/
func (vsh ReportsHandler) VisitsByLocationHandler(c *gin.Context) {
	format := formatStrToEnum(c.Param("format"))

	ft, err := extractFromTo(c)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	code, err := vsh.extractErplyIDFromSession(c, format)
	if err != nil {
		sendError(c, err, format, http.StatusBadRequest)
		return
	}

	data, err := vsh.ReportsProvider.VisitsByLocation(ft, code)
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
				Name: "Visits",
			},
		},
		GraphName: "Visits by location",
		XName:     "Location",
		YName:     "Count",
	})
}

func (vsh ReportsHandler) extractErplyIDFromSession(c *gin.Context, format Format) (string, error) {
	sessI, found := c.Get(auth.IdentityKeyBack)
	if !found {
		return "", errors.New("invalid session context")
	}

	sess := sessI.(auth.Session)
	return sess.ErplyID, nil
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


