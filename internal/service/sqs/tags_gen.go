// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists sqs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn sqsiface.SQSAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &sqs.ListQueueTagsInput{
		QueueUrl: aws.String(identifier),
	}

	output, err := conn.ListQueueTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.Tags), nil
}

// map[string]*string handling

// Tags returns sqs service tags.
func Tags(tags tftags.KeyValueTags) map[string]*string {
	return aws.StringMap(tags.Map())
}

// KeyValueTags creates KeyValueTags from sqs service tags.
func KeyValueTags(tags map[string]*string) tftags.KeyValueTags {
	return tftags.New(tags)
}

// UpdateTags updates sqs service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn sqsiface.SQSAPI, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &sqs.UntagQueueInput{
			QueueUrl: aws.String(identifier),
			TagKeys:  aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.UntagQueue(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &sqs.TagQueueInput{
			QueueUrl: aws.String(identifier),
			Tags:     Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagQueue(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
