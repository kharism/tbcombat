package tbcombat

type Skill interface {
	Command
	GetName() string
	GetDescription() string
}
