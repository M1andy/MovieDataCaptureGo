package proxy

import (
	c "MovieDataCaptureGo/internal/config"
	. "MovieDataCaptureGo/internal/logger"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

var ProxyFunc colly.ProxyFunc

func newProxyFunc(cfg *c.Config) colly.ProxyFunc {
	if !cfg.Proxy.Switch {
		return nil
	}
	ps, err := proxy.RoundRobinProxySwitcher(c.CFG.Proxy.URL...)
	if err != nil {
		Logger.Warnln(err)
	}
	return ps
}

func init() {
	ProxyFunc = newProxyFunc(c.CFG)
}
