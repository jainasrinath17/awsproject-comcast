output "instanceip" {
  description = "Public IP of the web server"
  value       = aws_instance.web_server.public_ip
}
