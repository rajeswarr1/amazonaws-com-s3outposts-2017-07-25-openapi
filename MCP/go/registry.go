package main

import (
	"github.com/amazon-s3-on-outposts/mcp-server/config"
	"github.com/amazon-s3-on-outposts/mcp-server/models"
	tools_s3outposts "github.com/amazon-s3-on-outposts/mcp-server/tools/s3outposts"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_s3outposts.CreateListendpointsTool(cfg),
		tools_s3outposts.CreateListoutpostswiths3Tool(cfg),
		tools_s3outposts.CreateListsharedendpointsTool(cfg),
		tools_s3outposts.CreateCreateendpointTool(cfg),
		tools_s3outposts.CreateDeleteendpointTool(cfg),
	}
}
