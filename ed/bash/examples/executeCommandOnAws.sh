#!/bin/bash

# sudo apt-get install awscli

# sudo chmod +x /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh
# sudo ln -s /home/kovpak/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh /usr/bin/ecoa

# OSX
# sudo chmod +x /Users/kvol/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh
# sudo ln -s /Users/kvol/web/kovpak/gh/ed/bash/examples/executeCommandOnAws.sh /usr/local/bin/ecoa

esc() {
    printf "\n";
    exit
}

error() {
    printf "\033[31m$1\033[0m"
}

printHost() {
    printf "\n \033[32m $1 \033[0m \033[34m $2 \033[0m \n\n"
}

highlightError() {
    printf "\033[31m$1\033[0m"
    return
}

if [ -z $1 ]; then
    error 'Please specify filter.'
    esc
fi
tagLike=$1

if [ -z "$2" ]; then
    error 'Please specify command.'
    esc
fi
commandStr=$2

keyFile='/home/kovpak/web/storage/ziipr.pem'
if [ ! -f $keyFile ]; then
    keyFile='/vagrant/ziipr.pem'
fi
if [ ! -f $keyFile ]; then
    keyFile='/Users/kvol/web/storage/ziipr.pem'
fi
keyFileStage="${keyFile/ziipr.pem/ziipr-stage.pem}"
keyFileProd="${keyFile/ziipr.pem/ziipr-prod.pem}"
declare -a keys=($keyFileProd $keyFileStage $keyFile)

hosts=$(\
    aws ec2 describe-instances \
    --output text --query 'Reservations[*].Instances[*].[PublicDnsName]' \
    --filter Name=tag:Name,Values=*$tagLike* \
)

checkKey() {
    try=$(ssh -i $1 $2 echo 200 2>&1)
    if [[ $try == *'@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@'* ]]; then
        printf "$try"
        return
    elif [[ $try == *'Permission denied'* ]]; then
        printf "$try"
        return
    else
        echo 'ok'
        return
    fi
}

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
    printHost "$tagName" "$host"
    uri="ec2-user@$host"
    #
    failMessages=()
    ok=0
    for key in "${keys[@]}"; do
        r=$(checkKey "$key" "$uri")
        if [[ $r = 'ok' ]]; then
            ok=1
            ssh -i $key $uri $commandStr
            break
        else
            msg=$(highlightError "$key $uri -> $r")
            failMessages+=("$msg")
        fi
    done
    #
    if [ $ok = '0' ]; then
        for m in "${failMessages[@]}"; do
            echo "$m"
        done
    fi
done
