// Code generated for Django model django_kea.Dhcp4Options. DO NOT EDIT.

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
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/euronetzrt/django-kea/models"
)

// Dhcp4options mirrors model django_kea.Dhcp4Options
type Dhcp4options struct {
	existsInDB bool

	optionId        int32
	Code            int32
	Value           sql.NullString
	FormattedValue  sql.NullString
	Space           sql.NullString
	Persistent      bool
	DhcpClientClass sql.NullString
	Dhcp4SubnetId   sql.NullInt64
	host            sql.NullInt32
	scope           int32
	UserContext     sql.NullString
	ModificationTs  time.Time
	Cancelled       bool
}

// Dhcp4optionsList is a list of Dhcp4options
type Dhcp4optionsList []*Dhcp4options

// Dhcp4optionsQS represents a queryset for django_kea.Dhcp4Options
type Dhcp4optionsQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs Dhcp4optionsQS) filter(c string, p interface{}) Dhcp4optionsQS {
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
func (qs Dhcp4optionsQS) Or(exprs ...Dhcp4optionsQS) Dhcp4optionsQS {
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

// BEGIN - django_kea.Dhcp4Options.option_id

// GetOptionId returns Dhcp4options.OptionId
func (d *Dhcp4options) GetOptionId() int32 {
	return d.optionId
}

// OptionIdEq filters for optionId being equal to argument
func (qs Dhcp4optionsQS) OptionIdEq(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" =`, v)
}

// OptionIdNe filters for optionId being not equal to argument
func (qs Dhcp4optionsQS) OptionIdNe(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" <>`, v)
}

// OptionIdLt filters for optionId being less than argument
func (qs Dhcp4optionsQS) OptionIdLt(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" <`, v)
}

// OptionIdLe filters for optionId being less than or equal to argument
func (qs Dhcp4optionsQS) OptionIdLe(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" <=`, v)
}

// OptionIdGt filters for optionId being greater than argument
func (qs Dhcp4optionsQS) OptionIdGt(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" >`, v)
}

// OptionIdGe filters for optionId being greater than or equal to argument
func (qs Dhcp4optionsQS) OptionIdGe(v int32) Dhcp4optionsQS {
	return qs.filter(`"option_id" >=`, v)
}

type inDhcp4optionsoptionId []interface{}

func (in inDhcp4optionsoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"option_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) OptionIdIn(values []int32) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsoptionId(vals),
	)

	return qs
}

type notinDhcp4optionsoptionId []interface{}

func (in notinDhcp4optionsoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"option_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) OptionIdNotIn(values []int32) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsoptionId(vals),
	)

	return qs
}

// OrderByOptionId sorts result by OptionId in ascending order
func (qs Dhcp4optionsQS) OrderByOptionId() Dhcp4optionsQS {
	qs.order = append(qs.order, `"option_id"`)

	return qs
}

// OrderByOptionIdDesc sorts result by OptionId in descending order
func (qs Dhcp4optionsQS) OrderByOptionIdDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"option_id" DESC`)

	return qs
}

// DistinctOnOptionId marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnOptionId() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"option_id"`)

	return qs
}

// END - django_kea.Dhcp4Options.option_id

// BEGIN - django_kea.Dhcp4Options.code

// CodeEq filters for Code being equal to argument
func (qs Dhcp4optionsQS) CodeEq(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" =`, v)
}

// CodeNe filters for Code being not equal to argument
func (qs Dhcp4optionsQS) CodeNe(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" <>`, v)
}

// CodeLt filters for Code being less than argument
func (qs Dhcp4optionsQS) CodeLt(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" <`, v)
}

// CodeLe filters for Code being less than or equal to argument
func (qs Dhcp4optionsQS) CodeLe(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" <=`, v)
}

// CodeGt filters for Code being greater than argument
func (qs Dhcp4optionsQS) CodeGt(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" >`, v)
}

// CodeGe filters for Code being greater than or equal to argument
func (qs Dhcp4optionsQS) CodeGe(v int32) Dhcp4optionsQS {
	return qs.filter(`"code" >=`, v)
}

type inDhcp4optionsCode []interface{}

func (in inDhcp4optionsCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"code" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) CodeIn(values []int32) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsCode(vals),
	)

	return qs
}

type notinDhcp4optionsCode []interface{}

func (in notinDhcp4optionsCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"code" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) CodeNotIn(values []int32) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsCode(vals),
	)

	return qs
}

