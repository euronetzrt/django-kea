/*
  AUTO-GENERATED file for Django model django_kea.Lease6Type

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

// Lease6type mirrors model django_kea.Lease6Type
type Lease6type struct {
	existsInDB bool

	LeaseType int32
	Name      sql.NullString
}

// Lease6typeQS represents a queryset for django_kea.Lease6Type
type Lease6typeQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Lease6typeQS) filter(c string, p interface{}) Lease6typeQS {
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
func (qs Lease6typeQS) Or(exprs ...Lease6typeQS) Lease6typeQS {
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

// LeaseTypeEq filters for LeaseType being equal to argument
func (qs Lease6typeQS) LeaseTypeEq(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" =`, v)
}

// LeaseTypeNe filters for LeaseType being not equal to argument
func (qs Lease6typeQS) LeaseTypeNe(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" <>`, v)
}

// LeaseTypeLt filters for LeaseType being less than argument
func (qs Lease6typeQS) LeaseTypeLt(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" <`, v)
}

// LeaseTypeLe filters for LeaseType being less than or equal to argument
func (qs Lease6typeQS) LeaseTypeLe(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" <=`, v)
}

// LeaseTypeGt filters for LeaseType being greater than argument
func (qs Lease6typeQS) LeaseTypeGt(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" >`, v)
}

// LeaseTypeGe filters for LeaseType being greater than or equal to argument
func (qs Lease6typeQS) LeaseTypeGe(v int32) Lease6typeQS {
	return qs.filter(`"lease_type" >=`, v)
}

type inLease6typeLeaseType struct {
	values []interface{}
}

func (in *inLease6typeLeaseType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"lease_type" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6typeQS) LeaseTypeIn(values []int32) Lease6typeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6typeLeaseType{
			values: vals,
		},
	)

	return qs
}

type notinLease6typeLeaseType struct {
	values []interface{}
}

func (in *notinLease6typeLeaseType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"lease_type" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6typeQS) LeaseTypeNotIn(values []int32) Lease6typeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6typeLeaseType{
			values: vals,
		},
	)

	return qs
}

// OrderByLeaseType sorts result by LeaseType in ascending order
func (qs Lease6typeQS) OrderByLeaseType() Lease6typeQS {
	qs.order = append(qs.order, `"lease_type"`)

	return qs
}

// OrderByLeaseTypeDesc sorts result by LeaseType in descending order
func (qs Lease6typeQS) OrderByLeaseTypeDesc() Lease6typeQS {
	qs.order = append(qs.order, `"lease_type" DESC`)

	return qs
}

// NameIsNull filters for Name being null
func (qs Lease6typeQS) NameIsNull() Lease6typeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NULL`,
		},
	)
	return qs
}

// NameIsNotNull filters for Name being not null
func (qs Lease6typeQS) NameIsNotNull() Lease6typeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NOT NULL`,
		},
	)
	return qs
}

// NameEq filters for Name being equal to argument
func (qs Lease6typeQS) NameEq(v string) Lease6typeQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs Lease6typeQS) NameNe(v string) Lease6typeQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs Lease6typeQS) NameLt(v string) Lease6typeQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs Lease6typeQS) NameLe(v string) Lease6typeQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs Lease6typeQS) NameGt(v string) Lease6typeQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs Lease6typeQS) NameGe(v string) Lease6typeQS {
	return qs.filter(`"name" >=`, v)
}

type inLease6typeName struct {
	values []interface{}
}

func (in *inLease6typeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6typeQS) NameIn(values []string) Lease6typeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6typeName{
			values: vals,
		},
	)

	return qs
}

type notinLease6typeName struct {
	values []interface{}
}

func (in *notinLease6typeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6typeQS) NameNotIn(values []string) Lease6typeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6typeName{
			values: vals,
		},
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs Lease6typeQS) OrderByName() Lease6typeQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs Lease6typeQS) OrderByNameDesc() Lease6typeQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Lease6typeQS) ForUpdate() Lease6typeQS {
	qs.forUpdate = true

	return qs
}

func (qs Lease6typeQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Lease6typeQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Lease6typeQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "lease_type", "name" FROM "lease6_types"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Lease6typeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "lease_type" FROM "lease6_types"` + s, p
}

// All returns all rows matching queryset filters
func (qs Lease6typeQS) All(db models.DBInterface) ([]*Lease6type, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Lease6type
	for rows.Next() {
		obj := Lease6type{existsInDB: true}
		if err = rows.Scan(&obj.LeaseType, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Lease6typeQS) First(db models.DBInterface) (*Lease6type, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Lease6type{existsInDB: true}
	err := row.Scan(&obj.LeaseType, &obj.Name)
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
func (qs Lease6typeQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease6_types"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Lease6typeQS) Update() Lease6typeUpdateQS {
	return Lease6typeUpdateQS{condFragments: qs.condFragments}
}

// Lease6typeUpdateQS represents an updated queryset for django_kea.Lease6Type
type Lease6typeUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Lease6typeUpdateQS) update(c string, v interface{}) Lease6typeUpdateQS {
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
func (uqs Lease6typeUpdateQS) SetLeaseType(v int32) Lease6typeUpdateQS {
	return uqs.update(`"lease_type"`, v)
}

// SetName sets Name to the given value
func (uqs Lease6typeUpdateQS) SetName(v sql.NullString) Lease6typeUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs Lease6typeUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Lease6typeQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease6_types" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (l *Lease6type) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "lease6_types" ("name", "lease_type") VALUES ($1, $2)`, l.Name, l.LeaseType)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Lease6type) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "lease6_types" SET "name" = $1 WHERE "lease_type" = $2`, l.Name, l.LeaseType)

	return err
}

// Save inserts or updates record
func (l *Lease6type) Save(db models.DBInterface) error {
	if l.existsInDB {
		return l.update(db)
	}

	return l.insert(db)
}

// Delete removes row from database
func (l *Lease6type) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "lease6_types" WHERE "lease_type" = $1`, l.LeaseType)

	l.existsInDB = false

	return err
}

// Lease6 returns the set of Lease6 referencing this Lease6type instance
func (l *Lease6type) Lease6() Lease6QS {
	return Lease6QS{}.LeaseTypeEq(l)
}
