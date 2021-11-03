/*
  AUTO-GENERATED file for Django model django_kea.Lease4

  Command used to generate:

  DJANGO_SETTINGS_MODULE=keatest.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/euronetzrt/django-kea django_kea

  https://github.com/rkojedzinszky/djan-go-rm
*/

package django_kea

import (
	"database/sql"
	"github.com/euronetzrt/django-kea/models"
	"strings"
	"time"
)

// Lease4 mirrors model django_kea.Lease4
type Lease4 struct {
	existsInDB bool

	Address       string
	Hwaddr        sql.NullString
	ClientId      sql.NullString
	ValidLifetime sql.NullInt64
	Expire        sql.NullTime
	SubnetId      sql.NullInt64
	FqdnFwd       sql.NullBool
	FqdnRev       sql.NullBool
	Hostname      sql.NullString
	state         sql.NullInt64
	UserContext   sql.NullString
}

// Lease4QS represents a queryset for django_kea.Lease4
type Lease4QS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Lease4QS) filter(c string, p interface{}) Lease4QS {
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
func (qs Lease4QS) Or(exprs ...Lease4QS) Lease4QS {
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

// AddressEq filters for Address being equal to argument
func (qs Lease4QS) AddressEq(v string) Lease4QS {
	return qs.filter(`"address" =`, v)
}

// AddressNe filters for Address being not equal to argument
func (qs Lease4QS) AddressNe(v string) Lease4QS {
	return qs.filter(`"address" <>`, v)
}

// AddressLt filters for Address being less than argument
func (qs Lease4QS) AddressLt(v string) Lease4QS {
	return qs.filter(`"address" <`, v)
}

// AddressLe filters for Address being less than or equal to argument
func (qs Lease4QS) AddressLe(v string) Lease4QS {
	return qs.filter(`"address" <=`, v)
}

// AddressGt filters for Address being greater than argument
func (qs Lease4QS) AddressGt(v string) Lease4QS {
	return qs.filter(`"address" >`, v)
}

// AddressGe filters for Address being greater than or equal to argument
func (qs Lease4QS) AddressGe(v string) Lease4QS {
	return qs.filter(`"address" >=`, v)
}

type inLease4Address struct {
	values []interface{}
}

func (in *inLease4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) AddressIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4Address{
			values: vals,
		},
	)

	return qs
}

type notinLease4Address struct {
	values []interface{}
}

func (in *notinLease4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) AddressNotIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4Address{
			values: vals,
		},
	)

	return qs
}

// OrderByAddress sorts result by Address in ascending order
func (qs Lease4QS) OrderByAddress() Lease4QS {
	qs.order = append(qs.order, `"address"`)

	return qs
}

// OrderByAddressDesc sorts result by Address in descending order
func (qs Lease4QS) OrderByAddressDesc() Lease4QS {
	qs.order = append(qs.order, `"address" DESC`)

	return qs
}

// HwaddrIsNull filters for Hwaddr being null
func (qs Lease4QS) HwaddrIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr" IS NULL`,
		},
	)
	return qs
}

// HwaddrIsNotNull filters for Hwaddr being not null
func (qs Lease4QS) HwaddrIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr" IS NOT NULL`,
		},
	)
	return qs
}

// HwaddrEq filters for Hwaddr being equal to argument
func (qs Lease4QS) HwaddrEq(v string) Lease4QS {
	return qs.filter(`"hwaddr" =`, v)
}

// HwaddrNe filters for Hwaddr being not equal to argument
func (qs Lease4QS) HwaddrNe(v string) Lease4QS {
	return qs.filter(`"hwaddr" <>`, v)
}

// HwaddrLt filters for Hwaddr being less than argument
func (qs Lease4QS) HwaddrLt(v string) Lease4QS {
	return qs.filter(`"hwaddr" <`, v)
}

// HwaddrLe filters for Hwaddr being less than or equal to argument
func (qs Lease4QS) HwaddrLe(v string) Lease4QS {
	return qs.filter(`"hwaddr" <=`, v)
}

// HwaddrGt filters for Hwaddr being greater than argument
func (qs Lease4QS) HwaddrGt(v string) Lease4QS {
	return qs.filter(`"hwaddr" >`, v)
}

