package utils

import (
	"github.com/thatisuday/commando"
)

// something unique that describes an "empty" flag value in CLI
// this is an issue of "commando" library and only a workarround
const EmptyFlagStringValue = "\u0000023d137f-9a5c-4a5f-9daf-c2b8034c00bc\u0000"

// GetStringFlag() - returns a string flag value without error
func GetStringFlag(flags map[string]commando.FlagValue, name string, defaultValue string) string {
	item, itemExists := flags[name]
	if !itemExists {
		return defaultValue // entry does not exist
	}

	value, err := item.GetString()
	if err == nil {
		if value == EmptyFlagStringValue {
			return defaultValue
		}

		return value
	} else {
		return defaultValue // invalid => use default
	}
}
