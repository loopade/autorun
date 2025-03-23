package autorun

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func getDesktopFilePath(cfg *AutoRunConfig) string {
	return fmt.Sprintf(".config/autostart/%s.desktop", strings.ToLower(cfg.AppName))
}

func queryAutoRun(cfg *AutoRunConfig) (bool, error) {
	usr, err := user.Current()
	if err != nil {
		return false, err
	}
	desktopPath := filepath.Join(usr.HomeDir, getDesktopFilePath(cfg))
	_, err = os.Stat(desktopPath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func enableAutoRun(cfg *AutoRunConfig) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	desktopPath := filepath.Join(usr.HomeDir, getDesktopFilePath(cfg))
	desktopContent := fmt.Sprintf(`[Desktop Entry]
Type=Application
Name=%s
Exec=%s
StartupNotify=false`, cfg.AppName, cfg.ExecutablePath)

	return os.WriteFile(desktopPath, []byte(desktopContent), 0644)
}

func disableAutoRun(cfg *AutoRunConfig) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	desktopPath := filepath.Join(usr.HomeDir, getDesktopFilePath(cfg))
	return os.Remove(desktopPath)
}
