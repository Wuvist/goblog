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

// Blogcategory is an object representing the database table.
type Blogcategory struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Cate string `boil:"cate" json:"cate" toml:"cate" yaml:"cate"`

	R *blogcategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blogcategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlogcategoryColumns = struct {
	ID   string
	Cate string
}{
	ID:   "id",
	Cate: "cate",
}

// blogcategoryR is where relationships are stored.
type blogcategoryR struct {
}

// blogcategoryL is where Load methods for each relationship are stored.
type blogcategoryL struct{}

var (
	blogcategoryColumns               = []string{"id", "cate"}
	blogcategoryColumnsWithoutDefault = []string{"cate"}
	blogcategoryColumnsWithDefault    = []string{"id"}
	blogcategoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// BlogcategorySlice is an alias for a slice of pointers to Blogcategory.
	// This should generally be used opposed to []Blogcategory.
	BlogcategorySlice []*Blogcategory
	// BlogcategoryHook is the signature for custom Blogcategory hook methods
	BlogcategoryHook func(boil.Executor, *Blogcategory) error

	blogcategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blogcategoryType                 = reflect.TypeOf(&Blogcategory{})
	blogcategoryMapping              = queries.MakeStructMapping(blogcategoryType)
	blogcategoryPrimaryKeyMapping, _ = queries.BindMapping(blogcategoryType, blogcategoryMapping, blogcategoryPrimaryKeyColumns)
	blogcategoryInsertCacheMut       sync.RWMutex
	blogcategoryInsertCache          = make(map[string]insertCache)
	blogcategoryUpdateCacheMut       sync.RWMutex
	blogcategoryUpdateCache          = make(map[string]updateCache)
	blogcategoryUpsertCacheMut       sync.RWMutex
	blogcategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var blogcategoryBeforeInsertHooks []BlogcategoryHook
var blogcategoryBeforeUpdateHooks []BlogcategoryHook
var blogcategoryBeforeDeleteHooks []BlogcategoryHook
var blogcategoryBeforeUpsertHooks []BlogcategoryHook

var blogcategoryAfterInsertHooks []BlogcategoryHook
var blogcategoryAfterSelectHooks []BlogcategoryHook
var blogcategoryAfterUpdateHooks []BlogcategoryHook
var blogcategoryAfterDeleteHooks []BlogcategoryHook
var blogcategoryAfterUpsertHooks []BlogcategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Blogcategory) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Blogcategory) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Blogcategory) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Blogcategory) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Blogcategory) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Blogcategory) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Blogcategory) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Blogcategory) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Blogcategory) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogcategoryAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBlogcategoryHook registers your hook function for all future operations.
func AddBlogcategoryHook(hookPoint boil.HookPoint, blogcategoryHook BlogcategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		blogcategoryBeforeInsertHooks = append(blogcategoryBeforeInsertHooks, blogcategoryHook)
	case boil.BeforeUpdateHook:
		blogcategoryBeforeUpdateHooks = append(blogcategoryBeforeUpdateHooks, blogcategoryHook)
	case boil.BeforeDeleteHook:
		blogcategoryBeforeDeleteHooks = append(blogcategoryBeforeDeleteHooks, blogcategoryHook)
	case boil.BeforeUpsertHook:
		blogcategoryBeforeUpsertHooks = append(blogcategoryBeforeUpsertHooks, blogcategoryHook)
	case boil.AfterInsertHook:
		blogcategoryAfterInsertHooks = append(blogcategoryAfterInsertHooks, blogcategoryHook)
	case boil.AfterSelectHook:
		blogcategoryAfterSelectHooks = append(blogcategoryAfterSelectHooks, blogcategoryHook)
	case boil.AfterUpdateHook:
		blogcategoryAfterUpdateHooks = append(blogcategoryAfterUpdateHooks, blogcategoryHook)
	case boil.AfterDeleteHook:
		blogcategoryAfterDeleteHooks = append(blogcategoryAfterDeleteHooks, blogcategoryHook)
	case boil.AfterUpsertHook:
		blogcategoryAfterUpsertHooks = append(blogcategoryAfterUpsertHooks, blogcategoryHook)
	}
}

