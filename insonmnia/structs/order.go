package structs

import (
	"errors"
	"math/big"
	"time"

	pb "github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
)

// Order represents a safe order wrapper.
//
// This is used to decompose the validation out of the protocol. All
// methods must return the valid sub-structures.
type Order struct {
	*pb.Order
}

// ByPrice implements sort.Interface; it allows for sorting Orders by Price field.
type ByPrice []*Order

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].GetPrice().Cmp(a[j].GetPrice()) == 1 }

func (o *Order) Unwrap() *pb.Order {
	return o.Order
}

func NewOrder(o *pb.Order) (*Order, error) {
	if err := validateOrder(o); err != nil {
		return nil, err
	} else {
		return &Order{o}, nil
	}
}

func validateOrder(o *pb.Order) error {
	if o == nil {
		return errors.New("order cannot be nil")
	}

	if o.PricePerSecond == nil {
		return errors.New("price cannot be nil")
	}

	if o.PricePerSecond.Unwrap().Sign() <= 0 {
		return errors.New("price/sec must be positive")
	}

	if err := validateSlot(o.Slot); err != nil {
		return err
	}

	return nil
}

func (o *Order) GetID() string {
	return o.Order.GetId()
}

func (o *Order) SetID(ID string) {
	o.Order.Id = ID
}

func (o *Order) GetPrice() *big.Int {
	v, err := util.ParseBigInt(o.Price)
	// We've already checked that during construction, but who knows...
	if err != nil {
		panic("order price has failed big.Int conversion")
	}
	return v
}

func (o *Order) GetSlot() *Slot {
	slot, err := NewSlot(o.Order.GetSlot())
	if err != nil {
		panic("validation has failed")
	}
	return slot
}

func (o *Order) IsBid() bool {
	return o.GetOrderType() == pb.OrderType_BID
}

func (o *Order) GetDuration() time.Duration {
	return time.Duration(o.Slot.Duration) * time.Second
}

func (o *Order) GetTotalPrice() *big.Int {
	bigDuration := big.NewInt(int64(o.GetDuration().Seconds()))
	return big.NewInt(0).Mul(o.PricePerSecond.Unwrap(), bigDuration)
}
