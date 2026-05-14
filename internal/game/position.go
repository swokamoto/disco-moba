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
	PlayerAttack int // direct damage to players
	CreepAttack  int // damage to lane creeps
	Defense      int // damage absorbed per debuff level
	DebuffMax    int // max debuff before morale loss
	Mobility     int // max lanes a player can move per wave (linear map, max 2)
	CreepPush    int // creep units added to your team's lane push each wave you are present
}

// BaseStats returns the starting stats for a given position.
// Pure function – no side effects.
func BaseStats(p Position) Stats {
	switch p {
	case PositionCarry:
		return Stats{PlayerAttack: 6, CreepAttack: 2, Defense: 1, DebuffMax: 5, Mobility: 1, CreepPush: 0}
	case PositionMid:
		return Stats{PlayerAttack: 4, CreepAttack: 4, Defense: 1, DebuffMax: 5, Mobility: 1, CreepPush: 0}
	case PositionOfflane:
		return Stats{PlayerAttack: 4, CreepAttack: 4, Defense: 3, DebuffMax: 5, Mobility: 2, CreepPush: 0}
	case PositionSoftSupport:
		return Stats{PlayerAttack: 2, CreepAttack: 4, Defense: 2, DebuffMax: 5, Mobility: 2, CreepPush: 0}
	case PositionHardSupport:
		return Stats{PlayerAttack: 2, CreepAttack: 4, Defense: 1, DebuffMax: 5, Mobility: 1, CreepPush: 2}
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
