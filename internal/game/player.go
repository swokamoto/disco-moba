package game

// PlayerID is a Discord user ID.
type PlayerID string

// Player represents one participant in a match.
// Players have no HP – only towers take damage.
//
// Skill selection: At match start, for each player's position, a list of skill IDs is shuffled.
// When a skill offer is needed, the last 2 IDs are presented as options, then popped off after selection.
type Player struct {
	ID           PlayerID
	Name         string
	TeamID       TeamID
	Position     Position
	Lane         Lane
	PreviousLane Lane      // for movement-based skill tracking
	Stats        Stats     // derived from Position at match start
	DebuffLevel  int       // cumulative debuff; reduces effectiveness, can be healed by support
	SkillIDs     []SkillID // skills picked so far (max 4: 3 regular + 1 special)
	// At match start, RemainingSkillIDs is initialized to [0,1,2,3,4,5,6] (as SkillIDs), then shuffled.
	RemainingSkillIDs []SkillID // shuffled, pop last 2 for each offer
}
