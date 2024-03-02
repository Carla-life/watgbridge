package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/exp/constraints"
)

type OptionValidator func(opt ConfigOption, v any) error

func NonEmptyStringValidator(opt ConfigOption, v any) error {
	switch val := v.(type) {
	case string:
		if strings.TrimSpace(val) == "" {
			return fmt.Errorf("'%s' should not be empty", opt.ViperKey())
		}
	default:
		return fmt.Errorf("'%s' should be of type '%s', found '%T'", opt.ViperKey(), opt.Type, reflect.TypeOf(val))
	}
	return nil
}

func NonZeroNumberValidator[T constraints.Integer | constraints.Float](opt ConfigOption, v any) error {
	switch val := v.(type) {
	case T:
		if val == 0 {
			return fmt.Errorf("'%s' should be non-zero", opt.ViperKey())
		}
	default:
		return fmt.Errorf("'%s' should be of type '%s', found '%T'", opt.ViperKey(), opt.Type, reflect.TypeOf(val))
	}
	return nil
}

func DatabaseValidator(opt ConfigOption, v any) error {
	var val DatabaseConfig
	err := viper.UnmarshalKey(opt.ViperKey(), &val)
	if err != nil {
		return fmt.Errorf("'%s' should be of type '%s', but failed to Unmarshal: %w", opt.ViperKey(), opt.Type, err)
	}

	if val.Sqlite != nil {
		if val.Sqlite.Path == "" {
			return fmt.Errorf("'%s.sqlite.path' not found or is empty", opt.ViperKey())
		}

	} else if val.PostgresQL != nil {
		if val.PostgresQL.DBName == "" || val.PostgresQL.Host == "" || val.PostgresQL.Password == "" || val.PostgresQL.User == "" || val.PostgresQL.Port == 0 {
			return fmt.Errorf("%s.postgresql has missing or empty values", opt.ViperKey())
		}

	} else if val.MySQL != nil {
		if val.MySQL.DBName == "" || val.MySQL.Host == "" || val.MySQL.Password == "" || val.MySQL.Port == 0 || val.MySQL.User == "" {
			return fmt.Errorf("%s.mysql has missing or empty values", opt.ViperKey())
		}

	} else {
		return fmt.Errorf("no database configuration found under %s", opt.ViperKey())
	}

	return nil
}
