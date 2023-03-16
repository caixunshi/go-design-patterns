package prototype

import "sync"

type ShapeCache struct {
	cache map[ShapeType]Shape
}

var (
	shapeCache *ShapeCache
	once       sync.Once
)

type ShapeType string

const (
	rectangle ShapeType = "rectangle"
	circle              = "circle"
	square              = "square"
)

func GetShapeCache() *ShapeCache {
	if shapeCache == nil {
		once.Do(func() {
			cache := make(map[ShapeType]Shape)
			cache[rectangle] = &Rectangle{
				resource: []interface{}{},
			}
			cache[circle] = &Circle{
				resource: []interface{}{},
			}
			cache[square] = &Square{
				resource: []interface{}{},
			}
			shapeCache = &ShapeCache{
				cache: cache,
			}

		})
	}
	return shapeCache
}

func (s *ShapeCache) getShape(shapeType ShapeType) Shape {
	return s.cache[shapeType].clone()
}
