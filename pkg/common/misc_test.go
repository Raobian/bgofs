package common

import (
	"testing"
)

func TestStrToSize(t *testing.T) {
	type args struct {
		sizeStr string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{
			name: "testa",
			args: args{
				sizeStr: "1024",
			},
			want: 1 << 10,
		}, {
			name: "testb",
			args: args{
				sizeStr: "284TB",
			},
			want: 284 << 40,
		}, {
			name: "testc",
			args: args{
				sizeStr: "567B",
			},
			want: 567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToSize(tt.args.sizeStr); got != tt.want {
				t.Errorf("StrToSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSizeToStr(t *testing.T) {
	type args struct {
		size uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testa",
			args: args{
				size: 2048,
			},
			want: "2.00KB",
		}, {
			name: "testb",
			args: args{
				size: 20,
			},
			want: "20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SizeToStr(tt.args.size); got != tt.want {
				t.Errorf("SizeToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
