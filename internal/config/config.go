package config

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"sort"
)

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	filename := fmt.Sprintf("./config/.env.%v.yml", env)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	dump(config)

	return &config, nil
}

func validateConfig(config *Config) error {
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		yamlKey := field.Tag.Get("yaml")
		if yamlKey != "" {
			fieldValue := v.Field(i)

			if isEmptyValue(fieldValue) {
				fmt.Printf("Error: Field %s is empty\n", field.Name)
				return fmt.Errorf("missing value for field: %s", field.Name)
			}

			if fieldValue.Kind() == reflect.Struct {
				if err := validateNestedStruct(fieldValue.Addr().Interface()); err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("missing YAML key for field: %s", field.Name)
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	}
	return false
}

func validateNestedStruct(i interface{}) error {
	v := reflect.ValueOf(i).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		yamlKey := field.Tag.Get("yaml")
		if yamlKey == "" {
			return fmt.Errorf("missing YAML key for nested field: %s", field.Name)
		}
	}
	return nil
}

func dump(config interface{}) {
	configMap := structToMap(config)

	var keys []string
	for key := range configMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Config Key", "Config Value"})
	t.SetBorder(true)

	for _, key := range keys {
		element := configMap[key]
		t.Append([]string{key, fmt.Sprintf("%v", element)})
	}

	t.Render()
}

func structToMap(config interface{}) map[string]interface{} {
	configMap := make(map[string]interface{})

	v := reflect.ValueOf(config)

	// Если передан указатель на структуру, работаем с её значением
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			// Для вложенных структур рекурсивно добавляем поля с точечной нотацией
			nestedMap := structToMap(fieldValue.Interface())
			for nestedKey, nestedValue := range nestedMap {
				configMap[fmt.Sprintf("%s.%s", field.Name, nestedKey)] = nestedValue
			}
		} else {
			// Для других типов просто добавляем значение
			configMap[field.Name] = fieldValue.Interface()
		}
	}

	return configMap
}
