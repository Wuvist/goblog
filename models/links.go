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

// Link is an object representing the database table.
type Link struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Blogger int    `boil:"blogger" json:"blogger" toml:"blogger" yaml:"blogger"`
	Link    string `boil:"link" json:"link" toml:"link" yaml:"link"`
	URL     string `boil:"URL" json:"URL" toml:"URL" yaml:"URL"`
	Reveal  int8   `boil:"reveal" json:"reveal" toml:"reveal" yaml:"reveal"`

	R *linkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L linkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LinkColumns = struct {
	ID      string
	Blogger string
	Link    string
	URL     string
	Reveal  string
}{
	ID:      "id",
	Blogger: "blogger",
	Link:    "link",
	URL:     "URL",
	Reveal:  "reveal",
}

// linkR is where relationships are stored.
type linkR struct {
}

// linkL is where Load methods for each relationship are stored.
type linkL struct{}

var (
	linkColumns               = []string{"id", "blogger", "link", "URL", "reveal"}
	linkColumnsWithoutDefault = []string{"blogger", "link", "URL"}
	linkColumnsWithDefault    = []string{"id", "reveal"}
	linkPrimaryKeyColumns     = []string{"id"}
)

type (
	// LinkSlice is an alias for a slice of pointers to Link.
	// This should generally be used opposed to []Link.
	LinkSlice []*Link
	// LinkHook is the signature for custom Link hook methods
	LinkHook func(boil.Executor, *Link) error

	linkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	linkType                 = reflect.TypeOf(&Link{})
	linkMapping              = queries.MakeStructMapping(linkType)
	linkPrimaryKeyMapping, _ = queries.BindMapping(linkType, linkMapping, linkPrimaryKeyColumns)
	linkInsertCacheMut       sync.RWMutex
	linkInsertCache          = make(map[string]insertCache)
	linkUpdateCacheMut       sync.RWMutex
	linkUpdateCache          = make(map[string]updateCache)
	linkUpsertCacheMut       sync.RWMutex
	linkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var linkBeforeInsertHooks []LinkHook
var linkBeforeUpdateHooks []LinkHook
var linkBeforeDeleteHooks []LinkHook
var linkBeforeUpsertHooks []LinkHook

var linkAfterInsertHooks []LinkHook
var linkAfterSelectHooks []LinkHook
var linkAfterUpdateHooks []LinkHook
var linkAfterDeleteHooks []LinkHook
var linkAfterUpsertHooks []LinkHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Link) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range linkBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Link) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range linkBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Link) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range linkBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Link) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range linkBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Link) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range linkAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Link) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range linkAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Link) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range linkAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Link) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range linkAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Link) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range linkAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLinkHook registers your hook function for all future operations.
func AddLinkHook(hookPoint boil.HookPoint, linkHook LinkHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		linkBeforeInsertHooks = append(linkBeforeInsertHooks, linkHook)
	case boil.BeforeUpdateHook:
		linkBeforeUpdateHooks = append(linkBeforeUpdateHooks, linkHook)
	case boil.BeforeDeleteHook:
		linkBeforeDeleteHooks = append(linkBeforeDeleteHooks, linkHook)
	case boil.BeforeUpsertHook:
		linkBeforeUpsertHooks = append(linkBeforeUpsertHooks, linkHook)
	case boil.AfterInsertHook:
		linkAfterInsertHooks = append(linkAfterInsertHooks, linkHook)
	case boil.AfterSelectHook:
		linkAfterSelectHooks = append(linkAfterSelectHooks, linkHook)
	case boil.AfterUpdateHook:
		linkAfterUpdateHooks = append(linkAfterUpdateHooks, linkHook)
	case boil.AfterDeleteHook:
		linkAfterDeleteHooks = append(linkAfterDeleteHooks, linkHook)
	case boil.AfterUpsertHook:
		linkAfterUpsertHooks = append(linkAfterUpsertHooks, linkHook)
	}
}

