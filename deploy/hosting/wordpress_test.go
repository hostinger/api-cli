package hosting

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hostinger/api-cli/site"
)

func wordPressProject(t *testing.T) string {
	dir := t.TempDir()
	linkHere(t, dir)
	write(t, dir, "wp-config.php", "<?php")
	write(t, dir, "wp-content/themes/mine/style.css", "body{}")
	write(t, dir, "database.sql", "CREATE TABLE wp_posts (id int);")
	return dir
}

func TestWordPressSiteIsImportedWithItsDatabase(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	dir := wordPressProject(t)
	ui := &fakeUI{t: t}

	res, err := Run(context.Background(), api, ui, Options{Path: dir, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if res.Type != WordPressSite || res.Source != SourceUpload {
		t.Errorf("result = %+v", res)
	}
	// The dump goes next to the archive, under a name that cannot be guessed.
	if res.Files != 2 || !strings.HasPrefix(res.Database, "database_") || !strings.HasSuffix(res.Database, ".sql") {
		t.Errorf("files = %d, database = %q", res.Files, res.Database)
	}
	calls := strings.Join(api.calls, "\n")
	if !strings.Contains(calls, "import u1 example.com "+res.Archive+" "+res.Database) || strings.Contains(calls, "deploy ") {
		t.Errorf("calls:\n%s", calls)
	}
	// An empty website loses nothing: no warning.
	if strings.Contains(strings.Join(ui.infos, "\n"), "since the dump was exported") {
		t.Errorf("warned about an empty website: %v", ui.infos)
	}
	if !strings.Contains(res.Next, "few minutes") {
		t.Errorf("next = %q", res.Next)
	}
	if p, err := site.Load(dir); err != nil || p == nil || p.Type != string(WordPressSite) {
		t.Errorf("project file = %+v, %v", p, err)
	}
}

// A link is consent to redeploying files, not to replacing a live database.
func TestLinkedWordPressImportStillAsks(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.content = true
	dir := wordPressProject(t)
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", OrderID: 11, Type: string(WordPressSite)}); err != nil {
		t.Fatal(err)
	}

	plain := &fakeUI{t: t}
	_, err := Run(context.Background(), api, plain, Options{Path: dir})
	if err == nil || !strings.Contains(err.Error(), "database") || !strings.Contains(err.Error(), "--yes") {
		t.Errorf("err = %v", err)
	}
	// What is lost is spelled out, not only "replaces its database".
	if infos := strings.Join(plain.infos, "\n"); !strings.Contains(infos, "since the dump was exported") || !strings.Contains(infos, "not a way to ship theme or plugin changes") {
		t.Errorf("infos:\n%s", infos)
	}
	ui := &fakeUI{t: t, interactive: true, confirms: []bool{false}}
	if _, err := Run(context.Background(), api, ui, Options{Path: dir}); err == nil {
		t.Error("expected the declined import to fail")
	}
	if len(ui.asked) != 1 || !strings.Contains(ui.asked[0], "along with its database") {
		t.Errorf("asked %v", ui.asked)
	}
	if strings.Contains(strings.Join(api.calls, "\n"), "import ") {
		t.Errorf("imported without consent: %v", api.calls)
	}

	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Yes: true}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), "import u1 example.com") {
		t.Errorf("calls: %v", api.calls)
	}
}

