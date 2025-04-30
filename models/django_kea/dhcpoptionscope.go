// Code generated for Django model django_kea.DhcpOptionScope. DO NOT EDIT.

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

// Dhcpoptionscope mirrors model django_kea.DhcpOptionScope
type Dhcpoptionscope struct {
	existsInDB bool

	ScopeId   int32
	ScopeName sql.NullString
}

// DhcpoptionscopeList is a list of Dhcpoptionscope
type DhcpoptionscopeList []*Dhcpoptionscope

// DhcpoptionscopeQS represents a queryset for django_kea.DhcpOptionScope
type DhcpoptionscopeQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs DhcpoptionscopeQS) filter(c string, p interface{}) DhcpoptionscopeQS {
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
func (qs DhcpoptionscopeQS) Or(exprs ...DhcpoptionscopeQS) DhcpoptionscopeQS {
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

// BEGIN - django_kea.DhcpOptionScope.scope_id

// ScopeIdEq filters for ScopeId being equal to argument
func (qs DhcpoptionscopeQS) ScopeIdEq(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" =`, v)
}

// ScopeIdNe filters for ScopeId being not equal to argument
func (qs DhcpoptionscopeQS) ScopeIdNe(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" <>`, v)
}

// ScopeIdLt filters for ScopeId being less than argument
func (qs DhcpoptionscopeQS) ScopeIdLt(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" <`, v)
}

// ScopeIdLe filters for ScopeId being less than or equal to argument
func (qs DhcpoptionscopeQS) ScopeIdLe(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" <=`, v)
}

// ScopeIdGt filters for ScopeId being greater than argument
func (qs DhcpoptionscopeQS) ScopeIdGt(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" >`, v)
}

// ScopeIdGe filters for ScopeId being greater than or equal to argument
func (qs DhcpoptionscopeQS) ScopeIdGe(v int32) DhcpoptionscopeQS {
	return qs.filter(`"scope_id" >=`, v)
}

type inDhcpoptionscopeScopeId []interface{}

func (in inDhcpoptionscopeScopeId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"scope_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DhcpoptionscopeQS) ScopeIdIn(values []int32) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcpoptionscopeScopeId(vals),
	)

	return qs
}

type notinDhcpoptionscopeScopeId []interface{}

func (in notinDhcpoptionscopeScopeId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"scope_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DhcpoptionscopeQS) ScopeIdNotIn(values []int32) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcpoptionscopeScopeId(vals),
	)

	return qs
}

// OrderByScopeId sorts result by ScopeId in ascending order
func (qs DhcpoptionscopeQS) OrderByScopeId() DhcpoptionscopeQS {
	qs.order = append(qs.order, `"scope_id"`)

	return qs
}

// OrderByScopeIdDesc sorts result by ScopeId in descending order
func (qs DhcpoptionscopeQS) OrderByScopeIdDesc() DhcpoptionscopeQS {
	qs.order = append(qs.order, `"scope_id" DESC`)

	return qs
}

// DistinctOnScopeId marks field in queries to add to DISTINCT ON clause
func (qs DhcpoptionscopeQS) DistinctOnScopeId() DhcpoptionscopeQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"scope_id"`)

	return qs
}

// END - django_kea.DhcpOptionScope.scope_id

// BEGIN - django_kea.DhcpOptionScope.scope_name

// ScopeNameIsNull filters for ScopeName being null
func (qs DhcpoptionscopeQS) ScopeNameIsNull() DhcpoptionscopeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"scope_name" IS NULL`,
		},
	)
	return qs
}

// ScopeNameIsNotNull filters for ScopeName being not null
func (qs DhcpoptionscopeQS) ScopeNameIsNotNull() DhcpoptionscopeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"scope_name" IS NOT NULL`,
		},
	)
	return qs
}

// ScopeNameEq filters for ScopeName being equal to argument
func (qs DhcpoptionscopeQS) ScopeNameEq(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" =`, v)
}

// ScopeNameNe filters for ScopeName being not equal to argument
func (qs DhcpoptionscopeQS) ScopeNameNe(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" <>`, v)
}

// ScopeNameLt filters for ScopeName being less than argument
func (qs DhcpoptionscopeQS) ScopeNameLt(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" <`, v)
}

// ScopeNameLe filters for ScopeName being less than or equal to argument
func (qs DhcpoptionscopeQS) ScopeNameLe(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" <=`, v)
}

// ScopeNameGt filters for ScopeName being greater than argument
func (qs DhcpoptionscopeQS) ScopeNameGt(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" >`, v)
}

// ScopeNameGe filters for ScopeName being greater than or equal to argument
func (qs DhcpoptionscopeQS) ScopeNameGe(v string) DhcpoptionscopeQS {
	return qs.filter(`"scope_name" >=`, v)
}

type inDhcpoptionscopeScopeName []interface{}

func (in inDhcpoptionscopeScopeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"scope_name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DhcpoptionscopeQS) ScopeNameIn(values []string) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcpoptionscopeScopeName(vals),
	)

	return qs
}

type notinDhcpoptionscopeScopeName []interface{}

func (in notinDhcpoptionscopeScopeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"scope_name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DhcpoptionscopeQS) ScopeNameNotIn(values []string) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcpoptionscopeScopeName(vals),
	)

	return qs
}

// OrderByScopeName sorts result by ScopeName in ascending order
func (qs DhcpoptionscopeQS) OrderByScopeName() DhcpoptionscopeQS {
	qs.order = append(qs.order, `"scope_name"`)

	return qs
}

// OrderByScopeNameDesc sorts result by ScopeName in descending order
func (qs DhcpoptionscopeQS) OrderByScopeNameDesc() DhcpoptionscopeQS {
	qs.order = append(qs.order, `"scope_name" DESC`)

	return qs
}

// DistinctOnScopeName marks field in queries to add to DISTINCT ON clause
func (qs DhcpoptionscopeQS) DistinctOnScopeName() DhcpoptionscopeQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"scope_name"`)

	return qs
}

// END - django_kea.DhcpOptionScope.scope_name

// OrderByRandom randomizes result
func (qs DhcpoptionscopeQS) OrderByRandom() DhcpoptionscopeQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs DhcpoptionscopeQS) ForUpdate() DhcpoptionscopeQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs DhcpoptionscopeQS) ForUpdateNowait() DhcpoptionscopeQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs DhcpoptionscopeQS) ForUpdateSkipLocked() DhcpoptionscopeQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs DhcpoptionscopeQS) ClearForUpdate() DhcpoptionscopeQS {
	qs.forClause = ""

	return qs
}

func (qs DhcpoptionscopeQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs DhcpoptionscopeQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs DhcpoptionscopeQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"scope_id", "scope_name" FROM "dhcp_option_scope"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs DhcpoptionscopeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "scope_id" FROM "dhcp_option_scope"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs DhcpoptionscopeQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"scope_id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "dhcp_option_scope"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs DhcpoptionscopeQS) All(ctx context.Context, db models.DBInterface) (DhcpoptionscopeList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret DhcpoptionscopeList
	for rows.Next() {
		obj := Dhcpoptionscope{existsInDB: true}
		if err = rows.Scan(&obj.ScopeId, &obj.ScopeName); err != nil {
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
func (qs DhcpoptionscopeQS) First(ctx context.Context, db models.DBInterface) (*Dhcpoptionscope, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Dhcpoptionscope{existsInDB: true}
	err := row.Scan(&obj.ScopeId, &obj.ScopeName)
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
func (qs DhcpoptionscopeQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp_option_scope"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs DhcpoptionscopeQS) Update() DhcpoptionscopeUpdateQS {
	return DhcpoptionscopeUpdateQS{condFragments: qs.condFragments}
}

// DhcpoptionscopeUpdateQS represents an updated queryset for django_kea.DhcpOptionScope
type DhcpoptionscopeUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs DhcpoptionscopeUpdateQS) update(c string, v interface{}) DhcpoptionscopeUpdateQS {
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

// SetScopeId sets ScopeId to the given value
func (uqs DhcpoptionscopeUpdateQS) SetScopeId(v int32) DhcpoptionscopeUpdateQS {
	return uqs.update(`"scope_id"`, v)
}

// SetScopeName sets ScopeName to the given value
func (uqs DhcpoptionscopeUpdateQS) SetScopeName(v sql.NullString) DhcpoptionscopeUpdateQS {
	return uqs.update(`"scope_name"`, v)
}

// Exec executes the update operation
func (uqs DhcpoptionscopeUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := DhcpoptionscopeQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "dhcp_option_scope" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (d *Dhcpoptionscope) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "dhcp_option_scope" ("scope_name", "scope_id") VALUES ($1, $2)`, d.ScopeName, d.ScopeId)

	if err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcpoptionscope) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "dhcp_option_scope" SET "scope_name" = $1 WHERE "scope_id" = $2`, d.ScopeName, d.ScopeId)

	return err
}

// Save inserts or updates record
func (d *Dhcpoptionscope) Save(ctx context.Context, db models.DBInterface) error {
	if d.existsInDB {
		return d.update(ctx, db)
	}

	return d.insert(ctx, db)
}

// Delete removes row from database
func (d *Dhcpoptionscope) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "dhcp_option_scope" WHERE "scope_id" = $1`, d.ScopeId)

	d.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (dl DhcpoptionscopeList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts DhcpoptionscopeList

	for _, d := range dl {
		if d.existsInDB {
			if err := d.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, d)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, d := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, d.ScopeName, d.ScopeId)
		offs += 2
	}

	qs := `INSERT INTO "dhcp_option_scope" ("scope_name", "scope_id") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, d := range inserts {
		d.existsInDB = true
	}

	return nil

}

// Dhcp4options returns the set of Dhcp4options referencing this Dhcpoptionscope instance
func (d *Dhcpoptionscope) Dhcp4options() Dhcp4optionsQS {
	return Dhcp4optionsQS{}.ScopeEq(d)
}

// Dhcp6options returns the set of Dhcp6options referencing this Dhcpoptionscope instance
func (d *Dhcpoptionscope) Dhcp6options() Dhcp6optionsQS {
	return Dhcp6optionsQS{}.ScopeEq(d)
}