// OrderByCode sorts result by Code in ascending order
func (qs Dhcp4optionsQS) OrderByCode() Dhcp4optionsQS {
	qs.order = append(qs.order, `"code"`)

	return qs
}

// OrderByCodeDesc sorts result by Code in descending order
func (qs Dhcp4optionsQS) OrderByCodeDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"code" DESC`)

	return qs
}

// DistinctOnCode marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnCode() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"code"`)

	return qs
}

// END - django_kea.Dhcp4Options.code

// BEGIN - django_kea.Dhcp4Options.value

// ValueIsNull filters for Value being null
func (qs Dhcp4optionsQS) ValueIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NULL`,
		},
	)
	return qs
}

// ValueIsNotNull filters for Value being not null
func (qs Dhcp4optionsQS) ValueIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NOT NULL`,
		},
	)
	return qs
}

// ValueEq filters for Value being equal to argument
func (qs Dhcp4optionsQS) ValueEq(v string) Dhcp4optionsQS {
	return qs.filter(`"value" =`, v)
}

// ValueNe filters for Value being not equal to argument
func (qs Dhcp4optionsQS) ValueNe(v string) Dhcp4optionsQS {
	return qs.filter(`"value" <>`, v)
}

// ValueLt filters for Value being less than argument
func (qs Dhcp4optionsQS) ValueLt(v string) Dhcp4optionsQS {
	return qs.filter(`"value" <`, v)
}

// ValueLe filters for Value being less than or equal to argument
func (qs Dhcp4optionsQS) ValueLe(v string) Dhcp4optionsQS {
	return qs.filter(`"value" <=`, v)
}

// ValueGt filters for Value being greater than argument
func (qs Dhcp4optionsQS) ValueGt(v string) Dhcp4optionsQS {
	return qs.filter(`"value" >`, v)
}

// ValueGe filters for Value being greater than or equal to argument
func (qs Dhcp4optionsQS) ValueGe(v string) Dhcp4optionsQS {
	return qs.filter(`"value" >=`, v)
}

type inDhcp4optionsValue []interface{}

func (in inDhcp4optionsValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"value" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) ValueIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsValue(vals),
	)

	return qs
}

type notinDhcp4optionsValue []interface{}

func (in notinDhcp4optionsValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"value" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) ValueNotIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsValue(vals),
	)

	return qs
}

// OrderByValue sorts result by Value in ascending order
func (qs Dhcp4optionsQS) OrderByValue() Dhcp4optionsQS {
	qs.order = append(qs.order, `"value"`)

	return qs
}

// OrderByValueDesc sorts result by Value in descending order
func (qs Dhcp4optionsQS) OrderByValueDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"value" DESC`)

	return qs
}

// DistinctOnValue marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnValue() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"value"`)

	return qs
}

// END - django_kea.Dhcp4Options.value

// BEGIN - django_kea.Dhcp4Options.formatted_value

// FormattedValueIsNull filters for FormattedValue being null
func (qs Dhcp4optionsQS) FormattedValueIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NULL`,
		},
	)
	return qs
}

// FormattedValueIsNotNull filters for FormattedValue being not null
func (qs Dhcp4optionsQS) FormattedValueIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NOT NULL`,
		},
	)
	return qs
}

// FormattedValueEq filters for FormattedValue being equal to argument
func (qs Dhcp4optionsQS) FormattedValueEq(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" =`, v)
}

// FormattedValueNe filters for FormattedValue being not equal to argument
func (qs Dhcp4optionsQS) FormattedValueNe(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" <>`, v)
}

// FormattedValueLt filters for FormattedValue being less than argument
func (qs Dhcp4optionsQS) FormattedValueLt(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" <`, v)
}

// FormattedValueLe filters for FormattedValue being less than or equal to argument
func (qs Dhcp4optionsQS) FormattedValueLe(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" <=`, v)
}

// FormattedValueGt filters for FormattedValue being greater than argument
func (qs Dhcp4optionsQS) FormattedValueGt(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" >`, v)
}

// FormattedValueGe filters for FormattedValue being greater than or equal to argument
func (qs Dhcp4optionsQS) FormattedValueGe(v string) Dhcp4optionsQS {
	return qs.filter(`"formatted_value" >=`, v)
}

type inDhcp4optionsFormattedValue []interface{}

