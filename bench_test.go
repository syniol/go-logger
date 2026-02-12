package logger_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/syniol/go-logger"
)

// BenchmarkSyniolLogger tests the performance of your custom implementation.
func BenchmarkSyniolLogger(b *testing.B) {
	// Redirect output to Discard to measure logic, not I/O
	sylog.SetOutput(io.Discard)

	b.ResetTimer()
	b.ReportAllocs() // This is critical for the "Top 1%" look

	for i := 0; i < b.N; i++ {
		sylog.LogInfo("pay-microservice", "transaction processed", "id", 12345)
	}
}

// BenchmarkSlogJSON measures the standard library for comparison.
func BenchmarkSlogJSON(b *testing.B) {
	handler := slog.NewJSONHandler(io.Discard, nil)
	l := slog.New(handler)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l.Info("transaction processed", "facility", "pay-microservice", "id", 12345)
	}
}

// BenchmarkSlogWithSource measures slog with source-code tracing enabled.
// This is the most "apples-to-apples" comparison for your logger's trace feature.
func BenchmarkSlogWithSource(b *testing.B) {
	opts := &slog.HandlerOptions{AddSource: true}
	handler := slog.NewJSONHandler(io.Discard, opts)
	l := slog.New(handler)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l.Info("transaction processed", "facility", "pay-microservice", "id", 12345)
	}
}
