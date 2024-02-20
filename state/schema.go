package state

import (
	"fmt"
	"reflect"
	"strings"
)

type ConfigOption struct {
	Name        string
	Description string
	Required    bool
	Mutable     bool
	Default     any
	Type        reflect.Type
	Hierarchy   []string
	Validator   OptionValidator
}

func (opt ConfigOption) ViperKey() string {
	keyPath := strings.Join(opt.Hierarchy, ".")
	return keyPath + "." + opt.Name
}

var AllConfigOptions = []ConfigOption{
	{
		Name:      "bot_token",
		Required:  true,
		Mutable:   false,
		Default:   nil,
		Type:      reflect.TypeOf(""),
		Hierarchy: []string{"telegram"},
		Validator: NonEmptyStringValidator,
	},
	{
		Name:      "api_url",
		Required:  false,
		Mutable:   false,
		Default:   "https://api.telegram.org/",
		Type:      reflect.TypeOf(""),
		Hierarchy: []string{"telegram"},
		Validator: NonEmptyStringValidator,
	},
	{
		Name:      "owner_id",
		Required:  true,
		Mutable:   false,
		Default:   nil,
		Type:      reflect.TypeOf(int64(0)),
		Hierarchy: []string{"telegram"},
		Validator: NonZeroNumberValidator[int64],
	},
	{
		Name:      "sudo_users_id",
		Required:  false,
		Mutable:   true,
		Default:   []int64{},
		Type:      reflect.TypeOf([]int64{}),
		Hierarchy: []string{"telegram"},
		Validator: nil,
	},
	{
		Name:      "update_commands",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"telegram"},
		Validator: nil,
	},
	{
		Name:      "send_online_presence",
		Required:  false,
		Mutable:   true,
		Default:   false,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"telegram"},
		Validator: nil,
	},
	{
		Name:      "send_read_receipts",
		Required:  false,
		Mutable:   true,
		Default:   false,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"telegram"},
		Validator: nil,
	},
	{
		Name:      "session_name",
		Required:  false,
		Mutable:   false,
		Default:   "WaTgBridge",
		Type:      reflect.TypeOf(""),
		Hierarchy: []string{"whatsapp"},
		Validator: NonEmptyStringValidator,
	},
	{
		Name:      "login_database",
		Required:  true,
		Mutable:   false,
		Default:   nil,
		Type:      reflect.TypeOf(DatabaseConfig{}),
		Hierarchy: []string{"whatsapp"},
		Validator: DatabaseValidator,
	},
	{
		Name:      "documents",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "images",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "gifs",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "videos",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "voice_notes",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "audios",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "status",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "stickers",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "video_stickers",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "contacts",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "locations",
		Required:  false,
		Mutable:   true,
		Default:   true,
		Type:      reflect.TypeOf(false),
		Hierarchy: []string{"whatsapp", "bridging"},
		Validator: nil,
	},
	{
		Name:      "pack_name",
		Required:  false,
		Mutable:   true,
		Default:   "WaTgBridge",
		Type:      reflect.TypeOf(""),
		Hierarchy: []string{"whatsapp", "sticker_metadata"},
		Validator: NonEmptyStringValidator,
	},
	{
		Name:      "author_name",
		Required:  false,
		Mutable:   true,
		Default:   "WaTgBridge",
		Type:      reflect.TypeOf(""),
		Hierarchy: []string{"whatsapp", "sticker_metadata"},
		Validator: NonEmptyStringValidator,
	},
}

func init() {
	for _, configOpt := range AllConfigOptions {
		fmt.Println(configOpt.ViperKey(), "=>", configOpt.Type)
	}
}
