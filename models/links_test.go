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

func testLinks(t *testing.T) {
	t.Parallel()

	query := Links(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testLinksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = link.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Links(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LinkSlice{link}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testLinksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := LinkExists(tx, link.ID)
	if err != nil {
		t.Errorf("Unable to check if Link exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LinkExistsG to return true, but got false.")
	}
}
func testLinksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	linkFound, err := FindLink(tx, link.ID)
	if err != nil {
		t.Error(err)
	}

	if linkFound == nil {
		t.Error("want a record, got nil")
	}
}
func testLinksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Links(tx).Bind(link); err != nil {
		t.Error(err)
	}
}

func testLinksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Links(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLinksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	linkOne := &Link{}
	linkTwo := &Link{}
	if err = randomize.Struct(seed, linkOne, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTwo, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = linkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = linkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Links(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	linkOne := &Link{}
	linkTwo := &Link{}
	if err = randomize.Struct(seed, linkOne, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTwo, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = linkOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = linkTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func linkBeforeInsertHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterInsertHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterSelectHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeUpdateHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterUpdateHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeDeleteHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterDeleteHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeUpsertHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterUpsertHook(e boil.Executor, o *Link) error {
	*o = Link{}
	return nil
}

func testLinksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Link{}
	o := &Link{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, linkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Link object: %s", err)
	}

	AddLinkHook(boil.BeforeInsertHook, linkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	linkBeforeInsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterInsertHook, linkAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	linkAfterInsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterSelectHook, linkAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	linkAfterSelectHooks = []LinkHook{}

	AddLinkHook(boil.BeforeUpdateHook, linkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	linkBeforeUpdateHooks = []LinkHook{}

	AddLinkHook(boil.AfterUpdateHook, linkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	linkAfterUpdateHooks = []LinkHook{}

	AddLinkHook(boil.BeforeDeleteHook, linkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	linkBeforeDeleteHooks = []LinkHook{}

	AddLinkHook(boil.AfterDeleteHook, linkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	linkAfterDeleteHooks = []LinkHook{}

	AddLinkHook(boil.BeforeUpsertHook, linkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	linkBeforeUpsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterUpsertHook, linkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	linkAfterUpsertHooks = []LinkHook{}
}
func testLinksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx, linkColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = link.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testLinksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := LinkSlice{link}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testLinksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Links(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	linkDBTypes = map[string]string{`Blogger`: `int`, `ID`: `int`, `Link`: `varchar`, `Reveal`: `tinyint`, `URL`: `varchar`}
	_           = bytes.MinRead
)

func testLinksUpdate(t *testing.T) {
	t.Parallel()

	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err = link.Update(tx); err != nil {
		t.Error(err)
	}
}

func testLinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	link := &Link{}
	if err = randomize.Struct(seed, link, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, link, linkDBTypes, true, linkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(linkColumns, linkPrimaryKeyColumns) {
		fields = linkColumns
	} else {
		fields = strmangle.SetComplement(
			linkColumns,
			linkPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(link))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := LinkSlice{link}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testLinksUpsert(t *testing.T) {
	t.Parallel()

	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	link := Link{}
	if err = randomize.Struct(seed, &link, linkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = link.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Link: %s", err)
	}

	count, err := Links(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &link, linkDBTypes, false, linkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err = link.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Link: %s", err)
	}

	count, err = Links(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