// HwaddrGe filters for Hwaddr being greater than or equal to argument
func (qs Lease4QS) HwaddrGe(v string) Lease4QS {
	return qs.filter(`"hwaddr" >=`, v)
}

type inLease4Hwaddr struct {
	values []interface{}
}

func (in *inLease4Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) HwaddrIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4Hwaddr{
			values: vals,
		},
	)

	return qs
}

type notinLease4Hwaddr struct {
	values []interface{}
}

func (in *notinLease4Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) HwaddrNotIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4Hwaddr{
			values: vals,
		},
	)

	return qs
}

// OrderByHwaddr sorts result by Hwaddr in ascending order
func (qs Lease4QS) OrderByHwaddr() Lease4QS {
	qs.order = append(qs.order, `"hwaddr"`)

	return qs
}

// OrderByHwaddrDesc sorts result by Hwaddr in descending order
func (qs Lease4QS) OrderByHwaddrDesc() Lease4QS {
	qs.order = append(qs.order, `"hwaddr" DESC`)

	return qs
}

// ClientIdIsNull filters for ClientId being null
func (qs Lease4QS) ClientIdIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"client_id" IS NULL`,
		},
	)
	return qs
}

// ClientIdIsNotNull filters for ClientId being not null
func (qs Lease4QS) ClientIdIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"client_id" IS NOT NULL`,
		},
	)
	return qs
}

// ClientIdEq filters for ClientId being equal to argument
func (qs Lease4QS) ClientIdEq(v string) Lease4QS {
	return qs.filter(`"client_id" =`, v)
}

// ClientIdNe filters for ClientId being not equal to argument
func (qs Lease4QS) ClientIdNe(v string) Lease4QS {
	return qs.filter(`"client_id" <>`, v)
}

// ClientIdLt filters for ClientId being less than argument
func (qs Lease4QS) ClientIdLt(v string) Lease4QS {
	return qs.filter(`"client_id" <`, v)
}

// ClientIdLe filters for ClientId being less than or equal to argument
func (qs Lease4QS) ClientIdLe(v string) Lease4QS {
	return qs.filter(`"client_id" <=`, v)
}

// ClientIdGt filters for ClientId being greater than argument
func (qs Lease4QS) ClientIdGt(v string) Lease4QS {
	return qs.filter(`"client_id" >`, v)
}

// ClientIdGe filters for ClientId being greater than or equal to argument
func (qs Lease4QS) ClientIdGe(v string) Lease4QS {
	return qs.filter(`"client_id" >=`, v)
}

type inLease4ClientId struct {
	values []interface{}
}

func (in *inLease4ClientId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"client_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ClientIdIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4ClientId{
			values: vals,
		},
	)

	return qs
}

type notinLease4ClientId struct {
	values []interface{}
}

func (in *notinLease4ClientId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"client_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ClientIdNotIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4ClientId{
			values: vals,
		},
	)

	return qs
}

// OrderByClientId sorts result by ClientId in ascending order
func (qs Lease4QS) OrderByClientId() Lease4QS {
	qs.order = append(qs.order, `"client_id"`)

	return qs
}

// OrderByClientIdDesc sorts result by ClientId in descending order
func (qs Lease4QS) OrderByClientIdDesc() Lease4QS {
	qs.order = append(qs.order, `"client_id" DESC`)

	return qs
}

// ValidLifetimeIsNull filters for ValidLifetime being null
func (qs Lease4QS) ValidLifetimeIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"valid_lifetime" IS NULL`,
		},
	)
	return qs
}

// ValidLifetimeIsNotNull filters for ValidLifetime being not null
func (qs Lease4QS) ValidLifetimeIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"valid_lifetime" IS NOT NULL`,
		},
	)
	return qs
}

// ValidLifetimeEq filters for ValidLifetime being equal to argument
func (qs Lease4QS) ValidLifetimeEq(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" =`, v)
}

// ValidLifetimeNe filters for ValidLifetime being not equal to argument
func (qs Lease4QS) ValidLifetimeNe(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" <>`, v)
}

