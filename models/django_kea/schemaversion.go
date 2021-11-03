/*
  AUTO-GENERATED file for Django model django_kea.SchemaVersion

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

// Schemaversion mirrors model django_kea.SchemaVersion
type Schemaversion struct {
	existsInDB bool

	Version int32
	Minor   sql.NullInt32
}

// SchemaversionQS represents a queryset for django_kea.SchemaVersion
type SchemaversionQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs SchemaversionQS) filter(c string, p interface{}) SchemaversionQS {
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
func (qs SchemaversionQS) Or(exprs ...SchemaversionQS) SchemaversionQS {
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

// VersionEq filters for Version being equal to argument
func (qs SchemaversionQS) VersionEq(v int32) SchemaversionQS {
	return qs.filter(`"version" =`, v)
}

// VersionNe filters for Version being not equal to argument
func (qs SchemaversionQS) VersionNe(v int32) SchemaversionQS {
	return qs.filter(`"version" <>`, v)
}

// VersionLt filters for Version being less than argument
func (qs SchemaversionQS) VersionLt(v int32) SchemaversionQS {
	return qs.filter(`"version" <`, v)
}

// VersionLe filters for Version being less than or equal to argument
func (qs SchemaversionQS) VersionLe(v int32) SchemaversionQS {
	return qs.filter(`"version" <=`, v)
}

// VersionGt filters for Version being greater than argument
func (qs SchemaversionQS) VersionGt(v int32) SchemaversionQS {
	return qs.filter(`"version" >`, v)
}

// VersionGe filters for Version being greater than or equal to argument
func (qs SchemaversionQS) VersionGe(v int32) SchemaversionQS {
	return qs.filter(`"version" >=`, v)
}

type inSchemaversionVersion struct {
	values []interface{}
}

func (in *inSchemaversionVersion) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"version" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SchemaversionQS) VersionIn(values []int32) SchemaversionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSchemaversionVersion{
			values: vals,
		},
	)

	return qs
}

type notinSchemaversionVersion struct {
	values []interface{}
}

func (in *notinSchemaversionVersion) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"version" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SchemaversionQS) VersionNotIn(values []int32) SchemaversionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSchemaversionVersion{
			values: vals,
		},
	)

	return qs
}

// OrderByVersion sorts result by Version in ascending order
func (qs SchemaversionQS) OrderByVersion() SchemaversionQS {
	qs.order = append(qs.order, `"version"`)

	return qs
}

// OrderByVersionDesc sorts result by Version in descending order
func (qs SchemaversionQS) OrderByVersionDesc() SchemaversionQS {
	qs.order = append(qs.order, `"version" DESC`)

	return qs
}

// MinorIsNull filters for Minor being null
func (qs SchemaversionQS) MinorIsNull() SchemaversionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"minor" IS NULL`,
		},
	)
	return qs
}

// MinorIsNotNull filters for Minor being not null
func (qs SchemaversionQS) MinorIsNotNull() SchemaversionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"minor" IS NOT NULL`,
		},
	)
	return qs
}

// MinorEq filters for Minor being equal to argument
func (qs SchemaversionQS) MinorEq(v int32) SchemaversionQS {
	return qs.filter(`"minor" =`, v)
}

// MinorNe filters for Minor being not equal to argument
func (qs SchemaversionQS) MinorNe(v int32) SchemaversionQS {
	return qs.filter(`"minor" <>`, v)
}

// MinorLt filters for Minor being less than argument
func (qs SchemaversionQS) MinorLt(v int32) SchemaversionQS {
	return qs.filter(`"minor" <`, v)
}

// MinorLe filters for Minor being less than or equal to argument
func (qs SchemaversionQS) MinorLe(v int32) SchemaversionQS {
	return qs.filter(`"minor" <=`, v)
}

// MinorGt filters for Minor being greater than argument
func (qs SchemaversionQS) MinorGt(v int32) SchemaversionQS {
	return qs.filter(`"minor" >`, v)
}

// MinorGe filters for Minor being greater than or equal to argument
func (qs SchemaversionQS) MinorGe(v int32) SchemaversionQS {
	return qs.filter(`"minor" >=`, v)
}

type inSchemaversionMinor struct {
	values []interface{}
}

func (in *inSchemaversionMinor) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"minor" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SchemaversionQS) MinorIn(values []int32) SchemaversionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSchemaversionMinor{
			values: vals,
		},
	)

	return qs
}

type notinSchemaversionMinor struct {
	values []interface{}
}

func (in *notinSchemaversionMinor) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"minor" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SchemaversionQS) MinorNotIn(values []int32) SchemaversionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSchemaversionMinor{
			values: vals,
		},
	)

	return qs
}

// OrderByMinor sorts result by Minor in ascending order
func (qs SchemaversionQS) OrderByMinor() SchemaversionQS {
	qs.order = append(qs.order, `"minor"`)

	return qs
}

// OrderByMinorDesc sorts result by Minor in descending order
func (qs SchemaversionQS) OrderByMinorDesc() SchemaversionQS {
	qs.order = append(qs.order, `"minor" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs SchemaversionQS) ForUpdate() SchemaversionQS {
	qs.forUpdate = true

	return qs
}

func (qs SchemaversionQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs SchemaversionQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs SchemaversionQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "version", "minor" FROM "schema_version"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs SchemaversionQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "version" FROM "schema_version"` + s, p
}

// All returns all rows matching queryset filters
func (qs SchemaversionQS) All(db models.DBInterface) ([]*Schemaversion, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Schemaversion
	for rows.Next() {
		obj := Schemaversion{existsInDB: true}
		if err = rows.Scan(&obj.Version, &obj.Minor); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs SchemaversionQS) First(db models.DBInterface) (*Schemaversion, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Schemaversion{existsInDB: true}
	err := row.Scan(&obj.Version, &obj.Minor)
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
func (qs SchemaversionQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "schema_version"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs SchemaversionQS) Update() SchemaversionUpdateQS {
	return SchemaversionUpdateQS{condFragments: qs.condFragments}
}

// SchemaversionUpdateQS represents an updated queryset for django_kea.SchemaVersion
type SchemaversionUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs SchemaversionUpdateQS) update(c string, v interface{}) SchemaversionUpdateQS {
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

// SetVersion sets Version to the given value
func (uqs SchemaversionUpdateQS) SetVersion(v int32) SchemaversionUpdateQS {
	return uqs.update(`"version"`, v)
}

// SetMinor sets Minor to the given value
func (uqs SchemaversionUpdateQS) SetMinor(v sql.NullInt32) SchemaversionUpdateQS {
	return uqs.update(`"minor"`, v)
}

// Exec executes the update operation
func (uqs SchemaversionUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := SchemaversionQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "schema_version" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (s *Schemaversion) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "schema_version" ("minor", "version") VALUES ($1, $2)`, s.Minor, s.Version)

	if err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Schemaversion) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "schema_version" SET "minor" = $1 WHERE "version" = $2`, s.Minor, s.Version)

	return err
}

// Save inserts or updates record
func (s *Schemaversion) Save(db models.DBInterface) error {
	if s.existsInDB {
		return s.update(db)
	}

	return s.insert(db)
}

// Delete removes row from database
func (s *Schemaversion) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "schema_version" WHERE "version" = $1`, s.Version)

	s.existsInDB = false

	return err
}
