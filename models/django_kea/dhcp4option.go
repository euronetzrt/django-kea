/*
  AUTO-GENERATED file for Django model django_kea.Dhcp4Option

  Command used to generate:

  DJANGO_SETTINGS_MODULE=keatest.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/euronetzrt/django-kea django_kea

  https://github.com/rkojedzinszky/djan-go-rm
*/

package django_kea

import (
	"database/sql"
	"fmt"
	"github.com/euronetzrt/django-kea/models"
	"strings"
)

// Dhcp4option mirrors model django_kea.Dhcp4Option
type Dhcp4option struct {
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
}

// Dhcp4optionQS represents a queryset for django_kea.Dhcp4Option
type Dhcp4optionQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Dhcp4optionQS) filter(c string, p interface{}) Dhcp4optionQS {
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
func (qs Dhcp4optionQS) Or(exprs ...Dhcp4optionQS) Dhcp4optionQS {
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

// GetOptionId returns Dhcp4option.OptionId
func (d *Dhcp4option) GetOptionId() int32 {
	return d.optionId
}

// OptionIdEq filters for optionId being equal to argument
func (qs Dhcp4optionQS) OptionIdEq(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" =`, v)
}

// OptionIdNe filters for optionId being not equal to argument
func (qs Dhcp4optionQS) OptionIdNe(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" <>`, v)
}

// OptionIdLt filters for optionId being less than argument
func (qs Dhcp4optionQS) OptionIdLt(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" <`, v)
}

// OptionIdLe filters for optionId being less than or equal to argument
func (qs Dhcp4optionQS) OptionIdLe(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" <=`, v)
}

// OptionIdGt filters for optionId being greater than argument
func (qs Dhcp4optionQS) OptionIdGt(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" >`, v)
}

// OptionIdGe filters for optionId being greater than or equal to argument
func (qs Dhcp4optionQS) OptionIdGe(v int32) Dhcp4optionQS {
	return qs.filter(`"option_id" >=`, v)
}

type inDhcp4optionoptionId struct {
	values []interface{}
}

func (in *inDhcp4optionoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"option_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) OptionIdIn(values []int32) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionoptionId{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionoptionId struct {
	values []interface{}
}

func (in *notinDhcp4optionoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"option_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) OptionIdNotIn(values []int32) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionoptionId{
			values: vals,
		},
	)

	return qs
}

// OrderByOptionId sorts result by OptionId in ascending order
func (qs Dhcp4optionQS) OrderByOptionId() Dhcp4optionQS {
	qs.order = append(qs.order, `"option_id"`)

	return qs
}

// OrderByOptionIdDesc sorts result by OptionId in descending order
func (qs Dhcp4optionQS) OrderByOptionIdDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"option_id" DESC`)

	return qs
}

// CodeEq filters for Code being equal to argument
func (qs Dhcp4optionQS) CodeEq(v int32) Dhcp4optionQS {
	return qs.filter(`"code" =`, v)
}

// CodeNe filters for Code being not equal to argument
func (qs Dhcp4optionQS) CodeNe(v int32) Dhcp4optionQS {
	return qs.filter(`"code" <>`, v)
}

// CodeLt filters for Code being less than argument
func (qs Dhcp4optionQS) CodeLt(v int32) Dhcp4optionQS {
	return qs.filter(`"code" <`, v)
}

// CodeLe filters for Code being less than or equal to argument
func (qs Dhcp4optionQS) CodeLe(v int32) Dhcp4optionQS {
	return qs.filter(`"code" <=`, v)
}

// CodeGt filters for Code being greater than argument
func (qs Dhcp4optionQS) CodeGt(v int32) Dhcp4optionQS {
	return qs.filter(`"code" >`, v)
}

// CodeGe filters for Code being greater than or equal to argument
func (qs Dhcp4optionQS) CodeGe(v int32) Dhcp4optionQS {
	return qs.filter(`"code" >=`, v)
}

type inDhcp4optionCode struct {
	values []interface{}
}

func (in *inDhcp4optionCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"code" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) CodeIn(values []int32) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionCode{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionCode struct {
	values []interface{}
}

func (in *notinDhcp4optionCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"code" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) CodeNotIn(values []int32) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionCode{
			values: vals,
		},
	)

	return qs
}

// OrderByCode sorts result by Code in ascending order
func (qs Dhcp4optionQS) OrderByCode() Dhcp4optionQS {
	qs.order = append(qs.order, `"code"`)

	return qs
}

// OrderByCodeDesc sorts result by Code in descending order
func (qs Dhcp4optionQS) OrderByCodeDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"code" DESC`)

	return qs
}

