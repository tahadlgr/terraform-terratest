output "ec2_instance_id" {
  value = aws_instance.ec2_sample.id
}

output "s3_bucket_name" {
  value = aws_s3_bucket.bucket_sample.bucket
}