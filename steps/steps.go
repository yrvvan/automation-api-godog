package steps

import (
	"context"
	"errors"
	"fmt"
)

type Item struct {
	ID   int
	Name string
}

var items []Item
var nextID int = 1

type godogsCtxKey struct{}

func IEat(ctx context.Context, num int) (context.Context, error) {
	available, ok := ctx.Value(godogsCtxKey{}).(int)
	if !ok {
		return ctx, errors.New("there are no godogs available")
	}

	if available < num {
		return ctx, fmt.Errorf("you cannot eat %d godogs, there are %d available", num, available)
	}

	available -= num

	return context.WithValue(ctx, godogsCtxKey{}, available), nil
}

func ThereShouldBeRemaining(ctx context.Context, remaining int) error {
	available, ok := ctx.Value(godogsCtxKey{}).(int)
	if !ok {
		return errors.New("there are no godogs available")
	}

	if available != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, available)
	}

	return nil
}

// CRUD functions
func ThereAreNoItems() error {
	items = []Item{}
	nextID = 1
	return nil
}

func ICreateAnItemNamed(qty int, name string) error {
	if len(items) > 0 {
		items = []Item{}
	}

	nextID = 1
	i := 0
	for i < qty {
		items = append(items, Item{ID: nextID, Name: name})
		nextID++
		i++
	}
	return nil
}

func IShouldHaveItems(count int) error {
	if len(items) != count {
		return fmt.Errorf("expected %d items, but got %d", count, len(items))
	}
	return nil
}

func AnItemNamed(name string) error {
	for _, item := range items {
		if item.Name == name {
			return nil
		}
	}
	return fmt.Errorf("no item named %s found", name)
}

func IReadTheItemWithID(id int) error {
	for _, item := range items {
		if item.ID == id {
			return nil
		}
	}
	return fmt.Errorf("no item with ID %d found", id)
}

func IShouldGetTheItemWithName(name string) error {
	for _, item := range items {
		if item.Name == name {
			return nil
		}
	}
	return fmt.Errorf("expected item with name %s not found", name)
}

func IUpdateTheItemWithIDToHaveName(id int, name string) error {
	for i, item := range items {
		if item.ID == id {
			items[i].Name = name
			return nil
		}
	}
	return fmt.Errorf("no item with ID %d found", id)
}

func TheItemWithIDShouldHaveName(id int, name string) error {
	for _, item := range items {
		if item.ID == id {
			if item.Name == name {
				return nil
			}
			return fmt.Errorf("expected item with ID %d to have name %s, but got %s", id, name, item.Name)
		}
	}
	return fmt.Errorf("no item with ID %d found", id)
}

func IDeleteTheItemWithID(id int) error {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no item with ID %d found", id)
}
