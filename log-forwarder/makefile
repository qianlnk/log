pack := log-forwarder
configs := $(subst .dist,,$(wildcard conf/*.*.dist))
remote := jp5

all:
	go build

package:
	GOOS=linux GOARCH=amd64 go build
	tar zcf $(pack)_`date +%Y-%m-%dT%H%M`.tar.gz $(pack)

clean:
	rm -f $(uf) $(pack) $(pack)_*.tar.gz

deploy: package
	scp $(pack)_*.tar.gz ${remote}:/data/go/log-forwarder/release
	@rm $(pack)_*.tar.gz
