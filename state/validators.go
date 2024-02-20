package state

import (
	"fmt"
	"reflect"
	"strings"

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
		return fmt.Errorf("'%s' should be of type '%T', found '%T'", opt.ViperKey(), opt.Type, reflect.TypeOf(val))
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
		return fmt.Errorf("'%s' should be of type '%T', found '%T'", opt.ViperKey(), opt.Type, reflect.TypeOf(val))
	}
	return nil
}

func DatabaseValidator(opt ConfigOption, v any) error {
	switch val := v.(type) {
	case DatabaseConfig:
		switch val.Type {
		case "sqlite":
			if val.Settings.Sqlite == nil {
				return fmt.Errorf("'%s.settings.sqlite' not found", opt.ViperKey())
			}

			sqlSettings := val.Settings.Sqlite
			if sqlSettings.Path == "" {
				return fmt.Errorf("'%s.settings.sqlite.path' not found or is empty", opt.ViperKey())
			}

		case "postgresql":
			if val.Settings.PostgresQL == nil {
				return fmt.Errorf("'%s.settings.postgresql' not found", opt.ViperKey())
			}

			pgSettings := val.Settings.PostgresQL
			if pgSettings.DBName == "" || pgSettings.Host == "" || pgSettings.Password == "" || pgSettings.User == "" || pgSettings.Port == 0 {
				return fmt.Errorf("%s.settings.postgresql has missing or empty values", opt.ViperKey())
			}

		case "mysql":
			if val.Settings.MySQL == nil {
				return fmt.Errorf("'%s.settings.mysql' not found", opt.ViperKey())
			}

			msSettings := val.Settings.MySQL
			if msSettings.DBName == "" || msSettings.Host == "" || msSettings.Password == "" || msSettings.Port == 0 || msSettings.User == "" {
				return fmt.Errorf("%s.settings.mysql has missing or empty values", opt.ViperKey())
			}
		}

	default:
		return fmt.Errorf("'%s' should be of type '%T', found '%T'", opt.ViperKey(), opt.Type, reflect.TypeOf(val))
	}
	return nil
}
