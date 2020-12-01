# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"

  config.vm.provider "virtualbox" do |v|
    v.memory = 1024
    v.cpus = 1
    v.linked_clone = true
  end

  boxes = [
    { :name => "mongo1", :ip => "192.168.33.71"},
    { :name => "mongo2", :ip => "192.168.33.72"},
    { :name => "mongo3", :ip => "192.168.33.73"}
  ]

  boxes.each do |opts|
    config.vm.define opts[:name] do |config|
      config.vm.hostname = opts[:name]
      config.vm.network :private_network, ip: opts[:ip]

      if opts[:name] == "mongo1"

        config.vm.provision "ansible" do |ansible|
          ansible.playbook = "mongo.yaml"
          ansible.limit = "all"
          ansible.become = true
          ansible.groups = {
            "mongo" => ["mongo1", "mongo2", "mongo3"]
          }
          ansible.host_vars = {
            "mongo1" => {"master" => true}
        }
      end
    end
  end
end
end






