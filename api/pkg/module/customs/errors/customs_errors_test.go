package errors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	type args struct {
		origin error
		new    error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				origin: NewCustomError(ERROR000),
				new:    NewCustomError(UNDEFINED)},
			wantErr: true,
		},
		{
			name: "origin nil",
			args: args{
				origin: nil,
				new:    NewCustomError(ERROR000)},
			wantErr: true,
		},
		{
			name: "new nil",
			args: args{
				origin: NewCustomError(ERROR000),
				new:    nil,
			},
			wantErr: true,
		},
		{
			name: "origin other error",
			args: args{
				origin: errors.New("origin error"),
				new:    nil,
			},
			wantErr: true,
		},
		{
			name: "new other error",
			args: args{
				origin: nil,
				new:    errors.New("new error"),
			},
			wantErr: true,
		},
		{
			name: "origin and new other error",
			args: args{
				origin: errors.New("origin error"),
				new:    errors.New("new error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Combine(tt.args.origin, tt.args.new); (err != nil) == tt.wantErr {
				return
			}
			t.Errorf("error origin:%#v, new:%#v", tt.args.origin, tt.args.new)
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		err  error
		code string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				err:  NewCustomError(ERROR000),
				code: UNDEFINED,
			},
			wantErr: true,
		},
		{
			name: "err is nil",
			args: args{
				err:  nil,
				code: UNDEFINED,
			},
			wantErr: true,
		},
		{
			name: "code is empty",
			args: args{
				err:  nil,
				code: "",
			},
			wantErr: true,
		},
		{
			name: "err is original",
			args: args{
				err:  errors.New("original error"),
				code: UNDEFINED,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Add(tt.args.err, tt.args.code); (err != nil) == tt.wantErr {
				fmt.Printf("%#v\n", err)
				return
			}
			t.Errorf("no error")
		})
	}
}


func TestToMap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "normal",
			args: args{
				err: NewCustomError(ERROR000),
			},
			want: map[string]string{ERROR000: ErrMap[ERROR000]},
		},
		{
			name: "original error",
			args: args{
				err: errors.New("original error"),
			},
			want: map[string]string{UNDEFINED: "original error"},
		},
		{
			name: "nil",
			args: args{
				err: nil,
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				return 
			}
		})
	}
}

