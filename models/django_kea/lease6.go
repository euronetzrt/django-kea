// Code generated for Django model django_kea.Lease6. DO NOT EDIT.

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
	"net"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/euronetzrt/django-kea/models"
)

// Lease6 mirrors model django_kea.Lease6
type Lease6 struct {
	existsInDB bool

	Address       net.IP
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
	hwaddrSource  sql.NullInt32
	UserContext   sql.NullString
	PoolId        int64
}

// Lease6List is a list of Lease6
type Lease6List []*Lease6

// Lease6QS represents a queryset for django_kea.Lease6
type Lease6QS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
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

// BEGIN - django_kea.Lease6.address

// AddressEq filters for Address being equal to argument
func (qs Lease6QS) AddressEq(v net.IP) Lease6QS {
	return qs.filter(`"address" =`, v)
}

// AddressNe filters for Address being not equal to argument
func (qs Lease6QS) AddressNe(v net.IP) Lease6QS {
	return qs.filter(`"address" <>`, v)
}

// AddressLt filters for Address being less than argument
func (qs Lease6QS) AddressLt(v net.IP) Lease6QS {
	return qs.filter(`"address" <`, v)
}

// AddressLe filters for Address being less than or equal to argument
func (qs Lease6QS) AddressLe(v net.IP) Lease6QS {
	return qs.filter(`"address" <=`, v)
}

// AddressGt filters for Address being greater than argument
func (qs Lease6QS) AddressGt(v net.IP) Lease6QS {
	return qs.filter(`"address" >`, v)
}

// AddressGe filters for Address being greater than or equal to argument
func (qs Lease6QS) AddressGe(v net.IP) Lease6QS {
	return qs.filter(`"address" >=`, v)
}

type inLease6Address []interface{}

func (in inLease6Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"address" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) AddressIn(values []net.IP) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Address(vals),
	)

	return qs
}

type notinLease6Address []interface{}

func (in notinLease6Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"address" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) AddressNotIn(values []net.IP) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Address(vals),
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

// DistinctOnAddress marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnAddress() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"address"`)

	return qs
}

// END - django_kea.Lease6.address

// BEGIN - django_kea.Lease6.duid

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

type inLease6Duid []interface{}

func (in inLease6Duid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"duid" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) DuidIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Duid(vals),
	)

	return qs
}

type notinLease6Duid []interface{}

func (in notinLease6Duid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"duid" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) DuidNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Duid(vals),
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

// DistinctOnDuid marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnDuid() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"duid"`)

	return qs
}

// END - django_kea.Lease6.duid

// BEGIN - django_kea.Lease6.valid_lifetime

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

type inLease6ValidLifetime []interface{}

func (in inLease6ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) ValidLifetimeIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6ValidLifetime(vals),
	)

	return qs
}

type notinLease6ValidLifetime []interface{}

func (in notinLease6ValidLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"valid_lifetime" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) ValidLifetimeNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6ValidLifetime(vals),
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

// DistinctOnValidLifetime marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnValidLifetime() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"valid_lifetime"`)

	return qs
}

// END - django_kea.Lease6.valid_lifetime

// BEGIN - django_kea.Lease6.expire

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

type inLease6Expire []interface{}

func (in inLease6Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"expire" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) ExpireIn(values []time.Time) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Expire(vals),
	)

	return qs
}

type notinLease6Expire []interface{}

func (in notinLease6Expire) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"expire" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) ExpireNotIn(values []time.Time) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Expire(vals),
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

// DistinctOnExpire marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnExpire() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"expire"`)

	return qs
}

// END - django_kea.Lease6.expire

// BEGIN - django_kea.Lease6.subnet_id

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

type inLease6SubnetId []interface{}

func (in inLease6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"subnet_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) SubnetIdIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6SubnetId(vals),
	)

	return qs
}

type notinLease6SubnetId []interface{}

func (in notinLease6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) SubnetIdNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6SubnetId(vals),
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

// DistinctOnSubnetId marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnSubnetId() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"subnet_id"`)

	return qs
}

// END - django_kea.Lease6.subnet_id

// BEGIN - django_kea.Lease6.pref_lifetime

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

type inLease6PrefLifetime []interface{}

func (in inLease6PrefLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"pref_lifetime" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PrefLifetimeIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6PrefLifetime(vals),
	)

	return qs
}

type notinLease6PrefLifetime []interface{}

func (in notinLease6PrefLifetime) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"pref_lifetime" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PrefLifetimeNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6PrefLifetime(vals),
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

