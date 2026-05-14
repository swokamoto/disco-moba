package game

// Environmental effects for match variety and replayability
// Add to Match struct: ActiveEnvironments []EnvironmentEffect

type EnvironmentEffect int

const (
	EnvNone         EnvironmentEffect = iota
	EnvFrostbite                      // If you don't move, gain 2 debuff
	EnvOvercrowding                   // 4+ on same team in lane: all take 2 debuff
	EnvBlazingSun                     // Healing is halved
	EnvMoraleSurge                    // Morale loss is doubled
)

var EnvironmentEffectNames = map[EnvironmentEffect]string{
	EnvFrostbite:    "Frostbite (Cold)",
	EnvOvercrowding: "Overcrowding",
	EnvBlazingSun:   "Blazing Sun",
	EnvMoraleSurge:  "Morale Surge",
}

// Example: Add to Match struct
// ActiveEnvironments []EnvironmentEffect // set at match start

// Example: How to check/apply in resolution logic
// for _, env := range match.ActiveEnvironments {
//   switch env {
//   case EnvFrostbite:
//     if player.Lane == player.PreviousLane { player.DebuffLevel += 2 }
//   case EnvOvercrowding:
//     if countTeamInLane(team, lane) >= 4 { ... }
//   case EnvBlazingSun:
//     healing = healing / 2
//   case EnvMoraleSurge:
//     moraleLoss = moraleLoss * 2
//   }
// }