func (in inDhcp4optionsFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"formatted_value" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) FormattedValueIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsFormattedValue(vals),
	)

	return qs
}

type notinDhcp4optionsFormattedValue []interface{}

func (in notinDhcp4optionsFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"formatted_value" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) FormattedValueNotIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsFormattedValue(vals),
	)

	return qs
}

// OrderByFormattedValue sorts result by FormattedValue in ascending order
func (qs Dhcp4optionsQS) OrderByFormattedValue() Dhcp4optionsQS {
	qs.order = append(qs.order, `"formatted_value"`)

	return qs
}

// OrderByFormattedValueDesc sorts result by FormattedValue in descending order
func (qs Dhcp4optionsQS) OrderByFormattedValueDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"formatted_value" DESC`)

	return qs
}

// DistinctOnFormattedValue marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnFormattedValue() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"formatted_value"`)

	return qs
}

// END - django_kea.Dhcp4Options.formatted_value

// BEGIN - django_kea.Dhcp4Options.space

// SpaceIsNull filters for Space being null
func (qs Dhcp4optionsQS) SpaceIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NULL`,
		},
	)
	return qs
}

// SpaceIsNotNull filters for Space being not null
func (qs Dhcp4optionsQS) SpaceIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NOT NULL`,
		},
	)
	return qs
}

// SpaceEq filters for Space being equal to argument
func (qs Dhcp4optionsQS) SpaceEq(v string) Dhcp4optionsQS {
	return qs.filter(`"space" =`, v)
}

// SpaceNe filters for Space being not equal to argument
func (qs Dhcp4optionsQS) SpaceNe(v string) Dhcp4optionsQS {
	return qs.filter(`"space" <>`, v)
}

// SpaceLt filters for Space being less than argument
func (qs Dhcp4optionsQS) SpaceLt(v string) Dhcp4optionsQS {
	return qs.filter(`"space" <`, v)
}

// SpaceLe filters for Space being less than or equal to argument
func (qs Dhcp4optionsQS) SpaceLe(v string) Dhcp4optionsQS {
	return qs.filter(`"space" <=`, v)
}

// SpaceGt filters for Space being greater than argument
func (qs Dhcp4optionsQS) SpaceGt(v string) Dhcp4optionsQS {
	return qs.filter(`"space" >`, v)
}

// SpaceGe filters for Space being greater than or equal to argument
func (qs Dhcp4optionsQS) SpaceGe(v string) Dhcp4optionsQS {
	return qs.filter(`"space" >=`, v)
}

type inDhcp4optionsSpace []interface{}

func (in inDhcp4optionsSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"space" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) SpaceIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsSpace(vals),
	)

	return qs
}

type notinDhcp4optionsSpace []interface{}

func (in notinDhcp4optionsSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"space" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) SpaceNotIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsSpace(vals),
	)

	return qs
}

// OrderBySpace sorts result by Space in ascending order
func (qs Dhcp4optionsQS) OrderBySpace() Dhcp4optionsQS {
	qs.order = append(qs.order, `"space"`)

	return qs
}

// OrderBySpaceDesc sorts result by Space in descending order
func (qs Dhcp4optionsQS) OrderBySpaceDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"space" DESC`)

	return qs
}

// DistinctOnSpace marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnSpace() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"space"`)

	return qs
}

// END - django_kea.Dhcp4Options.space

// BEGIN - django_kea.Dhcp4Options.persistent

// PersistentEq filters for Persistent being equal to argument
func (qs Dhcp4optionsQS) PersistentEq(v bool) Dhcp4optionsQS {
	return qs.filter(`"persistent" =`, v)
}

// PersistentNe filters for Persistent being not equal to argument
func (qs Dhcp4optionsQS) PersistentNe(v bool) Dhcp4optionsQS {
	return qs.filter(`"persistent" <>`, v)
}

type inDhcp4optionsPersistent []interface{}

func (in inDhcp4optionsPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"persistent" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) PersistentIn(values []bool) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsPersistent(vals),
	)

	return qs
}

type notinDhcp4optionsPersistent []interface{}

func (in notinDhcp4optionsPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"persistent" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) PersistentNotIn(values []bool) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsPersistent(vals),
	)

	return qs
}

// OrderByPersistent sorts result by Persistent in ascending order
func (qs Dhcp4optionsQS) OrderByPersistent() Dhcp4optionsQS {
	qs.order = append(qs.order, `"persistent"`)

	return qs
}

