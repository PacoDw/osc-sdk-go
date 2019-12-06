/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// ClientGateway Information about the client gateway.
type ClientGateway struct {
	// An unsigned 32-bits Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find out the path to your client gateway through the Internet network.
	BgpAsn int32 `json:"BgpAsn,omitempty"`
	// The ID of the client gateway.
	ClientGatewayId string `json:"ClientGatewayId,omitempty"`
	// The type of communication tunnel used by the client gateway (only `ipsec.1` is supported).
	ConnectionType string `json:"ConnectionType,omitempty"`
	// The public IPv4 address of the client gateway (must be a fixed address into a NATed network).
	PublicIp string `json:"PublicIp,omitempty"`
	// The state of the client gateway (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the client gateway.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
