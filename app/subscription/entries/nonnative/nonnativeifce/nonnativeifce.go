package nonnativeifce

import (
	"io/fs"

	"github.com/imannamdari/v2ray-core/v5/app/subscription/entries"
)

type NonNativeConverterConstructorT func(fs fs.FS) (entries.Converter, error)

var NewNonNativeConverterConstructor NonNativeConverterConstructorT
