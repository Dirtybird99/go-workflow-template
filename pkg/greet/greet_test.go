package greet

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"World", "Hello, World!"},
		{"Go", "Hello, Go!"},
		{"", "Hello, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.name)
			if got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Benchmark")
	}
}
