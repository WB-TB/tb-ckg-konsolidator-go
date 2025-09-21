package main

import (
	"fhir-sirs/pkg/api"
	//"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	// Start Datadog to analyse traffic
	//tracer.Start(tracer.WithDebugMode(false))
	//defer tracer.Stop()

	// Call the entrypoint
	api.Start()
}
