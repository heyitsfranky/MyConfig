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
			fieldElem := reflect.ValueOf(obj).Elem().FieldByName(field.Name)
			if err := convertFieldValue(fieldElem, value); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("missing key '%s' in '%s'", fieldName, configPath)
		}
	}
	return nil
}

// convertFieldValue converts the value to the desired field type and sets it in the field.
func convertFieldValue(fieldElem reflect.Value, value interface{}) error {
	switch fieldElem.Kind() {
	case reflect.String:
		if stringValue, ok := value.(string); ok {
			fieldElem.SetString(stringValue)
		} else {
			return fmt.Errorf("expected a string, got '%v'", value)
		}
	case reflect.Int:
		if intValue, ok := value.(int); ok {
			fieldElem.SetInt(int64(intValue))
		} else {
			return fmt.Errorf("expected an int, got '%v'", value)
		}
	case reflect.Slice:
		slice := reflect.MakeSlice(fieldElem.Type(), 0, 0)
		if valueSlice, ok := value.([]interface{}); ok {
			for _, elem := range valueSlice {
				convertedElem, err := convertToType(elem, fieldElem.Type().Elem())
				if err != nil {
					return err
				}
				slice = reflect.Append(slice, convertedElem)
			}
			fieldElem.Set(slice)
		} else {
			return fmt.Errorf("expected a slice, got '%v'", value)
		}
	case reflect.Interface:
		if value != nil {
			valueOf := reflect.ValueOf(value)
			fieldElem.Set(valueOf)
		}
	default:
		return fmt.Errorf("unsupported field type")
	}
	return nil
}

func convertToType(value interface{}, targetType reflect.Type) (reflect.Value, error) {
	valueOf := reflect.ValueOf(value)
	convertedValue := valueOf.Convert(targetType)
	return convertedValue, nil
}
