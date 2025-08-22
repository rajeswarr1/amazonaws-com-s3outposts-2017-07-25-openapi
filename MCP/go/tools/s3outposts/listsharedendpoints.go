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

func ListsharedendpointsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["outpostId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outpostId=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/S3Outposts/ListSharedEndpoints#outpostId%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ListSharedEndpointsResult
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

func CreateListsharedendpointsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_S3Outposts_ListSharedEndpoints#outpostId",
		mcp.WithDescription("<p>Lists all endpoints associated with an Outpost that has been shared by Amazon Web Services Resource Access Manager (RAM).</p> <p>Related actions include:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html">CreateEndpoint</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html">DeleteEndpoint</a> </p> </li> </ul>"),
		mcp.WithString("nextToken", mcp.Description("If a previous response from this operation included a <code>NextToken</code> value, you can provide that value here to retrieve the next page of results.")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of endpoints that will be returned in the response.")),
		mcp.WithString("outpostId", mcp.Required(), mcp.Description("The ID of the Amazon Web Services Outpost.")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListsharedendpointsHandler(cfg),
	}
}
