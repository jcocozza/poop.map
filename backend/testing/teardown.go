package testing

import (
	"os"
	"path/filepath"
)

func TearDownTest() error {
	logPattern := "cassidy-log-*-test.log"
	logs, err := filepath.Glob(logPattern)
	if err != nil { panic(err) }
	for _, log := range logs {
		err := os.Remove(log)
		if err != nil {
			panic(err)
		}
	}
	return os.Remove("poop_locations_database_test.db")
}
