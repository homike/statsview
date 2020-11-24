package statsview

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// GoroutinesViewer collects the goroutine number metric via `runtime.NumGoroutine()`
type BasicViewer struct {
	ViewName     string
	graph        *charts.Line
	valueBuilder ValueBuilderFunc
}

// NewGoroutinesViewer returns the GoroutinesViewer instance
// Series: Goroutines
func NewBasicViewer(name string, Series []string, b ValueBuilderFunc) Viewer {
	graph := newBasicView(name)
	graph.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{Name: "Num"}),
		charts.WithTitleOpts(opts.Title{Title: name}),
	)

	if Series == nil || len(Series) == 0 { // one line
		graph.AddSeries(name, []opts.LineData{})
	} else { // multi line
		for _, v := range Series {
			graph.AddSeries(v, []opts.LineData{})
		}
	}

	return &BasicViewer{
		ViewName:     name,
		graph:        graph,
		valueBuilder: b,
	}
}

func (b *BasicViewer) Name() string {
	return b.ViewName
}

func (b *BasicViewer) View() *charts.Line {
	return b.graph
}

func (b *BasicViewer) Serve(w http.ResponseWriter, _ *http.Request) {
	metrics := Metrics{
		Values: b.valueBuilder(),
		Time:   time.Now().Format(defaultCfg.TimeFormat),
	}

	bs, _ := json.Marshal(metrics)
	w.Write(bs)
}
