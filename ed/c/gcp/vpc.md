VPC (Virtual Private Cloud)
-

````sh
# allow ssh
name:                  default-ssh
network:               default
direction:             INGRESS
sourceRanges[0]:       0.0.0.0/0
allowed[0].IPProtocol: tcp
allowed[0].ports[0]:   22
````

````sh
gcloud compute networks subnets list
````

Subnet is per region.
Dynamic routing mode can be global or regional.
