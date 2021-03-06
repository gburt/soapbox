// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// Environment represents a row from 'public.environments'.
type Environment struct {
	ID            int           `json:"id"`             // id
	ApplicationID sql.NullInt64 `json:"application_id"` // application_id
	Name          string        `json:"name"`           // name
	Slug          string        `json:"slug"`           // slug
	CreatedAt     time.Time     `json:"created_at"`     // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Environment exists in the database.
func (e *Environment) Exists() bool {
	return e._exists
}

// Deleted provides information if the Environment has been deleted from the database.
func (e *Environment) Deleted() bool {
	return e._deleted
}

// Insert inserts the Environment to the database.
func (e *Environment) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.environments (` +
		`application_id, name, slug, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, e.ApplicationID, e.Name, e.Slug, e.CreatedAt)
	err = db.QueryRow(sqlstr, e.ApplicationID, e.Name, e.Slug, e.CreatedAt).Scan(&e.ID)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Update updates the Environment in the database.
func (e *Environment) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if e._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.environments SET (` +
		`application_id, name, slug, created_at` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	XOLog(sqlstr, e.ApplicationID, e.Name, e.Slug, e.CreatedAt, e.ID)
	_, err = db.Exec(sqlstr, e.ApplicationID, e.Name, e.Slug, e.CreatedAt, e.ID)
	return err
}

// Save saves the Environment to the database.
func (e *Environment) Save(db XODB) error {
	if e.Exists() {
		return e.Update(db)
	}

	return e.Insert(db)
}

// Upsert performs an upsert for Environment.
//
// NOTE: PostgreSQL 9.5+ only
func (e *Environment) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if e._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.environments (` +
		`id, application_id, name, slug, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, application_id, name, slug, created_at` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.application_id, EXCLUDED.name, EXCLUDED.slug, EXCLUDED.created_at` +
		`)`

	// run query
	XOLog(sqlstr, e.ID, e.ApplicationID, e.Name, e.Slug, e.CreatedAt)
	_, err = db.Exec(sqlstr, e.ID, e.ApplicationID, e.Name, e.Slug, e.CreatedAt)
	if err != nil {
		return err
	}

	// set existence
	e._exists = true

	return nil
}

// Delete deletes the Environment from the database.
func (e *Environment) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !e._exists {
		return nil
	}

	// if deleted, bail
	if e._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.environments WHERE id = $1`

	// run query
	XOLog(sqlstr, e.ID)
	_, err = db.Exec(sqlstr, e.ID)
	if err != nil {
		return err
	}

	// set deleted
	e._deleted = true

	return nil
}

// Application returns the Application associated with the Environment's ApplicationID (application_id).
//
// Generated from foreign key 'environments_application_id_fkey'.
func (e *Environment) Application(db XODB) (*Application, error) {
	return ApplicationByID(db, int(e.ApplicationID.Int64))
}

// EnvironmentByApplicationIDName retrieves a row from 'public.environments' as a Environment.
//
// Generated from index 'environments_application_id_name_key'.
func EnvironmentByApplicationIDName(db XODB, applicationID sql.NullInt64, name string) (*Environment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, application_id, name, slug, created_at ` +
		`FROM public.environments ` +
		`WHERE application_id = $1 AND name = $2`

	// run query
	XOLog(sqlstr, applicationID, name)
	e := Environment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, applicationID, name).Scan(&e.ID, &e.ApplicationID, &e.Name, &e.Slug, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// EnvironmentByApplicationIDSlug retrieves a row from 'public.environments' as a Environment.
//
// Generated from index 'environments_application_id_slug_key'.
func EnvironmentByApplicationIDSlug(db XODB, applicationID sql.NullInt64, slug string) (*Environment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, application_id, name, slug, created_at ` +
		`FROM public.environments ` +
		`WHERE application_id = $1 AND slug = $2`

	// run query
	XOLog(sqlstr, applicationID, slug)
	e := Environment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, applicationID, slug).Scan(&e.ID, &e.ApplicationID, &e.Name, &e.Slug, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// EnvironmentByID retrieves a row from 'public.environments' as a Environment.
//
// Generated from index 'environments_pkey'.
func EnvironmentByID(db XODB, id int) (*Environment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, application_id, name, slug, created_at ` +
		`FROM public.environments ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	e := Environment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&e.ID, &e.ApplicationID, &e.Name, &e.Slug, &e.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
