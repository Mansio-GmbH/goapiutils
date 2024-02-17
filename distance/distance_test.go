package distance_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/distance"
	"github.com/stretchr/testify/require"
)

func TestDistance_Meters(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Test 1",
			fields: fields{
				value: 1000,
			},
			want: 1000,
		},
		{
			name: "Test 2",
			fields: fields{
				value: 2000,
			},
			want: 2000,
		},
		{
			name: "Test 3",
			fields: fields{
				value: 3000,
			},
			want: 3000,
		},
		{
			name: "Test 4",
			fields: fields{
				value: 4000,
			},
			want: 4000,
		},
		{
			name: "Test 5",
			fields: fields{
				value: 5000,
			},
			want: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Distance(tt.fields.value) * distance.Meter
			require.Equal(t, tt.want, d.Meters())
		})
	}
}

func TestDistance_Kilometers(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "Test 1",
			fields: fields{
				value: 1000,
			},
			want: 1,
		},
		{
			name: "Test 2",
			fields: fields{
				value: 2000,
			},
			want: 2,
		},
		{
			name: "Test 3",
			fields: fields{
				value: 3000,
			},
			want: 3,
		},
		{
			name: "Test 4",
			fields: fields{
				value: 4000,
			},
			want: 4,
		},
		{
			name: "Test 5",
			fields: fields{
				value: 5000,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Distance(tt.fields.value) * distance.Meter
			require.Equal(t, tt.want, d.Kilometers())
		})
	}
}

func TestDistance_Truncate(t *testing.T) {
	type fields struct {
		value int64
	}
	type args struct {
		m distance.Distance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   distance.Distance
	}{
		{
			name: "Test 1",
			fields: fields{
				value: 1000,
			},
			args: args{
				m: 100,
			},
			want: 1000,
		},
		{
			name: "Test 2",
			fields: fields{
				value: 2000,
			},
			args: args{
				m: 100,
			},
			want: 2000,
		},
		{
			name: "Test 3",
			fields: fields{
				value: 3000,
			},
			args: args{
				m: 100,
			},
			want: 3000,
		},
		{
			name: "Test 4",
			fields: fields{
				value: 4000,
			},
			args: args{
				m: 100,
			},
			want: 4000,
		},
		{
			name: "Test 5",
			fields: fields{
				value: 5000,
			},
			args: args{
				m: 100,
			},
			want: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Distance(tt.fields.value) * distance.Meter
			require.Equal(t, tt.want, d.Truncate(tt.args.m))
		})
	}
}
