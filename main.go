package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	mcpServer := server.NewMCPServer("weather", "0.0.1", server.WithToolCapabilities(true))

	// Add a weather forecast tool
	weatherForecastTool := mcp.NewTool("weather_forecast",
		mcp.WithDescription("Get weather forecast"),
		mcp.WithString("Latitude",
			mcp.Required(),
			mcp.Description("Latitude of the location"),
		),
		mcp.WithString("Longitude",
			mcp.Required(),
			mcp.Description("Longitude of the location"),
		),
		mcp.WithString("Time",
			mcp.Required(),
			mcp.Description("Time of the forecast"),
		),
	)

	mcpServer.AddTool(weatherForecastTool, weatherForecastHandler)

	// transport
	httpServer := server.NewStreamableHTTPServer(mcpServer)
	if err := httpServer.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func weatherForecastHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// extract parameters from the request
	arguments := request.Params.Arguments.(map[string]interface{})
	latitude := arguments["Latitude"].(string)
	longitude := arguments["Longitude"].(string)
	time := arguments["Time"].(string)

	// call to open-meteo API to get weather forecast
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m&timezone=auto", latitude, longitude)
	resp, err := http.Get(url)
	if err != nil {
		return mcp.NewToolResultError("failed to fetch weather data"), nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return mcp.NewToolResultError("failed to fetch weather data"), nil
	}
	var data struct {
		Hourly struct {
			Temperature      []float64 `json:"temperature_2m"`
			RelativeHumidity []float64 `json:"relative_humidity_2m"`
			WindSpeed        []float64 `json:"wind_speed_10m"`
		} `json:"hourly"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return mcp.NewToolResultError("failed to parse weather data"), nil
	}

	// Prepare response
	if len(data.Hourly.Temperature) == 0 {
		return mcp.NewToolResultError("no temperature data available"), nil
	}
	temperature := data.Hourly.Temperature[0]
	humidity := data.Hourly.RelativeHumidity[0]
	windSpeed := data.Hourly.WindSpeed[0]
	return mcp.NewToolResultText(fmt.Sprintf("The temperature at %s, %s on %s is %.2f°C with %.2f%% humidity and a wind speed of %.2f m/s", latitude, longitude, time, temperature, humidity, windSpeed)), nil
}
