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

func testInfos(t *testing.T) {
	t.Parallel()

	query := Infos(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = info.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Infos(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InfoSlice{info}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := InfoExists(tx, info.ID)
	if err != nil {
		t.Errorf("Unable to check if Info exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InfoExistsG to return true, but got false.")
	}
}
func testInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	infoFound, err := FindInfo(tx, info.ID)
	if err != nil {
		t.Error(err)
	}

	if infoFound == nil {
		t.Error("want a record, got nil")
	}
}
func testInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Infos(tx).Bind(info); err != nil {
		t.Error(err)
	}
}

func testInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Infos(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	infoOne := &Info{}
	infoTwo := &Info{}
	if err = randomize.Struct(seed, infoOne, infoDBTypes, false, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}
	if err = randomize.Struct(seed, infoTwo, infoDBTypes, false, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = infoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = infoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Infos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	infoOne := &Info{}
	infoTwo := &Info{}
	if err = randomize.Struct(seed, infoOne, infoDBTypes, false, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}
	if err = randomize.Struct(seed, infoTwo, infoDBTypes, false, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = infoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = infoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func infoBeforeInsertHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoAfterInsertHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoAfterSelectHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoBeforeUpdateHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoAfterUpdateHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoBeforeDeleteHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoAfterDeleteHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoBeforeUpsertHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func infoAfterUpsertHook(e boil.Executor, o *Info) error {
	*o = Info{}
	return nil
}

func testInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Info{}
	o := &Info{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, infoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Info object: %s", err)
	}

	AddInfoHook(boil.BeforeInsertHook, infoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	infoBeforeInsertHooks = []InfoHook{}

	AddInfoHook(boil.AfterInsertHook, infoAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	infoAfterInsertHooks = []InfoHook{}

	AddInfoHook(boil.AfterSelectHook, infoAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	infoAfterSelectHooks = []InfoHook{}

	AddInfoHook(boil.BeforeUpdateHook, infoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	infoBeforeUpdateHooks = []InfoHook{}

	AddInfoHook(boil.AfterUpdateHook, infoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	infoAfterUpdateHooks = []InfoHook{}

	AddInfoHook(boil.BeforeDeleteHook, infoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	infoBeforeDeleteHooks = []InfoHook{}

	AddInfoHook(boil.AfterDeleteHook, infoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	infoAfterDeleteHooks = []InfoHook{}

	AddInfoHook(boil.BeforeUpsertHook, infoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	infoBeforeUpsertHooks = []InfoHook{}

	AddInfoHook(boil.AfterUpsertHook, infoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	infoAfterUpsertHooks = []InfoHook{}
}
func testInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx, infoColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = info.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InfoSlice{info}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Infos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	infoDBTypes = map[string]string{`Count`: `int`, `ID`: `smallint`, `Notes`: `varchar`}
	_           = bytes.MinRead
)

func testInfosUpdate(t *testing.T) {
	t.Parallel()

	if len(infoColumns) == len(infoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	if err = info.Update(tx); err != nil {
		t.Error(err)
	}
}

func testInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(infoColumns) == len(infoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	info := &Info{}
	if err = randomize.Struct(seed, info, infoDBTypes, true, infoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, info, infoDBTypes, true, infoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(infoColumns, infoPrimaryKeyColumns) {
		fields = infoColumns
	} else {
		fields = strmangle.SetComplement(
			infoColumns,
			infoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(info))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := InfoSlice{info}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(infoColumns) == len(infoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	info := Info{}
	if err = randomize.Struct(seed, &info, infoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = info.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Info: %s", err)
	}

	count, err := Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &info, infoDBTypes, false, infoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Info struct: %s", err)
	}

	if err = info.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Info: %s", err)
	}

	count, err = Infos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
