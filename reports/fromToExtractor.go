package reports

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

func extractFromTo(c *gin.Context) (FromTo, error) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
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
		ft.From.IsNull = false
		ft.From.Time = fromT
	}

	if toStr != "" {
		toT, err := ExtractDate(toStr)
		if err != nil {
			return ft, err
		}
		ft.To.IsNull = false
		ft.To.Time = toT
	}

	return ft, nil
}

func ExtractDate(rawDate string) (t time.Time, err error) {
	t, err = time.Parse("2006-01-02T15:04:05", rawDate)
	if err == nil {
		return t, nil
	}

	return
}
