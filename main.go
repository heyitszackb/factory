package main

import "time"

func main() {
	entityManager := NewEntityManager("./grid_template.txt")
	o := NewOrchestrator(entityManager)
	v := NewVisualizer(entityManager)
	step := 0

	for {
		o.Step()
		v.Draw(step)
		v.DrawDebug(step)
		step++
		time.Sleep(200 * time.Millisecond)
	}
}
