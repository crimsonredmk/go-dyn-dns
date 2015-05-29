# go-dyn-dns #

Tool to update AWS Route53 with an EC2 instance's public IP address.

# Configuration #

Create a configuration file at `/etc/go-dyn-dns-conf.json` with contents:

    {
        "AWS_ACCESS_KEY": "YOUR_ACCESS_KEY",
        "AWS_SECRET_ACCESS_KEY": "YOUR_SECRET_ACCESS_KEY",
        "HOSTED_ZONE_ID": "YOUR_HOSTED_ZONE_ID",
        "SUBDOMAIN": "mysubdomain.mydomain.com"
    }
