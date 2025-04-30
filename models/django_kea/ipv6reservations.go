// Code generated for Django model django_kea.Ipv6Reservations. DO NOT EDIT.

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

	"github.com/jackc/pgx/v5"

	"github.com/euronetzrt/django-kea/models"
)

// Ipv6reservations mirrors model django_kea.Ipv6Reservations
type Ipv6reservations struct {
	existsInDB bool

	reservationId int32
	Address       net.IP
	PrefixLen     int32
	Type          int32
	Dhcp6Iaid     sql.NullInt32
	host          int32
}

// Ipv6reservationsList is a list of Ipv6reservations
type Ipv6reservationsList []*Ipv6reservations

// Ipv6reservationsQS represents a queryset for django_kea.Ipv6Reservations
type Ipv6reservationsQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs Ipv6reservationsQS) filter(c string, p interface{}) Ipv6reservationsQS {
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
func (qs Ipv6reservationsQS) Or(exprs ...Ipv6reservationsQS) Ipv6reservationsQS {
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

// BEGIN - django_kea.Ipv6Reservations.reservation_id

// GetReservationId returns Ipv6reservations.ReservationId
func (i *Ipv6reservations) GetReservationId() int32 {
	return i.reservationId
}

// ReservationIdEq filters for reservationId being equal to argument
func (qs Ipv6reservationsQS) ReservationIdEq(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" =`, v)
}

// ReservationIdNe filters for reservationId being not equal to argument
func (qs Ipv6reservationsQS) ReservationIdNe(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" <>`, v)
}

// ReservationIdLt filters for reservationId being less than argument
func (qs Ipv6reservationsQS) ReservationIdLt(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" <`, v)
}

// ReservationIdLe filters for reservationId being less than or equal to argument
func (qs Ipv6reservationsQS) ReservationIdLe(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" <=`, v)
}

// ReservationIdGt filters for reservationId being greater than argument
func (qs Ipv6reservationsQS) ReservationIdGt(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" >`, v)
}

// ReservationIdGe filters for reservationId being greater than or equal to argument
func (qs Ipv6reservationsQS) ReservationIdGe(v int32) Ipv6reservationsQS {
	return qs.filter(`"reservation_id" >=`, v)
}

type inIpv6reservationsreservationId []interface{}

func (in inIpv6reservationsreservationId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"reservation_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) ReservationIdIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inIpv6reservationsreservationId(vals),
	)

	return qs
}

type notinIpv6reservationsreservationId []interface{}

func (in notinIpv6reservationsreservationId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"reservation_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) ReservationIdNotIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinIpv6reservationsreservationId(vals),
	)

	return qs
}

// OrderByReservationId sorts result by ReservationId in ascending order
func (qs Ipv6reservationsQS) OrderByReservationId() Ipv6reservationsQS {
	qs.order = append(qs.order, `"reservation_id"`)

	return qs
}

// OrderByReservationIdDesc sorts result by ReservationId in descending order
func (qs Ipv6reservationsQS) OrderByReservationIdDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"reservation_id" DESC`)

	return qs
}

// DistinctOnReservationId marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnReservationId() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"reservation_id"`)

	return qs
}

// END - django_kea.Ipv6Reservations.reservation_id

// BEGIN - django_kea.Ipv6Reservations.address

// AddressEq filters for Address being equal to argument
func (qs Ipv6reservationsQS) AddressEq(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" =`, v)
}

// AddressNe filters for Address being not equal to argument
func (qs Ipv6reservationsQS) AddressNe(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" <>`, v)
}

// AddressLt filters for Address being less than argument
func (qs Ipv6reservationsQS) AddressLt(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" <`, v)
}

// AddressLe filters for Address being less than or equal to argument
func (qs Ipv6reservationsQS) AddressLe(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" <=`, v)
}

// AddressGt filters for Address being greater than argument
func (qs Ipv6reservationsQS) AddressGt(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" >`, v)
}

// AddressGe filters for Address being greater than or equal to argument
func (qs Ipv6reservationsQS) AddressGe(v net.IP) Ipv6reservationsQS {
	return qs.filter(`"address" >=`, v)
}

type inIpv6reservationsAddress []interface{}

func (in inIpv6reservationsAddress) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"address" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) AddressIn(values []net.IP) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inIpv6reservationsAddress(vals),
	)

	return qs
}

type notinIpv6reservationsAddress []interface{}

func (in notinIpv6reservationsAddress) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"address" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) AddressNotIn(values []net.IP) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinIpv6reservationsAddress(vals),
	)

	return qs
}

// OrderByAddress sorts result by Address in ascending order
func (qs Ipv6reservationsQS) OrderByAddress() Ipv6reservationsQS {
	qs.order = append(qs.order, `"address"`)

	return qs
}

