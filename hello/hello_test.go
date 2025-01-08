package hello

import (
	"testing"
)

func TestGetHello1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "return 'Hello1 DDD!!' if you set 'DDD'",
			args: args{
				"DDD",
			},
			want: "Hello1 DDD!!",
		},
		{
			name: "空文字を指定したら Hello1 が返ってくる",
			args: args{
				"",
			},
			want: "Hello1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHello1(tt.args.s); got != tt.want {
				t.Errorf("GetHello1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHello_GetHello2(t *testing.T) {
	type fields struct {
		Say string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "return 'Hello2 DDD!!' if you set 'DDD'",
			fields: fields{
				Say: "DDD",
			},
			want: "Hello2 DDD!!",
		},
		{
			name: "空文字を指定したら Hello2 が返ってくる",
			fields: fields{
				Say: "",
			},
			want: "Hello2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Hello{
				Say: tt.fields.Say,
			}
			if got := h.GetHello2(); got != tt.want {
				t.Errorf("Hello.GetHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
