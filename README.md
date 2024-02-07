# Backend Assignment | FamPay

## Project Overview
This project aims to create an API for fetching the latest videos from YouTube based on a predefined search query/tag. The fetched videos are stored in a database with relevant information such as video title, description, publishing datetime, thumbnail URLs, etc. Users can retrieve the stored video data, search for videos by title and description, and access the API endpoints in a paginated response.

## Features
- Continuously fetches the latest videos from YouTube using the YouTube API with a configurable interval.
- Stores video data in a database with proper indexing.
- Provides endpoints for retrieving stored video data, searching videos by title and description, and accessing paginated responses.
- Dockerized for easy deployment and scalability.

## Technologies Used
- Go (Golang) - Backend language
- Gin - HTTP web framework
- PostgreSQL - Database for storing video data
- Docker - Containerization for deployment
- YouTube Data API v3 - For fetching video data from YouTube

## Setup Instructions
1. Clone the repository:
   ```
   git clone https://github.com/AnujSingh28/FamPayApp.git
   ```

2. Build and run the Docker container:
   ```
   docker-compose up --build
   ```

3. Access the API endpoints:
   - The API endpoints will be available at `http://localhost:8080`.

## API Endpoints
1. A cron job has been set up to fetch YouTube videos without relying on an external API trigger. This job runs every minute, retrieving video details and storing them in the database. The decision to run it at one-minute intervals was made due to the limitations of the YouTube Data v3 API, which imposes a daily threshold.
   - Using my API key for the purpose.
   - Can be changed if needed. `./constants/common.go GCPApiKey`
3. GET `/allVideos`: Retrieve stored video data in a paginated response sorted by publishing datetime in descending order.
   ```
   curl --location 'http://localhost:8080/allVideos?page=1&recordsPerPage=8'
   ```
4. GET `/getVideo`: Search stored video by search slug.
   ```
   curl --location 'http://localhost:8080/getVideo?slug=India'
   ```
