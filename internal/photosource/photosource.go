// Package photosource defines the PhotoSource abstraction (ADR-0002).
//
// All Memorie code that needs photo data interacts via this interface —
// never directly with Immich's API or filesystem. This keeps the door
// open for future providers (PhotoPrism, raw filesystem) without
// rippling changes through every layer.
package photosource

import "context"

type PhotoSource interface {
	Health(ctx context.Context) error
}
