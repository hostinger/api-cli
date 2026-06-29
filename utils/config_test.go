package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func write(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestMigrateLegacyConfig_NewExists_NoMigration(t *testing.T) {
	home := t.TempDir()
	newPath := filepath.Join(home, ".hostinger.yaml")
	oldPath := filepath.Join(home, ".hapi.yaml")
	write(t, newPath, "api_token: new\n")
	write(t, oldPath, "api_token: old\n")

	var buf bytes.Buffer
	got, migrated := MigrateLegacyConfig(home, &buf)

	if got != newPath {
		t.Errorf("path = %q, want %q", got, newPath)
	}
	if migrated {
		t.Error("migrated = true, want false")
	}
	if _, err := os.Stat(oldPath); err != nil {
		t.Error("old file must be left untouched when new exists")
	}
	if buf.Len() != 0 {
		t.Errorf("unexpected notice: %q", buf.String())
	}
}

func TestMigrateLegacyConfig_OldOnly_Migrates(t *testing.T) {
	home := t.TempDir()
	newPath := filepath.Join(home, ".hostinger.yaml")
	oldPath := filepath.Join(home, ".hapi.yaml")
	write(t, oldPath, "api_token: old\n")

	var buf bytes.Buffer
	got, migrated := MigrateLegacyConfig(home, &buf)

	if got != newPath {
		t.Errorf("path = %q, want %q", got, newPath)
	}
	if !migrated {
		t.Error("migrated = false, want true")
	}
	if _, err := os.Stat(oldPath); !os.IsNotExist(err) {
		t.Error("old file should have been moved")
	}
	if b, _ := os.ReadFile(newPath); string(b) != "api_token: old\n" {
		t.Errorf("new file content = %q", string(b))
	}
	if !strings.Contains(buf.String(), "migrated") {
		t.Errorf("expected migration notice, got %q", buf.String())
	}
}

func TestMigrateLegacyConfig_NeitherExists_NoOp(t *testing.T) {
	home := t.TempDir()
	newPath := filepath.Join(home, ".hostinger.yaml")

	var buf bytes.Buffer
	got, migrated := MigrateLegacyConfig(home, &buf)

	if got != newPath {
		t.Errorf("path = %q, want %q", got, newPath)
	}
	if migrated {
		t.Error("migrated = true, want false")
	}
	if buf.Len() != 0 {
		t.Errorf("unexpected notice: %q", buf.String())
	}
}

func TestLegacyEnvWarning(t *testing.T) {
	if got := legacyEnvWarning("", "tok", "HOSTINGER_API_TOKEN", "HAPI_API_TOKEN"); !strings.Contains(got, "HAPI_API_TOKEN") || !strings.Contains(got, "HOSTINGER_API_TOKEN") {
		t.Errorf("expected deprecation warning naming both vars, got %q", got)
	}
	if got := legacyEnvWarning("new", "old", "A", "B"); got != "" {
		t.Errorf("new set => no warning, got %q", got)
	}
	if got := legacyEnvWarning("", "", "A", "B"); got != "" {
		t.Errorf("neither set => no warning, got %q", got)
	}
}

func TestBindLegacyEnv_PrimaryEnv(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)

	t.Setenv("HOSTINGER_API_TOKEN", "primary")
	t.Setenv("HAPI_API_TOKEN", "")

	var buf bytes.Buffer
	BindLegacyEnv(&buf)

	if got := viper.GetString("api_token"); got != "primary" {
		t.Errorf("api_token = %q, want %q", got, "primary")
	}
	if strings.Contains(buf.String(), "api_token") {
		t.Errorf("expected no deprecation notice for api_token, got %q", buf.String())
	}
}

func TestBindLegacyEnv_LegacyEnvFallback(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)

	t.Setenv("HOSTINGER_API_URL", "")
	t.Setenv("HAPI_API_URL", "legacy")

	var buf bytes.Buffer
	BindLegacyEnv(&buf)

	if got := viper.GetString("api_url"); got != "legacy" {
		t.Errorf("api_url = %q, want %q", got, "legacy")
	}
	notice := buf.String()
	if !strings.Contains(notice, "HAPI_API_URL") {
		t.Errorf("expected deprecation notice to name HAPI_API_URL, got %q", notice)
	}
	if !strings.Contains(notice, "HOSTINGER_API_URL") {
		t.Errorf("expected deprecation notice to name HOSTINGER_API_URL, got %q", notice)
	}
}

func TestMigrateLegacyConfig_RenameFails_FallsBackToOld(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("rename cannot be blocked by perms as root")
	}

	home := t.TempDir()
	oldPath := filepath.Join(home, ".hapi.yaml")
	newPath := filepath.Join(home, ".hostinger.yaml")
	write(t, oldPath, "api_token: old\n")

	// Make home non-writable so rename fails (can't create new file).
	if err := os.Chmod(home, 0o500); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chmod(home, 0o700) })

	var buf bytes.Buffer
	got, migrated := MigrateLegacyConfig(home, &buf)

	if got != oldPath {
		t.Errorf("path = %q, want %q", got, oldPath)
	}
	if !migrated {
		t.Error("migrated = false, want true (legacy fallback in effect)")
	}
	if _, err := os.Stat(newPath); !os.IsNotExist(err) {
		t.Error("new file should not have been created")
	}
	notice := buf.String()
	if !strings.Contains(notice, "warning:") {
		t.Errorf("expected warning notice, got %q", notice)
	}
}
