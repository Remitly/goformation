// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_CurrencyDisplayFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.CurrencyDisplayFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html
type Analysis_CurrencyDisplayFormatConfiguration struct {

	// DecimalPlacesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-decimalplacesconfiguration
	DecimalPlacesConfiguration *Analysis_DecimalPlacesConfiguration `json:"DecimalPlacesConfiguration,omitempty"`

	// NegativeValueConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-negativevalueconfiguration
	NegativeValueConfiguration *Analysis_NegativeValueConfiguration `json:"NegativeValueConfiguration,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Analysis_NullValueFormatConfiguration `json:"NullValueFormatConfiguration,omitempty"`

	// NumberScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-numberscale
	NumberScale *string `json:"NumberScale,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-prefix
	Prefix *string `json:"Prefix,omitempty"`

	// SeparatorConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-separatorconfiguration
	SeparatorConfiguration *Analysis_NumericSeparatorConfiguration `json:"SeparatorConfiguration,omitempty"`

	// Suffix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-suffix
	Suffix *string `json:"Suffix,omitempty"`

	// Symbol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-currencydisplayformatconfiguration.html#cfn-quicksight-analysis-currencydisplayformatconfiguration-symbol
	Symbol *string `json:"Symbol,omitempty"`

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
func (r *Analysis_CurrencyDisplayFormatConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.CurrencyDisplayFormatConfiguration"
}
