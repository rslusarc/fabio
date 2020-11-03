package prometheus

import (
	gkm "github.com/go-kit/kit/metrics"
	prommetrics "github.com/go-kit/kit/metrics/prometheus"
	promclient "github.com/prometheus/client_golang/prometheus"
)

type Provider struct {
	Opts    promclient.Opts
	Buckets []float64
}

func (p *Provider) NewCounter(name string, labels ...string) gkm.Counter {
	copts := promclient.CounterOpts(p.Opts)
	copts.Name = name
	return prommetrics.NewCounterFrom(copts, labels)
}

func (p *Provider) NewGauge(name string, labels ...string) gkm.Gauge {
	gopts := promclient.GaugeOpts(p.Opts)
	gopts.Name = name
	return prommetrics.NewGaugeFrom(gopts, labels)
}

func (p *Provider) NewHistogram(name string, labels ...string) gkm.Histogram {
	hopts := promclient.HistogramOpts{
		Namespace:   p.Opts.Namespace,
		Subsystem:   p.Opts.Subsystem,
		Name:        name,
		Help:        p.Opts.Help,
		ConstLabels: p.Opts.ConstLabels,
		Buckets:     p.Buckets,
	}
	return prommetrics.NewHistogramFrom(hopts, labels)
}

func (p *Provider) Unregister(v interface{}) {
	// noop
}