func TestWordPressDumpIsFoundOrNamed(t *testing.T) {
	fastPolling(t)
	outside := filepath.Join(t.TempDir(), "backup.sql")
	write(t, filepath.Dir(outside), "backup.sql", "--")
	notSQL := filepath.Join(t.TempDir(), "backup.txt")
	write(t, filepath.Dir(notSQL), "backup.txt", "--")

	for name, tc := range map[string]struct {
		extra   []string // more dumps in the folder
		gz      string   // a gzipped dump in the folder
		remove  bool     // no database.sql in the folder
		sql     string
		want    string // the uploaded dump's name prefix
		wantErr string
	}{
		"several dumps":            {extra: []string{"old.sql"}, wantErr: "several database dumps"},
		"several dumps, one named": {extra: []string{"old.sql"}, sql: "old.sql", want: "old_"},
		"no dump":                  {remove: true, wantErr: "no database dump"},
		"dump outside the folder":  {remove: true, sql: outside, want: "backup_"},
		"missing dump":             {sql: "nope.sql", wantErr: "not found"},
		"not a .sql file":          {sql: notSQL, wantErr: "not a .sql or .sql.gz file"},
		"any case, with its .gz":   {remove: true, extra: []string{"DB.SQL"}, gz: "DB.sql.gz", want: "DB_"},
		"gzipped only":             {remove: true, gz: "backup.sql.gz", want: "backup_"},
	} {
		t.Run(name, func(t *testing.T) {
			api := newFake(t)
			dir := wordPressProject(t)
			for _, f := range tc.extra {
				write(t, dir, f, "--")
			}
			if tc.gz != "" {
				gzipFile(t, filepath.Join(dir, tc.gz), "CREATE TABLE wp_posts (id int);")
			}
			if tc.remove {
				if err := os.Remove(filepath.Join(dir, "database.sql")); err != nil {
					t.Fatal(err)
				}
			}
			sql := tc.sql
			if sql != "" && !filepath.IsAbs(sql) {
				sql = filepath.Join(dir, sql)
			}
			res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", SQL: sql})
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("err = %v", err)
				}
				if strings.Contains(strings.Join(api.calls, "\n"), "websites") {
					t.Errorf("the website was looked up before the dump was checked: %v", api.calls)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			// Only the site is packed: no dump lands in the document root. A .sql.gz
			// is uploaded decompressed.
			if res.Files != 2 || !strings.HasPrefix(res.Database, tc.want) || !strings.HasSuffix(strings.ToLower(res.Database), ".sql") {
				t.Errorf("files = %d, database = %q", res.Files, res.Database)
			}
			if !strings.Contains(strings.Join(api.calls, "\n"), "import u1 example.com "+res.Archive+" "+res.Database) {
				t.Errorf("calls: %v", api.calls)
			}
		})
	}
}

// Without a link only the site's folder is searched: a dump that happens to
// be in the current directory would replace the live database.
func TestUnlinkedImportIgnoresADumpInTheCurrentDirectory(t *testing.T) {
	fastPolling(t)
	cwd := t.TempDir()
	linkHere(t, cwd)
	write(t, cwd, "unrelated.sql", "CREATE TABLE other (id int);")
	write(t, cwd, "site/wp-config.php", "<?php")
	write(t, cwd, "site/wp-content/themes/mine/style.css", "body{}")
	api := newFake(t)
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: filepath.Join(cwd, "site"), Domain: "example.com", Yes: true})
	if err == nil || !strings.Contains(err.Error(), "no database dump") {
		t.Errorf("err = %v", err)
	}
	if calls := strings.Join(api.calls, "\n"); strings.Contains(calls, "creds") {
		t.Errorf("uploaded:\n%s", calls)
	}
}

func TestWordPressArchiveNeedsItsDump(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	tmp := t.TempDir()
	archive := filepath.Join(tmp, "site.zip")
	write(t, tmp, "site.zip", "PK")
	write(t, tmp, "db.sql", "--")

	if _, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: archive, Domain: "example.com", Type: WordPressSite}); err == nil || !strings.Contains(err.Error(), "--sql") {
		t.Errorf("err = %v", err)
	}
	res, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: archive, Domain: "example.com", Type: WordPressSite, SQL: filepath.Join(tmp, "db.sql")})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(api.calls, "\n"), "import u1 example.com "+res.Archive+" "+res.Database) {
		t.Errorf("calls: %v", api.calls)
	}
}

