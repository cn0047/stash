Vagrant
-

1.8.1
1.7.4

[download virtualbox](https://www.virtualbox.org/wiki/Linux_Downloads),
[download vagrant](https://www.vagrantup.com/downloads.html).

````
vboxmanage --version
````

````
vagrant plugin list

vagrant plugin install vagrant-vbguest
vagrant plugin install vagrant-cachier
vagrant plugin install vagrant-hosts
````

````
vagrant global-status

vagrant up
vagrant reload
vagrant provision

vagrant rsync # forces a re-sync of any rsync synced folders

vagrant ssh
vagrant ssh-config

vagrant suspend # stop machine
vagrant halt    # power down the guest machine
vagrant destroy # power down, and remove all of the guest hard disks

vagrant resume # after suspend
````

guest machine - virtual box.
host machine - your laptop.

https://www.vagrantup.com/docs/multi-machine/
