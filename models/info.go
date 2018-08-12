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
	"gopkg.in/volatiletech/null.v6"
)

// Info is an object representing the database table.
type Info struct {
	ID    int16       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Count int         `boil:"count" json:"count" toml:"count" yaml:"count"`
	Notes null.String `boil:"notes" json:"notes,omitempty" toml:"notes" yaml:"notes,omitempty"`

	R *infoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L infoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InfoColumns = struct {
	ID    string
	Count string
	Notes string
}{
	ID:    "id",
	Count: "count",
	Notes: "notes",
}

// infoR is where relationships are stored.
type infoR struct {
}

// infoL is where Load methods for each relationship are stored.
type infoL struct{}

var (
	infoColumns               = []string{"id", "count", "notes"}
	infoColumnsWithoutDefault = []string{"notes"}
	infoColumnsWithDefault    = []string{"id", "count"}
	infoPrimaryKeyColumns     = []string{"id"}
)

type (
	// InfoSlice is an alias for a slice of pointers to Info.
	// This should generally be used opposed to []Info.
	InfoSlice []*Info
	// InfoHook is the signature for custom Info hook methods
	InfoHook func(boil.Executor, *Info) error

	infoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	infoType                 = reflect.TypeOf(&Info{})
	infoMapping              = queries.MakeStructMapping(infoType)
	infoPrimaryKeyMapping, _ = queries.BindMapping(infoType, infoMapping, infoPrimaryKeyColumns)
	infoInsertCacheMut       sync.RWMutex
	infoInsertCache          = make(map[string]insertCache)
	infoUpdateCacheMut       sync.RWMutex
	infoUpdateCache          = make(map[string]updateCache)
	infoUpsertCacheMut       sync.RWMutex
	infoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var infoBeforeInsertHooks []InfoHook
var infoBeforeUpdateHooks []InfoHook
var infoBeforeDeleteHooks []InfoHook
var infoBeforeUpsertHooks []InfoHook

var infoAfterInsertHooks []InfoHook
var infoAfterSelectHooks []InfoHook
var infoAfterUpdateHooks []InfoHook
var infoAfterDeleteHooks []InfoHook
var infoAfterUpsertHooks []InfoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Info) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range infoBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Info) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range infoBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Info) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range infoBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Info) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range infoBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Info) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range infoAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Info) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range infoAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Info) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range infoAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Info) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range infoAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Info) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range infoAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInfoHook registers your hook function for all future operations.
func AddInfoHook(hookPoint boil.HookPoint, infoHook InfoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		infoBeforeInsertHooks = append(infoBeforeInsertHooks, infoHook)
	case boil.BeforeUpdateHook:
		infoBeforeUpdateHooks = append(infoBeforeUpdateHooks, infoHook)
	case boil.BeforeDeleteHook:
		infoBeforeDeleteHooks = append(infoBeforeDeleteHooks, infoHook)
	case boil.BeforeUpsertHook:
		infoBeforeUpsertHooks = append(infoBeforeUpsertHooks, infoHook)
	case boil.AfterInsertHook:
		infoAfterInsertHooks = append(infoAfterInsertHooks, infoHook)
	case boil.AfterSelectHook:
		infoAfterSelectHooks = append(infoAfterSelectHooks, infoHook)
	case boil.AfterUpdateHook:
		infoAfterUpdateHooks = append(infoAfterUpdateHooks, infoHook)
	case boil.AfterDeleteHook:
		infoAfterDeleteHooks = append(infoAfterDeleteHooks, infoHook)
	case boil.AfterUpsertHook:
		infoAfterUpsertHooks = append(infoAfterUpsertHooks, infoHook)
	}
}

