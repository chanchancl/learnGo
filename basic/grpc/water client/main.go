package main

import (
	"log"

	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:xxxx"
)

func main() {
	// 1. Dial a connection
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer con.Close()

	// // 2. Create a client use con
	// client := pm.NewMetricsClient(con)

	// metric := &pm.Metric{
	// 	Name:        "Serve_Request",
	// 	Description: "Count of Server_Request",
	// 	Type:        pm.MetricType_COUNTER,
	// 	Labels:      []string{"type"},
	// 	Values:      []string{"200"},
	// }

	// // 3. create a context for operation
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // 4. register metric
	// client.Register(ctx, metric)

	// // 5. operate the metric
	// go func() {
	// 	for {
	// 		metric.Values[0] = "200"
	// 		log.Printf("Increase :%v", metric)
	// 		client.Increase(ctx, metric)

	// 		if rand.Float32() < 0.5 {
	// 			metric.Values[0] = "404"
	// 			log.Printf("Increase :%v", metric)
	// 			client.Increase(ctx, metric)
	// 		}

	// 		time.Sleep(time.Second)
	// 	}
	// }()

	select {}
}
