terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "lightsail"
    region = "us-east-1"
  }
}