// OneP returns a single link record from the query, and panics on error.
func (q linkQuery) OneP() *Link {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single link record from the query.
func (q linkQuery) One() (*Link, error) {
	o := &Link{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for links")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Link records from the query, and panics on error.
func (q linkQuery) AllP() LinkSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Link records from the query.
func (q linkQuery) All() (LinkSlice, error) {
	var o []*Link

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Link slice")
	}

	if len(linkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Link records in the query, and panics on error.
func (q linkQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Link records in the query.
func (q linkQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count links rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q linkQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q linkQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if links exists")
	}

	return count > 0, nil
}

// LinksG retrieves all records.
func LinksG(mods ...qm.QueryMod) linkQuery {
	return Links(boil.GetDB(), mods...)
}

// Links retrieves all the records using an executor.
func Links(exec boil.Executor, mods ...qm.QueryMod) linkQuery {
	mods = append(mods, qm.From("`links`"))
	return linkQuery{NewQuery(exec, mods...)}
}

// FindLinkG retrieves a single record by ID.
func FindLinkG(id int, selectCols ...string) (*Link, error) {
	return FindLink(boil.GetDB(), id, selectCols...)
}

// FindLinkGP retrieves a single record by ID, and panics on error.
func FindLinkGP(id int, selectCols ...string) *Link {
	retobj, err := FindLink(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLink retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLink(exec boil.Executor, id int, selectCols ...string) (*Link, error) {
	linkObj := &Link{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `links` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(linkObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from links")
	}

	return linkObj, nil
}

// FindLinkP retrieves a single record by ID with an executor, and panics on error.
func FindLinkP(exec boil.Executor, id int, selectCols ...string) *Link {
	retobj, err := FindLink(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Link) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Link) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Link) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Link) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no links provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	linkInsertCacheMut.RLock()
	cache, cached := linkInsertCache[key]
	linkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			linkColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(linkType, linkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `links` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `links` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `links` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, linkPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into links")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == linkMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for links")
	}

CacheNoHooks:
	if !cached {
		linkInsertCacheMut.Lock()
		linkInsertCache[key] = cache
		linkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Link record. See Update for
// whitelist behavior description.
func (o *Link) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Link record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Link) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Link, and panics on error.
// See Update for whitelist behavior description.
func (o *Link) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Link.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Link) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	linkUpdateCacheMut.RLock()
	cache, cached := linkUpdateCache[key]
	linkUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			linkColumns,
			linkPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update links, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `links` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, linkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, append(wl, linkPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update links row")
	}

	if !cached {
		linkUpdateCacheMut.Lock()
		linkUpdateCache[key] = cache
		linkUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q linkQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q linkQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for links")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LinkSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o LinkSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LinkSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LinkSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `links` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in link slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Link) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Link) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Link) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Link) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no links provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)

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

	linkUpsertCacheMut.RLock()
	cache, cached := linkUpsertCache[key]
	linkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			linkColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			linkColumns,
			linkPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert links, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "links", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `links` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(linkType, linkMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for links")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == linkMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for links")
	}

CacheNoHooks:
	if !cached {
		linkUpsertCacheMut.Lock()
		linkUpsertCache[key] = cache
		linkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Link record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Link) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Link record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Link) DeleteG() error {
	if o == nil {
		return errors.New("models: no Link provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Link record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Link) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Link record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Link) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Link provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), linkPrimaryKeyMapping)
	sql := "DELETE FROM `links` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from links")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q linkQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q linkQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no linkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from links")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o LinkSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o LinkSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Link slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LinkSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LinkSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Link slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(linkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from link slice")
	}

	if len(linkAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Link) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Link) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Link) ReloadG() error {
	if o == nil {
		return errors.New("models: no Link provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Link) Reload(exec boil.Executor) error {
	ret, err := FindLink(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LinkSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LinkSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LinkSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty LinkSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LinkSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	links := LinkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `links`.* FROM `links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, linkPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&links)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LinkSlice")
	}

	*o = links

	return nil
}

// LinkExists checks if the Link row exists.
func LinkExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `links` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if links exists")
	}

	return exists, nil
}

// LinkExistsG checks if the Link row exists.
func LinkExistsG(id int) (bool, error) {
	return LinkExists(boil.GetDB(), id)
}

// LinkExistsGP checks if the Link row exists. Panics on error.
func LinkExistsGP(id int) bool {
	e, err := LinkExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LinkExistsP checks if the Link row exists. Panics on error.
func LinkExistsP(exec boil.Executor, id int) bool {
	e, err := LinkExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}