package profile

import g "github.com/scottstamp/goearth"

// Args contains the event arguments for profile events.
type Args struct {
	Profile Profile
}

// Updated registers an event handler that is invoked when the user's profile is updated.
func (mgr *Manager) Updated(handler g.EventHandler[Args]) {
	mgr.updated.Register(handler)
}
