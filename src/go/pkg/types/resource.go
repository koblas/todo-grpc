package types

import (
	"fmt"

	"github.com/rs/xid"
)

type ResourceType interface{ Prefix() string }

type ID[T ResourceType] xid.ID

func New[T ResourceType]() ID[T] { return ID[T](xid.New()) }

func (id ID[T]) String() string {
	var resourceType T // create the default value for the resource type

	return fmt.Sprintf(
		"%s_%s",
		resourceType.Prefix(), // Extract the "prefix" we want from the resource type
		xid.ID(id).String(),   // Use XID's string marshalling
	)
}
