// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Channel_ColorCorrection AWS CloudFormation Resource (AWS::MediaLive::Channel.ColorCorrection)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html
type Channel_ColorCorrection struct {

	// InputColorSpace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-inputcolorspace
	InputColorSpace *string `json:"InputColorSpace,omitempty"`

	// OutputColorSpace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-outputcolorspace
	OutputColorSpace *string `json:"OutputColorSpace,omitempty"`

	// Uri AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-colorcorrection.html#cfn-medialive-channel-colorcorrection-uri
	Uri *string `json:"Uri,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Channel_ColorCorrection) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.ColorCorrection"
}
