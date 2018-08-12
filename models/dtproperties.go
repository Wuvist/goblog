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

// Dtproperty is an object representing the database table.
type Dtproperty struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Objectid null.Int    `boil:"objectid" json:"objectid,omitempty" toml:"objectid" yaml:"objectid,omitempty"`
	Property string      `boil:"property" json:"property" toml:"property" yaml:"property"`
	Value    null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Uvalue   null.String `boil:"uvalue" json:"uvalue,omitempty" toml:"uvalue" yaml:"uvalue,omitempty"`
	Lvalue   null.Bytes  `boil:"lvalue" json:"lvalue,omitempty" toml:"lvalue" yaml:"lvalue,omitempty"`
	Version  int         `boil:"version" json:"version" toml:"version" yaml:"version"`

	R *dtpropertyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dtpropertyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DtpropertyColumns = struct {
	ID       string
	Objectid string
	Property string
	Value    string
	Uvalue   string
	Lvalue   string
	Version  string
}{
	ID:       "id",
	Objectid: "objectid",
	Property: "property",
	Value:    "value",
	Uvalue:   "uvalue",
	Lvalue:   "lvalue",
	Version:  "version",
}

// dtpropertyR is where relationships are stored.
type dtpropertyR struct {
}

// dtpropertyL is where Load methods for each relationship are stored.
type dtpropertyL struct{}

var (
	dtpropertyColumns               = []string{"id", "objectid", "property", "value", "uvalue", "lvalue", "version"}
	dtpropertyColumnsWithoutDefault = []string{"objectid", "property", "value", "uvalue", "lvalue"}
	dtpropertyColumnsWithDefault    = []string{"id", "version"}
	dtpropertyPrimaryKeyColumns     = []string{"id", "property"}
)

type (
	// DtpropertySlice is an alias for a slice of pointers to Dtproperty.
	// This should generally be used opposed to []Dtproperty.
	DtpropertySlice []*Dtproperty
	// DtpropertyHook is the signature for custom Dtproperty hook methods
	DtpropertyHook func(boil.Executor, *Dtproperty) error

	dtpropertyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dtpropertyType                 = reflect.TypeOf(&Dtproperty{})
	dtpropertyMapping              = queries.MakeStructMapping(dtpropertyType)
	dtpropertyPrimaryKeyMapping, _ = queries.BindMapping(dtpropertyType, dtpropertyMapping, dtpropertyPrimaryKeyColumns)
	dtpropertyInsertCacheMut       sync.RWMutex
	dtpropertyInsertCache          = make(map[string]insertCache)
	dtpropertyUpdateCacheMut       sync.RWMutex
	dtpropertyUpdateCache          = make(map[string]updateCache)
	dtpropertyUpsertCacheMut       sync.RWMutex
	dtpropertyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var dtpropertyBeforeInsertHooks []DtpropertyHook
var dtpropertyBeforeUpdateHooks []DtpropertyHook
var dtpropertyBeforeDeleteHooks []DtpropertyHook
var dtpropertyBeforeUpsertHooks []DtpropertyHook

var dtpropertyAfterInsertHooks []DtpropertyHook
var dtpropertyAfterSelectHooks []DtpropertyHook
var dtpropertyAfterUpdateHooks []DtpropertyHook
var dtpropertyAfterDeleteHooks []DtpropertyHook
var dtpropertyAfterUpsertHooks []DtpropertyHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Dtproperty) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Dtproperty) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Dtproperty) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Dtproperty) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Dtproperty) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Dtproperty) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Dtproperty) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Dtproperty) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Dtproperty) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dtpropertyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDtpropertyHook registers your hook function for all future operations.
func AddDtpropertyHook(hookPoint boil.HookPoint, dtpropertyHook DtpropertyHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dtpropertyBeforeInsertHooks = append(dtpropertyBeforeInsertHooks, dtpropertyHook)
	case boil.BeforeUpdateHook:
		dtpropertyBeforeUpdateHooks = append(dtpropertyBeforeUpdateHooks, dtpropertyHook)
	case boil.BeforeDeleteHook:
		dtpropertyBeforeDeleteHooks = append(dtpropertyBeforeDeleteHooks, dtpropertyHook)
	case boil.BeforeUpsertHook:
		dtpropertyBeforeUpsertHooks = append(dtpropertyBeforeUpsertHooks, dtpropertyHook)
	case boil.AfterInsertHook:
		dtpropertyAfterInsertHooks = append(dtpropertyAfterInsertHooks, dtpropertyHook)
	case boil.AfterSelectHook:
		dtpropertyAfterSelectHooks = append(dtpropertyAfterSelectHooks, dtpropertyHook)
	case boil.AfterUpdateHook:
		dtpropertyAfterUpdateHooks = append(dtpropertyAfterUpdateHooks, dtpropertyHook)
	case boil.AfterDeleteHook:
		dtpropertyAfterDeleteHooks = append(dtpropertyAfterDeleteHooks, dtpropertyHook)
	case boil.AfterUpsertHook:
		dtpropertyAfterUpsertHooks = append(dtpropertyAfterUpsertHooks, dtpropertyHook)
	}
}

