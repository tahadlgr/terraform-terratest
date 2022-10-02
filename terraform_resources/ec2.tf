resource "aws_instance" "ec2_sample" {
  instance_type = var.instance_type
  ami           = var.instance_ami
  subnet_id = var.subnet_name

  tags = var.tags
}

