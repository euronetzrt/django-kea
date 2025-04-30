// Code generated for Django model django_kea.Dhcp6Options. DO NOT EDIT.

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

// Dhcp6options mirrors model django_kea.Dhcp6Options
type Dhcp6options struct {
	existsInDB bool

	optionId        int32
	Code            int32
	Value           sql.NullString
	FormattedValue  sql.NullString
	Space           sql.NullString
	Persistent      bool
	DhcpClientClass sql.NullString
	Dhcp6SubnetId   sql.NullInt64
	host            sql.NullInt32
	scope           int32
	UserContext     sql.NullString
	Cancelled       bool
}

// Dhcp6optionsList is a list of Dhcp6options
type Dhcp6optionsList []*Dhcp6options

// Dhcp6optionsQS represents a queryset for django_kea.Dhcp6Options
type Dhcp6optionsQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs Dhcp6optionsQS) filter(c string, p interface{}) Dhcp6optionsQS {
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
func (qs Dhcp6optionsQS) Or(exprs ...Dhcp6optionsQS) Dhcp6optionsQS {
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

// BEGIN - django_kea.Dhcp6Options.option_id

// GetOptionId returns Dhcp6options.OptionId
func (d *Dhcp6options) GetOptionId() int32 {
	return d.optionId
}

// OptionIdEq filters for optionId being equal to argument
func (qs Dhcp6optionsQS) OptionIdEq(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" =`, v)
}

// OptionIdNe filters for optionId being not equal to argument
func (qs Dhcp6optionsQS) OptionIdNe(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" <>`, v)
}

// OptionIdLt filters for optionId being less than argument
func (qs Dhcp6optionsQS) OptionIdLt(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" <`, v)
}

// OptionIdLe filters for optionId being less than or equal to argument
func (qs Dhcp6optionsQS) OptionIdLe(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" <=`, v)
}

// OptionIdGt filters for optionId being greater than argument
func (qs Dhcp6optionsQS) OptionIdGt(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" >`, v)
}

// OptionIdGe filters for optionId being greater than or equal to argument
func (qs Dhcp6optionsQS) OptionIdGe(v int32) Dhcp6optionsQS {
	return qs.filter(`"option_id" >=`, v)
}

type inDhcp6optionsoptionId []interface{}

func (in inDhcp6optionsoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"option_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) OptionIdIn(values []int32) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsoptionId(vals),
	)

	return qs
}

type notinDhcp6optionsoptionId []interface{}

func (in notinDhcp6optionsoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"option_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) OptionIdNotIn(values []int32) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsoptionId(vals),
	)

	return qs
}

// OrderByOptionId sorts result by OptionId in ascending order
func (qs Dhcp6optionsQS) OrderByOptionId() Dhcp6optionsQS {
	qs.order = append(qs.order, `"option_id"`)

	return qs
}

// OrderByOptionIdDesc sorts result by OptionId in descending order
func (qs Dhcp6optionsQS) OrderByOptionIdDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"option_id" DESC`)

	return qs
}

// DistinctOnOptionId marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnOptionId() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"option_id"`)

	return qs
}

// END - django_kea.Dhcp6Options.option_id

// BEGIN - django_kea.Dhcp6Options.code

// CodeEq filters for Code being equal to argument
func (qs Dhcp6optionsQS) CodeEq(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" =`, v)
}

// CodeNe filters for Code being not equal to argument
func (qs Dhcp6optionsQS) CodeNe(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" <>`, v)
}

// CodeLt filters for Code being less than argument
func (qs Dhcp6optionsQS) CodeLt(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" <`, v)
}

// CodeLe filters for Code being less than or equal to argument
func (qs Dhcp6optionsQS) CodeLe(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" <=`, v)
}

// CodeGt filters for Code being greater than argument
func (qs Dhcp6optionsQS) CodeGt(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" >`, v)
}

// CodeGe filters for Code being greater than or equal to argument
func (qs Dhcp6optionsQS) CodeGe(v int32) Dhcp6optionsQS {
	return qs.filter(`"code" >=`, v)
}

