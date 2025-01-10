package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func initTracer() (*sdktrace.TracerProvider, error) {
	// Set up the Jaeger exporter
	endpoint := os.Getenv("JAEGER_ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:14268/api/traces"
	}
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, err
	}

	// Set up the resource (application/service identity)
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName("my-go-app"),
		),
	)
	if err != nil {
		return nil, err
	}

	// Create a trace provider with the Jaeger exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()

	tracer = otel.Tracer("my-go-app")

	r := chi.NewRouter()

	// Middleware to automatically create spans for HTTP requests
	r.Use(otelhttp.NewMiddleware("my-server"))

	// Define a handler with tracing
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Start a new span for this request
		ctx, span := tracer.Start(r.Context(), "GET /")
		defer span.End()

		// Simulate some processing and add attributes to the span
		logger := log.New(os.Stdout, "", log.LstdFlags)
		logger.Println("Received request at /")

		span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
		span.SetAttributes(semconv.HTTPRouteKey.String("/"))
		span.AddEvent("Processing request")
		// Simulate a database call (or any other action)
		databaseCall(ctx)

		w.Write([]byte("Hello, World!"))

		span.AddEvent("Response sent")
	})

	// Start the server with tracing middleware
	http.ListenAndServe(":8080", otelhttp.NewHandler(r, "my-server"))
}

// Simulating a function that would trigger a new span
func databaseCall(ctx context.Context) {
	_, span := tracer.Start(ctx, "Database Call")
	defer span.End()

	// Simulate work
	fmt.Println("Simulating a database call...")
}