// DistinctOnPrefLifetime marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnPrefLifetime() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"pref_lifetime"`)

	return qs
}

// END - django_kea.Lease6.pref_lifetime

// BEGIN - django_kea.Lease6.lease_type

// GetLeaseType returns Lease6types
func (l *Lease6) GetLeaseType(ctx context.Context, db models.DBInterface) (*Lease6types, error) {
	if !l.leaseType.Valid {
		return nil, nil
	}

	return Lease6typesQS{}.LeaseTypeEq(l.leaseType.Int32).First(ctx, db)
}

// SetLeaseType sets foreign key pointer to Lease6types
func (l *Lease6) SetLeaseType(ptr *Lease6types) error {
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
func (qs Lease6QS) LeaseTypeEq(v *Lease6types) Lease6QS {
	return qs.filter(`"lease_type" =`, v.LeaseType)
}

// LeaseTypeRawEq filters for leaseType being equal to raw argument
func (qs Lease6QS) LeaseTypeRawEq(v int32) Lease6QS {
	return qs.filter(`"lease_type" =`, v)
}

type inLease6leaseTypeLease6types struct {
	qs Lease6typesQS
}

func (in *inLease6leaseTypeLease6types) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"lease_type" IN (` + s + `)`, p
}

func (qs Lease6QS) LeaseTypeIn(oqs Lease6typesQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&inLease6leaseTypeLease6types{
			qs: oqs,
		},
	)

	return qs
}

type notinLease6leaseTypeLease6types struct {
	qs Lease6typesQS
}

func (nin *notinLease6leaseTypeLease6types) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"lease_type" NOT IN (` + s + `)`, p
}

func (qs Lease6QS) LeaseTypeNotIn(oqs Lease6typesQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&notinLease6leaseTypeLease6types{
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

// DistinctOnLeaseType marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnLeaseType() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"lease_type"`)

	return qs
}

// END - django_kea.Lease6.lease_type

// BEGIN - django_kea.Lease6.iaid

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

type inLease6Iaid []interface{}

func (in inLease6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"iaid" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) IaidIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Iaid(vals),
	)

	return qs
}

type notinLease6Iaid []interface{}

func (in notinLease6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"iaid" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) IaidNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Iaid(vals),
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

// DistinctOnIaid marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnIaid() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"iaid"`)

	return qs
}

// END - django_kea.Lease6.iaid

// BEGIN - django_kea.Lease6.prefix_len

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

type inLease6PrefixLen []interface{}

func (in inLease6PrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"prefix_len" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PrefixLenIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6PrefixLen(vals),
	)

	return qs
}

type notinLease6PrefixLen []interface{}

func (in notinLease6PrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"prefix_len" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PrefixLenNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6PrefixLen(vals),
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

// DistinctOnPrefixLen marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnPrefixLen() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"prefix_len"`)

	return qs
}

// END - django_kea.Lease6.prefix_len

// BEGIN - django_kea.Lease6.fqdn_fwd

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

type inLease6FqdnFwd []interface{}

func (in inLease6FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) FqdnFwdIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6FqdnFwd(vals),
	)

	return qs
}

type notinLease6FqdnFwd []interface{}

func (in notinLease6FqdnFwd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"fqdn_fwd" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) FqdnFwdNotIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6FqdnFwd(vals),
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

// DistinctOnFqdnFwd marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnFqdnFwd() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"fqdn_fwd"`)

	return qs
}

// END - django_kea.Lease6.fqdn_fwd

// BEGIN - django_kea.Lease6.fqdn_rev

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

type inLease6FqdnRev []interface{}

func (in inLease6FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) FqdnRevIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6FqdnRev(vals),
	)

	return qs
}

type notinLease6FqdnRev []interface{}

func (in notinLease6FqdnRev) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"fqdn_rev" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) FqdnRevNotIn(values []bool) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6FqdnRev(vals),
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

// DistinctOnFqdnRev marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnFqdnRev() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"fqdn_rev"`)

	return qs
}

// END - django_kea.Lease6.fqdn_rev

// BEGIN - django_kea.Lease6.hostname

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

type inLease6Hostname []interface{}

func (in inLease6Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hostname" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HostnameIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Hostname(vals),
	)

	return qs
}

type notinLease6Hostname []interface{}

func (in notinLease6Hostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HostnameNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Hostname(vals),
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

// DistinctOnHostname marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnHostname() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hostname"`)

	return qs
}

// END - django_kea.Lease6.hostname

// BEGIN - django_kea.Lease6.state