type inDhcp6optionsCode []interface{}

func (in inDhcp6optionsCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"code" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) CodeIn(values []int32) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsCode(vals),
	)

	return qs
}

type notinDhcp6optionsCode []interface{}

func (in notinDhcp6optionsCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"code" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) CodeNotIn(values []int32) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsCode(vals),
	)

	return qs
}

// OrderByCode sorts result by Code in ascending order
func (qs Dhcp6optionsQS) OrderByCode() Dhcp6optionsQS {
	qs.order = append(qs.order, `"code"`)

	return qs
}

// OrderByCodeDesc sorts result by Code in descending order
func (qs Dhcp6optionsQS) OrderByCodeDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"code" DESC`)

	return qs
}

// DistinctOnCode marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnCode() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"code"`)

	return qs
}

// END - django_kea.Dhcp6Options.code

// BEGIN - django_kea.Dhcp6Options.value

// ValueIsNull filters for Value being null
func (qs Dhcp6optionsQS) ValueIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NULL`,
		},
	)
	return qs
}

// ValueIsNotNull filters for Value being not null
func (qs Dhcp6optionsQS) ValueIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NOT NULL`,
		},
	)
	return qs
}

// ValueEq filters for Value being equal to argument
func (qs Dhcp6optionsQS) ValueEq(v string) Dhcp6optionsQS {
	return qs.filter(`"value" =`, v)
}

// ValueNe filters for Value being not equal to argument
func (qs Dhcp6optionsQS) ValueNe(v string) Dhcp6optionsQS {
	return qs.filter(`"value" <>`, v)
}

// ValueLt filters for Value being less than argument
func (qs Dhcp6optionsQS) ValueLt(v string) Dhcp6optionsQS {
	return qs.filter(`"value" <`, v)
}

// ValueLe filters for Value being less than or equal to argument
func (qs Dhcp6optionsQS) ValueLe(v string) Dhcp6optionsQS {
	return qs.filter(`"value" <=`, v)
}

// ValueGt filters for Value being greater than argument
func (qs Dhcp6optionsQS) ValueGt(v string) Dhcp6optionsQS {
	return qs.filter(`"value" >`, v)
}

// ValueGe filters for Value being greater than or equal to argument
func (qs Dhcp6optionsQS) ValueGe(v string) Dhcp6optionsQS {
	return qs.filter(`"value" >=`, v)
}

type inDhcp6optionsValue []interface{}

func (in inDhcp6optionsValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"value" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) ValueIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsValue(vals),
	)

	return qs
}

type notinDhcp6optionsValue []interface{}

func (in notinDhcp6optionsValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"value" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) ValueNotIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsValue(vals),
	)

	return qs
}

// OrderByValue sorts result by Value in ascending order
func (qs Dhcp6optionsQS) OrderByValue() Dhcp6optionsQS {
	qs.order = append(qs.order, `"value"`)

	return qs
}

// OrderByValueDesc sorts result by Value in descending order
func (qs Dhcp6optionsQS) OrderByValueDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"value" DESC`)

	return qs
}

// DistinctOnValue marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnValue() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"value"`)

	return qs
}

// END - django_kea.Dhcp6Options.value

// BEGIN - django_kea.Dhcp6Options.formatted_value

// FormattedValueIsNull filters for FormattedValue being null
func (qs Dhcp6optionsQS) FormattedValueIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NULL`,
		},
	)
	return qs
}

// FormattedValueIsNotNull filters for FormattedValue being not null
func (qs Dhcp6optionsQS) FormattedValueIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NOT NULL`,
		},
	)
	return qs
}

// FormattedValueEq filters for FormattedValue being equal to argument
func (qs Dhcp6optionsQS) FormattedValueEq(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" =`, v)
}

// FormattedValueNe filters for FormattedValue being not equal to argument
func (qs Dhcp6optionsQS) FormattedValueNe(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" <>`, v)
}

// FormattedValueLt filters for FormattedValue being less than argument
func (qs Dhcp6optionsQS) FormattedValueLt(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" <`, v)
}

