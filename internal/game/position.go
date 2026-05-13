package game

// Position represents one of the five roles, like Dota positions.
// No two players on the same team may share a position.
type Position string

const (
	PositionCarry       Position = "carry"        // pos 1 – high damage, scales late
	PositionMid         Position = "mid"          // pos 2 – versatile, solo lane
	PositionOfflane     Position = "offlane"      // pos 3 – durable initiator
	PositionSoftSupport Position = "soft_support" // pos 4 – utility / roamer
	PositionHardSupport Position = "hard_support" // pos 5 – protector / healer
)

// AllPositions lists every valid position.
var AllPositions = []Position{
	PositionCarry,
	PositionMid,
	PositionOfflane,
	PositionSoftSupport,
	PositionHardSupport,
}

// Stats are the base attributes for a position.
//
// Damage chain per wave:
//   player Attack → opponent's LaneCreep in that lane → opponent's Tower health → opponent's Morale
//
// CreepPush adds to your team's LaneCreep each wave you are present in a lane,
// increasing pressure on the opponent's tower even without a direct attack action.
type Stats struct {
	Attack    int // damage dealt per wave, hits opponent LaneCreep first, then tower, then morale
	Mobility  int // max lanes a player can move per wave (linear map, max 2)
	CreepPush int // creep units added to your team's lane push each wave you are present
}

// BaseStats returns the starting stats for a given position.
// Pure function – no side effects.
func BaseStats(p Position) Stats {
	switch p {
	case PositionCarry:
		return Stats{Attack: 8, Mobility: 1, CreepPush: 2}
	case PositionMid:
		return Stats{Attack: 6, Mobility: 2, CreepPush: 3}
	case PositionOfflane:
		return Stats{Attack: 5, Mobility: 1, CreepPush: 4}
	case PositionSoftSupport:
		return Stats{Attack: 3, Mobility: 2, CreepPush: 5}
	case PositionHardSupport:
		return Stats{Attack: 2, Mobility: 1, CreepPush: 7}
	default:
		return Stats{}
	}
}

// ValidPosition returns true if p is a known position.
func ValidPosition(p Position) bool {
	for _, pos := range AllPositions {
		if p == pos {
			return true
		}
	}
	return false
}
