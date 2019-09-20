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

func TestResourceList_Load(t *testing.T) {
	type args struct {
		window string
	}
	tests := []struct {
		name    string
		r       ResourceList
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Normal float value",
			r: ResourceList{
				{
					Name:  "load_avg",
					Type:  "hl",
					Value: "31.63",
				},
			},
			args: args{
				window: "avg",
			},
			want:    31.63,
			wantErr: false,
		},
		{
			name: "Unknown window",
			r: ResourceList{
				{
					Name:  "load_avg",
					Type:  "hl",
					Value: "31.63",
				},
			},
			args: args{
				window: "meow",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Unparseable float",
			r: ResourceList{
				{
					Name:  "load_avg",
					Type:  "hl",
					Value: "31.meow",
				},
			},
			args: args{
				window: "avg",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Load(tt.args.window)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceList_getStorageValueFromList(t *testing.T) {
	type args struct {
		KeyName string
	}
	tests := []struct {
		name    string
		r       ResourceList
		args    args
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Normal Gigabyte Memory Value",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "3.2G",
				},
			},
			args: args{
				KeyName: "free_mem",
			},
			want: StorageValue{
				Size:  3.2,
				Scale: "G",
				Bytes: 3200000000,
			},
			wantErr: false,
		},
		{
			name: "Can't Locate Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "3.2G",
				},
			},
			args: args{
				KeyName: "meow",
			},
			want: StorageValue{
				Size:  0,
				Scale: "",
				Bytes: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.getStorageValueFromList(tt.args.KeyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.getStorageValueFromList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.getStorageValueFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}
