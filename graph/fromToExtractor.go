package graph

import (
	"github.com/gin-gonic/gin"
	"time"
)

type NullTime struct {
	IsNull bool
	Time   time.Time
}

type FromTo struct {
	From NullTime
	To   NullTime
}

func ExtractFromTo(c *gin.Context) (FromTo, error) {
	fromStr := c.Param("from")
	toStr := c.Param("to")
	ft := FromTo{
		From: NullTime{
			IsNull: true,
		},
		To: NullTime{
			IsNull: true,
		},
	}
	if fromStr != "" {
		fromT, err := ExtractDate(fromStr)
		if err != nil {
			return ft, err
		}
		ft.From.IsNull = true
		ft.From.Time = fromT
	}

	if toStr != "" {
		toT, err := ExtractDate(toStr)
		if err != nil {
			return ft, err
		}
		ft.To.IsNull = true
		ft.To.Time = toT
	}

	return ft, nil
}

func ExtractDate(rawDate string) (t time.Time, err error) {
	formats := []string{"2006-01-02T15:04:05", "2006-01-02T15:04", "2006-01-02 15:04:05", "2006-01-02", "2006-01-02 15:04"}
	for _, format := range formats {
		t, err = time.Parse(rawDate, format)
		if err == nil {
			return t, nil
		}
	}

	return
}
