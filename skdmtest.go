package skdmtest

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, from v1 branch %s", name)
}
