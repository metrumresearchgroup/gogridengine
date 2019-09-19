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
		{
			name: "4 Teranutes",
			args: args{
				input: "4.76T",
			},
			want: StorageValue{
				Bytes: 4760000000000,
				Scale: "T",
				Size:  4.76,
			},
		},
		{
			name: "Invalid Storage Value",
			args: args{
				input: "1.01M~",
			},
			want: StorageValue{
				Bytes: 0,
				Scale: "",
				Size:  0,
			},
			wantErr: true,
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

func TestResourceList_NumberofProcessors(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    int32
		wantErr bool
	}{
		{
			name: "Testing for 2",
			r: ResourceList{
				{
					Name:  "num_proc",
					Type:  "hl",
					Value: "2",
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Testing for large int 32",
			r: ResourceList{
				{
					Name:  "num_proc",
					Type:  "hl",
					Value: "16820",
				},
			},
			want:    16820,
			wantErr: false,
		},
		{
			name: "Testing for invalid string",
			r: ResourceList{
				{
					Name:  "num_proc",
					Type:  "hl",
					Value: "meow",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Testing for unlocateable key",
			r: ResourceList{
				{
					Name:  "meow",
					Type:  "hl",
					Value: "2",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.NumberofProcessors()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.NumberofProcessors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.NumberofProcessors() = %v, want %v", got, tt.want)
			}
		})
	}
}