// OneP returns a single dtproperty record from the query, and panics on error.
func (q dtpropertyQuery) OneP() *Dtproperty {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dtproperty record from the query.
func (q dtpropertyQuery) One() (*Dtproperty, error) {
	o := &Dtproperty{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dtproperties")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Dtproperty records from the query, and panics on error.
func (q dtpropertyQuery) AllP() DtpropertySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Dtproperty records from the query.
func (q dtpropertyQuery) All() (DtpropertySlice, error) {
	var o []*Dtproperty

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Dtproperty slice")
	}

	if len(dtpropertyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Dtproperty records in the query, and panics on error.
func (q dtpropertyQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Dtproperty records in the query.
func (q dtpropertyQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dtproperties rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dtpropertyQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dtpropertyQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dtproperties exists")
	}

	return count > 0, nil
}

// DtpropertiesG retrieves all records.
func DtpropertiesG(mods ...qm.QueryMod) dtpropertyQuery {
	return Dtproperties(boil.GetDB(), mods...)
}

// Dtproperties retrieves all the records using an executor.
func Dtproperties(exec boil.Executor, mods ...qm.QueryMod) dtpropertyQuery {
	mods = append(mods, qm.From("`dtproperties`"))
	return dtpropertyQuery{NewQuery(exec, mods...)}
}

// FindDtpropertyG retrieves a single record by ID.
func FindDtpropertyG(id int, property string, selectCols ...string) (*Dtproperty, error) {
	return FindDtproperty(boil.GetDB(), id, property, selectCols...)
}

// FindDtpropertyGP retrieves a single record by ID, and panics on error.
func FindDtpropertyGP(id int, property string, selectCols ...string) *Dtproperty {
	retobj, err := FindDtproperty(boil.GetDB(), id, property, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDtproperty retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDtproperty(exec boil.Executor, id int, property string, selectCols ...string) (*Dtproperty, error) {
	dtpropertyObj := &Dtproperty{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `dtproperties` where `id`=? AND `property`=?", sel,
	)

	q := queries.Raw(exec, query, id, property)

	err := q.Bind(dtpropertyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dtproperties")
	}

	return dtpropertyObj, nil
}

// FindDtpropertyP retrieves a single record by ID with an executor, and panics on error.
func FindDtpropertyP(exec boil.Executor, id int, property string, selectCols ...string) *Dtproperty {
	retobj, err := FindDtproperty(exec, id, property, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Dtproperty) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Dtproperty) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Dtproperty) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Dtproperty) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no dtproperties provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtpropertyColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dtpropertyInsertCacheMut.RLock()
	cache, cached := dtpropertyInsertCache[key]
	dtpropertyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dtpropertyColumns,
			dtpropertyColumnsWithDefault,
			dtpropertyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dtpropertyType, dtpropertyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dtpropertyType, dtpropertyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `dtproperties` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `dtproperties` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `dtproperties` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dtpropertyPrimaryKeyColumns))
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

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "models: unable to insert into dtproperties")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
		o.Property,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for dtproperties")
	}

