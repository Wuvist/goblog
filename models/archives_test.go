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

func testArchives(t *testing.T) {
	t.Parallel()

	query := Archives(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testArchivesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = archive.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArchivesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Archives(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArchivesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ArchiveSlice{archive}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testArchivesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ArchiveExists(tx, archive.Index)
	if err != nil {
		t.Errorf("Unable to check if Archive exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArchiveExistsG to return true, but got false.")
	}
}
func testArchivesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	archiveFound, err := FindArchive(tx, archive.Index)
	if err != nil {
		t.Error(err)
	}

	if archiveFound == nil {
		t.Error("want a record, got nil")
	}
}
func testArchivesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Archives(tx).Bind(archive); err != nil {
		t.Error(err)
	}
}

func testArchivesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Archives(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArchivesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archiveOne := &Archive{}
	archiveTwo := &Archive{}
	if err = randomize.Struct(seed, archiveOne, archiveDBTypes, false, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}
	if err = randomize.Struct(seed, archiveTwo, archiveDBTypes, false, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archiveOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = archiveTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Archives(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArchivesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	archiveOne := &Archive{}
	archiveTwo := &Archive{}
	if err = randomize.Struct(seed, archiveOne, archiveDBTypes, false, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}
	if err = randomize.Struct(seed, archiveTwo, archiveDBTypes, false, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archiveOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = archiveTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func archiveBeforeInsertHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveAfterInsertHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveAfterSelectHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveBeforeUpdateHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveAfterUpdateHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveBeforeDeleteHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveAfterDeleteHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveBeforeUpsertHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func archiveAfterUpsertHook(e boil.Executor, o *Archive) error {
	*o = Archive{}
	return nil
}

func testArchivesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Archive{}
	o := &Archive{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, archiveDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Archive object: %s", err)
	}

	AddArchiveHook(boil.BeforeInsertHook, archiveBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	archiveBeforeInsertHooks = []ArchiveHook{}

	AddArchiveHook(boil.AfterInsertHook, archiveAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	archiveAfterInsertHooks = []ArchiveHook{}

	AddArchiveHook(boil.AfterSelectHook, archiveAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	archiveAfterSelectHooks = []ArchiveHook{}

	AddArchiveHook(boil.BeforeUpdateHook, archiveBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	archiveBeforeUpdateHooks = []ArchiveHook{}

	AddArchiveHook(boil.AfterUpdateHook, archiveAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	archiveAfterUpdateHooks = []ArchiveHook{}

	AddArchiveHook(boil.BeforeDeleteHook, archiveBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	archiveBeforeDeleteHooks = []ArchiveHook{}

	AddArchiveHook(boil.AfterDeleteHook, archiveAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	archiveAfterDeleteHooks = []ArchiveHook{}

	AddArchiveHook(boil.BeforeUpsertHook, archiveBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	archiveBeforeUpsertHooks = []ArchiveHook{}

	AddArchiveHook(boil.AfterUpsertHook, archiveAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	archiveAfterUpsertHooks = []ArchiveHook{}
}
func testArchivesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArchivesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx, archiveColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArchivesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = archive.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testArchivesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ArchiveSlice{archive}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testArchivesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Archives(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	archiveDBTypes = map[string]string{`AddDate`: `datetime`, `Author`: `int`, `Content`: `longtext`, `Emot`: `int`, `Index`: `int`, `Ref`: `int`, `Title`: `varchar`}
	_              = bytes.MinRead
)

func testArchivesUpdate(t *testing.T) {
	t.Parallel()

	if len(archiveColumns) == len(archivePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	if err = archive.Update(tx); err != nil {
		t.Error(err)
	}
}

func testArchivesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(archiveColumns) == len(archivePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	archive := &Archive{}
	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archiveColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, archive, archiveDBTypes, true, archivePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(archiveColumns, archivePrimaryKeyColumns) {
		fields = archiveColumns
	} else {
		fields = strmangle.SetComplement(
			archiveColumns,
			archivePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(archive))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ArchiveSlice{archive}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testArchivesUpsert(t *testing.T) {
	t.Parallel()

	if len(archiveColumns) == len(archivePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	archive := Archive{}
	if err = randomize.Struct(seed, &archive, archiveDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = archive.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Archive: %s", err)
	}

	count, err := Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &archive, archiveDBTypes, false, archivePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Archive struct: %s", err)
	}

	if err = archive.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Archive: %s", err)
	}

	count, err = Archives(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
