package statsview

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"text/template"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Metrics
type Metrics struct {
	Values []float64 `json:"values"`
	Time   string    `json:"time"`
}

// Viewer is the abstraction of a Graph which in charge of collecting metrics from somewhere
type Viewer interface {
	Name() string
	View() *charts.Line
	Serve(w http.ResponseWriter, _ *http.Request)
}

type statsEntity struct {
	T     string
	Stats *runtime.MemStats
}

var rtStats = &statsEntity{Stats: &runtime.MemStats{}}

func StartRTCollect() {
	runtime.ReadMemStats(rtStats.Stats)
	rtStats.T = time.Now().Format(defaultCfg.TimeFormat)
}

func genViewTemplate(vid, route string) string {
	tpl, err := template.New("view").Parse(defaultCfg.Template)
	if err != nil {
		panic("statsview: failed to parse template " + err.Error())
	}

	var c = struct {
		Interval  int
		MaxPoints int
		Addr      string
		Route     string
		ViewID    string
	}{
		Interval:  defaultCfg.Interval,
		MaxPoints: defaultCfg.MaxPoints,
		Addr:      defaultCfg.Addr,
		Route:     route,
		ViewID:    vid,
	}

	buf := bytes.Buffer{}
	if err := tpl.Execute(&buf, c); err != nil {
		panic("statsview: failed to execute template " + err.Error())
	}

	return buf.String()
}

func fixedPrecision(n float64, p int) float64 {
	var r float64
	switch p {
	case 2:
		r, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	case 6:
		r, _ = strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)
	}
	return r
}

func newBasicView(route string) *charts.Line {
	graph := charts.NewLine()
	graph.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, Trigger: "axis"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Time"}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "600px",
			Height: "400px",
			Theme:  string(defaultCfg.Theme),
		}),
	)
	graph.SetXAxis([]string{}).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	graph.AddJSFuncs(genViewTemplate(graph.ChartID, route))
	return graph
}

type ValueBuilderFunc func() []float64
