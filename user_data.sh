#!/bin/bash
sudo yum update -y
sudo yum install -y httpd mod_ssl
sudo systemctl start httpd
sudo systemctl enable httpd

# Create SSL certificates (self-signed)
sudo mkdir /etc/ssl/private
sudo openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 \
  -keyout /etc/ssl/private/apache-selfsigned.key \
  -out /etc/ssl/certs/apache-selfsigned.crt \
  -subj "/C=US/ST=New York/L=New York/O=Company Name/OU=Org/CN=example.com"

# Update Apache config to redirect HTTP to HTTPS
sudo bash -c 'cat > /etc/httpd/conf.d/ssl.conf << EOF
<VirtualHost *:80>
  ServerName example.com
  Redirect / https://example.com/
</VirtualHost>

<VirtualHost *:443>
  DocumentRoot "/var/www/html"
  ServerName example.com
  SSLEngine on
  SSLCertificateFile /etc/ssl/certs/apache-selfsigned.crt
  SSLCertificateKeyFile /etc/ssl/private/apache-selfsigned.key
</VirtualHost>
EOF'

# Add content to the web server
sudo bash -c 'cat > /var/www/html/index.html << EOF
<html>
<head>
<title>Hello World</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>
EOF'

sudo systemctl restart httpd
