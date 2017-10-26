// Package mmmdatastructures implements common data structures
// that are hard-coded to particular Go types rather than
// using the empty interface and casting.
//
// This means that there is a lot of repetition between
// implementations; no effort has been made to DRY up the code.
//
// On the other hand, if you need precisely one of these
// data structrues, there's no extraneous fluff/indirection/abstraction.
//
// Also, because of the simplistic nature of these data structures,
// they may be good starting places for copying and customizing
// for your own projects.
package mmmdatastructures
