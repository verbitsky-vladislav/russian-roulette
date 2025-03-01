package types

import (
	"encoding/json"
	"fmt"
	"github.com/ericlagergren/decimal"
	"strconv"
)

type Decimal struct {
	*decimal.Big
}

func NewDecimal(big *decimal.Big) *Decimal {
	return &Decimal{big}
}

func (d *Decimal) LessThan(other *Decimal) bool {
	return d.Cmp(other.Big) < 0
}

func (d *Decimal) GreaterThan(other *Decimal) bool {
	return d.Cmp(other.Big) > 0
}

func (d *Decimal) Equal(other *Decimal) bool {
	return d.Cmp(other.Big) == 0
}

func (d *Decimal) Neg() *Decimal {
	return &Decimal{new(decimal.Big).Neg(d.Big)}
}

func (d *Decimal) Abs() *Decimal {
	return &Decimal{new(decimal.Big).Abs(d.Big)}
}

func (d *Decimal) Sub(other *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Sub(d.Big, other.Big)}
}

func (d *Decimal) LessThanOrEqual(other *Decimal) bool {
	return d.Cmp(other.Big) <= 0
}

func (d *Decimal) GreaterThanOrEqual(other *Decimal) bool {
	return d.Cmp(other.Big) >= 0
}

func (d *Decimal) Add(other *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Add(d.Big, other.Big)}
}

func (d *Decimal) Mul(other *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Mul(d.Big, other.Big)}
}

func (d *Decimal) Div(other *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Quo(d.Big, other.Big)}
}

func SumDecimal(a, b *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Add(a.Big, b.Big)}
}

func MulDecimal(a, b *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Mul(a.Big, b.Big)}
}

func DivDecimal(a, b *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Quo(a.Big, b.Big)}
}

func SubDecimal(a, b *Decimal) *Decimal {
	return &Decimal{new(decimal.Big).Sub(a.Big, b.Big)}
}

func String(d *Decimal) string {
	if d.Big == nil {
		return "0.00000000000000000000"
	}
	return DecimalToString(d.Big)
}

func ValidString(d *Decimal) string {
	if d.Big == nil {
		return "0.00000000000000000000"
	}
	return DecimalToString(d.Big)
}

func NewDecimalFromString(s string) *Decimal {
	return &Decimal{BigFromString(s)}
}

func NewDecimalFromInt(i int) *Decimal {
	return &Decimal{decimal.New(int64(i), 0)}
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(DecimalToString(d.Big))
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)

	if err != nil {
		str = string(data)
	}

	dec := BigFromString(str)
	d.Big = dec
	return nil
}

func DecimalToString(d *decimal.Big) string {
	if d == nil {
		return "0.00000000000000000000"
	}
	return fmt.Sprintf("%.20f", d)
}

func DecimalToInt(d *Decimal) int {
	decimalStr := DecimalToString(d.Big)
	//parse string to int
	numberFlt, err := strconv.ParseFloat(decimalStr, 64)
	if err != nil {
		panic(err)
	}
	return int(numberFlt)

}

func BigFromString(s string) *decimal.Big {
	if s == "" {
		return nil
	}
	d := new(decimal.Big)
	_, ok := d.SetString(s)

	if !ok {
		return nil
	}
	return d
}

func (d *Decimal) String() string {
	return DecimalToString(d.Big)
}

var ZeroDecimal = NewDecimalFromInt(0)