// ValidLifetimeLt filters for ValidLifetime being less than argument
func (qs Lease4QS) ValidLifetimeLt(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" <`, v)
}

// ValidLifetimeLe filters for ValidLifetime being less than or equal to argument
func (qs Lease4QS) ValidLifetimeLe(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" <=`, v)
}

// ValidLifetimeGt filters for ValidLifetime being greater than argument
func (qs Lease4QS) ValidLifetimeGt(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" >`, v)
}

// ValidLifetimeGe filters for ValidLifetime being greater than or equal to argument
func (qs Lease4QS) ValidLifetimeGe(v int64) Lease4QS {
	return qs.filter(`"valid_lifetime" >=`, v)
}

type inLease4ValidLifetime struct {
	values []interface{}
}

func (in *inLease4ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ValidLifetimeIn(values []int64) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4ValidLifetime{
			values: vals,
		},
	)

	return qs
}

type notinLease4ValidLifetime struct {
	values []interface{}
}

func (in *notinLease4ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ValidLifetimeNotIn(values []int64) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4ValidLifetime{
			values: vals,
		},
	)

	return qs
}

// OrderByValidLifetime sorts result by ValidLifetime in ascending order
func (qs Lease4QS) OrderByValidLifetime() Lease4QS {
	qs.order = append(qs.order, `"valid_lifetime"`)

	return qs
}

// OrderByValidLifetimeDesc sorts result by ValidLifetime in descending order
func (qs Lease4QS) OrderByValidLifetimeDesc() Lease4QS {
	qs.order = append(qs.order, `"valid_lifetime" DESC`)

	return qs
}

// ExpireIsNull filters for Expire being null
func (qs Lease4QS) ExpireIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"expire" IS NULL`,
		},
	)
	return qs
}

// ExpireIsNotNull filters for Expire being not null
func (qs Lease4QS) ExpireIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"expire" IS NOT NULL`,
		},
	)
	return qs
}

// ExpireEq filters for Expire being equal to argument
func (qs Lease4QS) ExpireEq(v time.Time) Lease4QS {
	return qs.filter(`"expire" =`, v)
}

// ExpireNe filters for Expire being not equal to argument
func (qs Lease4QS) ExpireNe(v time.Time) Lease4QS {
	return qs.filter(`"expire" <>`, v)
}

// ExpireLt filters for Expire being less than argument
func (qs Lease4QS) ExpireLt(v time.Time) Lease4QS {
	return qs.filter(`"expire" <`, v)
}

// ExpireLe filters for Expire being less than or equal to argument
func (qs Lease4QS) ExpireLe(v time.Time) Lease4QS {
	return qs.filter(`"expire" <=`, v)
}

// ExpireGt filters for Expire being greater than argument
func (qs Lease4QS) ExpireGt(v time.Time) Lease4QS {
	return qs.filter(`"expire" >`, v)
}

// ExpireGe filters for Expire being greater than or equal to argument
func (qs Lease4QS) ExpireGe(v time.Time) Lease4QS {
	return qs.filter(`"expire" >=`, v)
}

type inLease4Expire struct {
	values []interface{}
}

func (in *inLease4Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"expire" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ExpireIn(values []time.Time) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4Expire{
			values: vals,
		},
	)

	return qs
}

type notinLease4Expire struct {
	values []interface{}
}

func (in *notinLease4Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"expire" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) ExpireNotIn(values []time.Time) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4Expire{
			values: vals,
		},
	)

	return qs
}

// OrderByExpire sorts result by Expire in ascending order
func (qs Lease4QS) OrderByExpire() Lease4QS {
	qs.order = append(qs.order, `"expire"`)

	return qs
}

// OrderByExpireDesc sorts result by Expire in descending order
func (qs Lease4QS) OrderByExpireDesc() Lease4QS {
	qs.order = append(qs.order, `"expire" DESC`)

	return qs
}

