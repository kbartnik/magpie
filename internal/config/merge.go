package config

import "maps"

// Merge combines a global and local Config. Non-zero local fields override
// global; zero-value local fields fall through to global.
func Merge(global, local *Config) *Config {
	if global == nil {
		global = &Config{}
	}
	if local == nil {
		local = &Config{}
	}

	out := *global

	if local.DefaultVault != "" {
		out.DefaultVault = local.DefaultVault
	}
	if local.InboxPath != "" {
		out.InboxPath = local.InboxPath
	}
	if local.ArchivePath != "" {
		out.ArchivePath = local.ArchivePath
	}
	if local.LogPath != "" {
		out.LogPath = local.LogPath
	}

	if global.Plugins != nil || local.Plugins != nil {
		out.Plugins = make(map[string]string, len(global.Plugins)+len(local.Plugins))
		maps.Copy(out.Plugins, global.Plugins)
		maps.Copy(out.Plugins, local.Plugins)
	}

	return &out
}
