package dreamer

import "math/rand"


type Orchestrator struct {
	entityManager *EntityManager
}

func NewOrchestrator(entityManager *EntityManager) *Orchestrator {
	return &Orchestrator{
		entityManager: entityManager,
	}
}

func (o *Orchestrator) update() {
	o.updateAdders()
	
	o.updateDeleters()
}

func (o *Orchestrator) move() {
	entityIdsAlreadyMoved := make([]EntityID,0,100)
	moves = o.generateMoves(entityIdsAlreadyMoved)

	for len(moves) > 0 {
		o.executeMoves(moves)
		o.updateEntityIdsAlreadyMoved(&entityIdsAlreadyMoved, moves)
		moves = o.generateMoves(entityIdsAlreadyMoved)
	}
}


func (o *Orchestrator) updateEntityIdsAlreadyMoved(entityIdsAlreadyMoved *[]EntityID, moves []Move) {
	// Add unique entity IDs from the moves
	for _, move := range moves {
		// Check if this entity ID is not already in the list
		if !containsEntity(*entityIdsAlreadyMoved, move.EntityID) {
			*entityIdsAlreadyMoved = append(*entityIdsAlreadyMoved, move.EntityID)
		}
	}
}

func (o *Orchestrator) executeMoves(moves []Move) {
	for _, move := range moves {
		o.entityManager.MoveEntityWithIDToCoord(move.EntityID, move.Coord)
	}
}

func (o *Orchestrator) generateMoves(entityIdsAlreadyMoved []EntityID) []Move {
	allValidPossibleMoves := o.getAllValidPossibleMoves(entityIdsAlreadyMoved)
	finalMoves := o.resolveFinalMovesFromAllValidMoves(allValidPossibleMoves)
	return finalMoves
}

func (o *Orchestrator) getAllValidPossibleMoves(entityIdsAlreadyMoved []EntityID) []Move {
	entityIds = o.entity_manager.GetAllEntityIDs()
	allValidPossibleMoves := make([]Move, 0, 100)
	for _, entityId := range entityIds {
		if containsEntity(entityIdsAlreadyMoved, entityId) {
            continue // skip entities that already moved
        }
		validMoveCoords := o.getAllValidMoveCoordsForEntity(entityId)
		for _, validMoveCoord := range validMoveCoords {
			allValidPossibleMoves = append(allValidPossibleMoves, Move{
                EntityID: entityId,
                Coord:    validMoveCoord,
            })
		}
	}
	return allValidPossibleMoves
}


func (o *Orchestrator) getAllValidMoveCoordsForEntity(entityId EntityID) {
	coordsOfValidMovesForEntity := make([]Coord, 0, 50)
	outputDirections := o.getOutputDirectionsForCellWithEntityID(entityId)
	for outputDirection := range outputDirections {
		coordToMoveTo := o.mapOutputDirectionToCoord(entityId, outputDirection)
		isValidMove := o.isEntityMoveValid(entityId, coordToMoveTo)
		if isValidMove {
			coords_of_valid_moves_for_entity := append(coords_of_valid_moves_for_entity, coordToMoveTo)
		}
	}
	return coordsOfValidMovesForEntity
}

