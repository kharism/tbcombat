package tbcombat

const (
	CMD_ATK = "CMD_ATK"
)

// Command is
type Command interface {
	// Handle
	Execute(target *Character)
	// speed determine the priority of an action in a queue
	GetSpeed() int
}

type AttackCommand struct {
	attacker *Character
}

func (c *AttackCommand) Execute(target *Character) {

}
func (c *AttackCommand) GetSpeed() int {
	return 1
}
