package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ACoin struct {
	Name string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewACoin() ACoin {
	return ACoin{
		Name: "Acoin",
		Owner: nil,
	}
}

func (a ACoin) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s, Owner: %s`, a.Name, a.Owner))
}