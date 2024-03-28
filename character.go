package tbcombat

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kharism/hanashi/core"
)

// character in this combat
type Character struct {
	Name           string
	HP             int
	MaxHP          int
	Image          *core.MovableImage
	OriImageWidth  int
	OriImageHeight int
	Stats          map[string]int

	CustomCommandName []string
	CustomCommandFunc []Skill
}

func (c *Character) Draw(screen *ebiten.Image) {
	c.Image.Draw(screen)
}
