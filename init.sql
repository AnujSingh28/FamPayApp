-- Create database
CREATE DATABASE IF NOT EXISTS youtubestore;
GRANT ALL PRIVILEGES ON DATABASE youtubestore TO "postgres";

-- Connect to youtubestore database
\c youtubestore;

-- Create table
CREATE TABLE IF NOT EXISTS youtube_video (
    id UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    description VARCHAR(255),
    video_id VARCHAR(255),
    video_link VARCHAR(255),
    published_at TIMESTAMP,
    thumbnail VARCHAR(255)
    );
