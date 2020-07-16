# Post_svc Service

This is the Post_svc service

Generated with

```
micro new post_svc --namespace=go.micro --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.post_svc
- Type: service
- Alias: post_svc

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./post_svc-service
```

Build a docker image
```
make docker
```