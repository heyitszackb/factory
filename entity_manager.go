package main

import (
	"math"
	"os"
	"regexp"
	"strings"
)

/*
There are no restrictions on the grid in the entity manager.
It just exposes an API for the orchestrator to interact with the entities
and fetch certain information about the entities inside of the grid.
Natively, the grid allows any entities to be placed in any cells.
Validation must happen at the orchestrator level.
The entity manager just knows how to get/set different information on entities.
*/
type EntityManager struct {
	coordToEntityIDs    map[Coord][]EntityID
	entityIDToCoord     map[EntityID]Coord
	entityIDToEntity    map[EntityID]Entity
	propertyToEntityIDs map[Property][]EntityID
	coordToEntities     map[Coord][]Entity
}

func NewEntityManager(filepath string) *EntityManager {
	entityManager := &EntityManager{
		coordToEntityIDs:    make(map[Coord][]EntityID),
		entityIDToCoord:     make(map[EntityID]Coord),
		entityIDToEntity:    make(map[EntityID]Entity),
		propertyToEntityIDs: make(map[Property][]EntityID),
		coordToEntities:     make(map[Coord][]Entity),
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		// handle error
	}

	fileContents := string(content)

	re := regexp.MustCompile(`\((.*?)\)`)
	matches := re.FindAllStringSubmatch(fileContents, -1)
	var encodedEntityList []string

	for _, match := range matches {
		// match[1] is the inside of the parentheses
		trimmed := strings.TrimSpace(match[1])
		encodedEntityList = append(encodedEntityList, trimmed)
	}

	gridSize := int(math.Sqrt(float64(len(encodedEntityList))))
	row := 0
	col := 0

	for _, encodedEntity := range encodedEntityList {
		properties := []Property{}

		// range yields runes, so compare to rune literals like 'A'
		for _, ch := range encodedEntity {
			switch ch {
			case 'A':
				properties = append(properties, ADDER)
			case 'D':
				properties = append(properties, DELETER)
			case 'M':
				properties = append(properties, MOVABLE)
			case '>':
				properties = append(properties, OUTPUT_RIGHT)
			case '<':
				properties = append(properties, OUTPUT_LEFT)
			case '^':
				properties = append(properties, OUTPUT_TOP)
			case 'v':
				properties = append(properties, OUTPUT_BOTTOM)
			}
		}
		entityManager.AddEntityAtCoordWithProperties(NewCoord(row, col), properties)
		col++
		if col == gridSize {
			row++
			col = 0
		}
	}
	return entityManager
}

// ZCB OPTIMIZATION - should these be returning pointers?
func (e *EntityManager) GetCoordToEntities() map[Coord][]Entity {
	return e.coordToEntities
}

func (e *EntityManager) GetAllEntityIDsWithProperty(property Property) []EntityID {
	return e.propertyToEntityIDs[property]
}

func (e *EntityManager) GetCoordOfEntityByID(entityID EntityID) Coord {
	return e.entityIDToCoord[entityID]
}

func (e *EntityManager) GetEntityIDsAtCoord(coord Coord) []EntityID {
	return e.coordToEntityIDs[coord]
}

func (e *EntityManager) GetAllPropertiesOfEntityWithID(entityID EntityID) []Property {
	entity := e.entityIDToEntity[entityID]
	return entity.Properties
}

func (e *EntityManager) GetAllEntityIDs() []EntityID {
	ids := make([]EntityID, 0, len(e.entityIDToEntity)) // preallocate with map size
	for id := range e.entityIDToEntity {
		ids = append(ids, id)
	}
	return ids
}

// MoveEntityWithIDToCoord moves an entity from its current coord to a new coord.
func (e *EntityManager) MoveEntityWithIDToCoord(entityID EntityID, newCoord Coord) {
	// Step 0: Find the current coord
	oldCoord, found := e.entityIDToCoord[entityID]
	if !found {
		return // nothing to move
	}

	// Step 1: Remove entityID from the old coord's list
	entityIDsAtOldCoord, found := e.coordToEntityIDs[oldCoord]
	if found {
		updatedIDs := removeEntityIDFromSlice(entityIDsAtOldCoord, entityID)
		if len(updatedIDs) == 0 {
			delete(e.coordToEntityIDs, oldCoord)
		} else {
			e.coordToEntityIDs[oldCoord] = updatedIDs
		}
	}

	// Step 2: Add entityID to the new coord's list
	e.coordToEntityIDs[newCoord] = append(e.coordToEntityIDs[newCoord], entityID)

	// Step 3: Update entityID → coord mapping
	e.entityIDToCoord[entityID] = newCoord
}

