/*
  AUTO-GENERATED file for Django model django_kea.Dhcp6Option

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

// Dhcp6option mirrors model django_kea.Dhcp6Option
type Dhcp6option struct {
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
}

// Dhcp6optionQS represents a queryset for django_kea.Dhcp6Option
type Dhcp6optionQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Dhcp6optionQS) filter(c string, p interface{}) Dhcp6optionQS {
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
func (qs Dhcp6optionQS) Or(exprs ...Dhcp6optionQS) Dhcp6optionQS {
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

// GetOptionId returns Dhcp6option.OptionId
func (d *Dhcp6option) GetOptionId() int32 {
	return d.optionId
}

// OptionIdEq filters for optionId being equal to argument
func (qs Dhcp6optionQS) OptionIdEq(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" =`, v)
}

// OptionIdNe filters for optionId being not equal to argument
func (qs Dhcp6optionQS) OptionIdNe(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" <>`, v)
}

// OptionIdLt filters for optionId being less than argument
func (qs Dhcp6optionQS) OptionIdLt(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" <`, v)
}

// OptionIdLe filters for optionId being less than or equal to argument
func (qs Dhcp6optionQS) OptionIdLe(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" <=`, v)
}

// OptionIdGt filters for optionId being greater than argument
func (qs Dhcp6optionQS) OptionIdGt(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" >`, v)
}

// OptionIdGe filters for optionId being greater than or equal to argument
func (qs Dhcp6optionQS) OptionIdGe(v int32) Dhcp6optionQS {
	return qs.filter(`"option_id" >=`, v)
}

type inDhcp6optionoptionId struct {
	values []interface{}
}

func (in *inDhcp6optionoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"option_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) OptionIdIn(values []int32) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionoptionId{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionoptionId struct {
	values []interface{}
}

func (in *notinDhcp6optionoptionId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"option_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) OptionIdNotIn(values []int32) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionoptionId{
			values: vals,
		},
	)

	return qs
}

// OrderByOptionId sorts result by OptionId in ascending order
func (qs Dhcp6optionQS) OrderByOptionId() Dhcp6optionQS {
	qs.order = append(qs.order, `"option_id"`)

	return qs
}

// OrderByOptionIdDesc sorts result by OptionId in descending order
func (qs Dhcp6optionQS) OrderByOptionIdDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"option_id" DESC`)

	return qs
}

// CodeEq filters for Code being equal to argument
func (qs Dhcp6optionQS) CodeEq(v int32) Dhcp6optionQS {
	return qs.filter(`"code" =`, v)
}

// CodeNe filters for Code being not equal to argument
func (qs Dhcp6optionQS) CodeNe(v int32) Dhcp6optionQS {
	return qs.filter(`"code" <>`, v)
}

// CodeLt filters for Code being less than argument
func (qs Dhcp6optionQS) CodeLt(v int32) Dhcp6optionQS {
	return qs.filter(`"code" <`, v)
}

// CodeLe filters for Code being less than or equal to argument
func (qs Dhcp6optionQS) CodeLe(v int32) Dhcp6optionQS {
	return qs.filter(`"code" <=`, v)
}

// CodeGt filters for Code being greater than argument
func (qs Dhcp6optionQS) CodeGt(v int32) Dhcp6optionQS {
	return qs.filter(`"code" >`, v)
}

// CodeGe filters for Code being greater than or equal to argument
func (qs Dhcp6optionQS) CodeGe(v int32) Dhcp6optionQS {
	return qs.filter(`"code" >=`, v)
}

type inDhcp6optionCode struct {
	values []interface{}
}

func (in *inDhcp6optionCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"code" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) CodeIn(values []int32) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionCode{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionCode struct {
	values []interface{}
}

func (in *notinDhcp6optionCode) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"code" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) CodeNotIn(values []int32) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionCode{
			values: vals,
		},
	)

	return qs
}

// OrderByCode sorts result by Code in ascending order
func (qs Dhcp6optionQS) OrderByCode() Dhcp6optionQS {
	qs.order = append(qs.order, `"code"`)

	return qs
}

// OrderByCodeDesc sorts result by Code in descending order
func (qs Dhcp6optionQS) OrderByCodeDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"code" DESC`)

	return qs
}

// ValueIsNull filters for Value being null
func (qs Dhcp6optionQS) ValueIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NULL`,
		},
	)
	return qs
}

// ValueIsNotNull filters for Value being not null
func (qs Dhcp6optionQS) ValueIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"value" IS NOT NULL`,
		},
	)
	return qs
}

// ValueEq filters for Value being equal to argument
func (qs Dhcp6optionQS) ValueEq(v string) Dhcp6optionQS {
	return qs.filter(`"value" =`, v)
}

// ValueNe filters for Value being not equal to argument
func (qs Dhcp6optionQS) ValueNe(v string) Dhcp6optionQS {
	return qs.filter(`"value" <>`, v)
}

// ValueLt filters for Value being less than argument
func (qs Dhcp6optionQS) ValueLt(v string) Dhcp6optionQS {
	return qs.filter(`"value" <`, v)
}

// ValueLe filters for Value being less than or equal to argument
func (qs Dhcp6optionQS) ValueLe(v string) Dhcp6optionQS {
	return qs.filter(`"value" <=`, v)
}

// ValueGt filters for Value being greater than argument
func (qs Dhcp6optionQS) ValueGt(v string) Dhcp6optionQS {
	return qs.filter(`"value" >`, v)
}

// ValueGe filters for Value being greater than or equal to argument
func (qs Dhcp6optionQS) ValueGe(v string) Dhcp6optionQS {
	return qs.filter(`"value" >=`, v)
}

type inDhcp6optionValue struct {
	values []interface{}
}

func (in *inDhcp6optionValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"value" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) ValueIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionValue{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionValue struct {
	values []interface{}
}

func (in *notinDhcp6optionValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"value" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) ValueNotIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionValue{
			values: vals,
		},
	)

	return qs
}

// OrderByValue sorts result by Value in ascending order
func (qs Dhcp6optionQS) OrderByValue() Dhcp6optionQS {
	qs.order = append(qs.order, `"value"`)

	return qs
}

// OrderByValueDesc sorts result by Value in descending order
func (qs Dhcp6optionQS) OrderByValueDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"value" DESC`)

	return qs
}

// FormattedValueIsNull filters for FormattedValue being null
func (qs Dhcp6optionQS) FormattedValueIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NULL`,
		},
	)
	return qs
}

// FormattedValueIsNotNull filters for FormattedValue being not null
func (qs Dhcp6optionQS) FormattedValueIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"formatted_value" IS NOT NULL`,
		},
	)
	return qs
}

// FormattedValueEq filters for FormattedValue being equal to argument
func (qs Dhcp6optionQS) FormattedValueEq(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" =`, v)
}

// FormattedValueNe filters for FormattedValue being not equal to argument
func (qs Dhcp6optionQS) FormattedValueNe(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" <>`, v)
}

// FormattedValueLt filters for FormattedValue being less than argument
func (qs Dhcp6optionQS) FormattedValueLt(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" <`, v)
}

// FormattedValueLe filters for FormattedValue being less than or equal to argument
func (qs Dhcp6optionQS) FormattedValueLe(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" <=`, v)
}

// FormattedValueGt filters for FormattedValue being greater than argument
func (qs Dhcp6optionQS) FormattedValueGt(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" >`, v)
}

// FormattedValueGe filters for FormattedValue being greater than or equal to argument
func (qs Dhcp6optionQS) FormattedValueGe(v string) Dhcp6optionQS {
	return qs.filter(`"formatted_value" >=`, v)
}

type inDhcp6optionFormattedValue struct {
	values []interface{}
}

func (in *inDhcp6optionFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"formatted_value" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) FormattedValueIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionFormattedValue{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionFormattedValue struct {
	values []interface{}
}

func (in *notinDhcp6optionFormattedValue) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"formatted_value" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) FormattedValueNotIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionFormattedValue{
			values: vals,
		},
	)

	return qs
}

// OrderByFormattedValue sorts result by FormattedValue in ascending order
func (qs Dhcp6optionQS) OrderByFormattedValue() Dhcp6optionQS {
	qs.order = append(qs.order, `"formatted_value"`)

	return qs
}

// OrderByFormattedValueDesc sorts result by FormattedValue in descending order
func (qs Dhcp6optionQS) OrderByFormattedValueDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"formatted_value" DESC`)

	return qs
}

// SpaceIsNull filters for Space being null
func (qs Dhcp6optionQS) SpaceIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NULL`,
		},
	)
	return qs
}

// SpaceIsNotNull filters for Space being not null
func (qs Dhcp6optionQS) SpaceIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"space" IS NOT NULL`,
		},
	)
	return qs
}

// SpaceEq filters for Space being equal to argument
func (qs Dhcp6optionQS) SpaceEq(v string) Dhcp6optionQS {
	return qs.filter(`"space" =`, v)
}

// SpaceNe filters for Space being not equal to argument
func (qs Dhcp6optionQS) SpaceNe(v string) Dhcp6optionQS {
	return qs.filter(`"space" <>`, v)
}

// SpaceLt filters for Space being less than argument
func (qs Dhcp6optionQS) SpaceLt(v string) Dhcp6optionQS {
	return qs.filter(`"space" <`, v)
}

// SpaceLe filters for Space being less than or equal to argument
func (qs Dhcp6optionQS) SpaceLe(v string) Dhcp6optionQS {
	return qs.filter(`"space" <=`, v)
}

// SpaceGt filters for Space being greater than argument
func (qs Dhcp6optionQS) SpaceGt(v string) Dhcp6optionQS {
	return qs.filter(`"space" >`, v)
}

// SpaceGe filters for Space being greater than or equal to argument
func (qs Dhcp6optionQS) SpaceGe(v string) Dhcp6optionQS {
	return qs.filter(`"space" >=`, v)
}

type inDhcp6optionSpace struct {
	values []interface{}
}

func (in *inDhcp6optionSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"space" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) SpaceIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionSpace{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionSpace struct {
	values []interface{}
}

func (in *notinDhcp6optionSpace) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"space" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) SpaceNotIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionSpace{
			values: vals,
		},
	)

	return qs
}

// OrderBySpace sorts result by Space in ascending order
func (qs Dhcp6optionQS) OrderBySpace() Dhcp6optionQS {
	qs.order = append(qs.order, `"space"`)

	return qs
}

// OrderBySpaceDesc sorts result by Space in descending order
func (qs Dhcp6optionQS) OrderBySpaceDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"space" DESC`)

	return qs
}

// PersistentEq filters for Persistent being equal to argument
func (qs Dhcp6optionQS) PersistentEq(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" =`, v)
}

// PersistentNe filters for Persistent being not equal to argument
func (qs Dhcp6optionQS) PersistentNe(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" <>`, v)
}

// PersistentLt filters for Persistent being less than argument
func (qs Dhcp6optionQS) PersistentLt(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" <`, v)
}

// PersistentLe filters for Persistent being less than or equal to argument
func (qs Dhcp6optionQS) PersistentLe(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" <=`, v)
}

// PersistentGt filters for Persistent being greater than argument
func (qs Dhcp6optionQS) PersistentGt(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" >`, v)
}

// PersistentGe filters for Persistent being greater than or equal to argument
func (qs Dhcp6optionQS) PersistentGe(v bool) Dhcp6optionQS {
	return qs.filter(`"persistent" >=`, v)
}

type inDhcp6optionPersistent struct {
	values []interface{}
}

func (in *inDhcp6optionPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"persistent" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) PersistentIn(values []bool) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionPersistent{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionPersistent struct {
	values []interface{}
}

func (in *notinDhcp6optionPersistent) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"persistent" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) PersistentNotIn(values []bool) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionPersistent{
			values: vals,
		},
	)

	return qs
}

// OrderByPersistent sorts result by Persistent in ascending order
func (qs Dhcp6optionQS) OrderByPersistent() Dhcp6optionQS {
	qs.order = append(qs.order, `"persistent"`)

	return qs
}

// OrderByPersistentDesc sorts result by Persistent in descending order
func (qs Dhcp6optionQS) OrderByPersistentDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"persistent" DESC`)

	return qs
}

// DhcpClientClassIsNull filters for DhcpClientClass being null
func (qs Dhcp6optionQS) DhcpClientClassIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NULL`,
		},
	)
	return qs
}

// DhcpClientClassIsNotNull filters for DhcpClientClass being not null
func (qs Dhcp6optionQS) DhcpClientClassIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp_client_class" IS NOT NULL`,
		},
	)
	return qs
}

// DhcpClientClassEq filters for DhcpClientClass being equal to argument
func (qs Dhcp6optionQS) DhcpClientClassEq(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" =`, v)
}

// DhcpClientClassNe filters for DhcpClientClass being not equal to argument
func (qs Dhcp6optionQS) DhcpClientClassNe(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" <>`, v)
}

// DhcpClientClassLt filters for DhcpClientClass being less than argument
func (qs Dhcp6optionQS) DhcpClientClassLt(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" <`, v)
}

// DhcpClientClassLe filters for DhcpClientClass being less than or equal to argument
func (qs Dhcp6optionQS) DhcpClientClassLe(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" <=`, v)
}

// DhcpClientClassGt filters for DhcpClientClass being greater than argument
func (qs Dhcp6optionQS) DhcpClientClassGt(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" >`, v)
}

// DhcpClientClassGe filters for DhcpClientClass being greater than or equal to argument
func (qs Dhcp6optionQS) DhcpClientClassGe(v string) Dhcp6optionQS {
	return qs.filter(`"dhcp_client_class" >=`, v)
}

type inDhcp6optionDhcpClientClass struct {
	values []interface{}
}

func (in *inDhcp6optionDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) DhcpClientClassIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionDhcpClientClass{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionDhcpClientClass struct {
	values []interface{}
}

func (in *notinDhcp6optionDhcpClientClass) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_client_class" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) DhcpClientClassNotIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionDhcpClientClass{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcpClientClass sorts result by DhcpClientClass in ascending order
func (qs Dhcp6optionQS) OrderByDhcpClientClass() Dhcp6optionQS {
	qs.order = append(qs.order, `"dhcp_client_class"`)

	return qs
}

// OrderByDhcpClientClassDesc sorts result by DhcpClientClass in descending order
func (qs Dhcp6optionQS) OrderByDhcpClientClassDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"dhcp_client_class" DESC`)

	return qs
}

// Dhcp6SubnetIdIsNull filters for Dhcp6SubnetId being null
func (qs Dhcp6optionQS) Dhcp6SubnetIdIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdIsNotNull filters for Dhcp6SubnetId being not null
func (qs Dhcp6optionQS) Dhcp6SubnetIdIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdEq filters for Dhcp6SubnetId being equal to argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdEq(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" =`, v)
}

// Dhcp6SubnetIdNe filters for Dhcp6SubnetId being not equal to argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdNe(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" <>`, v)
}

// Dhcp6SubnetIdLt filters for Dhcp6SubnetId being less than argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdLt(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" <`, v)
}

// Dhcp6SubnetIdLe filters for Dhcp6SubnetId being less than or equal to argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdLe(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" <=`, v)
}

// Dhcp6SubnetIdGt filters for Dhcp6SubnetId being greater than argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdGt(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" >`, v)
}

// Dhcp6SubnetIdGe filters for Dhcp6SubnetId being greater than or equal to argument
func (qs Dhcp6optionQS) Dhcp6SubnetIdGe(v int64) Dhcp6optionQS {
	return qs.filter(`"dhcp6_subnet_id" >=`, v)
}

type inDhcp6optionDhcp6SubnetId struct {
	values []interface{}
}

func (in *inDhcp6optionDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) Dhcp6SubnetIdIn(values []int64) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionDhcp6SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionDhcp6SubnetId struct {
	values []interface{}
}

func (in *notinDhcp6optionDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) Dhcp6SubnetIdNotIn(values []int64) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionDhcp6SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp6SubnetId sorts result by Dhcp6SubnetId in ascending order
func (qs Dhcp6optionQS) OrderByDhcp6SubnetId() Dhcp6optionQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id"`)

	return qs
}

// OrderByDhcp6SubnetIdDesc sorts result by Dhcp6SubnetId in descending order
func (qs Dhcp6optionQS) OrderByDhcp6SubnetIdDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id" DESC`)

	return qs
}

// GetHost returns Host
func (d *Dhcp6option) GetHost(db models.DBInterface) (*Host, error) {
	if !d.host.Valid {
		return nil, nil
	}

	return HostQS{}.HostIdEq(d.host.Int32).First(db)
}

// SetHost sets foreign key pointer to Host
func (d *Dhcp6option) SetHost(ptr *Host) error {
	if ptr != nil {
		d.host.Int32 = ptr.GetHostId()
		d.host.Valid = true
	} else {
		d.host.Valid = false
	}

	return nil
}

// GetHostRaw returns Dhcp6option.Host
func (d *Dhcp6option) GetHostRaw() sql.NullInt32 {
	return d.host
}

// HostIsNull filters for host being null
func (qs Dhcp6optionQS) HostIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NULL`,
		},
	)
	return qs
}

// HostIsNotNull filters for host being not null
func (qs Dhcp6optionQS) HostIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"host_id" IS NOT NULL`,
		},
	)
	return qs
}

// HostEq filters for host being equal to argument
func (qs Dhcp6optionQS) HostEq(v *Host) Dhcp6optionQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Dhcp6optionQS) HostRawEq(v int32) Dhcp6optionQS {
	return qs.filter(`"host_id" =`, v)
}

type inDhcp6optionhostHost struct {
	qs HostQS
}

func (in *inDhcp6optionhostHost) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Dhcp6optionQS) HostIn(oqs HostQS) Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionhostHost{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Dhcp6optionQS) OrderByHost() Dhcp6optionQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Dhcp6optionQS) OrderByHostDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// GetScope returns Dhcpoptionscope
func (d *Dhcp6option) GetScope(db models.DBInterface) (*Dhcpoptionscope, error) {
	return DhcpoptionscopeQS{}.ScopeIdEq(d.scope).First(db)
}

// SetScope sets foreign key pointer to Dhcpoptionscope
func (d *Dhcp6option) SetScope(ptr *Dhcpoptionscope) error {
	if ptr != nil {
		d.scope = ptr.ScopeId
	} else {
		return fmt.Errorf("Dhcp6option.SetScope: non-null field received null value")
	}

	return nil
}

// GetScopeRaw returns Dhcp6option.Scope
func (d *Dhcp6option) GetScopeRaw() int32 {
	return d.scope
}

// ScopeEq filters for scope being equal to argument
func (qs Dhcp6optionQS) ScopeEq(v *Dhcpoptionscope) Dhcp6optionQS {
	return qs.filter(`"scope_id" =`, v.ScopeId)
}

// ScopeRawEq filters for scope being equal to raw argument
func (qs Dhcp6optionQS) ScopeRawEq(v int32) Dhcp6optionQS {
	return qs.filter(`"scope_id" =`, v)
}

type inDhcp6optionscopeDhcpoptionscope struct {
	qs DhcpoptionscopeQS
}

func (in *inDhcp6optionscopeDhcpoptionscope) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"scope_id" IN (` + s + `)`, p
}

func (qs Dhcp6optionQS) ScopeIn(oqs DhcpoptionscopeQS) Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionscopeDhcpoptionscope{
			qs: oqs,
		},
	)

	return qs
}

// OrderByScope sorts result by Scope in ascending order
func (qs Dhcp6optionQS) OrderByScope() Dhcp6optionQS {
	qs.order = append(qs.order, `"scope_id"`)

	return qs
}

// OrderByScopeDesc sorts result by Scope in descending order
func (qs Dhcp6optionQS) OrderByScopeDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"scope_id" DESC`)

	return qs
}

// UserContextIsNull filters for UserContext being null
func (qs Dhcp6optionQS) UserContextIsNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Dhcp6optionQS) UserContextIsNotNull() Dhcp6optionQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Dhcp6optionQS) UserContextEq(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Dhcp6optionQS) UserContextNe(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Dhcp6optionQS) UserContextLt(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Dhcp6optionQS) UserContextLe(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Dhcp6optionQS) UserContextGt(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Dhcp6optionQS) UserContextGe(v string) Dhcp6optionQS {
	return qs.filter(`"user_context" >=`, v)
}

type inDhcp6optionUserContext struct {
	values []interface{}
}

func (in *inDhcp6optionUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) UserContextIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDhcp6optionUserContext{
			values: vals,
		},
	)

	return qs
}

type notinDhcp6optionUserContext struct {
	values []interface{}
}

func (in *notinDhcp6optionUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Dhcp6optionQS) UserContextNotIn(values []string) Dhcp6optionQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDhcp6optionUserContext{
			values: vals,
		},
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Dhcp6optionQS) OrderByUserContext() Dhcp6optionQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Dhcp6optionQS) OrderByUserContextDesc() Dhcp6optionQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Dhcp6optionQS) ForUpdate() Dhcp6optionQS {
	qs.forUpdate = true

	return qs
}

func (qs Dhcp6optionQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Dhcp6optionQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Dhcp6optionQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "option_id", "code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp6_subnet_id", "host_id", "scope_id", "user_context" FROM "dhcp6_options"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Dhcp6optionQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "option_id" FROM "dhcp6_options"` + s, p
}

// All returns all rows matching queryset filters
func (qs Dhcp6optionQS) All(db models.DBInterface) ([]*Dhcp6option, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Dhcp6option
	for rows.Next() {
		obj := Dhcp6option{existsInDB: true}
		if err = rows.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp6SubnetId, &obj.host, &obj.scope, &obj.UserContext); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Dhcp6optionQS) First(db models.DBInterface) (*Dhcp6option, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Dhcp6option{existsInDB: true}
	err := row.Scan(&obj.optionId, &obj.Code, &obj.Value, &obj.FormattedValue, &obj.Space, &obj.Persistent, &obj.DhcpClientClass, &obj.Dhcp6SubnetId, &obj.host, &obj.scope, &obj.UserContext)
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
func (qs Dhcp6optionQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "dhcp6_options"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Dhcp6optionQS) Update() Dhcp6optionUpdateQS {
	return Dhcp6optionUpdateQS{condFragments: qs.condFragments}
}

// Dhcp6optionUpdateQS represents an updated queryset for django_kea.Dhcp6Option
type Dhcp6optionUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Dhcp6optionUpdateQS) update(c string, v interface{}) Dhcp6optionUpdateQS {
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
func (uqs Dhcp6optionUpdateQS) SetOptionId(v int32) Dhcp6optionUpdateQS {
	return uqs.update(`"option_id"`, v)
}

// SetCode sets Code to the given value
func (uqs Dhcp6optionUpdateQS) SetCode(v int32) Dhcp6optionUpdateQS {
	return uqs.update(`"code"`, v)
}

// SetValue sets Value to the given value
func (uqs Dhcp6optionUpdateQS) SetValue(v sql.NullString) Dhcp6optionUpdateQS {
	return uqs.update(`"value"`, v)
}

// SetFormattedValue sets FormattedValue to the given value
func (uqs Dhcp6optionUpdateQS) SetFormattedValue(v sql.NullString) Dhcp6optionUpdateQS {
	return uqs.update(`"formatted_value"`, v)
}

// SetSpace sets Space to the given value
func (uqs Dhcp6optionUpdateQS) SetSpace(v sql.NullString) Dhcp6optionUpdateQS {
	return uqs.update(`"space"`, v)
}

// SetPersistent sets Persistent to the given value
func (uqs Dhcp6optionUpdateQS) SetPersistent(v bool) Dhcp6optionUpdateQS {
	return uqs.update(`"persistent"`, v)
}

// SetDhcpClientClass sets DhcpClientClass to the given value
func (uqs Dhcp6optionUpdateQS) SetDhcpClientClass(v sql.NullString) Dhcp6optionUpdateQS {
	return uqs.update(`"dhcp_client_class"`, v)
}

// SetDhcp6SubnetId sets Dhcp6SubnetId to the given value
func (uqs Dhcp6optionUpdateQS) SetDhcp6SubnetId(v sql.NullInt64) Dhcp6optionUpdateQS {
	return uqs.update(`"dhcp6_subnet_id"`, v)
}

// SetHost sets foreign key pointer to Host
func (uqs Dhcp6optionUpdateQS) SetHost(ptr *Host) Dhcp6optionUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // SetScope sets foreign key pointer to Dhcpoptionscope
func (uqs Dhcp6optionUpdateQS) SetScope(ptr *Dhcpoptionscope) Dhcp6optionUpdateQS {
	if ptr != nil {
		return uqs.update(`"scope_id"`, ptr.ScopeId)
	}

	return uqs.update(`"scope_id"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Dhcp6optionUpdateQS) SetUserContext(v sql.NullString) Dhcp6optionUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// Exec executes the update operation
func (uqs Dhcp6optionUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Dhcp6optionQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "dhcp6_options" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (d *Dhcp6option) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "dhcp6_options" ("code", "value", "formatted_value", "space", "persistent", "dhcp_client_class", "dhcp6_subnet_id", "host_id", "scope_id", "user_context") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "option_id"`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp6SubnetId, d.host, d.scope, d.UserContext)

	if err := row.Scan(&d.optionId); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Dhcp6option) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "dhcp6_options" SET "code" = $1, "value" = $2, "formatted_value" = $3, "space" = $4, "persistent" = $5, "dhcp_client_class" = $6, "dhcp6_subnet_id" = $7, "host_id" = $8, "scope_id" = $9, "user_context" = $10 WHERE "option_id" = $11`, d.Code, d.Value, d.FormattedValue, d.Space, d.Persistent, d.DhcpClientClass, d.Dhcp6SubnetId, d.host, d.scope, d.UserContext, d.optionId)

	return err
}

// Save inserts or updates record
func (d *Dhcp6option) Save(db models.DBInterface) error {
	if d.existsInDB {
		return d.update(db)
	}

	return d.insert(db)
}

// Delete removes row from database
func (d *Dhcp6option) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "dhcp6_options" WHERE "option_id" = $1`, d.optionId)

	d.existsInDB = false

	return err
}