// OrderByPersistentDesc sorts result by Persistent in descending order
func (qs Dhcp4optionsQS) OrderByPersistentDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"persistent" DESC`)

	return qs
}

// DistinctOnPersistent marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnPersistent() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"persistent"`)

	return qs
}

// END - django_kea.Dhcp4Options.persistent

// BEGIN - django_kea.Dhcp4Options.dhcp_client_class

// DhcpClientClassIsNull filters for DhcpClientClass being null
func (qs Dhcp4optionsQS) DhcpClientClassIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NULL`,
		},
	)
	return qs
}

// DhcpClientClassIsNotNull filters for DhcpClientClass being not null
func (qs Dhcp4optionsQS) DhcpClientClassIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NOT NULL`,
		},
	)
	return qs
}

// DhcpClientClassEq filters for DhcpClientClass being equal to argument
func (qs Dhcp4optionsQS) DhcpClientClassEq(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" =`, v)
}

// DhcpClientClassNe filters for DhcpClientClass being not equal to argument
func (qs Dhcp4optionsQS) DhcpClientClassNe(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" <>`, v)
}

// DhcpClientClassLt filters for DhcpClientClass being less than argument
func (qs Dhcp4optionsQS) DhcpClientClassLt(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" <`, v)
}

// DhcpClientClassLe filters for DhcpClientClass being less than or equal to argument
func (qs Dhcp4optionsQS) DhcpClientClassLe(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" <=`, v)
}

// DhcpClientClassGt filters for DhcpClientClass being greater than argument
func (qs Dhcp4optionsQS) DhcpClientClassGt(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" >`, v)
}

// DhcpClientClassGe filters for DhcpClientClass being greater than or equal to argument
func (qs Dhcp4optionsQS) DhcpClientClassGe(v string) Dhcp4optionsQS {
	return qs.filter(`"dhcp_client_class" >=`, v)
}

type inDhcp4optionsDhcpClientClass []interface{}

func (in inDhcp4optionsDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) DhcpClientClassIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsDhcpClientClass(vals),
	)

	return qs
}

type notinDhcp4optionsDhcpClientClass []interface{}

func (in notinDhcp4optionsDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) DhcpClientClassNotIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsDhcpClientClass(vals),
	)

	return qs
}

// OrderByDhcpClientClass sorts result by DhcpClientClass in ascending order
func (qs Dhcp4optionsQS) OrderByDhcpClientClass() Dhcp4optionsQS {
	qs.order = append(qs.order, `"dhcp_client_class"`)

	return qs
}

// OrderByDhcpClientClassDesc sorts result by DhcpClientClass in descending order
func (qs Dhcp4optionsQS) OrderByDhcpClientClassDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"dhcp_client_class" DESC`)

	return qs
}

// DistinctOnDhcpClientClass marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnDhcpClientClass() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp_client_class"`)

	return qs
}

// END - django_kea.Dhcp4Options.dhcp_client_class

// BEGIN - django_kea.Dhcp4Options.dhcp4_subnet_id

// Dhcp4SubnetIdIsNull filters for Dhcp4SubnetId being null
func (qs Dhcp4optionsQS) Dhcp4SubnetIdIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdIsNotNull filters for Dhcp4SubnetId being not null
func (qs Dhcp4optionsQS) Dhcp4SubnetIdIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdEq filters for Dhcp4SubnetId being equal to argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdEq(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" =`, v)
}

// Dhcp4SubnetIdNe filters for Dhcp4SubnetId being not equal to argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdNe(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" <>`, v)
}

// Dhcp4SubnetIdLt filters for Dhcp4SubnetId being less than argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdLt(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" <`, v)
}

// Dhcp4SubnetIdLe filters for Dhcp4SubnetId being less than or equal to argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdLe(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" <=`, v)
}

// Dhcp4SubnetIdGt filters for Dhcp4SubnetId being greater than argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdGt(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" >`, v)
}

// Dhcp4SubnetIdGe filters for Dhcp4SubnetId being greater than or equal to argument
func (qs Dhcp4optionsQS) Dhcp4SubnetIdGe(v int64) Dhcp4optionsQS {
	return qs.filter(`"dhcp4_subnet_id" >=`, v)
}

type inDhcp4optionsDhcp4SubnetId []interface{}

func (in inDhcp4optionsDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) Dhcp4SubnetIdIn(values []int64) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsDhcp4SubnetId(vals),
	)

	return qs
}

