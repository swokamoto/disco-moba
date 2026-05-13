package game

// Lane represents one of the three main lanes.
type Lane string

const (
	LaneLeft  Lane = "left"
	LaneMid   Lane = "mid"
	LaneRight Lane = "right"
)

// AllLanes lists lanes in order from left to right.
var AllLanes = []Lane{LaneLeft, LaneMid, LaneRight}

// LaneIndex returns the 0-based index of a lane (used for array lookups).
// Pure function.
func LaneIndex(l Lane) int {
	for i, lane := range AllLanes {
		if lane == l {
			return i
		}
	}
	return -1
}

// ValidLane returns true if l is a known lane.
func ValidLane(l Lane) bool {
	for _, lane := range AllLanes {
		if l == lane {
			return true
		}
	}
	return false
}
