package staticanalysis_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestContains(t *testing.T) {
	t.Parallel()

	type args struct {
		i    string
		want bool
	}

	cases := []struct {
		name string
		args args
	}{
		{
			name: "true",
			args: args{
				i:    "\"select * from users where user=?\"",
				want: true,
			},
		},
		{
			name: "false",
			args: args{
				i:    "\"delete * from users where user=?\"",
				want: false,
			},
		},
	}

	for n, tt := range cases {
		tt := tt
		n := n
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			t.Parallel()
			got := strings.Contains("select", tt.args.i)
			if diff := cmp.Diff(tt.args.want, got); diff != "" {
				t.Errorf("strings.Contains does not notice content: (-got +want)\n%s", diff)
			}
		})
	}
}
