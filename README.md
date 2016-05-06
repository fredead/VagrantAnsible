# Vagrant / Ansible
This is an example of using vagrant with anisble to create 2 web servers and an nginx proxy to round robin between the 2 servers.
##### Software requirements
   - vagrant > 1.8.1
   - ansible > 2.0.2.0
   - go compiler
# How to run
```
make
```
# What happens
   - Compiles the go program for i386 linux 
   - Starts 2 web servers using ansible playlist
   - These servers copy the go binary and start it using systemd
   - Starts up a proxy server
   - Installs nginx from EPEL
   - Configures config based on information about the previous webservers
   - Does a quick test to see if it works


# Imporvements
   - The configuration is very dependant on the number of webservers partly due to the way vagrant and ansible play with each other it maybe possible to do this a better way.
   - The checks are very basic 
   - Security has not being a concerns so currently the app runs as root

# notes 
If you decide to try vagrant and ansible some quick notes of things that will help
- As Vagrant has a different private key per server it is importnat to keep a fact cache. see [ansible.cfg]
- Setting hostname in a Vagrant loop does not see possible (Bug?)

