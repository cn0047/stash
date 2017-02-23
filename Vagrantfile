Vagrant.configure(2) do |config|

  # config.vm.box = "ubuntu/trusty64"
  # config.vm.provision :shell, path: "vagrant/ubuntu-14.04/provision.sh"
  config.vm.box = "bento/ubuntu-16.04"
  config.vm.provision :shell, path: "vagrant/ubuntu-16.04/provision.sh"

  config.vm.synced_folder "ed", "/var/www/html"
  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
    v.cpus = 2
  end

  config.vm.define "ed", primary: true do |node|
    node.vm.hostname = 'ed'
    node.vm.network :private_network, ip: "192.168.56.101"
    node.vm.network :forwarded_port, guest: 22, host: 10122, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
  end

  config.vm.define "mysqlMaster1" do |node|
    node.vm.hostname = "mysqlMaster1"
    node.vm.network :private_network, ip: "192.168.56.102"
    node.vm.network :forwarded_port, guest: 22, host: 10123, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
  end

end
