// Package twofer is short for two for one. One for you and one for me.
package twofer

// ShareWith will return two-fer of a given name, whom to share. If the name is
// missing, the name will be re-assigned with "you"
func ShareWith(name string) string {
	// overwrite the name if it missing
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
