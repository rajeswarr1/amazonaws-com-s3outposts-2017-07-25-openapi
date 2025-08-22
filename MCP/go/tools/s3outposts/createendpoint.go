package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-s3-on-outposts/mcp-server/config"
	"github.com/amazon-s3-on-outposts/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateendpointHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/S3Outposts/CreateEndpoint", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CreateEndpointResult
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

func CreateCreateendpointTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_S3Outposts_CreateEndpoint",
		mcp.WithDescription("<p>Creates an endpoint and associates it with the specified Outpost.</p> <note> <p>It can take up to 5 minutes for this action to finish.</p> </note> <p/> <p>Related actions include:</p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html">DeleteEndpoint</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_ListEndpoints.html">ListEndpoints</a> </p> </li> </ul>"),
		mcp.WithString("AccessType", mcp.Description("Input parameter: <p>The type of access for the network connectivity for the Amazon S3 on Outposts endpoint. To use the Amazon Web Services VPC, choose <code>Private</code>. To use the endpoint with an on-premises network, choose <code>CustomerOwnedIp</code>. If you choose <code>CustomerOwnedIp</code>, you must also provide the customer-owned IP address pool (CoIP pool).</p> <note> <p> <code>Private</code> is the default access type value.</p> </note>")),
		mcp.WithString("CustomerOwnedIpv4Pool", mcp.Description("Input parameter: The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint. IP addresses are allocated from this pool for the endpoint.")),
		mcp.WithString("OutpostId", mcp.Required(), mcp.Description("Input parameter: The ID of the Outposts. ")),
		mcp.WithString("SecurityGroupId", mcp.Required(), mcp.Description("Input parameter: The ID of the security group to use with the endpoint.")),
		mcp.WithString("SubnetId", mcp.Required(), mcp.Description("Input parameter: The ID of the subnet in the selected VPC. The endpoint subnet must belong to the Outpost that has Amazon S3 on Outposts provisioned.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateendpointHandler(cfg),
	}
}
