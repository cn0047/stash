Vagrant.configure(2) do |config|

  # config.vm.box = "ubuntu/trusty64"
  config.vm.box = "bento/ubuntu-16.04"
  config.vm.provision :shell, path: "vagrant/provision.sh"
  config.vm.synced_folder "ed", "/var/www/html"
  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
    v.cpus = 2
  end

  config.vm.define "ed", primary: true do |node|
    node.vm.hostname = 'ed'
    node.vm.network :private_network, ip: "192.168.56.101"
    node.vm.network :forwarded_port, guest: 22, host: 10121, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
  end

  config.vm.define "mm" do |node|
    node.vm.hostname = "mm"
    node.vm.network :private_network, ip: "192.168.56.102"
    node.vm.network :forwarded_port, guest: 22, host: 10122, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
    node.vm.provider "virtualbox" do |v|
      v.memory = 256
      v.cpus = 1
    end
  end

  config.vm.define "ms1" do |node|
    node.vm.hostname = "ms1"
    node.vm.network :private_network, ip: "192.168.56.103"
    node.vm.network :forwarded_port, guest: 22, host: 10123, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
    node.vm.provider "virtualbox" do |v|
      v.memory = 256
      v.cpus = 1
    end
  end

  config.vm.define "ms2" do |node|
    node.vm.hostname = "ms2"
    node.vm.network :private_network, ip: "192.168.56.104"
    node.vm.network :forwarded_port, guest: 22, host: 10124, id: "ssh"
    node.vm.provision :hosts, :sync_hosts => true
    node.vm.provider "virtualbox" do |v|
      v.memory = 256
      v.cpus = 1
    end
  end

end
