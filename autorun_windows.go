package autorun

import (
	"golang.org/x/sys/windows/registry"
)

func queryAutoRun(cfg *AutoRunConfig) (bool, error) {
	key, err := openAutoRunKey(registry.QUERY_VALUE)
	if err != nil {
		return false, err
	}
	defer key.Close()
	val, _, err := key.GetStringValue(cfg.AppName)
	if err != nil {
		if err == registry.ErrNotExist {
			return false, nil
		}
		return false, err
	}
	if val == cfg.ExecutablePath {
		return true, nil
	}
	return false, nil
}

func openAutoRunKey(access uint32) (registry.Key, error) {
	autorunKey := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	key, err := registry.OpenKey(registry.CURRENT_USER, autorunKey, access)
	if err != nil {
		return 0, err
	}
	return key, nil
}

func enableAutoRun(cfg *AutoRunConfig) error {
	key, err := openAutoRunKey(registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	return key.SetStringValue(cfg.AppName, cfg.ExecutablePath)
}

func disableAutoRun(cfg *AutoRunConfig) error {
	key, err := openAutoRunKey(registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	return key.DeleteValue(cfg.AppName)
}
