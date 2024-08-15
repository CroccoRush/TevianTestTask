# Tevian Test Task

## Overview
This service provides a REST API for managing tasks that involve image processing to detect faces, determine gender and age, and generate statistics based on the processed data. The service interacts with the FaceСloud API for face detection.

## Features
- Create, and delete tasks.
- Add images to tasks.
- Start asynchronous processing of tasks.
- Retrieve task status and statistics.

[//]: # (- Secure API access with HTTP Basic Auth.)

## Requirements
- Go 1.22 or higher
- PostgreSQL 16.4 or higher
- FaceCloud API credentials

or

- Docker
- FaceCloud API credentials

## Installation

### Prerequisites
- Ensure Go is installed on your system. You can download it from [golang.org](https://golang.org/).
- Install PostgreSQL 16.4. You can download it from the [PostgreSQL official site](https://www.postgresql.org/download/).
- Obtain FaceCloud API credentials by following the instructions in the [FaceCloud documentation](https://docs.facecloud.tevian.ru).

### Run source code
1. Set up environment variables:
   ```sh
   export FACECLOUD_EMAIL=your_facecloud_email
   export FACECLOUD_PASSWORD=your_facecloud_password
   export FACECLOUD_API_URL=https://backend.facecloud.tevian.ru/api

   export POSTGRES_HOST=postgres
   export POSTGRES_PORT=5432
   export POSTGRES_USER=postgres
   export POSTGRES_PASSWORD=password
   export POSTGRES_DB=database
   ```

2. Run the docker container:
   ```sh
   docker run --name db_postgres_container -p 5432:5432 -e POSTGRES_PASSWORD=<POSTGRES_PASSWORD> -e POSTGRES_USER=<POSTGRES_USER> -e POSTGRES_DB=<POSTGRES_DB> -d postgres:16.4-alpine3.20
   ```
   
3. Build the service:
   ```sh
   go build -o tevian_test_task
   ```

4. Run the service:
   ```sh
   ./tevian_test_task
   ```

### Run docker-compose.yml
1. Run from the root of the project:
   ```sh
   docker-compose -f ./build/package/docker-compose.yml up
   ```



## API Documentation
The API documentation is available in Swagger format. Once the service is running, you can access the Swagger UI by navigating to:

```
http://localhost:3000/swagger
```

