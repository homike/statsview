package statsview

import "github.com/go-echarts/go-echarts/v2/types"

const (
	DefaultTemplate = `
$(function () { setInterval({{ .ViewID }}_sync, {{ .Interval }}); });
function {{ .ViewID }}_sync() {
    $.ajax({
        type: "GET",
        url: "http://{{ .Addr }}/statsview/view/{{ .Route }}",
        dataType: "json",
        success: function (result) {
            let opt = goecharts_{{ .ViewID }}.getOption();

            let x = opt.xAxis[0].data;
            x.push(result.time);
            if (x.length > {{ .MaxPoints }}) {
                x = x.slice(1);
            }
            opt.xAxis[0].data = x;

            for (let i = 0; i < result.values.length; i++) {
                let y = opt.series[i].data;
                y.push({ value: result.values[i] });
                if (y.length > {{ .MaxPoints }}) {
                    y = y.slice(1);
                }
                opt.series[i].data = y;

                goecharts_{{ .ViewID }}.setOption(opt);
            }
        }
    });
}`
	DefaultMaxPoints  = 200
	DefaultTimeFormat = "15:04:05"
	DefaultInterval   = 2000
	DefaultAddr       = "localhost:18066"
	DefaultTheme      = ThemeMacarons
)

type config struct {
	Interval   int
	MaxPoints  int
	Template   string
	Addr       string
	TimeFormat string
	Theme      Theme
}

type Theme string

const (
	ThemeWesteros Theme = types.ThemeWesteros
	ThemeMacarons Theme = types.ThemeMacarons
)

func SetConfiguration(opts ...Option) {
	for _, opt := range opts {
		opt(defaultCfg)
	}
}

var defaultCfg = &config{
	Interval:   DefaultInterval,
	MaxPoints:  DefaultMaxPoints,
	Template:   DefaultTemplate,
	Addr:       DefaultAddr,
	TimeFormat: DefaultTimeFormat,
	Theme:      DefaultTheme,
}

type Option func(c *config)

// Addr returns the default server listening address
func Addr() string {
	return defaultCfg.Addr
}

// Interval returns the default collecting interval of ViewManager
func Interval() int {
	return defaultCfg.Interval
}

// WithInterval sets the interval of collecting and pulling metrics
func WithInterval(interval int) Option {
	return func(c *config) {
		c.Interval = interval
	}
}

// WithMaxPoints sets the maximum points of each chart series
func WithMaxPoints(n int) Option {
	return func(c *config) {
		c.MaxPoints = n
	}
}

// WithTemplate sets the rendered template which fetching stats from the server and
// handling the metrics data
func WithTemplate(t string) Option {
	return func(c *config) {
		c.Template = t
	}
}

// WithAddr sets the listening address
func WithAddr(addr string) Option {
	return func(c *config) {
		c.Addr = addr
	}
}

// WithTimeFormat sets the time format for the line-chart Y-axis label
func WithTimeFormat(s string) Option {
	return func(c *config) {
		c.TimeFormat = s
	}
}

// WithTheme sets the theme of the charts
func WithTheme(theme Theme) Option {
	return func(c *config) {
		c.Theme = theme
	}
}
