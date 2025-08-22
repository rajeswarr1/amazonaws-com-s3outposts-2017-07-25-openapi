package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-s3-on-outposts/mcp-server/config"
	"github.com/amazon-s3-on-outposts/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeleteendpointHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["endpointId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endpointId=%v", val))
		}
		if val, ok := args["outpostId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outpostId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/S3Outposts/DeleteEndpoint#endpointId&outpostId%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDeleteendpointTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_S3Outposts_DeleteEndpoint#endpointId&outpostId",
		mcp.WithDescription("<p>Deletes an endpoint.</p> <note> <p>It can take up to 5 minutes for this action to finish.</p> </note> <p/> <p>Related actions include:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html">CreateEndpoint</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_ListEndpoints.html">ListEndpoints</a> </p> </li> </ul>"),
		mcp.WithString("endpointId", mcp.Required(), mcp.Description("The ID of the endpoint.")),
		mcp.WithString("outpostId", mcp.Required(), mcp.Description("The ID of the Outposts. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteendpointHandler(cfg),
	}
}
