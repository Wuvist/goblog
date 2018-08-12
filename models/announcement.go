// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Announcement is an object representing the database table.
type Announcement struct {
	ID      int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Subject string    `boil:"Subject" json:"Subject" toml:"Subject" yaml:"Subject"`
	Content string    `boil:"Content" json:"Content" toml:"Content" yaml:"Content"`
	Addtime time.Time `boil:"addtime" json:"addtime" toml:"addtime" yaml:"addtime"`

	R *announcementR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L announcementL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnnouncementColumns = struct {
	ID      string
	Subject string
	Content string
	Addtime string
}{
	ID:      "id",
	Subject: "Subject",
	Content: "Content",
	Addtime: "addtime",
}

// announcementR is where relationships are stored.
type announcementR struct {
}

// announcementL is where Load methods for each relationship are stored.
type announcementL struct{}

var (
	announcementColumns               = []string{"id", "Subject", "Content", "addtime"}
	announcementColumnsWithoutDefault = []string{"Subject", "Content"}
	announcementColumnsWithDefault    = []string{"id", "addtime"}
	announcementPrimaryKeyColumns     = []string{"id"}
)

type (
	// AnnouncementSlice is an alias for a slice of pointers to Announcement.
	// This should generally be used opposed to []Announcement.
	AnnouncementSlice []*Announcement
	// AnnouncementHook is the signature for custom Announcement hook methods
	AnnouncementHook func(boil.Executor, *Announcement) error

	announcementQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	announcementType                 = reflect.TypeOf(&Announcement{})
	announcementMapping              = queries.MakeStructMapping(announcementType)
	announcementPrimaryKeyMapping, _ = queries.BindMapping(announcementType, announcementMapping, announcementPrimaryKeyColumns)
	announcementInsertCacheMut       sync.RWMutex
	announcementInsertCache          = make(map[string]insertCache)
	announcementUpdateCacheMut       sync.RWMutex
	announcementUpdateCache          = make(map[string]updateCache)
	announcementUpsertCacheMut       sync.RWMutex
	announcementUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var announcementBeforeInsertHooks []AnnouncementHook
var announcementBeforeUpdateHooks []AnnouncementHook
var announcementBeforeDeleteHooks []AnnouncementHook
var announcementBeforeUpsertHooks []AnnouncementHook

var announcementAfterInsertHooks []AnnouncementHook
var announcementAfterSelectHooks []AnnouncementHook
var announcementAfterUpdateHooks []AnnouncementHook
var announcementAfterDeleteHooks []AnnouncementHook
var announcementAfterUpsertHooks []AnnouncementHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Announcement) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Announcement) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Announcement) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Announcement) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Announcement) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Announcement) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Announcement) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Announcement) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Announcement) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range announcementAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnnouncementHook registers your hook function for all future operations.
func AddAnnouncementHook(hookPoint boil.HookPoint, announcementHook AnnouncementHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		announcementBeforeInsertHooks = append(announcementBeforeInsertHooks, announcementHook)
	case boil.BeforeUpdateHook:
		announcementBeforeUpdateHooks = append(announcementBeforeUpdateHooks, announcementHook)
	case boil.BeforeDeleteHook:
		announcementBeforeDeleteHooks = append(announcementBeforeDeleteHooks, announcementHook)
	case boil.BeforeUpsertHook:
		announcementBeforeUpsertHooks = append(announcementBeforeUpsertHooks, announcementHook)
	case boil.AfterInsertHook:
		announcementAfterInsertHooks = append(announcementAfterInsertHooks, announcementHook)
	case boil.AfterSelectHook:
		announcementAfterSelectHooks = append(announcementAfterSelectHooks, announcementHook)
	case boil.AfterUpdateHook:
		announcementAfterUpdateHooks = append(announcementAfterUpdateHooks, announcementHook)
	case boil.AfterDeleteHook:
		announcementAfterDeleteHooks = append(announcementAfterDeleteHooks, announcementHook)
	case boil.AfterUpsertHook:
		announcementAfterUpsertHooks = append(announcementAfterUpsertHooks, announcementHook)
	}
}

