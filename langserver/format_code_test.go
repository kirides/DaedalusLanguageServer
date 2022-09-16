package langserver_test

import (
	"testing"

	"github.com/kirides/DaedalusLanguageServer/langserver"
)

func TestFormatCode(t *testing.T) {
	langserver.FormatCode(`
	func void
	DoStuff   (var int val,
		var float b) {

	};
	
	class C_ITEM { var int v1; var int v2; var string str;
	};
	`)
}
