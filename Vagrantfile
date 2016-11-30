Vagrant.configure(2) do |config|
  config.vm.define "ed" do |t|
  end
  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
    v.cpus = 2
  end
  config.vm.network "forwarded_port", guest: 80, host: 8080
  config.vm.network "forwarded_port", guest: 22, host: 1234
  config.vm.synced_folder "ed", "/var/www/html"

  config.vm.box = "ubuntu/trusty64"
  config.vm.provision :shell, path: "vagrant/ubuntu-14.04/provision.sh"

  #config.vm.box = "bento/ubuntu-16.04"
  #config.vm.provision :shell, path: "vagrant/ubuntu-16.04/provision.sh"

end