// SubnetIdIsNull filters for SubnetId being null
func (qs Lease4QS) SubnetIdIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"subnet_id" IS NULL`,
		},
	)
	return qs
}

// SubnetIdIsNotNull filters for SubnetId being not null
func (qs Lease4QS) SubnetIdIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// SubnetIdEq filters for SubnetId being equal to argument
func (qs Lease4QS) SubnetIdEq(v int64) Lease4QS {
	return qs.filter(`"subnet_id" =`, v)
}

// SubnetIdNe filters for SubnetId being not equal to argument
func (qs Lease4QS) SubnetIdNe(v int64) Lease4QS {
	return qs.filter(`"subnet_id" <>`, v)
}

// SubnetIdLt filters for SubnetId being less than argument
func (qs Lease4QS) SubnetIdLt(v int64) Lease4QS {
	return qs.filter(`"subnet_id" <`, v)
}

// SubnetIdLe filters for SubnetId being less than or equal to argument
func (qs Lease4QS) SubnetIdLe(v int64) Lease4QS {
	return qs.filter(`"subnet_id" <=`, v)
}

// SubnetIdGt filters for SubnetId being greater than argument
func (qs Lease4QS) SubnetIdGt(v int64) Lease4QS {
	return qs.filter(`"subnet_id" >`, v)
}

// SubnetIdGe filters for SubnetId being greater than or equal to argument
func (qs Lease4QS) SubnetIdGe(v int64) Lease4QS {
	return qs.filter(`"subnet_id" >=`, v)
}

type inLease4SubnetId struct {
	values []interface{}
}

func (in *inLease4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) SubnetIdIn(values []int64) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinLease4SubnetId struct {
	values []interface{}
}

func (in *notinLease4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) SubnetIdNotIn(values []int64) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderBySubnetId sorts result by SubnetId in ascending order
func (qs Lease4QS) OrderBySubnetId() Lease4QS {
	qs.order = append(qs.order, `"subnet_id"`)

	return qs
}

// OrderBySubnetIdDesc sorts result by SubnetId in descending order
func (qs Lease4QS) OrderBySubnetIdDesc() Lease4QS {
	qs.order = append(qs.order, `"subnet_id" DESC`)

	return qs
}

// FqdnFwdIsNull filters for FqdnFwd being null
func (qs Lease4QS) FqdnFwdIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_fwd" IS NULL`,
		},
	)
	return qs
}

// FqdnFwdIsNotNull filters for FqdnFwd being not null
func (qs Lease4QS) FqdnFwdIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_fwd" IS NOT NULL`,
		},
	)
	return qs
}

// FqdnFwdEq filters for FqdnFwd being equal to argument
func (qs Lease4QS) FqdnFwdEq(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" =`, v)
}

// FqdnFwdNe filters for FqdnFwd being not equal to argument
func (qs Lease4QS) FqdnFwdNe(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" <>`, v)
}

// FqdnFwdLt filters for FqdnFwd being less than argument
func (qs Lease4QS) FqdnFwdLt(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" <`, v)
}

// FqdnFwdLe filters for FqdnFwd being less than or equal to argument
func (qs Lease4QS) FqdnFwdLe(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" <=`, v)
}

// FqdnFwdGt filters for FqdnFwd being greater than argument
func (qs Lease4QS) FqdnFwdGt(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" >`, v)
}

// FqdnFwdGe filters for FqdnFwd being greater than or equal to argument
func (qs Lease4QS) FqdnFwdGe(v bool) Lease4QS {
	return qs.filter(`"fqdn_fwd" >=`, v)
}

type inLease4FqdnFwd struct {
	values []interface{}
}

func (in *inLease4FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) FqdnFwdIn(values []bool) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4FqdnFwd{
			values: vals,
		},
	)

	return qs
}

type notinLease4FqdnFwd struct {
	values []interface{}
}

func (in *notinLease4FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) FqdnFwdNotIn(values []bool) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4FqdnFwd{
			values: vals,
		},
	)

	return qs
}

// OrderByFqdnFwd sorts result by FqdnFwd in ascending order
func (qs Lease4QS) OrderByFqdnFwd() Lease4QS {
	qs.order = append(qs.order, `"fqdn_fwd"`)

	return qs
}

// OrderByFqdnFwdDesc sorts result by FqdnFwd in descending order
func (qs Lease4QS) OrderByFqdnFwdDesc() Lease4QS {
	qs.order = append(qs.order, `"fqdn_fwd" DESC`)

	return qs
}