func TestWordPressImportRefusals(t *testing.T) {
	fastPolling(t)
	t.Run("onto a Node.js app", func(t *testing.T) {
		api := newFake(t)
		api.websites[0].WebsiteType = "nodejs"
		_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: wordPressProject(t), Domain: "example.com", Yes: true})
		if err == nil || !strings.Contains(err.Error(), "Node.js app") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("from Git", func(t *testing.T) {
		_, err := Run(context.Background(), newFake(t), &fakeUI{t: t}, Options{Path: wordPressProject(t), Domain: "example.com", Source: SourceGit})
		if err == nil || !strings.Contains(err.Error(), "--source git") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("onto a headless website", func(t *testing.T) {
		for _, typ := range []Type{"", WordPressSite} { // detected, and passed with --type
			api := newFake(t)
			api.websites[0].WebsiteType = "headless"
			dir := wordPressProject(t)
			write(t, dir, "package.json", `{"name":"frontend"}`)
			_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", Type: typ, Yes: true})
			if err == nil || !strings.Contains(err.Error(), "no WordPress import") {
				t.Errorf("type %q: err = %v", typ, err)
			}
			if calls := strings.Join(api.calls, "\n"); strings.Contains(calls, "creds") {
				t.Errorf("type %q uploaded:\n%s", typ, calls)
			}
		}
	})
	t.Run("--sql with a static site", func(t *testing.T) {
		_, err := Run(context.Background(), newFake(t), &fakeUI{t: t}, Options{Path: staticProject(t), Domain: "example.com", SQL: "db.sql"})
		if err == nil || !strings.Contains(err.Error(), "--sql is only") {
			t.Errorf("err = %v", err)
		}
	})
}

func TestWordPressDryRunSaysItOverwrites(t *testing.T) {
	api := newFake(t)
	api.content = true
	dir := wordPressProject(t)
	if err := site.Save(dir, &site.Project{Kind: "website", Domain: "example.com", Username: "u1", Type: string(WordPressSite)}); err != nil {
		t.Fatal(err)
	}
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), api, ui, Options{Path: dir, DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(strings.Join(ui.infos, "\n"), "since the dump was exported") {
		t.Errorf("the plan does not say what is lost: %v", ui.infos)
	}
	if !res.OverwritesContent || res.Database == "" || strings.Contains(strings.Join(api.calls, "\n"), "creds") {
		t.Errorf("result = %+v, calls = %v", res, api.calls)
	}
}

// A failed import leaves the dump on the server: the error says where.
func TestFailedWordPressImportNamesTheUploadedDump(t *testing.T) {
	fastPolling(t)
	api := newFake(t)
	api.importErr = errors.New("import refused")
	_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: wordPressProject(t), Domain: "example.com", Yes: true})
	if err == nil || !strings.Contains(err.Error(), "import refused") || !strings.Contains(err.Error(), "uploaded as database_") {
		t.Errorf("err = %v", err)
	}
}

