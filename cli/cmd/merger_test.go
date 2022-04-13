package cmd

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func Test_merger(t *testing.T) {
	type args struct {
		old io.Reader
		new io.Reader
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name: "when new lines added",
			args: args{
				old: strings.NewReader("a\nb\n"),
				new: strings.NewReader("a\nc\nd"),
			},
			expected: []string{"a", "c", "d"},
		},
		{
			name: "when old lines deleted",
			args: args{
				old: strings.NewReader("a\nb\nc\nd\ne\n"),
				new: strings.NewReader("a\nc\nd\ne\n"),
			},
			expected: []string{"a", "c", "d", "e", "e"},
		},
		{
			name: "when no changes detected",
			args: args{
				old: strings.NewReader("a\nb\nc\n"),
				new: strings.NewReader("a\nb\nc\n"),
			},
			expected: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merger(tt.args.old, tt.args.new)
			assert.EqualValues(t, tt.expected, result)
		})
	}
}