// FqdnRevIsNull filters for FqdnRev being null
func (qs Lease4QS) FqdnRevIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_rev" IS NULL`,
		},
	)
	return qs
}

// FqdnRevIsNotNull filters for FqdnRev being not null
func (qs Lease4QS) FqdnRevIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_rev" IS NOT NULL`,
		},
	)
	return qs
}

// FqdnRevEq filters for FqdnRev being equal to argument
func (qs Lease4QS) FqdnRevEq(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" =`, v)
}

// FqdnRevNe filters for FqdnRev being not equal to argument
func (qs Lease4QS) FqdnRevNe(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" <>`, v)
}

// FqdnRevLt filters for FqdnRev being less than argument
func (qs Lease4QS) FqdnRevLt(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" <`, v)
}

// FqdnRevLe filters for FqdnRev being less than or equal to argument
func (qs Lease4QS) FqdnRevLe(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" <=`, v)
}

// FqdnRevGt filters for FqdnRev being greater than argument
func (qs Lease4QS) FqdnRevGt(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" >`, v)
}

// FqdnRevGe filters for FqdnRev being greater than or equal to argument
func (qs Lease4QS) FqdnRevGe(v bool) Lease4QS {
	return qs.filter(`"fqdn_rev" >=`, v)
}

type inLease4FqdnRev struct {
	values []interface{}
}

func (in *inLease4FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) FqdnRevIn(values []bool) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4FqdnRev{
			values: vals,
		},
	)

	return qs
}

type notinLease4FqdnRev struct {
	values []interface{}
}

func (in *notinLease4FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) FqdnRevNotIn(values []bool) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4FqdnRev{
			values: vals,
		},
	)

	return qs
}

// OrderByFqdnRev sorts result by FqdnRev in ascending order
func (qs Lease4QS) OrderByFqdnRev() Lease4QS {
	qs.order = append(qs.order, `"fqdn_rev"`)

	return qs
}

// OrderByFqdnRevDesc sorts result by FqdnRev in descending order
func (qs Lease4QS) OrderByFqdnRevDesc() Lease4QS {
	qs.order = append(qs.order, `"fqdn_rev" DESC`)

	return qs
}

// HostnameIsNull filters for Hostname being null
func (qs Lease4QS) HostnameIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NULL`,
		},
	)
	return qs
}

// HostnameIsNotNull filters for Hostname being not null
func (qs Lease4QS) HostnameIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NOT NULL`,
		},
	)
	return qs
}

// HostnameEq filters for Hostname being equal to argument
func (qs Lease4QS) HostnameEq(v string) Lease4QS {
	return qs.filter(`"hostname" =`, v)
}

// HostnameNe filters for Hostname being not equal to argument
func (qs Lease4QS) HostnameNe(v string) Lease4QS {
	return qs.filter(`"hostname" <>`, v)
}

// HostnameLt filters for Hostname being less than argument
func (qs Lease4QS) HostnameLt(v string) Lease4QS {
	return qs.filter(`"hostname" <`, v)
}

// HostnameLe filters for Hostname being less than or equal to argument
func (qs Lease4QS) HostnameLe(v string) Lease4QS {
	return qs.filter(`"hostname" <=`, v)
}

// HostnameGt filters for Hostname being greater than argument
func (qs Lease4QS) HostnameGt(v string) Lease4QS {
	return qs.filter(`"hostname" >`, v)
}

// HostnameGe filters for Hostname being greater than or equal to argument
func (qs Lease4QS) HostnameGe(v string) Lease4QS {
	return qs.filter(`"hostname" >=`, v)
}

type inLease4Hostname struct {
	values []interface{}
}

func (in *inLease4Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) HostnameIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4Hostname{
			values: vals,
		},
	)

	return qs
}

type notinLease4Hostname struct {
	values []interface{}
}

func (in *notinLease4Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) HostnameNotIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4Hostname{
			values: vals,
		},
	)

	return qs
}

// OrderByHostname sorts result by Hostname in ascending order
func (qs Lease4QS) OrderByHostname() Lease4QS {
	qs.order = append(qs.order, `"hostname"`)

	return qs
}

