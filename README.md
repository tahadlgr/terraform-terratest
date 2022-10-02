## terraform-terratest-assignment

While preparing this assignment, followed the steps below:

- Prepared desired terraform files to create EC2 instance and S3 bucket. EC2 instance type will be "t2.micro" and associated Amazon Machine Image found in AMI Catalog on AWS. For S3 bucket, a unique bucket name is given. Both resources will be created on eu-central-1 region and both resources will have "Name" : "Flugel", "Owner": "InfraTeam" tags. All variables in this assignment are defined in variables.tf to be used in other files.

- Prepared outputs.tf to be used while testing. Here EC2 instance id and S3 bucket name that we created will be outputted.

- In the local, run **terraform init** to initialize a working directory containing Terraform configuration files, and **terraform plan** command to see which actions will work appropriately.

- Since terratest is a Go library, used Go to prepare test file. All steps were commented in this file to explain code. Here we use terratest to deploy our resources. After that, we fetched terraform outputs that we wrote in outputs.tf file. The test will examine that whether the tags of our resources tagged as expected. At the end of the test, resources will be undeployed on AWS.

- Used **go mod init** and **go mod tidy** to initialize and to get necessary imports for the test, then used **go test -v** command to check whether tests work appropriately.

- Prepared **workflow.yml** file under .github/workflows to setup a GitHub Action and run the test on it. Repository secrets are created to use AWS access and secret keys in the workflow.