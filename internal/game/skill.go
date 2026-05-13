package game

// SkillID uniquely identifies a skill.
type SkillID string

// SkillTier indicates when the skill becomes available.
type SkillTier string

const (
	SkillTierRegular SkillTier = "regular" // offered at waves 0, 2, 4 (7 per position)
	SkillTierSpecial SkillTier = "special" // offered after wave 6 (2 per position)
)

// Skill describes a single learnable ability.
//
// For each player, RemainingSkillIDs is a shuffled list of skill IDs for their position.
// When a skill offer is needed, the last 2 IDs are presented as options, then both are popped off after selection.
// This keeps skill selection simple and ensures variety between matches.
type Skill struct {
	ID       SkillID
	Position Position
	Tier     SkillTier
	Name     string
	Desc     string
}

// SkillSelectWaves are the wave numbers where regular skill selection occurs.
var SkillSelectWaves = []int{0, 2, 4}

// SpecialSkillWave is the wave after which special skill selection occurs.
const SpecialSkillWave = 6

// SkillsPerOffer is how many choices are presented at each selection.
const SkillsPerOffer = 2

// PicksPerOffer is how many the player chooses from each offer.
const PicksPerOffer = 1

// Skill selection logic:
// At each selection, present the last 2 SkillIDs from RemainingSkillIDs as options.
// After a pick, pop both IDs off the slice.
