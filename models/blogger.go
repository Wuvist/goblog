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

// Blogger is an object representing the database table.
type Blogger struct {
	Index    int         `boil:"index" json:"index" toml:"index" yaml:"index"`
	ID       string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Nick     null.String `boil:"nick" json:"nick,omitempty" toml:"nick" yaml:"nick,omitempty"`
	DOB      time.Time   `boil:"DOB" json:"DOB" toml:"DOB" yaml:"DOB"`
	Blogname string      `boil:"blogname" json:"blogname" toml:"blogname" yaml:"blogname"`
	Blogskin int         `boil:"blogskin" json:"blogskin" toml:"blogskin" yaml:"blogskin"`
	Email    string      `boil:"email" json:"email" toml:"email" yaml:"email"`
	Visitor  int         `boil:"visitor" json:"visitor" toml:"visitor" yaml:"visitor"`
	Intro    null.String `boil:"intro" json:"intro,omitempty" toml:"intro" yaml:"intro,omitempty"`
	Blogs    int         `boil:"blogs" json:"blogs" toml:"blogs" yaml:"blogs"`
	IP       null.String `boil:"IP" json:"IP,omitempty" toml:"IP" yaml:"IP,omitempty"`
	TS       int8        `boil:"TS" json:"TS" toml:"TS" yaml:"TS"`
	LastLog  time.Time   `boil:"Last_log" json:"Last_log" toml:"Last_log" yaml:"Last_log"`
	LastPost null.Time   `boil:"Last_post" json:"Last_post,omitempty" toml:"Last_post" yaml:"Last_post,omitempty"`
	Activate int8        `boil:"Activate" json:"Activate" toml:"Activate" yaml:"Activate"`
	Reveal   int8        `boil:"Reveal" json:"Reveal" toml:"Reveal" yaml:"Reveal"`
	Userpic  int         `boil:"userpic" json:"userpic" toml:"userpic" yaml:"userpic"`
	Lang     string      `boil:"lang" json:"lang" toml:"lang" yaml:"lang"`
	Widget   null.String `boil:"widget" json:"widget,omitempty" toml:"widget" yaml:"widget,omitempty"`

	R *bloggerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bloggerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BloggerColumns = struct {
	Index    string
	ID       string
	Nick     string
	DOB      string
	Blogname string
	Blogskin string
	Email    string
	Visitor  string
	Intro    string
	Blogs    string
	IP       string
	TS       string
	LastLog  string
	LastPost string
	Activate string
	Reveal   string
	Userpic  string
	Lang     string
	Widget   string
}{
	Index:    "index",
	ID:       "id",
	Nick:     "nick",
	DOB:      "DOB",
	Blogname: "blogname",
	Blogskin: "blogskin",
	Email:    "email",
	Visitor:  "visitor",
	Intro:    "intro",
	Blogs:    "blogs",
	IP:       "IP",
	TS:       "TS",
	LastLog:  "Last_log",
	LastPost: "Last_post",
	Activate: "Activate",
	Reveal:   "Reveal",
	Userpic:  "userpic",
	Lang:     "lang",
	Widget:   "widget",
}

// bloggerR is where relationships are stored.
type bloggerR struct {
}

// bloggerL is where Load methods for each relationship are stored.
type bloggerL struct{}

var (
	bloggerColumns               = []string{"index", "id", "nick", "DOB", "blogname", "blogskin", "email", "visitor", "intro", "blogs", "IP", "TS", "Last_log", "Last_post", "Activate", "Reveal", "userpic", "lang", "widget"}
	bloggerColumnsWithoutDefault = []string{"id", "nick", "DOB", "blogname", "blogskin", "email", "intro", "IP", "Last_post", "widget"}
	bloggerColumnsWithDefault    = []string{"index", "visitor", "blogs", "TS", "Last_log", "Activate", "Reveal", "userpic", "lang"}
	bloggerPrimaryKeyColumns     = []string{"index"}
)

type (
	// BloggerSlice is an alias for a slice of pointers to Blogger.
	// This should generally be used opposed to []Blogger.
	BloggerSlice []*Blogger
	// BloggerHook is the signature for custom Blogger hook methods
	BloggerHook func(boil.Executor, *Blogger) error

	bloggerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bloggerType                 = reflect.TypeOf(&Blogger{})
	bloggerMapping              = queries.MakeStructMapping(bloggerType)
	bloggerPrimaryKeyMapping, _ = queries.BindMapping(bloggerType, bloggerMapping, bloggerPrimaryKeyColumns)
	bloggerInsertCacheMut       sync.RWMutex
	bloggerInsertCache          = make(map[string]insertCache)
	bloggerUpdateCacheMut       sync.RWMutex
	bloggerUpdateCache          = make(map[string]updateCache)
	bloggerUpsertCacheMut       sync.RWMutex
	bloggerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var bloggerBeforeInsertHooks []BloggerHook
var bloggerBeforeUpdateHooks []BloggerHook
var bloggerBeforeDeleteHooks []BloggerHook
var bloggerBeforeUpsertHooks []BloggerHook

var bloggerAfterInsertHooks []BloggerHook
var bloggerAfterSelectHooks []BloggerHook
var bloggerAfterUpdateHooks []BloggerHook
var bloggerAfterDeleteHooks []BloggerHook
var bloggerAfterUpsertHooks []BloggerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Blogger) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Blogger) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Blogger) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Blogger) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Blogger) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Blogger) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Blogger) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Blogger) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Blogger) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bloggerAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBloggerHook registers your hook function for all future operations.
func AddBloggerHook(hookPoint boil.HookPoint, bloggerHook BloggerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bloggerBeforeInsertHooks = append(bloggerBeforeInsertHooks, bloggerHook)
	case boil.BeforeUpdateHook:
		bloggerBeforeUpdateHooks = append(bloggerBeforeUpdateHooks, bloggerHook)
	case boil.BeforeDeleteHook:
		bloggerBeforeDeleteHooks = append(bloggerBeforeDeleteHooks, bloggerHook)
	case boil.BeforeUpsertHook:
		bloggerBeforeUpsertHooks = append(bloggerBeforeUpsertHooks, bloggerHook)
	case boil.AfterInsertHook:
		bloggerAfterInsertHooks = append(bloggerAfterInsertHooks, bloggerHook)
	case boil.AfterSelectHook:
		bloggerAfterSelectHooks = append(bloggerAfterSelectHooks, bloggerHook)
	case boil.AfterUpdateHook:
		bloggerAfterUpdateHooks = append(bloggerAfterUpdateHooks, bloggerHook)
	case boil.AfterDeleteHook:
		bloggerAfterDeleteHooks = append(bloggerAfterDeleteHooks, bloggerHook)
	case boil.AfterUpsertHook:
		bloggerAfterUpsertHooks = append(bloggerAfterUpsertHooks, bloggerHook)
	}
}

