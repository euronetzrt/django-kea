// Code generated for Django model django_kea.Lease6Types. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=keatest.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/euronetzrt/django-kea django_kea

  https://github.com/rkojedzinszky/djan-go-rm
*/

package django_kea

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/euronetzrt/django-kea/models"
)

// Lease6types mirrors model django_kea.Lease6Types
type Lease6types struct {
	existsInDB bool

	LeaseType int32
	Name      sql.NullString
}

// Lease6typesList is a list of Lease6types
type Lease6typesList []*Lease6types

// Lease6typesQS represents a queryset for django_kea.Lease6Types
type Lease6typesQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs Lease6typesQS) filter(c string, p interface{}) Lease6typesQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// Or combines given expressions with OR operator
func (qs Lease6typesQS) Or(exprs ...Lease6typesQS) Lease6typesQS {
	var o models.OrFragment

	for _, expr := range exprs {
		o = append(o, expr.condFragments)
	}

	qs.condFragments = append(
		qs.condFragments,
		o,
	)

	return qs
}

// BEGIN - django_kea.Lease6Types.lease_type

// LeaseTypeEq filters for LeaseType being equal to argument
func (qs Lease6typesQS) LeaseTypeEq(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" =`, v)
}

// LeaseTypeNe filters for LeaseType being not equal to argument
func (qs Lease6typesQS) LeaseTypeNe(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" <>`, v)
}

// LeaseTypeLt filters for LeaseType being less than argument
func (qs Lease6typesQS) LeaseTypeLt(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" <`, v)
}

// LeaseTypeLe filters for LeaseType being less than or equal to argument
func (qs Lease6typesQS) LeaseTypeLe(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" <=`, v)
}

// LeaseTypeGt filters for LeaseType being greater than argument
func (qs Lease6typesQS) LeaseTypeGt(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" >`, v)
}

// LeaseTypeGe filters for LeaseType being greater than or equal to argument
func (qs Lease6typesQS) LeaseTypeGe(v int32) Lease6typesQS {
	return qs.filter(`"lease_type" >=`, v)
}

type inLease6typesLeaseType []interface{}

func (in inLease6typesLeaseType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"lease_type" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6typesQS) LeaseTypeIn(values []int32) Lease6typesQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6typesLeaseType(vals),
	)

	return qs
}

type notinLease6typesLeaseType []interface{}

func (in notinLease6typesLeaseType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"lease_type" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6typesQS) LeaseTypeNotIn(values []int32) Lease6typesQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6typesLeaseType(vals),
	)

	return qs
}

// OrderByLeaseType sorts result by LeaseType in ascending order
func (qs Lease6typesQS) OrderByLeaseType() Lease6typesQS {
	qs.order = append(qs.order, `"lease_type"`)

	return qs
}

// OrderByLeaseTypeDesc sorts result by LeaseType in descending order
func (qs Lease6typesQS) OrderByLeaseTypeDesc() Lease6typesQS {
	qs.order = append(qs.order, `"lease_type" DESC`)

	return qs
}

// DistinctOnLeaseType marks field in queries to add to DISTINCT ON clause
func (qs Lease6typesQS) DistinctOnLeaseType() Lease6typesQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"lease_type"`)

	return qs
}

// END - django_kea.Lease6Types.lease_type

// BEGIN - django_kea.Lease6Types.name

// NameIsNull filters for Name being null
func (qs Lease6typesQS) NameIsNull() Lease6typesQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NULL`,
		},
	)
	return qs
}

