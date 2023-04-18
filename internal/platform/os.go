package platform

import "runtime"

//goland:noinspection ALL
func IsOsDarwin() bool {
	return runtime.GOOS == "darwin"
}
