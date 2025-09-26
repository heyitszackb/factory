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


func (o *Orchestrator) Step() {
	o.updateAdders()
	
	o.updateDeleters()
}


/*

	update_adders() {
		adder_entity_ids = entitiesdb.get_all_entity_ids_with_property('adder')
		for adder_entity_id in adder_entity_ids:	
			// 0. create an entity to spawn (in the future this might be the type of adder)
			// Once we create the entity we can assume the props are valid as the check would be on the constructor
            if !should_place_new_entity():
                continue

            props_for_potential_new_entity = ['movable']

            // 0. validate that the props combination is legal
            if !is_valid_props(props_for_potential_new_entity):
                continue
			
			
			// 1. validate to check if we can spawn an entity at the coord (generic, used by adder) with the props
			coord_of_adder_entity = entity_manager.get_coord_of_entity_by_id(adder_entity_id)

			can_spawn = can_spawn_entity_at_coord_with_props(props_for_potential_new_entity, coord_of_adder_entity)
			// 2. actually spawn an entity at that coord
			entity_manager.add_entity_at_coord_with_props(props_for_potential_new_entity, coord_of_adder_entity)
	}

*/


func (o *Orchestrator) updateAdders() {
	adderEntityIds := o.entityManager.GetAllEntityIDsWithProperty(ADDER)
	for adderEntityId := range adderEntityIds {
		// 0. create an entity to spawn (in the future this might be the type of adder)
		// Once we create the entity we can assume the props are valid as the check would be on the constructor
		
		if !shouldPlaceNewEntity {
			continue
		}
		
		propertiesForPotentialNewEntity := []Property{MOVABLE}
		// 0. validate that the props combination is legal
		if !IsValidProps(propertiesForPotentialNewEntity) {
			continue
		}

		// 1. validate to check if we can spawn an entity at the coord (generic, used by adder) with the props
		coordOfAdderEntity := o.entityManager.GetCoordOfEntityByID(adderEntityId)
		/*
		// ZCB ADD
		canSpawn := can_spawn_entity_at_coord_with_props(props_for_potential_new_entity, coord)
		*/

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

    // checks to make sure props are internally consistent inside of the entity
    is_valid_props(props) {
        /*
        Business rules:
            0. Entities must have at least one prop
            1. Entities can only be: output_left, output_right, output_top, output_bottom, adder , deleter, movable
            2. Entities can have only one of: output_left, output_right, output_top, output_bottom, adder , deleter, movable
            3. Entities can only have 1 prop unless it is a combinatino of (output_left, output_right, output_top and output_bottom)
        */
        valid_props = ['output_left','output_right','output_top','output_bottom','adder','deleter','movable']
        output_props = ['output_left','output_right','output_top','output_bottom']
        
        // Rule 0:
        if len(props) == 0:
            return false

        // Rule 3:
        if len(props) > 1:
            for prop in props:
                if prop not in output_props:
                    return false

        seen_props = []
        for prop in props:
            // Rule 1:
            if prop not in valid_props:
                return false
            // Rule 2:
            if prop in seen_props:
                return false
            seen_props.append(prop)

        return true
    }

    // for now, hardcode to some fixed value but can come from other props in the future
    should_place_new_entity() {
        if random.randint(1) < 0.01: // for now, spawn a new entity at a 1% chance each frame
            return true
        return false
    }
	
	can_spawn_entity_at_coord_with_props(props: []string, coord): bool {
		// before we add an entity we need to make sure the entity is valid given the other entities on it's coord		
		
		entity_ids_at_coord = entity_manager.get_entity_ids_at_coord(coord)
		/* As a business logic rule, I will say for now:
		- there can only by one entity with property movable in a cell
		- there can only be one entity with property adder in a cell
		- there can only be one entity with property deleter in a cell
		*/
		
		num_adders = 0
		num_deleters = 0
		num_movable = 0
		for entitiy_id_at_coord in entity_ids_at_coord:
            if entity_manager.entity_has_property(entitiy_id_at_coord, 'adder')
				num_adders++
                if entity_manager.entity_has_property(entitiy_id_at_coord, 'deleter')
				num_deleters++
                if entity_manager.entity_has_property(entitiy_id_at_coord, 'movable')
				num_movable++
		
        if 'adder' in props and num_adder > 0:
			return false
        if 'deleter' in props and num_deleters > 0:
			return false
        if 'movable' in props and num_movable > 0:
			return false
			
		return true
	}
	
	move(): void {
		entity_ids_already_moved = []
		moves = generate_moves(entities_already_moved)
	
		while len(moves) > 0:
			execute_moves(moves)
			entity_ids_already_moved = update_entities_already_moved(entity_ids_already_moved, moves)
		
			moves = generate_moves(entity_ids_already_moved)
	}
	
	update_entities_already_moved(entity_ids_already_moved,moves) {
		// 1. Get unique entity IDs from the moves
		new_entity_ids_already_moved = entity_ids_already_moved
		for entity_id, coord in moves:
			if entity_id not in unique_entity_ids:
				new_entity_ids_already_moved.append(entity_id)
		return new_entity_ids_already_moved
	}
	
	
	generate_moves(entities_already_moved): [](entity_id, Coord) {
		all_valid_possible_moves = get_all_valid_possible_moves(entities_already_moved)			
		final_moves = resolve_final_moves_from_all_valid_moves(all_valid_possible_moves) 	
		return final_moves
	}
	
	get_all_valid_possible_moves(entities_already_moved) {
		entity_ids = entity_manager.get_all_entity_ids() <- returns entity_ids
		all_valid_possible_moves = []
		for entity_id in entity_ids:
			if entity_id in entities_already_moved:
				valid_moves = get_all_valid_moves_for_entity(entity_id)
				for valid_move in valid_moves:
					all_valid_possible_moves.append((entity_id,valid_move))
		return all_valid_possible_moves
	}
	
	
	
	execute_final_moves(final_moves) {
		for entity_id, Coord in final_moves:
			entity_manager.move_entity_with_id_to_coord(entity_id, Coord)
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
	
	get_all_valid_moves_for_entity(entity_id): []Coord {
		coords_of_valid_moves_for_entity = []
		output_directions = get_output_directions_for_cell_with_entity(entity_id)
		for output_direction in output_directions:
			coord_to_move_to = map_output_direction_to_coord(entity_id, output_direction)
			is_valid_move = is_entity_move_valid(entity, coord_to_move_to)
			if is_valid_move:
				coords_of_valid_moves_for_entity.append(coord_to_move_to)
		return coords_of_valid_moves_for_entity
	}
	
	// exclude considering output directions of properties of self
	get_output_directions_for_cell_with_entity(entity_id) {
		output_direction_properties = []
		coord_of_entity = entity_manager.get_coord_of_entity_by_id(entity_id)
		entity_ids_on_cell_with_entity = entity_manager.get_entity_ids_at_coord(coord_of_entity)
		output_properties = [
			"output_left",
			"output_right",
			"output_top",
			"output_bottom",
		]
		for entity_id_on_cell_with_entity in entity_ids_on_cell_with_entity:
            if entity_id_on_cell_with_entity == entity_id:
                continue
			properties = entity_manager.get_all_properties_of_entity_with_id(entity_id_on_cell_with_entity)
			for property1 in output_properties:
				for property2 in properties:
					if property1 == property2:
						if property1 not in output_direction_properties:
							output_direction_properties.append(property1)
		return output_direction_properties
	}
	
	is_entity_move_valid(entity_id, coord_to_move_to) {
        if !entity_manager.entity_has_property(entity_id, 'movable'):
            return false
		coord_of_entity = entity_manager.get_coord_of_entity_by_id(entity_id)
		if coord_of_entity.isAtRightOf(coord_to_move_to):
			coord_to_move_to = coord_to_right <- alias the variable
			entity_ids_at_right_of_entity = entity_manager.get_entity_ids_at_coord(coord_to_right)
			for entity_id_to_right in entity_ids_at_right_of_entity:
                if entity_manager.entity_has_property(entity_id_to_right, 'output_left'):
                    return false
                if entity_manager.entity_has_property(entity_id_to_right, 'movable'):
                    return false
			return true
		elif coord_of_entity.isATLeftOf(coord_to_move_to):
            coord_to_move_to = coord_to_left <- alias the variable
            entity_ids_at_left_of_entity = entity_manager.get_entity_ids_at_coord(coord_to_left)
            for entity_id_to_left in entity_ids_at_left_of_entity:
                if entity_manager.entity_has_property(entity_id_to_left, 'output_right'):
                    return false
                if entity_manager.entity_has_property(entity_id_to_left, 'movable'):
                    return false
            return true
		elif coord_of_entity.isAtTopOf(coord_to_move_to):
            coord_to_move_to = coord_to_top <- alias the variable
            entity_ids_at_top_of_entity = entity_manager.get_entity_ids_at_coord(coord_to_top)
            for entity_id_to_top in entity_ids_at_top_of_entity:
                if entity_manager.entity_has_property(entity_id_to_top, 'output_bottom'):
                    return false
                if entity_manager.entity_has_property(entity_id_to_top, 'movable'):
                    return false
            return true
		elif coord_of_entity.isAtBottomOf(coord_to_move_to):
            coord_to_move_to = coord_to_bottom <- alias the variable
            entity_ids_at_bottom_of_entity = entity_manager.get_entity_ids_at_coord(coord_to_bottom)
            for entity_id_to_bottom in entity_ids_at_bottom_of_entity:
                if entity_manager.entity_has_property(entity_id_to_bottom, 'output_top'):
                    return false
                if entity_manager.entity_has_property(entity_id_to_bottom, 'movable'):
                    return false
            return true
		return false
	}
	
	
	
	
	map_output_direction_to_coord(entity_id, output_direction) {
		if output_direction == "output_right":
            // 0. get coord at entity
            // 1. set coord to be the right of that entity ZCB (Come back to fix this)
			coord = get_coord_at_right_of(entity_id)
		elif output_direction == "output_left":
			coord = get_coord_at_left_of(entity_id)
		elif output_direction == "output_top":
			coord = get_coord_at_top_of(entity_id)
		elif output_direction == "output_bottom":
			coord = get_coord_at_bottom_of(entity_id)
		return coord
	}
	
	get_coord_at_right_of(entity_id) {
		coord_of_entity = get_coord_of_entity_by_id(entity_id)
		
	}
}