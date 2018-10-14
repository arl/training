package main

import (
	"math"
	"testing"
)

func Test_digitValue(t *testing.T) {
	tests := []struct {
		d      byte
		want   uint64
		wantOk bool
	}{
		{
			d:      'a',
			want:   10,
			wantOk: true,
		},
		{
			d:      '0',
			want:   0,
			wantOk: true,
		},
		{
			d:      'Z',
			want:   61,
			wantOk: true,
		},
		{
			d:      0,
			want:   0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		gotVal, gotOk := digitValue(tt.d)
		if gotOk != tt.wantOk {
			t.Fatalf("digitValue(%c) got ok = %v, want %v\n", tt.d, gotOk, tt.wantOk)
		}
		if gotVal != tt.want {
			t.Errorf("digitValue(%c) got val = %v, want %v\n", tt.d, gotVal, tt.want)
		}
	}
}

func Test_toBase62(t *testing.T) {
	tests := []struct {
		id   uint64
		want string
	}{
		{
			id:   0,
			want: "00000000000",
		},
		{
			id:   61,
			want: "0000000000Z",
		},
		{
			id:   62,
			want: "00000000010",
		},
		{
			id:   63,
			want: "00000000011",
		},
		{
			id:   math.MaxUint64,
			want: "lYGhA16ahyf",
		},
	}

	for _, tt := range tests {
		if got := toBase62(tt.id); got != tt.want {
			t.Errorf("toBase62(%v) = '%v', want '%v'\n", tt.id, got, tt.want)
		}
	}
}

func Test_fromBase62(t *testing.T) {
	tests := []struct {
		name    string
		digits  string
		want    uint64
		wantErr bool
	}{
		{
			digits: "00000000000",
			want:   0,
		},
		{
			digits: "0000000000Z",
			want:   61,
		},
		{
			digits: "00000000010",
			want:   62,
		},
		{
			digits: "11",
			want:   63,
		},
		{
			digits: "lYGhA16ahyf",
			want:   math.MaxUint64,
		},
		// {
		// 	digits: "ZYGhA16ahyf",
		// 	want:   math.MaxUint64,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromBase62(tt.digits)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromBase62('%s') error = %v, wantErr %v\n", tt.digits, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fromBase62('%s') = %v, want %v\n", tt.digits, got, tt.want)
			}
		})
	}
}