type notinDhcp4optionsDhcp4SubnetId []interface{}

func (in notinDhcp4optionsDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) Dhcp4SubnetIdNotIn(values []int64) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsDhcp4SubnetId(vals),
	)

	return qs
}

// OrderByDhcp4SubnetId sorts result by Dhcp4SubnetId in ascending order
func (qs Dhcp4optionsQS) OrderByDhcp4SubnetId() Dhcp4optionsQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id"`)

	return qs
}

// OrderByDhcp4SubnetIdDesc sorts result by Dhcp4SubnetId in descending order
func (qs Dhcp4optionsQS) OrderByDhcp4SubnetIdDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id" DESC`)

	return qs
}

// DistinctOnDhcp4SubnetId marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnDhcp4SubnetId() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_subnet_id"`)

	return qs
}

// END - django_kea.Dhcp4Options.dhcp4_subnet_id

// BEGIN - django_kea.Dhcp4Options.host

// GetHost returns Hosts
func (d *Dhcp4options) GetHost(ctx context.Context, db models.DBInterface) (*Hosts, error) {
	if !d.host.Valid {
		return nil, nil
	}

	return HostsQS{}.HostIdEq(d.host.Int32).First(ctx, db)
}

// SetHost sets foreign key pointer to Hosts
func (d *Dhcp4options) SetHost(ptr *Hosts) error {
	if ptr != nil {
		d.host.Int32 = ptr.GetHostId()
		d.host.Valid = true
	} else {
		d.host.Valid = false
	}

	return nil
}

// GetHostRaw returns Dhcp4options.Host
func (d *Dhcp4options) GetHostRaw() sql.NullInt32 {
	return d.host
}

// HostIsNull filters for host being null
func (qs Dhcp4optionsQS) HostIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NULL`,
		},
	)
	return qs
}

// HostIsNotNull filters for host being not null
func (qs Dhcp4optionsQS) HostIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NOT NULL`,
		},
	)
	return qs
}

// HostEq filters for host being equal to argument
func (qs Dhcp4optionsQS) HostEq(v *Hosts) Dhcp4optionsQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Dhcp4optionsQS) HostRawEq(v int32) Dhcp4optionsQS {
	return qs.filter(`"host_id" =`, v)
}

type inDhcp4optionshostHosts struct {
	qs HostsQS
}

func (in *inDhcp4optionshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Dhcp4optionsQS) HostIn(oqs HostsQS) Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionshostHosts{
			qs: oqs,
		},
	)

	return qs
}

type notinDhcp4optionshostHosts struct {
	qs HostsQS
}

func (nin *notinDhcp4optionshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"host_id" NOT IN (` + s + `)`, p
}

func (qs Dhcp4optionsQS) HostNotIn(oqs HostsQS) Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionshostHosts{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Dhcp4optionsQS) OrderByHost() Dhcp4optionsQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Dhcp4optionsQS) OrderByHostDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// DistinctOnHost marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnHost() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"host_id"`)

	return qs
}

// END - django_kea.Dhcp4Options.host

// BEGIN - django_kea.Dhcp4Options.scope

// GetScope returns Dhcpoptionscope
func (d *Dhcp4options) GetScope(ctx context.Context, db models.DBInterface) (*Dhcpoptionscope, error) {
	return DhcpoptionscopeQS{}.ScopeIdEq(d.scope).First(ctx, db)
}

// SetScope sets foreign key pointer to Dhcpoptionscope
func (d *Dhcp4options) SetScope(ptr *Dhcpoptionscope) error {
	if ptr != nil {
		d.scope = ptr.ScopeId
	} else {
		return fmt.Errorf("Dhcp4options.SetScope: non-null field received null value")
	}

	return nil
}

// GetScopeRaw returns Dhcp4options.Scope
func (d *Dhcp4options) GetScopeRaw() int32 {
	return d.scope
}

// ScopeEq filters for scope being equal to argument
func (qs Dhcp4optionsQS) ScopeEq(v *Dhcpoptionscope) Dhcp4optionsQS {
	return qs.filter(`"scope_id" =`, v.ScopeId)
}

// ScopeRawEq filters for scope being equal to raw argument
func (qs Dhcp4optionsQS) ScopeRawEq(v int32) Dhcp4optionsQS {
	return qs.filter(`"scope_id" =`, v)
}