// OrderByHostnameDesc sorts result by Hostname in descending order
func (qs Lease4QS) OrderByHostnameDesc() Lease4QS {
	qs.order = append(qs.order, `"hostname" DESC`)

	return qs
}

// GetState returns Leasestate
func (l *Lease4) GetState(db models.DBInterface) (*Leasestate, error) {
	if !l.state.Valid {
		return nil, nil
	}

	return LeasestateQS{}.StateEq(l.state.Int64).First(db)
}

// SetState sets foreign key pointer to Leasestate
func (l *Lease4) SetState(ptr *Leasestate) error {
	if ptr != nil {
		l.state.Int64 = ptr.State
		l.state.Valid = true
	} else {
		l.state.Valid = false
	}

	return nil
}

// GetStateRaw returns Lease4.State
func (l *Lease4) GetStateRaw() sql.NullInt64 {
	return l.state
}

// StateIsNull filters for state being null
func (qs Lease4QS) StateIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"state" IS NULL`,
		},
	)
	return qs
}

// StateIsNotNull filters for state being not null
func (qs Lease4QS) StateIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"state" IS NOT NULL`,
		},
	)
	return qs
}

// StateEq filters for state being equal to argument
func (qs Lease4QS) StateEq(v *Leasestate) Lease4QS {
	return qs.filter(`"state" =`, v.State)
}

// StateRawEq filters for state being equal to raw argument
func (qs Lease4QS) StateRawEq(v int64) Lease4QS {
	return qs.filter(`"state" =`, v)
}

type inLease4stateLeasestate struct {
	qs LeasestateQS
}

func (in *inLease4stateLeasestate) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"state" IN (` + s + `)`, p
}

func (qs Lease4QS) StateIn(oqs LeasestateQS) Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&inLease4stateLeasestate{
			qs: oqs,
		},
	)

	return qs
}

// OrderByState sorts result by State in ascending order
func (qs Lease4QS) OrderByState() Lease4QS {
	qs.order = append(qs.order, `"state"`)

	return qs
}

// OrderByStateDesc sorts result by State in descending order
func (qs Lease4QS) OrderByStateDesc() Lease4QS {
	qs.order = append(qs.order, `"state" DESC`)

	return qs
}

// UserContextIsNull filters for UserContext being null
func (qs Lease4QS) UserContextIsNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Lease4QS) UserContextIsNotNull() Lease4QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Lease4QS) UserContextEq(v string) Lease4QS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Lease4QS) UserContextNe(v string) Lease4QS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Lease4QS) UserContextLt(v string) Lease4QS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Lease4QS) UserContextLe(v string) Lease4QS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Lease4QS) UserContextGt(v string) Lease4QS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Lease4QS) UserContextGe(v string) Lease4QS {
	return qs.filter(`"user_context" >=`, v)
}

type inLease4UserContext struct {
	values []interface{}
}

func (in *inLease4UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) UserContextIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease4UserContext{
			values: vals,
		},
	)

	return qs
}

type notinLease4UserContext struct {
	values []interface{}
}

func (in *notinLease4UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease4QS) UserContextNotIn(values []string) Lease4QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease4UserContext{
			values: vals,
		},
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Lease4QS) OrderByUserContext() Lease4QS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Lease4QS) OrderByUserContextDesc() Lease4QS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Lease4QS) ForUpdate() Lease4QS {
	qs.forUpdate = true

	return qs
}

func (qs Lease4QS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Lease4QS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Lease4QS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "address", "hwaddr", "client_id", "valid_lifetime", "expire", "subnet_id", "fqdn_fwd", "fqdn_rev", "hostname", "state", "user_context" FROM "lease4"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Lease4QS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "address" FROM "lease4"` + s, p
}

