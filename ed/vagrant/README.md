Vagrant
-

1.8.1
1.7.4

[download virtualbox](https://www.virtualbox.org/wiki/Linux_Downloads),
[download vagrant](https://www.vagrantup.com/downloads.html).

vagrant plugin install vagrant-vbguest
vagrant plugin install vagrant-cachier

````
vagrant global-status

vagrant up
vagrant reload
vagrant provision

vagrant ssh

vagrant suspend # stop machine
vagrant halt    # power down the guest machine
vagrant destroy # power down, and remove all of the guest hard disks

vagrant resume # after suspend
````

guest machine - virtual box.
host machine - your laptop.