// GetState returns Leasestate
func (l *Lease6) GetState(ctx context.Context, db models.DBInterface) (*Leasestate, error) {
	if !l.state.Valid {
		return nil, nil
	}

	return LeasestateQS{}.StateEq(l.state.Int64).First(ctx, db)
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

type notinLease6stateLeasestate struct {
	qs LeasestateQS
}

func (nin *notinLease6stateLeasestate) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"state" NOT IN (` + s + `)`, p
}

func (qs Lease6QS) StateNotIn(oqs LeasestateQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&notinLease6stateLeasestate{
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

// DistinctOnState marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnState() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"state"`)

	return qs
}

// END - django_kea.Lease6.state

// BEGIN - django_kea.Lease6.hwaddr

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

type inLease6Hwaddr []interface{}

func (in inLease6Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwaddr" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HwaddrIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Hwaddr(vals),
	)

	return qs
}

type notinLease6Hwaddr []interface{}

func (in notinLease6Hwaddr) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwaddr" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HwaddrNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Hwaddr(vals),
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

// DistinctOnHwaddr marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnHwaddr() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hwaddr"`)

	return qs
}

// END - django_kea.Lease6.hwaddr

// BEGIN - django_kea.Lease6.hwtype

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

type inLease6Hwtype []interface{}

func (in inLease6Hwtype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwtype" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HwtypeIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6Hwtype(vals),
	)

	return qs
}

type notinLease6Hwtype []interface{}

func (in notinLease6Hwtype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hwtype" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) HwtypeNotIn(values []int32) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6Hwtype(vals),
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

// DistinctOnHwtype marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnHwtype() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hwtype"`)

	return qs
}

// END - django_kea.Lease6.hwtype

// BEGIN - django_kea.Lease6.hwaddr_source

// GetHwaddrSource returns Leasehwaddrsource
func (l *Lease6) GetHwaddrSource(ctx context.Context, db models.DBInterface) (*Leasehwaddrsource, error) {
	if !l.hwaddrSource.Valid {
		return nil, nil
	}

	return LeasehwaddrsourceQS{}.HwaddrSourceEq(l.hwaddrSource.Int32).First(ctx, db)
}

// SetHwaddrSource sets foreign key pointer to Leasehwaddrsource
func (l *Lease6) SetHwaddrSource(ptr *Leasehwaddrsource) error {
	if ptr != nil {
		l.hwaddrSource.Int32 = ptr.HwaddrSource
		l.hwaddrSource.Valid = true
	} else {
		l.hwaddrSource.Valid = false
	}

	return nil
}

// GetHwaddrSourceRaw returns Lease6.HwaddrSource
func (l *Lease6) GetHwaddrSourceRaw() sql.NullInt32 {
	return l.hwaddrSource
}

// HwaddrSourceIsNull filters for hwaddrSource being null
func (qs Lease6QS) HwaddrSourceIsNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr_source" IS NULL`,
		},
	)
	return qs
}

// HwaddrSourceIsNotNull filters for hwaddrSource being not null
func (qs Lease6QS) HwaddrSourceIsNotNull() Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hwaddr_source" IS NOT NULL`,
		},
	)
	return qs
}

// HwaddrSourceEq filters for hwaddrSource being equal to argument
func (qs Lease6QS) HwaddrSourceEq(v *Leasehwaddrsource) Lease6QS {
	return qs.filter(`"hwaddr_source" =`, v.HwaddrSource)
}

// HwaddrSourceRawEq filters for hwaddrSource being equal to raw argument
func (qs Lease6QS) HwaddrSourceRawEq(v int32) Lease6QS {
	return qs.filter(`"hwaddr_source" =`, v)
}

type inLease6hwaddrSourceLeasehwaddrsource struct {
	qs LeasehwaddrsourceQS
}

func (in *inLease6hwaddrSourceLeasehwaddrsource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"hwaddr_source" IN (` + s + `)`, p
}

func (qs Lease6QS) HwaddrSourceIn(oqs LeasehwaddrsourceQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&inLease6hwaddrSourceLeasehwaddrsource{
			qs: oqs,
		},
	)

	return qs
}

type notinLease6hwaddrSourceLeasehwaddrsource struct {
	qs LeasehwaddrsourceQS
}

func (nin *notinLease6hwaddrSourceLeasehwaddrsource) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"hwaddr_source" NOT IN (` + s + `)`, p
}

