package main

import (
	"github.com/kharism/tbcombat"

	"github.com/kharism/hanashi/core"
)

func NewSlime(name string) *tbcombat.Character {
	s := &tbcombat.Character{}
	s.Name = "Slime " + name
	s.HP = 10
	s.MaxHP = 10
	param := core.MovableImageParams{}
	param.WithScale(&core.ScaleParam{Sx: 0.25, Sy: 0.25})
	image, _ := imgPool.GetImage("./img/opp/slime.png")
	s.OriImageWidth = image.Bounds().Dx()
	s.OriImageHeight = image.Bounds().Dy()
	s.Image = core.NewMovableImage(image, &param)
	return s
}
