/*
  AUTO-GENERATED file for Django model django_kea.Ipv6Reservation

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

// Ipv6reservation mirrors model django_kea.Ipv6Reservation
type Ipv6reservation struct {
	existsInDB bool

	reservationId int32
	Address       string
	PrefixLen     int32
	Type          int32
	Dhcp6Iaid     sql.NullInt32
	host          int32
}

// Ipv6reservationQS represents a queryset for django_kea.Ipv6Reservation
type Ipv6reservationQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs Ipv6reservationQS) filter(c string, p interface{}) Ipv6reservationQS {
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
func (qs Ipv6reservationQS) Or(exprs ...Ipv6reservationQS) Ipv6reservationQS {
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

// GetReservationId returns Ipv6reservation.ReservationId
func (i *Ipv6reservation) GetReservationId() int32 {
	return i.reservationId
}

// ReservationIdEq filters for reservationId being equal to argument
func (qs Ipv6reservationQS) ReservationIdEq(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" =`, v)
}

// ReservationIdNe filters for reservationId being not equal to argument
func (qs Ipv6reservationQS) ReservationIdNe(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" <>`, v)
}

// ReservationIdLt filters for reservationId being less than argument
func (qs Ipv6reservationQS) ReservationIdLt(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" <`, v)
}

// ReservationIdLe filters for reservationId being less than or equal to argument
func (qs Ipv6reservationQS) ReservationIdLe(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" <=`, v)
}

// ReservationIdGt filters for reservationId being greater than argument
func (qs Ipv6reservationQS) ReservationIdGt(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" >`, v)
}

// ReservationIdGe filters for reservationId being greater than or equal to argument
func (qs Ipv6reservationQS) ReservationIdGe(v int32) Ipv6reservationQS {
	return qs.filter(`"reservation_id" >=`, v)
}

type inIpv6reservationreservationId struct {
	values []interface{}
}

func (in *inIpv6reservationreservationId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"reservation_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) ReservationIdIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationreservationId{
			values: vals,
		},
	)

	return qs
}

type notinIpv6reservationreservationId struct {
	values []interface{}
}

func (in *notinIpv6reservationreservationId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"reservation_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) ReservationIdNotIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationreservationId{
			values: vals,
		},
	)

	return qs
}

// OrderByReservationId sorts result by ReservationId in ascending order
func (qs Ipv6reservationQS) OrderByReservationId() Ipv6reservationQS {
	qs.order = append(qs.order, `"reservation_id"`)

	return qs
}

// OrderByReservationIdDesc sorts result by ReservationId in descending order
func (qs Ipv6reservationQS) OrderByReservationIdDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"reservation_id" DESC`)

	return qs
}

// AddressEq filters for Address being equal to argument
func (qs Ipv6reservationQS) AddressEq(v string) Ipv6reservationQS {
	return qs.filter(`"address" =`, v)
}

// AddressNe filters for Address being not equal to argument
func (qs Ipv6reservationQS) AddressNe(v string) Ipv6reservationQS {
	return qs.filter(`"address" <>`, v)
}

// AddressLt filters for Address being less than argument
func (qs Ipv6reservationQS) AddressLt(v string) Ipv6reservationQS {
	return qs.filter(`"address" <`, v)
}

// AddressLe filters for Address being less than or equal to argument
func (qs Ipv6reservationQS) AddressLe(v string) Ipv6reservationQS {
	return qs.filter(`"address" <=`, v)
}

// AddressGt filters for Address being greater than argument
func (qs Ipv6reservationQS) AddressGt(v string) Ipv6reservationQS {
	return qs.filter(`"address" >`, v)
}

// AddressGe filters for Address being greater than or equal to argument
func (qs Ipv6reservationQS) AddressGe(v string) Ipv6reservationQS {
	return qs.filter(`"address" >=`, v)
}

type inIpv6reservationAddress struct {
	values []interface{}
}

func (in *inIpv6reservationAddress) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) AddressIn(values []string) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationAddress{
			values: vals,
		},
	)

	return qs
}

type notinIpv6reservationAddress struct {
	values []interface{}
}

func (in *notinIpv6reservationAddress) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"address" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) AddressNotIn(values []string) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationAddress{
			values: vals,
		},
	)

	return qs
}

// OrderByAddress sorts result by Address in ascending order
func (qs Ipv6reservationQS) OrderByAddress() Ipv6reservationQS {
	qs.order = append(qs.order, `"address"`)

	return qs
}

// OrderByAddressDesc sorts result by Address in descending order
func (qs Ipv6reservationQS) OrderByAddressDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"address" DESC`)

	return qs
}

// PrefixLenEq filters for PrefixLen being equal to argument
func (qs Ipv6reservationQS) PrefixLenEq(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" =`, v)
}

// PrefixLenNe filters for PrefixLen being not equal to argument
func (qs Ipv6reservationQS) PrefixLenNe(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" <>`, v)
}

// PrefixLenLt filters for PrefixLen being less than argument
func (qs Ipv6reservationQS) PrefixLenLt(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" <`, v)
}

// PrefixLenLe filters for PrefixLen being less than or equal to argument
func (qs Ipv6reservationQS) PrefixLenLe(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" <=`, v)
}

// PrefixLenGt filters for PrefixLen being greater than argument
func (qs Ipv6reservationQS) PrefixLenGt(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" >`, v)
}

// PrefixLenGe filters for PrefixLen being greater than or equal to argument
func (qs Ipv6reservationQS) PrefixLenGe(v int32) Ipv6reservationQS {
	return qs.filter(`"prefix_len" >=`, v)
}

type inIpv6reservationPrefixLen struct {
	values []interface{}
}

func (in *inIpv6reservationPrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"prefix_len" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) PrefixLenIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationPrefixLen{
			values: vals,
		},
	)

	return qs
}

type notinIpv6reservationPrefixLen struct {
	values []interface{}
}

func (in *notinIpv6reservationPrefixLen) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"prefix_len" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) PrefixLenNotIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationPrefixLen{
			values: vals,
		},
	)

	return qs
}

// OrderByPrefixLen sorts result by PrefixLen in ascending order
func (qs Ipv6reservationQS) OrderByPrefixLen() Ipv6reservationQS {
	qs.order = append(qs.order, `"prefix_len"`)

	return qs
}

// OrderByPrefixLenDesc sorts result by PrefixLen in descending order
func (qs Ipv6reservationQS) OrderByPrefixLenDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"prefix_len" DESC`)

	return qs
}

// TypeEq filters for Type being equal to argument
func (qs Ipv6reservationQS) TypeEq(v int32) Ipv6reservationQS {
	return qs.filter(`"type" =`, v)
}

// TypeNe filters for Type being not equal to argument
func (qs Ipv6reservationQS) TypeNe(v int32) Ipv6reservationQS {
	return qs.filter(`"type" <>`, v)
}

// TypeLt filters for Type being less than argument
func (qs Ipv6reservationQS) TypeLt(v int32) Ipv6reservationQS {
	return qs.filter(`"type" <`, v)
}

// TypeLe filters for Type being less than or equal to argument
func (qs Ipv6reservationQS) TypeLe(v int32) Ipv6reservationQS {
	return qs.filter(`"type" <=`, v)
}

// TypeGt filters for Type being greater than argument
func (qs Ipv6reservationQS) TypeGt(v int32) Ipv6reservationQS {
	return qs.filter(`"type" >`, v)
}

// TypeGe filters for Type being greater than or equal to argument
func (qs Ipv6reservationQS) TypeGe(v int32) Ipv6reservationQS {
	return qs.filter(`"type" >=`, v)
}

type inIpv6reservationType struct {
	values []interface{}
}

func (in *inIpv6reservationType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"type" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) TypeIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationType{
			values: vals,
		},
	)

	return qs
}

type notinIpv6reservationType struct {
	values []interface{}
}

func (in *notinIpv6reservationType) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"type" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) TypeNotIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationType{
			values: vals,
		},
	)

	return qs
}

// OrderByType sorts result by Type in ascending order
func (qs Ipv6reservationQS) OrderByType() Ipv6reservationQS {
	qs.order = append(qs.order, `"type"`)

	return qs
}

// OrderByTypeDesc sorts result by Type in descending order
func (qs Ipv6reservationQS) OrderByTypeDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"type" DESC`)

	return qs
}

// Dhcp6IaidIsNull filters for Dhcp6Iaid being null
func (qs Ipv6reservationQS) Dhcp6IaidIsNull() Ipv6reservationQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_iaid" IS NULL`,
		},
	)
	return qs
}

// Dhcp6IaidIsNotNull filters for Dhcp6Iaid being not null
func (qs Ipv6reservationQS) Dhcp6IaidIsNotNull() Ipv6reservationQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_iaid" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6IaidEq filters for Dhcp6Iaid being equal to argument
func (qs Ipv6reservationQS) Dhcp6IaidEq(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" =`, v)
}

// Dhcp6IaidNe filters for Dhcp6Iaid being not equal to argument
func (qs Ipv6reservationQS) Dhcp6IaidNe(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" <>`, v)
}

// Dhcp6IaidLt filters for Dhcp6Iaid being less than argument
func (qs Ipv6reservationQS) Dhcp6IaidLt(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" <`, v)
}

// Dhcp6IaidLe filters for Dhcp6Iaid being less than or equal to argument
func (qs Ipv6reservationQS) Dhcp6IaidLe(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" <=`, v)
}

// Dhcp6IaidGt filters for Dhcp6Iaid being greater than argument
func (qs Ipv6reservationQS) Dhcp6IaidGt(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" >`, v)
}

// Dhcp6IaidGe filters for Dhcp6Iaid being greater than or equal to argument
func (qs Ipv6reservationQS) Dhcp6IaidGe(v int32) Ipv6reservationQS {
	return qs.filter(`"dhcp6_iaid" >=`, v)
}

type inIpv6reservationDhcp6Iaid struct {
	values []interface{}
}

func (in *inIpv6reservationDhcp6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_iaid" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) Dhcp6IaidIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationDhcp6Iaid{
			values: vals,
		},
	)

	return qs
}

type notinIpv6reservationDhcp6Iaid struct {
	values []interface{}
}

func (in *notinIpv6reservationDhcp6Iaid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_iaid" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs Ipv6reservationQS) Dhcp6IaidNotIn(values []int32) Ipv6reservationQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinIpv6reservationDhcp6Iaid{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp6Iaid sorts result by Dhcp6Iaid in ascending order
func (qs Ipv6reservationQS) OrderByDhcp6Iaid() Ipv6reservationQS {
	qs.order = append(qs.order, `"dhcp6_iaid"`)

	return qs
}

// OrderByDhcp6IaidDesc sorts result by Dhcp6Iaid in descending order
func (qs Ipv6reservationQS) OrderByDhcp6IaidDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"dhcp6_iaid" DESC`)

	return qs
}

// GetHost returns Host
func (i *Ipv6reservation) GetHost(db models.DBInterface) (*Host, error) {
	return HostQS{}.HostIdEq(i.host).First(db)
}

// SetHost sets foreign key pointer to Host
func (i *Ipv6reservation) SetHost(ptr *Host) error {
	if ptr != nil {
		i.host = ptr.GetHostId()
	} else {
		return fmt.Errorf("Ipv6reservation.SetHost: non-null field received null value")
	}

	return nil
}

// GetHostRaw returns Ipv6reservation.Host
func (i *Ipv6reservation) GetHostRaw() int32 {
	return i.host
}

// HostEq filters for host being equal to argument
func (qs Ipv6reservationQS) HostEq(v *Host) Ipv6reservationQS {
	return qs.filter(`"host_id" =`, v.GetHostId())
}

// HostRawEq filters for host being equal to raw argument
func (qs Ipv6reservationQS) HostRawEq(v int32) Ipv6reservationQS {
	return qs.filter(`"host_id" =`, v)
}

type inIpv6reservationhostHost struct {
	qs HostQS
}

func (in *inIpv6reservationhostHost) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"host_id" IN (` + s + `)`, p
}

func (qs Ipv6reservationQS) HostIn(oqs HostQS) Ipv6reservationQS {
	qs.condFragments = append(
		qs.condFragments,
		&inIpv6reservationhostHost{
			qs: oqs,
		},
	)

	return qs
}

// OrderByHost sorts result by Host in ascending order
func (qs Ipv6reservationQS) OrderByHost() Ipv6reservationQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostDesc sorts result by Host in descending order
func (qs Ipv6reservationQS) OrderByHostDesc() Ipv6reservationQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs Ipv6reservationQS) ForUpdate() Ipv6reservationQS {
	qs.forUpdate = true

	return qs
}

func (qs Ipv6reservationQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs Ipv6reservationQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs Ipv6reservationQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "reservation_id", "address", "prefix_len", "type", "dhcp6_iaid", "host_id" FROM "ipv6_reservations"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs Ipv6reservationQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "reservation_id" FROM "ipv6_reservations"` + s, p
}

// All returns all rows matching queryset filters
func (qs Ipv6reservationQS) All(db models.DBInterface) ([]*Ipv6reservation, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Ipv6reservation
	for rows.Next() {
		obj := Ipv6reservation{existsInDB: true}
		if err = rows.Scan(&obj.reservationId, &obj.Address, &obj.PrefixLen, &obj.Type, &obj.Dhcp6Iaid, &obj.host); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs Ipv6reservationQS) First(db models.DBInterface) (*Ipv6reservation, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Ipv6reservation{existsInDB: true}
	err := row.Scan(&obj.reservationId, &obj.Address, &obj.PrefixLen, &obj.Type, &obj.Dhcp6Iaid, &obj.host)
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
func (qs Ipv6reservationQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "ipv6_reservations"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs Ipv6reservationQS) Update() Ipv6reservationUpdateQS {
	return Ipv6reservationUpdateQS{condFragments: qs.condFragments}
}

// Ipv6reservationUpdateQS represents an updated queryset for django_kea.Ipv6Reservation
type Ipv6reservationUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs Ipv6reservationUpdateQS) update(c string, v interface{}) Ipv6reservationUpdateQS {
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
func (uqs Ipv6reservationUpdateQS) SetReservationId(v int32) Ipv6reservationUpdateQS {
	return uqs.update(`"reservation_id"`, v)
}

// SetAddress sets Address to the given value
func (uqs Ipv6reservationUpdateQS) SetAddress(v string) Ipv6reservationUpdateQS {
	return uqs.update(`"address"`, v)
}

// SetPrefixLen sets PrefixLen to the given value
func (uqs Ipv6reservationUpdateQS) SetPrefixLen(v int32) Ipv6reservationUpdateQS {
	return uqs.update(`"prefix_len"`, v)
}

// SetType sets Type to the given value
func (uqs Ipv6reservationUpdateQS) SetType(v int32) Ipv6reservationUpdateQS {
	return uqs.update(`"type"`, v)
}

// SetDhcp6Iaid sets Dhcp6Iaid to the given value
func (uqs Ipv6reservationUpdateQS) SetDhcp6Iaid(v sql.NullInt32) Ipv6reservationUpdateQS {
	return uqs.update(`"dhcp6_iaid"`, v)
}

// SetHost sets foreign key pointer to Host
func (uqs Ipv6reservationUpdateQS) SetHost(ptr *Host) Ipv6reservationUpdateQS {
	if ptr != nil {
		return uqs.update(`"host_id"`, ptr.GetHostId())
	}

	return uqs.update(`"host_id"`, nil)
} // Exec executes the update operation
func (uqs Ipv6reservationUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := Ipv6reservationQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "ipv6_reservations" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (i *Ipv6reservation) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "ipv6_reservations" ("address", "prefix_len", "type", "dhcp6_iaid", "host_id") VALUES ($1, $2, $3, $4, $5) RETURNING "reservation_id"`, i.Address, i.PrefixLen, i.Type, i.Dhcp6Iaid, i.host)

	if err := row.Scan(&i.reservationId); err != nil {
		return err
	}

	i.existsInDB = true

	return nil
}

// update operation
func (i *Ipv6reservation) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "ipv6_reservations" SET "address" = $1, "prefix_len" = $2, "type" = $3, "dhcp6_iaid" = $4, "host_id" = $5 WHERE "reservation_id" = $6`, i.Address, i.PrefixLen, i.Type, i.Dhcp6Iaid, i.host, i.reservationId)

	return err
}

// Save inserts or updates record
func (i *Ipv6reservation) Save(db models.DBInterface) error {
	if i.existsInDB {
		return i.update(db)
	}

	return i.insert(db)
}

// Delete removes row from database
func (i *Ipv6reservation) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "ipv6_reservations" WHERE "reservation_id" = $1`, i.reservationId)

	i.existsInDB = false

	return err
}
