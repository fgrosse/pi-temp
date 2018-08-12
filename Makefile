.PHONY: release

VERSION ?= $(shell git describe --dirty)

release:
	test -n "$(VERSION)"
	mkdir -p releases

	# Prepare template directory
	mkdir -p release-$(VERSION)
	cp LICENSE release-$(VERSION)
	cp README.md release-$(VERSION)

	# Linux 64
	GOOS=linux GOARCH=amd64 go build -o release-$(VERSION)/pi-temp
	tar -czf pi-temp-$(VERSION).linux-amd64.tar.gz -C release-$(VERSION) .
	mv pi-temp-$(VERSION).*.tar.gz releases

	# Linux arm
	GOOS=linux GOARCH=arm go build -o release-$(VERSION)/pi-temp
	tar -czf pi-temp-$(VERSION).linux-arm.tar.gz -C release-$(VERSION) .
	mv pi-temp-$(VERSION).*.tar.gz releases

	# Linux arm64
	GOOS=linux GOARCH=arm64 go build -o release-$(VERSION)/pi-temp
	tar -czf pi-temp-$(VERSION).linux-arm64.tar.gz -C release-$(VERSION) .
	mv pi-temp-$(VERSION).*.tar.gz releases

	rm -R release-$(VERSION)
