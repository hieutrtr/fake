package fake

// AdState generates random ad's state
func AdState() string {
	return lookup(lang, "states", false)
}
