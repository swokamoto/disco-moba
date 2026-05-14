package game

// SkillEffect describes how a skill modifies lane stats.
type SkillEffect string

const (
	EffectStaticBonus    SkillEffect = "static_bonus"  // flat bonus
	EffectProcChance     SkillEffect = "proc_chance"   // % chance to apply Value
	EffectUnitsInLane    SkillEffect = "units_in_lane" // bonus per unit present (TargetTeam: self/enemy)
	EffectCreepCount     SkillEffect = "creep_count"   // bonus per N creep
	EffectTowerHealth    SkillEffect = "tower_health"  // bonus based on tower health
	EffectMorale         SkillEffect = "morale"        // bonus if morale below threshold
	EffectMovement       SkillEffect = "movement"      // special case for mobility-based skills	
)

// SkillType indicates what stat the skill modifies.
type SkillType string

const (
	SkillTypeAttack   SkillType = "attack"
	SkillTypeHeal     SkillType = "heal" // removes debuffs from allies
)

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
// Skill struct now includes Type (Attack or Heal) and Value (amount to add).
// Skill struct now supports different effect types and extra fields for proc, scaling, and threshold skills.
// TargetTeam distinguishes whether a threshold applies to "self" or "enemy" team.
type Skill struct {
	ID         SkillID
	Position   Position
	Tier       SkillTier
	Type       SkillType   // "attack" or "heal"
	Effect     SkillEffect // how the bonus is applied
	Value      int         // base amount to add
	Chance     float64     // for proc_chance (0.0–1.0)
	Threshold  int         // for morale/tower/creep threshold skills
	Multiplier float64     // for scaling skills (e.g., per N creep)
	TargetTeam string      // "self" or "enemy" (for threshold-based skills)
	Name       string
	Desc       string
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

// Example positive counter and synergy skills:
// Example positive counter and synergy skills:

// Organized skill pools by position for easy editing and lookup.
var CarrySkills = []Skill{
	{
		ID:       "power_surge",
		Position: PositionCarry,
		Tier:     SkillTierRegular,
		Type:     SkillTypeAttack,
		Effect:   EffectStaticBonus,
		Value:    2,
		Name:     "Power Surge",
		Desc:     "+2 Attack.",
	},
	{
		ID:         "threat_response",
		Position:   PositionCarry,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      1,
		TargetTeam: "enemy",
		Name:       "Threat Response",
		Desc:       "+1 Attack for each enemy in your lane.",
	},
	{
		ID:         "overwhelm",
		Position:   PositionCarry,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      1,
		TargetTeam: "self",
		Name:       "Overwhelm",
		Desc:       "+1 Attack per ally in lane.",
	},
	{
		ID:         "comeback_strike",
		Position:   PositionCarry,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectMorale,
		Value:      3,
		Threshold:  40,
		TargetTeam: "enemy",
		Name:       "Comeback Strike",
		Desc:       "+3 Attack if enemy morale < 40.",
	},
	{
		ID:         "creep_cleaver",
		Position:   PositionCarry,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectCreepCount,
		Value:      1,
		Multiplier: 5,
		TargetTeam: "enemy",
		Name:       "Creep Cleaver",
		Desc:       "+1 Attack for every 5 enemy creep in lane.",
	},
	{
		ID:         "finisher",
		Position:   PositionCarry,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectTowerHealth,
		Value:      3,
		Threshold:  30,
		TargetTeam: "enemy",
		Name:       "Finisher",
		Desc:       "+3 Attack if enemy tower < 30 health.",
	},
	{
		ID:       "lucky_hit",
		Position: PositionCarry,
		Tier:     SkillTierRegular,
		Type:     SkillTypeAttack,
		Effect:   EffectProcChance,
		Value:    4,
		Chance:   0.25,
		Name:     "Lucky Hit",
		Desc:     "25% chance to gain +4 Attack this wave.",
	},
	// Special skills
	{
		ID:         "last_chance",
		Position:   PositionCarry,
		Tier:       SkillTierSpecial,
		Type:       SkillTypeAttack,
		Effect:     EffectMorale,
		Value:      6,
		Threshold:  30,
		TargetTeam: "self",
		Name:       "Last Chance",
		Desc:       "+6 Attack if your team's morale is below 30.",
	},
	{
		ID:       "berserker_luck",
		Position: PositionCarry,
		Tier:     SkillTierSpecial,
		Type:     SkillTypeAttack,
		Effect:   EffectProcChance,
		Value:    8,
		Chance:   0.5,
		Name:     "Berserker's Luck",
		Desc:     "50% chance to gain +8 Attack this wave.",
	},
}

var MidSkills = []Skill{
	// Solo-focused
	{
		ID:         "arcane_isolation",
		Position:   PositionMid,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      3,
		Threshold:  0, // no allies in lane
		TargetTeam: "self",
		Name:       "Arcane Isolation",
		Desc:       "+3 Attack if you are the only ally in your lane.",
	},
	{
		ID:         "pressure_response",
		Position:   PositionMid,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      1,
		TargetTeam: "enemy",
		Name:       "Pressure Response",
		Desc:       "+1 Attack for each enemy in your lane.",
	},
	{
		ID:       "momentum_burst",
		Position: PositionMid,
		Tier:     SkillTierRegular,
		Type:     SkillTypeAttack,
		Effect:   EffectProcChance,
		Value:    4,
		Chance:   0.33,
		Name:     "Momentum Burst",
		Desc:     "33% chance to gain +4 Attack this wave.",
	},
	// Teamplay-focused
	{
		ID:         "rally_strike",
		Position:   PositionMid,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      1,
		TargetTeam: "self",
		Name:       "Rally Strike",
		Desc:       "+1 Attack for each ally in your lane.",
	},
	{
		ID:       "lane_leadership",
		Position: PositionMid,
		Tier:     SkillTierRegular,
		Type:     SkillTypeAttack,
		Effect:   EffectStaticBonus,
		Value:    2,
		Name:     "Lane Leadership",
		Desc:     "+2 Attack if you have more allies than enemies in your lane.",
	},
	{
		ID:         "push_amplifier",
		Position:   PositionMid,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectCreepCount,
		Value:      1,
		Multiplier: 5, // per 5 friendly creep
		TargetTeam: "self",
		Name:       "Push Amplifier",
		Desc:       "+1 Attack for every 5 friendly creep in your lane.",
	},
	{
		ID:         "tactical_overwatch",
		Position:   PositionMid,
		Tier:       SkillTierRegular,
		Type:       SkillTypeAttack,
		Effect:     EffectMorale,
		Value:      2,
		Threshold:  60,
		TargetTeam: "self",
		Name:       "Tactical Overwatch",
		Desc:       "+2 Attack if your team's morale is above 60.",
	},
	// Special skills
	{
		ID:         "solo_supremacy",
		Position:   PositionMid,
		Tier:       SkillTierSpecial,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      6,
		Threshold:  0, // no allies in lane
		TargetTeam: "self",
		Name:       "Solo Supremacy",
		Desc:       "+6 Attack if you are the only ally in your lane.",
	},
	{
		ID:         "teamplay_maestro",
		Position:   PositionMid,
		Tier:       SkillTierSpecial,
		Type:       SkillTypeAttack,
		Effect:     EffectUnitsInLane,
		Value:      2,
		TargetTeam: "self",
		Name:       "Teamplay Maestro",
		Desc:       "+2 Attack for each ally in your lane.",
	},
}

var OfflaneSkills = []Skill{
	   // Tank/soak skills
	   {
		   ID:         "iron_wall",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeHeal,
		   Effect:     EffectStaticBonus,
		   Value:      2, // Remove 2 debuffs from self
		   Name:       "Iron Wall",
		   Desc:       "Remove 2 debuffs from yourself at the start of each wave.",
	   },
	   {
		   ID:         "damage_sponger",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeHeal,
		   Effect:     EffectUnitsInLane,
		   Value:      1, // Remove 1 debuff per ally present
		   TargetTeam: "self",
		   Name:       "Damage Sponger",
		   Desc:       "Remove 1 debuff from each ally in your lane.",
	   },
	   {
		   ID:         "tower_guardian",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeAttack,
		   Effect:     EffectTowerHealth,
		   Value:      2, // +2 Attack if your tower is below 50 health
		   Threshold:  50,
		   Name:       "Tower Guardian",
		   Desc:       "+2 Attack if your tower is below 50 health.",
	   },
	   // Movement/empty lane skills
	   {
		   ID:         "lane_swapper",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeAttack,
		   Effect:     EffectMovement,
		   Value:      2, // +2 Attack if you changed lanes this wave
		   Name:       "Lane Swapper",
		   Desc:       "+2 Attack if you moved to a new lane this wave.",
	   },
	   {
		   ID:         "ghost_push",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeAttack,
		   Effect:     EffectUnitsInLane,
		   Value:      3, // +3 Attack if no enemies in lane
		   Threshold:  0,
		   TargetTeam: "enemy",
		   Name:       "Ghost Push",
		   Desc:       "+3 Attack if there are no enemies in your lane.",
	   },
	   {
		   ID:         "tower_sapper",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeAttack,
		   Effect:     EffectCreepCount,
		   Value:      1, // +1 Attack per 5 friendly creep
		   Multiplier: 5,
		   TargetTeam: "self",
		   Name:       "Tower Sapper",
		   Desc:       "+1 Attack for every 5 friendly creep in your lane.",
	   },
	   {
		   ID:         "relentless_roamer",
		   Position:   PositionOfflane,
		   Tier:       SkillTierRegular,
		   Type:       SkillTypeAttack,
		   Effect:     EffectMovement,
		   Value:      3, // +3 Attack if you changed lanes this wave (proc logic in resolution)
		   Chance:     0.33, // 33% chance if you moved
		   Name:       "Relentless Roamer",
		   Desc:       "33% chance to gain +3 Attack if you moved to a new lane this wave.",
	   },
	   // Special skills
	   {
		   ID:         "fortress",
		   Position:   PositionOfflane,
		   Tier:       SkillTierSpecial,
		   Type:       SkillTypeHeal,
		   Effect:     EffectStaticBonus,
		   Value:      4, // Remove 4 debuffs from self
		   Name:       "Fortress",
		   Desc:       "Remove 4 debuffs from yourself at the start of each wave.",
	   },
	   {
		   ID:         "backdoor_artist",
		   Position:   PositionOfflane,
		   Tier:       SkillTierSpecial,
		   Type:       SkillTypeAttack,
		   Effect:     EffectUnitsInLane,
		   Value:      6, // +6 Attack if no enemies in lane
		   Threshold:  0,
		   TargetTeam: "enemy",
		   Name:       "Backdoor Artist",
		   Desc:       "+6 Attack if there are no enemies in your lane.",
	   },
}

var SoftSupportSkills = []Skill{
	// Add Soft Support skills here
}

var HardSupportSkills = []Skill{
	// Add Hard Support skills here
}

// Organized skill pools by position for easy editing and lookup.
// Influence skills removed. All skills now use Attack, Heal, or debuff/block mechanics.
// TODO: Update each skill to use only Attack, Heal, or new stat model.
// (Skill definitions to be refactored in next step)