// ValueIsNull filters for Value being null
func (qs Dhcp4optionQS) ValueIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NULL`,
		},
	)
	return qs
}

// ValueIsNotNull filters for Value being not null
func (qs Dhcp4optionQS) ValueIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NOT NULL`,
		},
	)
	return qs
}

// ValueEq filters for Value being equal to argument
func (qs Dhcp4optionQS) ValueEq(v string) Dhcp4optionQS {
	return qs.filter(`"value" =`, v)
}

// ValueNe filters for Value being not equal to argument
func (qs Dhcp4optionQS) ValueNe(v string) Dhcp4optionQS {
	return qs.filter(`"value" <>`, v)
}

// ValueLt filters for Value being less than argument
func (qs Dhcp4optionQS) ValueLt(v string) Dhcp4optionQS {
	return qs.filter(`"value" <`, v)
}

// ValueLe filters for Value being less than or equal to argument
func (qs Dhcp4optionQS) ValueLe(v string) Dhcp4optionQS {
	return qs.filter(`"value" <=`, v)
}

// ValueGt filters for Value being greater than argument
func (qs Dhcp4optionQS) ValueGt(v string) Dhcp4optionQS {
	return qs.filter(`"value" >`, v)
}

// ValueGe filters for Value being greater than or equal to argument
func (qs Dhcp4optionQS) ValueGe(v string) Dhcp4optionQS {
	return qs.filter(`"value" >=`, v)
}

type inDhcp4optionValue struct {
	values []interface{}
}

func (in *inDhcp4optionValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"value" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) ValueIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionValue{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionValue struct {
	values []interface{}
}

func (in *notinDhcp4optionValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"value" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) ValueNotIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionValue{
			values: vals,
		},
	)

	return qs
}

// OrderByValue sorts result by Value in ascending order
func (qs Dhcp4optionQS) OrderByValue() Dhcp4optionQS {
	qs.order = append(qs.order, `"value"`)

	return qs
}

// OrderByValueDesc sorts result by Value in descending order
func (qs Dhcp4optionQS) OrderByValueDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"value" DESC`)

	return qs
}

// FormattedValueIsNull filters for FormattedValue being null
func (qs Dhcp4optionQS) FormattedValueIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NULL`,
		},
	)
	return qs
}

// FormattedValueIsNotNull filters for FormattedValue being not null
func (qs Dhcp4optionQS) FormattedValueIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NOT NULL`,
		},
	)
	return qs
}

// FormattedValueEq filters for FormattedValue being equal to argument
func (qs Dhcp4optionQS) FormattedValueEq(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" =`, v)
}

// FormattedValueNe filters for FormattedValue being not equal to argument
func (qs Dhcp4optionQS) FormattedValueNe(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" <>`, v)
}

// FormattedValueLt filters for FormattedValue being less than argument
func (qs Dhcp4optionQS) FormattedValueLt(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" <`, v)
}

// FormattedValueLe filters for FormattedValue being less than or equal to argument
func (qs Dhcp4optionQS) FormattedValueLe(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" <=`, v)
}

// FormattedValueGt filters for FormattedValue being greater than argument
func (qs Dhcp4optionQS) FormattedValueGt(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" >`, v)
}

// FormattedValueGe filters for FormattedValue being greater than or equal to argument
func (qs Dhcp4optionQS) FormattedValueGe(v string) Dhcp4optionQS {
	return qs.filter(`"formatted_value" >=`, v)
}

type inDhcp4optionFormattedValue struct {
	values []interface{}
}

func (in *inDhcp4optionFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"formatted_value" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) FormattedValueIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionFormattedValue{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionFormattedValue struct {
	values []interface{}
}

func (in *notinDhcp4optionFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"formatted_value" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) FormattedValueNotIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionFormattedValue{
			values: vals,
		},
	)

	return qs
}

// OrderByFormattedValue sorts result by FormattedValue in ascending order
func (qs Dhcp4optionQS) OrderByFormattedValue() Dhcp4optionQS {
	qs.order = append(qs.order, `"formatted_value"`)

	return qs
}