// FormattedValueLe filters for FormattedValue being less than or equal to argument
func (qs Dhcp6optionsQS) FormattedValueLe(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" <=`, v)
}

// FormattedValueGt filters for FormattedValue being greater than argument
func (qs Dhcp6optionsQS) FormattedValueGt(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" >`, v)
}

// FormattedValueGe filters for FormattedValue being greater than or equal to argument
func (qs Dhcp6optionsQS) FormattedValueGe(v string) Dhcp6optionsQS {
	return qs.filter(`"formatted_value" >=`, v)
}

type inDhcp6optionsFormattedValue []interface{}

func (in inDhcp6optionsFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"formatted_value" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) FormattedValueIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsFormattedValue(vals),
	)

	return qs
}

type notinDhcp6optionsFormattedValue []interface{}

func (in notinDhcp6optionsFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"formatted_value" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) FormattedValueNotIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsFormattedValue(vals),
	)

	return qs
}

// OrderByFormattedValue sorts result by FormattedValue in ascending order
func (qs Dhcp6optionsQS) OrderByFormattedValue() Dhcp6optionsQS {
	qs.order = append(qs.order, `"formatted_value"`)

	return qs
}

// OrderByFormattedValueDesc sorts result by FormattedValue in descending order
func (qs Dhcp6optionsQS) OrderByFormattedValueDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"formatted_value" DESC`)

	return qs
}

// DistinctOnFormattedValue marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnFormattedValue() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"formatted_value"`)

	return qs
}

// END - django_kea.Dhcp6Options.formatted_value

// BEGIN - django_kea.Dhcp6Options.space

// SpaceIsNull filters for Space being null
func (qs Dhcp6optionsQS) SpaceIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NULL`,
		},
	)
	return qs
}

// SpaceIsNotNull filters for Space being not null
func (qs Dhcp6optionsQS) SpaceIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NOT NULL`,
		},
	)
	return qs
}

// SpaceEq filters for Space being equal to argument
func (qs Dhcp6optionsQS) SpaceEq(v string) Dhcp6optionsQS {
	return qs.filter(`"space" =`, v)
}

// SpaceNe filters for Space being not equal to argument
func (qs Dhcp6optionsQS) SpaceNe(v string) Dhcp6optionsQS {
	return qs.filter(`"space" <>`, v)
}

// SpaceLt filters for Space being less than argument
func (qs Dhcp6optionsQS) SpaceLt(v string) Dhcp6optionsQS {
	return qs.filter(`"space" <`, v)
}

// SpaceLe filters for Space being less than or equal to argument
func (qs Dhcp6optionsQS) SpaceLe(v string) Dhcp6optionsQS {
	return qs.filter(`"space" <=`, v)
}

// SpaceGt filters for Space being greater than argument
func (qs Dhcp6optionsQS) SpaceGt(v string) Dhcp6optionsQS {
	return qs.filter(`"space" >`, v)
}

// SpaceGe filters for Space being greater than or equal to argument
func (qs Dhcp6optionsQS) SpaceGe(v string) Dhcp6optionsQS {
	return qs.filter(`"space" >=`, v)
}

type inDhcp6optionsSpace []interface{}

func (in inDhcp6optionsSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"space" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) SpaceIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsSpace(vals),
	)

	return qs
}

type notinDhcp6optionsSpace []interface{}

func (in notinDhcp6optionsSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"space" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) SpaceNotIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsSpace(vals),
	)

	return qs
}

// OrderBySpace sorts result by Space in ascending order
func (qs Dhcp6optionsQS) OrderBySpace() Dhcp6optionsQS {
	qs.order = append(qs.order, `"space"`)

	return qs
}

// OrderBySpaceDesc sorts result by Space in descending order
func (qs Dhcp6optionsQS) OrderBySpaceDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"space" DESC`)

	return qs
}

// DistinctOnSpace marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnSpace() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"space"`)

	return qs
}

// END - django_kea.Dhcp6Options.space

// BEGIN - django_kea.Dhcp6Options.persistent

