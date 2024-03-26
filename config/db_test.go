package config_test

import (
	"miniproject3-markasbali/config"
	"testing"
)

func TestConnection(t *testing.T) {
	testing.Init()
	config.OpenDB()
}
