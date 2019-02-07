//go:generate go run ../../../vendor/github.com/go-bindata/go-bindata/go-bindata -pkg cmd ../../../gmailctl.libsonnet ../../../default-config.jsonnet

package cmd

// gmailctlLib returns the embedded gmailctl.libsonnet file
func gmailctlLib() string {
	return string(MustAsset("../../../gmailctl.libsonnet"))
}

// defaultConfig returns the embedded default configuration file
func defaultConfig() string {
	return string(MustAsset("../../../default-config.jsonnet"))
}
