// Code generated for Django model django_kea.LeaseHwaddrSource. DO NOT EDIT.

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

// Leasehwaddrsource mirrors model django_kea.LeaseHwaddrSource
type Leasehwaddrsource struct {
	existsInDB bool

	HwaddrSource int32
	Name         sql.NullString
}

// LeasehwaddrsourceList is a list of Leasehwaddrsource
type LeasehwaddrsourceList []*Leasehwaddrsource

// LeasehwaddrsourceQS represents a queryset for django_kea.LeaseHwaddrSource
type LeasehwaddrsourceQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs LeasehwaddrsourceQS) filter(c string, p interface{}) LeasehwaddrsourceQS {
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
func (qs LeasehwaddrsourceQS) Or(exprs ...LeasehwaddrsourceQS) LeasehwaddrsourceQS {
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

// BEGIN - django_kea.LeaseHwaddrSource.hwaddr_source

// HwaddrSourceEq filters for HwaddrSource being equal to argument
func (qs LeasehwaddrsourceQS) HwaddrSourceEq(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" =`, v)
}

// HwaddrSourceNe filters for HwaddrSource being not equal to argument
func (qs LeasehwaddrsourceQS) HwaddrSourceNe(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" <>`, v)
}

// HwaddrSourceLt filters for HwaddrSource being less than argument
func (qs LeasehwaddrsourceQS) HwaddrSourceLt(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" <`, v)
}

// HwaddrSourceLe filters for HwaddrSource being less than or equal to argument
func (qs LeasehwaddrsourceQS) HwaddrSourceLe(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" <=`, v)
}

// HwaddrSourceGt filters for HwaddrSource being greater than argument
func (qs LeasehwaddrsourceQS) HwaddrSourceGt(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" >`, v)
}

// HwaddrSourceGe filters for HwaddrSource being greater than or equal to argument
func (qs LeasehwaddrsourceQS) HwaddrSourceGe(v int32) LeasehwaddrsourceQS {
	return qs.filter(`"hwaddr_source" >=`, v)
}

type inLeasehwaddrsourceHwaddrSource []interface{}

func (in inLeasehwaddrsourceHwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasehwaddrsourceQS) HwaddrSourceIn(values []int32) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLeasehwaddrsourceHwaddrSource(vals),
	)

	return qs
}

type notinLeasehwaddrsourceHwaddrSource []interface{}

func (in notinLeasehwaddrsourceHwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasehwaddrsourceQS) HwaddrSourceNotIn(values []int32) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLeasehwaddrsourceHwaddrSource(vals),
	)

	return qs
}

// OrderByHwaddrSource sorts result by HwaddrSource in ascending order
func (qs LeasehwaddrsourceQS) OrderByHwaddrSource() LeasehwaddrsourceQS {
	qs.order = append(qs.order, `"hwaddr_source"`)

	return qs
}

// OrderByHwaddrSourceDesc sorts result by HwaddrSource in descending order
func (qs LeasehwaddrsourceQS) OrderByHwaddrSourceDesc() LeasehwaddrsourceQS {
	qs.order = append(qs.order, `"hwaddr_source" DESC`)

	return qs
}

// DistinctOnHwaddrSource marks field in queries to add to DISTINCT ON clause
func (qs LeasehwaddrsourceQS) DistinctOnHwaddrSource() LeasehwaddrsourceQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hwaddr_source"`)

	return qs
}

// END - django_kea.LeaseHwaddrSource.hwaddr_source

// BEGIN - django_kea.LeaseHwaddrSource.name

// NameIsNull filters for Name being null
func (qs LeasehwaddrsourceQS) NameIsNull() LeasehwaddrsourceQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NULL`,
		},
	)
	return qs
}

// NameIsNotNull filters for Name being not null
func (qs LeasehwaddrsourceQS) NameIsNotNull() LeasehwaddrsourceQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NOT NULL`,
		},
	)
	return qs
}

// NameEq filters for Name being equal to argument
func (qs LeasehwaddrsourceQS) NameEq(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs LeasehwaddrsourceQS) NameNe(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs LeasehwaddrsourceQS) NameLt(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs LeasehwaddrsourceQS) NameLe(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs LeasehwaddrsourceQS) NameGt(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs LeasehwaddrsourceQS) NameGe(v string) LeasehwaddrsourceQS {
	return qs.filter(`"name" >=`, v)
}

type inLeasehwaddrsourceName []interface{}

func (in inLeasehwaddrsourceName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasehwaddrsourceQS) NameIn(values []string) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLeasehwaddrsourceName(vals),
	)

	return qs
}

