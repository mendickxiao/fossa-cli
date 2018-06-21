#!/usr/bin/env bash

set -e

# Test known good Go projects:
## FOSSA CLI (dep)
cd $GOPATH/src/github.com/fossas/fossa-cli
fossa init
time fossa analyze --output go:./cmd/fossa

## Kubernetes (godep)
cd $GOPATH/src/k8s.io/kubernetes
fossa init
time fossa analyze --output --option allow-unresolved-prefix:k8s.io go:./cmd/kube-apiserver

## Consul (govendor)
cd $GOPATH/src/github.com/hashicorp/consul
fossa init
time fossa analyze --output --option allow-nested-vendor:true --option allow-deep-vendor:true go:.

## Docker (vndr)
cd $GOPATH/src/github.com/docker/docker
fossa init
time fossa analyze --output --option allow-unresolved-prefix:"github.com/docker archive/tar" go:./cmd/dockerd

## Docker CE (vndr)
cd $GOPATH/src/github.com/docker/docker-ce
fossa init
time fossa analyze --output --option allow-unresolved-prefix:"github.com/docker archive/tar" --option allow-external-vendor-prefix:github.com/docker/docker go:./components/engine/cmd/dockerd

## InfluxDB (gdm)
cd $GOPATH/src/github.com/influxdata/influxdb
fossa init
time fossa analyze --output --option allow-unresolved-prefix:github.com/influxdata go:./cmd/influxd

## rkt (glide)
cd $GOPATH/src/github.com/rkt/rkt
fossa init
time fossa analyze --output go:./rkt

# Test known good NodeJS projects:
cd $HOME/puppeteer
fossa init
time fossa analyze --output

cd $HOME/pkg
fossa init
time fossa analyze --output

cd $HOME/faker.js
fossa init
time fossa analyze --output

cd $HOME/fastify
fossa init
time fossa analyze --output

cd $HOME/nest
fossa init
time fossa analyze --output

cd $HOME/ohm
fossa init
time fossa analyze --output

cd $HOME/express
fossa init
time fossa analyze --output

## standard has missing peer dependencies.
cd $HOME/standard
fossa init
time fossa analyze --output --option allow-npm-err:true

cd $HOME/sodium-encryption
fossa init
time fossa analyze --output

cd $HOME/request
fossa init
time fossa analyze --output
