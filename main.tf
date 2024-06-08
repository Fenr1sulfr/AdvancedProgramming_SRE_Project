"provider "aws" {
  region = "eu-north-1"
}


resource "aws_instance" "web_server" {
  ami           = "ami-0705384c0b33c194c"  # Example Ubuntu AMI ID (ensure it's a valid Ubuntu AMI ID for your region)
  instance_type = "t3.micro"
  security_groups = ["default"]
  key_name = "keypair"

  tags = {
    Name = "url-shortener-server"
  }

}

output "instance_ip" {
  value = aws_instance.web_server.public_ip
}
"