// PersistentEq filters for Persistent being equal to argument
func (qs Dhcp6optionsQS) PersistentEq(v bool) Dhcp6optionsQS {
	return qs.filter(`"persistent" =`, v)
}

// PersistentNe filters for Persistent being not equal to argument
func (qs Dhcp6optionsQS) PersistentNe(v bool) Dhcp6optionsQS {
	return qs.filter(`"persistent" <>`, v)
}

type inDhcp6optionsPersistent []interface{}

func (in inDhcp6optionsPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"persistent" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) PersistentIn(values []bool) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsPersistent(vals),
	)

	return qs
}

type notinDhcp6optionsPersistent []interface{}

func (in notinDhcp6optionsPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"persistent" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) PersistentNotIn(values []bool) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsPersistent(vals),
	)

	return qs
}

// OrderByPersistent sorts result by Persistent in ascending order
func (qs Dhcp6optionsQS) OrderByPersistent() Dhcp6optionsQS {
	qs.order = append(qs.order, `"persistent"`)

	return qs
}

// OrderByPersistentDesc sorts result by Persistent in descending order
func (qs Dhcp6optionsQS) OrderByPersistentDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"persistent" DESC`)

	return qs
}

// DistinctOnPersistent marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnPersistent() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"persistent"`)

	return qs
}

// END - django_kea.Dhcp6Options.persistent

// BEGIN - django_kea.Dhcp6Options.dhcp_client_class

// DhcpClientClassIsNull filters for DhcpClientClass being null
func (qs Dhcp6optionsQS) DhcpClientClassIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NULL`,
		},
	)
	return qs
}

// DhcpClientClassIsNotNull filters for DhcpClientClass being not null
func (qs Dhcp6optionsQS) DhcpClientClassIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NOT NULL`,
		},
	)
	return qs
}

// DhcpClientClassEq filters for DhcpClientClass being equal to argument
func (qs Dhcp6optionsQS) DhcpClientClassEq(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" =`, v)
}

// DhcpClientClassNe filters for DhcpClientClass being not equal to argument
func (qs Dhcp6optionsQS) DhcpClientClassNe(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" <>`, v)
}

// DhcpClientClassLt filters for DhcpClientClass being less than argument
func (qs Dhcp6optionsQS) DhcpClientClassLt(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" <`, v)
}

// DhcpClientClassLe filters for DhcpClientClass being less than or equal to argument
func (qs Dhcp6optionsQS) DhcpClientClassLe(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" <=`, v)
}

// DhcpClientClassGt filters for DhcpClientClass being greater than argument
func (qs Dhcp6optionsQS) DhcpClientClassGt(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" >`, v)
}

// DhcpClientClassGe filters for DhcpClientClass being greater than or equal to argument
func (qs Dhcp6optionsQS) DhcpClientClassGe(v string) Dhcp6optionsQS {
	return qs.filter(`"dhcp_client_class" >=`, v)
}

type inDhcp6optionsDhcpClientClass []interface{}

func (in inDhcp6optionsDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) DhcpClientClassIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsDhcpClientClass(vals),
	)

	return qs
}

type notinDhcp6optionsDhcpClientClass []interface{}

func (in notinDhcp6optionsDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) DhcpClientClassNotIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsDhcpClientClass(vals),
	)

	return qs
}

// OrderByDhcpClientClass sorts result by DhcpClientClass in ascending order
func (qs Dhcp6optionsQS) OrderByDhcpClientClass() Dhcp6optionsQS {
	qs.order = append(qs.order, `"dhcp_client_class"`)

	return qs
}

// OrderByDhcpClientClassDesc sorts result by DhcpClientClass in descending order
func (qs Dhcp6optionsQS) OrderByDhcpClientClassDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"dhcp_client_class" DESC`)

	return qs
}

// DistinctOnDhcpClientClass marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnDhcpClientClass() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp_client_class"`)

	return qs
}

// END - django_kea.Dhcp6Options.dhcp_client_class

// BEGIN - django_kea.Dhcp6Options.dhcp6_subnet_id