type inDhcp4optionsscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (in *inDhcp4optionsscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"scope_id" IN (` + s + `)`, p
}

func (qs Dhcp4optionsQS) ScopeIn(oqs DhcpoptionscopeQS) Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionsscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

type notinDhcp4optionsscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (nin *notinDhcp4optionsscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"scope_id" NOT IN (` + s + `)`, p
}

func (qs Dhcp4optionsQS) ScopeNotIn(oqs DhcpoptionscopeQS) Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionsscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

// OrderByScope sorts result by Scope in ascending order
func (qs Dhcp4optionsQS) OrderByScope() Dhcp4optionsQS {
	qs.order = append(qs.order, `"scope_id"`)

	return qs
}

// OrderByScopeDesc sorts result by Scope in descending order
func (qs Dhcp4optionsQS) OrderByScopeDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"scope_id" DESC`)

	return qs
}

// DistinctOnScope marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnScope() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"scope_id"`)

	return qs
}

// END - django_kea.Dhcp4Options.scope

// BEGIN - django_kea.Dhcp4Options.user_context

// UserContextIsNull filters for UserContext being null
func (qs Dhcp4optionsQS) UserContextIsNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Dhcp4optionsQS) UserContextIsNotNull() Dhcp4optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Dhcp4optionsQS) UserContextEq(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Dhcp4optionsQS) UserContextNe(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Dhcp4optionsQS) UserContextLt(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Dhcp4optionsQS) UserContextLe(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Dhcp4optionsQS) UserContextGt(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Dhcp4optionsQS) UserContextGe(v string) Dhcp4optionsQS {
	return qs.filter(`"user_context" >=`, v)
}

type inDhcp4optionsUserContext []interface{}

func (in inDhcp4optionsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) UserContextIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsUserContext(vals),
	)

	return qs
}

type notinDhcp4optionsUserContext []interface{}

func (in notinDhcp4optionsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) UserContextNotIn(values []string) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsUserContext(vals),
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Dhcp4optionsQS) OrderByUserContext() Dhcp4optionsQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Dhcp4optionsQS) OrderByUserContextDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// DistinctOnUserContext marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnUserContext() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"user_context"`)

	return qs
}

// END - django_kea.Dhcp4Options.user_context

// BEGIN - django_kea.Dhcp4Options.modification_ts

// ModificationTsEq filters for ModificationTs being equal to argument
func (qs Dhcp4optionsQS) ModificationTsEq(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" =`, v)
}

// ModificationTsNe filters for ModificationTs being not equal to argument
func (qs Dhcp4optionsQS) ModificationTsNe(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" <>`, v)
}

// ModificationTsLt filters for ModificationTs being less than argument
func (qs Dhcp4optionsQS) ModificationTsLt(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" <`, v)
}

// ModificationTsLe filters for ModificationTs being less than or equal to argument
func (qs Dhcp4optionsQS) ModificationTsLe(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" <=`, v)
}

// ModificationTsGt filters for ModificationTs being greater than argument
func (qs Dhcp4optionsQS) ModificationTsGt(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" >`, v)
}

// ModificationTsGe filters for ModificationTs being greater than or equal to argument
func (qs Dhcp4optionsQS) ModificationTsGe(v time.Time) Dhcp4optionsQS {
	return qs.filter(`"modification_ts" >=`, v)
}

type inDhcp4optionsModificationTs []interface{}

func (in inDhcp4optionsModificationTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"modification_ts" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) ModificationTsIn(values []time.Time) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsModificationTs(vals),
	)

	return qs
}

type notinDhcp4optionsModificationTs []interface{}

func (in notinDhcp4optionsModificationTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"modification_ts" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) ModificationTsNotIn(values []time.Time) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsModificationTs(vals),
	)

	return qs
}

// OrderByModificationTs sorts result by ModificationTs in ascending order
func (qs Dhcp4optionsQS) OrderByModificationTs() Dhcp4optionsQS {
	qs.order = append(qs.order, `"modification_ts"`)

	return qs
}

// OrderByModificationTsDesc sorts result by ModificationTs in descending order
func (qs Dhcp4optionsQS) OrderByModificationTsDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"modification_ts" DESC`)

	return qs
}

// DistinctOnModificationTs marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnModificationTs() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"modification_ts"`)

	return qs
}

// END - django_kea.Dhcp4Options.modification_ts

// BEGIN - django_kea.Dhcp4Options.cancelled

