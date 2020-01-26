package card

import (
	"fmt"
	"github.com/go-numb/go-optional-setter"
)

type Base struct {
	// 配布全体の分布とレアリティ
	ID     int
	Rarity optional.Float64
	Artist optional.String

	// Card informations
	Cost     optional.Float64
	Name     optional.String
	Descript optional.String
	Type     optional.String
	Tap      optional.Bool

	Attack       optional.Float64
	Defence      optional.Float64
	Magic        optional.Float64
	MagicDefence optional.Float64
	// Effect is post effective
	Effect Effect

	Filename optional.String

	// Attr is 属性
	Attr       optional.Int
	Attributes []string
}

// Effect effect values to enemy or me.
type Effect struct {
	Type     optional.Int
	Add      optional.Float64
	Sub      optional.Float64
	Descript optional.String
}

// NewBase creates card data
func NewBase(attributes []string) *Base {
	return &Base{
		Attributes: attributes,
	}
}

// Status is interface, does printf status.
func (p *Base) Status() string {
	return fmt.Sprintf(`コスト: %sx%.f Tap: %t
%s
%s
攻撃力: %.f
防御力: %.f
魔力: %.f
魔防力: %.f
%s`,
		p.Attributes[p.Attr.Value()],
		p.Cost.Value(),
		p.Tap.Value(),
		p.Name.Value(),
		p.Descript.Value(),
		p.Attack.Value(),
		p.Defence.Value(),
		p.Magic.Value(),
		p.MagicDefence.Value(),
		p.Effect.Descript.Value())
}
