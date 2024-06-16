# Upcoming Changes

This file lists upcoming changes and improvements for the CLI Application that generates a `Golang REST API service` connected to a database in a Docker container. The generated API service uses the `Gin-gonic` framework, `Viper` library, `SQLC` for query generation, and `Golang-migrate` for database migration.

## Planned Changes
- [ ] **Add MongoDB Support**
  - Description: Extend the CLI to support MongoDB databases.
  - Purpose: Provide more database options for the generated API service.

- [ ] **Restructure to Optimized Plugin Architecture**
  - Description: Refactor the CLI application to use a plugin architecture.
  - Impact: Improve maintainability and extensibility of the CLI tool.

- [ ] **Add Dockerfile and Manifests**
  - Description: Include Dockerfile and Kubernetes manifests (deployments, services, secrets, etc.) for the generated API service.
  - Purpose: Simplify deployment and integration into Kubernetes environments.

## In Progress
- [ ] **Feature: MongoDB Support**
  - Description: Development work to integrate MongoDB as a supported database.
  - Status: Researching and prototyping integration methods.

- [ ] **Restructure to Optimized Plugin Architecture**
  - Description: Refactor the CLI application to use a plugin architecture.
  - Impact: Improve maintainability and extensibility of the CLI tool.

## Proposed Changes
- [ ] **Add Jenkinsfile**
  - Description: Add a Jenkinsfile for continuous integration and continuous delivery (CI/CD) pipelines.
  - Impact: Automate the build, test, and deployment processes.

- [ ] **Multi-Language Support (Node.js)**
  - Description: Extend the CLI to generate API services in Node.js.
  - Purpose: Provide support for generating API services in multiple programming languages.

- [ ] **Proposal: GraphQL Support**
  - Description: Add an option to generate API services with GraphQL endpoints.
  - Consideration: Assess the demand and feasibility of implementing GraphQL.

## Completed Changes
- [x] **Initial Release**
  - Description: Initial version of the CLI application with support for PostgreSQL database.
  - Status: Released and available for use.

- [x] **v2.0.0**
  - Description: Renamed the subcommand from `template` to `go-template`
  - Status: Released and available for use.

- [x] **v2.1.0**
  - Description: Added MySQL database support to the CLI application.
  - Status: Released and available for use.
