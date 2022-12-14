//go:build (!windows && !darwin) || (darwin && !cgo)
// +build !windows,!darwin darwin,!cgo

package ieproxy

func (psc *ProxyScriptConf) findProxyForURL(URL string) string {
	return ""
}
