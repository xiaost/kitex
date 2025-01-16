package env

import (
	"os"
	"strconv"
)

// UseProtoc returns true if we'd like to use protoc instead of prutal.
//
// protoc will be deprecated in the future, and this func will be removed.
func UseProtoc() bool {
	v := os.Getenv("KITEX_TOOL_USE_PROTOC")
	if v == "" {
		return true // use protoc by default
	}
	ok, _ := strconv.ParseBool(v)
	return ok
}