func gzipFile(t *testing.T, path, content string) {
	t.Helper()
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	zw := gzip.NewWriter(f)
	if _, err := zw.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

// A project that keeps WordPress in a subfolder, with the export next to it
// (as a Hostinger backup and a docker-compose setup leave it).
func TestWordPressInASubfolderIsFoundWithTheDumpBesideIt(t *testing.T) {
	fastPolling(t)
	root := t.TempDir()
	linkHere(t, root)
	write(t, root, "docker-compose.yml", "services: {}")
	write(t, root, "db-init/dump.sql", "--")
	write(t, root, "wordpress/wp-config.php", "<?php")
	write(t, root, "wordpress/wp-content/themes/mine/style.css", "body{}")
	write(t, root, "u1_db.20261111.sql", "CREATE TABLE wp_posts (id int);")
	gzipFile(t, filepath.Join(root, "u1_db.20261111.sql.gz"), "CREATE TABLE wp_posts (id int);")
	api := newFake(t)
	ui := &fakeUI{t: t}

	res, err := Run(context.Background(), api, ui, Options{Path: root, Domain: "example.com"})
	if err != nil {
		t.Fatal(err)
	}
	// Only the WordPress folder is packed; the .sql is taken over its .gz copy.
	if res.Type != WordPressSite || res.Files != 2 || !strings.HasPrefix(res.Database, "u1_db.20261111_") || !strings.HasSuffix(res.Database, ".sql") {
		t.Errorf("type = %s, files = %d, database = %q", res.Type, res.Files, res.Database)
	}
	if infos := strings.Join(ui.infos, "\n"); !strings.Contains(infos, "The WordPress site is in wordpress/") {
		t.Errorf("infos:\n%s", infos)
	}
	p, err := site.Load(root)
	if err != nil || p == nil || p.Path != "wordpress" {
		t.Fatalf("project file = %+v, %v", p, err)
	}

	// A plain up from the root or from inside wordpress/ redeploys the linked
	// folder and still finds the dump beside it.
	for _, cwd := range []string{root, filepath.Join(root, "wordpress")} {
		linkHere(t, cwd)
		res, err = Run(context.Background(), newFake(t), &fakeUI{t: t}, Options{Path: ".", Yes: true})
		if err != nil || res.Type != WordPressSite || !strings.HasPrefix(res.Database, "u1_db.20261111_") {
			t.Errorf("from %s: %+v, %v", cwd, res, err)
		}
	}

	// Two WordPress sites: neither is guessed.
	write(t, root, "other/wp-config.php", "<?php")
	if sub := wordPressSubdir(root); sub != "" {
		t.Errorf("picked %s", sub)
	}
}

// A damaged or empty dump fails before anything is uploaded, in a dry run too:
// the API cannot delete an archive left behind by a later failure.
func TestBadDumpFailsBeforeAnyUpload(t *testing.T) {
	fastPolling(t)
	truncated := func(t *testing.T, dir string) {
		gzipFile(t, filepath.Join(dir, "backup.sql.gz"), strings.Repeat("INSERT INTO wp_posts VALUES (1);\n", 500))
		raw, _ := os.ReadFile(filepath.Join(dir, "backup.sql.gz"))
		write(t, dir, "backup.sql.gz", string(raw[:len(raw)/2]))
	}
	for name, tc := range map[string]struct {
		setup   func(t *testing.T, dir string)
		wantErr string
	}{
		"not gzip":       {func(t *testing.T, dir string) { write(t, dir, "backup.sql.gz", "plain text") }, "not a gzip file"},
		"truncated gzip": {truncated, "is damaged"},
		"empty gzip":     {func(t *testing.T, dir string) { gzipFile(t, filepath.Join(dir, "backup.sql.gz"), "") }, "is empty"},
		"empty .sql":     {func(t *testing.T, dir string) { write(t, dir, "backup.sql", "") }, "missing or empty"},
	} {
		for _, dryRun := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s, dry run %v", name, dryRun), func(t *testing.T) {
				dir := wordPressProject(t)
				if err := os.Remove(filepath.Join(dir, "database.sql")); err != nil {
					t.Fatal(err)
				}
				tc.setup(t, dir)
				api := newFake(t)
				_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: dryRun, Yes: true})
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("err = %v", err)
				}
				if calls := strings.Join(api.calls, "\n"); strings.Contains(calls, "creds") {
					t.Errorf("uploaded before the dump was checked:\n%s", calls)
				}
			})
		}
		t.Run(name+", --create", func(t *testing.T) {
			dir := wordPressProject(t)
			if err := os.Remove(filepath.Join(dir, "database.sql")); err != nil {
				t.Fatal(err)
			}
			tc.setup(t, dir)
			api := newFake(t)
			_, err := Run(context.Background(), api, &fakeUI{t: t}, Options{Path: dir, Domain: "new.example", Create: true, OrderID: 11, Type: WordPressSite, Yes: true})
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("err = %v", err)
			}
			if calls := strings.Join(api.calls, "\n"); strings.Contains(calls, "setup ") || strings.Contains(calls, "websites") {
				t.Errorf("a website was chosen before the dump was checked:\n%s", calls)
			}
		})
	}
}

