package monketype

import (
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"

	"encoding/json"
)

type Content struct {
	ID           string   `json:"id"               db:"id"`
	FileURL      string   `json:"file_url"         db:"file_url"`
	Author       string   `json:"author"           db:"author"`
	Mime         string   `json:"mime"             db:"mime"`
	Tags         []string `json:"tags"             db:"tags"`
	LikeCount    int      `json:"like_count"       db:"like_count"`
	DislikeCount int      `json:"dislike_count"    db:"dislike_count"`
	RepubCount   int      `json:"repub_count"      db:"repub_count"`
	ViewCount    int      `json:"view_count"       db:"view_count"`
	CommentCount int      `json:"comment_count"    db:"comment_count"`
	Created      int64    `json:"created"          db:"created"`
	Featured     bool     `json:"featured"         db:"featured"`
	Featurable   bool     `json:"featurable"       db:"featurable"`
	Removed      bool     `json:"removed"          db:"removed"`
	NSFW         bool     `json:"nsfw"             db:"nsfw"`
}

func (_ Content) FromMap(data map[string]interface{}) (it MonkeType, err error) {
	var config mapstructure.DecoderConfig = mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  "json",
		Result:   &it,
	}

	var decoder *mapstructure.Decoder
	if decoder, err = mapstructure.NewDecoder(&config); err == nil {
		err = decoder.Decode(data)
	}

	return
}

func (content Content) Map() (data map[string]interface{}, err error) {
	var bytes []byte
	if bytes, err = json.Marshal(content); err == nil {
		err = json.Unmarshal(bytes, &data)
	}

	return
}

func (content Content) JSON() (data []byte, err error) {
	data, err = json.Marshal(content)
	return
}

func NewContent(file_url, author, mime string, tags []string, featurable, nsfw bool) (created Content) {
	created = Content{
		FileURL:    file_url,
		Author:     author,
		Mime:       mime,
		Tags:       tags,
		Featurable: featurable,
		NSFW:       nsfw,

		Created: time.Now().Unix(),
		ID:      uuid.New().String(),
	}

	return
}
