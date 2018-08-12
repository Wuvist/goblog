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

// BlogTag is an object representing the database table.
type BlogTag struct {
	Aid int `boil:"aid" json:"aid" toml:"aid" yaml:"aid"`
	Tid int `boil:"tid" json:"tid" toml:"tid" yaml:"tid"`

	R *blogTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blogTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlogTagColumns = struct {
	Aid string
	Tid string
}{
	Aid: "aid",
	Tid: "tid",
}

// blogTagR is where relationships are stored.
type blogTagR struct {
}

// blogTagL is where Load methods for each relationship are stored.
type blogTagL struct{}

var (
	blogTagColumns               = []string{"aid", "tid"}
	blogTagColumnsWithoutDefault = []string{"aid", "tid"}
	blogTagColumnsWithDefault    = []string{}
	blogTagPrimaryKeyColumns     = []string{"aid", "tid"}
)

type (
	// BlogTagSlice is an alias for a slice of pointers to BlogTag.
	// This should generally be used opposed to []BlogTag.
	BlogTagSlice []*BlogTag
	// BlogTagHook is the signature for custom BlogTag hook methods
	BlogTagHook func(boil.Executor, *BlogTag) error

	blogTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blogTagType                 = reflect.TypeOf(&BlogTag{})
	blogTagMapping              = queries.MakeStructMapping(blogTagType)
	blogTagPrimaryKeyMapping, _ = queries.BindMapping(blogTagType, blogTagMapping, blogTagPrimaryKeyColumns)
	blogTagInsertCacheMut       sync.RWMutex
	blogTagInsertCache          = make(map[string]insertCache)
	blogTagUpdateCacheMut       sync.RWMutex
	blogTagUpdateCache          = make(map[string]updateCache)
	blogTagUpsertCacheMut       sync.RWMutex
	blogTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var blogTagBeforeInsertHooks []BlogTagHook
var blogTagBeforeUpdateHooks []BlogTagHook
var blogTagBeforeDeleteHooks []BlogTagHook
var blogTagBeforeUpsertHooks []BlogTagHook

var blogTagAfterInsertHooks []BlogTagHook
var blogTagAfterSelectHooks []BlogTagHook
var blogTagAfterUpdateHooks []BlogTagHook
var blogTagAfterDeleteHooks []BlogTagHook
var blogTagAfterUpsertHooks []BlogTagHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BlogTag) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BlogTag) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BlogTag) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BlogTag) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BlogTag) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BlogTag) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BlogTag) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BlogTag) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BlogTag) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range blogTagAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBlogTagHook registers your hook function for all future operations.
func AddBlogTagHook(hookPoint boil.HookPoint, blogTagHook BlogTagHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		blogTagBeforeInsertHooks = append(blogTagBeforeInsertHooks, blogTagHook)
	case boil.BeforeUpdateHook:
		blogTagBeforeUpdateHooks = append(blogTagBeforeUpdateHooks, blogTagHook)
	case boil.BeforeDeleteHook:
		blogTagBeforeDeleteHooks = append(blogTagBeforeDeleteHooks, blogTagHook)
	case boil.BeforeUpsertHook:
		blogTagBeforeUpsertHooks = append(blogTagBeforeUpsertHooks, blogTagHook)
	case boil.AfterInsertHook:
		blogTagAfterInsertHooks = append(blogTagAfterInsertHooks, blogTagHook)
	case boil.AfterSelectHook:
		blogTagAfterSelectHooks = append(blogTagAfterSelectHooks, blogTagHook)
	case boil.AfterUpdateHook:
		blogTagAfterUpdateHooks = append(blogTagAfterUpdateHooks, blogTagHook)
	case boil.AfterDeleteHook:
		blogTagAfterDeleteHooks = append(blogTagAfterDeleteHooks, blogTagHook)
	case boil.AfterUpsertHook:
		blogTagAfterUpsertHooks = append(blogTagAfterUpsertHooks, blogTagHook)
	}
}