// .gitignore still leaves out what the site does not need; dumps and Docker
// files are left out even when it does not list them.
func TestWordPressImportFollowsGitignoreOutsideTheSite(t *testing.T) {
	dir := wordPressProject(t)
	write(t, dir, ".gitignore", "wp-content/*\n!wp-content/themes/\n*.bak\n*.log\n.idea/\n")
	write(t, dir, "wp-content/uploads/a.jpg", "jpg")
	write(t, dir, "wp-content/plugins/shop/shop.php", "<?php")
	write(t, dir, "wp-content/plugins/shop/install.sql", "CREATE TABLE shop (id int);") // the plugin's code
	write(t, dir, "wp-content/plugins/shop/debug.log", "customer data")                 // *.log still applies
	write(t, dir, "wp-content/db.php", "<?php")
	write(t, dir, "wp-content/cache/page.html", "cached")
	write(t, dir, "wp-content/debug.log", "secrets")
	write(t, dir, "notes.bak", "old")
	write(t, dir, ".idea/workspace.xml", "<x/>")
	write(t, dir, "docker-compose.yml", "MYSQL_ROOT_PASSWORD: pw")
	write(t, dir, "wp-content/themes/mine/compose.yaml", "x")
	write(t, dir, "extra.sql.bak", "--") // left out by *.bak already
	write(t, dir, "old/dump.sql~", "--")
	write(t, dir, "mysql.php", "<?php")
	write(t, dir, ".env", "SECRET=1")
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), newFake(t), ui, Options{Path: dir, Domain: "example.com", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	// .gitignore, wp-config.php, mysql.php, the theme's style.css, the upload,
	// the plugin with its install.sql and the drop-in; not the cache, the
	// debug.logs, .bak, .idea, Docker files or dumps.
	if res.Files != 8 {
		t.Errorf("files = %d", res.Files)
	}
	// Scripts see what was left out, not only the terminal.
	if got := strings.Join(res.SkippedFiles, ","); got != ".env,docker-compose.yml,old/dump.sql~,wp-content/themes/mine/compose.yaml" {
		t.Errorf("skipped = %s", got)
	}
	infos := strings.Join(ui.infos, "\n")
	for _, want := range []string{
		"Packing what the site needs although .gitignore leaves it out: wp-content/db.php, wp-content/plugins/shop/install.sql, wp-content/plugins/shop/shop.php and 1 more",
		"database dumps would be served publicly): old/dump.sql~",
		"local development files, which often hold database passwords): docker-compose.yml, wp-content/themes/mine/compose.yaml",
	} {
		if !strings.Contains(infos, want) {
			t.Errorf("missing %q in:\n%s", want, infos)
		}
	}

	// .hostingerignore is written for deploys, so it applies in full.
	write(t, dir, IgnoreFile, "wp-content/uploads/\n")
	res, err = Run(context.Background(), newFake(t), &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	// The 8 above plus .hostingerignore and what only .gitignore left out
	// (cache, both debug.logs, .bak, .idea); still no upload, .env, dump or Docker file.
	if res.Files != 13 {
		t.Errorf("with %s: files = %d", IgnoreFile, res.Files)
	}
}

// A rule over all of a needed folder's contents (uploads/*) is overridden like
// one over the folder. In plugins and themes everything but logs and caches is
// packed: vendor/ and dist/ are code the site runs.
func TestWordPressImportRestoresAWholeFolderOnly(t *testing.T) {
	dir := wordPressProject(t)
	write(t, dir, ".gitignore", "wp-content/uploads/*\n!wp-content/uploads/.gitkeep\nwp-content/plugins/shop/cache/\nvendor/\ndist/\n")
	write(t, dir, "wp-content/plugins/shop/vendor/autoload.php", "<?php")
	write(t, dir, "wp-content/plugins/shop/vendor/psr/cache/src/CacheItemInterface.php", "<?php") // code, not a cache
	write(t, dir, "wp-content/themes/mine/dist/app.js", "x")
	write(t, dir, "vendor/tool.php", "<?php") // outside plugins and themes .gitignore applies
	write(t, dir, "wp-content/uploads/.gitkeep", "")
	write(t, dir, "wp-content/uploads/2026/a.jpg", "jpg")
	write(t, dir, "wp-content/plugins/shop/shop.php", "<?php")
	write(t, dir, "wp-content/plugins/shop/cache/x.html", "cached")
	res, err := Run(context.Background(), newFake(t), &fakeUI{t: t}, Options{Path: dir, Domain: "example.com", DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	// .gitignore, wp-config.php, style.css, .gitkeep, the upload, shop.php,
	// autoload.php, psr/cache and app.js; not the plugin's cache or the top vendor/.
	if res.Files != 9 {
		t.Errorf("files = %d", res.Files)
	}
}

// Each left-out file is reported once, with its own reason: dumps and Docker
// files are not secrets, and a needed symlink is not packed.
func TestWordPressImportReportsEachLeftOutFileOnce(t *testing.T) {
	fastPolling(t)
	dir := wordPressProject(t)
	write(t, dir, ".gitignore", "/wp-content/uploads\n")
	write(t, dir, "docker-compose.yml", "MYSQL_ROOT_PASSWORD: pw")
	write(t, dir, ".env", "SECRET=1")
	write(t, dir, "wp-content/themes/mine/.npmrc", "//registry.npmjs.org/:_authToken=x") // served like any file
	shared := t.TempDir()
	if err := os.Symlink(shared, filepath.Join(dir, "wp-content", "uploads")); err != nil {
		t.Fatal(err)
	}
	ui := &fakeUI{t: t}
	res, err := Run(context.Background(), newFake(t), ui, Options{Path: dir, Domain: "example.com", Yes: true})
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(res.SkippedFiles, ","); got != ".env,docker-compose.yml,wp-content/themes/mine/.npmrc" {
		t.Errorf("skipped = %s", got)
	}
	infos := strings.Join(ui.infos, "\n")
	for _, want := range []string{
		"Left out (secrets): .env, wp-content/themes/mine/.npmrc. ",
		"local development files, which often hold database passwords): docker-compose.yml",
		"Left out (symlinks are not followed): wp-content/uploads.",
	} {
		if !strings.Contains(infos, want) {
			t.Errorf("missing %q in:\n%s", want, infos)
		}
	}
	if strings.Contains(infos, "Packing what the site needs") {
		t.Errorf("the symlink was reported as packed:\n%s", infos)
	}
}

func TestDumpAndDevFileNames(t *testing.T) {
	for name, want := range map[string]bool{
		"db.sql": true, "DB.SQL.GZ": true, "x.sql.bak": true, "x.sql~": true, "x.sql.2024-01-01": true,
		"mysql.php": false, "sqlite.php": false, "sql": false, "notes.sqlx": false,
	} {
		if isDump(name) != want {
			t.Errorf("isDump(%q) = %v", name, !want)
		}
	}
	for name, want := range map[string]bool{
		"docker-compose.yml": true, "docker-compose.override.yaml": true, "compose.yaml": true, "Dockerfile": true,
		"Dockerfile.dev": true, ".dockerignore": true, "composer.json": false, "compose.php": false, "config.yml": false,
	} {
		if isDevFile(name) != want {
			t.Errorf("isDevFile(%q) = %v", name, !want)
		}
	}
}