// OneP returns a single blogcategory record from the query, and panics on error.
func (q blogcategoryQuery) OneP() *Blogcategory {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single blogcategory record from the query.
func (q blogcategoryQuery) One() (*Blogcategory, error) {
	o := &Blogcategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blogcategory")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Blogcategory records from the query, and panics on error.
func (q blogcategoryQuery) AllP() BlogcategorySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Blogcategory records from the query.
func (q blogcategoryQuery) All() (BlogcategorySlice, error) {
	var o []*Blogcategory

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Blogcategory slice")
	}

	if len(blogcategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Blogcategory records in the query, and panics on error.
func (q blogcategoryQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Blogcategory records in the query.
func (q blogcategoryQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blogcategory rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q blogcategoryQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q blogcategoryQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blogcategory exists")
	}

	return count > 0, nil
}

// BlogcategoriesG retrieves all records.
func BlogcategoriesG(mods ...qm.QueryMod) blogcategoryQuery {
	return Blogcategories(boil.GetDB(), mods...)
}

// Blogcategories retrieves all the records using an executor.
func Blogcategories(exec boil.Executor, mods ...qm.QueryMod) blogcategoryQuery {
	mods = append(mods, qm.From("`blogcategory`"))
	return blogcategoryQuery{NewQuery(exec, mods...)}
}

// FindBlogcategoryG retrieves a single record by ID.
func FindBlogcategoryG(id int, selectCols ...string) (*Blogcategory, error) {
	return FindBlogcategory(boil.GetDB(), id, selectCols...)
}

// FindBlogcategoryGP retrieves a single record by ID, and panics on error.
func FindBlogcategoryGP(id int, selectCols ...string) *Blogcategory {
	retobj, err := FindBlogcategory(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindBlogcategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlogcategory(exec boil.Executor, id int, selectCols ...string) (*Blogcategory, error) {
	blogcategoryObj := &Blogcategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `blogcategory` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(blogcategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blogcategory")
	}

	return blogcategoryObj, nil
}

// FindBlogcategoryP retrieves a single record by ID with an executor, and panics on error.
func FindBlogcategoryP(exec boil.Executor, id int, selectCols ...string) *Blogcategory {
	retobj, err := FindBlogcategory(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Blogcategory) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Blogcategory) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Blogcategory) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Blogcategory) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blogcategory provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogcategoryColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	blogcategoryInsertCacheMut.RLock()
	cache, cached := blogcategoryInsertCache[key]
	blogcategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			blogcategoryColumns,
			blogcategoryColumnsWithDefault,
			blogcategoryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(blogcategoryType, blogcategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blogcategoryType, blogcategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `blogcategory` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `blogcategory` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `blogcategory` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, blogcategoryPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into blogcategory")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == blogcategoryMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for blogcategory")
	}

CacheNoHooks:
	if !cached {
		blogcategoryInsertCacheMut.Lock()
		blogcategoryInsertCache[key] = cache
		blogcategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Blogcategory record. See Update for
// whitelist behavior description.
func (o *Blogcategory) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Blogcategory record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Blogcategory) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Blogcategory, and panics on error.
// See Update for whitelist behavior description.
func (o *Blogcategory) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Blogcategory.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Blogcategory) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	blogcategoryUpdateCacheMut.RLock()
	cache, cached := blogcategoryUpdateCache[key]
	blogcategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			blogcategoryColumns,
			blogcategoryPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update blogcategory, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `blogcategory` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, blogcategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blogcategoryType, blogcategoryMapping, append(wl, blogcategoryPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update blogcategory row")
	}

	if !cached {
		blogcategoryUpdateCacheMut.Lock()
		blogcategoryUpdateCache[key] = cache
		blogcategoryUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q blogcategoryQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q blogcategoryQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for blogcategory")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BlogcategorySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o BlogcategorySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o BlogcategorySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlogcategorySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogcategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `blogcategory` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogcategoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in blogcategory slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Blogcategory) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Blogcategory) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Blogcategory) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Blogcategory) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blogcategory provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogcategoryColumnsWithDefault, o)

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

	blogcategoryUpsertCacheMut.RLock()
	cache, cached := blogcategoryUpsertCache[key]
	blogcategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			blogcategoryColumns,
			blogcategoryColumnsWithDefault,
			blogcategoryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			blogcategoryColumns,
			blogcategoryPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert blogcategory, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "blogcategory", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `blogcategory` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(blogcategoryType, blogcategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blogcategoryType, blogcategoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for blogcategory")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == blogcategoryMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for blogcategory")
	}

CacheNoHooks:
	if !cached {
		blogcategoryUpsertCacheMut.Lock()
		blogcategoryUpsertCache[key] = cache
		blogcategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Blogcategory record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Blogcategory) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Blogcategory record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Blogcategory) DeleteG() error {
	if o == nil {
		return errors.New("models: no Blogcategory provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Blogcategory record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Blogcategory) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Blogcategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Blogcategory) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Blogcategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blogcategoryPrimaryKeyMapping)
	sql := "DELETE FROM `blogcategory` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from blogcategory")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q blogcategoryQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q blogcategoryQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no blogcategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blogcategory")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o BlogcategorySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o BlogcategorySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Blogcategory slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o BlogcategorySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlogcategorySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Blogcategory slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(blogcategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogcategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `blogcategory` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogcategoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blogcategory slice")
	}

	if len(blogcategoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Blogcategory) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Blogcategory) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Blogcategory) ReloadG() error {
	if o == nil {
		return errors.New("models: no Blogcategory provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Blogcategory) Reload(exec boil.Executor) error {
	ret, err := FindBlogcategory(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BlogcategorySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BlogcategorySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogcategorySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BlogcategorySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogcategorySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	blogcategories := BlogcategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogcategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `blogcategory`.* FROM `blogcategory` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogcategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&blogcategories)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlogcategorySlice")
	}

	*o = blogcategories

	return nil
}

// BlogcategoryExists checks if the Blogcategory row exists.
func BlogcategoryExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `blogcategory` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blogcategory exists")
	}

	return exists, nil
}

// BlogcategoryExistsG checks if the Blogcategory row exists.
func BlogcategoryExistsG(id int) (bool, error) {
	return BlogcategoryExists(boil.GetDB(), id)
}

// BlogcategoryExistsGP checks if the Blogcategory row exists. Panics on error.
func BlogcategoryExistsGP(id int) bool {
	e, err := BlogcategoryExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// BlogcategoryExistsP checks if the Blogcategory row exists. Panics on error.
func BlogcategoryExistsP(exec boil.Executor, id int) bool {
	e, err := BlogcategoryExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}