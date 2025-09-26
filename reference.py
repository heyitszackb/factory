/*
[DONE] ZCB: Need to think about the API for entitiesdb and the grid (w/ entity_id vs entity usage)
[DONE] ZCB: Grid, entitiesdb API.
[DONE] ZCB: Create Entity struct definition
[DONE] ZCB: Implement Coord methods
[DONE] ZCB: Implement methods on Entity_manager that work (DO NOT hyper optimize)
[DONE] ZBC: Create grid initializer
ZCB: Create ceel visualizer
ZBC: Break out each component into it's own file as-is (entity_manager, orchestrator, visualizer, coord, etc)
ZCB: Convert each section to Go structs (add NO unit tests)
ZCB: Create a sample input text file
ZCB: Run and test!

ZCB: When initializing the grid from a file input, we should have some way to validate the input in a similar way that we validate it
in the adder (although I am fine keeping this logic in 2 places because one means what the orchestrator can do, and the other means
what the entity initializer can do which we may or may not want to infinitely couple so let's just keep it seperate for now).
*/




main() {
	entity_manager = NewEntityManager('./grid_template.txt') <- init grid w/ template
	o = Orchestrator(entity_manager)
	v = Visualizer()
	for 1:
		grid = o.Step()
		v.draw(entity_manager)
}


Visualizer() {
	draw(entity_manager) {
        coord_to_entities = entity_manager.get_coord_to_entity
        for coord, entities in coord_to_entities:
            // ZCB COME BACK!
            // if any entity has the adder property, draw A

            // 0. Create a list of properties on this coord (can optimize later to pre-populate this)
            properties_on_coord = []
            for entity in entities:
                for property in entity.properties 
                    properties_on_coord.append(property)
            if properties_on_coord.contains('adder'):
                draw('A')
            elif properties_on_coord.contains('deleter'):
                draw('D')
            elif properties_on_coord.contains('movable'):
                draw('M')
            elif len(properties_on_coord) > 0:
                // at this point, we know it doesn't contain adder, deleter, or movable so it must only contain output properties
                num_output_properties = len(properties_on_coord)
                draw
            else:
                draw('.')
            // else if any entity has 

	 for each cell in the grid:
		 create a character to represent that cell based on the following rules:
			 - if the cell contains an "adder" property, draw it first (A)
			 - else if the cell contains a "deleter" property, draw it after (D)
			 - else if the cell contains a movable property, draw that
			 - else if the cell contains any other entity, draw it based on outputs:
				 - if the entity has 1 output, draw an arrow pointing that wayAll
				 - else if the entity has 2, 3, or 4 outputs, display the number 2,3 or 4
			 - else if the cell has no entities, show a period for the cell
	
	}
}

type Coord struct {
    Row int
    Col int
}

func (c *Coord) isAtRightOf(coord: Coord): bool {
    return coord.row == c.row && (c.row + 1 = coord.row)
}

func (c *Coord) isAtLeftOf(coord: Coord): bool {
    return coord.row == c.row && (c.row = coord.row + 1)
}

func (c *Coord) isAtTopOf(coord: Coord): bool {
    return (coord.row + 1 == c.row) && c.row = coord.row
}

func (c *Coord) isAtBottomOf(coord: Coord): bool {
    return (coord.row == c.row + 1) && c.row = coord.row
}


type Entity struct {
    Id string
    Properties []string
}

func (*Entity) NewEntity() {
    new_entity_id = GenerateRandomId()
    return &Entity{
        Id: new_entity_id
        Properties: []
    }
}


enum Property {
    - output_left, output_right, output_top, output_bottom, adder , deleter
}

/*
There are no restrictions on the grid in the entity manager.
It just exposes an API for the orchestrator to interact with the entities
and fetch certain information about the entities inside of the grid.
Natively, the grid allows any entities to be placed in any cells.
Validation must happen at the orchestrator level.
The entity manager just knows how to get/set different information on entities.
*/
EntityManager() {
    coord_to_entity_ids: map[coord]:[]entity_id // maps one coord to a list of entities at that coord
    entity_id_to_coord: map[entity_id]:Coord // only allow an entity to be in one cell for now
    entity_id_to_entity: map[entity_id]:Entity // Entity is an entity
    property_to_entity_ids: map[Property]:[]entity_id
    coord_to_entities: map[Coord]:[]Entities    // used only by the visualizer. maps coord to a full entity object

    // used only by the visualizer
    get_coord_to_entities() map[coord]:[]Entity {
        return coord_to_entity
    }

    // given a property, return all entityIDs with that property
    get_all_entity_ids_with_property(property: string): []entity_id {
        return property_to_entity_ids[property]
    }

    // given entity_id, return coord
    get_coord_of_entity_by_id(entity_id): Coord {
        return entity_id_to_coord[entity_id]
    }

    // return all entity IDs at coord
    get_entity_ids_at_coord(coord): []entity_id {
        return coord_to_entity_ids[coord]
    }

    // Remove all entities with ID of entity ID
    remove_entity(entity_id): void {
        // 0. Get the coord where the entity is
        coord = entity_id_to_coord[entity_id]
        // 0.5 Remove from coord_to_entity
        coord_to_entities[coord].remove(where=entity.Id=entity_id) // Remove WHERE entity.id == entity_id
        // 1. Go to that coord and remove the entity from that coord list (because we are only allowing one coord per entity, this works!)
        coord_to_entity_ids[coord].remove(entity_id) // or something like this
        // 2. Remove entity_id
        entity_id_to_coord[entity_id].delete() // or something like this
        // 3. Get all properties the entity was in
        entity = entity_id_to_entity[entity_id]
        // 4. Remove the entity ID from each property this entity_id is in
        for prop in entity.properties:
            property_to_entity_ids[prop].remove(entity_id)
        // 5. Remove the entity
        entity_id_to_entity[entity_id].delete() // or something like this

    }

    // returns true if entity with entity ID has property 'property'
    entity_has_property(entity_id, property: string): bool {
        entity_ids_with_prop = property_to_entity_ids[property]
        if entity_ids_with_prop:
            for entity_id_with_prop in entity_id_with_prop:
                if entity_id_with_prop == entity_id:
                    return true
        return false
    }

    get_all_properties_of_entity_with_id(entity_id: string): []string {
        entity = entity_id_to_entity[entity_id]
        return entity.properties
    }

    // adds an entity with props at cord
    add_entity_at_coord_with_props(props: []string, coord: Coord): void {
        // do in transaction:
        new_entity = NewEntity()
        coord_to_entity_ids[coord].append(new_entity.Id) // if exists, append, if not already exists, add + append
        entity_id_to_coord[new_entity.Id] = coord
        entity_id_to_entity[new_entity.Id] = new_entity
        coord_to_entities[coord].append(new_entity)

        // assign properties
        for property in props:
            property_to_entity_ids[prop].append(new_entity.Id) // if prop already exists in map then append, if not already exists, add + append
        
    }

    // returns all entity IDs in an indeterminant order
    get_all_entity_ids(): []entity_id {
        return entity_id_to_entity.keys()
    }

    // moves the entity ID to the coord
    move_entity_with_id_to_coord(entity_id, coord): void {
        // Do in transaction
        // 0. Remove the entity from the coord it is currently at
        coord_to_entity_ids[coord].remove(entity_id) // or something like this

        // 1. Overwrite the entity to the new coord
        entity_id_to_coord[entity_id] = coord
    }
}