// Dhcp6SubnetIdIsNull filters for Dhcp6SubnetId being null
func (qs Dhcp6optionsQS) Dhcp6SubnetIdIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdIsNotNull filters for Dhcp6SubnetId being not null
func (qs Dhcp6optionsQS) Dhcp6SubnetIdIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdEq filters for Dhcp6SubnetId being equal to argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdEq(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" =`, v)
}

// Dhcp6SubnetIdNe filters for Dhcp6SubnetId being not equal to argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdNe(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" <>`, v)
}

// Dhcp6SubnetIdLt filters for Dhcp6SubnetId being less than argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdLt(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" <`, v)
}

// Dhcp6SubnetIdLe filters for Dhcp6SubnetId being less than or equal to argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdLe(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" <=`, v)
}

// Dhcp6SubnetIdGt filters for Dhcp6SubnetId being greater than argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdGt(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" >`, v)
}

// Dhcp6SubnetIdGe filters for Dhcp6SubnetId being greater than or equal to argument
func (qs Dhcp6optionsQS) Dhcp6SubnetIdGe(v int64) Dhcp6optionsQS {
	return qs.filter(`"dhcp6_subnet_id" >=`, v)
}

type inDhcp6optionsDhcp6SubnetId []interface{}

func (in inDhcp6optionsDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) Dhcp6SubnetIdIn(values []int64) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsDhcp6SubnetId(vals),
	)

	return qs
}

type notinDhcp6optionsDhcp6SubnetId []interface{}

func (in notinDhcp6optionsDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) Dhcp6SubnetIdNotIn(values []int64) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsDhcp6SubnetId(vals),
	)

	return qs
}

// OrderByDhcp6SubnetId sorts result by Dhcp6SubnetId in ascending order
func (qs Dhcp6optionsQS) OrderByDhcp6SubnetId() Dhcp6optionsQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id"`)

	return qs
}

// OrderByDhcp6SubnetIdDesc sorts result by Dhcp6SubnetId in descending order
func (qs Dhcp6optionsQS) OrderByDhcp6SubnetIdDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id" DESC`)

	return qs
}

// DistinctOnDhcp6SubnetId marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnDhcp6SubnetId() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp6_subnet_id"`)

	return qs
}

// END - django_kea.Dhcp6Options.dhcp6_subnet_id

// BEGIN - django_kea.Dhcp6Options.host

// GetHost returns Hosts
func (d *Dhcp6options) GetHost(ctx context.Context, db models.DBInterface) (*Hosts, error) {
	if !d.host.Valid {
		return nil, nil
	}

	return HostsQS{}.HostIdEq(d.host.Int32).First(ctx, db)
}

// SetHost sets foreign key pointer to Hosts
func (d *Dhcp6options) SetHost(ptr *Hosts) error {
	if ptr != nil {
		d.host.Int32 = ptr.GetHostId()
		d.host.Valid = true
	} else {
		d.host.Valid = false
	}

	return nil
}

// GetHostRaw returns Dhcp6options.Host
func (d *Dhcp6options) GetHostRaw() sql.NullInt32 {
	return d.host
}

// HostIsNull filters for host being null
func (qs Dhcp6optionsQS) HostIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NULL`,
		},
	)
	return qs
}

// HostIsNotNull filters for host being not null
func (qs Dhcp6optionsQS) HostIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NOT NULL`,
		},
	)
	return qs
}

// HostEq filters for host being equal to argument
func (qs Dhcp6optionsQS) HostEq(v *Hosts) Dhcp6optionsQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Dhcp6optionsQS) HostRawEq(v int32) Dhcp6optionsQS {
	return qs.filter(`"host_id" =`, v)
}

type inDhcp6optionshostHosts struct {
	qs HostsQS
}

func (in *inDhcp6optionshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Dhcp6optionsQS) HostIn(oqs HostsQS) Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionshostHosts{
			qs: oqs,
		},
	)

	return qs
}

type notinDhcp6optionshostHosts struct {
	qs HostsQS
}

