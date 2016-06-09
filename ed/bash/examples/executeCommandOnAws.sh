#!/bin/bash

# sudo chmod +x /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh
# sudo ln -s /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh /usr/bin/ecoa

if [ -z "$1" ]; then
    echo 'Please specify filter.'
    exit 1
fi
tagLike=$1

if [ -z "$2" ]; then
    echo 'Please specify command.'
    exit 2
fi
commandStr=$2

hosts=$(\
aws ec2 describe-instances \
--output text --query 'Reservations[*].Instances[*].[PublicDnsName]' \
--filter Name=tag:Name,Values=*$tagLike* \
)

for host in $hosts; do
    printf "\n \033[34m $host \033[0m \n\n"
    ssh -i /home/kovpak/web/storage/ziipr.pem ec2-user@$host $commandStr
done
