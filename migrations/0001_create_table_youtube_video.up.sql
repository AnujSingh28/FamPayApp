CREATE TABLE IF NOT EXISTS youtube_video(
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category VARCHAR(255),
    description VARCHAR(255),
    video_id VARCHAR(255),
    video_link VARCHAR(255),
    published_at time ,
    thumbnail VARCHAR(255),
);