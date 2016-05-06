
PROG=helloworld
GOOS=linux
GOARCH=386

up: $(PROG)
	vagrant up && echo "Test url should be on port 80 on :- " && vagrant ssh proxy -c 'ip addr show dev eth1' | grep "inet "  | awk '{ print $2 }' | sed -e 's/\/[0-9]*//'

$(PROG): helloWorld.go
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o $@ $^


clean:
	vagrant destroy -f && rm -r $(PROG) .ansible_cache 
