/*
  AUTO-GENERATED file for Django model django_kea.Lease6

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

// Lease6 mirrors model django_kea.Lease6
type Lease6 struct {
	existsInDB bool

	Address       string
	Duid          sql.NullString
	ValidLifetime sql.NullInt64
	Expire        sql.NullTime
	SubnetId      sql.NullInt64
	PrefLifetime  sql.NullInt64
	leaseType     sql.NullInt32
	Iaid          sql.NullInt32
	PrefixLen     sql.NullInt32
	FqdnFwd       sql.NullBool
	FqdnRev       sql.NullBool
	Hostname      sql.NullString
	state         sql.NullInt64
	Hwaddr        sql.NullString
	Hwtype        sql.NullInt32
	HwaddrSource  sql.NullInt32
	UserContext   sql.NullString
}

// Lease6QS represents a queryset for django_kea.Lease6
type Lease6QS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Lease6QS) filter(c string, p interface{}) Lease6QS {
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
func (qs Lease6QS) Or(exprs ...Lease6QS) Lease6QS {
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
func (qs Lease6QS) AddressEq(v string) Lease6QS {
	return qs.filter(`"address" =`, v)
}

// AddressNe filters for Address being not equal to argument
func (qs Lease6QS) AddressNe(v string) Lease6QS {
	return qs.filter(`"address" <>`, v)
}

// AddressLt filters for Address being less than argument
func (qs Lease6QS) AddressLt(v string) Lease6QS {
	return qs.filter(`"address" <`, v)
}

// AddressLe filters for Address being less than or equal to argument
func (qs Lease6QS) AddressLe(v string) Lease6QS {
	return qs.filter(`"address" <=`, v)
}

// AddressGt filters for Address being greater than argument
func (qs Lease6QS) AddressGt(v string) Lease6QS {
	return qs.filter(`"address" >`, v)
}

// AddressGe filters for Address being greater than or equal to argument
func (qs Lease6QS) AddressGe(v string) Lease6QS {
	return qs.filter(`"address" >=`, v)
}

type inLease6Address struct {
	values []interface{}
}

func (in *inLease6Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) AddressIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Address{
			values: vals,
		},
	)

	return qs
}

type notinLease6Address struct {
	values []interface{}
}

func (in *notinLease6Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) AddressNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Address{
			values: vals,
		},
	)

	return qs
}

// OrderByAddress sorts result by Address in ascending order
func (qs Lease6QS) OrderByAddress() Lease6QS {
	qs.order = append(qs.order, `"address"`)

	return qs
}

// OrderByAddressDesc sorts result by Address in descending order
func (qs Lease6QS) OrderByAddressDesc() Lease6QS {
	qs.order = append(qs.order, `"address" DESC`)

	return qs
}

// DuidIsNull filters for Duid being null
func (qs Lease6QS) DuidIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"duid" IS NULL`,
		},
	)
	return qs
}

// DuidIsNotNull filters for Duid being not null
func (qs Lease6QS) DuidIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"duid" IS NOT NULL`,
		},
	)
	return qs
}

// DuidEq filters for Duid being equal to argument
func (qs Lease6QS) DuidEq(v string) Lease6QS {
	return qs.filter(`"duid" =`, v)
}

// DuidNe filters for Duid being not equal to argument
func (qs Lease6QS) DuidNe(v string) Lease6QS {
	return qs.filter(`"duid" <>`, v)
}

// DuidLt filters for Duid being less than argument
func (qs Lease6QS) DuidLt(v string) Lease6QS {
	return qs.filter(`"duid" <`, v)
}

// DuidLe filters for Duid being less than or equal to argument
func (qs Lease6QS) DuidLe(v string) Lease6QS {
	return qs.filter(`"duid" <=`, v)
}

// DuidGt filters for Duid being greater than argument
func (qs Lease6QS) DuidGt(v string) Lease6QS {
	return qs.filter(`"duid" >`, v)
}

// DuidGe filters for Duid being greater than or equal to argument
func (qs Lease6QS) DuidGe(v string) Lease6QS {
	return qs.filter(`"duid" >=`, v)
}

type inLease6Duid struct {
	values []interface{}
}

func (in *inLease6Duid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"duid" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) DuidIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Duid{
			values: vals,
		},
	)

	return qs
}

type notinLease6Duid struct {
	values []interface{}
}

func (in *notinLease6Duid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"duid" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) DuidNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Duid{
			values: vals,
		},
	)

	return qs
}

// OrderByDuid sorts result by Duid in ascending order
func (qs Lease6QS) OrderByDuid() Lease6QS {
	qs.order = append(qs.order, `"duid"`)

	return qs
}

// OrderByDuidDesc sorts result by Duid in descending order
func (qs Lease6QS) OrderByDuidDesc() Lease6QS {
	qs.order = append(qs.order, `"duid" DESC`)

	return qs
}

// ValidLifetimeIsNull filters for ValidLifetime being null
func (qs Lease6QS) ValidLifetimeIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"valid_lifetime" IS NULL`,
		},
	)
	return qs
}

// ValidLifetimeIsNotNull filters for ValidLifetime being not null
func (qs Lease6QS) ValidLifetimeIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"valid_lifetime" IS NOT NULL`,
		},
	)
	return qs
}

// ValidLifetimeEq filters for ValidLifetime being equal to argument
func (qs Lease6QS) ValidLifetimeEq(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" =`, v)
}

// ValidLifetimeNe filters for ValidLifetime being not equal to argument
func (qs Lease6QS) ValidLifetimeNe(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" <>`, v)
}

// ValidLifetimeLt filters for ValidLifetime being less than argument
func (qs Lease6QS) ValidLifetimeLt(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" <`, v)
}

// ValidLifetimeLe filters for ValidLifetime being less than or equal to argument
func (qs Lease6QS) ValidLifetimeLe(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" <=`, v)
}

// ValidLifetimeGt filters for ValidLifetime being greater than argument
func (qs Lease6QS) ValidLifetimeGt(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" >`, v)
}

// ValidLifetimeGe filters for ValidLifetime being greater than or equal to argument
func (qs Lease6QS) ValidLifetimeGe(v int64) Lease6QS {
	return qs.filter(`"valid_lifetime" >=`, v)
}

type inLease6ValidLifetime struct {
	values []interface{}
}

func (in *inLease6ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) ValidLifetimeIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6ValidLifetime{
			values: vals,
		},
	)

	return qs
}

type notinLease6ValidLifetime struct {
	values []interface{}
}

func (in *notinLease6ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) ValidLifetimeNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6ValidLifetime{
			values: vals,
		},
	)

	return qs
}

// OrderByValidLifetime sorts result by ValidLifetime in ascending order
func (qs Lease6QS) OrderByValidLifetime() Lease6QS {
	qs.order = append(qs.order, `"valid_lifetime"`)

	return qs
}

// OrderByValidLifetimeDesc sorts result by ValidLifetime in descending order
func (qs Lease6QS) OrderByValidLifetimeDesc() Lease6QS {
	qs.order = append(qs.order, `"valid_lifetime" DESC`)

	return qs
}

// ExpireIsNull filters for Expire being null
func (qs Lease6QS) ExpireIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"expire" IS NULL`,
		},
	)
	return qs
}

// ExpireIsNotNull filters for Expire being not null
func (qs Lease6QS) ExpireIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"expire" IS NOT NULL`,
		},
	)
	return qs
}

// ExpireEq filters for Expire being equal to argument
func (qs Lease6QS) ExpireEq(v time.Time) Lease6QS {
	return qs.filter(`"expire" =`, v)
}

// ExpireNe filters for Expire being not equal to argument
func (qs Lease6QS) ExpireNe(v time.Time) Lease6QS {
	return qs.filter(`"expire" <>`, v)
}

// ExpireLt filters for Expire being less than argument
func (qs Lease6QS) ExpireLt(v time.Time) Lease6QS {
	return qs.filter(`"expire" <`, v)
}

// ExpireLe filters for Expire being less than or equal to argument
func (qs Lease6QS) ExpireLe(v time.Time) Lease6QS {
	return qs.filter(`"expire" <=`, v)
}

// ExpireGt filters for Expire being greater than argument
func (qs Lease6QS) ExpireGt(v time.Time) Lease6QS {
	return qs.filter(`"expire" >`, v)
}

// ExpireGe filters for Expire being greater than or equal to argument
func (qs Lease6QS) ExpireGe(v time.Time) Lease6QS {
	return qs.filter(`"expire" >=`, v)
}

type inLease6Expire struct {
	values []interface{}
}

func (in *inLease6Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"expire" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) ExpireIn(values []time.Time) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Expire{
			values: vals,
		},
	)

	return qs
}

type notinLease6Expire struct {
	values []interface{}
}

func (in *notinLease6Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"expire" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) ExpireNotIn(values []time.Time) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Expire{
			values: vals,
		},
	)

	return qs
}

// OrderByExpire sorts result by Expire in ascending order
func (qs Lease6QS) OrderByExpire() Lease6QS {
	qs.order = append(qs.order, `"expire"`)

	return qs
}

// OrderByExpireDesc sorts result by Expire in descending order
func (qs Lease6QS) OrderByExpireDesc() Lease6QS {
	qs.order = append(qs.order, `"expire" DESC`)

	return qs
}

// SubnetIdIsNull filters for SubnetId being null
func (qs Lease6QS) SubnetIdIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"subnet_id" IS NULL`,
		},
	)
	return qs
}

// SubnetIdIsNotNull filters for SubnetId being not null
func (qs Lease6QS) SubnetIdIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// SubnetIdEq filters for SubnetId being equal to argument
func (qs Lease6QS) SubnetIdEq(v int64) Lease6QS {
	return qs.filter(`"subnet_id" =`, v)
}

// SubnetIdNe filters for SubnetId being not equal to argument
func (qs Lease6QS) SubnetIdNe(v int64) Lease6QS {
	return qs.filter(`"subnet_id" <>`, v)
}

// SubnetIdLt filters for SubnetId being less than argument
func (qs Lease6QS) SubnetIdLt(v int64) Lease6QS {
	return qs.filter(`"subnet_id" <`, v)
}

// SubnetIdLe filters for SubnetId being less than or equal to argument
func (qs Lease6QS) SubnetIdLe(v int64) Lease6QS {
	return qs.filter(`"subnet_id" <=`, v)
}

// SubnetIdGt filters for SubnetId being greater than argument
func (qs Lease6QS) SubnetIdGt(v int64) Lease6QS {
	return qs.filter(`"subnet_id" >`, v)
}

// SubnetIdGe filters for SubnetId being greater than or equal to argument
func (qs Lease6QS) SubnetIdGe(v int64) Lease6QS {
	return qs.filter(`"subnet_id" >=`, v)
}

type inLease6SubnetId struct {
	values []interface{}
}

func (in *inLease6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) SubnetIdIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinLease6SubnetId struct {
	values []interface{}
}

func (in *notinLease6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) SubnetIdNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderBySubnetId sorts result by SubnetId in ascending order
func (qs Lease6QS) OrderBySubnetId() Lease6QS {
	qs.order = append(qs.order, `"subnet_id"`)

	return qs
}

// OrderBySubnetIdDesc sorts result by SubnetId in descending order
func (qs Lease6QS) OrderBySubnetIdDesc() Lease6QS {
	qs.order = append(qs.order, `"subnet_id" DESC`)

	return qs
}

// PrefLifetimeIsNull filters for PrefLifetime being null
func (qs Lease6QS) PrefLifetimeIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"pref_lifetime" IS NULL`,
		},
	)
	return qs
}

// PrefLifetimeIsNotNull filters for PrefLifetime being not null
func (qs Lease6QS) PrefLifetimeIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"pref_lifetime" IS NOT NULL`,
		},
	)
	return qs
}

// PrefLifetimeEq filters for PrefLifetime being equal to argument
func (qs Lease6QS) PrefLifetimeEq(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" =`, v)
}

// PrefLifetimeNe filters for PrefLifetime being not equal to argument
func (qs Lease6QS) PrefLifetimeNe(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" <>`, v)
}

// PrefLifetimeLt filters for PrefLifetime being less than argument
func (qs Lease6QS) PrefLifetimeLt(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" <`, v)
}

// PrefLifetimeLe filters for PrefLifetime being less than or equal to argument
func (qs Lease6QS) PrefLifetimeLe(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" <=`, v)
}

// PrefLifetimeGt filters for PrefLifetime being greater than argument
func (qs Lease6QS) PrefLifetimeGt(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" >`, v)
}

// PrefLifetimeGe filters for PrefLifetime being greater than or equal to argument
func (qs Lease6QS) PrefLifetimeGe(v int64) Lease6QS {
	return qs.filter(`"pref_lifetime" >=`, v)
}

type inLease6PrefLifetime struct {
	values []interface{}
}

func (in *inLease6PrefLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"pref_lifetime" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) PrefLifetimeIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6PrefLifetime{
			values: vals,
		},
	)

	return qs
}

type notinLease6PrefLifetime struct {
	values []interface{}
}

func (in *notinLease6PrefLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"pref_lifetime" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) PrefLifetimeNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6PrefLifetime{
			values: vals,
		},
	)

	return qs
}

// OrderByPrefLifetime sorts result by PrefLifetime in ascending order
func (qs Lease6QS) OrderByPrefLifetime() Lease6QS {
	qs.order = append(qs.order, `"pref_lifetime"`)

	return qs
}

// OrderByPrefLifetimeDesc sorts result by PrefLifetime in descending order
func (qs Lease6QS) OrderByPrefLifetimeDesc() Lease6QS {
	qs.order = append(qs.order, `"pref_lifetime" DESC`)

	return qs
}

// GetLeaseType returns Lease6type
func (l *Lease6) GetLeaseType(db models.DBInterface) (*Lease6type, error) {
	if !l.leaseType.Valid {
		return nil, nil
	}

	return Lease6typeQS{}.LeaseTypeEq(l.leaseType.Int32).First(db)
}

// SetLeaseType sets foreign key pointer to Lease6type
func (l *Lease6) SetLeaseType(ptr *Lease6type) error {
	if ptr != nil {
		l.leaseType.Int32 = ptr.LeaseType
		l.leaseType.Valid = true
	} else {
		l.leaseType.Valid = false
	}

	return nil
}

// GetLeaseTypeRaw returns Lease6.LeaseType
func (l *Lease6) GetLeaseTypeRaw() sql.NullInt32 {
	return l.leaseType
}

// LeaseTypeIsNull filters for leaseType being null
func (qs Lease6QS) LeaseTypeIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"lease_type" IS NULL`,
		},
	)
	return qs
}

// LeaseTypeIsNotNull filters for leaseType being not null
func (qs Lease6QS) LeaseTypeIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"lease_type" IS NOT NULL`,
		},
	)
	return qs
}

// LeaseTypeEq filters for leaseType being equal to argument
func (qs Lease6QS) LeaseTypeEq(v *Lease6type) Lease6QS {
	return qs.filter(`"lease_type" =`, v.LeaseType)
}

// LeaseTypeRawEq filters for leaseType being equal to raw argument
func (qs Lease6QS) LeaseTypeRawEq(v int32) Lease6QS {
	return qs.filter(`"lease_type" =`, v)
}

type inLease6leaseTypeLease6type struct {
	qs Lease6typeQS
}

func (in *inLease6leaseTypeLease6type) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"lease_type" IN (` + s + `)`, p
}

func (qs Lease6QS) LeaseTypeIn(oqs Lease6typeQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&inLease6leaseTypeLease6type{
			qs: oqs,
		},
	)

	return qs
}

// OrderByLeaseType sorts result by LeaseType in ascending order
func (qs Lease6QS) OrderByLeaseType() Lease6QS {
	qs.order = append(qs.order, `"lease_type"`)

	return qs
}

// OrderByLeaseTypeDesc sorts result by LeaseType in descending order
func (qs Lease6QS) OrderByLeaseTypeDesc() Lease6QS {
	qs.order = append(qs.order, `"lease_type" DESC`)

	return qs
}

// IaidIsNull filters for Iaid being null
func (qs Lease6QS) IaidIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"iaid" IS NULL`,
		},
	)
	return qs
}

// IaidIsNotNull filters for Iaid being not null
func (qs Lease6QS) IaidIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"iaid" IS NOT NULL`,
		},
	)
	return qs
}

// IaidEq filters for Iaid being equal to argument
func (qs Lease6QS) IaidEq(v int32) Lease6QS {
	return qs.filter(`"iaid" =`, v)
}

// IaidNe filters for Iaid being not equal to argument
func (qs Lease6QS) IaidNe(v int32) Lease6QS {
	return qs.filter(`"iaid" <>`, v)
}

// IaidLt filters for Iaid being less than argument
func (qs Lease6QS) IaidLt(v int32) Lease6QS {
	return qs.filter(`"iaid" <`, v)
}

// IaidLe filters for Iaid being less than or equal to argument
func (qs Lease6QS) IaidLe(v int32) Lease6QS {
	return qs.filter(`"iaid" <=`, v)
}

// IaidGt filters for Iaid being greater than argument
func (qs Lease6QS) IaidGt(v int32) Lease6QS {
	return qs.filter(`"iaid" >`, v)
}

// IaidGe filters for Iaid being greater than or equal to argument
func (qs Lease6QS) IaidGe(v int32) Lease6QS {
	return qs.filter(`"iaid" >=`, v)
}

type inLease6Iaid struct {
	values []interface{}
}

func (in *inLease6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"iaid" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) IaidIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Iaid{
			values: vals,
		},
	)

	return qs
}

type notinLease6Iaid struct {
	values []interface{}
}

func (in *notinLease6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"iaid" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) IaidNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Iaid{
			values: vals,
		},
	)

	return qs
}

// OrderByIaid sorts result by Iaid in ascending order
func (qs Lease6QS) OrderByIaid() Lease6QS {
	qs.order = append(qs.order, `"iaid"`)

	return qs
}

// OrderByIaidDesc sorts result by Iaid in descending order
func (qs Lease6QS) OrderByIaidDesc() Lease6QS {
	qs.order = append(qs.order, `"iaid" DESC`)

	return qs
}

// PrefixLenIsNull filters for PrefixLen being null
func (qs Lease6QS) PrefixLenIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"prefix_len" IS NULL`,
		},
	)
	return qs
}

// PrefixLenIsNotNull filters for PrefixLen being not null
func (qs Lease6QS) PrefixLenIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"prefix_len" IS NOT NULL`,
		},
	)
	return qs
}

// PrefixLenEq filters for PrefixLen being equal to argument
func (qs Lease6QS) PrefixLenEq(v int32) Lease6QS {
	return qs.filter(`"prefix_len" =`, v)
}

// PrefixLenNe filters for PrefixLen being not equal to argument
func (qs Lease6QS) PrefixLenNe(v int32) Lease6QS {
	return qs.filter(`"prefix_len" <>`, v)
}

// PrefixLenLt filters for PrefixLen being less than argument
func (qs Lease6QS) PrefixLenLt(v int32) Lease6QS {
	return qs.filter(`"prefix_len" <`, v)
}

// PrefixLenLe filters for PrefixLen being less than or equal to argument
func (qs Lease6QS) PrefixLenLe(v int32) Lease6QS {
	return qs.filter(`"prefix_len" <=`, v)
}

// PrefixLenGt filters for PrefixLen being greater than argument
func (qs Lease6QS) PrefixLenGt(v int32) Lease6QS {
	return qs.filter(`"prefix_len" >`, v)
}

// PrefixLenGe filters for PrefixLen being greater than or equal to argument
func (qs Lease6QS) PrefixLenGe(v int32) Lease6QS {
	return qs.filter(`"prefix_len" >=`, v)
}

type inLease6PrefixLen struct {
	values []interface{}
}

func (in *inLease6PrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"prefix_len" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) PrefixLenIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6PrefixLen{
			values: vals,
		},
	)

	return qs
}

type notinLease6PrefixLen struct {
	values []interface{}
}

func (in *notinLease6PrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"prefix_len" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) PrefixLenNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6PrefixLen{
			values: vals,
		},
	)

	return qs
}

// OrderByPrefixLen sorts result by PrefixLen in ascending order
func (qs Lease6QS) OrderByPrefixLen() Lease6QS {
	qs.order = append(qs.order, `"prefix_len"`)

	return qs
}

// OrderByPrefixLenDesc sorts result by PrefixLen in descending order
func (qs Lease6QS) OrderByPrefixLenDesc() Lease6QS {
	qs.order = append(qs.order, `"prefix_len" DESC`)

	return qs
}

// FqdnFwdIsNull filters for FqdnFwd being null
func (qs Lease6QS) FqdnFwdIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_fwd" IS NULL`,
		},
	)
	return qs
}

// FqdnFwdIsNotNull filters for FqdnFwd being not null
func (qs Lease6QS) FqdnFwdIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_fwd" IS NOT NULL`,
		},
	)
	return qs
}

// FqdnFwdEq filters for FqdnFwd being equal to argument
func (qs Lease6QS) FqdnFwdEq(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" =`, v)
}

// FqdnFwdNe filters for FqdnFwd being not equal to argument
func (qs Lease6QS) FqdnFwdNe(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" <>`, v)
}

// FqdnFwdLt filters for FqdnFwd being less than argument
func (qs Lease6QS) FqdnFwdLt(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" <`, v)
}

// FqdnFwdLe filters for FqdnFwd being less than or equal to argument
func (qs Lease6QS) FqdnFwdLe(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" <=`, v)
}

// FqdnFwdGt filters for FqdnFwd being greater than argument
func (qs Lease6QS) FqdnFwdGt(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" >`, v)
}

// FqdnFwdGe filters for FqdnFwd being greater than or equal to argument
func (qs Lease6QS) FqdnFwdGe(v bool) Lease6QS {
	return qs.filter(`"fqdn_fwd" >=`, v)
}

type inLease6FqdnFwd struct {
	values []interface{}
}

func (in *inLease6FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) FqdnFwdIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6FqdnFwd{
			values: vals,
		},
	)

	return qs
}

type notinLease6FqdnFwd struct {
	values []interface{}
}

func (in *notinLease6FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) FqdnFwdNotIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6FqdnFwd{
			values: vals,
		},
	)

	return qs
}

// OrderByFqdnFwd sorts result by FqdnFwd in ascending order
func (qs Lease6QS) OrderByFqdnFwd() Lease6QS {
	qs.order = append(qs.order, `"fqdn_fwd"`)

	return qs
}

// OrderByFqdnFwdDesc sorts result by FqdnFwd in descending order
func (qs Lease6QS) OrderByFqdnFwdDesc() Lease6QS {
	qs.order = append(qs.order, `"fqdn_fwd" DESC`)

	return qs
}

// FqdnRevIsNull filters for FqdnRev being null
func (qs Lease6QS) FqdnRevIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_rev" IS NULL`,
		},
	)
	return qs
}

// FqdnRevIsNotNull filters for FqdnRev being not null
func (qs Lease6QS) FqdnRevIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"fqdn_rev" IS NOT NULL`,
		},
	)
	return qs
}

// FqdnRevEq filters for FqdnRev being equal to argument
func (qs Lease6QS) FqdnRevEq(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" =`, v)
}

// FqdnRevNe filters for FqdnRev being not equal to argument
func (qs Lease6QS) FqdnRevNe(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" <>`, v)
}

// FqdnRevLt filters for FqdnRev being less than argument
func (qs Lease6QS) FqdnRevLt(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" <`, v)
}

// FqdnRevLe filters for FqdnRev being less than or equal to argument
func (qs Lease6QS) FqdnRevLe(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" <=`, v)
}

// FqdnRevGt filters for FqdnRev being greater than argument
func (qs Lease6QS) FqdnRevGt(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" >`, v)
}

// FqdnRevGe filters for FqdnRev being greater than or equal to argument
func (qs Lease6QS) FqdnRevGe(v bool) Lease6QS {
	return qs.filter(`"fqdn_rev" >=`, v)
}

type inLease6FqdnRev struct {
	values []interface{}
}

func (in *inLease6FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) FqdnRevIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6FqdnRev{
			values: vals,
		},
	)

	return qs
}

type notinLease6FqdnRev struct {
	values []interface{}
}

func (in *notinLease6FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) FqdnRevNotIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6FqdnRev{
			values: vals,
		},
	)

	return qs
}

// OrderByFqdnRev sorts result by FqdnRev in ascending order
func (qs Lease6QS) OrderByFqdnRev() Lease6QS {
	qs.order = append(qs.order, `"fqdn_rev"`)

	return qs
}

// OrderByFqdnRevDesc sorts result by FqdnRev in descending order
func (qs Lease6QS) OrderByFqdnRevDesc() Lease6QS {
	qs.order = append(qs.order, `"fqdn_rev" DESC`)

	return qs
}

// HostnameIsNull filters for Hostname being null
func (qs Lease6QS) HostnameIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NULL`,
		},
	)
	return qs
}

// HostnameIsNotNull filters for Hostname being not null
func (qs Lease6QS) HostnameIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NOT NULL`,
		},
	)
	return qs
}

// HostnameEq filters for Hostname being equal to argument
func (qs Lease6QS) HostnameEq(v string) Lease6QS {
	return qs.filter(`"hostname" =`, v)
}

// HostnameNe filters for Hostname being not equal to argument
func (qs Lease6QS) HostnameNe(v string) Lease6QS {
	return qs.filter(`"hostname" <>`, v)
}

// HostnameLt filters for Hostname being less than argument
func (qs Lease6QS) HostnameLt(v string) Lease6QS {
	return qs.filter(`"hostname" <`, v)
}

// HostnameLe filters for Hostname being less than or equal to argument
func (qs Lease6QS) HostnameLe(v string) Lease6QS {
	return qs.filter(`"hostname" <=`, v)
}

// HostnameGt filters for Hostname being greater than argument
func (qs Lease6QS) HostnameGt(v string) Lease6QS {
	return qs.filter(`"hostname" >`, v)
}

// HostnameGe filters for Hostname being greater than or equal to argument
func (qs Lease6QS) HostnameGe(v string) Lease6QS {
	return qs.filter(`"hostname" >=`, v)
}

type inLease6Hostname struct {
	values []interface{}
}

func (in *inLease6Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HostnameIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Hostname{
			values: vals,
		},
	)

	return qs
}

type notinLease6Hostname struct {
	values []interface{}
}

func (in *notinLease6Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HostnameNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Hostname{
			values: vals,
		},
	)

	return qs
}

// OrderByHostname sorts result by Hostname in ascending order
func (qs Lease6QS) OrderByHostname() Lease6QS {
	qs.order = append(qs.order, `"hostname"`)

	return qs
}

// OrderByHostnameDesc sorts result by Hostname in descending order
func (qs Lease6QS) OrderByHostnameDesc() Lease6QS {
	qs.order = append(qs.order, `"hostname" DESC`)

	return qs
}

// GetState returns Leasestate
func (l *Lease6) GetState(db models.DBInterface) (*Leasestate, error) {
	if !l.state.Valid {
		return nil, nil
	}

	return LeasestateQS{}.StateEq(l.state.Int64).First(db)
}

// SetState sets foreign key pointer to Leasestate
func (l *Lease6) SetState(ptr *Leasestate) error {
	if ptr != nil {
		l.state.Int64 = ptr.State
		l.state.Valid = true
	} else {
		l.state.Valid = false
	}

	return nil
}

// GetStateRaw returns Lease6.State
func (l *Lease6) GetStateRaw() sql.NullInt64 {
	return l.state
}

// StateIsNull filters for state being null
func (qs Lease6QS) StateIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"state" IS NULL`,
		},
	)
	return qs
}

// StateIsNotNull filters for state being not null
func (qs Lease6QS) StateIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"state" IS NOT NULL`,
		},
	)
	return qs
}

// StateEq filters for state being equal to argument
func (qs Lease6QS) StateEq(v *Leasestate) Lease6QS {
	return qs.filter(`"state" =`, v.State)
}

// StateRawEq filters for state being equal to raw argument
func (qs Lease6QS) StateRawEq(v int64) Lease6QS {
	return qs.filter(`"state" =`, v)
}

type inLease6stateLeasestate struct {
	qs LeasestateQS
}

func (in *inLease6stateLeasestate) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"state" IN (` + s + `)`, p
}

func (qs Lease6QS) StateIn(oqs LeasestateQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&inLease6stateLeasestate{
			qs: oqs,
		},
	)

	return qs
}

// OrderByState sorts result by State in ascending order
func (qs Lease6QS) OrderByState() Lease6QS {
	qs.order = append(qs.order, `"state"`)

	return qs
}

// OrderByStateDesc sorts result by State in descending order
func (qs Lease6QS) OrderByStateDesc() Lease6QS {
	qs.order = append(qs.order, `"state" DESC`)

	return qs
}

// HwaddrIsNull filters for Hwaddr being null
func (qs Lease6QS) HwaddrIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr" IS NULL`,
		},
	)
	return qs
}

// HwaddrIsNotNull filters for Hwaddr being not null
func (qs Lease6QS) HwaddrIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr" IS NOT NULL`,
		},
	)
	return qs
}

// HwaddrEq filters for Hwaddr being equal to argument
func (qs Lease6QS) HwaddrEq(v string) Lease6QS {
	return qs.filter(`"hwaddr" =`, v)
}

// HwaddrNe filters for Hwaddr being not equal to argument
func (qs Lease6QS) HwaddrNe(v string) Lease6QS {
	return qs.filter(`"hwaddr" <>`, v)
}

// HwaddrLt filters for Hwaddr being less than argument
func (qs Lease6QS) HwaddrLt(v string) Lease6QS {
	return qs.filter(`"hwaddr" <`, v)
}

// HwaddrLe filters for Hwaddr being less than or equal to argument
func (qs Lease6QS) HwaddrLe(v string) Lease6QS {
	return qs.filter(`"hwaddr" <=`, v)
}

// HwaddrGt filters for Hwaddr being greater than argument
func (qs Lease6QS) HwaddrGt(v string) Lease6QS {
	return qs.filter(`"hwaddr" >`, v)
}

// HwaddrGe filters for Hwaddr being greater than or equal to argument
func (qs Lease6QS) HwaddrGe(v string) Lease6QS {
	return qs.filter(`"hwaddr" >=`, v)
}

type inLease6Hwaddr struct {
	values []interface{}
}

func (in *inLease6Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwaddrIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Hwaddr{
			values: vals,
		},
	)

	return qs
}

type notinLease6Hwaddr struct {
	values []interface{}
}

func (in *notinLease6Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwaddrNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Hwaddr{
			values: vals,
		},
	)

	return qs
}

// OrderByHwaddr sorts result by Hwaddr in ascending order
func (qs Lease6QS) OrderByHwaddr() Lease6QS {
	qs.order = append(qs.order, `"hwaddr"`)

	return qs
}

// OrderByHwaddrDesc sorts result by Hwaddr in descending order
func (qs Lease6QS) OrderByHwaddrDesc() Lease6QS {
	qs.order = append(qs.order, `"hwaddr" DESC`)

	return qs
}

// HwtypeIsNull filters for Hwtype being null
func (qs Lease6QS) HwtypeIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwtype" IS NULL`,
		},
	)
	return qs
}

// HwtypeIsNotNull filters for Hwtype being not null
func (qs Lease6QS) HwtypeIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwtype" IS NOT NULL`,
		},
	)
	return qs
}

// HwtypeEq filters for Hwtype being equal to argument
func (qs Lease6QS) HwtypeEq(v int32) Lease6QS {
	return qs.filter(`"hwtype" =`, v)
}

// HwtypeNe filters for Hwtype being not equal to argument
func (qs Lease6QS) HwtypeNe(v int32) Lease6QS {
	return qs.filter(`"hwtype" <>`, v)
}

// HwtypeLt filters for Hwtype being less than argument
func (qs Lease6QS) HwtypeLt(v int32) Lease6QS {
	return qs.filter(`"hwtype" <`, v)
}

// HwtypeLe filters for Hwtype being less than or equal to argument
func (qs Lease6QS) HwtypeLe(v int32) Lease6QS {
	return qs.filter(`"hwtype" <=`, v)
}

// HwtypeGt filters for Hwtype being greater than argument
func (qs Lease6QS) HwtypeGt(v int32) Lease6QS {
	return qs.filter(`"hwtype" >`, v)
}

// HwtypeGe filters for Hwtype being greater than or equal to argument
func (qs Lease6QS) HwtypeGe(v int32) Lease6QS {
	return qs.filter(`"hwtype" >=`, v)
}

type inLease6Hwtype struct {
	values []interface{}
}

func (in *inLease6Hwtype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwtype" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwtypeIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6Hwtype{
			values: vals,
		},
	)

	return qs
}

type notinLease6Hwtype struct {
	values []interface{}
}

func (in *notinLease6Hwtype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwtype" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwtypeNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6Hwtype{
			values: vals,
		},
	)

	return qs
}

// OrderByHwtype sorts result by Hwtype in ascending order
func (qs Lease6QS) OrderByHwtype() Lease6QS {
	qs.order = append(qs.order, `"hwtype"`)

	return qs
}

// OrderByHwtypeDesc sorts result by Hwtype in descending order
func (qs Lease6QS) OrderByHwtypeDesc() Lease6QS {
	qs.order = append(qs.order, `"hwtype" DESC`)

	return qs
}

// HwaddrSourceIsNull filters for HwaddrSource being null
func (qs Lease6QS) HwaddrSourceIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr_source" IS NULL`,
		},
	)
	return qs
}

// HwaddrSourceIsNotNull filters for HwaddrSource being not null
func (qs Lease6QS) HwaddrSourceIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr_source" IS NOT NULL`,
		},
	)
	return qs
}

// HwaddrSourceEq filters for HwaddrSource being equal to argument
func (qs Lease6QS) HwaddrSourceEq(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" =`, v)
}

// HwaddrSourceNe filters for HwaddrSource being not equal to argument
func (qs Lease6QS) HwaddrSourceNe(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" <>`, v)
}

// HwaddrSourceLt filters for HwaddrSource being less than argument
func (qs Lease6QS) HwaddrSourceLt(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" <`, v)
}

// HwaddrSourceLe filters for HwaddrSource being less than or equal to argument
func (qs Lease6QS) HwaddrSourceLe(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" <=`, v)
}

// HwaddrSourceGt filters for HwaddrSource being greater than argument
func (qs Lease6QS) HwaddrSourceGt(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" >`, v)
}

// HwaddrSourceGe filters for HwaddrSource being greater than or equal to argument
func (qs Lease6QS) HwaddrSourceGe(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" >=`, v)
}

type inLease6HwaddrSource struct {
	values []interface{}
}

func (in *inLease6HwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwaddrSourceIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6HwaddrSource{
			values: vals,
		},
	)

	return qs
}

type notinLease6HwaddrSource struct {
	values []interface{}
}

func (in *notinLease6HwaddrSource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hwaddr_source" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) HwaddrSourceNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6HwaddrSource{
			values: vals,
		},
	)

	return qs
}

// OrderByHwaddrSource sorts result by HwaddrSource in ascending order
func (qs Lease6QS) OrderByHwaddrSource() Lease6QS {
	qs.order = append(qs.order, `"hwaddr_source"`)

	return qs
}

// OrderByHwaddrSourceDesc sorts result by HwaddrSource in descending order
func (qs Lease6QS) OrderByHwaddrSourceDesc() Lease6QS {
	qs.order = append(qs.order, `"hwaddr_source" DESC`)

	return qs
}

// UserContextIsNull filters for UserContext being null
func (qs Lease6QS) UserContextIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs Lease6QS) UserContextIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs Lease6QS) UserContextEq(v string) Lease6QS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs Lease6QS) UserContextNe(v string) Lease6QS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs Lease6QS) UserContextLt(v string) Lease6QS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs Lease6QS) UserContextLe(v string) Lease6QS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs Lease6QS) UserContextGt(v string) Lease6QS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs Lease6QS) UserContextGe(v string) Lease6QS {
	return qs.filter(`"user_context" >=`, v)
}

type inLease6UserContext struct {
	values []interface{}
}

func (in *inLease6UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) UserContextIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inLease6UserContext{
			values: vals,
		},
	)

	return qs
}

type notinLease6UserContext struct {
	values []interface{}
}

func (in *notinLease6UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Lease6QS) UserContextNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinLease6UserContext{
			values: vals,
		},
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs Lease6QS) OrderByUserContext() Lease6QS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs Lease6QS) OrderByUserContextDesc() Lease6QS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Lease6QS) ForUpdate() Lease6QS {
	qs.forUpdate = true

	return qs
}

func (qs Lease6QS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Lease6QS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Lease6QS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "address", "duid", "valid_lifetime", "expire", "subnet_id", "pref_lifetime", "lease_type", "iaid", "prefix_len", "fqdn_fwd", "fqdn_rev", "hostname", "state", "hwaddr", "hwtype", "hwaddr_source", "user_context" FROM "lease6"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Lease6QS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "address" FROM "lease6"` + s, p
}

// All returns all rows matching queryset filters
func (qs Lease6QS) All(db models.DBInterface) ([]*Lease6, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Lease6
	for rows.Next() {
		obj := Lease6{existsInDB: true}
		if err = rows.Scan(&obj.Address, &obj.Duid, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.PrefLifetime, &obj.leaseType, &obj.Iaid, &obj.PrefixLen, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.Hwaddr, &obj.Hwtype, &obj.HwaddrSource, &obj.UserContext); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Lease6QS) First(db models.DBInterface) (*Lease6, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Lease6{existsInDB: true}
	err := row.Scan(&obj.Address, &obj.Duid, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.PrefLifetime, &obj.leaseType, &obj.Iaid, &obj.PrefixLen, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.Hwaddr, &obj.Hwtype, &obj.HwaddrSource, &obj.UserContext)
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
func (qs Lease6QS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease6"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Lease6QS) Update() Lease6UpdateQS {
	return Lease6UpdateQS{condFragments: qs.condFragments}
}

// Lease6UpdateQS represents an updated queryset for django_kea.Lease6
type Lease6UpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Lease6UpdateQS) update(c string, v interface{}) Lease6UpdateQS {
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
func (uqs Lease6UpdateQS) SetAddress(v string) Lease6UpdateQS {
	return uqs.update(`"address"`, v)
}

// SetDuid sets Duid to the given value
func (uqs Lease6UpdateQS) SetDuid(v sql.NullString) Lease6UpdateQS {
	return uqs.update(`"duid"`, v)
}

// SetValidLifetime sets ValidLifetime to the given value
func (uqs Lease6UpdateQS) SetValidLifetime(v sql.NullInt64) Lease6UpdateQS {
	return uqs.update(`"valid_lifetime"`, v)
}

// SetExpire sets Expire to the given value
func (uqs Lease6UpdateQS) SetExpire(v sql.NullTime) Lease6UpdateQS {
	return uqs.update(`"expire"`, v)
}

// SetSubnetId sets SubnetId to the given value
func (uqs Lease6UpdateQS) SetSubnetId(v sql.NullInt64) Lease6UpdateQS {
	return uqs.update(`"subnet_id"`, v)
}

// SetPrefLifetime sets PrefLifetime to the given value
func (uqs Lease6UpdateQS) SetPrefLifetime(v sql.NullInt64) Lease6UpdateQS {
	return uqs.update(`"pref_lifetime"`, v)
}

// SetLeaseType sets foreign key pointer to Lease6type
func (uqs Lease6UpdateQS) SetLeaseType(ptr *Lease6type) Lease6UpdateQS {
	if ptr != nil {
		return uqs.update(`"lease_type"`, ptr.LeaseType)
	}

	return uqs.update(`"lease_type"`, nil)
} // SetIaid sets Iaid to the given value
func (uqs Lease6UpdateQS) SetIaid(v sql.NullInt32) Lease6UpdateQS {
	return uqs.update(`"iaid"`, v)
}

// SetPrefixLen sets PrefixLen to the given value
func (uqs Lease6UpdateQS) SetPrefixLen(v sql.NullInt32) Lease6UpdateQS {
	return uqs.update(`"prefix_len"`, v)
}

// SetFqdnFwd sets FqdnFwd to the given value
func (uqs Lease6UpdateQS) SetFqdnFwd(v sql.NullBool) Lease6UpdateQS {
	return uqs.update(`"fqdn_fwd"`, v)
}

// SetFqdnRev sets FqdnRev to the given value
func (uqs Lease6UpdateQS) SetFqdnRev(v sql.NullBool) Lease6UpdateQS {
	return uqs.update(`"fqdn_rev"`, v)
}

// SetHostname sets Hostname to the given value
func (uqs Lease6UpdateQS) SetHostname(v sql.NullString) Lease6UpdateQS {
	return uqs.update(`"hostname"`, v)
}

// SetState sets foreign key pointer to Leasestate
func (uqs Lease6UpdateQS) SetState(ptr *Leasestate) Lease6UpdateQS {
	if ptr != nil {
		return uqs.update(`"state"`, ptr.State)
	}

	return uqs.update(`"state"`, nil)
} // SetHwaddr sets Hwaddr to the given value
func (uqs Lease6UpdateQS) SetHwaddr(v sql.NullString) Lease6UpdateQS {
	return uqs.update(`"hwaddr"`, v)
}

// SetHwtype sets Hwtype to the given value
func (uqs Lease6UpdateQS) SetHwtype(v sql.NullInt32) Lease6UpdateQS {
	return uqs.update(`"hwtype"`, v)
}

// SetHwaddrSource sets HwaddrSource to the given value
func (uqs Lease6UpdateQS) SetHwaddrSource(v sql.NullInt32) Lease6UpdateQS {
	return uqs.update(`"hwaddr_source"`, v)
}

// SetUserContext sets UserContext to the given value
func (uqs Lease6UpdateQS) SetUserContext(v sql.NullString) Lease6UpdateQS {
	return uqs.update(`"user_context"`, v)
}

// Exec executes the update operation
func (uqs Lease6UpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Lease6QS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "lease6" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (l *Lease6) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "lease6" ("duid", "valid_lifetime", "expire", "subnet_id", "pref_lifetime", "lease_type", "iaid", "prefix_len", "fqdn_fwd", "fqdn_rev", "hostname", "state", "hwaddr", "hwtype", "hwaddr_source", "user_context", "address") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`, l.Duid, l.ValidLifetime, l.Expire, l.SubnetId, l.PrefLifetime, l.leaseType, l.Iaid, l.PrefixLen, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.Hwaddr, l.Hwtype, l.HwaddrSource, l.UserContext, l.Address)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Lease6) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "lease6" SET "duid" = $1, "valid_lifetime" = $2, "expire" = $3, "subnet_id" = $4, "pref_lifetime" = $5, "lease_type" = $6, "iaid" = $7, "prefix_len" = $8, "fqdn_fwd" = $9, "fqdn_rev" = $10, "hostname" = $11, "state" = $12, "hwaddr" = $13, "hwtype" = $14, "hwaddr_source" = $15, "user_context" = $16 WHERE "address" = $17`, l.Duid, l.ValidLifetime, l.Expire, l.SubnetId, l.PrefLifetime, l.leaseType, l.Iaid, l.PrefixLen, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.Hwaddr, l.Hwtype, l.HwaddrSource, l.UserContext, l.Address)

	return err
}

// Save inserts or updates record
func (l *Lease6) Save(db models.DBInterface) error {
	if l.existsInDB {
		return l.update(db)
	}

	return l.insert(db)
}

// Delete removes row from database
func (l *Lease6) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "lease6" WHERE "address" = $1`, l.Address)

	l.existsInDB = false

	return err
}
