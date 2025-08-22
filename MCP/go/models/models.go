package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteEndpointRequest represents the DeleteEndpointRequest schema from the OpenAPI specification
type DeleteEndpointRequest struct {
}

// FailedReason represents the FailedReason schema from the OpenAPI specification
type FailedReason struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

// NetworkInterface represents the NetworkInterface schema from the OpenAPI specification
type NetworkInterface struct {
	Networkinterfaceid interface{} `json:"NetworkInterfaceId,omitempty"`
}

// CreateEndpointResult represents the CreateEndpointResult schema from the OpenAPI specification
type CreateEndpointResult struct {
	Endpointarn interface{} `json:"EndpointArn,omitempty"`
}

// ListEndpointsRequest represents the ListEndpointsRequest schema from the OpenAPI specification
type ListEndpointsRequest struct {
}

// CreateEndpointRequest represents the CreateEndpointRequest schema from the OpenAPI specification
type CreateEndpointRequest struct {
	Accesstype interface{} `json:"AccessType,omitempty"`
	Customerownedipv4pool interface{} `json:"CustomerOwnedIpv4Pool,omitempty"`
	Outpostid interface{} `json:"OutpostId"`
	Securitygroupid interface{} `json:"SecurityGroupId"`
	Subnetid interface{} `json:"SubnetId"`
}

// ListSharedEndpointsResult represents the ListSharedEndpointsResult schema from the OpenAPI specification
type ListSharedEndpointsResult struct {
	Endpoints interface{} `json:"Endpoints,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListEndpointsResult represents the ListEndpointsResult schema from the OpenAPI specification
type ListEndpointsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Endpoints interface{} `json:"Endpoints,omitempty"`
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Vpcid interface{} `json:"VpcId,omitempty"`
	Accesstype interface{} `json:"AccessType,omitempty"`
	Cidrblock interface{} `json:"CidrBlock,omitempty"`
	Customerownedipv4pool interface{} `json:"CustomerOwnedIpv4Pool,omitempty"`
	Failedreason interface{} `json:"FailedReason,omitempty"`
	Outpostsid interface{} `json:"OutpostsId,omitempty"`
	Endpointarn interface{} `json:"EndpointArn,omitempty"`
	Networkinterfaces interface{} `json:"NetworkInterfaces,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Subnetid interface{} `json:"SubnetId,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Securitygroupid interface{} `json:"SecurityGroupId,omitempty"`
}

// ListSharedEndpointsRequest represents the ListSharedEndpointsRequest schema from the OpenAPI specification
type ListSharedEndpointsRequest struct {
}

// Outpost represents the Outpost schema from the OpenAPI specification
type Outpost struct {
	Outpostarn interface{} `json:"OutpostArn,omitempty"`
	Outpostid interface{} `json:"OutpostId,omitempty"`
	Ownerid interface{} `json:"OwnerId,omitempty"`
	Capacityinbytes interface{} `json:"CapacityInBytes,omitempty"`
}

// ListOutpostsWithS3Request represents the ListOutpostsWithS3Request schema from the OpenAPI specification
type ListOutpostsWithS3Request struct {
}

// ListOutpostsWithS3Result represents the ListOutpostsWithS3Result schema from the OpenAPI specification
type ListOutpostsWithS3Result struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Outposts interface{} `json:"Outposts,omitempty"`
}
