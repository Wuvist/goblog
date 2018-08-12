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

// Archive is an object representing the database table.
type Archive struct {
	Index   int       `boil:"index" json:"index" toml:"index" yaml:"index"`
	Title   string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Author  int       `boil:"author" json:"author" toml:"author" yaml:"author"`
	Content string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	Emot    null.Int  `boil:"emot" json:"emot,omitempty" toml:"emot" yaml:"emot,omitempty"`
	Ref     null.Int  `boil:"ref" json:"ref,omitempty" toml:"ref" yaml:"ref,omitempty"`
	AddDate time.Time `boil:"add_date" json:"add_date" toml:"add_date" yaml:"add_date"`

	R *archiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L archiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArchiveColumns = struct {
	Index   string
	Title   string
	Author  string
	Content string
	Emot    string
	Ref     string
	AddDate string
}{
	Index:   "index",
	Title:   "title",
	Author:  "author",
	Content: "content",
	Emot:    "emot",
	Ref:     "ref",
	AddDate: "add_date",
}

// archiveR is where relationships are stored.
type archiveR struct {
}

// archiveL is where Load methods for each relationship are stored.
type archiveL struct{}

var (
	archiveColumns               = []string{"index", "title", "author", "content", "emot", "ref", "add_date"}
	archiveColumnsWithoutDefault = []string{"title", "author", "content", "emot", "ref", "add_date"}
	archiveColumnsWithDefault    = []string{"index"}
	archivePrimaryKeyColumns     = []string{"index"}
)

type (
	// ArchiveSlice is an alias for a slice of pointers to Archive.
	// This should generally be used opposed to []Archive.
	ArchiveSlice []*Archive
	// ArchiveHook is the signature for custom Archive hook methods
	ArchiveHook func(boil.Executor, *Archive) error

	archiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	archiveType                 = reflect.TypeOf(&Archive{})
	archiveMapping              = queries.MakeStructMapping(archiveType)
	archivePrimaryKeyMapping, _ = queries.BindMapping(archiveType, archiveMapping, archivePrimaryKeyColumns)
	archiveInsertCacheMut       sync.RWMutex
	archiveInsertCache          = make(map[string]insertCache)
	archiveUpdateCacheMut       sync.RWMutex
	archiveUpdateCache          = make(map[string]updateCache)
	archiveUpsertCacheMut       sync.RWMutex
	archiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var archiveBeforeInsertHooks []ArchiveHook
var archiveBeforeUpdateHooks []ArchiveHook
var archiveBeforeDeleteHooks []ArchiveHook
var archiveBeforeUpsertHooks []ArchiveHook

var archiveAfterInsertHooks []ArchiveHook
var archiveAfterSelectHooks []ArchiveHook
var archiveAfterUpdateHooks []ArchiveHook
var archiveAfterDeleteHooks []ArchiveHook
var archiveAfterUpsertHooks []ArchiveHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Archive) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Archive) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Archive) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Archive) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Archive) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Archive) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Archive) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Archive) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Archive) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range archiveAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArchiveHook registers your hook function for all future operations.
func AddArchiveHook(hookPoint boil.HookPoint, archiveHook ArchiveHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		archiveBeforeInsertHooks = append(archiveBeforeInsertHooks, archiveHook)
	case boil.BeforeUpdateHook:
		archiveBeforeUpdateHooks = append(archiveBeforeUpdateHooks, archiveHook)
	case boil.BeforeDeleteHook:
		archiveBeforeDeleteHooks = append(archiveBeforeDeleteHooks, archiveHook)
	case boil.BeforeUpsertHook:
		archiveBeforeUpsertHooks = append(archiveBeforeUpsertHooks, archiveHook)
	case boil.AfterInsertHook:
		archiveAfterInsertHooks = append(archiveAfterInsertHooks, archiveHook)
	case boil.AfterSelectHook:
		archiveAfterSelectHooks = append(archiveAfterSelectHooks, archiveHook)
	case boil.AfterUpdateHook:
		archiveAfterUpdateHooks = append(archiveAfterUpdateHooks, archiveHook)
	case boil.AfterDeleteHook:
		archiveAfterDeleteHooks = append(archiveAfterDeleteHooks, archiveHook)
	case boil.AfterUpsertHook:
		archiveAfterUpsertHooks = append(archiveAfterUpsertHooks, archiveHook)
	}
}

