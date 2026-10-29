package cmd

import (
	"os"

	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

var unlinkGlobal bool

var UnlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Forget the website this directory is linked to",
	Long: `Remove the .hostinger directory that links this directory to a website, so
commands ask again or need a domain. With --global, remove default_domain from the
config file instead.`,
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if unlinkGlobal {
			path := configPath()
			if !configHasKey(path, "default_domain") {
				if env := utils.EnvPrefix + "DEFAULT_DOMAIN"; os.Getenv(env) != "" {
					return printResult(cmd, "default_domain comes from "+env+", not "+path+": unset it there", map[string]any{"unlinked": false, "env": env})
				}
				return printResult(cmd, "No default website is set in "+path, map[string]any{"unlinked": false})
			}
			removed, err := writeConfigKey(path, "default_domain", "")
			if err != nil {
				return err
			}
			text := "Default website removed from " + path
			if removed {
				text = "Default website removed; " + path + " deleted as it held nothing else"
			}
			return printResult(cmd, text, map[string]any{"unlinked": true, "configFile": path})
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		p, linkedDir, err := site.Find(dir)
		if err != nil {
			// unlink is how a broken link is removed, so don't require reading it.
			broken := brokenLink(dir)
			if broken == "" {
				return err
			}
			if err := site.Remove(broken); err != nil {
				return err
			}
			return printResult(cmd, "Removed the unreadable link in "+broken, map[string]any{"unlinked": true, "linkFile": site.Path(broken)})
		}
		if p == nil {
			return printResult(cmd, "This directory is not linked to a website", map[string]any{"unlinked": false})
		}
		if err := site.Remove(linkedDir); err != nil {
			return err
		}
		return printResult(cmd, "Unlinked "+linkedDir+" from "+p.Domain, map[string]any{"unlinked": true, "domain": p.Domain, "linkFile": site.Path(linkedDir)})
	},
}

// configHasKey reports whether the config file itself sets key; viper also
// answers from the environment.
func configHasKey(path, key string) bool {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	m := map[string]any{}
	if yaml.Unmarshal(raw, &m) != nil {
		return false
	}
	_, ok := m[key]
	return ok
}

// brokenLink is the nearest directory from start up whose link cannot be read.
func brokenLink(start string) string {
	dirs, _ := site.LinkDirs(start)
	for _, dir := range dirs {
		if _, err := site.Load(dir); err != nil {
			return dir
		}
	}
	return ""
}

func init() {
	UnlinkCmd.Flags().BoolVar(&unlinkGlobal, "global", false, "Remove default_domain from the config file instead")
	UnlinkCmd.GroupID = groupDeploy
	RootCmd.AddCommand(UnlinkCmd)
}
