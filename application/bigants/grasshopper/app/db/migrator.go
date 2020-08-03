package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v4"
)

func NewMigrator(db beginner) *Migrator {
	return &Migrator{db, nil}
}

type Migrator struct {
	db  beginner
	Log func(fmt string, a ...interface{})
}

// MigrateTo 함수는 주어진 버전까지의 마이그레이션을 모두 적용한다.
func (m *Migrator) MigrateTo(ctx context.Context, sqlDir string, to string) (err error) {

	return withTx(ctx, m.db, func(tx pgx.Tx) (err error) {
		var from string
		if from, err = m.CurrentVersion(ctx); err != nil {
			return err
		}
		mgs, err := m.MigrationsFromTo(sqlDir, from, to)
		if err != nil {
			return err
		}

		m.log("migrate: found %d migrations to exec from %s to %s", len(mgs), from, to)

		for _, mg := range mgs {
			if _, err = tx.Exec(ctx, mg.SQL()); err != nil {
				m.log("migrate: %s_%s FAILED\n", mg.Version(), mg.Name())
				return err
			}
			if mg.isUp {
				m.log("migrate: %s_%s UP\n", mg.Version(), mg.Name())
			} else {
				m.log("migrate: %s_%s DOWN\n", mg.Version(), mg.Name())
			}
		}
		return nil
	})
}

func (m *Migrator) MigrateToLatestVersion(ctx context.Context, sqlDir string) (err error) {
	to, err := m.LatestVersion(sqlDir)
	if err != nil {
		return err
	}
	return m.MigrateTo(ctx, sqlDir, to)
}

func (m *Migrator) CurrentVersion(ctx context.Context) (currentVersion string, err error) {
	err = withTx(ctx, m.db, func(tx pgx.Tx) error {
		_, err = tx.Exec(ctx, initialMigrationSQL)
		if err != nil {
			return err
		}
		err = tx.QueryRow(ctx,
			`SELECT version FROM gh.migration ORDER BY version DESC LIMIT 1`,
		).Scan(&currentVersion)
		if err != nil && err != pgx.ErrNoRows {
			return err
		}
		return nil
	})
	return currentVersion, err
}

func (m *Migrator) LatestVersion(sqlDir string) (latest string, err error) {
	files, err := findMigrations(sqlDir)
	if err != nil {
		return latest, err
	}
	for _, fp := range files {
		version, _, _ := parseFilepath(fp)
		if version > latest {
			latest = version
		}
	}
	return latest, nil
}

// MigrationsFromTo returns all Migrations which is required migrate database schema from `from` to `to`
func (m *Migrator) MigrationsFromTo(sqlDir string, from, to string) (mgs []*Migration, err error) {
	var lower, upper string
	var isUp bool

	if from < to {
		isUp, lower, upper = true, from, to
	} else {
		isUp, lower, upper = false, to, from
	}

	files, err := findMigrations(sqlDir)
	if err != nil {
		return nil, err
	}

	for _, fp := range files {
		mg := Migration{}
		mg.version, mg.name, mg.isUp = parseFilepath(fp)
		if lower < mg.version && mg.version <= upper && mg.isUp == isUp {
			mg.sql, err = readAndFormatSQL(fp)
			if err != nil {
				return nil, err
			}
			mgs = append(mgs, &mg)
		}
	}
	sort.Slice(mgs, func(i, j int) bool {
		if isUp {
			return mgs[i].version < mgs[j].version
		}
		return mgs[i].version > mgs[j].version
	})

	return mgs, nil
}

func (m *Migrator) log(fmt string, a ...interface{}) {
	if m.Log != nil {
		m.Log(fmt, a...)
	}
}

func findMigrations(sqlDir string) ([]string, error) {
	return filepath.Glob(filepath.Join(string(sqlDir), "*.sql"))
}

func parseFilepath(fp string) (version, name string, isUp bool) {
	f := filepath.Base(fp)
	ext := filepath.Ext(f) // .sql
	f = strings.TrimSuffix(f, ext)
	upOrDown := filepath.Ext(f) // .up or .down
	versionAndName := strings.Split(strings.TrimSuffix(f, upOrDown), "_")
	return versionAndName[0], versionAndName[1], upOrDown != ".down"
}

func readAndFormatSQL(fp string) (sql string, err error) {
	version, _, isUp := parseFilepath(fp)

	b, err := ioutil.ReadFile(fp)
	if err != nil {
		return sql, err
	}
	if isUp {
		sql = fmt.Sprintf(
			"DO $__migration__$ BEGIN PERFORM migrate_up(\n'%s'\n); %s END $__migration__$;",
			version, string(b),
		)
	} else {
		sql = fmt.Sprintf(
			"DO $__migration__$ BEGIN PERFORM migrate_down(\n'%s'\n); %s END $__migration__$;",
			version, string(b),
		)
	}
	return sql, nil
}

type Migration struct {
	version string
	name    string
	isUp    bool
	sql     string
}

func (m *Migration) Name() string      { return m.name }
func (m *Migration) Version() string   { return m.version }
func (m *Migration) SQL() (sql string) { return m.sql }

var initialMigrationSQL = `
	DO $A$ BEGIN

	CREATE SCHEMA IF NOT EXISTS gh;

	CREATE TABLE IF NOT EXISTS gh.migration (
		"version" TEXT PRIMARY KEY,
		"created" TIMESTAMPTZ DEFAULT NOW()
	);

	CREATE OR REPLACE FUNCTION migrate_up(text) returns void AS $$
	BEGIN
		INSERT INTO gh.migration VALUES ($1);
	END;
	$$ LANGUAGE plpgsql;

	CREATE OR REPLACE FUNCTION migrate_down(text) returns void AS $$
	BEGIN
		IF (SELECT EXISTS(SELECT version FROM gh.migration WHERE version = $1)) THEN
			DELETE FROM gh.migration WHERE version = $1;
		ELSE
			RAISE EXCEPTION 'Cannot drop gh.migration %. it does not exists in database', $1;
		END IF;
	END;
	$$ LANGUAGE plpgsql;

	END $A$;
`
