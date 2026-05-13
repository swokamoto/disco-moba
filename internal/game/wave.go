package game

// Action is a single player's decision for a wave.
// Only one action is allowed per player per wave.
//
// Since all skills are passive, the only action is movement (choosing a lane).
type Action struct {
	PlayerID   PlayerID
	TargetLane Lane // lane the player moves to for this wave
}

// Wave holds all submitted actions for a single day-wave.
type Wave struct {
	Number  int
	Actions map[PlayerID]Action
}

// IsSkillSelectWave returns true on waves where players choose a new skill (0, 2, 4, 6).
func IsSkillSelectWave(waveNumber int) bool {
	return waveNumber%2 == 0
}
