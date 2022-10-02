package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsExample(t *testing.T) {
	t.Parallel()

	// Path for terraform resources
	terraformFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "/terraform_resources")

	awsRegion := "eu-central-1"

	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: terraformFolder,
	})

	// "terraform destroy" at the end of the test in order to clear resources
	defer terraform.Destroy(t, terraformOptions)

	// init and apply of terraform resources
	terraform.InitAndApply(t, terraformOptions)

	// terraform output for our created resources
	instanceID := terraform.Output(t, terraformOptions, "ec2_instance_id")
	bucketName := terraform.Output(t, terraformOptions, "s3_bucket_name")

	// Look up the tags for the given instance id and bucket name
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)
	bucketTags := aws.GetS3BucketTags(t, awsRegion, bucketName)

	// Expected Tags
	tags := map[string]string(map[string]string{"Name": "Flugel", "Owner": "InfraTeam"})

	// Tests
	assert.Equal(t, tags, instanceTags)
	assert.Equal(t, tags, bucketTags)
}
