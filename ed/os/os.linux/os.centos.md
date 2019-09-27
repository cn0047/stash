centos
-

````sh
cat /etc/system-release

yum --showduplicates list available elasticsearch\*

sudo yum check-update

rpm -qa # list installed rpms
yum list installed

# like sudo service sshd status in ubuntu
chkconfig --list sshd
chkconfig sshd on
chkconfig sshd off
````
