// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testBlogTags(t *testing.T) {
	t.Parallel()

	query := BlogTags(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBlogTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blogTag.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BlogTags(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlogTagSlice{blogTag}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBlogTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BlogTagExists(tx, blogTag.Aid, blogTag.Tid)
	if err != nil {
		t.Errorf("Unable to check if BlogTag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlogTagExistsG to return true, but got false.")
	}
}
func testBlogTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	blogTagFound, err := FindBlogTag(tx, blogTag.Aid, blogTag.Tid)
	if err != nil {
		t.Error(err)
	}

	if blogTagFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBlogTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BlogTags(tx).Bind(blogTag); err != nil {
		t.Error(err)
	}
}

func testBlogTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := BlogTags(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlogTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTagOne := &BlogTag{}
	blogTagTwo := &BlogTag{}
	if err = randomize.Struct(seed, blogTagOne, blogTagDBTypes, false, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTagTwo, blogTagDBTypes, false, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTagOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blogTagTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BlogTags(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlogTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blogTagOne := &BlogTag{}
	blogTagTwo := &BlogTag{}
	if err = randomize.Struct(seed, blogTagOne, blogTagDBTypes, false, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTagTwo, blogTagDBTypes, false, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTagOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blogTagTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func blogTagBeforeInsertHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagAfterInsertHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagAfterSelectHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagBeforeUpdateHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagAfterUpdateHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagBeforeDeleteHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagAfterDeleteHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagBeforeUpsertHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func blogTagAfterUpsertHook(e boil.Executor, o *BlogTag) error {
	*o = BlogTag{}
	return nil
}

func testBlogTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &BlogTag{}
	o := &BlogTag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, blogTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BlogTag object: %s", err)
	}

	AddBlogTagHook(boil.BeforeInsertHook, blogTagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	blogTagBeforeInsertHooks = []BlogTagHook{}

	AddBlogTagHook(boil.AfterInsertHook, blogTagAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	blogTagAfterInsertHooks = []BlogTagHook{}

	AddBlogTagHook(boil.AfterSelectHook, blogTagAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	blogTagAfterSelectHooks = []BlogTagHook{}

	AddBlogTagHook(boil.BeforeUpdateHook, blogTagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	blogTagBeforeUpdateHooks = []BlogTagHook{}

	AddBlogTagHook(boil.AfterUpdateHook, blogTagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	blogTagAfterUpdateHooks = []BlogTagHook{}

	AddBlogTagHook(boil.BeforeDeleteHook, blogTagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	blogTagBeforeDeleteHooks = []BlogTagHook{}

	AddBlogTagHook(boil.AfterDeleteHook, blogTagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	blogTagAfterDeleteHooks = []BlogTagHook{}

	AddBlogTagHook(boil.BeforeUpsertHook, blogTagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	blogTagBeforeUpsertHooks = []BlogTagHook{}

	AddBlogTagHook(boil.AfterUpsertHook, blogTagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	blogTagAfterUpsertHooks = []BlogTagHook{}
}
func testBlogTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx, blogTagColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blogTag.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBlogTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlogTagSlice{blogTag}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBlogTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BlogTags(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blogTagDBTypes = map[string]string{`Aid`: `int`, `Tid`: `int`}
	_              = bytes.MinRead
)

func testBlogTagsUpdate(t *testing.T) {
	t.Parallel()

	if len(blogTagColumns) == len(blogTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	if err = blogTag.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBlogTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blogTagColumns) == len(blogTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blogTag := &BlogTag{}
	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blogTag, blogTagDBTypes, true, blogTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blogTagColumns, blogTagPrimaryKeyColumns) {
		fields = blogTagColumns
	} else {
		fields = strmangle.SetComplement(
			blogTagColumns,
			blogTagPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(blogTag))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BlogTagSlice{blogTag}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBlogTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(blogTagColumns) == len(blogTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	blogTag := BlogTag{}
	if err = randomize.Struct(seed, &blogTag, blogTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blogTag.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert BlogTag: %s", err)
	}

	count, err := BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &blogTag, blogTagDBTypes, false, blogTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BlogTag struct: %s", err)
	}

	if err = blogTag.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert BlogTag: %s", err)
	}

	count, err = BlogTags(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
