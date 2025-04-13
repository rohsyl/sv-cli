//go:build windows
// +build windows

package metrics

import (
	"golang.org/x/sys/windows/registry"
)

func getWindowsOSInfo() map[string]string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.READ)
	if err != nil {
		return map[string]string{}
	}
	defer k.Close()

	info := map[string]string{}

	if productName, _, err := k.GetStringValue("ProductName"); err == nil {
		info["PRETTY_NAME"] = productName
	}
	if currentVersion, _, err := k.GetStringValue("CurrentVersion"); err == nil {
		info["VERSION_ID"] = currentVersion
	}
	if _, _, err := k.GetStringValue("EditionID"); err == nil {
		info["ID"] = "windows"
		info["ID_LIKE"] = "windows"
	}

	return info
}
