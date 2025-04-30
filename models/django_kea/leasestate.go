// Code generated for Django model django_kea.LeaseState. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=keatest.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/euronetzrt/django-kea django_kea

  https://github.com/rkojedzinszky/djan-go-rm
*/

package django_kea

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/euronetzrt/django-kea/models"
)

// Leasestate mirrors model django_kea.LeaseState
type Leasestate struct {
	existsInDB bool

	State int64
	Name  string
}

// LeasestateList is a list of Leasestate
type LeasestateList []*Leasestate

// LeasestateQS represents a queryset for django_kea.LeaseState
type LeasestateQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs LeasestateQS) filter(c string, p interface{}) LeasestateQS {
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
func (qs LeasestateQS) Or(exprs ...LeasestateQS) LeasestateQS {
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

// BEGIN - django_kea.LeaseState.state

// StateEq filters for State being equal to argument
func (qs LeasestateQS) StateEq(v int64) LeasestateQS {
	return qs.filter(`"state" =`, v)
}

// StateNe filters for State being not equal to argument
func (qs LeasestateQS) StateNe(v int64) LeasestateQS {
	return qs.filter(`"state" <>`, v)
}

// StateLt filters for State being less than argument
func (qs LeasestateQS) StateLt(v int64) LeasestateQS {
	return qs.filter(`"state" <`, v)
}

// StateLe filters for State being less than or equal to argument
func (qs LeasestateQS) StateLe(v int64) LeasestateQS {
	return qs.filter(`"state" <=`, v)
}

// StateGt filters for State being greater than argument
func (qs LeasestateQS) StateGt(v int64) LeasestateQS {
	return qs.filter(`"state" >`, v)
}

// StateGe filters for State being greater than or equal to argument
func (qs LeasestateQS) StateGe(v int64) LeasestateQS {
	return qs.filter(`"state" >=`, v)
}

type inLeasestateState []interface{}

func (in inLeasestateState) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"state" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasestateQS) StateIn(values []int64) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLeasestateState(vals),
	)

	return qs
}

type notinLeasestateState []interface{}

func (in notinLeasestateState) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"state" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasestateQS) StateNotIn(values []int64) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLeasestateState(vals),
	)

	return qs
}

// OrderByState sorts result by State in ascending order
func (qs LeasestateQS) OrderByState() LeasestateQS {
	qs.order = append(qs.order, `"state"`)

	return qs
}

// OrderByStateDesc sorts result by State in descending order
func (qs LeasestateQS) OrderByStateDesc() LeasestateQS {
	qs.order = append(qs.order, `"state" DESC`)

	return qs
}

// DistinctOnState marks field in queries to add to DISTINCT ON clause
func (qs LeasestateQS) DistinctOnState() LeasestateQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"state"`)

	return qs
}

// END - django_kea.LeaseState.state

// BEGIN - django_kea.LeaseState.name

// NameEq filters for Name being equal to argument
func (qs LeasestateQS) NameEq(v string) LeasestateQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs LeasestateQS) NameNe(v string) LeasestateQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs LeasestateQS) NameLt(v string) LeasestateQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs LeasestateQS) NameLe(v string) LeasestateQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs LeasestateQS) NameGt(v string) LeasestateQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs LeasestateQS) NameGe(v string) LeasestateQS {
	return qs.filter(`"name" >=`, v)
}

type inLeasestateName []interface{}

func (in inLeasestateName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasestateQS) NameIn(values []string) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLeasestateName(vals),
	)

	return qs
}

type notinLeasestateName []interface{}

func (in notinLeasestateName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs LeasestateQS) NameNotIn(values []string) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLeasestateName(vals),
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs LeasestateQS) OrderByName() LeasestateQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs LeasestateQS) OrderByNameDesc() LeasestateQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs LeasestateQS) DistinctOnName() LeasestateQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - django_kea.LeaseState.name

// OrderByRandom randomizes result
func (qs LeasestateQS) OrderByRandom() LeasestateQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs LeasestateQS) ForUpdate() LeasestateQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs LeasestateQS) ForUpdateNowait() LeasestateQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs LeasestateQS) ForUpdateSkipLocked() LeasestateQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs LeasestateQS) ClearForUpdate() LeasestateQS {
	qs.forClause = ""

	return qs
}

func (qs LeasestateQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs LeasestateQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs LeasestateQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"state", "name" FROM "lease_state"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs LeasestateQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "state" FROM "lease_state"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs LeasestateQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"state"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "lease_state"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs LeasestateQS) All(ctx context.Context, db models.DBInterface) (LeasestateList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret LeasestateList
	for rows.Next() {
		obj := Leasestate{existsInDB: true}
		if err = rows.Scan(&obj.State, &obj.Name); err != nil {
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
func (qs LeasestateQS) First(ctx context.Context, db models.DBInterface) (*Leasestate, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Leasestate{existsInDB: true}
	err := row.Scan(&obj.State, &obj.Name)
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
func (qs LeasestateQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease_state"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs LeasestateQS) Update() LeasestateUpdateQS {
	return LeasestateUpdateQS{condFragments: qs.condFragments}
}

// LeasestateUpdateQS represents an updated queryset for django_kea.LeaseState
type LeasestateUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs LeasestateUpdateQS) update(c string, v interface{}) LeasestateUpdateQS {
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

// SetState sets State to the given value
func (uqs LeasestateUpdateQS) SetState(v int64) LeasestateUpdateQS {
	return uqs.update(`"state"`, v)
}

// SetName sets Name to the given value
func (uqs LeasestateUpdateQS) SetName(v string) LeasestateUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs LeasestateUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := LeasestateQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease_state" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (l *Leasestate) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "lease_state" ("name", "state") VALUES ($1, $2)`, l.Name, l.State)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Leasestate) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "lease_state" SET "name" = $1 WHERE "state" = $2`, l.Name, l.State)

	return err
}

// Save inserts or updates record
func (l *Leasestate) Save(ctx context.Context, db models.DBInterface) error {
	if l.existsInDB {
		return l.update(ctx, db)
	}

	return l.insert(ctx, db)
}

// Delete removes row from database
func (l *Leasestate) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "lease_state" WHERE "state" = $1`, l.State)

	l.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (ll LeasestateList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts LeasestateList

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
		vaa = append(vaa, l.Name, l.State)
		offs += 2
	}

	qs := `INSERT INTO "lease_state" ("name", "state") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, l := range inserts {
		l.existsInDB = true
	}

	return nil

}

// Lease4 returns the set of Lease4 referencing this Leasestate instance
func (l *Leasestate) Lease4() Lease4QS {
	return Lease4QS{}.StateEq(l)
}

// Lease6 returns the set of Lease6 referencing this Leasestate instance
func (l *Leasestate) Lease6() Lease6QS {
	return Lease6QS{}.StateEq(l)
}