CacheNoHooks:
	if !cached {
		dtpropertyInsertCacheMut.Lock()
		dtpropertyInsertCache[key] = cache
		dtpropertyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Dtproperty record. See Update for
// whitelist behavior description.
func (o *Dtproperty) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Dtproperty record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Dtproperty) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Dtproperty, and panics on error.
// See Update for whitelist behavior description.
func (o *Dtproperty) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Dtproperty.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Dtproperty) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	dtpropertyUpdateCacheMut.RLock()
	cache, cached := dtpropertyUpdateCache[key]
	dtpropertyUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			dtpropertyColumns,
			dtpropertyPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update dtproperties, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `dtproperties` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dtpropertyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dtpropertyType, dtpropertyMapping, append(wl, dtpropertyPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update dtproperties row")
	}

	if !cached {
		dtpropertyUpdateCacheMut.Lock()
		dtpropertyUpdateCache[key] = cache
		dtpropertyUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dtpropertyQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dtpropertyQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for dtproperties")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DtpropertySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DtpropertySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DtpropertySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DtpropertySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtpropertyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `dtproperties` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtpropertyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in dtproperty slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Dtproperty) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Dtproperty) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Dtproperty) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Dtproperty) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no dtproperties provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dtpropertyColumnsWithDefault, o)

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

	dtpropertyUpsertCacheMut.RLock()
	cache, cached := dtpropertyUpsertCache[key]
	dtpropertyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			dtpropertyColumns,
			dtpropertyColumnsWithDefault,
			dtpropertyColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			dtpropertyColumns,
			dtpropertyPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert dtproperties, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "dtproperties", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `dtproperties` WHERE `id`=? AND `property`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(dtpropertyType, dtpropertyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dtpropertyType, dtpropertyMapping, ret)
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

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for dtproperties")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
		o.Property,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for dtproperties")
	}

CacheNoHooks:
	if !cached {
		dtpropertyUpsertCacheMut.Lock()
		dtpropertyUpsertCache[key] = cache
		dtpropertyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Dtproperty record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dtproperty) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Dtproperty record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Dtproperty) DeleteG() error {
	if o == nil {
		return errors.New("models: no Dtproperty provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Dtproperty record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dtproperty) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Dtproperty record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Dtproperty) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Dtproperty provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dtpropertyPrimaryKeyMapping)
	sql := "DELETE FROM `dtproperties` WHERE `id`=? AND `property`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from dtproperties")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dtpropertyQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dtpropertyQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dtpropertyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dtproperties")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DtpropertySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DtpropertySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Dtproperty slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DtpropertySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DtpropertySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Dtproperty slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(dtpropertyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtpropertyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `dtproperties` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtpropertyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from dtproperty slice")
	}

	if len(dtpropertyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Dtproperty) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Dtproperty) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Dtproperty) ReloadG() error {
	if o == nil {
		return errors.New("models: no Dtproperty provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Dtproperty) Reload(exec boil.Executor) error {
	ret, err := FindDtproperty(exec, o.ID, o.Property)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DtpropertySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DtpropertySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DtpropertySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DtpropertySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DtpropertySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dtproperties := DtpropertySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dtpropertyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `dtproperties`.* FROM `dtproperties` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dtpropertyPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dtproperties)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DtpropertySlice")
	}

	*o = dtproperties

	return nil
}

// DtpropertyExists checks if the Dtproperty row exists.
func DtpropertyExists(exec boil.Executor, id int, property string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `dtproperties` where `id`=? AND `property`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id, property)
	}

	row := exec.QueryRow(sql, id, property)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dtproperties exists")
	}

	return exists, nil
}

// DtpropertyExistsG checks if the Dtproperty row exists.
func DtpropertyExistsG(id int, property string) (bool, error) {
	return DtpropertyExists(boil.GetDB(), id, property)
}

// DtpropertyExistsGP checks if the Dtproperty row exists. Panics on error.
func DtpropertyExistsGP(id int, property string) bool {
	e, err := DtpropertyExists(boil.GetDB(), id, property)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DtpropertyExistsP checks if the Dtproperty row exists. Panics on error.
func DtpropertyExistsP(exec boil.Executor, id int, property string) bool {
	e, err := DtpropertyExists(exec, id, property)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
