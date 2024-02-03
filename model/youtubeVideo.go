package model

import (
	"github.com/google/uuid"
	"time"
)

type YoutubeVideo struct {
	ID        uuid.UUID `gorm:"type:uuid;size:36;primaryKey;index;not null;default:gen_random_uuid();" json:"id"`
	Title string `gorm:"type:varchar(100);size:100;not null;default:null;" json:"title"`
	Category string `gorm:"type:varchar(100);size:100;not null;default:null;" json:"category"`
	Description string `gorm:"type:varchar(100);size:100;not null;default:null;" json:"description"`
	VideoId string `gorm:"type:varchar(100);size:500;not null;default:null;" json:"video_id"`
	VideoLink string `gorm:"type:varchar(500);size:500;not null;default:null;" json:"video_link"`
	PublishedAt time.Time `gorm:"type:varchar(100);size:100;not null;default:null;" json:"published_at"`
	Thumbnail string `gorm:"type:varchar(200);size:200;not null;default:null;" json:"thumbnail"`
}

func (c YoutubeVideo) TableName() string {
	return "youtube_video"
}

