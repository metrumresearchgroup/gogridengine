package gogridengine

import (
	"reflect"
	"testing"
)

func TestResourceListFreeMemory(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "mem_free",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.FreeMemory()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.FreeMemory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.FreeMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListFreeSwap(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "swap_free",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.FreeSwap()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.FreeSwap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.FreeSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListFreeVirtualMemory(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "virtual_free",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.FreeVirtualMemory()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.FreeVirtualMemory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.FreeVirtualMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListTotalMemory(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "mem_total",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.TotalMemory()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.TotalMemory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.TotalMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListTotalSwap(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "swap_total",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.TotalSwap()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.TotalSwap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.TotalSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListTotalVirtual(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "virtual_total",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.TotalVirtual()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.TotalVirtual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.TotalVirtual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListMemoryUsed(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "mem_used",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.MemoryUsed()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.MemoryUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.MemoryUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListSwapUsed(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "swap_used",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.SwapUsed()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.SwapUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.SwapUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListVirtualUsed(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    StorageValue
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "virtual_used",
					Type:  "hl",
					Value: "2G",
				},
			},
			want: StorageValue{
				Size:  2,
				Scale: "G",
				Bytes: 2000000000,
			},
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
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
			got, err := tt.r.VirtualUsed()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.VirtualUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceList.VirtualUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListCPU(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    float64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "cpu",
					Type:  "hl",
					Value: "1.04",
				},
			},
			want:    1.04,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CPU()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.CPU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.CPU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListNPLoadAverage(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    float64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "np_load_avg",
					Type:  "hl",
					Value: "1.04",
				},
			},
			want:    1.04,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.NPLoadAverage()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.NPLoadAverage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.NPLoadAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListNPLoadShort(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    float64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "np_load_short",
					Type:  "hl",
					Value: "1.04",
				},
			},
			want:    1.04,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.NPLoadShort()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.NPLoadShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.NPLoadShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListNPLoadMedium(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    float64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "np_load_medium",
					Type:  "hl",
					Value: "1.04",
				},
			},
			want:    1.04,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.NPLoadMedium()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.NPLoadMedium() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.NPLoadMedium() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListNPLoadLong(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    float64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "np_load_long",
					Type:  "hl",
					Value: "1.04",
				},
			},
			want:    1.04,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.NPLoadLong()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.NPLoadLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.NPLoadLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListProcessorCount(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    int64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "num_proc",
					Type:  "hl",
					Value: "14",
				},
			},
			want:    14,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ProcessorCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.ProcessorCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.ProcessorCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListMSocketCount(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    int64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "m_socket",
					Type:  "hl",
					Value: "14",
				},
			},
			want:    14,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MSocketCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.MSocketCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.MSocketCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListMThreadCount(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    int64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "m_thread",
					Type:  "hl",
					Value: "14",
				},
			},
			want:    14,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MThreadCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.MThreadCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.MThreadCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceListMCoreCount(t *testing.T) {
	tests := []struct {
		name    string
		r       ResourceList
		want    int64
		wantErr bool
	}{
		{
			name: "Valid extraction",
			r: ResourceList{
				{
					Name:  "m_core",
					Type:  "hl",
					Value: "14",
				},
			},
			want:    14,
			wantErr: false,
		},
		{
			name: "No Key",
			r: ResourceList{
				{
					Name:  "free_mem",
					Type:  "hl",
					Value: "2G",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MCoreCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceList.MCoreCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceList.MCoreCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