type notinLeasehwaddrsourceName []interface{}

func (in notinLeasehwaddrsourceName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasehwaddrsourceQS) NameNotIn(values []string) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLeasehwaddrsourceName(vals),
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs LeasehwaddrsourceQS) OrderByName() LeasehwaddrsourceQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs LeasehwaddrsourceQS) OrderByNameDesc() LeasehwaddrsourceQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs LeasehwaddrsourceQS) DistinctOnName() LeasehwaddrsourceQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - django_kea.LeaseHwaddrSource.name

// OrderByRandom randomizes result
func (qs LeasehwaddrsourceQS) OrderByRandom() LeasehwaddrsourceQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs LeasehwaddrsourceQS) ForUpdate() LeasehwaddrsourceQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs LeasehwaddrsourceQS) ForUpdateNowait() LeasehwaddrsourceQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs LeasehwaddrsourceQS) ForUpdateSkipLocked() LeasehwaddrsourceQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs LeasehwaddrsourceQS) ClearForUpdate() LeasehwaddrsourceQS {
	qs.forClause = ""

	return qs
}

func (qs LeasehwaddrsourceQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs LeasehwaddrsourceQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs LeasehwaddrsourceQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"hwaddr_source", "name" FROM "lease_hwaddr_source"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs LeasehwaddrsourceQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "hwaddr_source" FROM "lease_hwaddr_source"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs LeasehwaddrsourceQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"hwaddr_source"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "lease_hwaddr_source"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs LeasehwaddrsourceQS) All(ctx context.Context, db models.DBInterface) (LeasehwaddrsourceList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret LeasehwaddrsourceList
	for rows.Next() {
		obj := Leasehwaddrsource{existsInDB: true}
		if err = rows.Scan(&obj.HwaddrSource, &obj.Name); err != nil {
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
func (qs LeasehwaddrsourceQS) First(ctx context.Context, db models.DBInterface) (*Leasehwaddrsource, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Leasehwaddrsource{existsInDB: true}
	err := row.Scan(&obj.HwaddrSource, &obj.Name)
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
func (qs LeasehwaddrsourceQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease_hwaddr_source"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs LeasehwaddrsourceQS) Update() LeasehwaddrsourceUpdateQS {
	return LeasehwaddrsourceUpdateQS{condFragments: qs.condFragments}
}

// LeasehwaddrsourceUpdateQS represents an updated queryset for django_kea.LeaseHwaddrSource
type LeasehwaddrsourceUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs LeasehwaddrsourceUpdateQS) update(c string, v interface{}) LeasehwaddrsourceUpdateQS {
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

// SetHwaddrSource sets HwaddrSource to the given value
func (uqs LeasehwaddrsourceUpdateQS) SetHwaddrSource(v int32) LeasehwaddrsourceUpdateQS {
	return uqs.update(`"hwaddr_source"`, v)
}

// SetName sets Name to the given value
func (uqs LeasehwaddrsourceUpdateQS) SetName(v sql.NullString) LeasehwaddrsourceUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs LeasehwaddrsourceUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := LeasehwaddrsourceQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease_hwaddr_source" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (l *Leasehwaddrsource) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "lease_hwaddr_source" ("name", "hwaddr_source") VALUES ($1, $2)`, l.Name, l.HwaddrSource)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Leasehwaddrsource) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "lease_hwaddr_source" SET "name" = $1 WHERE "hwaddr_source" = $2`, l.Name, l.HwaddrSource)

	return err
}

// Save inserts or updates record
func (l *Leasehwaddrsource) Save(ctx context.Context, db models.DBInterface) error {
	if l.existsInDB {
		return l.update(ctx, db)
	}

	return l.insert(ctx, db)
}

// Delete removes row from database
func (l *Leasehwaddrsource) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "lease_hwaddr_source" WHERE "hwaddr_source" = $1`, l.HwaddrSource)

	l.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (ll LeasehwaddrsourceList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts LeasehwaddrsourceList

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
		vaa = append(vaa, l.Name, l.HwaddrSource)
		offs += 2
	}

	qs := `INSERT INTO "lease_hwaddr_source" ("name", "hwaddr_source") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, l := range inserts {
		l.existsInDB = true
	}

	return nil

}

// Lease6 returns the set of Lease6 referencing this Leasehwaddrsource instance
func (l *Leasehwaddrsource) Lease6() Lease6QS {
	return Lease6QS{}.HwaddrSourceEq(l)
}
