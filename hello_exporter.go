package helloexporter

import (
	"context"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// NewFactory creates a new factory for HelloExporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		"helloexporter",
		createDefaultConfig,
		exporterhelper.WithTraces(createTracesExporter),
		exporterhelper.WithLogs(createLogsExporter),
		exporterhelper.WithMetrics(createMetricsExporter),
	)
}

// 配置结构体
func createDefaultConfig() component.Config {
	return &struct{}{}
}

// 处理 Traces
func createTracesExporter(ctx context.Context, params component.ExporterCreateSettings, cfg component.Config) (component.TracesExporter, error) {
	return exporterhelper.NewTracesExporter(ctx, params, cfg, func(ctx context.Context, td ptrace.Traces) error {
		fmt.Println("Hello, World! Traces Exported")
		return nil
	})
}

// 处理 Logs
func createLogsExporter(ctx context.Context, params component.ExporterCreateSettings, cfg component.Config) (component.LogsExporter, error) {
	return exporterhelper.NewLogsExporter(ctx, params, cfg, func(ctx context.Context, ld plog.Logs) error {
		fmt.Println("Hello, World! Logs Exported")
		return nil
	})
}

// 处理 Metrics
func createMetricsExporter(ctx context.Context, params component.ExporterCreateSettings, cfg component.Config) (component.MetricsExporter, error) {
	return exporterhelper.NewMetricsExporter(ctx, params, cfg, func(ctx context.Context, md pmetric.Metrics) error {
		fmt.Println("Hello, World! Metrics Exported")
		return nil
	})
}

