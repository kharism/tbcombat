package main

import (
	"fmt"
	"log"

	"github.com/kharism/tbcombat"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cs *tbcombat.CombatScene
}

func (g *Game) Update() error {
	err := g.cs.Update()
	if err != nil {
		return err
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	g.cs.Draw(screen)
}
func (g *Game) Layout(width, length int) (int, int) {
	return 640, 480
}

var (
	imgPool *ImagePool
)

type Layout struct {
	oppStartX int
}

func NewLayout(Opponents []*tbcombat.Character) *Layout {
	totalOppWidth := 0
	for _, c := range Opponents {
		fmt.Println(c.OriImageWidth, int(c.Image.ScaleParam.Sx))
		totalOppWidth += int(float64(c.OriImageWidth) * c.Image.ScaleParam.Sx)
	}
	fmt.Println("totalOppWidth", totalOppWidth)
	l := &Layout{oppStartX: (640 / 2) - totalOppWidth/2}
	return l
}
func (l *Layout) StartOppPos() (int, int) {
	return l.oppStartX, 100
}

// starting party position
func (l *Layout) StartPartyPos() (int, int) {
	return 20, 400
}

// menu start pos
func (l *Layout) MenuStartPos() (int, int) {
	return 0, 0
}

// log location
func (l *Layout) LogStartPos() (int, int) {
	return 0, 300
}

type CombatInitalState struct {
	party    []*tbcombat.Character
	opp      []*tbcombat.Character
	bgImage  *ebiten.Image
	layouter *Layout
}

func (s *CombatInitalState) GetPartyChar() []*tbcombat.Character {
	return s.party
}
func (s *CombatInitalState) GetOpponents() []*tbcombat.Character {
	return s.opp
}
func (s *CombatInitalState) GetBgImage() *ebiten.Image {
	return s.bgImage
}
func (s *CombatInitalState) GetLayouter() tbcombat.Layouter {
	return s.layouter
}
func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	imgPool = &ImagePool{Map: map[string]*ebiten.Image{}}
	initialState := &CombatInitalState{
		party: []*tbcombat.Character{NewWarrior()},
		opp:   []*tbcombat.Character{NewSlime("A"), NewSlime("B"), NewSlime("C")},
	}
	initialState.layouter = NewLayout(initialState.opp)
	cs := tbcombat.NewCombatScene(initialState)
	g := &Game{cs: cs}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
