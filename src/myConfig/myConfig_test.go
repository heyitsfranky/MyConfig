package myConfig

import (
	"testing"
)

const TEST_INIT_CONFIG_PATH = "../../config.yaml"

type TestConfig struct {
	DBPassword   string        `yaml:"DB-password"`
	BuildVersion string        `yaml:"build-version"`
	AllowedIps   []string      `yaml:"allowed-ips"`
	AllowedIDs   []int         `yaml:"allowed-ids"`
	SpecialNr    int           `yaml:"special-nr"`
	DynamicArray []interface{} `yaml:"dynamic-array"`
	DynamicValue interface{}   `yaml:"dynamic-value"`
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

func Test_Keys_Missing_In_Config(t *testing.T) {
	paths := []string{
		"missing_string.yaml",
		"missing_string_slice.yaml",
		"missing_int.yaml",
		"missing_int_slice.yaml",
		"missing_dynamic.yaml",
		"missing_dynamic_slice.yaml",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			var testCfg *TestConfig
			err := Init("../../test-configs/"+path, &testCfg)
			if err == nil {
				t.Errorf("Expected an error for path %s, but got nothing", path)
			}
		})
	}
}
