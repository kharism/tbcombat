package tbcombat

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kharism/hanashi/core"
)

type CombatSubstate interface {
	Draw(screen *ebiten.Image)
	Update()
	OnLoad()
}
type MenuButton struct {
	*core.MovableImage
	Label       string
	cursorIn    bool
	onClickFunc func()
}

func (b *MenuButton) Draw(screen *ebiten.Image) {
	if b.cursorIn {
		b.ScaleParam.Sx = 1.1
	} else {
		b.ScaleParam.Sx = 1
	}
	btnX, btnY := b.MovableImage.GetPos()
	b.MovableImage.Draw(screen)
	text.Draw(screen, b.Label, core.DefaultFont, int(btnX)+10, int(btnY)+30, color.White)
}
func (b *MenuButton) Update() {
	curX, curY := ebiten.CursorPosition()
	butPosX, butPosY := b.GetPos()
	width, height := b.GetSize()
	// fmt.Println(width, height)
	if curX > int(butPosX) && curX < int(butPosX+width) && curY > int(butPosY) && curY < int(butPosY+height) {
		b.cursorIn = true
		// fmt.Println("Cursor In")
	} else {
		b.cursorIn = false
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		if b.cursorIn && b.onClickFunc != nil {
			b.onClickFunc()
		}
	}
}

type MainCombatMenu struct {
	combatScene *CombatScene

	buttons []*MenuButton
}

func (b *MainCombatMenu) OnLoad() {

}
func NewMainCombatMenu(combatScene *CombatScene, buttonBase *ebiten.Image) CombatSubstate {
	jj := MainCombatMenu{combatScene: combatScene}
	// btn, _ := imgPool.GetImage("icon/blue_button00.png")
	// attack button
	atkBtnImgParam := core.NewMovableImageParams().WithMoveParam(core.MoveParam{Sy: 80, Sx: 0})
	atkButton := &MenuButton{MovableImage: core.NewMovableImage(buttonBase, atkBtnImgParam), Label: "Attack"}
	atkButton.onClickFunc = func() {
		cc := &AttackCommand{attacker: combatScene.Party[combatScene.CurrentChrIdx]}
		combatScene.CombatQueue = append(combatScene.CombatQueue, cc)
		combatScene.CurrentChrIdx += 1
		for true {
			if combatScene.CurrentChrIdx >= len(combatScene.Party) {
				combatScene.CurrentChrIdx = -1
				combatScene.BeginCombat()
				break
			}
			if combatScene.Party[combatScene.CurrentChrIdx].HP <= 0 {
				combatScene.CurrentChrIdx += 1
				continue
			}
		}

	}
	// back button
	backBtnImgParam := core.NewMovableImageParams().WithMoveParam(core.MoveParam{Sy: 120, Sx: 0})
	backButton := &MenuButton{MovableImage: core.NewMovableImage(buttonBase, backBtnImgParam), Label: "Back"}
	backButton.onClickFunc = func() {
		combatScene.CurrentChrIdx -= 1
		combatScene.CombatQueue = combatScene.CombatQueue[:len(combatScene.CombatQueue)-1]
	}
	jj.buttons = append(jj.buttons, atkButton, backButton)
	return &jj
}
func (mm *MainCombatMenu) Draw(screen *ebiten.Image) {
	for _, b := range mm.buttons {
		b.Draw(screen)
	}
}
func (mm *MainCombatMenu) Update() {
	for _, b := range mm.buttons {
		b.Update()
	}
}