func (nin *notinDhcp6optionshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"host_id" NOT IN (` + s + `)`, p
}

func (qs Dhcp6optionsQS) HostNotIn(oqs HostsQS) Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionshostHosts{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Dhcp6optionsQS) OrderByHost() Dhcp6optionsQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Dhcp6optionsQS) OrderByHostDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// DistinctOnHost marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnHost() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"host_id"`)

	return qs
}

// END - django_kea.Dhcp6Options.host

// BEGIN - django_kea.Dhcp6Options.scope

// GetScope returns Dhcpoptionscope
func (d *Dhcp6options) GetScope(ctx context.Context, db models.DBInterface) (*Dhcpoptionscope, error) {
	return DhcpoptionscopeQS{}.ScopeIdEq(d.scope).First(ctx, db)
}

// SetScope sets foreign key pointer to Dhcpoptionscope
func (d *Dhcp6options) SetScope(ptr *Dhcpoptionscope) error {
	if ptr != nil {
		d.scope = ptr.ScopeId
	} else {
		return fmt.Errorf("Dhcp6options.SetScope: non-null field received null value")
	}

	return nil
}

// GetScopeRaw returns Dhcp6options.Scope
func (d *Dhcp6options) GetScopeRaw() int32 {
	return d.scope
}

// ScopeEq filters for scope being equal to argument
func (qs Dhcp6optionsQS) ScopeEq(v *Dhcpoptionscope) Dhcp6optionsQS {
	return qs.filter(`"scope_id" =`, v.ScopeId)
}

// ScopeRawEq filters for scope being equal to raw argument
func (qs Dhcp6optionsQS) ScopeRawEq(v int32) Dhcp6optionsQS {
	return qs.filter(`"scope_id" =`, v)
}

type inDhcp6optionsscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (in *inDhcp6optionsscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"scope_id" IN (` + s + `)`, p
}

func (qs Dhcp6optionsQS) ScopeIn(oqs DhcpoptionscopeQS) Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionsscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

type notinDhcp6optionsscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (nin *notinDhcp6optionsscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"scope_id" NOT IN (` + s + `)`, p
}

func (qs Dhcp6optionsQS) ScopeNotIn(oqs DhcpoptionscopeQS) Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionsscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

// OrderByScope sorts result by Scope in ascending order
func (qs Dhcp6optionsQS) OrderByScope() Dhcp6optionsQS {
	qs.order = append(qs.order, `"scope_id"`)

	return qs
}

// OrderByScopeDesc sorts result by Scope in descending order
func (qs Dhcp6optionsQS) OrderByScopeDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"scope_id" DESC`)

	return qs
}

// DistinctOnScope marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnScope() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"scope_id"`)

	return qs
}

// END - django_kea.Dhcp6Options.scope

// BEGIN - django_kea.Dhcp6Options.user_context

// UserContextIsNull filters for UserContext being null
func (qs Dhcp6optionsQS) UserContextIsNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Dhcp6optionsQS) UserContextIsNotNull() Dhcp6optionsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Dhcp6optionsQS) UserContextEq(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Dhcp6optionsQS) UserContextNe(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Dhcp6optionsQS) UserContextLt(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Dhcp6optionsQS) UserContextLe(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Dhcp6optionsQS) UserContextGt(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Dhcp6optionsQS) UserContextGe(v string) Dhcp6optionsQS {
	return qs.filter(`"user_context" >=`, v)
}

type inDhcp6optionsUserContext []interface{}

func (in inDhcp6optionsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) UserContextIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsUserContext(vals),
	)

	return qs
}

type notinDhcp6optionsUserContext []interface{}

func (in notinDhcp6optionsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) UserContextNotIn(values []string) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsUserContext(vals),
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Dhcp6optionsQS) OrderByUserContext() Dhcp6optionsQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Dhcp6optionsQS) OrderByUserContextDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// DistinctOnUserContext marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnUserContext() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"user_context"`)

	return qs
}

// END - django_kea.Dhcp6Options.user_context

// BEGIN - django_kea.Dhcp6Options.cancelled