// OrderByFormattedValueDesc sorts result by FormattedValue in descending order
func (qs Dhcp4optionQS) OrderByFormattedValueDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"formatted_value" DESC`)

	return qs
}

// SpaceIsNull filters for Space being null
func (qs Dhcp4optionQS) SpaceIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NULL`,
		},
	)
	return qs
}

// SpaceIsNotNull filters for Space being not null
func (qs Dhcp4optionQS) SpaceIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NOT NULL`,
		},
	)
	return qs
}

// SpaceEq filters for Space being equal to argument
func (qs Dhcp4optionQS) SpaceEq(v string) Dhcp4optionQS {
	return qs.filter(`"space" =`, v)
}

// SpaceNe filters for Space being not equal to argument
func (qs Dhcp4optionQS) SpaceNe(v string) Dhcp4optionQS {
	return qs.filter(`"space" <>`, v)
}

// SpaceLt filters for Space being less than argument
func (qs Dhcp4optionQS) SpaceLt(v string) Dhcp4optionQS {
	return qs.filter(`"space" <`, v)
}

// SpaceLe filters for Space being less than or equal to argument
func (qs Dhcp4optionQS) SpaceLe(v string) Dhcp4optionQS {
	return qs.filter(`"space" <=`, v)
}

// SpaceGt filters for Space being greater than argument
func (qs Dhcp4optionQS) SpaceGt(v string) Dhcp4optionQS {
	return qs.filter(`"space" >`, v)
}

// SpaceGe filters for Space being greater than or equal to argument
func (qs Dhcp4optionQS) SpaceGe(v string) Dhcp4optionQS {
	return qs.filter(`"space" >=`, v)
}

type inDhcp4optionSpace struct {
	values []interface{}
}

func (in *inDhcp4optionSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"space" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) SpaceIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionSpace{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionSpace struct {
	values []interface{}
}

func (in *notinDhcp4optionSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"space" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) SpaceNotIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionSpace{
			values: vals,
		},
	)

	return qs
}

// OrderBySpace sorts result by Space in ascending order
func (qs Dhcp4optionQS) OrderBySpace() Dhcp4optionQS {
	qs.order = append(qs.order, `"space"`)

	return qs
}

// OrderBySpaceDesc sorts result by Space in descending order
func (qs Dhcp4optionQS) OrderBySpaceDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"space" DESC`)

	return qs
}

// PersistentEq filters for Persistent being equal to argument
func (qs Dhcp4optionQS) PersistentEq(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" =`, v)
}

// PersistentNe filters for Persistent being not equal to argument
func (qs Dhcp4optionQS) PersistentNe(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" <>`, v)
}

// PersistentLt filters for Persistent being less than argument
func (qs Dhcp4optionQS) PersistentLt(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" <`, v)
}

// PersistentLe filters for Persistent being less than or equal to argument
func (qs Dhcp4optionQS) PersistentLe(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" <=`, v)
}

// PersistentGt filters for Persistent being greater than argument
func (qs Dhcp4optionQS) PersistentGt(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" >`, v)
}

// PersistentGe filters for Persistent being greater than or equal to argument
func (qs Dhcp4optionQS) PersistentGe(v bool) Dhcp4optionQS {
	return qs.filter(`"persistent" >=`, v)
}

type inDhcp4optionPersistent struct {
	values []interface{}
}

func (in *inDhcp4optionPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"persistent" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) PersistentIn(values []bool) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionPersistent{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionPersistent struct {
	values []interface{}
}

func (in *notinDhcp4optionPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"persistent" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) PersistentNotIn(values []bool) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionPersistent{
			values: vals,
		},
	)

	return qs
}

// OrderByPersistent sorts result by Persistent in ascending order
func (qs Dhcp4optionQS) OrderByPersistent() Dhcp4optionQS {
	qs.order = append(qs.order, `"persistent"`)

	return qs
}

// OrderByPersistentDesc sorts result by Persistent in descending order
func (qs Dhcp4optionQS) OrderByPersistentDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"persistent" DESC`)

	return qs
}

// DhcpClientClassIsNull filters for DhcpClientClass being null
func (qs Dhcp4optionQS) DhcpClientClassIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NULL`,
		},
	)
	return qs
}

// DhcpClientClassIsNotNull filters for DhcpClientClass being not null
func (qs Dhcp4optionQS) DhcpClientClassIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NOT NULL`,
		},
	)
	return qs
}

// DhcpClientClassEq filters for DhcpClientClass being equal to argument
func (qs Dhcp4optionQS) DhcpClientClassEq(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" =`, v)
}

