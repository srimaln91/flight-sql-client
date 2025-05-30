package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/flight/flightsql"
	"google.golang.org/grpc"
)

func main() {
	// CLI flags
	host := flag.String("host", "localhost", "Flight SQL server hostname")
	port := flag.Int("port", 32010, "Flight SQL server port")
	query := flag.String("query", "", "SQL query to execute")
	timeout := flag.Duration("timeout", 10*time.Second, "Query timeout")
	tls := flag.Bool("tls", false, "Enable TLS (not implemented yet)")

	flag.Parse()

	if *query == "" {
		fmt.Fprintln(os.Stderr, "Error: --query must be provided")
		flag.Usage()
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%d", *host, *port)

	// Dial options (TLS disabled for now)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if *tls {
		fmt.Fprintln(os.Stderr, "TLS is not yet supported in this CLI")
		os.Exit(1)
	}

	// Create FlightSQL client
	client, err := flightsql.NewClient(address, nil, nil, opts...)
	if err != nil {
		log.Fatalf("Failed to connect to FlightSQL server: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// allocator := memory.DefaultAllocator

	fmt.Printf("Querying %s ...\n", address)
	fmt.Printf("SQL: %s\n\n", *query)

	info, err := client.Execute(ctx, *query)
	if err != nil {
		log.Fatalf("Query execution failed: %v", err)
	}

	reader, err := client.DoGet(ctx, info.Endpoint[0].Ticket)
	if err != nil {
		log.Fatalf("Failed to fetch results: %v", err)
	}
	defer reader.Release()

	// Print column headers
	schema := reader.Schema()
	var headers []string
	for _, field := range schema.Fields() {
		headers = append(headers, field.Name)
	}
	fmt.Println(strings.Join(headers, "\t"))
	fmt.Println(strings.Repeat("-", len(strings.Join(headers, "\t"))))

	// Print each row
	for reader.Next() {
		record := reader.Record()
		numRows := int(record.NumRows())
		numCols := int(record.NumCols())

		for i := 0; i < numRows; i++ {
			var row []string
			for j := 0; j < numCols; j++ {
				col := record.Column(j)
				row = append(row, formatValue(col, i))
			}
			fmt.Println(strings.Join(row, "\t"))
		}
	}
	if err := reader.Err(); err != nil {
		log.Fatalf("Error reading result: %v", err)
	}
}

// formatValue handles nulls and formats Arrow array values into strings
func formatValue(col arrow.Array, idx int) string {
	if col.IsNull(idx) {
		return "NULL"
	}
	return fmt.Sprint(col.ValueStr(idx))
}