// CancelledEq filters for Cancelled being equal to argument
func (qs Dhcp6optionsQS) CancelledEq(v bool) Dhcp6optionsQS {
	return qs.filter(`"cancelled" =`, v)
}

// CancelledNe filters for Cancelled being not equal to argument
func (qs Dhcp6optionsQS) CancelledNe(v bool) Dhcp6optionsQS {
	return qs.filter(`"cancelled" <>`, v)
}

type inDhcp6optionsCancelled []interface{}

func (in inDhcp6optionsCancelled) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"cancelled" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) CancelledIn(values []bool) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDhcp6optionsCancelled(vals),
	)

	return qs
}

type notinDhcp6optionsCancelled []interface{}

func (in notinDhcp6optionsCancelled) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"cancelled" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Dhcp6optionsQS) CancelledNotIn(values []bool) Dhcp6optionsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDhcp6optionsCancelled(vals),
	)

	return qs
}

// OrderByCancelled sorts result by Cancelled in ascending order
func (qs Dhcp6optionsQS) OrderByCancelled() Dhcp6optionsQS {
	qs.order = append(qs.order, `"cancelled"`)

	return qs
}

// OrderByCancelledDesc sorts result by Cancelled in descending order
func (qs Dhcp6optionsQS) OrderByCancelledDesc() Dhcp6optionsQS {
	qs.order = append(qs.order, `"cancelled" DESC`)

	return qs
}

// DistinctOnCancelled marks field in queries to add to DISTINCT ON clause
func (qs Dhcp6optionsQS) DistinctOnCancelled() Dhcp6optionsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"cancelled"`)

	return qs
}

// END - django_kea.Dhcp6Options.cancelled

// OrderByRandom randomizes result
func (qs Dhcp6optionsQS) OrderByRandom() Dhcp6optionsQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Dhcp6optionsQS) ForUpdate() Dhcp6optionsQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs Dhcp6optionsQS) ForUpdateNowait() Dhcp6optionsQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs Dhcp6optionsQS) ForUpdateSkipLocked() Dhcp6optionsQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs Dhcp6optionsQS) ClearForUpdate() Dhcp6optionsQS {
	qs.forClause = ""

	return qs
}

func (qs Dhcp6optionsQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Dhcp6optionsQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Dhcp6optionsQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"option_id", "code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp6_subnet_id", "host_id", "scope_id", "user_context", "cancelled" FROM "dhcp6_options"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Dhcp6optionsQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "option_id" FROM "dhcp6_options"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs Dhcp6optionsQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"option_id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "dhcp6_options"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs Dhcp6optionsQS) All(ctx context.Context, db models.DBInterface) (Dhcp6optionsList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret Dhcp6optionsList
	for rows.Next() {
		obj := Dhcp6options{existsInDB: true}
		if err = rows.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp6SubnetId, &obj.host, &obj.scope, &obj.UserContext, &obj.Cancelled); err != nil {
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
func (qs Dhcp6optionsQS) First(ctx context.Context, db models.DBInterface) (*Dhcp6options, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Dhcp6options{existsInDB: true}
	err := row.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp6SubnetId, &obj.host, &obj.scope, &obj.UserContext, &obj.Cancelled)
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
func (qs Dhcp6optionsQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp6_options"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Dhcp6optionsQS) Update() Dhcp6optionsUpdateQS {
	return Dhcp6optionsUpdateQS{condFragments: qs.condFragments}
}

// Dhcp6optionsUpdateQS represents an updated queryset for django_kea.Dhcp6Options
type Dhcp6optionsUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Dhcp6optionsUpdateQS) update(c string, v interface{}) Dhcp6optionsUpdateQS {
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
func (uqs Dhcp6optionsUpdateQS) SetOptionId(v int32) Dhcp6optionsUpdateQS {
	return uqs.update(`"option_id"`, v)
}

// SetCode sets Code to the given value
func (uqs Dhcp6optionsUpdateQS) SetCode(v int32) Dhcp6optionsUpdateQS {
	return uqs.update(`"code"`, v)
}

// SetValue sets Value to the given value
func (uqs Dhcp6optionsUpdateQS) SetValue(v sql.NullString) Dhcp6optionsUpdateQS {
	return uqs.update(`"value"`, v)
}

// SetFormattedValue sets FormattedValue to the given value
func (uqs Dhcp6optionsUpdateQS) SetFormattedValue(v sql.NullString) Dhcp6optionsUpdateQS {
	return uqs.update(`"formatted_value"`, v)
}

// SetSpace sets Space to the given value
func (uqs Dhcp6optionsUpdateQS) SetSpace(v sql.NullString) Dhcp6optionsUpdateQS {
	return uqs.update(`"space"`, v)
}

// SetPersistent sets Persistent to the given value
func (uqs Dhcp6optionsUpdateQS) SetPersistent(v bool) Dhcp6optionsUpdateQS {
	return uqs.update(`"persistent"`, v)
}

// SetDhcpClientClass sets DhcpClientClass to the given value
func (uqs Dhcp6optionsUpdateQS) SetDhcpClientClass(v sql.NullString) Dhcp6optionsUpdateQS {
	return uqs.update(`"dhcp_client_class"`, v)
}

// SetDhcp6SubnetId sets Dhcp6SubnetId to the given value
func (uqs Dhcp6optionsUpdateQS) SetDhcp6SubnetId(v sql.NullInt64) Dhcp6optionsUpdateQS {
	return uqs.update(`"dhcp6_subnet_id"`, v)
}

// SetHost sets foreign key pointer to Hosts
func (uqs Dhcp6optionsUpdateQS) SetHost(ptr *Hosts) Dhcp6optionsUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // SetScope sets foreign key pointer to Dhcpoptionscope
func (uqs Dhcp6optionsUpdateQS) SetScope(ptr *Dhcpoptionscope) Dhcp6optionsUpdateQS {
	if ptr != nil {
		return uqs.update(`"scope_id"`, ptr.ScopeId)
	}

	return uqs.update(`"scope_id"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Dhcp6optionsUpdateQS) SetUserContext(v sql.NullString) Dhcp6optionsUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// SetCancelled sets Cancelled to the given value
