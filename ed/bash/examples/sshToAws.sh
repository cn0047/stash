#!/bin/bash

# sudo chmod +x /home/kovpak/web/kovpak/gh/ed/bash/examples/sshToAws.sh
# sudo ln -s /home/kovpak/web/kovpak/gh/ed/bash/examples/sshToAws.sh /usr/bin/sshToAws

if [ -z "$1" ]; then
    echo 'Please specify host.'
    break
fi
host=$1
user='ec2-user'
if [ "$2" ]; then
    user=$2
fi
uri="$user@$host.eu-west-1.compute.amazonaws.com"
ssh -i /home/kovpak/web/storage/ziipr.pem "$uri"