/*
   // returns true if entity with entity ID has property 'property'
   entity_has_property(entity_id, property: string): bool {
       entity_ids_with_prop = property_to_entity_ids[property]
       if entity_ids_with_prop:
           for entity_id_with_prop in entity_ids_with_prop:
               if entity_id_with_prop == entity_id:
                   return true
       return false
   }



*/

func (e EntityManager) EntityHasProperty(entity_id EntityID, property Property) bool {
	entityIdsWithProperty, found := e.propertyToEntityIDs[property]
	if found {
		for _, entityIdWithProperty := range entityIdsWithProperty {
			if entityIdWithProperty == entity_id {
				return true
			}
		}
	}
	return false
}

func (e *EntityManager) AddEntityAtCoordWithProperties(coord Coord, properties []Property) {
	newEntity := NewEntity(properties)

	e.coordToEntityIDs[coord] = append(e.coordToEntityIDs[coord], newEntity.ID)
	e.entityIDToCoord[newEntity.ID] = coord
	e.entityIDToEntity[newEntity.ID] = *newEntity
	e.coordToEntities[coord] = append(e.coordToEntities[coord], *newEntity)
	for _, property := range properties {
		e.propertyToEntityIDs[property] = append(e.propertyToEntityIDs[property], newEntity.ID)
	}
}

func (e *EntityManager) RemoveEntity(entityID EntityID) {
	// 0) Locate the entity and coord; bail if not found.
	coord, ok := e.entityIDToCoord[entityID]
	if !ok {
		return
	}
	entity, ok := e.entityIDToEntity[entityID]
	if !ok {
		return
	}

	// 1. Remove from coord_to_entity
	entitiesAtCoord, found := e.coordToEntities[coord]
	if found {
		updatedEntitiesAtCoord := removeEntityFromSlice(entitiesAtCoord, entityID)
		if len(updatedEntitiesAtCoord) == 0 {
			delete(e.coordToEntities, coord)
		} else {
			e.coordToEntities[coord] = updatedEntitiesAtCoord
		}
	}

	// 2. Go to that coord and remove the entity from that coord list (because we are only allowing one coord per entity, this works!)
	// coord_to_entity_ids[coord].remove(entity_id)
	entityIdsAtCoord, found := e.coordToEntityIDs[coord]
	if found {
		updatedEntityIdsAtCoord := removeEntityIDFromSlice(entityIdsAtCoord, entityID)
		if len(updatedEntityIdsAtCoord) == 0 {
			delete(e.coordToEntityIDs, coord)
		} else {
			e.coordToEntityIDs[coord] = updatedEntityIdsAtCoord
		}
	}

	// 3. Remove Entity ID
	delete(e.entityIDToCoord, entityID)

	// 4. Remove the entity ID for each property this entity_id is in
	for _, property := range entity.Properties {
		if entityIdsForProperty, found := e.propertyToEntityIDs[property]; found {
			updatedIDs := removeEntityIDFromSlice(entityIdsForProperty, entityID)
			if len(updatedIDs) == 0 {
				delete(e.propertyToEntityIDs, property)
			} else {
				e.propertyToEntityIDs[property] = updatedIDs
			}
		}
	}

	// 5. Remove the entity
	delete(e.entityIDToEntity, entityID)
}

// removeEntityIDFromSlice returns a new slice without the target EntityID.
func removeEntityIDFromSlice(ids []EntityID, targetID EntityID) []EntityID {
	writeIndex := 0
	for _, currentID := range ids {
		if currentID != targetID {
			ids[writeIndex] = currentID
			writeIndex++
		}
	}
	return ids[:writeIndex]
}

// Helper method
// removeEntityFromSlice returns a new slice with the entity removed.
func removeEntityFromSlice(entities []Entity, targetID EntityID) []Entity {
	writeIndex := 0
	for _, currentEntity := range entities {
		if currentEntity.ID != targetID {
			entities[writeIndex] = currentEntity
			writeIndex++
		}
	}
	return entities[:writeIndex]
}