func (uqs Dhcp6optionsUpdateQS) SetCancelled(v bool) Dhcp6optionsUpdateQS {
	return uqs.update(`"cancelled"`, v)
}

// Exec executes the update operation
func (uqs Dhcp6optionsUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := Dhcp6optionsQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "dhcp6_options" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (d *Dhcp6options) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "dhcp6_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp6_subnet_id", "host_id", "scope_id", "user_context", "cancelled") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING "option_id"`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp6SubnetId, d.host, d.scope, d.UserContext, d.Cancelled)

	if err := row.Scan(&d.optionId); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcp6options) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "dhcp6_options" SET "code" = $1, "value" = $2, "formatted_value" = $3, "space" = $4, "persistent" = $5, "dhcp_client_class" = $6, "dhcp6_subnet_id" = $7, "host_id" = $8, "scope_id" = $9, "user_context" = $10, "cancelled" = $11 WHERE "option_id" = $12`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp6SubnetId, d.host, d.scope, d.UserContext, d.Cancelled, d.optionId)

	return err
}

// Save inserts or updates record
func (d *Dhcp6options) Save(ctx context.Context, db models.DBInterface) error {
	if d.existsInDB {
		return d.update(ctx, db)
	}

	return d.insert(ctx, db)
}

// Delete removes row from database
func (d *Dhcp6options) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "dhcp6_options" WHERE "option_id" = $1`, d.optionId)

	d.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (dl Dhcp6optionsList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts Dhcp6optionsList

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
	vaa := make([]any, 0, 11*len(inserts))
	offs := 1
	for _, d := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4, offs+5, offs+6, offs+7, offs+8, offs+9, offs+10))
		vaa = append(vaa, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp6SubnetId, d.host, d.scope, d.UserContext, d.Cancelled)
		offs += 11
	}

	qs := `INSERT INTO "dhcp6_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp6_subnet_id", "host_id", "scope_id", "user_context", "cancelled") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "option_id"`
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
