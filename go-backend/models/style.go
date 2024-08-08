package models

import (
	"fmt"
)

type Style struct {
	id   int
	name string
}

var styles = getAllStyles()

func NewStyle(id int, name string) *Style {
	return &Style{id: id, name: name}
}

// these methods and var styles need to be refactored to service and repository
func GetStyleById(id int) (*Style, error) {
	for i := 0; i < len(styles); i++ {
		if styles[i].id == id {
			return &styles[i], nil
		}
	}
	return nil, fmt.Errorf("invalid id {%v} for style", id)
}

func getAllStyles() []Style {
	values := []string{"Slab", "Überhang", "Slopey", "Lang", "Kurz", "Statisch", "Dynamisch", "Stretchy", "Kante", "Kraftig", "Einfach hoch", "Um Ecke"}
	styleArray := []Style{}
	for i := 0; i < len(values); i++ {
		styleArray = append(styleArray, *NewStyle(i, values[i]))
	}
	return styleArray
}
