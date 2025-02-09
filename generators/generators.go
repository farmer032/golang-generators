package generators

type SliceGenerator struct {
	currentIndex int
	slice        *[]int
}

func NewSliceGenerator(slice *[]int) *SliceGenerator {
	return &SliceGenerator{currentIndex: 0, slice: slice}
}

func (iter *SliceGenerator) HasNext() bool {
	return iter.currentIndex < len(*iter.slice)
}

func (iter *SliceGenerator) Next() int {
	slice := *iter.slice
	value := slice[iter.currentIndex]
	iter.currentIndex++
	return value
}

type MapGenerator struct {
	CurrentIndex int
	Map          *map[string]string
}

type GeneratorEntry struct {
	Key   string
	Value string
}

func NewMapGenerator(m *map[string]string) *MapGenerator {
	return &MapGenerator{CurrentIndex: 0, Map: m}
}

func (iter *MapGenerator) HasNext() bool {
	return iter.CurrentIndex < len(*iter.Map)
}

func (iter *MapGenerator) Next() (string, string) {

	channel := make(chan GeneratorEntry)

	go func() {
		defer close(channel)

		for key, value := range *iter.Map {
			channel <- GeneratorEntry{Key: key, Value: value}
			iter.CurrentIndex++
		}
	}()

	entry := <-channel

	return entry.Key, entry.Value
}
