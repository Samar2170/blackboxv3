package models

import "gorm.io/gorm"

type Channel struct {
	*gorm.Model
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type MediaMessage struct {
	*gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	ChannelID uint    `json:"chat_id"`
	Channel   Channel `json:"channel" gorm:"foreignKey:ChannelID"`
	Text      string  `json:"text"`
	MediaName string  `json:"media_name"`
	MediaType string  `json:"media_type"`
	MediaSize int64   `json:"media_size"`
	MediaURL  string  `json:"media_url"`
}

type TextMessage struct {
	*gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	ChannelID uint    `json:"chat_id"`
	Channel   Channel `json:"channel" gorm:"foreignKey:ChannelID"`
	Text      string  `json:"text"`
}
