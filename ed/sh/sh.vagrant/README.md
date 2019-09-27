Vagrant
-

1.8.1
1.7.4

[download vagrant](https://www.vagrantup.com/downloads.html).

````sh
vboxmanage --version
````

````sh
vagrant plugin list

vagrant plugin install vagrant-vbguest
vagrant plugin install vagrant-cachier
vagrant plugin install vagrant-hosts
````

````sh
vagrant global-status

vagrant up
vagrant up --provider virtualbox
vagrant provision

vagrant rsync # forces a re-sync of any rsync synced folders

vagrant ssh
vagrant ssh-config

vagrant suspend # stop machine
vagrant halt    # power down the guest machine
vagrant reload  # The equivalent of running a halt followed by an up
vagrant destroy # power down, and remove all of the guest hard disks

vagrant resume # after suspend
````

guest machine - virtual box.
host machine - your laptop.

https://www.vagrantup.com/docs/multi-machine/