func (o *Orchestrator) isEntityMoveValid(entityId EntityID, coordToMoveTo Coord) bool {
	// Business rule: we can't move the entity if it doesn't have the movable property
	if !o.entityManager.EntityHasProperty(entityId, MOVABLE) {
		return false
	}
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityId)
	
	if coordOfEntity.IsAtRightOf(coordToMoveTo) {
		// Moving to the right
		entityIdsAtRightOfEntity := o.entityManager.GetEntityIDsAtCoord(coordToMoveTo)
		for _, entityIdToRight := range entityIdsAtRightOfEntity {
			if o.entityManager.EntityHasProperty(entityIdToRight, OUTPUT_LEFT) {
				return false
			}
			if o.entityManager.EntityHasProperty(entityIdToRight, MOVABLE) {
				return false
			}
		}
		return true
	} else if coordOfEntity.IsAtLeftOf(coordToMoveTo) {
		// Moving to the left
		entityIdsAtLeftOfEntity := o.entityManager.GetEntityIDsAtCoord(coordToMoveTo)
		for _, entityIdToLeft := range entityIdsAtLeftOfEntity {
			if o.entityManager.EntityHasProperty(entityIdToLeft, OUTPUT_RIGHT) {
				return false
			}
			if o.entityManager.EntityHasProperty(entityIdToLeft, MOVABLE) {
				return false
			}
		}
		return true
	} else if coordOfEntity.IsAtTopOf(coordToMoveTo) {
		// Moving to the top
		entityIdsAtTopOfEntity := o.entityManager.GetEntityIDsAtCoord(coordToMoveTo)
		for _, entityIdToTop := range entityIdsAtTopOfEntity {
			if o.entityManager.EntityHasProperty(entityIdToTop, OUTPUT_BOTTOM) {
				return false
			}
			if o.entityManager.EntityHasProperty(entityIdToTop, MOVABLE) {
				return false
			}
		}
		return true
	} else if coordOfEntity.IsAtBottomOf(coordToMoveTo) {
		// Moving to the bottom
		entityIdsAtBottomOfEntity := o.entityManager.GetEntityIDsAtCoord(coordToMoveTo)
		for _, entityIdToBottom := range entityIdsAtBottomOfEntity {
			if o.entityManager.EntityHasProperty(entityIdToBottom, OUTPUT_TOP) {
				return false
			}
			if o.entityManager.EntityHasProperty(entityIdToBottom, MOVABLE) {
				return false
			}
		}
		return true
	}
	return false
}

func (o *Orchestrator) mapOutputDirectionToCoord(entityID EntityID, outputDirection Property) Coord {
	switch outputDirection {
	case OUTPUT_RIGHT:
		return o.getCoordAtRightOf(entityID)
	case OUTPUT_LEFT:
		return o.getCoordAtLeftOf(entityID)
	case OUTPUT_TOP:
		return o.getCoordAtTopOf(entityID)
	case OUTPUT_BOTTOM:
		return o.getCoordAtBottomOf(entityID)
	default:
		// Return the current coordinate if direction is not recognized
		return o.entityManager.GetCoordOfEntityByID(entityID)
	}
}

func (o *Orchestrator) getCoordAtRightOf(entityID EntityID) Coord {
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityID)
	return NewCoord(coordOfEntity.Row, coordOfEntity.Col+1)
}

func (o *Orchestrator) getCoordAtLeftOf(entityID EntityID) Coord {
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityID)
	return NewCoord(coordOfEntity.Row, coordOfEntity.Col-1)
}

func (o *Orchestrator) getCoordAtTopOf(entityID EntityID) Coord {
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityID)
	return NewCoord(coordOfEntity.Row-1, coordOfEntity.Col)
}

func (o *Orchestrator) getCoordAtBottomOf(entityID EntityID) Coord {
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityID)
	return NewCoord(coordOfEntity.Row+1, coordOfEntity.Col)
}


func (o *Orchestrator) getOutputDirectionsForCellWithEntityID(entityID EntityID) []Property {
	outputDirectionProperties := []Property{}
	coordOfEntity := o.entityManager.GetCoordOfEntityByID(entityID)
	entityIdsOnCellWithEntity := o.entityManager.GetEntityIDsAtCoord(coordOfEntity)
	
	outputProperties := []Property{
		OUTPUT_LEFT,
		OUTPUT_RIGHT,
		OUTPUT_TOP,
		OUTPUT_BOTTOM,
	}
	
	for _, entityIdOnCellWithEntity := range entityIdsOnCellWithEntity {
		if entityIdOnCellWithEntity == entityID {
			continue
		}
		entityProperties := o.entityManager.GetAllPropertiesOfEntityWithID(entityIdOnCellWithEntity)
		for _, outputProperty := range outputProperties {
			for _, entityProperty := range entityProperties {
				if outputProperty == entityProperty {
					// Check if outputProperty is not already in outputDirectionProperties
					found := false
					for _, existingProperty := range outputDirectionProperties {
						if existingProperty == outputProperty {
							found = true
							break
						}
					}
					if !found {
						outputDirectionProperties = append(outputDirectionProperties, property1)
					}
				}
			}
		}
	}
	return outputDirectionProperties
}