// OrderByAddressDesc sorts result by Address in descending order
func (qs Ipv6reservationsQS) OrderByAddressDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"address" DESC`)

	return qs
}

// DistinctOnAddress marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnAddress() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"address"`)

	return qs
}

// END - django_kea.Ipv6Reservations.address

// BEGIN - django_kea.Ipv6Reservations.prefix_len

// PrefixLenEq filters for PrefixLen being equal to argument
func (qs Ipv6reservationsQS) PrefixLenEq(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" =`, v)
}

// PrefixLenNe filters for PrefixLen being not equal to argument
func (qs Ipv6reservationsQS) PrefixLenNe(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" <>`, v)
}

// PrefixLenLt filters for PrefixLen being less than argument
func (qs Ipv6reservationsQS) PrefixLenLt(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" <`, v)
}

// PrefixLenLe filters for PrefixLen being less than or equal to argument
func (qs Ipv6reservationsQS) PrefixLenLe(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" <=`, v)
}

// PrefixLenGt filters for PrefixLen being greater than argument
func (qs Ipv6reservationsQS) PrefixLenGt(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" >`, v)
}

// PrefixLenGe filters for PrefixLen being greater than or equal to argument
func (qs Ipv6reservationsQS) PrefixLenGe(v int32) Ipv6reservationsQS {
	return qs.filter(`"prefix_len" >=`, v)
}

type inIpv6reservationsPrefixLen []interface{}

func (in inIpv6reservationsPrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"prefix_len" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) PrefixLenIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inIpv6reservationsPrefixLen(vals),
	)

	return qs
}

type notinIpv6reservationsPrefixLen []interface{}

func (in notinIpv6reservationsPrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"prefix_len" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) PrefixLenNotIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinIpv6reservationsPrefixLen(vals),
	)

	return qs
}

// OrderByPrefixLen sorts result by PrefixLen in ascending order
func (qs Ipv6reservationsQS) OrderByPrefixLen() Ipv6reservationsQS {
	qs.order = append(qs.order, `"prefix_len"`)

	return qs
}

// OrderByPrefixLenDesc sorts result by PrefixLen in descending order
func (qs Ipv6reservationsQS) OrderByPrefixLenDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"prefix_len" DESC`)

	return qs
}

// DistinctOnPrefixLen marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnPrefixLen() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"prefix_len"`)

	return qs
}

// END - django_kea.Ipv6Reservations.prefix_len

// BEGIN - django_kea.Ipv6Reservations.type

// TypeEq filters for Type being equal to argument
func (qs Ipv6reservationsQS) TypeEq(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" =`, v)
}

// TypeNe filters for Type being not equal to argument
func (qs Ipv6reservationsQS) TypeNe(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" <>`, v)
}

// TypeLt filters for Type being less than argument
func (qs Ipv6reservationsQS) TypeLt(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" <`, v)
}

// TypeLe filters for Type being less than or equal to argument
func (qs Ipv6reservationsQS) TypeLe(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" <=`, v)
}

// TypeGt filters for Type being greater than argument
func (qs Ipv6reservationsQS) TypeGt(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" >`, v)
}

// TypeGe filters for Type being greater than or equal to argument
func (qs Ipv6reservationsQS) TypeGe(v int32) Ipv6reservationsQS {
	return qs.filter(`"type" >=`, v)
}

type inIpv6reservationsType []interface{}

func (in inIpv6reservationsType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"type" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) TypeIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inIpv6reservationsType(vals),
	)

	return qs
}

type notinIpv6reservationsType []interface{}

func (in notinIpv6reservationsType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"type" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) TypeNotIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinIpv6reservationsType(vals),
	)

	return qs
}

// OrderByType sorts result by Type in ascending order
func (qs Ipv6reservationsQS) OrderByType() Ipv6reservationsQS {
	qs.order = append(qs.order, `"type"`)

	return qs
}

// OrderByTypeDesc sorts result by Type in descending order
func (qs Ipv6reservationsQS) OrderByTypeDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"type" DESC`)

	return qs
}

// DistinctOnType marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnType() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"type"`)

	return qs
}

// END - django_kea.Ipv6Reservations.type

// BEGIN - django_kea.Ipv6Reservations.dhcp6_iaid

// Dhcp6IaidIsNull filters for Dhcp6Iaid being null
func (qs Ipv6reservationsQS) Dhcp6IaidIsNull() Ipv6reservationsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_iaid" IS NULL`,
		},
	)
	return qs
}

// Dhcp6IaidIsNotNull filters for Dhcp6Iaid being not null
func (qs Ipv6reservationsQS) Dhcp6IaidIsNotNull() Ipv6reservationsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_iaid" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6IaidEq filters for Dhcp6Iaid being equal to argument
func (qs Ipv6reservationsQS) Dhcp6IaidEq(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" =`, v)
}

