resource "aws_s3_bucket" "bucket_sample" {
  bucket = var.bucket_name

  tags = var.tags
}