// CancelledEq filters for Cancelled being equal to argument
func (qs Dhcp4optionsQS) CancelledEq(v bool) Dhcp4optionsQS {
	return qs.filter(`"cancelled" =`, v)
}

// CancelledNe filters for Cancelled being not equal to argument
func (qs Dhcp4optionsQS) CancelledNe(v bool) Dhcp4optionsQS {
	return qs.filter(`"cancelled" <>`, v)
}

type inDhcp4optionsCancelled []interface{}

func (in inDhcp4optionsCancelled) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"cancelled" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) CancelledIn(values []bool) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp4optionsCancelled(vals),
	)

	return qs
}

type notinDhcp4optionsCancelled []interface{}

func (in notinDhcp4optionsCancelled) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"cancelled" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp4optionsQS) CancelledNotIn(values []bool) Dhcp4optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp4optionsCancelled(vals),
	)

	return qs
}

// OrderByCancelled sorts result by Cancelled in ascending order
func (qs Dhcp4optionsQS) OrderByCancelled() Dhcp4optionsQS {
	qs.order = append(qs.order, `"cancelled"`)

	return qs
}

// OrderByCancelledDesc sorts result by Cancelled in descending order
func (qs Dhcp4optionsQS) OrderByCancelledDesc() Dhcp4optionsQS {
	qs.order = append(qs.order, `"cancelled" DESC`)

	return qs
}

// DistinctOnCancelled marks field in queries to add to DISTINCT ON clause
func (qs Dhcp4optionsQS) DistinctOnCancelled() Dhcp4optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"cancelled"`)

	return qs
}

// END - django_kea.Dhcp4Options.cancelled

// OrderByRandom randomizes result
func (qs Dhcp4optionsQS) OrderByRandom() Dhcp4optionsQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Dhcp4optionsQS) ForUpdate() Dhcp4optionsQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs Dhcp4optionsQS) ForUpdateNowait() Dhcp4optionsQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs Dhcp4optionsQS) ForUpdateSkipLocked() Dhcp4optionsQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs Dhcp4optionsQS) ClearForUpdate() Dhcp4optionsQS {
	qs.forClause = ""

	return qs
}

func (qs Dhcp4optionsQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Dhcp4optionsQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Dhcp4optionsQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"option_id", "code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp4_subnet_id", "host_id", "scope_id", "user_context", "modification_ts", "cancelled" FROM "dhcp4_options"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Dhcp4optionsQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "option_id" FROM "dhcp4_options"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs Dhcp4optionsQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"option_id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "dhcp4_options"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs Dhcp4optionsQS) All(ctx context.Context, db models.DBInterface) (Dhcp4optionsList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret Dhcp4optionsList
	for rows.Next() {
		obj := Dhcp4options{existsInDB: true}
		if err = rows.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp4SubnetId, &obj.host, &obj.scope, &obj.UserContext, &obj.ModificationTs, &obj.Cancelled); err != nil {
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
func (qs Dhcp4optionsQS) First(ctx context.Context, db models.DBInterface) (*Dhcp4options, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Dhcp4options{existsInDB: true}
	err := row.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp4SubnetId, &obj.host, &obj.scope, &obj.UserContext, &obj.ModificationTs, &obj.Cancelled)
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
func (qs Dhcp4optionsQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp4_options"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Dhcp4optionsQS) Update() Dhcp4optionsUpdateQS {
	return Dhcp4optionsUpdateQS{condFragments: qs.condFragments}
}

// Dhcp4optionsUpdateQS represents an updated queryset for django_kea.Dhcp4Options
type Dhcp4optionsUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Dhcp4optionsUpdateQS) update(c string, v interface{}) Dhcp4optionsUpdateQS {
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

// SetOptionId sets OptionId to the given value
func (uqs Dhcp4optionsUpdateQS) SetOptionId(v int32) Dhcp4optionsUpdateQS {
	return uqs.update(`"option_id"`, v)
}

// SetCode sets Code to the given value
func (uqs Dhcp4optionsUpdateQS) SetCode(v int32) Dhcp4optionsUpdateQS {
	return uqs.update(`"code"`, v)
}

// SetValue sets Value to the given value
func (uqs Dhcp4optionsUpdateQS) SetValue(v sql.NullString) Dhcp4optionsUpdateQS {
	return uqs.update(`"value"`, v)
}

// SetFormattedValue sets FormattedValue to the given value
func (uqs Dhcp4optionsUpdateQS) SetFormattedValue(v sql.NullString) Dhcp4optionsUpdateQS {
	return uqs.update(`"formatted_value"`, v)
}

// SetSpace sets Space to the given value
func (uqs Dhcp4optionsUpdateQS) SetSpace(v sql.NullString) Dhcp4optionsUpdateQS {
	return uqs.update(`"space"`, v)
}

// SetPersistent sets Persistent to the given value
func (uqs Dhcp4optionsUpdateQS) SetPersistent(v bool) Dhcp4optionsUpdateQS {
	return uqs.update(`"persistent"`, v)
}

// SetDhcpClientClass sets DhcpClientClass to the given value
func (uqs Dhcp4optionsUpdateQS) SetDhcpClientClass(v sql.NullString) Dhcp4optionsUpdateQS {
	return uqs.update(`"dhcp_client_class"`, v)
}

// SetDhcp4SubnetId sets Dhcp4SubnetId to the given value
func (uqs Dhcp4optionsUpdateQS) SetDhcp4SubnetId(v sql.NullInt64) Dhcp4optionsUpdateQS {
	return uqs.update(`"dhcp4_subnet_id"`, v)
}

// SetHost sets foreign key pointer to Hosts
func (uqs Dhcp4optionsUpdateQS) SetHost(ptr *Hosts) Dhcp4optionsUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // SetScope sets foreign key pointer to Dhcpoptionscope
func (uqs Dhcp4optionsUpdateQS) SetScope(ptr *Dhcpoptionscope) Dhcp4optionsUpdateQS {
	if ptr != nil {
		return uqs.update(`"scope_id"`, ptr.ScopeId)
	}

	return uqs.update(`"scope_id"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Dhcp4optionsUpdateQS) SetUserContext(v sql.NullString) Dhcp4optionsUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// SetModificationTs sets ModificationTs to the given value
