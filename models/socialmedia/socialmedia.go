package socialmedia

import (
	"time"
)

type MoodState int

const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

type AuditableContent struct {
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
}

type Post struct {
	AuditableContent
	Caption      string    `json:"caption"`
	MessageBody  string    `json:"messageBody"`
	URL          string    `json:"url"`
	ImageURI     string    `json:"imageURI"`
	ThumbnailURI string    `json:"thumbnailURI"`
	Keywords     []string  `json:"keywords"`
	Likers       []string  `json:"Likers"`
	AuthorMood   MoodState `json:"authorMood"`
	AuthorMoodEmoji string `json:"authorMoodEmoji"`
}

var Moods map[string]MoodState
var MoodsEmoji map[MoodState]string

func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}

	MoodsEmoji = map[MoodState]string{MoodStateNeutral: "\xF0\x9F\x98\x90", MoodStateHappy: "\xF0\x9F\x98\x8A", MoodStateSad: "\xF0\x9F\x98\x9E", MoodStateAngry: "\xF0\x9F\x98\xA0", MoodStateHopeful: "\xF0\x9F\x98\x8C", MoodStateThrilled: "\xF0\x9F\x98\x81", MoodStateBored: "\xF0\x9F\x98\x84", MoodStateShy: "\xF0\x9f\x98\x83", MoodStateComical: "\xF0\x9F\x98\x9C", MoodStateOnCloudNine: "\xF0\x9F\x98\x82"}
}

func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post{
	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent, AuthorMoodEmoji: MoodsEmoji[mood]}
}
