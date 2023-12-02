package goids

import (
	"math"
	"testing"
)

func TestLen(t *testing.T) {
	type args struct {
		v Vector
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test 1", args{Vector{3, 4}}, 5},
		{"Test 2", args{Vector{0, 0}}, 0},
		{"Test 3", args{Vector{3, -1}}, math.Sqrt(10)},
		{"Test 4", args{Vector{-1, 2}}, math.Sqrt(5)},
		{"Test 5", args{Vector{2, 2}}, math.Sqrt(8)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.v.Len(); got != tt.want {
				t.Errorf("Vector.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScalarMul(t *testing.T) {
	type args struct {
		v Vector
		c float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, 2}, Vector{6, 8}},
		{"Test 2", args{Vector{0, 0}, 2}, Vector{0, 0}},
		{"Test 3", args{Vector{3, -1}, 2}, Vector{6, -2}},
		{"Test 4", args{Vector{-1, 2}, 2}, Vector{-2, 4}},
		{"Test 5", args{Vector{2, 2}, 2}, Vector{4, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.ScalarMul(tt.args.c)
			if tt.args.v != tt.want {
				t.Errorf("Vector.ScalarMul() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		v Vector
		l float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, 2}, Vector{6.0 / 5.0, 8.0 / 5.0}},
		{"Test 2", args{Vector{0, 0}, 2}, Vector{0.0, 0.0}},
		{"Test 3", args{Vector{3, -1}, 2}, Vector{6 / math.Sqrt(10), -2 / math.Sqrt(10)}},
		{"Test 4", args{Vector{-1, 2}, 2}, Vector{-2 / math.Sqrt(5), 4 / math.Sqrt(5)}},
		{"Test 5", args{Vector{2, 2}, 2}, Vector{4 / math.Sqrt(8), 4 / math.Sqrt(8)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.Scale(tt.args.l)
			if (tt.args.v.X-tt.want.X) > 0.0000001 || (tt.args.v.Y-tt.want.Y) > 0.0000001 {
				t.Errorf("Vector.Scale() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	type args struct {
		v Vector
		l float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, 2}, Vector{6.0 / 5.0, 8.0 / 5.0}},
		{"Test 2", args{Vector{0, 0}, 2}, Vector{0.0, 0.0}},
		{"Test 3", args{Vector{3, -1}, 2}, Vector{6 / math.Sqrt(10), -2 / math.Sqrt(10)}},
		{"Test 4", args{Vector{-1, 2}, 5}, Vector{-1, 2}},
		{"Test 5", args{Vector{2, 2}, 2}, Vector{4 / math.Sqrt(8), 4 / math.Sqrt(8)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.Limit(tt.args.l)
			if (tt.args.v.X-tt.want.X) > 0.0000001 || (tt.args.v.Y-tt.want.Y) > 0.0000001 {
				t.Errorf("Vector.Limit() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		v  Vector
		v2 Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, Vector{1, 2}}, Vector{4, 6}},
		{"Test 2", args{Vector{0, 0}, Vector{1, 2}}, Vector{1, 2}},
		{"Test 3", args{Vector{3, -1}, Vector{1, 2}}, Vector{4, 1}},
		{"Test 4", args{Vector{-1, 2}, Vector{1, 2}}, Vector{0, 4}},
		{"Test 5", args{Vector{2, 2}, Vector{1, 2}}, Vector{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.Add(tt.args.v2)
			if tt.args.v != tt.want {
				t.Errorf("Vector.Add() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestSubMethod(t *testing.T) {
	type args struct {
		v  Vector
		v2 Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, Vector{1, 2}}, Vector{2, 2}},
		{"Test 2", args{Vector{0, 0}, Vector{1, 2}}, Vector{-1, -2}},
		{"Test 3", args{Vector{3, -1}, Vector{1, 2}}, Vector{2, -3}},
		{"Test 4", args{Vector{-1, 2}, Vector{1, 2}}, Vector{-2, 0}},
		{"Test 5", args{Vector{2, 2}, Vector{-1, 2}}, Vector{3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.Sub(tt.args.v2)
			if tt.args.v != tt.want {
				t.Errorf("Vector.Sub() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestSubFunc(t *testing.T) {
	type args struct {
		v2 Vector
		v1 Vector
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, Vector{1, 2}}, Vector{2, 2}},
		{"Test 2", args{Vector{0, 0}, Vector{1, 2}}, Vector{-1, -2}},
		{"Test 3", args{Vector{3, -1}, Vector{1, 2}}, Vector{2, -3}},
		{"Test 4", args{Vector{-1, 2}, Vector{1, 2}}, Vector{-2, 0}},
		{"Test 5", args{Vector{2, 2}, Vector{-1, 2}}, Vector{3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.v2, tt.args.v1); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
