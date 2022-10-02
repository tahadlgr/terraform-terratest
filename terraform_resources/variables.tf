variable "region" {
  default = "eu-central-1"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "instance_ami" {
  default = "ami-05ff5eaef6149df49"
}

variable "bucket_name" {
  default = "unique-bucket-name-flugel-0325266548636552332"
}

# variable "vpc_name" {
#   default = "vpc-0d3fdfda4ca58d9eb"
# }

variable "subnet_name" {
  default = "subnet-0603cfacef0e68a7c"
}

variable "tags" {
  default = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}
