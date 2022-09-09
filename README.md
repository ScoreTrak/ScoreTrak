![Alt text](./images/logo.svg "Logo")

<h1 align="center">Security Competition Scoring Engine</h1>

[![codecov](https://codecov.io/gh/ScoreTrak/ScoreTrak/branch/master/graph/badge.svg?token=WCHRWVG6B0)](https://codecov.io/gh/ScoreTrak/ScoreTrak)

## Key Features

* Fault Tolerant(Supports multiple scoring masters)
* Written in Golang
* Supports Integrations with platforms like docker, docker swarm, and kubernetes.
    * Automatically Spawns/Removes Workers.
* Supports running checks from workers and on masters directly
* Editing Service and Host configurations at runtime
* And More!
# Getting Started

## Clone the repo and install the dependencies
```bash
# Clone ScoreTrak repository
$ git clone https://github.com/ScoreTrak/ScoreTrak

$ cd ScoreTrak

# Install dependencies
$ go mod tidy
```

# Get to know ScoreTrak

Familiarize yourself with deployment options in the deployment [dir](./Deployment/README.md) and documentation [site](https://scoretrak.github.io).

