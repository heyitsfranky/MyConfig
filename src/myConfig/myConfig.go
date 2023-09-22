package myConfig

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// This method will only initialize nil pointers, therefore allowing it to be called multiple times without a file-reading performance penalty to ensure a config-ptr being non-nil.
func Init(configPath string, objPtr interface{}) error {
	objValue := reflect.ValueOf(objPtr).Elem()
	if objValue.IsNil() {
		newObj := reflect.New(objValue.Type().Elem())
		objValue.Set(newObj)
		err := read(configPath, objValue.Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func read(configPath string, obj interface{}) error {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	// To check if all keys are present, we have to load them into a tempConfig
	var tempConfig map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &tempConfig)
	if err != nil {
		return err
	}
	// Ensure obj is valid
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr || objValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("obj must be a pointer to a struct")
	}
	if objValue.IsNil() {
		objValue.Set(reflect.New(objValue.Elem().Type()))
	}

	// Actual checking step
	objType := reflect.TypeOf(obj).Elem()
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldName := field.Tag.Get("yaml")
		if value, exists := tempConfig[fieldName]; exists {
			configFieldType := field.Type
			if configFieldType.Kind() == reflect.String {
				fieldValue := reflect.ValueOf(obj).Elem().FieldByName(field.Name)
				fieldValue.SetString(value.(string))
			} else {
				return fmt.Errorf("unsupported field type for '%s' in '%s'", fieldName, configPath)
			}
		} else {
			return fmt.Errorf("missing key '%s' in '%s'", fieldName, configPath)
		}
	}
	return nil
}
