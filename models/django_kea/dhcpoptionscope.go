/*
  AUTO-GENERATED file for Django model django_kea.DhcpOptionScope

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

// Dhcpoptionscope mirrors model django_kea.DhcpOptionScope
type Dhcpoptionscope struct {
	existsInDB bool

	ScopeId   int32
	ScopeName sql.NullString
}

// DhcpoptionscopeQS represents a queryset for django_kea.DhcpOptionScope
type DhcpoptionscopeQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
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

type inDhcpoptionscopeScopeId struct {
	values []interface{}
}

func (in *inDhcpoptionscopeScopeId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"scope_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DhcpoptionscopeQS) ScopeIdIn(values []int32) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcpoptionscopeScopeId{
			values: vals,
		},
	)

	return qs
}

type notinDhcpoptionscopeScopeId struct {
	values []interface{}
}

func (in *notinDhcpoptionscopeScopeId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"scope_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DhcpoptionscopeQS) ScopeIdNotIn(values []int32) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcpoptionscopeScopeId{
			values: vals,
		},
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

type inDhcpoptionscopeScopeName struct {
	values []interface{}
}

func (in *inDhcpoptionscopeScopeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"scope_name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DhcpoptionscopeQS) ScopeNameIn(values []string) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcpoptionscopeScopeName{
			values: vals,
		},
	)

	return qs
}

type notinDhcpoptionscopeScopeName struct {
	values []interface{}
}

func (in *notinDhcpoptionscopeScopeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"scope_name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DhcpoptionscopeQS) ScopeNameNotIn(values []string) DhcpoptionscopeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcpoptionscopeScopeName{
			values: vals,
		},
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

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs DhcpoptionscopeQS) ForUpdate() DhcpoptionscopeQS {
	qs.forUpdate = true

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

func (qs DhcpoptionscopeQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "scope_id", "scope_name" FROM "dhcp_option_scope"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs DhcpoptionscopeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "scope_id" FROM "dhcp_option_scope"` + s, p
}

// All returns all rows matching queryset filters
func (qs DhcpoptionscopeQS) All(db models.DBInterface) ([]*Dhcpoptionscope, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Dhcpoptionscope
	for rows.Next() {
		obj := Dhcpoptionscope{existsInDB: true}
		if err = rows.Scan(&obj.ScopeId, &obj.ScopeName); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs DhcpoptionscopeQS) First(db models.DBInterface) (*Dhcpoptionscope, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Dhcpoptionscope{existsInDB: true}
	err := row.Scan(&obj.ScopeId, &obj.ScopeName)
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
func (qs DhcpoptionscopeQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp_option_scope"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
func (uqs DhcpoptionscopeUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (d *Dhcpoptionscope) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "dhcp_option_scope" ("scope_name", "scope_id") VALUES ($1, $2)`, d.ScopeName, d.ScopeId)

	if err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcpoptionscope) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "dhcp_option_scope" SET "scope_name" = $1 WHERE "scope_id" = $2`, d.ScopeName, d.ScopeId)

	return err
}

// Save inserts or updates record
func (d *Dhcpoptionscope) Save(db models.DBInterface) error {
	if d.existsInDB {
		return d.update(db)
	}

	return d.insert(db)
}

// Delete removes row from database
func (d *Dhcpoptionscope) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "dhcp_option_scope" WHERE "scope_id" = $1`, d.ScopeId)

	d.existsInDB = false

	return err
}

// Dhcp4option returns the set of Dhcp4option referencing this Dhcpoptionscope instance
func (d *Dhcpoptionscope) Dhcp4option() Dhcp4optionQS {
	return Dhcp4optionQS{}.ScopeEq(d)
}

// Dhcp6option returns the set of Dhcp6option referencing this Dhcpoptionscope instance
func (d *Dhcpoptionscope) Dhcp6option() Dhcp6optionQS {
	return Dhcp6optionQS{}.ScopeEq(d)
}