func (qs Lease6QS) HwaddrSourceNotIn(oqs LeasehwaddrsourceQS) Lease6QS {
	qs.condFragments = append(
		qs.condFragments,
		&notinLease6hwaddrSourceLeasehwaddrsource{
			qs: oqs,
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

// DistinctOnHwaddrSource marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnHwaddrSource() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hwaddr_source"`)

	return qs
}

// END - django_kea.Lease6.hwaddr_source

// BEGIN - django_kea.Lease6.user_context

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

type inLease6UserContext []interface{}

func (in inLease6UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) UserContextIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6UserContext(vals),
	)

	return qs
}

type notinLease6UserContext []interface{}

func (in notinLease6UserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) UserContextNotIn(values []string) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6UserContext(vals),
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

// DistinctOnUserContext marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnUserContext() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"user_context"`)

	return qs
}

// END - django_kea.Lease6.user_context

// BEGIN - django_kea.Lease6.pool_id

// PoolIdEq filters for PoolId being equal to argument
func (qs Lease6QS) PoolIdEq(v int64) Lease6QS {
	return qs.filter(`"pool_id" =`, v)
}

// PoolIdNe filters for PoolId being not equal to argument
func (qs Lease6QS) PoolIdNe(v int64) Lease6QS {
	return qs.filter(`"pool_id" <>`, v)
}

// PoolIdLt filters for PoolId being less than argument
func (qs Lease6QS) PoolIdLt(v int64) Lease6QS {
	return qs.filter(`"pool_id" <`, v)
}

// PoolIdLe filters for PoolId being less than or equal to argument
func (qs Lease6QS) PoolIdLe(v int64) Lease6QS {
	return qs.filter(`"pool_id" <=`, v)
}

// PoolIdGt filters for PoolId being greater than argument
func (qs Lease6QS) PoolIdGt(v int64) Lease6QS {
	return qs.filter(`"pool_id" >`, v)
}

// PoolIdGe filters for PoolId being greater than or equal to argument
func (qs Lease6QS) PoolIdGe(v int64) Lease6QS {
	return qs.filter(`"pool_id" >=`, v)
}

type inLease6PoolId []interface{}

func (in inLease6PoolId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"pool_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PoolIdIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inLease6PoolId(vals),
	)

	return qs
}

type notinLease6PoolId []interface{}

func (in notinLease6PoolId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"pool_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Lease6QS) PoolIdNotIn(values []int64) Lease6QS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinLease6PoolId(vals),
	)

	return qs
}

// OrderByPoolId sorts result by PoolId in ascending order
func (qs Lease6QS) OrderByPoolId() Lease6QS {
	qs.order = append(qs.order, `"pool_id"`)

	return qs
}

// OrderByPoolIdDesc sorts result by PoolId in descending order
func (qs Lease6QS) OrderByPoolIdDesc() Lease6QS {
	qs.order = append(qs.order, `"pool_id" DESC`)

	return qs
}

