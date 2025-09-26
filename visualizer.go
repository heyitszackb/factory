package dreamer

type Visualizer struct {
	entityManager *EntityManager
}

func NewVisualizer(entityManager *EntityManager) *Visualizer {
	return &Visualizer{
		entityManager: entityManager,
	}
}

func (v *Visualizer) Draw() {
}



Visualizer() {
	draw(entity_manager) {
        coord_to_entities = entity_manager.get_coord_to_entity
        for coord, entities in coord_to_entities:

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

}