package like

type LikeType int

const (
	Like          LikeType = 0
	Kinda_Like    LikeType = 1
	Kinda_Dislike LikeType = 2
	Dislike       LikeType = 3
)

func (l LikeType) String() string {
	likeTypes := [...]string{"Like", "Kinda_Like", "Kinda_Dislike", "Dislike"}

	if l < Like || l > Dislike {
		return "UNKNOWN"
	}

	return likeTypes[l]
}
