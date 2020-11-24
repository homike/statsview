# Statsview

Statsview 是一个简单的统计数据可视化工具, 基于[go-echarts](https://github.com/go-echarts/go-echarts).

## 📝 Usage

参考examples/demo

```golang
views := []statsview.Viewer{
    statsview.NewBasicViewer("Goroutine", nil, func() []float64 {
        return generateValues()
    })
}
statsview.Startup(views)

// Visit your browser at http://localhost:8088/statsview
```

## ⚙️ Configuration

Statsview gets a variety of configurations for the users. Everyone could customize their favorite charts style.

```golang
// WithInterval sets the interval(in Millisecond) of collecting and pulling metrics
// default -> 2000
WithInterval(interval int)

// WithMaxPoints sets the maximum points of each chart series
// default -> 30
WithMaxPoints(n int)

// WithTemplate sets the rendered template which fetching stats from the server and
// handling the metrics data
WithTemplate(t string)

// WithAddr sets the listen address
// default -> "localhost:18066"
WithAddr(addr string)

// WithTimeFormat sets the time format for the line-chart Y-axis label
// default -> "15:04:05"
WithTimeFormat(s string)

// WithTheme sets the theme of the charts
// default -> Macarons
//
// Optional:
// * ThemeWesteros
// * ThemeMacarons
WithTheme(theme Theme)
```

#### Set the options

```golang
statsview.SetConfiguration(
    statsview.WithAddr("192.168.0.1:8088"),
    statsview.WithInterval(10000))
```

## ������ Snapshot

![Macarons](https://github.com/homike/media/blob/main/statsview.png)
