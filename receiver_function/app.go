/* app.go */

package main

import "fmt"

// Create a custom type called app
type app []string

// Receiver function for app type
// When a variable of type app is created, it has access to this print() method
func (a app) print() {
	for i, app := range a {
		fmt.Println(i, app)
	}
}
