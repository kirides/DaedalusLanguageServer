package DaedalusLanguageServer

import "embed"

//go:embed DaedalusBuiltins/*
var BuiltinsFS embed.FS

const DaedalusBuiltinsPath = "DaedalusBuiltins"
