package chapter5

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKadai_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   cutEntity
		isError bool
	}{
		{
			name: "成功",
			input: cutEntity{
				args:      []string{"aaaaa", "bbbbb"},
				delimiter: ",",
				fields:    1,
			},
			isError: false,
		},
		{name: "要素なし",
			input: cutEntity{
				args:      nil,
				delimiter: "",
				fields:    0,
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isError {
				assert.Error(t, Validate(tt.input))
			} else {
				assert.NoError(t, Validate(tt.input))
			}
		})
	}
}

func TestKadai_Cut(t *testing.T) {
	type input struct {
		cutEntity
		str io.Reader
	}

	tests := []struct {
		name    string
		input   input
		expect  string
		isError bool
	}{
		{
			name: "成功",
			input: input{
				cutEntity: cutEntity{
					delimiter: ":",
					fields:    1,
				},
				str: bytes.NewBufferString("aaaaaa:bbbbbb"),
			},
			isError: false,
			expect:  "aaaaaa",
		},
		{name: "要素なし",
			input: input{
				cutEntity: cutEntity{
					delimiter: ":",
					fields:    3,
				},
				str: bytes.NewBufferString("aaaaaa:bbbbbb"),
			},
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if tt.isError {
				assert.Error(t, Cut(tt.input.cutEntity, tt.input.str, &out))
			} else {
				assert.NoError(t, Cut(tt.input.cutEntity, tt.input.str, &out))
				assert.Equal(t, tt.expect, out.String())
			}
		})
	}
}

func BenchmarkCut(b *testing.B) {
	entity := cutEntity{
		delimiter: ":",
		fields:    1,
	}
	str := bytes.NewBufferString("aaaaaa:bbbbbb")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Cut(entity, str, os.Stdout)
	}
}
