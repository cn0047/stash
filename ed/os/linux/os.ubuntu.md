ubuntu
-

````sh
apt-get update --fix-missing
apt-get install -y --force-yes

apt list vim
apt list --installed
# search package
apt-cache search htop

dpkg -l     # list installed debs
dpkg -L vim # list
dpkg -s vim # status

apt-get update
apt-get upgrade
apt-get full-upgrade
apt-get install -y software-properties-common build-essential
add-apt-repository "deb http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1804/x86_64/ /"

apt-get autoremove
apt-get autoclean
apt-get clean
apt-get clean packages
````
