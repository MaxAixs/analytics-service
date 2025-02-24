# Analytics Service

## Description
The Analytics Service is designed to collect and process data on user task completions. It accepts information about completed tasks, stores it in a PostgreSQL database, and generates weekly reports that aggregate user task completion statistics via a gRPC API.

## Features
- **Task Data Storage:** Receives and stores completed task data while tracking notification status.
- **Weekly Report Generation:** Aggregates tasks completed over the last 7 days for each user.
- **gRPC API:** Provides efficient and reliable communication between microservices.
- **Scalability:** Easily scalable and integrates seamlessly into a broader application ecosystem.

## Project Structure
- **analytics:** Contains data models (`TaskModel` and `CompletedTaskModel`) for tasks.
- **services:** Implements business logic for saving and retrieving task data, including model conversions between internal and gRPC formats.
- **repository:** Provides the data access layer for storing and fetching data from PostgreSQL.
- **handler:** Contains gRPC handlers that expose the API for saving task data and fetching weekly reports.
- **server:** Sets up and starts the gRPC server.
- **database:** Manages configuration and connection to the database.
- **proto:** Contains Protocol Buffer definitions and gRPC service contracts.
- **Docker Compose:** Includes configuration files for deploying the service, database, and migration tool.

## Requirements
- **Go**
- **Protocol Buffers:** `protoc` compiler installed (version 3.x or higher) for generating gRPC code.
- **PostgreSQL:** Can be deployed using Docker Compose.
- **Docker & Docker Compose:** For local deployment.

