#!/bin/sh

source /etc/bashrc

l=$(sudo /usr/local/bin/supervisorctl status | awk '{print $1}')
for p in $l; do
    v=$(sudo /usr/local/bin/supervisorctl status $p | grep RUNNING -c )
    `aws cloudwatch put-metric-data --namespace $APP_ENVIRONMENT'.supervisor' --metric-name $HOSTNAME'.'$p --value $v`
done
