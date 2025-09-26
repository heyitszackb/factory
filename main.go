package dreamer

import "time"

func main() {
	entityManager := NewEntityManager("./grid_template.txt")
	o := NewOrchestrator(entityManager)
	v := NewVisualizer(entityManager)

	for {
		o.Step()
		v.Draw()
		time.Sleep(200 * time.Millisecond)
	}
}
