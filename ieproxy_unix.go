//go:build (!windows && !darwin) || (darwin && !cgo)
// +build !windows,!darwin darwin,!cgo

package ieproxy

func getConf() ProxyConf {
	return ProxyConf{}
}

func reloadConf() ProxyConf {
	return getConf()
}

func overrideEnvWithStaticProxy(pc ProxyConf, setenv envSetter) {
}