// DhcpClientClassNe filters for DhcpClientClass being not equal to argument
func (qs Dhcp4optionQS) DhcpClientClassNe(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" <>`, v)
}

// DhcpClientClassLt filters for DhcpClientClass being less than argument
func (qs Dhcp4optionQS) DhcpClientClassLt(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" <`, v)
}

// DhcpClientClassLe filters for DhcpClientClass being less than or equal to argument
func (qs Dhcp4optionQS) DhcpClientClassLe(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" <=`, v)
}

// DhcpClientClassGt filters for DhcpClientClass being greater than argument
func (qs Dhcp4optionQS) DhcpClientClassGt(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" >`, v)
}

// DhcpClientClassGe filters for DhcpClientClass being greater than or equal to argument
func (qs Dhcp4optionQS) DhcpClientClassGe(v string) Dhcp4optionQS {
	return qs.filter(`"dhcp_client_class" >=`, v)
}

type inDhcp4optionDhcpClientClass struct {
	values []interface{}
}

func (in *inDhcp4optionDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) DhcpClientClassIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionDhcpClientClass{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionDhcpClientClass struct {
	values []interface{}
}

func (in *notinDhcp4optionDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) DhcpClientClassNotIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionDhcpClientClass{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcpClientClass sorts result by DhcpClientClass in ascending order
func (qs Dhcp4optionQS) OrderByDhcpClientClass() Dhcp4optionQS {
	qs.order = append(qs.order, `"dhcp_client_class"`)

	return qs
}

// OrderByDhcpClientClassDesc sorts result by DhcpClientClass in descending order
func (qs Dhcp4optionQS) OrderByDhcpClientClassDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"dhcp_client_class" DESC`)

	return qs
}

// Dhcp4SubnetIdIsNull filters for Dhcp4SubnetId being null
func (qs Dhcp4optionQS) Dhcp4SubnetIdIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdIsNotNull filters for Dhcp4SubnetId being not null
func (qs Dhcp4optionQS) Dhcp4SubnetIdIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdEq filters for Dhcp4SubnetId being equal to argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdEq(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" =`, v)
}

// Dhcp4SubnetIdNe filters for Dhcp4SubnetId being not equal to argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdNe(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" <>`, v)
}

// Dhcp4SubnetIdLt filters for Dhcp4SubnetId being less than argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdLt(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" <`, v)
}

// Dhcp4SubnetIdLe filters for Dhcp4SubnetId being less than or equal to argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdLe(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" <=`, v)
}

// Dhcp4SubnetIdGt filters for Dhcp4SubnetId being greater than argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdGt(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" >`, v)
}

// Dhcp4SubnetIdGe filters for Dhcp4SubnetId being greater than or equal to argument
func (qs Dhcp4optionQS) Dhcp4SubnetIdGe(v int64) Dhcp4optionQS {
	return qs.filter(`"dhcp4_subnet_id" >=`, v)
}

type inDhcp4optionDhcp4SubnetId struct {
	values []interface{}
}

func (in *inDhcp4optionDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) Dhcp4SubnetIdIn(values []int64) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionDhcp4SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionDhcp4SubnetId struct {
	values []interface{}
}

func (in *notinDhcp4optionDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) Dhcp4SubnetIdNotIn(values []int64) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionDhcp4SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4SubnetId sorts result by Dhcp4SubnetId in ascending order
func (qs Dhcp4optionQS) OrderByDhcp4SubnetId() Dhcp4optionQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id"`)

	return qs
}

// OrderByDhcp4SubnetIdDesc sorts result by Dhcp4SubnetId in descending order
func (qs Dhcp4optionQS) OrderByDhcp4SubnetIdDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id" DESC`)

	return qs
}

// GetHost returns Host
func (d *Dhcp4option) GetHost(db models.DBInterface) (*Host, error) {
	if !d.host.Valid {
		return nil, nil
	}

	return HostQS{}.HostIdEq(d.host.Int32).First(db)
}

// SetHost sets foreign key pointer to Host
func (d *Dhcp4option) SetHost(ptr *Host) error {
	if ptr != nil {
		d.host.Int32 = ptr.GetHostId()
		d.host.Valid = true
	} else {
		d.host.Valid = false
	}

	return nil
}

// GetHostRaw returns Dhcp4option.Host
func (d *Dhcp4option) GetHostRaw() sql.NullInt32 {
	return d.host
}

