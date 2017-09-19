H
-

````sh
# osx proxy

sudo networksetup -setwebproxy "Wi-Fi" 54.174.16.166 80
sudo networksetup -setwebproxystate "Wi-Fi" off

netstat -an

nc -zv 10.0.2.2 22

nc -zv 78.140.186.238 1-65535
nc -zv 10.0.2.2 22
````
