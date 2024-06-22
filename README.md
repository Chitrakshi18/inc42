This project includes three main services: a Next.js application, a Go application, and a WordPress application. The project uses Docker Compose for local development and GitHub Actions for CI/CD workflows, which deploy the services to Amazon ECS.

## Table of Contents

Overview

Services

Local Development

CI/CD Workflows

ECS Deployment

Ports

## Overview

This project is designed to facilitate the development, deployment, and continuous integration of three distinct applications: a Next.js app, a Go app, and a WordPress site. Docker Compose is used to manage the services locally, while GitHub Actions handles the CI/CD pipelines, automating the building, pushing of Docker images, and deploying to Amazon ECS.

## Services

### Next.js Application

The Next.js application is a React-based framework used for server-side rendering and generating static websites. It is configured to run on port 3000. The CI/CD pipeline for this service builds the Docker image, pushes it to Docker Hub, and updates the ECS service.

### Go Application

The Go application is a backend service that runs on port 8080. It uses the Go programming language and is set up similarly to the Next.js application, with a CI/CD pipeline that handles building, pushing the Docker image, and deploying to ECS.

### WordPress Application

The WordPress application is a content management system that runs on port 8081. It includes a MariaDB database for storing content. The CI/CD pipeline for WordPress builds the Docker image, pushes it to Docker Hub, and deploys it to ECS.

## Local Development

For local development, Docker Compose is used to orchestrate the services. Each service runs in its own container, and the environment variables required for each service are specified in a .env file. Running docker-compose up --build will start all the services, making them accessible on their respective ports.

## CI/CD Workflows

GitHub Actions are used for continuous integration and continuous deployment:

**Next.js**: On pushing to the main branch or merging a pull request, the workflow builds the Docker image, pushes it to Docker Hub, and updates the ECS service.

**Go**: Similar to Next.js, the workflow triggers on changes to the main branch and handles building, pushing, and deploying the Docker image.

**WordPress**: The workflow for WordPress also builds, pushes the Docker image to Docker Hub, and updates the ECS service on changes to the main branch.

## ECS Deployment

Amazon ECS (Elastic Container Service) is used for deploying the Docker images. Each service has a corresponding ECS service and task definition that defines how the Docker containers should be run in the cluster. The GitHub Actions workflows update these services with new Docker images whenever changes are pushed to the main branch.

## Ports

Next.js Application: Port 3000

Go Application: Port 8080

WordPress Application: Port 8081

These ports are configured in the docker-compose.yml file and expose the services to be accessed locally during development. In production, these services are managed and scaled by ECS, ensuring high availability and reliability.
