# Standalone ScoreTrak on local machine


 ## 1. Setup the database
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
 
 ## 2. Setup the config
 Scoretrak supports two types of configs: **Static Config** and **Dynamic Config**.
 The key difference is that configurations related to **Static Config** can be modified before scoretrak starts, while **Dynamic Config** can be modified at runtime.
 Copy 
 ```bash
 # Make a copy of a default-config.yml
 $ cp ./configs/default-config.yml ./configs/config.yml
 ```
 Note: 
 
 ## 3. Start the application
 ```bash
 $ go run ./cmd/master/main.go 
 ```
 