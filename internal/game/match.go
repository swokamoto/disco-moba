package game

import "time"

type MatchID string
type TeamID string

// WinCondition describes how a match ended.
type WinCondition string

const (
	WinTowersDestroyed WinCondition = "towers_destroyed" // all 3 opponent towers destroyed
	WinMoraleZero      WinCondition = "morale_zero"      // opponent morale reached 0
	WinCreepOverrun    WinCondition = "creep_overrun"    // opponent base overrun by creep
)

// MatchStatus is the lifecycle state of a match.
type MatchStatus string

const (
	MatchPending  MatchStatus = "pending"  // waiting for players to join
	MatchActive   MatchStatus = "active"   // waves in progress
	MatchFinished MatchStatus = "finished" // a win condition was met
)

// Tower represents one of a team's three towers.
type Tower struct {
	Index     int // 0, 1, 2
	Health    int // reaches 0 when destroyed
	Destroyed bool
}

// TowerMaxHealth is the starting health of each tower.
const TowerMaxHealth = 100

// Team represents one side in a match.
// LaneCreep[i] is this team's creep count pushing into the opponent's lane i.
// Indexed 0=Left, 1=Mid, 2=Right (matching AllLanes order).
type Team struct {
	ID        TeamID
	Players   []Player
	Towers    [3]Tower // indexed 0=Left, 1=Mid, 2=Right
	LaneCreep [3]int   // this team's outgoing creep per lane
	Morale    int      // 0–100; reaching 0 triggers WinMoraleZero
	ServerID  string   // Discord guild ID
	ThreadID  string   // Discord thread ID in that guild
}

// Match is the top-level game state.
type Match struct {
	ID           MatchID
	Teams        [2]Team
	CurrentWave  int
	Status       MatchStatus
	CreatedAt    time.Time
	Winner       TeamID       // empty until match is finished
	WinCondition WinCondition // empty until match is finished
	FastMode     bool         // if true, use fast timing and rules
}
