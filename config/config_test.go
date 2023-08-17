package config

import (
	"flag"
	"testing"
	"os"
)

func TestLoadConfig(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }() 

	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	origCommandLine := flag.CommandLine
	flag.CommandLine = flagSet 

	os.Args = []string{"cmd", "-rpcurl=http://test:8545", "-indexpast", "-db=test.db", "-apiport=3000", "-startblock=100", "-blockspencycle=20"}

	expectedConfig := &AppConfig{
		RPCURL:         "http://test:8545",
		IndexPast:      true,
		Subscribe:      false,
		DatabaseDSN:    "test.db",
		APIPort:        "3000",
		StartBlock:     100,
		BlocksPerCycle: 20,
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if *config != *expectedConfig {
		t.Fatalf("Mismatched configs: got %+v, want %+v", config, expectedConfig)
	}

	flag.CommandLine = origCommandLine
}