package proxy

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"

	c "MovieDataCaptureGo/internal/config"
	. "MovieDataCaptureGo/internal/logger"
)

var ProxyFunc colly.ProxyFunc

func newProxyFunc(cfg *c.Config) colly.ProxyFunc {
	if !cfg.Proxy.Switch {
		return nil
	}
	ps, err := proxy.RoundRobinProxySwitcher(c.CFG.Proxy.URL...)
	Logger.Debugf("Proxy List: %s", c.CFG.Proxy.URL)
	if err != nil {
		Logger.Warnln(err)
	}
	return ps
}

func init() {
	ProxyFunc = newProxyFunc(c.CFG)
}
