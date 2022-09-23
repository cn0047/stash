VPC
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