func containsEntity(list []EntityID, id EntityID) bool {
    for _, v := range list {
        if v == id {
            return true
        }
    }
    return false
}

func (o *Orchestrator) resolveFinalMovesFromAllValidMoves() []Move {

}


func (o *Orchestrator) Step() {
	o.update()
	o.move()
}


func (o *Orchestrator) updateAdders() {
	adderEntityIds := o.entityManager.GetAllEntityIDsWithProperty(ADDER)
	for adderEntityId := range adderEntityIds {
		// 0. create an entity to spawn (in the future this might be the type of adder)
		// Once we create the entity we can assume the props are valid as the check would be on the constructor
		
		if !shouldPlaceNewEntity() {
			continue
		}
		
		propertiesForPotentialNewEntity := []Property{MOVABLE}
		// 0. validate that the props combination is legal
		if !IsValidProps(propertiesForPotentialNewEntity) {
			continue
		}

		// 1. validate to check if we can spawn an entity at the coord (generic, used by adder) with the props
		coordOfAdderEntity := o.entityManager.GetCoordOfEntityByID(adderEntityId)
		
		canSpawn := canSpawnEntityAtCoordWithProperties(propertiesForPotentialNewEntity, coordOfAdderEntity)

		// 2. actually spawn an entity at that coord
		o.entityManager.AddEntityAtCoordWithProperties(propertiesForPotentialNewEntity, coordOfAdderEntity)

		

		
	}
}

// Helper
// IsValidProps validates an entity's properties against business rules:
//
// 0. Entities must have at least one prop
// 1. Entities can only be: output_left, output_right, output_top, output_bottom, adder, deleter, movable
// 2. Entities can only have one of each prop (no duplicates)
// 3. Entities can only have >1 prop if all are outputs (left, right, top, bottom)
func IsValidProps(properties []Property) bool {
	if len(properties) == 0 {
		return false // rule 0
	}

	// Define valid and output-only sets
	validProperties := map[Property]struct{}{
		OUTPUT_LEFT:   {},
		OUTPUT_RIGHT:  {},
		OUTPUT_TOP:    {},
		OUTPUT_BOTTOM: {},
		ADDER:         {},
		DELETER:       {},
		MOVABLE:       {},
	}

	outputProperties := map[Property]struct{}{
        OUTPUT_LEFT:   {},
        OUTPUT_RIGHT:  {},
        OUTPUT_TOP:    {},
        OUTPUT_BOTTOM: {},
    }

	// Rule 3: if more than one property, they must all be outputs
    if len(properties) > 1 {
        for _, property := range properties {
            _, isOutputProperty := outputProps[property];
			if !isOutputProperty {
                return false
            }
        }
    }

	seen := make(map[Property]struct{})
    for _, property := range properties {
        // Rule 1: must be valid
        if _, ok := validProperties[property]; !ok {
            return false
        }
        // Rule 2: no duplicates
        if _, already := seen[property]; already {
            return false
        }
        seen[property] = struct{}{}
    }

    return true

}

// Helper
func shouldPlaceNewEntity() bool {
	return rand.Float64() < 0.01
}

