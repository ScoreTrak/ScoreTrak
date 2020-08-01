![Alt text](./images/logo.svg "Logo")

<h1 align="center">A distributed scoring engine</h1>

## Key Features

* Fault Tolerant
* Written in Golang
* Supports Integrations with platforms like docker, docker swarm, and kubernetes(in dev).
    * Automatically Spawns/Removes Workers.
* Supports running checks from workers and on masters directly
* And More!
# Getting Started

## 1. Clone the repo and install the dependencies
```bash
# Clone ScoreTrak repository
$ git clone https://github.com/L1ghtman2k/ScoreTrak

$ cd ScoreTrak

# Install dependencies
$ go mod tidy
```


## 2. Setup the database
Currently, ScoreTrak has integrations with Cockroach DB.

To install cockroachdb go to https://www.cockroachlabs.com/docs/stable/install-cockroachdb, and select the installation that suites your operating system.

Once installed, run:
```bash
$ cockroach start-single-node --insecure
```
This will launch an insecure cockroach cluster that scoretrak can use to store its Data.

Create scoretrak database:
```bash
$ cockroach sql --insecure --host=localhost:26257
> CREATE DATABASE scoretrak;
```

## 3. Setup the config
Scoretrak supports two types of configs: **Static Config** and **Dynamic Config**.
The key difference is that configurations related to **Static Config** can be modified before scoretrak starts, while **Dynamic Config** can be modified at runtime.
Copy 
```bash
# Make a copy of a default-config.yml
$ cp ./configs/default-config.yml ./configs/config.yml
```
Note: 

## 4. Start the application
```bash
$ go run ./cmd/master/main.go 
```

## 5. Get to know ScoreTrak

Check out `./docs` directory for more docs!