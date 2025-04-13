//go:build !windows
// +build !windows

package metrics

func getWindowsOSInfo() map[string]string {
	return map[string]string{}
}
