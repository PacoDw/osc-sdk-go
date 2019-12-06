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
// LinkPublicIpRequest struct for LinkPublicIpRequest
type LinkPublicIpRequest struct {
	// - If `true`, allows the EIP to be associated with the VM or NIC that you specify even if it is already associated with another VM or NIC.<br /> - If `false`, prevents the EIP from being associated with the VM or NIC that you specify if it is already associated with another VM or NIC.<br /><br />  (By default, `true` in the public Cloud, `false` in a Net.)
	AllowRelink bool `json:"AllowRelink,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// (Net only) The ID of the NIC. This parameter is required if the VM has more than one NIC attached. Otherwise, you need to specify the `VmId` parameter instead. You cannot specify both parameters at the same time.
	NicId string `json:"NicId,omitempty"`
	// (Net only) The primary or secondary private IP address of the specified NIC. By default, the primary private IP address.
	PrivateIp string `json:"PrivateIp,omitempty"`
	// The EIP. In the public Cloud, this parameter is required.
	PublicIp string `json:"PublicIp,omitempty"`
	// The allocation ID of the EIP. In a Net, this parameter is required.
	PublicIpId string `json:"PublicIpId,omitempty"`
	// The ID of the VM.<br /> - In the public Cloud, this parameter is required.<br /> - In a Net, this parameter is required if the VM has only one NIC. Otherwise, you need to specify the `NicId` parameter instead. You cannot specify both parameters at the same time.
	VmId string `json:"VmId,omitempty"`
}
