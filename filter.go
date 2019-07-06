// Package hook provides wordpress like action and filter system for golang
package hook

import (
	"sort"
)

type Filter struct {
	ID       string // optional: unique id will generate random if not specified
	Tag      string // tag name
	Function func(args *map[string]interface{})
	Priority int
	// FunctionArgs map[string]interface{}
}

// Filters declares array of Filter
type Filters struct {
	List []Filter
}

// PrioritySorterF sorts filters by priority.
type PrioritySorterF []Filter

func (a PrioritySorterF) Len() int           { return len(a) }
func (a PrioritySorterF) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PrioritySorterF) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

// func (f *Filters) Hello() string {
// 	fmt.Println("Hello Filter")
// 	return "Hello Filter"
// }

// Add new filter struct to the Filters List
func (f *Filters) Add(tag string, funcToAdd func(args *map[string]interface{}), args ...map[string]interface{}) {

	// prevent panic: runtime error: index out of range
	var atts map[string]interface{}
	if len(args) != 0 {
		atts = args[0]
	}

	defaultArgs := map[string]interface{}{
		"id":       GenerateRandString(10),
		"priority": 10,
		// "functionArgs": make(map[string]interface{}),
	}

	for k, v := range atts {
		defaultArgs[k] = v
	}

	id := defaultArgs["id"].(string)
	priority := defaultArgs["priority"].(int)
	// functionArgs := defaultArgs["functionArgs"].(map[string]interface{})

	filter := Filter{
		ID:       id,
		Tag:      tag,
		Function: funcToAdd,
		Priority: priority,
	}
	f.List = append(f.List, filter)
	// fmt.Println(Filters)
}

// Execute the filter by specifying the tag name and args
func (f *Filters) Apply(tag string, args ...*map[string]interface{}) {
	// var atts *map[string]interface{}
	atts := make(map[string]interface{})
	atts1 := &atts // atts1 become *map[string]interface{}
	if len(args) != 0 {
		atts1 = (args[0]) // updates atts1 value from args[0] *map[string]interface{}
	}

	var filteredFilters []Filter

	// filter the Filters by tag
	for _, action := range f.List {
		// fmt.Println(action)
		if tag == action.Tag {
			filteredFilters = append(filteredFilters, action)
		}
	}

	// sort the filtered Filters by priority
	sort.Sort(PrioritySorterF(filteredFilters))
	// log.Println("by priority:", filteredFilters)

	for _, action := range filteredFilters {
		action.Function(atts1)
	}
}

// Remove the filter by ID
func (f *Filters) RemoveById(id string) {
	var filteredFilters []Filter
	for _, action := range f.List {
		if id != action.ID {
			filteredFilters = append(filteredFilters, action)
		}
	}
	f.List = filteredFilters
}

// Remove all the filters specified with particular tagname
func (f *Filters) Remove(tag string) {
	var filteredFilters []Filter
	// filter the Filters by tag
	for _, action := range f.List {
		// fmt.Println(action)
		if tag != action.Tag {
			filteredFilters = append(filteredFilters, action)
		}
	}

	f.List = filteredFilters
}

// Remove all filters
func (f *Filters) RemoveAll() {
	f.List = nil
}
