package hook

import (
	"log"
	"sort"
)

// Action : Array of actions
type Action struct {
	ID       string
	Tag      string
	Function func(args map[string]interface{})
	Priority int
	// FunctionArgs map[string]interface{}
}

// PrioritySorter sorts actions by priority.
type PrioritySorter []Action

func (a PrioritySorter) Len() int           { return len(a) }
func (a PrioritySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PrioritySorter) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

// https://github.com/knesklab/hook/blob/master/class/Action.js

var Actions []Action

func Add(tag string, funcToAdd func(args map[string]interface{}), args ...map[string]interface{}) {

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

	action := Action{
		ID:       id,
		Tag:      tag,
		Function: funcToAdd,
		Priority: priority,
	}
	Actions = append(Actions, action)
	// fmt.Println(Actions)
}

func Do(tag string, args ...map[string]interface{}) {
	var atts map[string]interface{}
	if len(args) != 0 {
		atts = args[0]
	}

	// fmt.Println(Actions)
	var filteredActions []Action

	// filter the actions by tag
	for _, action := range Actions {
		// fmt.Println(action)
		if tag == action.Tag {
			filteredActions = append(filteredActions, action)
		}
	}

	// sort the filtered actions by priority
	sort.Sort(PrioritySorter(filteredActions))
	log.Println("by priority:", filteredActions)

	for _, action := range filteredActions {
		action.Function(atts)
	}

}

func RemoveById(id string) {
	var filteredActions []Action
	for _, action := range Actions {
		if id != action.ID {
			filteredActions = append(filteredActions, action)
		}
	}
	Actions = filteredActions
}

func Remove(tag string) {
	var filteredActions []Action
	// filter the actions by tag
	for _, action := range Actions {
		// fmt.Println(action)
		if tag != action.Tag {
			filteredActions = append(filteredActions, action)
		}
	}

	Actions = filteredActions
}

func RemoveAll() {
	Actions = nil
}
