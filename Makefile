
.EXPORT_ALL_VARIABLES:

SPEEDTEST_DB_FILE="./speedtest.db"

clean:
	-rm -rf bin

bin/speedtest-wrapper-go:
	go build -v -o bin/speedtest-wrapper-go -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/speedtest-wrapper-go/config.BuildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/speedtest-wrapper-go/config.Version=$(git rev-parse --short HEAD)'" .

bin/speedtest-wrapper-go.linux.armv7:
	GOOS=linux \
	GOARCH=arm \
	GOARM=7 \
	go build -v -o bin/speedtest-wrapper-go.linux.armv7 -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/speedtest-wrapper-go/config.BuildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/speedtest-wrapper-go/config.Version=$(git rev-parse --short HEAD)'" .

bin/speedtest-wrapper-go.linux.amd64:
	GOOS=linux \
	GOARCH=amd64 \
	go build -v -o bin/speedtest-wrapper-go.linux.amd64 -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/speedtest-wrapper-go/config.BuildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/speedtest-wrapper-go/config.Version=$(git rev-parse --short HEAD)'" .

bin/speedtest-wrapper-go.linux.arm64:
	GOOS=linux \
	GOARCH=arm64 \
	GOARM=7 \
	go build -v -o bin/speedtest-wrapper-go.linux.arm64 -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/speedtest-wrapper-go/config.BuildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/speedtest-wrapper-go/config.Version=$(git rev-parse --short HEAD)'" .

buildall: clean bin/speedtest-wrapper-go.linux.amd64 bin/speedtest-wrapper-go.linux.armv7 bin/speedtest-wrapper-go.linux.arm64
	@echo "Build with success"

build: bin/speedtest-wrapper-go
	@echo "Build with success"

start:
	SPEEDTEST_DB_FILE=$(SPEEDTEST_DB_FILE) go run main.go test -p --config config/samples/config.yml

vagrantclean:
	-cd install_sample ; vagrant destroy -f
	-rm -rf install_sample/.tmp

vagrantssh:
	-cd install_sample ; vagrant ssh

vagrantsetup: vagrantclean build
	-mkdir -p install_sample/.tmp
	cp bin/speedtest-wrapper-go install_sample/.tmp/speedtest-wrapper-go
	cp config/samples/config.yml install_sample/.tmp/config.yml
	cp install_sample/scripts/*.sh install_sample/.tmp/
	cd install_sample ; vagrant up

vagrantlibvirt: vagrantclean build
	-mkdir -p install_sample/.tmp
	cp bin/speedtest-wrapper-go install_sample/.tmp/speedtest-wrapper-go
	cp config/samples/config.yml install_sample/.tmp/config.yml
	cp install_sample/scripts/*.sh install_sample/.tmp/
	cd install_sample ; vagrant up --provider=libvirt

molecule: build
	cd install_sample/ansible/roles/setup_speedtest_wrapper ; pwd ; molecule test
