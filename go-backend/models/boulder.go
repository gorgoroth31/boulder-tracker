package models

type Boulder struct {
	screwedDifficulty  Difficulty
	feltLikeDifficulty Difficulty
	attempts           int
	sessionsTried      int
	exhausting         bool
	style              []Style
}

func NewBoulder(screwedDifficulty Difficulty,
	feltLikeDifficulty Difficulty,
	attempts int, sessionsTried int,
	exhausting bool, style []Style) *Boulder {
	return &Boulder{screwedDifficulty: screwedDifficulty,
		feltLikeDifficulty: feltLikeDifficulty,
		attempts:           attempts, sessionsTried: sessionsTried,
		exhausting: exhausting, style: style}
}