// OneP returns a single announcement record from the query, and panics on error.
func (q announcementQuery) OneP() *Announcement {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single announcement record from the query.
func (q announcementQuery) One() (*Announcement, error) {
	o := &Announcement{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for announcement")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Announcement records from the query, and panics on error.
func (q announcementQuery) AllP() AnnouncementSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Announcement records from the query.
func (q announcementQuery) All() (AnnouncementSlice, error) {
	var o []*Announcement

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Announcement slice")
	}

	if len(announcementAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Announcement records in the query, and panics on error.
func (q announcementQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Announcement records in the query.
func (q announcementQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count announcement rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q announcementQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q announcementQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if announcement exists")
	}

	return count > 0, nil
}

// AnnouncementsG retrieves all records.
func AnnouncementsG(mods ...qm.QueryMod) announcementQuery {
	return Announcements(boil.GetDB(), mods...)
}

// Announcements retrieves all the records using an executor.
func Announcements(exec boil.Executor, mods ...qm.QueryMod) announcementQuery {
	mods = append(mods, qm.From("`announcement`"))
	return announcementQuery{NewQuery(exec, mods...)}
}

// FindAnnouncementG retrieves a single record by ID.
func FindAnnouncementG(id int, selectCols ...string) (*Announcement, error) {
	return FindAnnouncement(boil.GetDB(), id, selectCols...)
}

// FindAnnouncementGP retrieves a single record by ID, and panics on error.
func FindAnnouncementGP(id int, selectCols ...string) *Announcement {
	retobj, err := FindAnnouncement(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAnnouncement retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnnouncement(exec boil.Executor, id int, selectCols ...string) (*Announcement, error) {
	announcementObj := &Announcement{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `announcement` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(announcementObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from announcement")
	}

	return announcementObj, nil
}

// FindAnnouncementP retrieves a single record by ID with an executor, and panics on error.
func FindAnnouncementP(exec boil.Executor, id int, selectCols ...string) *Announcement {
	retobj, err := FindAnnouncement(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Announcement) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Announcement) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Announcement) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Announcement) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no announcement provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(announcementColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	announcementInsertCacheMut.RLock()
	cache, cached := announcementInsertCache[key]
	announcementInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			announcementColumns,
			announcementColumnsWithDefault,
			announcementColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(announcementType, announcementMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(announcementType, announcementMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `announcement` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `announcement` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `announcement` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, announcementPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into announcement")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == announcementMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for announcement")
	}

CacheNoHooks:
	if !cached {
		announcementInsertCacheMut.Lock()
		announcementInsertCache[key] = cache
		announcementInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Announcement record. See Update for
// whitelist behavior description.
func (o *Announcement) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Announcement record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Announcement) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Announcement, and panics on error.
// See Update for whitelist behavior description.
func (o *Announcement) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Announcement.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Announcement) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	announcementUpdateCacheMut.RLock()
	cache, cached := announcementUpdateCache[key]
	announcementUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			announcementColumns,
			announcementPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update announcement, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `announcement` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, announcementPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(announcementType, announcementMapping, append(wl, announcementPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update announcement row")
	}

	if !cached {
		announcementUpdateCacheMut.Lock()
		announcementUpdateCache[key] = cache
		announcementUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q announcementQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q announcementQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for announcement")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AnnouncementSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AnnouncementSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AnnouncementSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnnouncementSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), announcementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `announcement` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, announcementPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in announcement slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Announcement) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Announcement) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Announcement) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Announcement) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no announcement provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(announcementColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	announcementUpsertCacheMut.RLock()
	cache, cached := announcementUpsertCache[key]
	announcementUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			announcementColumns,
			announcementColumnsWithDefault,
			announcementColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			announcementColumns,
			announcementPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert announcement, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "announcement", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `announcement` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(announcementType, announcementMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(announcementType, announcementMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for announcement")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == announcementMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for announcement")
	}

CacheNoHooks:
	if !cached {
		announcementUpsertCacheMut.Lock()
		announcementUpsertCache[key] = cache
		announcementUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Announcement record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Announcement) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Announcement record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Announcement) DeleteG() error {
	if o == nil {
		return errors.New("models: no Announcement provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Announcement record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Announcement) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Announcement record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Announcement) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Announcement provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), announcementPrimaryKeyMapping)
	sql := "DELETE FROM `announcement` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from announcement")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q announcementQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q announcementQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no announcementQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from announcement")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AnnouncementSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AnnouncementSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Announcement slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AnnouncementSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnnouncementSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Announcement slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(announcementBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), announcementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `announcement` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, announcementPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from announcement slice")
	}

	if len(announcementAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Announcement) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Announcement) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Announcement) ReloadG() error {
	if o == nil {
		return errors.New("models: no Announcement provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Announcement) Reload(exec boil.Executor) error {
	ret, err := FindAnnouncement(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnnouncementSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnnouncementSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnnouncementSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AnnouncementSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnnouncementSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	announcements := AnnouncementSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), announcementPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `announcement`.* FROM `announcement` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, announcementPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&announcements)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnnouncementSlice")
	}

	*o = announcements

	return nil
}

// AnnouncementExists checks if the Announcement row exists.
func AnnouncementExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `announcement` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if announcement exists")
	}

	return exists, nil
}

// AnnouncementExistsG checks if the Announcement row exists.
func AnnouncementExistsG(id int) (bool, error) {
	return AnnouncementExists(boil.GetDB(), id)
}

// AnnouncementExistsGP checks if the Announcement row exists. Panics on error.
func AnnouncementExistsGP(id int) bool {
	e, err := AnnouncementExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AnnouncementExistsP checks if the Announcement row exists. Panics on error.
func AnnouncementExistsP(exec boil.Executor, id int) bool {
	e, err := AnnouncementExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
