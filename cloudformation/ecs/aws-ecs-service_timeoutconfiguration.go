// Code generated by "go generate". Please don't change this file directly.

package ecs

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Service_TimeoutConfiguration AWS CloudFormation Resource (AWS::ECS::Service.TimeoutConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html
type Service_TimeoutConfiguration struct {

	// IdleTimeoutSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html#cfn-ecs-service-timeoutconfiguration-idletimeoutseconds
	IdleTimeoutSeconds *int `json:"IdleTimeoutSeconds,omitempty"`

	// PerRequestTimeoutSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-timeoutconfiguration.html#cfn-ecs-service-timeoutconfiguration-perrequesttimeoutseconds
	PerRequestTimeoutSeconds *int `json:"PerRequestTimeoutSeconds,omitempty"`

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
func (r *Service_TimeoutConfiguration) AWSCloudFormationType() string {
	return "AWS::ECS::Service.TimeoutConfiguration"
}
