package autorun

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getLaunchAgentPath(cfg *AutoRunConfig) string {
	return fmt.Sprintf("Library/LaunchAgents/%s.%s.plist", strings.ToLower(cfg.CompanyName), strings.ToLower(cfg.AppName))
}

func getLaunchAgentContent(cfg *AutoRunConfig) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>%s.%s</string>
	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
</dict>
</plist>
`, strings.ToLower(cfg.CompanyName), strings.ToLower(cfg.AppName), cfg.ExecutablePath)
}

func queryAutoRun(cfg *AutoRunConfig) (bool, error) {
	_, err := os.Stat(filepath.Join(os.Getenv("HOME"), getLaunchAgentPath(cfg)))
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func enableAutoRun(cfg *AutoRunConfig) error {
	launchAgent := filepath.Join(os.Getenv("HOME"), getLaunchAgentPath(cfg))
	return os.WriteFile(launchAgent, []byte(getLaunchAgentContent(cfg)), 0644)
}

func disableAutoRun(cfg *AutoRunConfig) error {
	launchAgent := filepath.Join(os.Getenv("HOME"), getLaunchAgentPath(cfg))
	return os.Remove(launchAgent)
}