// OneP returns a single blogger record from the query, and panics on error.
func (q bloggerQuery) OneP() *Blogger {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single blogger record from the query.
func (q bloggerQuery) One() (*Blogger, error) {
	o := &Blogger{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for blogger")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Blogger records from the query, and panics on error.
func (q bloggerQuery) AllP() BloggerSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Blogger records from the query.
func (q bloggerQuery) All() (BloggerSlice, error) {
	var o []*Blogger

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Blogger slice")
	}

	if len(bloggerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Blogger records in the query, and panics on error.
func (q bloggerQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Blogger records in the query.
func (q bloggerQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count blogger rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q bloggerQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q bloggerQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if blogger exists")
	}

	return count > 0, nil
}

// BloggersG retrieves all records.
func BloggersG(mods ...qm.QueryMod) bloggerQuery {
	return Bloggers(boil.GetDB(), mods...)
}

// Bloggers retrieves all the records using an executor.
func Bloggers(exec boil.Executor, mods ...qm.QueryMod) bloggerQuery {
	mods = append(mods, qm.From("`blogger`"))
	return bloggerQuery{NewQuery(exec, mods...)}
}

// FindBloggerG retrieves a single record by ID.
func FindBloggerG(index int, selectCols ...string) (*Blogger, error) {
	return FindBlogger(boil.GetDB(), index, selectCols...)
}

// FindBloggerGP retrieves a single record by ID, and panics on error.
func FindBloggerGP(index int, selectCols ...string) *Blogger {
	retobj, err := FindBlogger(boil.GetDB(), index, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindBlogger retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlogger(exec boil.Executor, index int, selectCols ...string) (*Blogger, error) {
	bloggerObj := &Blogger{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `blogger` where `index`=?", sel,
	)

	q := queries.Raw(exec, query, index)

	err := q.Bind(bloggerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from blogger")
	}

	return bloggerObj, nil
}

// FindBloggerP retrieves a single record by ID with an executor, and panics on error.
func FindBloggerP(exec boil.Executor, index int, selectCols ...string) *Blogger {
	retobj, err := FindBlogger(exec, index, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Blogger) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Blogger) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Blogger) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Blogger) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blogger provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bloggerColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	bloggerInsertCacheMut.RLock()
	cache, cached := bloggerInsertCache[key]
	bloggerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			bloggerColumns,
			bloggerColumnsWithDefault,
			bloggerColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(bloggerType, bloggerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bloggerType, bloggerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `blogger` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `blogger` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `blogger` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, bloggerPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into blogger")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bloggerMapping["Index"] {
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
		return errors.Wrap(err, "models: unable to populate default values for blogger")
	}

CacheNoHooks:
	if !cached {
		bloggerInsertCacheMut.Lock()
		bloggerInsertCache[key] = cache
		bloggerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Blogger record. See Update for
// whitelist behavior description.
func (o *Blogger) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Blogger record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Blogger) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Blogger, and panics on error.
// See Update for whitelist behavior description.
func (o *Blogger) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Blogger.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Blogger) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	bloggerUpdateCacheMut.RLock()
	cache, cached := bloggerUpdateCache[key]
	bloggerUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			bloggerColumns,
			bloggerPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update blogger, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `blogger` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, bloggerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bloggerType, bloggerMapping, append(wl, bloggerPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update blogger row")
	}

	if !cached {
		bloggerUpdateCacheMut.Lock()
		bloggerUpdateCache[key] = cache
		bloggerUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q bloggerQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q bloggerQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for blogger")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BloggerSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o BloggerSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o BloggerSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BloggerSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloggerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `blogger` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloggerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in blogger slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Blogger) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Blogger) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Blogger) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Blogger) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no blogger provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bloggerColumnsWithDefault, o)

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

	bloggerUpsertCacheMut.RLock()
	cache, cached := bloggerUpsertCache[key]
	bloggerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			bloggerColumns,
			bloggerColumnsWithDefault,
			bloggerColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			bloggerColumns,
			bloggerPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert blogger, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "blogger", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `blogger` WHERE `index`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(bloggerType, bloggerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bloggerType, bloggerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for blogger")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == bloggerMapping["Index"] {
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
		return errors.Wrap(err, "models: unable to populate default values for blogger")
	}

CacheNoHooks:
	if !cached {
		bloggerUpsertCacheMut.Lock()
		bloggerUpsertCache[key] = cache
		bloggerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Blogger record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Blogger) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Blogger record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Blogger) DeleteG() error {
	if o == nil {
		return errors.New("models: no Blogger provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Blogger record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Blogger) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Blogger record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Blogger) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Blogger provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bloggerPrimaryKeyMapping)
	sql := "DELETE FROM `blogger` WHERE `index`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from blogger")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q bloggerQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q bloggerQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no bloggerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blogger")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o BloggerSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o BloggerSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Blogger slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o BloggerSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BloggerSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Blogger slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(bloggerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloggerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `blogger` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloggerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from blogger slice")
	}

	if len(bloggerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Blogger) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Blogger) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Blogger) ReloadG() error {
	if o == nil {
		return errors.New("models: no Blogger provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Blogger) Reload(exec boil.Executor) error {
	ret, err := FindBlogger(exec, o.Index)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BloggerSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *BloggerSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BloggerSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BloggerSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BloggerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	bloggers := BloggerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bloggerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `blogger`.* FROM `blogger` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, bloggerPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&bloggers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BloggerSlice")
	}

	*o = bloggers

	return nil
}

// BloggerExists checks if the Blogger row exists.
func BloggerExists(exec boil.Executor, index int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `blogger` where `index`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, index)
	}

	row := exec.QueryRow(sql, index)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if blogger exists")
	}

	return exists, nil
}

// BloggerExistsG checks if the Blogger row exists.
func BloggerExistsG(index int) (bool, error) {
	return BloggerExists(boil.GetDB(), index)
}

// BloggerExistsGP checks if the Blogger row exists. Panics on error.
func BloggerExistsGP(index int) bool {
	e, err := BloggerExists(boil.GetDB(), index)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// BloggerExistsP checks if the Blogger row exists. Panics on error.
func BloggerExistsP(exec boil.Executor, index int) bool {
	e, err := BloggerExists(exec, index)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