// DistinctOnPoolId marks field in queries to add to DISTINCT ON clause
func (qs Lease6QS) DistinctOnPoolId() Lease6QS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"pool_id"`)

	return qs
}

// END - django_kea.Lease6.pool_id

// OrderByRandom randomizes result
func (qs Lease6QS) OrderByRandom() Lease6QS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Lease6QS) ForUpdate() Lease6QS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs Lease6QS) ForUpdateNowait() Lease6QS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs Lease6QS) ForUpdateSkipLocked() Lease6QS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs Lease6QS) ClearForUpdate() Lease6QS {
	qs.forClause = ""

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

func (qs Lease6QS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"address", "duid", "valid_lifetime", "expire", "subnet_id", "pref_lifetime", "lease_type", "iaid", "prefix_len", "fqdn_fwd", "fqdn_rev", "hostname", "state", "hwaddr", "hwtype", "hwaddr_source", "user_context", "pool_id" FROM "lease6"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Lease6QS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "address" FROM "lease6"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs Lease6QS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"address"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "lease6"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs Lease6QS) All(ctx context.Context, db models.DBInterface) (Lease6List, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret Lease6List
	for rows.Next() {
		obj := Lease6{existsInDB: true}
		if err = rows.Scan(&obj.Address, &obj.Duid, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.PrefLifetime, &obj.leaseType, &obj.Iaid, &obj.PrefixLen, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.Hwaddr, &obj.Hwtype, &obj.hwaddrSource, &obj.UserContext, &obj.PoolId); err != nil {
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
func (qs Lease6QS) First(ctx context.Context, db models.DBInterface) (*Lease6, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Lease6{existsInDB: true}
	err := row.Scan(&obj.Address, &obj.Duid, &obj.ValidLifetime, &obj.Expire, &obj.SubnetId, &obj.PrefLifetime, &obj.leaseType, &obj.Iaid, &obj.PrefixLen, &obj.FqdnFwd, &obj.FqdnRev, &obj.Hostname, &obj.state, &obj.Hwaddr, &obj.Hwtype, &obj.hwaddrSource, &obj.UserContext, &obj.PoolId)
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
func (qs Lease6QS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "lease6"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs Lease6UpdateQS) SetAddress(v net.IP) Lease6UpdateQS {
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

// SetLeaseType sets foreign key pointer to Lease6types
func (uqs Lease6UpdateQS) SetLeaseType(ptr *Lease6types) Lease6UpdateQS {
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

// SetHwaddrSource sets foreign key pointer to Leasehwaddrsource
func (uqs Lease6UpdateQS) SetHwaddrSource(ptr *Leasehwaddrsource) Lease6UpdateQS {
	if ptr != nil {
		return uqs.update(`"hwaddr_source"`, ptr.HwaddrSource)
	}

	return uqs.update(`"hwaddr_source"`, nil)
} // SetUserContext sets UserContext to the given value
func (uqs Lease6UpdateQS) SetUserContext(v sql.NullString) Lease6UpdateQS {
	return uqs.update(`"user_context"`, v)
}

// SetPoolId sets PoolId to the given value
func (uqs Lease6UpdateQS) SetPoolId(v int64) Lease6UpdateQS {
	return uqs.update(`"pool_id"`, v)
}

// Exec executes the update operation
func (uqs Lease6UpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (l *Lease6) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "lease6" ("duid", "valid_lifetime", "expire", "subnet_id", "pref_lifetime", "lease_type", "iaid", "prefix_len", "fqdn_fwd", "fqdn_rev", "hostname", "state", "hwaddr", "hwtype", "hwaddr_source", "user_context", "pool_id", "address") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`, l.Duid, l.ValidLifetime, l.Expire, l.SubnetId, l.PrefLifetime, l.leaseType, l.Iaid, l.PrefixLen, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.Hwaddr, l.Hwtype, l.hwaddrSource, l.UserContext, l.PoolId, l.Address)

	if err != nil {
		return err
	}

	l.existsInDB = true

	return nil
}

// update operation
func (l *Lease6) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "lease6" SET "duid" = $1, "valid_lifetime" = $2, "expire" = $3, "subnet_id" = $4, "pref_lifetime" = $5, "lease_type" = $6, "iaid" = $7, "prefix_len" = $8, "fqdn_fwd" = $9, "fqdn_rev" = $10, "hostname" = $11, "state" = $12, "hwaddr" = $13, "hwtype" = $14, "hwaddr_source" = $15, "user_context" = $16, "pool_id" = $17 WHERE "address" = $18`, l.Duid, l.ValidLifetime, l.Expire, l.SubnetId, l.PrefLifetime, l.leaseType, l.Iaid, l.PrefixLen, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.Hwaddr, l.Hwtype, l.hwaddrSource, l.UserContext, l.PoolId, l.Address)

	return err
}

// Save inserts or updates record
func (l *Lease6) Save(ctx context.Context, db models.DBInterface) error {
	if l.existsInDB {
		return l.update(ctx, db)
	}

	return l.insert(ctx, db)
}

// Delete removes row from database
func (l *Lease6) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "lease6" WHERE "address" = $1`, l.Address)

	l.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (ll Lease6List) Save(ctx context.Context, db models.DBInterface) error {
	var inserts Lease6List

	for _, l := range ll {
		if l.existsInDB {
			if err := l.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, l)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 18*len(inserts))
	offs := 1
	for _, l := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4, offs+5, offs+6, offs+7, offs+8, offs+9, offs+10, offs+11, offs+12, offs+13, offs+14, offs+15, offs+16, offs+17))
		vaa = append(vaa, l.Duid, l.ValidLifetime, l.Expire, l.SubnetId, l.PrefLifetime, l.leaseType, l.Iaid, l.PrefixLen, l.FqdnFwd, l.FqdnRev, l.Hostname, l.state, l.Hwaddr, l.Hwtype, l.hwaddrSource, l.UserContext, l.PoolId, l.Address)
		offs += 18
	}

	qs := `INSERT INTO "lease6" ("duid", "valid_lifetime", "expire", "subnet_id", "pref_lifetime", "lease_type", "iaid", "prefix_len", "fqdn_fwd", "fqdn_rev", "hostname", "state", "hwaddr", "hwtype", "hwaddr_source", "user_context", "pool_id", "address") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, l := range inserts {
		l.existsInDB = true
	}

	return nil

}