// OneP returns a single info record from the query, and panics on error.
func (q infoQuery) OneP() *Info {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single info record from the query.
func (q infoQuery) One() (*Info, error) {
	o := &Info{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for info")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Info records from the query, and panics on error.
func (q infoQuery) AllP() InfoSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Info records from the query.
func (q infoQuery) All() (InfoSlice, error) {
	var o []*Info

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Info slice")
	}

	if len(infoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Info records in the query, and panics on error.
func (q infoQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Info records in the query.
func (q infoQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count info rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q infoQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q infoQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if info exists")
	}

	return count > 0, nil
}

// InfosG retrieves all records.
func InfosG(mods ...qm.QueryMod) infoQuery {
	return Infos(boil.GetDB(), mods...)
}

// Infos retrieves all the records using an executor.
func Infos(exec boil.Executor, mods ...qm.QueryMod) infoQuery {
	mods = append(mods, qm.From("`info`"))
	return infoQuery{NewQuery(exec, mods...)}
}

// FindInfoG retrieves a single record by ID.
func FindInfoG(id int16, selectCols ...string) (*Info, error) {
	return FindInfo(boil.GetDB(), id, selectCols...)
}

// FindInfoGP retrieves a single record by ID, and panics on error.
func FindInfoGP(id int16, selectCols ...string) *Info {
	retobj, err := FindInfo(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindInfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInfo(exec boil.Executor, id int16, selectCols ...string) (*Info, error) {
	infoObj := &Info{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `info` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(infoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from info")
	}

	return infoObj, nil
}

// FindInfoP retrieves a single record by ID with an executor, and panics on error.
func FindInfoP(exec boil.Executor, id int16, selectCols ...string) *Info {
	retobj, err := FindInfo(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Info) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Info) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Info) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Info) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no info provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(infoColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	infoInsertCacheMut.RLock()
	cache, cached := infoInsertCache[key]
	infoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			infoColumns,
			infoColumnsWithDefault,
			infoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(infoType, infoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(infoType, infoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `info` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `info` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `info` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, infoPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into info")
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

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == infoMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for info")
	}

CacheNoHooks:
	if !cached {
		infoInsertCacheMut.Lock()
		infoInsertCache[key] = cache
		infoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Info record. See Update for
// whitelist behavior description.
func (o *Info) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Info record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Info) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Info, and panics on error.
// See Update for whitelist behavior description.
func (o *Info) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Info.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Info) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	infoUpdateCacheMut.RLock()
	cache, cached := infoUpdateCache[key]
	infoUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			infoColumns,
			infoPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update info, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `info` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, infoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(infoType, infoMapping, append(wl, infoPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update info row")
	}

	if !cached {
		infoUpdateCacheMut.Lock()
		infoUpdateCache[key] = cache
		infoUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q infoQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q infoQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for info")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o InfoSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o InfoSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o InfoSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InfoSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), infoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `info` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, infoPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in info slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Info) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Info) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Info) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Info) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no info provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(infoColumnsWithDefault, o)

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

	infoUpsertCacheMut.RLock()
	cache, cached := infoUpsertCache[key]
	infoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			infoColumns,
			infoColumnsWithDefault,
			infoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			infoColumns,
			infoPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert info, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "info", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `info` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(infoType, infoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(infoType, infoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for info")
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

	o.ID = int16(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == infoMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for info")
	}

CacheNoHooks:
	if !cached {
		infoUpsertCacheMut.Lock()
		infoUpsertCache[key] = cache
		infoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Info record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Info) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Info record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Info) DeleteG() error {
	if o == nil {
		return errors.New("models: no Info provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Info record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Info) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Info record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Info) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Info provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), infoPrimaryKeyMapping)
	sql := "DELETE FROM `info` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from info")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q infoQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q infoQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no infoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from info")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o InfoSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o InfoSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Info slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o InfoSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InfoSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Info slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(infoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), infoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `info` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, infoPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from info slice")
	}

	if len(infoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Info) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Info) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Info) ReloadG() error {
	if o == nil {
		return errors.New("models: no Info provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Info) Reload(exec boil.Executor) error {
	ret, err := FindInfo(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *InfoSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *InfoSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InfoSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty InfoSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InfoSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	infos := InfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), infoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `info`.* FROM `info` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, infoPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&infos)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InfoSlice")
	}

	*o = infos

	return nil
}

// InfoExists checks if the Info row exists.
func InfoExists(exec boil.Executor, id int16) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `info` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if info exists")
	}

	return exists, nil
}

// InfoExistsG checks if the Info row exists.
func InfoExistsG(id int16) (bool, error) {
	return InfoExists(boil.GetDB(), id)
}

// InfoExistsGP checks if the Info row exists. Panics on error.
func InfoExistsGP(id int16) bool {
	e, err := InfoExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// InfoExistsP checks if the Info row exists. Panics on error.
func InfoExistsP(exec boil.Executor, id int16) bool {
	e, err := InfoExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}