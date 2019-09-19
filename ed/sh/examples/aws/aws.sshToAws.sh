#!/bin/bash

# sudo chmod +x /home/kovpak/web/kovpak/gh/ed/bash/examples/sshToAws.sh
# sudo ln -s /home/kovpak/web/kovpak/gh/ed/bash/examples/sshToAws.sh /usr/bin/sta

# OSX
# sudo chmod +x /Users/kvol/web/kovpak/gh/ed/bash/examples/sshToAws.sh
# sudo ln -s /Users/kvol/web/kovpak/gh/ed/bash/examples/sshToAws.sh /usr/local/bin/sta

esc() {
    printf "\n";
    exit
}

error() {
    printf "\033[31m$1\033[0m"
}

msg1() {
    printf "\033[31mFell down\033[0m with key: \033[33m$1\033[0m now gonna try with key: \033[36m$2\033[0m \n"
}

if [ -z $1 ]; then
    error 'Please specify host.'
    esc
fi
host=$1

user='ec2-user'
if [ $2 ]; then
    user=$2
fi

keyFile='/home/kovpak/web/storage/zii.pem'
if [ ! -f $keyFile ]; then
    keyFile='/vagrant/zii.pem'
fi
if [ ! -f $keyFile ]; then
    keyFile='/Users/kvol/web/storage/zii.pem'
fi
keyFileStage="${keyFile/zii.pem/zii-stage.pem}"
keyFileProd="${keyFile/zii.pem/zii-prod.pem}"
declare -a keys=($keyFileProd $keyFileStage $keyFile)

uri="$user@$host.eu-west-1.compute.amazonaws.com"

checkKey() {
    try=$(ssh -i $1 $uri echo 200 2>&1)
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

failMessages=()
for key in "${keys[@]}"; do
    r=$(checkKey "$key")
    if [[ $r = 'ok' ]]; then
        ssh -i $key $uri
        exit
    else
        failMessages+=("$key -> $r")
    fi
done

for m in "${failMessages[@]}"; do
    echo "$m"
done
