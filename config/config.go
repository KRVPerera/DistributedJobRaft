package config

type Config struct {
	ID               int
	Peers            []string
	ElectionTimeout  int
	HeartbeatTimeout int
}