// NameIsNotNull filters for Name being not null
func (qs Lease6typesQS) NameIsNotNull() Lease6typesQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NOT NULL`,
		},
	)
	return qs
}

// NameEq filters for Name being equal to argument
func (qs Lease6typesQS) NameEq(v string) Lease6typesQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs Lease6typesQS) NameNe(v string) Lease6typesQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs Lease6typesQS) NameLt(v string) Lease6typesQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs Lease6typesQS) NameLe(v string) Lease6typesQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs Lease6typesQS) NameGt(v string) Lease6typesQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs Lease6typesQS) NameGe(v string) Lease6typesQS {
	return qs.filter(`"name" >=`, v)
}

type inLease6typesName []interface{}

func (in inLease6typesName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6typesQS) NameIn(values []string) Lease6typesQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6typesName(vals),
	)

	return qs
}

type notinLease6typesName []interface{}

func (in notinLease6typesName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6typesQS) NameNotIn(values []string) Lease6typesQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6typesName(vals),
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs Lease6typesQS) OrderByName() Lease6typesQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs Lease6typesQS) OrderByNameDesc() Lease6typesQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs Lease6typesQS) DistinctOnName() Lease6typesQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - django_kea.Lease6Types.name

// OrderByRandom randomizes result
func (qs Lease6typesQS) OrderByRandom() Lease6typesQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Lease6typesQS) ForUpdate() Lease6typesQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs Lease6typesQS) ForUpdateNowait() Lease6typesQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs Lease6typesQS) ForUpdateSkipLocked() Lease6typesQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs Lease6typesQS) ClearForUpdate() Lease6typesQS {
	qs.forClause = ""

	return qs
}

func (qs Lease6typesQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Lease6typesQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Lease6typesQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"lease_type", "name" FROM "lease6_types"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Lease6typesQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "lease_type" FROM "lease6_types"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs Lease6typesQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"lease_type"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "lease6_types"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs Lease6typesQS) All(ctx context.Context, db models.DBInterface) (Lease6typesList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret Lease6typesList
	for rows.Next() {
		obj := Lease6types{existsInDB: true}
		if err = rows.Scan(&obj.LeaseType, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Lease6typesQS) First(ctx context.Context, db models.DBInterface) (*Lease6types, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Lease6types{existsInDB: true}
	err := row.Scan(&obj.LeaseType, &obj.Name)
	switch err {
	case nil:
		return &obj, nil
	case pgx.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

// Delete deletes rows matching queryset filters
func (qs Lease6typesQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease6_types"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Lease6typesQS) Update() Lease6typesUpdateQS {
	return Lease6typesUpdateQS{condFragments: qs.condFragments}
}

// Lease6typesUpdateQS represents an updated queryset for django_kea.Lease6Types
type Lease6typesUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Lease6typesUpdateQS) update(c string, v interface{}) Lease6typesUpdateQS {
	var frag models.ConditionFragment

	if v == nil {
		frag = &models.ConstantFragment{
			Constant: c + " = NULL",
		}
	} else {
		frag = &models.UnaryFragment{
			Frag:  c + " =",
			Param: v,
		}
	}

	uqs.updates = append(uqs.updates, frag)

	return uqs
}

// SetLeaseType sets LeaseType to the given value
func (uqs Lease6typesUpdateQS) SetLeaseType(v int32) Lease6typesUpdateQS {
	return uqs.update(`"lease_type"`, v)
}

// SetName sets Name to the given value
func (uqs Lease6typesUpdateQS) SetName(v sql.NullString) Lease6typesUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs Lease6typesUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
	if len(uqs.updates) == 0 {
		return 0, nil
	}

	c := &models.PositionalCounter{}

	var params []interface{}

	var sets []string
	for _, set := range uqs.updates {
		s, p := set.GetConditionFragment(c)

		sets = append(sets, s)
		params = append(params, p...)
	}

	ws, wp := Lease6typesQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease6_types" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (l *Lease6types) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "lease6_types" ("name", "lease_type") VALUES ($1, $2)`, l.Name, l.LeaseType)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Lease6types) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "lease6_types" SET "name" = $1 WHERE "lease_type" = $2`, l.Name, l.LeaseType)

	return err
}

// Save inserts or updates record
func (l *Lease6types) Save(ctx context.Context, db models.DBInterface) error {
	if l.existsInDB {
		return l.update(ctx, db)
	}

	return l.insert(ctx, db)
}

// Delete removes row from database
func (l *Lease6types) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "lease6_types" WHERE "lease_type" = $1`, l.LeaseType)

	l.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (ll Lease6typesList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts Lease6typesList

	for _, l := range ll {
		if l.existsInDB {
			if err := l.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, l)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, l := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, l.Name, l.LeaseType)
		offs += 2
	}

	qs := `INSERT INTO "lease6_types" ("name", "lease_type") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, l := range inserts {
		l.existsInDB = true
	}

	return nil

}

// Lease6 returns the set of Lease6 referencing this Lease6types instance
func (l *Lease6types) Lease6() Lease6QS {
	return Lease6QS{}.LeaseTypeEq(l)
}