// OneP returns a single archive record from the query, and panics on error.
func (q archiveQuery) OneP() *Archive {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single archive record from the query.
func (q archiveQuery) One() (*Archive, error) {
	o := &Archive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for archives")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Archive records from the query, and panics on error.
func (q archiveQuery) AllP() ArchiveSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Archive records from the query.
func (q archiveQuery) All() (ArchiveSlice, error) {
	var o []*Archive

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Archive slice")
	}

	if len(archiveAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Archive records in the query, and panics on error.
func (q archiveQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Archive records in the query.
func (q archiveQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count archives rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q archiveQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q archiveQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if archives exists")
	}

	return count > 0, nil
}

// ArchivesG retrieves all records.
func ArchivesG(mods ...qm.QueryMod) archiveQuery {
	return Archives(boil.GetDB(), mods...)
}

// Archives retrieves all the records using an executor.
func Archives(exec boil.Executor, mods ...qm.QueryMod) archiveQuery {
	mods = append(mods, qm.From("`archives`"))
	return archiveQuery{NewQuery(exec, mods...)}
}

// FindArchiveG retrieves a single record by ID.
func FindArchiveG(index int, selectCols ...string) (*Archive, error) {
	return FindArchive(boil.GetDB(), index, selectCols...)
}

// FindArchiveGP retrieves a single record by ID, and panics on error.
func FindArchiveGP(index int, selectCols ...string) *Archive {
	retobj, err := FindArchive(boil.GetDB(), index, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindArchive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArchive(exec boil.Executor, index int, selectCols ...string) (*Archive, error) {
	archiveObj := &Archive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `archives` where `index`=?", sel,
	)

	q := queries.Raw(exec, query, index)

	err := q.Bind(archiveObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from archives")
	}

	return archiveObj, nil
}

// FindArchiveP retrieves a single record by ID with an executor, and panics on error.
func FindArchiveP(exec boil.Executor, index int, selectCols ...string) *Archive {
	retobj, err := FindArchive(exec, index, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Archive) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Archive) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Archive) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Archive) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no archives provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	archiveInsertCacheMut.RLock()
	cache, cached := archiveInsertCache[key]
	archiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			archiveColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `archives` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `archives` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `archives` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into archives")
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

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == archiveMapping["Index"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Index,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for archives")
	}

CacheNoHooks:
	if !cached {
		archiveInsertCacheMut.Lock()
		archiveInsertCache[key] = cache
		archiveInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Archive record. See Update for
// whitelist behavior description.
func (o *Archive) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Archive record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Archive) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Archive, and panics on error.
// See Update for whitelist behavior description.
func (o *Archive) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Archive.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Archive) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	archiveUpdateCacheMut.RLock()
	cache, cached := archiveUpdateCache[key]
	archiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			archiveColumns,
			archivePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update archives, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `archives` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, append(wl, archivePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update archives row")
	}

	if !cached {
		archiveUpdateCacheMut.Lock()
		archiveUpdateCache[key] = cache
		archiveUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q archiveQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q archiveQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for archives")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ArchiveSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ArchiveSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ArchiveSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArchiveSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `archives` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in archive slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Archive) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Archive) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Archive) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Archive) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no archives provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)

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

	archiveUpsertCacheMut.RLock()
	cache, cached := archiveUpsertCache[key]
	archiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			archiveColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			archiveColumns,
			archivePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert archives, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "archives", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `archives` WHERE `index`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for archives")
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

	o.Index = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == archiveMapping["Index"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Index,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for archives")
	}

CacheNoHooks:
	if !cached {
		archiveUpsertCacheMut.Lock()
		archiveUpsertCache[key] = cache
		archiveUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Archive record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Archive) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Archive record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Archive) DeleteG() error {
	if o == nil {
		return errors.New("models: no Archive provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Archive record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Archive) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Archive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Archive) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Archive provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), archivePrimaryKeyMapping)
	sql := "DELETE FROM `archives` WHERE `index`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from archives")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q archiveQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q archiveQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no archiveQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from archives")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ArchiveSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ArchiveSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Archive slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ArchiveSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArchiveSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Archive slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(archiveBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `archives` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from archive slice")
	}

	if len(archiveAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Archive) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Archive) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Archive) ReloadG() error {
	if o == nil {
		return errors.New("models: no Archive provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Archive) Reload(exec boil.Executor) error {
	ret, err := FindArchive(exec, o.Index)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ArchiveSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ArchiveSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArchiveSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ArchiveSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArchiveSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	archives := ArchiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `archives`.* FROM `archives` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&archives)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ArchiveSlice")
	}

	*o = archives

	return nil
}

// ArchiveExists checks if the Archive row exists.
func ArchiveExists(exec boil.Executor, index int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `archives` where `index`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, index)
	}

	row := exec.QueryRow(sql, index)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if archives exists")
	}

	return exists, nil
}

// ArchiveExistsG checks if the Archive row exists.
func ArchiveExistsG(index int) (bool, error) {
	return ArchiveExists(boil.GetDB(), index)
}

// ArchiveExistsGP checks if the Archive row exists. Panics on error.
func ArchiveExistsGP(index int) bool {
	e, err := ArchiveExists(boil.GetDB(), index)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ArchiveExistsP checks if the Archive row exists. Panics on error.
func ArchiveExistsP(exec boil.Executor, index int) bool {
	e, err := ArchiveExists(exec, index)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
