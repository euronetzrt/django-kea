/*
  AUTO-GENERATED file for Django model django_kea.HostIdentifierType

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

// Hostidentifiertype mirrors model django_kea.HostIdentifierType
type Hostidentifiertype struct {
	existsInDB bool

	Type int32
	Name sql.NullString
}

// HostidentifiertypeQS represents a queryset for django_kea.HostIdentifierType
type HostidentifiertypeQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs HostidentifiertypeQS) filter(c string, p interface{}) HostidentifiertypeQS {
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
func (qs HostidentifiertypeQS) Or(exprs ...HostidentifiertypeQS) HostidentifiertypeQS {
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

// TypeEq filters for Type being equal to argument
func (qs HostidentifiertypeQS) TypeEq(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" =`, v)
}

// TypeNe filters for Type being not equal to argument
func (qs HostidentifiertypeQS) TypeNe(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" <>`, v)
}

// TypeLt filters for Type being less than argument
func (qs HostidentifiertypeQS) TypeLt(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" <`, v)
}

// TypeLe filters for Type being less than or equal to argument
func (qs HostidentifiertypeQS) TypeLe(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" <=`, v)
}

// TypeGt filters for Type being greater than argument
func (qs HostidentifiertypeQS) TypeGt(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" >`, v)
}

// TypeGe filters for Type being greater than or equal to argument
func (qs HostidentifiertypeQS) TypeGe(v int32) HostidentifiertypeQS {
	return qs.filter(`"type" >=`, v)
}

type inHostidentifiertypeType struct {
	values []interface{}
}

func (in *inHostidentifiertypeType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"type" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostidentifiertypeQS) TypeIn(values []int32) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostidentifiertypeType{
			values: vals,
		},
	)

	return qs
}

type notinHostidentifiertypeType struct {
	values []interface{}
}

func (in *notinHostidentifiertypeType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"type" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostidentifiertypeQS) TypeNotIn(values []int32) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostidentifiertypeType{
			values: vals,
		},
	)

	return qs
}

// OrderByType sorts result by Type in ascending order
func (qs HostidentifiertypeQS) OrderByType() HostidentifiertypeQS {
	qs.order = append(qs.order, `"type"`)

	return qs
}

// OrderByTypeDesc sorts result by Type in descending order
func (qs HostidentifiertypeQS) OrderByTypeDesc() HostidentifiertypeQS {
	qs.order = append(qs.order, `"type" DESC`)

	return qs
}

// NameIsNull filters for Name being null
func (qs HostidentifiertypeQS) NameIsNull() HostidentifiertypeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NULL`,
		},
	)
	return qs
}

// NameIsNotNull filters for Name being not null
func (qs HostidentifiertypeQS) NameIsNotNull() HostidentifiertypeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"name" IS NOT NULL`,
		},
	)
	return qs
}

// NameEq filters for Name being equal to argument
func (qs HostidentifiertypeQS) NameEq(v string) HostidentifiertypeQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs HostidentifiertypeQS) NameNe(v string) HostidentifiertypeQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs HostidentifiertypeQS) NameLt(v string) HostidentifiertypeQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs HostidentifiertypeQS) NameLe(v string) HostidentifiertypeQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs HostidentifiertypeQS) NameGt(v string) HostidentifiertypeQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs HostidentifiertypeQS) NameGe(v string) HostidentifiertypeQS {
	return qs.filter(`"name" >=`, v)
}

type inHostidentifiertypeName struct {
	values []interface{}
}

func (in *inHostidentifiertypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostidentifiertypeQS) NameIn(values []string) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostidentifiertypeName{
			values: vals,
		},
	)

	return qs
}

type notinHostidentifiertypeName struct {
	values []interface{}
}

func (in *notinHostidentifiertypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostidentifiertypeQS) NameNotIn(values []string) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostidentifiertypeName{
			values: vals,
		},
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs HostidentifiertypeQS) OrderByName() HostidentifiertypeQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs HostidentifiertypeQS) OrderByNameDesc() HostidentifiertypeQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs HostidentifiertypeQS) ForUpdate() HostidentifiertypeQS {
	qs.forUpdate = true

	return qs
}

func (qs HostidentifiertypeQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs HostidentifiertypeQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs HostidentifiertypeQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "type", "name" FROM "host_identifier_type"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs HostidentifiertypeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "type" FROM "host_identifier_type"` + s, p
}

// All returns all rows matching queryset filters
func (qs HostidentifiertypeQS) All(db models.DBInterface) ([]*Hostidentifiertype, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Hostidentifiertype
	for rows.Next() {
		obj := Hostidentifiertype{existsInDB: true}
		if err = rows.Scan(&obj.Type, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs HostidentifiertypeQS) First(db models.DBInterface) (*Hostidentifiertype, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Hostidentifiertype{existsInDB: true}
	err := row.Scan(&obj.Type, &obj.Name)
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
func (qs HostidentifiertypeQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "host_identifier_type"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs HostidentifiertypeQS) Update() HostidentifiertypeUpdateQS {
	return HostidentifiertypeUpdateQS{condFragments: qs.condFragments}
}

// HostidentifiertypeUpdateQS represents an updated queryset for django_kea.HostIdentifierType
type HostidentifiertypeUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs HostidentifiertypeUpdateQS) update(c string, v interface{}) HostidentifiertypeUpdateQS {
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

// SetType sets Type to the given value
func (uqs HostidentifiertypeUpdateQS) SetType(v int32) HostidentifiertypeUpdateQS {
	return uqs.update(`"type"`, v)
}

// SetName sets Name to the given value
func (uqs HostidentifiertypeUpdateQS) SetName(v sql.NullString) HostidentifiertypeUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs HostidentifiertypeUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := HostidentifiertypeQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "host_identifier_type" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (h *Hostidentifiertype) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "host_identifier_type" ("name", "type") VALUES ($1, $2)`, h.Name, h.Type)

	if err != nil {
		return err
	}

	h.existsInDB = true

	return nil
}

// update operation
func (h *Hostidentifiertype) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "host_identifier_type" SET "name" = $1 WHERE "type" = $2`, h.Name, h.Type)

	return err
}

// Save inserts or updates record
func (h *Hostidentifiertype) Save(db models.DBInterface) error {
	if h.existsInDB {
		return h.update(db)
	}

	return h.insert(db)
}

// Delete removes row from database
func (h *Hostidentifiertype) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "host_identifier_type" WHERE "type" = $1`, h.Type)

	h.existsInDB = false

	return err
}

// Host returns the set of Host referencing this Hostidentifiertype instance
func (h *Hostidentifiertype) Host() HostQS {
	return HostQS{}.DhcpIdentifierTypeEq(h)
}
