package myConfig

import (
	"testing"
)

const TEST_INIT_CONFIG_PATH = "../../config.yaml"

type TestConfig struct {
	DBPassword   string   `yaml:"DB-password"`
	BuildVersion string   `yaml:"build-version"`
	AllowedIps   []string `yaml:"allowed-ips"`
	AllowedIDs   []int    `yaml:"allowed-ids"`
	SpecialNr    int      `yaml:"special-nr"`
}

func Test_Init_With_Nil_Pointer(t *testing.T) {
	var testCfg *TestConfig
	err := Init(TEST_INIT_CONFIG_PATH, &testCfg)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	//find a certain ip in allowedips
}

func Test_Init_With_Valid_Pointer(t *testing.T) {
	testCfg := getTestConstructor()
	err := Init(TEST_INIT_CONFIG_PATH, &testCfg)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func getTestConstructor() *TestConfig {
	return &TestConfig{DBPassword: "", BuildVersion: "test"}
}
