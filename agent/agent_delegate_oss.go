// +build !ent

package agent

// The OSS delegate does not add any extra functions
type delegate interface {
	commonDelegate
}
