package parser

import "testing"

func TestParser(t *testing.T) {
	type tcase struct {
		sql string
	}
	fn := func(t *testing.T, name string, tc tcase) {
		ast, err := Parse(name, []byte(tc.sql))
		t.Logf("ast: %v, err: %v", ast, err)
	}
	tests := map[string]tcase{
		"simple tests": tcase{
			sql: "SELECT * FROM tble",
		},
	}
	for name, k := range tests {
		tc := tc // make a copy of tc.
		t.Run(name, func(t *testing.T) { fn(t, name, tc) })
	}

}
