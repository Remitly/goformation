// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Topic_SemanticEntityType AWS CloudFormation Resource (AWS::QuickSight::Topic.SemanticEntityType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html
type Topic_SemanticEntityType struct {

	// SubTypeName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-subtypename
	SubTypeName *string `json:"SubTypeName,omitempty"`

	// TypeName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-typename
	TypeName *string `json:"TypeName,omitempty"`

	// TypeParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-semanticentitytype.html#cfn-quicksight-topic-semanticentitytype-typeparameters
	TypeParameters map[string]string `json:"TypeParameters,omitempty"`

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
func (r *Topic_SemanticEntityType) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.SemanticEntityType"
}