// OneP returns a single blogTag record from the query, and panics on error.
func (q blogTagQuery) OneP() *BlogTag {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single blogTag record from the query.
func (q blogTagQuery) One() (*BlogTag, error) {
	o := &BlogTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blog_tag")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all BlogTag records from the query, and panics on error.
func (q blogTagQuery) AllP() BlogTagSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all BlogTag records from the query.
func (q blogTagQuery) All() (BlogTagSlice, error) {
	var o []*BlogTag

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BlogTag slice")
	}

	if len(blogTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all BlogTag records in the query, and panics on error.
func (q blogTagQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all BlogTag records in the query.
func (q blogTagQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blog_tag rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q blogTagQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q blogTagQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blog_tag exists")
	}

	return count > 0, nil
}

// BlogTagsG retrieves all records.
func BlogTagsG(mods ...qm.QueryMod) blogTagQuery {
	return BlogTags(boil.GetDB(), mods...)
}

// BlogTags retrieves all the records using an executor.
func BlogTags(exec boil.Executor, mods ...qm.QueryMod) blogTagQuery {
	mods = append(mods, qm.From("`blog_tag`"))
	return blogTagQuery{NewQuery(exec, mods...)}
}

// FindBlogTagG retrieves a single record by ID.
func FindBlogTagG(aid int, tid int, selectCols ...string) (*BlogTag, error) {
	return FindBlogTag(boil.GetDB(), aid, tid, selectCols...)
}

// FindBlogTagGP retrieves a single record by ID, and panics on error.
func FindBlogTagGP(aid int, tid int, selectCols ...string) *BlogTag {
	retobj, err := FindBlogTag(boil.GetDB(), aid, tid, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindBlogTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlogTag(exec boil.Executor, aid int, tid int, selectCols ...string) (*BlogTag, error) {
	blogTagObj := &BlogTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `blog_tag` where `aid`=? AND `tid`=?", sel,
	)

	q := queries.Raw(exec, query, aid, tid)

	err := q.Bind(blogTagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blog_tag")
	}

	return blogTagObj, nil
}

// FindBlogTagP retrieves a single record by ID with an executor, and panics on error.
func FindBlogTagP(exec boil.Executor, aid int, tid int, selectCols ...string) *BlogTag {
	retobj, err := FindBlogTag(exec, aid, tid, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BlogTag) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *BlogTag) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *BlogTag) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *BlogTag) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blog_tag provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogTagColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	blogTagInsertCacheMut.RLock()
	cache, cached := blogTagInsertCache[key]
	blogTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			blogTagColumns,
			blogTagColumnsWithDefault,
			blogTagColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(blogTagType, blogTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blogTagType, blogTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `blog_tag` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `blog_tag` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `blog_tag` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, blogTagPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into blog_tag")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Aid,
		o.Tid,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for blog_tag")
	}

CacheNoHooks:
	if !cached {
		blogTagInsertCacheMut.Lock()
		blogTagInsertCache[key] = cache
		blogTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single BlogTag record. See Update for
// whitelist behavior description.
func (o *BlogTag) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single BlogTag record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *BlogTag) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the BlogTag, and panics on error.
// See Update for whitelist behavior description.
func (o *BlogTag) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the BlogTag.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *BlogTag) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	blogTagUpdateCacheMut.RLock()
	cache, cached := blogTagUpdateCache[key]
	blogTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			blogTagColumns,
			blogTagPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update blog_tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `blog_tag` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, blogTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blogTagType, blogTagMapping, append(wl, blogTagPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update blog_tag row")
	}

	if !cached {
		blogTagUpdateCacheMut.Lock()
		blogTagUpdateCache[key] = cache
		blogTagUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q blogTagQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q blogTagQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for blog_tag")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BlogTagSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o BlogTagSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o BlogTagSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlogTagSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `blog_tag` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogTagPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in blogTag slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BlogTag) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *BlogTag) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *BlogTag) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *BlogTag) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blog_tag provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(blogTagColumnsWithDefault, o)

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

	blogTagUpsertCacheMut.RLock()
	cache, cached := blogTagUpsertCache[key]
	blogTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			blogTagColumns,
			blogTagColumnsWithDefault,
			blogTagColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			blogTagColumns,
			blogTagPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert blog_tag, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "blog_tag", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `blog_tag` WHERE `aid`=? AND `tid`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(blogTagType, blogTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blogTagType, blogTagMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for blog_tag")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Aid,
		o.Tid,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for blog_tag")
	}

CacheNoHooks:
	if !cached {
		blogTagUpsertCacheMut.Lock()
		blogTagUpsertCache[key] = cache
		blogTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single BlogTag record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BlogTag) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single BlogTag record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BlogTag) DeleteG() error {
	if o == nil {
		return errors.New("models: no BlogTag provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single BlogTag record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *BlogTag) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single BlogTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BlogTag) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BlogTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blogTagPrimaryKeyMapping)
	sql := "DELETE FROM `blog_tag` WHERE `aid`=? AND `tid`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from blog_tag")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q blogTagQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q blogTagQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no blogTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blog_tag")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o BlogTagSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o BlogTagSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no BlogTag slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o BlogTagSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlogTagSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no BlogTag slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(blogTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `blog_tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogTagPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blogTag slice")
	}

	if len(blogTagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *BlogTag) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *BlogTag) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BlogTag) ReloadG() error {
	if o == nil {
		return errors.New("models: no BlogTag provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BlogTag) Reload(exec boil.Executor) error {
	ret, err := FindBlogTag(exec, o.Aid, o.Tid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BlogTagSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BlogTagSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogTagSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BlogTagSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlogTagSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	blogTags := BlogTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blogTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `blog_tag`.* FROM `blog_tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, blogTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&blogTags)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlogTagSlice")
	}

	*o = blogTags

	return nil
}

// BlogTagExists checks if the BlogTag row exists.
func BlogTagExists(exec boil.Executor, aid int, tid int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `blog_tag` where `aid`=? AND `tid`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, aid, tid)
	}

	row := exec.QueryRow(sql, aid, tid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blog_tag exists")
	}

	return exists, nil
}

// BlogTagExistsG checks if the BlogTag row exists.
func BlogTagExistsG(aid int, tid int) (bool, error) {
	return BlogTagExists(boil.GetDB(), aid, tid)
}

// BlogTagExistsGP checks if the BlogTag row exists. Panics on error.
func BlogTagExistsGP(aid int, tid int) bool {
	e, err := BlogTagExists(boil.GetDB(), aid, tid)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// BlogTagExistsP checks if the BlogTag row exists. Panics on error.
func BlogTagExistsP(exec boil.Executor, aid int, tid int) bool {
	e, err := BlogTagExists(exec, aid, tid)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}