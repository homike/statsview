package statsview

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/rs/cors"

	"statsview/statics"
)

func init() {
	templates.PageTpl = `
{{- define "page" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
<style> .box { justify-content:center; display:flex; flex-wrap:wrap } </style>
<div class="box"> {{- range .Charts }} {{ template "base" . }} {{- end }} </div>
</body>
</html>
{{ end }}
`
}

func Startup(views []Viewer) {
	// page
	page := components.NewPage()
	page.PageTitle = "Statsview"
	page.AssetsHost = fmt.Sprintf("http://%s/statsview/statics/", Addr())
	page.Assets.JSAssets.Add("jquery.min.js")

	// views
	mux := http.NewServeMux()
	for _, v := range views {
		page.AddCharts(v.View())
		mux.HandleFunc("/statsview/view/"+v.Name(), v.Serve)
	}

	mux.HandleFunc("/statsview", func(w http.ResponseWriter, _ *http.Request) {
		page.Render(w)
	})
	staticsPrev := "/statsview/statics/"
	mux.HandleFunc(staticsPrev+"echarts.min.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.EchartJS))
	})
	mux.HandleFunc(staticsPrev+"jquery.min.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.JqueryJS))
	})
	mux.HandleFunc(staticsPrev+"themes/westeros.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.WesterosJS))
	})
	mux.HandleFunc(staticsPrev+"themes/macarons.js", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(statics.MacaronsJS))
	})

	srv := &http.Server{
		Addr:           Addr(),
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	srv.Handler = cors.Default().Handler(mux)
	srv.ListenAndServe()
}
