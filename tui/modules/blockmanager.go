package modules

type BlockManager struct {
	focusManagers []FocusManager
	active        int
}

func NewBlockManager(focusManagers ...FocusManager) BlockManager {
	return BlockManager{
		focusManagers: focusManagers,
		active:        0,
	}
}

func (b *BlockManager) Set(cursor int) {
	if cursor > len(b.focusManagers)-1 {
		b.active = 0
	} else if cursor < 0 {
		b.active = len(b.focusManagers) - 1
	} else {
		b.active = cursor
	}

	for _, fm := range b.focusManagers {
		fm.BlurAll()
	}

	fm := b.Active()

	fm.Set(fm.FocusedIndex())
}

func (b *BlockManager) Active() *FocusManager {
	return &b.focusManagers[b.active]
}

func (b BlockManager) ActiveIndex() int {
	return b.active
}

func (b *BlockManager) Next() {
	b.Set(b.active + 1)
}

func (b *BlockManager) Prev() {
	b.Set(b.active - 1)
}

func (b *BlockManager) Focus() {
	b.Active().Set(b.Active().FocusedIndex())
}

func (b *BlockManager) BlurAll() {
	for _, fm := range b.focusManagers {
		fm.BlurAll()
	}
}