// Dhcp6IaidNe filters for Dhcp6Iaid being not equal to argument
func (qs Ipv6reservationsQS) Dhcp6IaidNe(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" <>`, v)
}

// Dhcp6IaidLt filters for Dhcp6Iaid being less than argument
func (qs Ipv6reservationsQS) Dhcp6IaidLt(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" <`, v)
}

// Dhcp6IaidLe filters for Dhcp6Iaid being less than or equal to argument
func (qs Ipv6reservationsQS) Dhcp6IaidLe(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" <=`, v)
}

// Dhcp6IaidGt filters for Dhcp6Iaid being greater than argument
func (qs Ipv6reservationsQS) Dhcp6IaidGt(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" >`, v)
}

// Dhcp6IaidGe filters for Dhcp6Iaid being greater than or equal to argument
func (qs Ipv6reservationsQS) Dhcp6IaidGe(v int32) Ipv6reservationsQS {
	return qs.filter(`"dhcp6_iaid" >=`, v)
}

type inIpv6reservationsDhcp6Iaid []interface{}

func (in inIpv6reservationsDhcp6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_iaid" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) Dhcp6IaidIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inIpv6reservationsDhcp6Iaid(vals),
	)

	return qs
}

type notinIpv6reservationsDhcp6Iaid []interface{}

func (in notinIpv6reservationsDhcp6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_iaid" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs Ipv6reservationsQS) Dhcp6IaidNotIn(values []int32) Ipv6reservationsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinIpv6reservationsDhcp6Iaid(vals),
	)

	return qs
}

// OrderByDhcp6Iaid sorts result by Dhcp6Iaid in ascending order
func (qs Ipv6reservationsQS) OrderByDhcp6Iaid() Ipv6reservationsQS {
	qs.order = append(qs.order, `"dhcp6_iaid"`)

	return qs
}

// OrderByDhcp6IaidDesc sorts result by Dhcp6Iaid in descending order
func (qs Ipv6reservationsQS) OrderByDhcp6IaidDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"dhcp6_iaid" DESC`)

	return qs
}

// DistinctOnDhcp6Iaid marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnDhcp6Iaid() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp6_iaid"`)

	return qs
}

// END - django_kea.Ipv6Reservations.dhcp6_iaid

// BEGIN - django_kea.Ipv6Reservations.host

// GetHost returns Hosts
func (i *Ipv6reservations) GetHost(ctx context.Context, db models.DBInterface) (*Hosts, error) {
	return HostsQS{}.HostIdEq(i.host).First(ctx, db)
}

// SetHost sets foreign key pointer to Hosts
func (i *Ipv6reservations) SetHost(ptr *Hosts) error {
	if ptr != nil {
		i.host = ptr.GetHostId()
	} else {
		return fmt.Errorf("Ipv6reservations.SetHost: non-null field received null value")
	}

	return nil
}

// GetHostRaw returns Ipv6reservations.Host
func (i *Ipv6reservations) GetHostRaw() int32 {
	return i.host
}

// HostEq filters for host being equal to argument
func (qs Ipv6reservationsQS) HostEq(v *Hosts) Ipv6reservationsQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Ipv6reservationsQS) HostRawEq(v int32) Ipv6reservationsQS {
	return qs.filter(`"host_id" =`, v)
}

type inIpv6reservationshostHosts struct {
	qs HostsQS
}

func (in *inIpv6reservationshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Ipv6reservationsQS) HostIn(oqs HostsQS) Ipv6reservationsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationshostHosts{
			qs: oqs,
		},
	)

	return qs
}

type notinIpv6reservationshostHosts struct {
	qs HostsQS
}

func (nin *notinIpv6reservationshostHosts) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"host_id" NOT IN (` + s + `)`, p
}

func (qs Ipv6reservationsQS) HostNotIn(oqs HostsQS) Ipv6reservationsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationshostHosts{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Ipv6reservationsQS) OrderByHost() Ipv6reservationsQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Ipv6reservationsQS) OrderByHostDesc() Ipv6reservationsQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// DistinctOnHost marks field in queries to add to DISTINCT ON clause
func (qs Ipv6reservationsQS) DistinctOnHost() Ipv6reservationsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"host_id"`)

	return qs
}

// END - django_kea.Ipv6Reservations.host

// OrderByRandom randomizes result
func (qs Ipv6reservationsQS) OrderByRandom() Ipv6reservationsQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Ipv6reservationsQS) ForUpdate() Ipv6reservationsQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs Ipv6reservationsQS) ForUpdateNowait() Ipv6reservationsQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs Ipv6reservationsQS) ForUpdateSkipLocked() Ipv6reservationsQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs Ipv6reservationsQS) ClearForUpdate() Ipv6reservationsQS {
	qs.forClause = ""

	return qs
}

