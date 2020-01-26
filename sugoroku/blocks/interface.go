package blocks

import "fmt"

type Block interface {
	String() string
}

type BlockType int

const (
	SPACE BlockType = iota
	FIELD
	SPECIAL

	CITY      // 街
	RESTHOUSE // 宿屋
	SHOP      // 道具屋
	ARMSSHOP  // 武器防具屋
	MAGICSHOP // 魔法屋
	KEYPOINT  // 目的地
)

type Space struct{}

func New(t BlockType) Block {
	switch t {
	case SPACE:
		return &Space{}
	case FIELD:
		return &Field{}
	}

	return &Space{}
}

func (p *Space) String() string {
	return fmt.Sprint("space")
}

type Field struct {
}

func (p *Field) String() string {
	return fmt.Sprint("field")
}
