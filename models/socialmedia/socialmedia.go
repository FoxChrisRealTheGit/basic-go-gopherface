package socialmedia

import(
	"time"
)

type MoodState int

const(
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
type AuditableContent struct{
	TimeCreated time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	CreatedBy string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}

type Post struct{
	AuditableContent
	Caption string `json:"caption"`
	MessageBody string `json:"messageBody"`
	URL string `json:"url"`
	ImageURI string `json:"imageURI"`
	ThumbnailURI string `json:"thumbnailURI"`
	Keywords []string `json:"keywords"`
	Likers []string `json:"Likers"`
	AuthorMood MoodState `json:"authorMood"`
}

var Moods map[string]MoodState

func init(){

}