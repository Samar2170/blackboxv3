package models

import (
	"blackboxv3/pkg/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message interface {
	toJson() map[string]interface{}
}

func init() {
	db.DB.AutoMigrate(&User{}, &Channel{}, &MediaMessage{}, &TextMessage{})
}

type User struct {
	*gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username"`
	PIN      string    `json:"pin"`
	Email    string    `json:"email"`
}

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

func (t *TextMessage) toJson() map[string]interface{} {
	return map[string]interface{}{
		"id":        t.ID,
		"chat_id":   t.ChannelID,
		"channel":   t.Channel,
		"text":      t.Text,
		"timestamp": t.CreatedAt,
	}
}

func (m *MediaMessage) toJson() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.ID,
		"chat_id":    m.ChannelID,
		"channel":    m.Channel,
		"text":       m.Text,
		"media_name": m.MediaName,
		"media_type": m.MediaType,
		"media_size": m.MediaSize,
		"media_url":  m.MediaURL,
		"timestamp":  m.CreatedAt,
	}
}