// HostIsNull filters for host being null
func (qs Dhcp4optionQS) HostIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NULL`,
		},
	)
	return qs
}

// HostIsNotNull filters for host being not null
func (qs Dhcp4optionQS) HostIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NOT NULL`,
		},
	)
	return qs
}

// HostEq filters for host being equal to argument
func (qs Dhcp4optionQS) HostEq(v *Host) Dhcp4optionQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Dhcp4optionQS) HostRawEq(v int32) Dhcp4optionQS {
	return qs.filter(`"host_id" =`, v)
}

type inDhcp4optionhostHost struct {
	qs HostQS
}

func (in *inDhcp4optionhostHost) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Dhcp4optionQS) HostIn(oqs HostQS) Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionhostHost{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Dhcp4optionQS) OrderByHost() Dhcp4optionQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Dhcp4optionQS) OrderByHostDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// GetScope returns Dhcpoptionscope
func (d *Dhcp4option) GetScope(db models.DBInterface) (*Dhcpoptionscope, error) {
	return DhcpoptionscopeQS{}.ScopeIdEq(d.scope).First(db)
}

// SetScope sets foreign key pointer to Dhcpoptionscope
func (d *Dhcp4option) SetScope(ptr *Dhcpoptionscope) error {
	if ptr != nil {
		d.scope = ptr.ScopeId
	} else {
		return fmt.Errorf("Dhcp4option.SetScope: non-null field received null value")
	}

	return nil
}

// GetScopeRaw returns Dhcp4option.Scope
func (d *Dhcp4option) GetScopeRaw() int32 {
	return d.scope
}

// ScopeEq filters for scope being equal to argument
func (qs Dhcp4optionQS) ScopeEq(v *Dhcpoptionscope) Dhcp4optionQS {
	return qs.filter(`"scope_id" =`, v.ScopeId)
}

// ScopeRawEq filters for scope being equal to raw argument
func (qs Dhcp4optionQS) ScopeRawEq(v int32) Dhcp4optionQS {
	return qs.filter(`"scope_id" =`, v)
}

type inDhcp4optionscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (in *inDhcp4optionscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"scope_id" IN (` + s + `)`, p
}

func (qs Dhcp4optionQS) ScopeIn(oqs DhcpoptionscopeQS) Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

// OrderByScope sorts result by Scope in ascending order
func (qs Dhcp4optionQS) OrderByScope() Dhcp4optionQS {
	qs.order = append(qs.order, `"scope_id"`)

	return qs
}

// OrderByScopeDesc sorts result by Scope in descending order
func (qs Dhcp4optionQS) OrderByScopeDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"scope_id" DESC`)

	return qs
}

// UserContextIsNull filters for UserContext being null
func (qs Dhcp4optionQS) UserContextIsNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Dhcp4optionQS) UserContextIsNotNull() Dhcp4optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Dhcp4optionQS) UserContextEq(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Dhcp4optionQS) UserContextNe(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Dhcp4optionQS) UserContextLt(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Dhcp4optionQS) UserContextLe(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Dhcp4optionQS) UserContextGt(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Dhcp4optionQS) UserContextGe(v string) Dhcp4optionQS {
	return qs.filter(`"user_context" >=`, v)
}

type inDhcp4optionUserContext struct {
	values []interface{}
}

func (in *inDhcp4optionUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) UserContextIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp4optionUserContext{
			values: vals,
		},
	)

	return qs
}

type notinDhcp4optionUserContext struct {
	values []interface{}
}

func (in *notinDhcp4optionUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp4optionQS) UserContextNotIn(values []string) Dhcp4optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp4optionUserContext{
			values: vals,
		},
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Dhcp4optionQS) OrderByUserContext() Dhcp4optionQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Dhcp4optionQS) OrderByUserContextDesc() Dhcp4optionQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Dhcp4optionQS) ForUpdate() Dhcp4optionQS {
	qs.forUpdate = true

	return qs
}