func (qs Ipv6reservationsQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Ipv6reservationsQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Ipv6reservationsQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"reservation_id", "address", "prefix_len", "type", "dhcp6_iaid", "host_id" FROM "ipv6_reservations"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Ipv6reservationsQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "reservation_id" FROM "ipv6_reservations"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs Ipv6reservationsQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"reservation_id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "ipv6_reservations"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs Ipv6reservationsQS) All(ctx context.Context, db models.DBInterface) (Ipv6reservationsList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret Ipv6reservationsList
	for rows.Next() {
		obj := Ipv6reservations{existsInDB: true}
		if err = rows.Scan(&obj.reservationId, &obj.Address, &obj.PrefixLen, &obj.Type, &obj.Dhcp6Iaid, &obj.host); err != nil {
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
func (qs Ipv6reservationsQS) First(ctx context.Context, db models.DBInterface) (*Ipv6reservations, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Ipv6reservations{existsInDB: true}
	err := row.Scan(&obj.reservationId, &obj.Address, &obj.PrefixLen, &obj.Type, &obj.Dhcp6Iaid, &obj.host)
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
func (qs Ipv6reservationsQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "ipv6_reservations"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Ipv6reservationsQS) Update() Ipv6reservationsUpdateQS {
	return Ipv6reservationsUpdateQS{condFragments: qs.condFragments}
}

// Ipv6reservationsUpdateQS represents an updated queryset for django_kea.Ipv6Reservations
type Ipv6reservationsUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Ipv6reservationsUpdateQS) update(c string, v interface{}) Ipv6reservationsUpdateQS {
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

// SetReservationId sets ReservationId to the given value
func (uqs Ipv6reservationsUpdateQS) SetReservationId(v int32) Ipv6reservationsUpdateQS {
	return uqs.update(`"reservation_id"`, v)
}

// SetAddress sets Address to the given value
func (uqs Ipv6reservationsUpdateQS) SetAddress(v net.IP) Ipv6reservationsUpdateQS {
	return uqs.update(`"address"`, v)
}

// SetPrefixLen sets PrefixLen to the given value
func (uqs Ipv6reservationsUpdateQS) SetPrefixLen(v int32) Ipv6reservationsUpdateQS {
	return uqs.update(`"prefix_len"`, v)
}

// SetType sets Type to the given value
func (uqs Ipv6reservationsUpdateQS) SetType(v int32) Ipv6reservationsUpdateQS {
	return uqs.update(`"type"`, v)
}

// SetDhcp6Iaid sets Dhcp6Iaid to the given value
func (uqs Ipv6reservationsUpdateQS) SetDhcp6Iaid(v sql.NullInt32) Ipv6reservationsUpdateQS {
	return uqs.update(`"dhcp6_iaid"`, v)
}

// SetHost sets foreign key pointer to Hosts
func (uqs Ipv6reservationsUpdateQS) SetHost(ptr *Hosts) Ipv6reservationsUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // Exec executes the update operation
func (uqs Ipv6reservationsUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := Ipv6reservationsQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "ipv6_reservations" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (i *Ipv6reservations) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "ipv6_reservations" ("address", "prefix_len", "type", "dhcp6_iaid", "host_id") VALUES ($1, $2, $3, $4, $5) RETURNING "reservation_id"`, i.Address, i.PrefixLen, i.Type, i.Dhcp6Iaid, i.host)

	if err := row.Scan(&i.reservationId); err != nil {
		return err
	}

	i.existsInDB = true

	return nil
}

// update operation
func (i *Ipv6reservations) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "ipv6_reservations" SET "address" = $1, "prefix_len" = $2, "type" = $3, "dhcp6_iaid" = $4, "host_id" = $5 WHERE "reservation_id" = $6`, i.Address, i.PrefixLen, i.Type, i.Dhcp6Iaid, i.host, i.reservationId)

	return err
}

// Save inserts or updates record
func (i *Ipv6reservations) Save(ctx context.Context, db models.DBInterface) error {
	if i.existsInDB {
		return i.update(ctx, db)
	}

	return i.insert(ctx, db)
}

// Delete removes row from database
func (i *Ipv6reservations) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "ipv6_reservations" WHERE "reservation_id" = $1`, i.reservationId)

	i.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (il Ipv6reservationsList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts Ipv6reservationsList

	for _, i := range il {
		if i.existsInDB {
			if err := i.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, i)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 5*len(inserts))
	offs := 1
	for _, i := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4))
		vaa = append(vaa, i.Address, i.PrefixLen, i.Type, i.Dhcp6Iaid, i.host)
		offs += 5
	}

	qs := `INSERT INTO "ipv6_reservations" ("address", "prefix_len", "type", "dhcp6_iaid", "host_id") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "reservation_id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, i := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&i.reservationId); err != nil {
			return err
		}

		i.existsInDB = true
	}

	return rows.Err()
}
