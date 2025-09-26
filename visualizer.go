package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Visualizer struct {
	entityManager *EntityManager
}

func NewVisualizer(entityManager *EntityManager) *Visualizer {
	return &Visualizer{
		entityManager: entityManager,
	}
}

func (v *Visualizer) Draw(step int) {
	coordToEntities := v.entityManager.GetCoordToEntities()

	// Find the grid bounds
	maxRow, maxCol := 0, 0
	for coord := range coordToEntities {
		if coord.Row > maxRow {
			maxRow = coord.Row
		}
		if coord.Col > maxCol {
			maxCol = coord.Col
		}
	}

	// Clear screen
	v.clearScreen()

	// Draw the grid
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			coord := NewCoord(row, col)
			entities, exists := coordToEntities[coord]

			if !exists || len(entities) == 0 {
				fmt.Print(".")
				continue
			}

			// Create a list of properties on this coord
			propertiesOnCoord := make([]Property, 0)
			for _, entity := range entities {
				propertiesOnCoord = append(propertiesOnCoord, entity.Properties...)
			}

			// Determine what to draw based on properties
			if v.hasProperty(propertiesOnCoord, ADDER) {
				fmt.Print("A ")
			} else if v.hasProperty(propertiesOnCoord, DELETER) {
				fmt.Print("D ")
			} else if v.hasProperty(propertiesOnCoord, MOVABLE) {
				fmt.Print("M ")
			} else if len(propertiesOnCoord) > 0 {
				// At this point, we know it doesn't contain adder, deleter, or movable
				// so it must only contain output properties
				// if it contains 1 output property, output V, >, <, or ^ respectively
				if len(propertiesOnCoord) == 1 {
					// Display the specific output direction symbol
					switch propertiesOnCoord[0] {
					case OUTPUT_RIGHT:
						fmt.Print("> ")
					case OUTPUT_LEFT:
						fmt.Print("< ")
					case OUTPUT_TOP:
						fmt.Print("^ ")
					case OUTPUT_BOTTOM:
						fmt.Print("v ")
					default:
						fmt.Print("? ") // Unknown property
					}
				} else {
					// Multiple output properties - show count
					numOutputProperties := len(propertiesOnCoord)
					fmt.Printf("%d ", numOutputProperties)
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println() // New line after each row
	}

	// Display step number at the end
	fmt.Printf("Step: %d\n", step)
}

func (v *Visualizer) DrawDebug(step int) {
	coordToEntities := v.entityManager.GetCoordToEntities()

	// Clear screen
	v.clearScreen()

	fmt.Println("=== DEBUG VIEW ===")
	fmt.Printf("Step: %d\n\n", step)

	// Find the grid bounds
	maxRow, maxCol := 0, 0
	for coord := range coordToEntities {
		if coord.Row > maxRow {
			maxRow = coord.Row
		}
		if coord.Col > maxCol {
			maxCol = coord.Col
		}
	}

	// Display properties for each cell
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			coord := NewCoord(row, col)
			entities, exists := coordToEntities[coord]

			fmt.Printf("Coord %d, %d: ", row, col)

			if !exists || len(entities) == 0 {
				fmt.Println("EMPTY")
				continue
			}

			// Collect all properties from all entities at this coord
			allProperties := make([]Property, 0)
			for _, entity := range entities {
				allProperties = append(allProperties, entity.Properties...)
			}

			// Display properties
			if len(allProperties) == 0 {
				fmt.Println("NO_PROPERTIES")
			} else {
				for i, property := range allProperties {
					if i > 0 {
						fmt.Print(", ")
					}
					fmt.Print(v.propertyToString(property))
				}
				fmt.Println()
			}
		}
	}
}

func (v *Visualizer) propertyToString(property Property) string {
	switch property {
	case ADDER:
		return "ADDER"
	case DELETER:
		return "DELETER"
	case MOVABLE:
		return "MOVABLE"
	case OUTPUT_LEFT:
		return "OUTPUT_LEFT"
	case OUTPUT_RIGHT:
		return "OUTPUT_RIGHT"
	case OUTPUT_TOP:
		return "OUTPUT_TOP"
	case OUTPUT_BOTTOM:
		return "OUTPUT_BOTTOM"
	default:
		return "UNKNOWN"
	}
}

func (v *Visualizer) hasProperty(properties []Property, targetProperty Property) bool {
	for _, property := range properties {
		if property == targetProperty {
			return true
		}
	}
	return false
}

func (v *Visualizer) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