func (qs Dhcp4optionQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Dhcp4optionQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Dhcp4optionQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "option_id", "code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp4_subnet_id", "host_id", "scope_id", "user_context" FROM "dhcp4_options"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Dhcp4optionQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "option_id" FROM "dhcp4_options"` + s, p
}

// All returns all rows matching queryset filters
func (qs Dhcp4optionQS) All(db models.DBInterface) ([]*Dhcp4option, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Dhcp4option
	for rows.Next() {
		obj := Dhcp4option{existsInDB: true}
		if err = rows.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp4SubnetId, &obj.host, &obj.scope, &obj.UserContext); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Dhcp4optionQS) First(db models.DBInterface) (*Dhcp4option, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Dhcp4option{existsInDB: true}
	err := row.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp4SubnetId, &obj.host, &obj.scope, &obj.UserContext)
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
func (qs Dhcp4optionQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp4_options"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Dhcp4optionQS) Update() Dhcp4optionUpdateQS {
	return Dhcp4optionUpdateQS{condFragments: qs.condFragments}
}

// Dhcp4optionUpdateQS represents an updated queryset for django_kea.Dhcp4Option
type Dhcp4optionUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Dhcp4optionUpdateQS) update(c string, v interface{}) Dhcp4optionUpdateQS {
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
func (uqs Dhcp4optionUpdateQS) SetOptionId(v int32) Dhcp4optionUpdateQS {
	return uqs.update(`"option_id"`, v)
}

// SetCode sets Code to the given value
func (uqs Dhcp4optionUpdateQS) SetCode(v int32) Dhcp4optionUpdateQS {
	return uqs.update(`"code"`, v)
}

// SetValue sets Value to the given value
func (uqs Dhcp4optionUpdateQS) SetValue(v sql.NullString) Dhcp4optionUpdateQS {
	return uqs.update(`"value"`, v)
}

// SetFormattedValue sets FormattedValue to the given value
func (uqs Dhcp4optionUpdateQS) SetFormattedValue(v sql.NullString) Dhcp4optionUpdateQS {
	return uqs.update(`"formatted_value"`, v)
}

// SetSpace sets Space to the given value
func (uqs Dhcp4optionUpdateQS) SetSpace(v sql.NullString) Dhcp4optionUpdateQS {
	return uqs.update(`"space"`, v)
}

// SetPersistent sets Persistent to the given value
func (uqs Dhcp4optionUpdateQS) SetPersistent(v bool) Dhcp4optionUpdateQS {
	return uqs.update(`"persistent"`, v)
}

// SetDhcpClientClass sets DhcpClientClass to the given value
func (uqs Dhcp4optionUpdateQS) SetDhcpClientClass(v sql.NullString) Dhcp4optionUpdateQS {
	return uqs.update(`"dhcp_client_class"`, v)
}

// SetDhcp4SubnetId sets Dhcp4SubnetId to the given value
func (uqs Dhcp4optionUpdateQS) SetDhcp4SubnetId(v sql.NullInt64) Dhcp4optionUpdateQS {
	return uqs.update(`"dhcp4_subnet_id"`, v)
}

// SetHost sets foreign key pointer to Host
func (uqs Dhcp4optionUpdateQS) SetHost(ptr *Host) Dhcp4optionUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // SetScope sets foreign key pointer to Dhcpoptionscope
func (uqs Dhcp4optionUpdateQS) SetScope(ptr *Dhcpoptionscope) Dhcp4optionUpdateQS {
	if ptr != nil {
		return uqs.update(`"scope_id"`, ptr.ScopeId)
	}

	return uqs.update(`"scope_id"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Dhcp4optionUpdateQS) SetUserContext(v sql.NullString) Dhcp4optionUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// Exec executes the update operation
func (uqs Dhcp4optionUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Dhcp4optionQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "dhcp4_options" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (d *Dhcp4option) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "dhcp4_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp4_subnet_id", "host_id", "scope_id", "user_context") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "option_id"`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp4SubnetId, d.host, d.scope, d.UserContext)

	if err := row.Scan(&d.optionId); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcp4option) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "dhcp4_options" SET "code" = $1, "value" = $2, "formatted_value" = $3, "space" = $4, "persistent" = $5, "dhcp_client_class" = $6, "dhcp4_subnet_id" = $7, "host_id" = $8, "scope_id" = $9, "user_context" = $10 WHERE "option_id" = $11`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp4SubnetId, d.host, d.scope, d.UserContext, d.optionId)

	return err
}

// Save inserts or updates record
func (d *Dhcp4option) Save(db models.DBInterface) error {
	if d.existsInDB {
		return d.update(db)
	}

	return d.insert(db)
}

// Delete removes row from database
func (d *Dhcp4option) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "dhcp4_options" WHERE "option_id" = $1`, d.optionId)

	d.existsInDB = false

	return err
}
