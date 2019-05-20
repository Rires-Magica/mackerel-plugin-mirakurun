package mpmirakurun

import (
	"flag"
	"fmt"
	"gopkg.in/resty.v1"
	"github.com/bitly/go-simplejson"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type MirakurunPlugin struct {
	Prefix string
	Port int
}

func (pl MirakurunPlugin) MetricKeyPrefix() string {
	return pl.Prefix
}

func (pl MirakurunPlugin) FetchMetrics() (map[string]int, error) {
	url := fmt.Sprintf("http://localhost:%d/api/tuners", pl.Port)
	resp, err := resty.R().Get(url)
	if err != nil {
		return nil, err
	}
	bytes := []byte(resp.String())
	tuners, err := simplejson.NewJson(bytes)
	if err != nil {
		return nil, err
	}
	// fucking shit poop code start
	var (
		available int
		free int
		using int
		fault int
	)
	for index, _ := range tuners.MustArray() {
		if tuners.GetIndex(index).Get("isAvailable").MustBool() {
			available++
		}
		if tuners.GetIndex(index).Get("isFree").MustBool() {
			free++
		}
		if tuners.GetIndex(index).Get("isUsing").MustBool() {
			using++
		}
		if tuners.GetIndex(index).Get("isFault").MustBool() {
			fault++
		}
	}
	// fucking shit poop code end
	return map[string]int{
		"available": available,
		"free": free,
		"using": using,
		"fault": fault
	}, nil
}

func (pl MirakurunPlugin) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		"tuners": {
			Label: pl.Prefix,
			Unit: mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "available", Label: "Available"},
				{Name: "free", Label: "Free", Stacked: true},
				{Name: "using", Label: "Using", Stacked: true},
				{Name: "fault", Label: "Fault", Stacked: true}
			}
		}
	}
}

func Do() {
	optPort = flag.Int("mirakurun-port", 40772, "Mirakurun API Port (http)")
	optPrefix = flag.String("metric-key-prefix", "mirakurun", "Metric Key Prefix")
	flag.Parse()
	mirakurunpl := MirakurunPlugin{
		Prefix: *optPrefix
		Port: *optPort
	}
	plugin := mp.NewMackerelPlugin(mirakurunpl)
	plugin.Run()
}