func (uqs Dhcp4optionsUpdateQS) SetModificationTs(v time.Time) Dhcp4optionsUpdateQS {
	return uqs.update(`"modification_ts"`, v)
}

// SetCancelled sets Cancelled to the given value
func (uqs Dhcp4optionsUpdateQS) SetCancelled(v bool) Dhcp4optionsUpdateQS {
	return uqs.update(`"cancelled"`, v)
}

// Exec executes the update operation
func (uqs Dhcp4optionsUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := Dhcp4optionsQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "dhcp4_options" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (d *Dhcp4options) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "dhcp4_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp4_subnet_id", "host_id", "scope_id", "user_context", "modification_ts", "cancelled") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING "option_id"`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp4SubnetId, d.host, d.scope, d.UserContext, d.ModificationTs, d.Cancelled)

	if err := row.Scan(&d.optionId); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcp4options) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "dhcp4_options" SET "code" = $1, "value" = $2, "formatted_value" = $3, "space" = $4, "persistent" = $5, "dhcp_client_class" = $6, "dhcp4_subnet_id" = $7, "host_id" = $8, "scope_id" = $9, "user_context" = $10, "modification_ts" = $11, "cancelled" = $12 WHERE "option_id" = $13`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp4SubnetId, d.host, d.scope, d.UserContext, d.ModificationTs, d.Cancelled, d.optionId)

	return err
}

// Save inserts or updates record
func (d *Dhcp4options) Save(ctx context.Context, db models.DBInterface) error {
	if d.existsInDB {
		return d.update(ctx, db)
	}

	return d.insert(ctx, db)
}

// Delete removes row from database
func (d *Dhcp4options) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "dhcp4_options" WHERE "option_id" = $1`, d.optionId)

	d.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (dl Dhcp4optionsList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts Dhcp4optionsList

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
	vaa := make([]any, 0, 12*len(inserts))
	offs := 1
	for _, d := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4, offs+5, offs+6, offs+7, offs+8, offs+9, offs+10, offs+11))
		vaa = append(vaa, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp4SubnetId, d.host, d.scope, d.UserContext, d.ModificationTs, d.Cancelled)
		offs += 12
	}

	qs := `INSERT INTO "dhcp4_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp4_subnet_id", "host_id", "scope_id", "user_context", "modification_ts", "cancelled") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "option_id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, d := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&d.optionId); err != nil {
			return err
		}

		d.existsInDB = true
	}

	return rows.Err()
}
