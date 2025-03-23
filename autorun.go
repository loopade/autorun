package autorun

// AutoRunConfig is a struct that contains the configuration for the autorun
// AppName is the name of the application, e.g. "MyApp"
// ExecutablePath is the path to the executable, e.g. "/usr/bin/myapp"
// CompanyName is the name of the company, e.g. "com.company"
type AutoRunConfig struct {
	AppName        string
	ExecutablePath string
	CompanyName    string
}

func QueryAutoRun(cfg *AutoRunConfig) (bool, error) {
	return queryAutoRun(cfg)
}
func EnableAutoRun(cfg *AutoRunConfig) error {
	return enableAutoRun(cfg)
}
func DisableAutoRun(cfg *AutoRunConfig) error {
	return disableAutoRun(cfg)
}
