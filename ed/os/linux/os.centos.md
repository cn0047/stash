centos
-

[packages](http://mirror.centos.org/centos/7/extras/x86_64/Packages/)

````sh
cat /etc/system-release
rpm -E %{rhel}          # get version

yum --showduplicates list available elasticsearch\*

sudo yum check-update

rpm -qa # list installed rpms
yum list installed

# like sudo service sshd status in ubuntu
chkconfig --list sshd
chkconfig sshd on
chkconfig sshd off
````
