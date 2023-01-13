Security tools
-

````sh
msfconsole
tcpdump -i eth0 -n udp port 53

ss # tool for sockets

nc -zv 10.0.2.2 22
nc -zv 78.140.186.238 1-65535
nc -zv 10.0.2.2 22

# public DNS IP addresses:
8.8.8.8
8.8.4.4

# osx proxy
sudo networksetup -setwebproxy "Wi-Fi" 54.174.16.166 80
# '71.4.186.100:21288', '198.7.96.243:21239', '104.128.17.224:21237', '209.48.175.196:21325', '172.102.207.55:21279', '198.7.97.209:21297', '162.248.134.54:21326', '23.239.219.244:21325', '143.191.19.7:21238', '173.44.12.201:21259', '173.44.5.52:21310', '173.44.12.68:21305', '173.44.26.80:21269', '107.152.144.66:21291', '208.73.73.206:21257', '204.14.87.85:21326', '144.168.128.88:21244', '204.14.87.149:21271', '45.57.195.33:21232', '173.44.5.185:21247', '173.44.12.141:21280', '173.44.26.220:21318', '107.152.144.219:21274', '208.73.73.9:21278', '143.191.19.89:21263', '143.191.31.123:21304', '69.174.99.149:21322', '50.117.15.33:21318', '173.44.12.16:21297', '216.172.144.220:21285',
sudo networksetup -setwebproxystate "Wi-Fi" off

ping 8.8.8.8 -c 15
traceroute http://cn007b.tumblr.com # print the route packets take to network host
nslookup git.mdm.comodo.od.ua
host github.com # shows ip by host.
host -t A github.com
dig domain
whois ip
nmap --script=http-headers www.zii.com # scan ports

````
