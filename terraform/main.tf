provider "aws" {
  region = "eu-north-1"
}

resource "aws_security_group" "example_sg" {
  name        = "example-security-group"
  description = "Allow SSH, HTTP, HTTPS, and ICMP"

  ingress {
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow ICMP"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "example-security-group"
  }
}

resource "aws_instance" "web_server" {
  ami           = "ami-0705384c0b33c194c"  # Example Ubuntu AMI ID (ensure it's a valid Ubuntu AMI ID for your region)
  instance_type = "t3.micro"
  security_groups = [aws_security_group.example_sg.name]
  key_name = "keypair"

  tags = {
    Name = "url-shortener-server"
  }

}

output "instance_ip" {
  value = aws_instance.web_server.public_ip
}
