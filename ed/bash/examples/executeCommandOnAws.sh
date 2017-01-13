#!/bin/bash

# sudo apt-get install awscli
# sudo chmod +x /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh
# sudo ln -s /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh /usr/bin/ecoa

if [ -z "$1" ]; then
    echo 'Please specify filter.'
fi
tagLike=$1

if [ -z "$2" ]; then
    echo 'Please specify command.'
fi
commandStr=$2

keyFile='/home/kovpak/web/storage/ziipr.pem'
if [ ! -f "$file" ]; then
    keyFile='/vagrant/ziipr.pem'
fi
if [ ! -f "$file" ]; then
    keyFile='/Users/kvol/web/storage/ziipr.pem'
fi

hosts=$(\
aws ec2 describe-instances \
--output text --query 'Reservations[*].Instances[*].[PublicDnsName]' \
--filter Name=tag:Name,Values=*$tagLike* \
)

for host in $hosts; do
    if [ $host = 'None' ];
    then
        continue
    fi
    tagName=$(
        aws ec2 describe-instances \
        --output text --query 'Reservations[*].Instances[*].[Tags[0].Value]' \
        --filter Name=dns-name,Values=$host
    )
    printf "\n"
    printf "\033[32m $tagName \033[0m"
    printf "\033[34m $host \033[0m"
    printf "\n\n"
    ssh -i $keyFile ec2-user@$host $commandStr
done
