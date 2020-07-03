ubuntu
-

````sh
apt-get install -y --force-yes
apt-get update --fix-missing

apt list vim
apt list --installed
# search package
apt-cache search htop

dpkg -l     # list installed debs
dpkg -L vim # list
dpkg -s vim # status

apt-get update
apt-get install -y software-properties-common
add-apt-repository "deb http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1804/x86_64/ /"
````
