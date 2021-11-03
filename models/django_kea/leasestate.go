/*
  AUTO-GENERATED file for Django model django_kea.LeaseState

  Command used to generate:

  DJANGO_SETTINGS_MODULE=keatest.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/euronetzrt/django-kea django_kea

  https://github.com/rkojedzinszky/djan-go-rm
*/

package django_kea

import (
	"database/sql"
	"github.com/euronetzrt/django-kea/models"
	"strings"
)

// Leasestate mirrors model django_kea.LeaseState
type Leasestate struct {
	existsInDB bool

	State int64
	Name  string
}

// LeasestateQS represents a queryset for django_kea.LeaseState
type LeasestateQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
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

type inLeasestateState struct {
	values []interface{}
}

func (in *inLeasestateState) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"state" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasestateQS) StateIn(values []int64) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLeasestateState{
			values: vals,
		},
	)

	return qs
}

type notinLeasestateState struct {
	values []interface{}
}

func (in *notinLeasestateState) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"state" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasestateQS) StateNotIn(values []int64) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLeasestateState{
			values: vals,
		},
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

type inLeasestateName struct {
	values []interface{}
}

func (in *inLeasestateName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasestateQS) NameIn(values []string) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLeasestateName{
			values: vals,
		},
	)

	return qs
}

type notinLeasestateName struct {
	values []interface{}
}

func (in *notinLeasestateName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasestateQS) NameNotIn(values []string) LeasestateQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLeasestateName{
			values: vals,
		},
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

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs LeasestateQS) ForUpdate() LeasestateQS {
	qs.forUpdate = true

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

func (qs LeasestateQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "state", "name" FROM "lease_state"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs LeasestateQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "state" FROM "lease_state"` + s, p
}

// All returns all rows matching queryset filters
func (qs LeasestateQS) All(db models.DBInterface) ([]*Leasestate, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Leasestate
	for rows.Next() {
		obj := Leasestate{existsInDB: true}
		if err = rows.Scan(&obj.State, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs LeasestateQS) First(db models.DBInterface) (*Leasestate, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Leasestate{existsInDB: true}
	err := row.Scan(&obj.State, &obj.Name)
	switch err {
	case nil:
		return &obj, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

// Delete deletes rows matching queryset filters
func (qs LeasestateQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease_state"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
func (uqs LeasestateUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (l *Leasestate) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "lease_state" ("name", "state") VALUES ($1, $2)`, l.Name, l.State)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Leasestate) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "lease_state" SET "name" = $1 WHERE "state" = $2`, l.Name, l.State)

	return err
}

// Save inserts or updates record
func (l *Leasestate) Save(db models.DBInterface) error {
	if l.existsInDB {
		return l.update(db)
	}

	return l.insert(db)
}

// Delete removes row from database
func (l *Leasestate) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "lease_state" WHERE "state" = $1`, l.State)

	l.existsInDB = false

	return err
}

// Lease4 returns the set of Lease4 referencing this Leasestate instance
func (l *Leasestate) Lease4() Lease4QS {
	return Lease4QS{}.StateEq(l)
}

// Lease6 returns the set of Lease6 referencing this Leasestate instance
func (l *Leasestate) Lease6() Lease6QS {
	return Lease6QS{}.StateEq(l)
}
