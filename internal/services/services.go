package services

import (
	"blackboxv3/internal/models"
	"blackboxv3/pkg/db"
	"blackboxv3/pkg/utils"
	"os"
	"strings"

	"log"
)

func Signup(username, pin, email string) error {
	err := db.DB.Create(&models.User{
		Username: username,
		PIN:      pin,
		Email:    email,
	}).Error
	return err
}

func Login(username, pin string) (string, error) {
	var user models.User
	err := db.DB.Where("username = ? AND pin = ?", username, pin).First(&user).Error
	if err != nil {
		return "", err
	}
	token, err := utils.CreateToken(username, user.ID)
	return token, err
}

func CreateChannel(name string) error {
	var err error
	uploadsDir := os.Getenv("UPLOADSDIR")
	log.Println(uploadsDir)
	log.Println(name)
	err = os.MkdirAll(uploadsDir+"/"+name, 0777)
	if err != nil {
		return err
	}
	err = db.DB.Create(&models.Channel{
		Name: name,
	}).Error
	if err != nil {
		err = os.Remove(uploadsDir + "/" + name)
	}
	return err
}

func GetChannels() ([]models.Channel, error) {
	var channels []models.Channel
	err := db.DB.Find(&channels).Error
	return channels, err
}

func GetChannelMessages(channelID uint) ([]models.Message, error) {
	var textMsgs []models.TextMessage
	var mediaMsgs []models.MediaMessage
	var err error
	err = db.DB.Where("channel_id = ?", channelID).Find(&textMsgs).Error
	if err != nil {
		return nil, err
	}
	err = db.DB.Where("channel_id = ?", channelID).Find(&mediaMsgs).Error
	if err != nil {
		return nil, err
	}
	var messages []models.Message
	for _, msg := range textMsgs {
		messages = append(messages, models.Message(msg))
	}
	for _, msg := range mediaMsgs {
		messages = append(messages, models.Message(msg))
	}
	return messages, nil
}

func SendMessage(channelID uint, text string) error {
	err := db.DB.Create(&models.TextMessage{
		ChannelID: channelID,
		Text:      text,
	}).Error
	return err
}

func SaveMediaMessageMetadata(channelID uint, fileName string) error {
	var err error
	fileStringSplit := strings.Split(fileName, ".")
	fileType := fileStringSplit[len(fileStringSplit)-1]

	err = db.DB.Create(&models.MediaMessage{
		ChannelID: channelID,
		MediaName: fileName,
		MediaType: fileType,
	}).Error
	return err
}

func GetChannelByParam(param, value string) (models.Channel, error) {
	var channel models.Channel
	err := db.DB.Where(param+" = ?", value).First(&channel).Error
	return channel, err
}
