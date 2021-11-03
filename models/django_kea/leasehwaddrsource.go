/*
  AUTO-GENERATED file for Django model django_kea.LeaseHwaddrSource

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

// Leasehwaddrsource mirrors model django_kea.LeaseHwaddrSource
type Leasehwaddrsource struct {
	existsInDB bool

	HwaddrSource int32
	Name         sql.NullString
}

// LeasehwaddrsourceQS represents a queryset for django_kea.LeaseHwaddrSource
type LeasehwaddrsourceQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
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

type inLeasehwaddrsourceHwaddrSource struct {
	values []interface{}
}

func (in *inLeasehwaddrsourceHwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasehwaddrsourceQS) HwaddrSourceIn(values []int32) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLeasehwaddrsourceHwaddrSource{
			values: vals,
		},
	)

	return qs
}

type notinLeasehwaddrsourceHwaddrSource struct {
	values []interface{}
}

func (in *notinLeasehwaddrsourceHwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasehwaddrsourceQS) HwaddrSourceNotIn(values []int32) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLeasehwaddrsourceHwaddrSource{
			values: vals,
		},
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

type inLeasehwaddrsourceName struct {
	values []interface{}
}

func (in *inLeasehwaddrsourceName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasehwaddrsourceQS) NameIn(values []string) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLeasehwaddrsourceName{
			values: vals,
		},
	)

	return qs
}

type notinLeasehwaddrsourceName struct {
	values []interface{}
}

func (in *notinLeasehwaddrsourceName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs LeasehwaddrsourceQS) NameNotIn(values []string) LeasehwaddrsourceQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLeasehwaddrsourceName{
			values: vals,
		},
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

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs LeasehwaddrsourceQS) ForUpdate() LeasehwaddrsourceQS {
	qs.forUpdate = true

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

func (qs LeasehwaddrsourceQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "hwaddr_source", "name" FROM "lease_hwaddr_source"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs LeasehwaddrsourceQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "hwaddr_source" FROM "lease_hwaddr_source"` + s, p
}

// All returns all rows matching queryset filters
func (qs LeasehwaddrsourceQS) All(db models.DBInterface) ([]*Leasehwaddrsource, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Leasehwaddrsource
	for rows.Next() {
		obj := Leasehwaddrsource{existsInDB: true}
		if err = rows.Scan(&obj.HwaddrSource, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs LeasehwaddrsourceQS) First(db models.DBInterface) (*Leasehwaddrsource, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Leasehwaddrsource{existsInDB: true}
	err := row.Scan(&obj.HwaddrSource, &obj.Name)
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
func (qs LeasehwaddrsourceQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease_hwaddr_source"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
func (uqs LeasehwaddrsourceUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (l *Leasehwaddrsource) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "lease_hwaddr_source" ("name", "hwaddr_source") VALUES ($1, $2)`, l.Name, l.HwaddrSource)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Leasehwaddrsource) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "lease_hwaddr_source" SET "name" = $1 WHERE "hwaddr_source" = $2`, l.Name, l.HwaddrSource)

	return err
}

// Save inserts or updates record
func (l *Leasehwaddrsource) Save(db models.DBInterface) error {
	if l.existsInDB {
		return l.update(db)
	}

	return l.insert(db)
}

// Delete removes row from database
func (l *Leasehwaddrsource) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "lease_hwaddr_source" WHERE "hwaddr_source" = $1`, l.HwaddrSource)

	l.existsInDB = false

	return err
}
