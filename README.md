# Case Study

## Table of Contents

* [Task Overview](#task-overview)
* [System Design](#system-design)
* [APP Config](#app-config)
* [Quickstart](#quickstart)
* [Optimization](#optimization)
* [API Document](#api-document)
* [Tech Stack](#tech-stack)
* [Thanks](#thanks)

## Task Overview

You are tasked with developing two applications: a job and a microservice. The job's responsibility is to fetch several
JSON
Lines (jsonl) files from AWS S3, process them, and write the data to a database. On the other hand, the microservice
should
provide one endpoint which gets a record by id.

### Job

You should create a pipeline to process files concurrently with the following steps:

1. Fetch jsonl files from the AWS S3.
2. Read files line by line and convert each line (which is a JSON object) to a Go struct.
3. Write the converted data to a database.
   Note that the files in AWS may be larger and the exact record(s) may already exist in the database. Therefore, you
   should
   consider potential performance issues.
   HINT: If the job runs a second time with the same product files or a new file contains nearly identical records to
   those already in
   the database, it could be a waste of resources.

### Microservice

You should create a RESTful microservice by implementing only one GET endpoint to fetch one row by id from the database.
But you should design the microservice with scalability in mind.

## System Design

![Figure: System Design](./docs/images/system-design.png)

Actually, I thought of reading the data I read from the file in the 1000s lines and writing it to a lookup table.
In this way, I could easily catch things like how much of the file I read and whether I got an error.
I also aimed to facilitate things like searching on data using postgres. My strategy here was to guarantee the read part
by doing bulk insert. In this process, I came up with this solution in order not to be unsure and over-engineering.

## APP Config

You can change any value from the following files.

* `./case-study-api-service/config.yml`
* `./case-study-job-service/config.yml`

## Quickstart

Just run the command below to start the application.

```bash
docker-compose up -d
```

Once the project is up and running, we can make a province request. We must remember that there is an authentication
step here.

> username -> Y2ltcmktdXNlcm5hbWU=
>
> password -> c2NSRlRXNVJvVEFVOENJTHE2VWtrUQ==

```bash
$ curl --location 'http://localhost:8080/healthcheck' \
--header 'Authorization: Basic WTJsdGNta3RkWE5sY201aGJXVT06YzJOU1JsUlhOVkp2VkVGVk9FTkpUSEUyVld0clVRPT0='
```

If you want to clean or see the number of keys in redis, you can follow the steps below.

```bash
$ docker exec -it redis redis-cli

127.0.0.1:6379> INFO Keyspace | grep ^db  # keys info
127.0.0.1:6379> flushall                  # all key cleanup
```

## Optimization

### Worker Optimize

| * | worker | latency       | keys   |
|---|--------|---------------|--------|
|   | 3      | 10.029652615s | 169999 |
|   | 5      | 10.542316375s | 169999 |
|   | 7      | 11.085843981s | 169999 |
|   | 9      | 10.094823171s | 169999 |

### Test Coverage

To run tests and generate mocks

```bash
cd case-study-api-service
go generate ./...
go test ./...
```

```bash
cd case-study-job-service
go generate ./...
go test ./...
```

If you want to run integration tests

```bash
cd case-study-api-service
go test -tags=integration ./...
```

```bash
cd case-study-job-service
go test -tags=integration ./...
```

| * | service name | coverage |
|---|--------------|----------|
|   | api-service  | 78.8%    |
|   | job-service  | 78.8%    |

## API Document

Postman collection path is available below

* `./docs/postman/case-study.postman_collection.json`

## Tech Stack

* Infrastructure
  * aws s3
  * redis
  * docker
* Application
  * go 1.22
  * zap
  * testcontainers-go
  * echo
  * viper
  * testify
  * gomock

## Thanks

Finally, while doing this case, I had the chance to research many things and improve myself. Thank you for this chance.
