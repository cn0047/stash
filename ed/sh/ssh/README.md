ssh (Secure Shell)
-

`pssh` tool to run ssh command on few servers simultaneously

````sh
# generate new ssh key
ssh-keygen -t rsa

# get key's fingerprint
ssh-keygen -E md5 -lf keyFile.pem

# get public key from private one
ssh-keygen -y -f ~/.ssh/id_rsa > ~/.ssh/id_rsa.pub

locate sshd_config

ssh-add ~/.ssh/id_rsa
vim /etc/ssh/sshd_config

sshfs -o nonempty -p22 root@host:/home/host/www /home/user/web/www
fusermount -u /home/user/web/www
ps aux | grep -i sftp | grep -v grep

# for AWS EC2
chmod 400 key.pem

# add public key to remote machine
echo 'ssh-rsa AAAAB3...3gRDw3sQ== name@mail.com' >> ~/.ssh/authorized_keys

# on host machine
cat ~/.ssh/id_rsa.pub
# on remote machine
echo 'key from id_rsa.pub from host machine' >> ~/.ssh/authorized_keys
````

````sh
ssh user@server.com
ssh -i $k $u@$h

ssh -i $key -N -L 9229:127.0.0.1:9229 root@server
# -L local_socket:remote_socket

# exec cmd through
ssh -i $k ubuntu@$h "echo 200 > /tmp/x"

scp -rp -i $key user@host:~/dir/ ~/dir/
````

`vim ~/.ssh/config`
````sh
Host ec2
    Hostname ec2-52-211-26-56.eu-west-1.compute.amazonaws.com
    User ec2-user
    IdentifyFile ~/path_to_ssh_key
````

Examples:

````
RSA/EC/DSA/OPENSSH

-----BEGIN RSA PRIVATE KEY-----
94bdTb18N6Zi9l23UXIQKpIh0pwqFYS7...
-----END RSA PRIVATE KEY-----
````
