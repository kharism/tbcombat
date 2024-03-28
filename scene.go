package tbcombat

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kharism/hanashi/core"
)

type CombatInitalState interface {
	GetPartyChar() []*Character
	GetOpponents() []*Character
	GetBgImage() *ebiten.Image
	GetLayouter() Layouter
}
type Layouter interface {
	// starting opp position
	StartOppPos() (int, int)

	// starting party position
	StartPartyPos() (int, int)

	// menu start pos
	MenuStartPos() (int, int)

	// log location
	LogStartPos() (int, int)
}

type CombatScene struct {
	Party        []*Character
	Opponents    []*Character
	BgImg        *ebiten.Image
	CombatQueue  []Command
	Layouter     Layouter
	MenuSubState CombatSubstate

	CurrentChrIdx int

	CustomDrawOpp  func(*ebiten.Image, []*Character, *CombatScene)
	CustomDrawChar func(*ebiten.Image, []*Character, *CombatScene)
}

func NewCombatScene(is CombatInitalState) *CombatScene {
	cs := &CombatScene{Party: is.GetPartyChar(), Opponents: is.GetOpponents(), CombatQueue: []Command{}}
	cs.BgImg = is.GetBgImage()
	cs.Layouter = is.GetLayouter()
	return cs
}
func (cs *CombatScene) Update() error {
	if cs.MenuSubState != nil {
		cs.Update()
	}
	return nil
}
func (cs *CombatScene) BeginCombat() {

}
func (cs *CombatScene) DrawOpp(screen *ebiten.Image) {
	startX, startY := cs.Layouter.StartOppPos()
	for idx, c := range cs.Opponents {
		// fmt.Println("Opp Pos", float64(startX)+float64(idx*c.OriImageWidth)*c.Image.ScaleParam.Sx)
		c.Image.SetPos(float64(startX)+float64(idx*c.OriImageWidth)*c.Image.ScaleParam.Sx, float64(startY))
		c.Draw(screen)
	}
}
func (cs *CombatScene) DrawChar(screen *ebiten.Image) {
	curFont := core.DefaultFont

	for idx, c := range cs.Party {
		box := ebiten.NewImage(160, 80)
		box.Fill(color.RGBA{0, 169, 0, 255})
		opt := ebiten.DrawImageOptions{}
		posX, posY := cs.Layouter.StartPartyPos()
		boxPosX := float64(posX + idx*170)
		text.Draw(box, c.Name, curFont, 0, 15, color.White)
		text.Draw(box, "HP "+strconv.Itoa(c.HP), curFont, 0, 40, color.White)
		opt.GeoM.Translate(boxPosX, float64(posY))
		if idx == cs.CurrentChrIdx {
			box2 := ebiten.NewImage(170, 100)
			box2.Fill(color.RGBA{169, 169, 169, 255})
			opt := ebiten.DrawImageOptions{}
			boxPosX := float64(15 + idx*170)
			opt.GeoM.Translate(boxPosX, float64(posY-5))
			screen.DrawImage(box2, &opt)
		}
		screen.DrawImage(box, &opt)
	}
}
func (cs *CombatScene) Draw(screen *ebiten.Image) {
	if cs.BgImg != nil {
		screen.DrawImage(cs.BgImg, nil)
	}
	if cs.CustomDrawOpp != nil {
		cs.CustomDrawOpp(screen, cs.Opponents, cs)
	} else {
		cs.DrawOpp(screen)
	}

	if cs.MenuSubState != nil {
		cs.Draw(screen)
	}

	if cs.CustomDrawChar != nil {
		cs.CustomDrawChar(screen, cs.Opponents, cs)
	} else {
		cs.DrawChar((screen))
	}
}
