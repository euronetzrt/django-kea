// Code generated for Django model django_kea.HostIdentifierType. DO NOT EDIT.

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

// Hostidentifiertype mirrors model django_kea.HostIdentifierType
type Hostidentifiertype struct {
	existsInDB bool

	Type int32
	Name sql.NullString
}

// HostidentifiertypeList is a list of Hostidentifiertype
type HostidentifiertypeList []*Hostidentifiertype

// HostidentifiertypeQS represents a queryset for django_kea.HostIdentifierType
type HostidentifiertypeQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
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

// BEGIN - django_kea.HostIdentifierType.type

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

type inHostidentifiertypeType []interface{}

func (in inHostidentifiertypeType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"type" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostidentifiertypeQS) TypeIn(values []int32) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostidentifiertypeType(vals),
	)

	return qs
}

type notinHostidentifiertypeType []interface{}

func (in notinHostidentifiertypeType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"type" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostidentifiertypeQS) TypeNotIn(values []int32) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostidentifiertypeType(vals),
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

// DistinctOnType marks field in queries to add to DISTINCT ON clause
func (qs HostidentifiertypeQS) DistinctOnType() HostidentifiertypeQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"type"`)

	return qs
}

// END - django_kea.HostIdentifierType.type

// BEGIN - django_kea.HostIdentifierType.name

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

type inHostidentifiertypeName []interface{}

func (in inHostidentifiertypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostidentifiertypeQS) NameIn(values []string) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostidentifiertypeName(vals),
	)

	return qs
}

type notinHostidentifiertypeName []interface{}

func (in notinHostidentifiertypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostidentifiertypeQS) NameNotIn(values []string) HostidentifiertypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostidentifiertypeName(vals),
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

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs HostidentifiertypeQS) DistinctOnName() HostidentifiertypeQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - django_kea.HostIdentifierType.name

// OrderByRandom randomizes result
func (qs HostidentifiertypeQS) OrderByRandom() HostidentifiertypeQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs HostidentifiertypeQS) ForUpdate() HostidentifiertypeQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs HostidentifiertypeQS) ForUpdateNowait() HostidentifiertypeQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs HostidentifiertypeQS) ForUpdateSkipLocked() HostidentifiertypeQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs HostidentifiertypeQS) ClearForUpdate() HostidentifiertypeQS {
	qs.forClause = ""

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

func (qs HostidentifiertypeQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"type", "name" FROM "host_identifier_type"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs HostidentifiertypeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "type" FROM "host_identifier_type"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs HostidentifiertypeQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"type"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "host_identifier_type"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs HostidentifiertypeQS) All(ctx context.Context, db models.DBInterface) (HostidentifiertypeList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret HostidentifiertypeList
	for rows.Next() {
		obj := Hostidentifiertype{existsInDB: true}
		if err = rows.Scan(&obj.Type, &obj.Name); err != nil {
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
func (qs HostidentifiertypeQS) First(ctx context.Context, db models.DBInterface) (*Hostidentifiertype, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Hostidentifiertype{existsInDB: true}
	err := row.Scan(&obj.Type, &obj.Name)
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
func (qs HostidentifiertypeQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "host_identifier_type"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs HostidentifiertypeUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (h *Hostidentifiertype) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "host_identifier_type" ("name", "type") VALUES ($1, $2)`, h.Name, h.Type)

	if err != nil {
		return err
	}

	h.existsInDB = true

	return nil
}

// update operation
func (h *Hostidentifiertype) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "host_identifier_type" SET "name" = $1 WHERE "type" = $2`, h.Name, h.Type)

	return err
}

// Save inserts or updates record
func (h *Hostidentifiertype) Save(ctx context.Context, db models.DBInterface) error {
	if h.existsInDB {
		return h.update(ctx, db)
	}

	return h.insert(ctx, db)
}

// Delete removes row from database
func (h *Hostidentifiertype) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "host_identifier_type" WHERE "type" = $1`, h.Type)

	h.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (hl HostidentifiertypeList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts HostidentifiertypeList

	for _, h := range hl {
		if h.existsInDB {
			if err := h.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, h)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, h := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, h.Name, h.Type)
		offs += 2
	}

	qs := `INSERT INTO "host_identifier_type" ("name", "type") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, h := range inserts {
		h.existsInDB = true
	}

	return nil

}

// Hosts returns the set of Hosts referencing this Hostidentifiertype instance
func (h *Hostidentifiertype) Hosts() HostsQS {
	return HostsQS{}.DhcpIdentifierTypeEq(h)
}
