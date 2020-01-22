## Getting Started with icube-cybervision-plugin

### Requirements

* [Golang](https://golang.org/dl/) 1.9+


### Clone

To clone icube-cybervision-plugin from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/CiscoDevNet/icube-cybervision-plugin
git clone https://github.com/CiscoDevNet/icube-cybervision-plugin ${GOPATH}/src/github.com/CiscoDevNet/icube-cybervision-plugin
```

### Build

To build the binary for icube-cybervision-plugin run the command below. This will generate a binary
in the same directory with the name icube-cybervision-plugin.

```
make
```


### Run

To run icube-cybervision-plugin with debugging output enabled, run:

```
./icube-cybervision-plugin -c icube-cybervision-plugin.yml -e -d "*"
```


### build linux docker and run it with docker-compose


#### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```

#### Build docker image
```
make docker
```

#### Run docker image via docker-compose

To run icube-cybervision-plugin via docker-compose, please check out [CyberVision system status, and system analysis demo](./docker-compose/README.md).


## Beat plugin guide
For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Test

To test icube-cybervision-plugin, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  icube-cybervision-plugin source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

### Clone

To clone icube-cybervision-plugin from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/CiscoDevNet/icube-cybervision-plugin
git clone https://github.com/CiscoDevNet/icube-cybervision-plugin ${GOPATH}/src/github.com/CiscoDevNet/icube-cybervision-plugin
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.