func (o *Orchestrator) canSpawnEntityAtCoordWithProperties(properties []Property, coord Coord) bool {
	// before we add an entity we need to make sure the entity is valid given the other entities on it's coord		
	entityIdsAtCoord = o.entity_manager.GetEntityIDsAtCoord(coord)
	/* As a business logic rule, I will say for now:
		- there can only by one entity with property movable in a cell
		- there can only be one entity with property adder in a cell
		- there can only be one entity with property deleter in a cell
	*/
	numAdders := 0
	numDeleters := 0
	numMovable := 0

	for _, entityIdAtCoord := range entityIdsAtCoord {
		if o.entityManager.EntityHasProperty(entityIdAtCoord, ADDER) {
			numAdders++
		}
		if o.entityManager.EntityHasProperty(entityIdAtCoord, DELETER) {
			numDeleters++
		}
		if o.entityManager.EntityHasProperty(entityIdAtCoord, MOVABLE) {
			numMovable++
		}
	}

	// if the new entity wants any of these properties and there's already one, reject
	if hasProp(props, ADDER) && numAdders > 0 {
		return false
	}
	if hasProp(props, DELETER) && numDeleters > 0 {
		return false
	}
	if hasProp(props, MOVABLE) && numMovable > 0 {
		return false
	}

	return true
	
}

func hasProp(props []Property, p Property) bool {
	for _, v := range props {
		if v == p {
			return true
		}
	}
	return false
}





func (o *Orchestrator) updateDeleters() {
	deleterEntityIds := o.entityManager.GetAllEntityIDsWithProperty(DELETER)
	for deleterEntityId := range deleterEntityIds {
		// A business rule is now that we should delete all other entities on this cell other than this deleter
		// 1. Get all entities on the same cell as the deleter
		coordOfDeleter := o.entityManager.GetCoordOfEntityByID(deleterEntityId)
		entityIdsOnCellWithDeleter := o.entityManager.GetEntityIDsAtCoord(coordOfDeleter)

		// 2. Remove all other entities on the same cell as the deleter, but not the deleter itself
		for entityIdOnCellWithDeleter := range entityIdsOnCellWithDeleter {
			if entityIdOnCellWithDeleter != deleterEntityId {
				o.entityManager.RemoveEntity(entityIdOnCellWithDeleter)
			}
		}
	}
}
	
	
	resolve_final_moves_from_all_valid_moves(moves_list:[](entity_id, Coord)) { // returns one move per entity
		while !is_valid_moveset(moves_list):
			// 1. pick random unique entity
			entity_id_of_entity_to_move = get_random_entity_id_from_moves_list(moves_list)
			
			// 2. pick a random unique move of that entity
			random_move_from_selected_entity_id = get_random_move_from_entity_id(entity_id,moves_list)
			
			// Delete necessary moves from moves_list
			filtered_moves_list = filter_moves_list(random_move_from_selected_entity_id, moves_list)
			moves_list = filtered_moves_list
		return moves_list
	}
	
	
	filter_moves_list(move_to_filter_by, moves_list) {
		updated_moves_list = []
		for move in moves_list:
			// 3. Delete all other moves that have the same destination
			if move.entity_id == move_to_filter_by.entity_id:
				continue
			// 4. Delete all other moves that that entity could take
			if move.coord == move_to_filter_by.coord:
				continue
			updated_moves_list.append(move)
		return updated_moves_list
	}
	
	get_random_move_from_entity_id(moves_list) {
		moves_with_entity_id = get_moves_with_entity_id(entity_id, moves_list)
		return random.choice(moves_with_entity_id)
	}
	
	get_moves_with_entity_id(entity_id, moves_list) {
		moves_with_entity_id = []
		for move in moves_list:
			if move.entity_id == entity_id: // or move[0] to get the entity_id in the tuple
				moves_with_entity_id.append(move)
		return moves_with_entity_id
	}
	
	// could be implemented a few ways but this works for now
	get_random_entity_id_from_moves_list(moves_list): entity_id {
		return random_choice(moves_list.entity)
	}
	
	// a valid moveset is no repeated entities or destinations
	is_valid_moveset(moves_list) {
		unique_destinations = []Coord <- array of Coords
		unique_entity_ids = [] <- array of entity IDs
		for entity_id, coord in moves_list:
			if entity_id in unique_entity_ids:
				return false
			if coord in unique_destinations:
				return false
	
			unique_destinations.append(coord)
			unique_entity_ids.append(entity_id)
		return true
	}
}