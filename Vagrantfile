# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "centos/7"
  config.vm.box_check_update = true
  config.vm.network "private_network", type: "dhcp"

  # This should work
  #(1..2).each do |i|
  #  config.vm.define "web#{i}" do |web|
  #    web.vm.provision "ansible" do |ansible|
  #      ansible.playbook = "provisioning/site.yml"
  #      ansible.verbose = "vvv"
  #      ansible.sudo = true
  #    end
  #  end
  #end

  config.vm.define "web1" do |web1|
    web1.vm.hostname = "web1"
    web1.vm.provision "ansible" do |ansible|
      ansible.playbook = "provisioning/site.yml"
      ansible.verbose = "vvv"
      ansible.sudo = true
    end
  end

  config.vm.define "web2" do |web2|
    web2.vm.hostname = "web2"
    web2.vm.provision "ansible" do |ansible|
      ansible.playbook = "provisioning/site.yml"
      ansible.verbose = "vvv"
      ansible.sudo = true
    end
  end

  # Nginx proxy
  config.vm.define "proxy" do |proxy|
    #proxy.vm.network :forwarded_port, guest: 80, host: 8888
    proxy.vm.hostname = "proxy"
    proxy.vm.provision "ansible" do |ansible|
      ansible.groups = { 'webservers' => ["web[1:2]"] }
      ansible.playbook = "provisioning/site.yml"
      ansible.verbose = "v"
      ansible.sudo = true
    end
  end

end
