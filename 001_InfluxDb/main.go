package main

import (
    "context"
    "fmt"
    "time"

    "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
    // Create a new client using an InfluxDB server base URL and an authentication token
	const token = "VYwuwT3YNwffTJWTKytqeckOcRxPdW9X3ff8V72kx4Ykw9DvW-Wo_B875yyl00et81wzuCT9cSEC12RsSUT8sw=="
	const bucket = "test1"
	const org = "IDEA"

	client := influxdb2.NewClient("http://localhost:8086", token)

    // Use blocking write client for writes to desired bucket
    writeAPI := client.WriteAPIBlocking("IDEA", "test1")
    // Create point using full params constructor 
    p := influxdb2.NewPoint("stat",
        map[string]string{"unit": "temperature"},
        map[string]interface{}{"avg": 24.5, "max": 45.0},
        time.Now())
    // write point immediately 
    writeAPI.WritePoint(context.Background(), p)
    // Create point using fluent style
    p = influxdb2.NewPointWithMeasurement("stat").
        AddTag("unit", "temperature").
        AddField("avg", 23.2).
        AddField("max", 45.0).
        SetTime(time.Now())
    writeAPI.WritePoint(context.Background(), p)
    
    // Or write directly line protocol
    line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
    writeAPI.WriteRecord(context.Background(), line)

    
    // Ensures background processes finishes
    client.Close()
}