Orchestrator:
	
	update() {
		// adders produce an entity if possible
		update_adders()
		
		// deleters delete entity if possible
        update_deleters()
		
	}

    update_deleters() {
        deleter_entity_ids = entity_manager.get_all_entity_ids_with_property('deleter')
        for deleter_entity_id in deleter_entity_ids:
            // A business rule is now that we should delete all other entities on this cell other than this deleter
            // 1. Get all entities on the same cell as the deleter
            coord_of_deleter = entity_manager.get_coord_of_entity_by_id(deleter_entity_id)
            entity_ids_on_cell_with_deleter = entity_manager.get_entity_ids_at_coord(coord_of_deleter)

            // 2. Remove all other entities on the same cell as the deleter, but not the deleter itself
            for entity_id_on_cell_with_deleter in entity_ids_on_cell_with_deleter and entity_id_on_cell_with_deleter != deleter_entity_id:
                entity_manager.remove_entity(entity_id_on_cell_with_deleter) // 
    }
	
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

			can_spawn = can_spawn_entity_at_coord_with_props(props_for_potential_new_entity, coord)
			// 2. actually spawn an entity at that coord
			entity_manager.add_entity_at_coord_with_props(props_for_potential_new_entity, coord)
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

⇦ ⇨ ⇧ ⇩
Example 1:
. ⇨ D ⇦ ⇦ ⇦ ⇦ ⇦ 
. ⇧ . . ⇧ . . ⇧
. ⇧ . . 2 ⇨ ⇨ ⇧ 
. 2 ⇦ ⇦ A . . ⇧
. ⇩ . . . . . ⇧
. ⇨ ⇨ ⇨ ⇨ ⇨ ⇨ ⇧

Example 2:
A ⇨ ⇨ ⇩ ⇦ 2 ⇦ A
⇨ ⇨ ⇨ ⇩ . ⇩ . .
⇧ . . ⇩ . ⇩ . .
2 ⇨ ⇨ D ⇦ 3 ⇨ ⇩
⇧ . . . . ⇩ . ⇩
⇧ ⇦ ⇦ ⇦ ⇦ ⇦ ⇦ ⇦


Example 3:
A ⇨ D
. ⇩ .
A ⇨ D
=
(adder, output_right),(output_right, output_bottom),(deleter)
(),(output_bottom),()
(adder),(output_right),(deleter)

(A>) (>v) (D)
()   (v)  ()
(A)  (>)  (d)

Would be parsed as:

input: raw text string
output: pre-loaded entity manager!

func () NewEntityManager(filepath: string): EntityManager {
    // 0. Create entity manager object
    entity_manager = EntityManager()
    // 1. Based on the contents of the file, use the API of entity_manager

    // 0. Get the file
    file_string = encode_to_string(filepath) // whatever the command for this is, I don't know
    // encoded_entity_list looks like ['A>','>v','D'...]
    encoded_entity_list = file_string.parse("()") // whatever the command for this is, it should parse the tokens between parens, taking into acocunt spaces between the parens
    // execute the commands on the entity_manager
    for encoded_entity in encoded_entity_list:
        properties = []
        for char in encoded_entity:
            if char == 'A':
                properties.append('adder')
            if char == 'D':
                properties.append('deleter')
            if char == 'M':
                properties.append('movable')

            if char == '>':
                properties.append('output_right')
            if char == '<':
                properties.append('output_left')
            if char == '^':
                properties.append('output_top')
            if char == 'v':
                properties.append('output_bottom')
        entity_manager.add_entity_at_coord_with_props(Coord(i,j), properties) // get i and j from the iterations of the loop
    return entity_manager
}