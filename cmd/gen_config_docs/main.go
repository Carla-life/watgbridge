package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/akshettrj/watgbridge/pkg/config"
)

var MarkDownFileLines = []string{
	"# WaTgBridge Configuration Options",
	"",
	"This file has been auto-generated. Do not edit while contributing.",
	"",
	"## Config File",
	"",
	"### File Formats",
	"",
	"The project supports a variety of configuration file formats including **JSON, TOML, YAML, HCL, envfile and Java properties** (all thanks to [spf13/viper](https://github.com/spf13/viper)).",
	"",
	"### File Naming",
	"",
	"You can create a config file with the name `config` and the extension of your own choice (for file formats listed above). For example, to use JSON format, you will have to name the config file `config.json` and place it in the same directory from which you will execute your binary.",
	"",
	"**Make sure that there is only one of such config files at a time in the directory when running the bot as the project will not be able to figure out which file to read and update when changes in config are made through the bot.**",
	"",
	"### Environment Variables",
	"",
	"You can also use environment variables to configure this bot.",
	"",
	"For example, to set the value `bot_api` under `telegram` section, you can use the `WATGBRIDGE_telegram.bot_api` environment variable (i.e. add a `WATGBRIDGE_` prefix).",
	"",
	"*The `WATGBRIDGE` part in the environment variable name is case-sensitive, rest of the name is case-insensitive.*",
	"",
	"## Config Structure",
	"",
	"**Please read the full descriptions carefully.**",
	"",
	"The project looks for the following configuration keys:",
	"",
}

func main() {
	opts := slices.Clone(config.AllConfigOptions)
	slices.SortFunc(opts, func(a config.ConfigOption, b config.ConfigOption) int {
		return strings.Compare(a.ViperKey(), b.ViperKey())
	})

	MarkDownFileLines = append(
		MarkDownFileLines,
		"| Hierarchy | Name | Description | Type | Required | Default Value |",
		"|:---------:|:----:|-------------|:----:|:--------:|:-------------:|",
	)

	for _, opt := range opts {
		var (
			hierarchy    = "-"
			required     = "No"
			defaultValue = ""
		)

		if len(opt.Hierarchy) > 0 {
			hierarchy = opt.Hierarchy[0]
			for _, level := range opt.Hierarchy[1:len(opt.Hierarchy)] {
				hierarchy += " > " + level
			}
		}

		if opt.Required {
			required = "Yes"
		}

		if opt.Default == nil {
			defaultValue = "-"
		} else {
			defaultValue = fmt.Sprintf("%+v", opt.Default)
		}

		MarkDownFileLines = append(
			MarkDownFileLines,
			fmt.Sprintf(
				"| `%s` | `%s` | %s | `%v` | %s | `%s` |",
				hierarchy, opt.Name, opt.Description, opt.Type, required, defaultValue,
			),
		)
	}

	for _, line := range MarkDownFileLines {
		fmt.Println(line)
	}
}
