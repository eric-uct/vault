package logical

import "time"

// SystemView exposes system configuration information in a safe way
// for logical backends to consume
type SystemView interface {
	// DefaultLeaseTTL returns the default lease TTL set in Vault configuration
	DefaultLeaseTTL() time.Duration

	// MaxLeaseTTL returns the max lease TTL set in Vault configuration; backend
	// authors should take care not to issue credentials that last longer than
	// this value, as Vault will revoke them
	MaxLeaseTTL() time.Duration

	// SudoPrivilege returns true if given policy name contains sudo policy
	// for the given path
	SudoPrivilege(path, policy string) bool
}

type StaticSystemView struct {
	DefaultLeaseTTLVal time.Duration
	MaxLeaseTTLVal     time.Duration
}

func (d StaticSystemView) DefaultLeaseTTL() time.Duration {
	return d.DefaultLeaseTTLVal
}

func (d StaticSystemView) MaxLeaseTTL() time.Duration {
	return d.MaxLeaseTTLVal
}

func (d StaticSystemView) SudoPrivilege(path, policy string) bool {
	return policy == "root"
}
