package gogridengine

import (
	"reflect"
	"testing"
)

func Test_newStorageValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    StorageValue
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "57 Gigabytes",
			args: args{
				input: "57.00G",
			},
			want: StorageValue{
				Bytes: 57000000000,
				Scale: "G",
				Size:  57.00,
			},
		},
		{
			name: "1.01 Megabytes",
			args: args{
				input: "1.01M",
			},
			want: StorageValue{
				Bytes: 1010000,
				Scale: "M",
				Size:  1.01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newStorageValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStorageValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStorageValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
