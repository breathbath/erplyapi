package graph

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"io"
)

type Series struct {
	Data []KeyValue
	Name string
}

type KeyValue struct {
	Key   string `db:"key" json:"key"`
	Value interface{} `db:"val" json:"value"`
}

type Type int

const (
	Undefined Type = iota
	Line
	Bar
)

type Input struct {
	Type      Type
	Series    []Series
	GraphName string
	XName string
	YName string
}

//https://go-echarts.github.io/go-echarts/docs/bar
func Generate(w io.Writer, i Input) error {
	if i.Type == Line {
		return generateLine(w, i)
	}
	if i.Type == Bar {
		return generateBar(w, i)
	}
	return fmt.Errorf("unsupported graph type: %v", i.Type)
}

func generateLine(w io.Writer, i Input) error {
	line := charts.NewLine()
	xPoints := make([]string, 0)
	for _, ss := range i.Series {
		yPoints := make([]interface{}, 0, len(ss.Data))
		for _, kv := range ss.Data {
			yPoints = append(yPoints, kv.Value)
			xPoints = append(xPoints, kv.Key)
		}
		line.AddYAxis(ss.Name, yPoints, charts.LabelTextOpts{Show: true, Position: "top"})
	}
	line.AddXAxis(xPoints)
	line.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	line.SetGlobalOptions(
		charts.TitleOpts{Title: i.GraphName},
		charts.YAxisOpts{Name: i.YName},
		charts.XAxisOpts{Name: i.XName},
	)

	err := line.Render(w)
	if err != nil {
		return err
	}

	return nil
}

func generateBar(w io.Writer, i Input) error {
	bar := charts.NewBar()
	xPoints := make([]string, 0)
	xPointsCoord := make(map[string]int)
	for _, ss := range i.Series {
		for _, kv := range ss.Data {
			_, ok := xPointsCoord[kv.Key]
			if !ok {
				xPoints = append(xPoints, kv.Key)
				xPointsCoord[kv.Key] = len(xPoints) - 1
			}
		}
	}
	for _, ss := range i.Series {
		yPoints := make([]interface{}, len(xPoints))
		for _, kv := range ss.Data {
			col := xPointsCoord[kv.Key]
			yPoints[col] = kv.Value
		}
		bar.AddYAxis(ss.Name, yPoints, charts.LabelTextOpts{Show: true, Position: "top"})
	}

	bar.AddXAxis(xPoints)
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: i.GraphName},
		charts.YAxisOpts{Name: i.YName},
		charts.XAxisOpts{Name: i.XName},
	)

	err := bar.Render(w)
	if err != nil {
		return err
	}

	return nil
}
