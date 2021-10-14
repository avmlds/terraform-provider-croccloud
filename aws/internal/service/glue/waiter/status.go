package waiter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/service/glue/finder"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

const (
	MLTransformStatusUnknown   = "Unknown"
	RegistryStatusUnknown      = "Unknown"
	SchemaStatusUnknown        = "Unknown"
	SchemaVersionStatusUnknown = "Unknown"
	TriggerStatusUnknown       = "Unknown"
)

// MLTransformStatus fetches the MLTransform and its Status
func MLTransformStatus(conn *glue.Glue, transformId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &glue.GetMLTransformInput{
			TransformId: aws.String(transformId),
		}

		output, err := conn.GetMLTransform(input)

		if err != nil {
			return nil, MLTransformStatusUnknown, err
		}

		if output == nil {
			return output, MLTransformStatusUnknown, nil
		}

		return output, aws.StringValue(output.Status), nil
	}
}

// RegistryStatus fetches the Registry and its Status
func RegistryStatus(conn *glue.Glue, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := finder.RegistryByID(conn, id)
		if err != nil {
			return nil, RegistryStatusUnknown, err
		}

		if output == nil {
			return output, RegistryStatusUnknown, nil
		}

		return output, aws.StringValue(output.Status), nil
	}
}

// SchemaStatus fetches the Schema and its Status
func SchemaStatus(conn *glue.Glue, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := finder.SchemaByID(conn, id)
		if err != nil {
			return nil, SchemaStatusUnknown, err
		}

		if output == nil {
			return output, SchemaStatusUnknown, nil
		}

		return output, aws.StringValue(output.SchemaStatus), nil
	}
}

// SchemaVersionStatus fetches the Schema Version and its Status
func SchemaVersionStatus(conn *glue.Glue, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := finder.SchemaVersionByID(conn, id)
		if err != nil {
			return nil, SchemaVersionStatusUnknown, err
		}

		if output == nil {
			return output, SchemaVersionStatusUnknown, nil
		}

		return output, aws.StringValue(output.Status), nil
	}
}

// TriggerStatus fetches the Trigger and its Status
func TriggerStatus(conn *glue.Glue, triggerName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &glue.GetTriggerInput{
			Name: aws.String(triggerName),
		}

		output, err := conn.GetTrigger(input)

		if err != nil {
			return nil, TriggerStatusUnknown, err
		}

		if output == nil {
			return output, TriggerStatusUnknown, nil
		}

		return output, aws.StringValue(output.Trigger.State), nil
	}
}

func GlueDevEndpointStatus(conn *glue.Glue, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := finder.DevEndpointByName(conn, name)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}

func GluePartitionIndexStatus(conn *glue.Glue, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := finder.PartitionIndexByName(conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.IndexStatus), nil
	}
}
