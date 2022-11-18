package nomad

import (
	"strings"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// init registers this plugin.
func init() { plugin.Register(pluginName, setup) }

func setup(c *caddy.Controller) error {
	n := Nomad{}
	err := parse(c, n)

	if err != nil {
		return plugin.Error("nomad", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		n.Next = next
		return n
	})

	return nil
}

func parse(c *caddy.Controller, n Nomad) error {
	for c.Next() {
		for c.NextBlock() {
			selector := strings.ToLower(c.Val())

			switch selector {
			case "foo":
				n.Foo = c.RemainingArgs()[0]
			case "bar":
				n.Bar = c.RemainingArgs()[0]
			default:
				return c.Errf("unknown property '%s'", selector)
			}
		}
	}

	return nil
}
