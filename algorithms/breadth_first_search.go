package algorithms

// Graph — список смежности. map[string][]string.
type Graph map[string][]string

// BreadthFirstSearch — поиск в ширину от start до первой вершины,
// удовлетворяющей условию match.
// Очередь в Go: обычный слайс (queue = append(queue, x); x = queue[0]; queue = queue[1:]).
// Множество посещённых: map[string]struct{}.
//
// TODO: реализовать.
func BreadthFirstSearch(g Graph, start string, match func(string) bool) (string, bool) {
	panic("not implemented")
}
