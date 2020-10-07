package sprite

import (
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// WalkCounter プレイヤーの歩きモーションカウンタ 最大8
// ref: wWalkCounter
var WalkCounter uint

// PlayerMovingDirection ref: wPlayerMovingDirection
var PlayerMovingDirection uint

// UpdatePlayerSprite update sprite direction and anim counter
// if in moving, increment anim counter
// if player is starting moving, change direction and increment anim counter
func UpdatePlayerSprite() {
	p := store.SpriteData[0]
	if p == nil {
		return
	}

	if p.WalkAnimationCounter != 0 {
		if p.WalkAnimationCounter > 0 {
			p.WalkAnimationCounter--
		}
		DisableSprite(0)
		return
	}

	if WalkCounter == 0 {
		if PlayerMovingDirection == 0 {
			p.AnimationFrame = 0
			return
		}
		p.Direction = util.Direction(PlayerMovingDirection)
	}

	p.AnimationFrame++
	if p.AnimationCounter() == 4 {
		p.AnimationFrame = 0
	}
	p.VRAM.Index = int(p.Direction + (p.AnimationFrame >> 2))
}

// AdvancePlayerSprite advance player's walk by a frame
func AdvancePlayerSprite() {
	p := store.SpriteData[0]
	if WalkCounter == 0 {
		p.MapXCoord += p.DeltaX
		p.MapYCoord += p.DeltaY
	}

	store.SCX += p.DeltaX
	store.SCY += p.DeltaY

	for i, s := range store.SpriteData {
		if i == 0 {
			continue
		}
		s.ScreenXPixel -= p.DeltaX
		s.ScreenYPixel -= p.DeltaY
	}
}