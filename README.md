# go-route53-dyn-dns #

A simple command to update AWS Route53 with a device's public IP address.

# Uses #

* On boot of an EC2 instance, update the A record of a (sub)domain to point to the new IP address.
* Your own dynamic DNS.

# Requirements #

* Linux. Windows, OS X, and BSD are untested.
* An internet connection and a public IP address.
* An AWS account.
* An existing hosted zone in AWS Route53.

# Configuration #

Create a configuration file at `/etc/go-route53-dyn-dns-conf.json` with contents:

    {
        "AWS_ACCESS_KEY": "YOUR_ACCESS_KEY",
        "AWS_SECRET_ACCESS_KEY": "YOUR_SECRET_ACCESS_KEY",
        "HOSTED_ZONE_ID": "YOUR_HOSTED_ZONE_ID",
        "SUBDOMAIN": "mysubdomain.mydomain.com"
    }

# Running #

Once your configuration file is ready, just run go-route53-dyn-dns. If successful, the following will be printed:

    Successfully mapped mysubdomain.mydomain.com to IP 55.44.33.22.

# Running on Boot - Ubuntu 14.04 LTS #

To run the script on boot once the network is up, create the file `/etc/init/go-route53-dyn-dns.conf` and paste these contents into it:

    description "go-route53-dyn-dns"

    start on net-device-up IFACE=eth0

    script
        /usr/local/bin/go-route53-dyn-dns
    end script

Remember to replace the location of `go-route53-dyn-dns`.