// All returns all rows matching queryset filters
func (qs Lease4QS) All(db models.DBInterface) ([]*Lease4, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Lease4
	for rows.Next() {
		obj := Lease4{existsInDB: true}
		if err = rows.Scan(&obj.Address, &obj.Hwaddr, &obj.ClientId, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.UserContext); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Lease4QS) First(db models.DBInterface) (*Lease4, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Lease4{existsInDB: true}
	err := row.Scan(&obj.Address, &obj.Hwaddr, &obj.ClientId, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.UserContext)
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
func (qs Lease4QS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease4"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Lease4QS) Update() Lease4UpdateQS {
	return Lease4UpdateQS{condFragments: qs.condFragments}
}

// Lease4UpdateQS represents an updated queryset for django_kea.Lease4
type Lease4UpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Lease4UpdateQS) update(c string, v interface{}) Lease4UpdateQS {
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

// SetAddress sets Address to the given value
func (uqs Lease4UpdateQS) SetAddress(v string) Lease4UpdateQS {
	return uqs.update(`"address"`, v)
}

// SetHwaddr sets Hwaddr to the given value
func (uqs Lease4UpdateQS) SetHwaddr(v sql.NullString) Lease4UpdateQS {
	return uqs.update(`"hwaddr"`, v)
}

// SetClientId sets ClientId to the given value
func (uqs Lease4UpdateQS) SetClientId(v sql.NullString) Lease4UpdateQS {
	return uqs.update(`"client_id"`, v)
}

// SetValidLifetime sets ValidLifetime to the given value
func (uqs Lease4UpdateQS) SetValidLifetime(v sql.NullInt64) Lease4UpdateQS {
	return uqs.update(`"valid_lifetime"`, v)
}

// SetExpire sets Expire to the given value
func (uqs Lease4UpdateQS) SetExpire(v sql.NullTime) Lease4UpdateQS {
	return uqs.update(`"expire"`, v)
}

// SetSubnetId sets SubnetId to the given value
func (uqs Lease4UpdateQS) SetSubnetId(v sql.NullInt64) Lease4UpdateQS {
	return uqs.update(`"subnet_id"`, v)
}

// SetFqdnFwd sets FqdnFwd to the given value
func (uqs Lease4UpdateQS) SetFqdnFwd(v sql.NullBool) Lease4UpdateQS {
	return uqs.update(`"fqdn_fwd"`, v)
}

// SetFqdnRev sets FqdnRev to the given value
func (uqs Lease4UpdateQS) SetFqdnRev(v sql.NullBool) Lease4UpdateQS {
	return uqs.update(`"fqdn_rev"`, v)
}

// SetHostname sets Hostname to the given value
func (uqs Lease4UpdateQS) SetHostname(v sql.NullString) Lease4UpdateQS {
	return uqs.update(`"hostname"`, v)
}

// SetState sets foreign key pointer to Leasestate
func (uqs Lease4UpdateQS) SetState(ptr *Leasestate) Lease4UpdateQS {
	if ptr != nil {
		return uqs.update(`"state"`, ptr.State)
	}

	return uqs.update(`"state"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Lease4UpdateQS) SetUserContext(v sql.NullString) Lease4UpdateQS {
	return uqs.update(`"user_context"`, v)
}

// Exec executes the update operation
func (uqs Lease4UpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Lease4QS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease4" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (l *Lease4) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "lease4" ("hwaddr", "client_id", "valid_lifetime", "expire", "subnet_id", "fqdn_fwd", "fqdn_rev", "hostname", "state", "user_context", "address") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, l.Hwaddr, l.ClientId, l.ValidLifetime, l.Expire, l.SubnetId, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.UserContext, l.Address)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Lease4) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "lease4" SET "hwaddr" = $1, "client_id" = $2, "valid_lifetime" = $3, "expire" = $4, "subnet_id" = $5, "fqdn_fwd" = $6, "fqdn_rev" = $7, "hostname" = $8, "state" = $9, "user_context" = $10 WHERE "address" = $11`, l.Hwaddr, l.ClientId, l.ValidLifetime, l.Expire, l.SubnetId, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.UserContext, l.Address)

	return err
}

// Save inserts or updates record
func (l *Lease4) Save(db models.DBInterface) error {
	if l.existsInDB {
		return l.update(db)
	}

	return l.insert(db)
}

// Delete removes row from database
func (l *Lease4) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "lease4" WHERE "address" = $1`, l.Address)

	l.existsInDB = false

	return err
}
