// Code generated for Django model django_kea.Hosts. DO NOT EDIT.

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

// Hosts mirrors model django_kea.Hosts
type Hosts struct {
	existsInDB bool

	hostId              int32
	DhcpIdentifier      string
	dhcpIdentifierType  int32
	Dhcp4SubnetId       sql.NullInt64
	Dhcp6SubnetId       sql.NullInt64
	Ipv4Address         sql.NullString
	Hostname            sql.NullString
	Dhcp4ClientClasses  sql.NullString
	Dhcp6ClientClasses  sql.NullString
	Dhcp4NextServer     sql.NullString
	Dhcp4ServerHostname sql.NullString
	Dhcp4BootFileName   sql.NullString
	UserContext         sql.NullString
	AuthKey             sql.NullString
}

// HostsList is a list of Hosts
type HostsList []*Hosts

// HostsQS represents a queryset for django_kea.Hosts
type HostsQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs HostsQS) filter(c string, p interface{}) HostsQS {
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
func (qs HostsQS) Or(exprs ...HostsQS) HostsQS {
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

// BEGIN - django_kea.Hosts.host_id

// GetHostId returns Hosts.HostId
func (h *Hosts) GetHostId() int32 {
	return h.hostId
}

// HostIdEq filters for hostId being equal to argument
func (qs HostsQS) HostIdEq(v int32) HostsQS {
	return qs.filter(`"host_id" =`, v)
}

// HostIdNe filters for hostId being not equal to argument
func (qs HostsQS) HostIdNe(v int32) HostsQS {
	return qs.filter(`"host_id" <>`, v)
}

// HostIdLt filters for hostId being less than argument
func (qs HostsQS) HostIdLt(v int32) HostsQS {
	return qs.filter(`"host_id" <`, v)
}

// HostIdLe filters for hostId being less than or equal to argument
func (qs HostsQS) HostIdLe(v int32) HostsQS {
	return qs.filter(`"host_id" <=`, v)
}

// HostIdGt filters for hostId being greater than argument
func (qs HostsQS) HostIdGt(v int32) HostsQS {
	return qs.filter(`"host_id" >`, v)
}

// HostIdGe filters for hostId being greater than or equal to argument
func (qs HostsQS) HostIdGe(v int32) HostsQS {
	return qs.filter(`"host_id" >=`, v)
}

type inHostshostId []interface{}

func (in inHostshostId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"host_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) HostIdIn(values []int32) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostshostId(vals),
	)

	return qs
}

type notinHostshostId []interface{}

func (in notinHostshostId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"host_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) HostIdNotIn(values []int32) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostshostId(vals),
	)

	return qs
}

// OrderByHostId sorts result by HostId in ascending order
func (qs HostsQS) OrderByHostId() HostsQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostIdDesc sorts result by HostId in descending order
func (qs HostsQS) OrderByHostIdDesc() HostsQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// DistinctOnHostId marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnHostId() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"host_id"`)

	return qs
}

// END - django_kea.Hosts.host_id

// BEGIN - django_kea.Hosts.dhcp_identifier

// DhcpIdentifierEq filters for DhcpIdentifier being equal to argument
func (qs HostsQS) DhcpIdentifierEq(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" =`, v)
}

// DhcpIdentifierNe filters for DhcpIdentifier being not equal to argument
func (qs HostsQS) DhcpIdentifierNe(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" <>`, v)
}

// DhcpIdentifierLt filters for DhcpIdentifier being less than argument
func (qs HostsQS) DhcpIdentifierLt(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" <`, v)
}

// DhcpIdentifierLe filters for DhcpIdentifier being less than or equal to argument
func (qs HostsQS) DhcpIdentifierLe(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" <=`, v)
}

// DhcpIdentifierGt filters for DhcpIdentifier being greater than argument
func (qs HostsQS) DhcpIdentifierGt(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" >`, v)
}

// DhcpIdentifierGe filters for DhcpIdentifier being greater than or equal to argument
func (qs HostsQS) DhcpIdentifierGe(v string) HostsQS {
	return qs.filter(`"dhcp_identifier" >=`, v)
}

type inHostsDhcpIdentifier []interface{}

func (in inHostsDhcpIdentifier) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_identifier" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) DhcpIdentifierIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcpIdentifier(vals),
	)

	return qs
}

type notinHostsDhcpIdentifier []interface{}

func (in notinHostsDhcpIdentifier) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp_identifier" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) DhcpIdentifierNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcpIdentifier(vals),
	)

	return qs
}

// OrderByDhcpIdentifier sorts result by DhcpIdentifier in ascending order
func (qs HostsQS) OrderByDhcpIdentifier() HostsQS {
	qs.order = append(qs.order, `"dhcp_identifier"`)

	return qs
}

// OrderByDhcpIdentifierDesc sorts result by DhcpIdentifier in descending order
func (qs HostsQS) OrderByDhcpIdentifierDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp_identifier" DESC`)

	return qs
}

// DistinctOnDhcpIdentifier marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcpIdentifier() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp_identifier"`)

	return qs
}

// END - django_kea.Hosts.dhcp_identifier

// BEGIN - django_kea.Hosts.dhcp_identifier_type

// GetDhcpIdentifierType returns Hostidentifiertype
func (h *Hosts) GetDhcpIdentifierType(ctx context.Context, db models.DBInterface) (*Hostidentifiertype, error) {
	return HostidentifiertypeQS{}.TypeEq(h.dhcpIdentifierType).First(ctx, db)
}

// SetDhcpIdentifierType sets foreign key pointer to Hostidentifiertype
func (h *Hosts) SetDhcpIdentifierType(ptr *Hostidentifiertype) error {
	if ptr != nil {
		h.dhcpIdentifierType = ptr.Type
	} else {
		return fmt.Errorf("Hosts.SetDhcpIdentifierType: non-null field received null value")
	}

	return nil
}

// GetDhcpIdentifierTypeRaw returns Hosts.DhcpIdentifierType
func (h *Hosts) GetDhcpIdentifierTypeRaw() int32 {
	return h.dhcpIdentifierType
}

// DhcpIdentifierTypeEq filters for dhcpIdentifierType being equal to argument
func (qs HostsQS) DhcpIdentifierTypeEq(v *Hostidentifiertype) HostsQS {
	return qs.filter(`"dhcp_identifier_type" =`, v.Type)
}

// DhcpIdentifierTypeRawEq filters for dhcpIdentifierType being equal to raw argument
func (qs HostsQS) DhcpIdentifierTypeRawEq(v int32) HostsQS {
	return qs.filter(`"dhcp_identifier_type" =`, v)
}

type inHostsdhcpIdentifierTypeHostidentifiertype struct {
	qs HostidentifiertypeQS
}

func (in *inHostsdhcpIdentifierTypeHostidentifiertype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"dhcp_identifier_type" IN (` + s + `)`, p
}

func (qs HostsQS) DhcpIdentifierTypeIn(oqs HostidentifiertypeQS) HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&inHostsdhcpIdentifierTypeHostidentifiertype{
			qs: oqs,
		},
	)

	return qs
}

type notinHostsdhcpIdentifierTypeHostidentifiertype struct {
	qs HostidentifiertypeQS
}

func (nin *notinHostsdhcpIdentifierTypeHostidentifiertype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := nin.qs.QueryId(c)

	return `"dhcp_identifier_type" NOT IN (` + s + `)`, p
}

func (qs HostsQS) DhcpIdentifierTypeNotIn(oqs HostidentifiertypeQS) HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&notinHostsdhcpIdentifierTypeHostidentifiertype{
			qs: oqs,
		},
	)

	return qs
}

// OrderByDhcpIdentifierType sorts result by DhcpIdentifierType in ascending order
func (qs HostsQS) OrderByDhcpIdentifierType() HostsQS {
	qs.order = append(qs.order, `"dhcp_identifier_type"`)

	return qs
}

// OrderByDhcpIdentifierTypeDesc sorts result by DhcpIdentifierType in descending order
func (qs HostsQS) OrderByDhcpIdentifierTypeDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp_identifier_type" DESC`)

	return qs
}

// DistinctOnDhcpIdentifierType marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcpIdentifierType() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp_identifier_type"`)

	return qs
}

// END - django_kea.Hosts.dhcp_identifier_type

// BEGIN - django_kea.Hosts.dhcp4_subnet_id

// Dhcp4SubnetIdIsNull filters for Dhcp4SubnetId being null
func (qs HostsQS) Dhcp4SubnetIdIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdIsNotNull filters for Dhcp4SubnetId being not null
func (qs HostsQS) Dhcp4SubnetIdIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdEq filters for Dhcp4SubnetId being equal to argument
func (qs HostsQS) Dhcp4SubnetIdEq(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" =`, v)
}

// Dhcp4SubnetIdNe filters for Dhcp4SubnetId being not equal to argument
func (qs HostsQS) Dhcp4SubnetIdNe(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" <>`, v)
}

// Dhcp4SubnetIdLt filters for Dhcp4SubnetId being less than argument
func (qs HostsQS) Dhcp4SubnetIdLt(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" <`, v)
}

// Dhcp4SubnetIdLe filters for Dhcp4SubnetId being less than or equal to argument
func (qs HostsQS) Dhcp4SubnetIdLe(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" <=`, v)
}

// Dhcp4SubnetIdGt filters for Dhcp4SubnetId being greater than argument
func (qs HostsQS) Dhcp4SubnetIdGt(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" >`, v)
}

// Dhcp4SubnetIdGe filters for Dhcp4SubnetId being greater than or equal to argument
func (qs HostsQS) Dhcp4SubnetIdGe(v int64) HostsQS {
	return qs.filter(`"dhcp4_subnet_id" >=`, v)
}

type inHostsDhcp4SubnetId []interface{}

func (in inHostsDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4SubnetIdIn(values []int64) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp4SubnetId(vals),
	)

	return qs
}

type notinHostsDhcp4SubnetId []interface{}

func (in notinHostsDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4SubnetIdNotIn(values []int64) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp4SubnetId(vals),
	)

	return qs
}

// OrderByDhcp4SubnetId sorts result by Dhcp4SubnetId in ascending order
func (qs HostsQS) OrderByDhcp4SubnetId() HostsQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id"`)

	return qs
}

// OrderByDhcp4SubnetIdDesc sorts result by Dhcp4SubnetId in descending order
func (qs HostsQS) OrderByDhcp4SubnetIdDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id" DESC`)

	return qs
}

// DistinctOnDhcp4SubnetId marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp4SubnetId() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_subnet_id"`)

	return qs
}

// END - django_kea.Hosts.dhcp4_subnet_id

// BEGIN - django_kea.Hosts.dhcp6_subnet_id

// Dhcp6SubnetIdIsNull filters for Dhcp6SubnetId being null
func (qs HostsQS) Dhcp6SubnetIdIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdIsNotNull filters for Dhcp6SubnetId being not null
func (qs HostsQS) Dhcp6SubnetIdIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdEq filters for Dhcp6SubnetId being equal to argument
func (qs HostsQS) Dhcp6SubnetIdEq(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" =`, v)
}

// Dhcp6SubnetIdNe filters for Dhcp6SubnetId being not equal to argument
func (qs HostsQS) Dhcp6SubnetIdNe(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" <>`, v)
}

// Dhcp6SubnetIdLt filters for Dhcp6SubnetId being less than argument
func (qs HostsQS) Dhcp6SubnetIdLt(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" <`, v)
}

// Dhcp6SubnetIdLe filters for Dhcp6SubnetId being less than or equal to argument
func (qs HostsQS) Dhcp6SubnetIdLe(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" <=`, v)
}

// Dhcp6SubnetIdGt filters for Dhcp6SubnetId being greater than argument
func (qs HostsQS) Dhcp6SubnetIdGt(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" >`, v)
}

// Dhcp6SubnetIdGe filters for Dhcp6SubnetId being greater than or equal to argument
func (qs HostsQS) Dhcp6SubnetIdGe(v int64) HostsQS {
	return qs.filter(`"dhcp6_subnet_id" >=`, v)
}

type inHostsDhcp6SubnetId []interface{}

func (in inHostsDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp6SubnetIdIn(values []int64) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp6SubnetId(vals),
	)

	return qs
}

type notinHostsDhcp6SubnetId []interface{}

func (in notinHostsDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp6SubnetIdNotIn(values []int64) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp6SubnetId(vals),
	)

	return qs
}

// OrderByDhcp6SubnetId sorts result by Dhcp6SubnetId in ascending order
func (qs HostsQS) OrderByDhcp6SubnetId() HostsQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id"`)

	return qs
}

// OrderByDhcp6SubnetIdDesc sorts result by Dhcp6SubnetId in descending order
func (qs HostsQS) OrderByDhcp6SubnetIdDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id" DESC`)

	return qs
}

// DistinctOnDhcp6SubnetId marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp6SubnetId() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp6_subnet_id"`)

	return qs
}

// END - django_kea.Hosts.dhcp6_subnet_id

// BEGIN - django_kea.Hosts.ipv4_address

// Ipv4AddressIsNull filters for Ipv4Address being null
func (qs HostsQS) Ipv4AddressIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"ipv4_address" IS NULL`,
		},
	)
	return qs
}

// Ipv4AddressIsNotNull filters for Ipv4Address being not null
func (qs HostsQS) Ipv4AddressIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"ipv4_address" IS NOT NULL`,
		},
	)
	return qs
}

// Ipv4AddressEq filters for Ipv4Address being equal to argument
func (qs HostsQS) Ipv4AddressEq(v string) HostsQS {
	return qs.filter(`"ipv4_address" =`, v)
}

// Ipv4AddressNe filters for Ipv4Address being not equal to argument
func (qs HostsQS) Ipv4AddressNe(v string) HostsQS {
	return qs.filter(`"ipv4_address" <>`, v)
}

// Ipv4AddressLt filters for Ipv4Address being less than argument
func (qs HostsQS) Ipv4AddressLt(v string) HostsQS {
	return qs.filter(`"ipv4_address" <`, v)
}

// Ipv4AddressLe filters for Ipv4Address being less than or equal to argument
func (qs HostsQS) Ipv4AddressLe(v string) HostsQS {
	return qs.filter(`"ipv4_address" <=`, v)
}

// Ipv4AddressGt filters for Ipv4Address being greater than argument
func (qs HostsQS) Ipv4AddressGt(v string) HostsQS {
	return qs.filter(`"ipv4_address" >`, v)
}

// Ipv4AddressGe filters for Ipv4Address being greater than or equal to argument
func (qs HostsQS) Ipv4AddressGe(v string) HostsQS {
	return qs.filter(`"ipv4_address" >=`, v)
}

type inHostsIpv4Address []interface{}

func (in inHostsIpv4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"ipv4_address" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Ipv4AddressIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsIpv4Address(vals),
	)

	return qs
}

type notinHostsIpv4Address []interface{}

func (in notinHostsIpv4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"ipv4_address" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Ipv4AddressNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsIpv4Address(vals),
	)

	return qs
}

// OrderByIpv4Address sorts result by Ipv4Address in ascending order
func (qs HostsQS) OrderByIpv4Address() HostsQS {
	qs.order = append(qs.order, `"ipv4_address"`)

	return qs
}

// OrderByIpv4AddressDesc sorts result by Ipv4Address in descending order
func (qs HostsQS) OrderByIpv4AddressDesc() HostsQS {
	qs.order = append(qs.order, `"ipv4_address" DESC`)

	return qs
}

// DistinctOnIpv4Address marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnIpv4Address() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"ipv4_address"`)

	return qs
}

// END - django_kea.Hosts.ipv4_address

// BEGIN - django_kea.Hosts.hostname

// HostnameIsNull filters for Hostname being null
func (qs HostsQS) HostnameIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NULL`,
		},
	)
	return qs
}

// HostnameIsNotNull filters for Hostname being not null
func (qs HostsQS) HostnameIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NOT NULL`,
		},
	)
	return qs
}

// HostnameEq filters for Hostname being equal to argument
func (qs HostsQS) HostnameEq(v string) HostsQS {
	return qs.filter(`"hostname" =`, v)
}

// HostnameNe filters for Hostname being not equal to argument
func (qs HostsQS) HostnameNe(v string) HostsQS {
	return qs.filter(`"hostname" <>`, v)
}

// HostnameLt filters for Hostname being less than argument
func (qs HostsQS) HostnameLt(v string) HostsQS {
	return qs.filter(`"hostname" <`, v)
}

// HostnameLe filters for Hostname being less than or equal to argument
func (qs HostsQS) HostnameLe(v string) HostsQS {
	return qs.filter(`"hostname" <=`, v)
}

// HostnameGt filters for Hostname being greater than argument
func (qs HostsQS) HostnameGt(v string) HostsQS {
	return qs.filter(`"hostname" >`, v)
}

// HostnameGe filters for Hostname being greater than or equal to argument
func (qs HostsQS) HostnameGe(v string) HostsQS {
	return qs.filter(`"hostname" >=`, v)
}

type inHostsHostname []interface{}

func (in inHostsHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hostname" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) HostnameIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsHostname(vals),
	)

	return qs
}

type notinHostsHostname []interface{}

func (in notinHostsHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) HostnameNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsHostname(vals),
	)

	return qs
}

// OrderByHostname sorts result by Hostname in ascending order
func (qs HostsQS) OrderByHostname() HostsQS {
	qs.order = append(qs.order, `"hostname"`)

	return qs
}

// OrderByHostnameDesc sorts result by Hostname in descending order
func (qs HostsQS) OrderByHostnameDesc() HostsQS {
	qs.order = append(qs.order, `"hostname" DESC`)

	return qs
}

// DistinctOnHostname marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnHostname() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"hostname"`)

	return qs
}

// END - django_kea.Hosts.hostname

// BEGIN - django_kea.Hosts.dhcp4_client_classes

// Dhcp4ClientClassesIsNull filters for Dhcp4ClientClasses being null
func (qs HostsQS) Dhcp4ClientClassesIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_client_classes" IS NULL`,
		},
	)
	return qs
}

// Dhcp4ClientClassesIsNotNull filters for Dhcp4ClientClasses being not null
func (qs HostsQS) Dhcp4ClientClassesIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_client_classes" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4ClientClassesEq filters for Dhcp4ClientClasses being equal to argument
func (qs HostsQS) Dhcp4ClientClassesEq(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" =`, v)
}

// Dhcp4ClientClassesNe filters for Dhcp4ClientClasses being not equal to argument
func (qs HostsQS) Dhcp4ClientClassesNe(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" <>`, v)
}

// Dhcp4ClientClassesLt filters for Dhcp4ClientClasses being less than argument
func (qs HostsQS) Dhcp4ClientClassesLt(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" <`, v)
}

// Dhcp4ClientClassesLe filters for Dhcp4ClientClasses being less than or equal to argument
func (qs HostsQS) Dhcp4ClientClassesLe(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" <=`, v)
}

// Dhcp4ClientClassesGt filters for Dhcp4ClientClasses being greater than argument
func (qs HostsQS) Dhcp4ClientClassesGt(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" >`, v)
}

// Dhcp4ClientClassesGe filters for Dhcp4ClientClasses being greater than or equal to argument
func (qs HostsQS) Dhcp4ClientClassesGe(v string) HostsQS {
	return qs.filter(`"dhcp4_client_classes" >=`, v)
}

type inHostsDhcp4ClientClasses []interface{}

func (in inHostsDhcp4ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_client_classes" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4ClientClassesIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp4ClientClasses(vals),
	)

	return qs
}

type notinHostsDhcp4ClientClasses []interface{}

func (in notinHostsDhcp4ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_client_classes" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4ClientClassesNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp4ClientClasses(vals),
	)

	return qs
}

// OrderByDhcp4ClientClasses sorts result by Dhcp4ClientClasses in ascending order
func (qs HostsQS) OrderByDhcp4ClientClasses() HostsQS {
	qs.order = append(qs.order, `"dhcp4_client_classes"`)

	return qs
}

// OrderByDhcp4ClientClassesDesc sorts result by Dhcp4ClientClasses in descending order
func (qs HostsQS) OrderByDhcp4ClientClassesDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp4_client_classes" DESC`)

	return qs
}

// DistinctOnDhcp4ClientClasses marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp4ClientClasses() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_client_classes"`)

	return qs
}

// END - django_kea.Hosts.dhcp4_client_classes

// BEGIN - django_kea.Hosts.dhcp6_client_classes

// Dhcp6ClientClassesIsNull filters for Dhcp6ClientClasses being null
func (qs HostsQS) Dhcp6ClientClassesIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_client_classes" IS NULL`,
		},
	)
	return qs
}

// Dhcp6ClientClassesIsNotNull filters for Dhcp6ClientClasses being not null
func (qs HostsQS) Dhcp6ClientClassesIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_client_classes" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6ClientClassesEq filters for Dhcp6ClientClasses being equal to argument
func (qs HostsQS) Dhcp6ClientClassesEq(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" =`, v)
}

// Dhcp6ClientClassesNe filters for Dhcp6ClientClasses being not equal to argument
func (qs HostsQS) Dhcp6ClientClassesNe(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" <>`, v)
}

// Dhcp6ClientClassesLt filters for Dhcp6ClientClasses being less than argument
func (qs HostsQS) Dhcp6ClientClassesLt(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" <`, v)
}

// Dhcp6ClientClassesLe filters for Dhcp6ClientClasses being less than or equal to argument
func (qs HostsQS) Dhcp6ClientClassesLe(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" <=`, v)
}

// Dhcp6ClientClassesGt filters for Dhcp6ClientClasses being greater than argument
func (qs HostsQS) Dhcp6ClientClassesGt(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" >`, v)
}

// Dhcp6ClientClassesGe filters for Dhcp6ClientClasses being greater than or equal to argument
func (qs HostsQS) Dhcp6ClientClassesGe(v string) HostsQS {
	return qs.filter(`"dhcp6_client_classes" >=`, v)
}

type inHostsDhcp6ClientClasses []interface{}

func (in inHostsDhcp6ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_client_classes" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp6ClientClassesIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp6ClientClasses(vals),
	)

	return qs
}

type notinHostsDhcp6ClientClasses []interface{}

func (in notinHostsDhcp6ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp6_client_classes" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp6ClientClassesNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp6ClientClasses(vals),
	)

	return qs
}

// OrderByDhcp6ClientClasses sorts result by Dhcp6ClientClasses in ascending order
func (qs HostsQS) OrderByDhcp6ClientClasses() HostsQS {
	qs.order = append(qs.order, `"dhcp6_client_classes"`)

	return qs
}

// OrderByDhcp6ClientClassesDesc sorts result by Dhcp6ClientClasses in descending order
func (qs HostsQS) OrderByDhcp6ClientClassesDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp6_client_classes" DESC`)

	return qs
}

// DistinctOnDhcp6ClientClasses marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp6ClientClasses() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp6_client_classes"`)

	return qs
}

// END - django_kea.Hosts.dhcp6_client_classes

// BEGIN - django_kea.Hosts.dhcp4_next_server

// Dhcp4NextServerIsNull filters for Dhcp4NextServer being null
func (qs HostsQS) Dhcp4NextServerIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_next_server" IS NULL`,
		},
	)
	return qs
}

// Dhcp4NextServerIsNotNull filters for Dhcp4NextServer being not null
func (qs HostsQS) Dhcp4NextServerIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_next_server" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4NextServerEq filters for Dhcp4NextServer being equal to argument
func (qs HostsQS) Dhcp4NextServerEq(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" =`, v)
}

// Dhcp4NextServerNe filters for Dhcp4NextServer being not equal to argument
func (qs HostsQS) Dhcp4NextServerNe(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" <>`, v)
}

// Dhcp4NextServerLt filters for Dhcp4NextServer being less than argument
func (qs HostsQS) Dhcp4NextServerLt(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" <`, v)
}

// Dhcp4NextServerLe filters for Dhcp4NextServer being less than or equal to argument
func (qs HostsQS) Dhcp4NextServerLe(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" <=`, v)
}

// Dhcp4NextServerGt filters for Dhcp4NextServer being greater than argument
func (qs HostsQS) Dhcp4NextServerGt(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" >`, v)
}

// Dhcp4NextServerGe filters for Dhcp4NextServer being greater than or equal to argument
func (qs HostsQS) Dhcp4NextServerGe(v string) HostsQS {
	return qs.filter(`"dhcp4_next_server" >=`, v)
}

type inHostsDhcp4NextServer []interface{}

func (in inHostsDhcp4NextServer) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_next_server" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4NextServerIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp4NextServer(vals),
	)

	return qs
}

type notinHostsDhcp4NextServer []interface{}

func (in notinHostsDhcp4NextServer) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_next_server" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4NextServerNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp4NextServer(vals),
	)

	return qs
}

// OrderByDhcp4NextServer sorts result by Dhcp4NextServer in ascending order
func (qs HostsQS) OrderByDhcp4NextServer() HostsQS {
	qs.order = append(qs.order, `"dhcp4_next_server"`)

	return qs
}

// OrderByDhcp4NextServerDesc sorts result by Dhcp4NextServer in descending order
func (qs HostsQS) OrderByDhcp4NextServerDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp4_next_server" DESC`)

	return qs
}

// DistinctOnDhcp4NextServer marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp4NextServer() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_next_server"`)

	return qs
}

// END - django_kea.Hosts.dhcp4_next_server

// BEGIN - django_kea.Hosts.dhcp4_server_hostname

// Dhcp4ServerHostnameIsNull filters for Dhcp4ServerHostname being null
func (qs HostsQS) Dhcp4ServerHostnameIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_server_hostname" IS NULL`,
		},
	)
	return qs
}

// Dhcp4ServerHostnameIsNotNull filters for Dhcp4ServerHostname being not null
func (qs HostsQS) Dhcp4ServerHostnameIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_server_hostname" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4ServerHostnameEq filters for Dhcp4ServerHostname being equal to argument
func (qs HostsQS) Dhcp4ServerHostnameEq(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" =`, v)
}

// Dhcp4ServerHostnameNe filters for Dhcp4ServerHostname being not equal to argument
func (qs HostsQS) Dhcp4ServerHostnameNe(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" <>`, v)
}

// Dhcp4ServerHostnameLt filters for Dhcp4ServerHostname being less than argument
func (qs HostsQS) Dhcp4ServerHostnameLt(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" <`, v)
}

// Dhcp4ServerHostnameLe filters for Dhcp4ServerHostname being less than or equal to argument
func (qs HostsQS) Dhcp4ServerHostnameLe(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" <=`, v)
}

// Dhcp4ServerHostnameGt filters for Dhcp4ServerHostname being greater than argument
func (qs HostsQS) Dhcp4ServerHostnameGt(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" >`, v)
}

// Dhcp4ServerHostnameGe filters for Dhcp4ServerHostname being greater than or equal to argument
func (qs HostsQS) Dhcp4ServerHostnameGe(v string) HostsQS {
	return qs.filter(`"dhcp4_server_hostname" >=`, v)
}

type inHostsDhcp4ServerHostname []interface{}

func (in inHostsDhcp4ServerHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_server_hostname" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4ServerHostnameIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp4ServerHostname(vals),
	)

	return qs
}

type notinHostsDhcp4ServerHostname []interface{}

func (in notinHostsDhcp4ServerHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_server_hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4ServerHostnameNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp4ServerHostname(vals),
	)

	return qs
}

// OrderByDhcp4ServerHostname sorts result by Dhcp4ServerHostname in ascending order
func (qs HostsQS) OrderByDhcp4ServerHostname() HostsQS {
	qs.order = append(qs.order, `"dhcp4_server_hostname"`)

	return qs
}

// OrderByDhcp4ServerHostnameDesc sorts result by Dhcp4ServerHostname in descending order
func (qs HostsQS) OrderByDhcp4ServerHostnameDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp4_server_hostname" DESC`)

	return qs
}

// DistinctOnDhcp4ServerHostname marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp4ServerHostname() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_server_hostname"`)

	return qs
}

// END - django_kea.Hosts.dhcp4_server_hostname

// BEGIN - django_kea.Hosts.dhcp4_boot_file_name

// Dhcp4BootFileNameIsNull filters for Dhcp4BootFileName being null
func (qs HostsQS) Dhcp4BootFileNameIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_boot_file_name" IS NULL`,
		},
	)
	return qs
}

// Dhcp4BootFileNameIsNotNull filters for Dhcp4BootFileName being not null
func (qs HostsQS) Dhcp4BootFileNameIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_boot_file_name" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4BootFileNameEq filters for Dhcp4BootFileName being equal to argument
func (qs HostsQS) Dhcp4BootFileNameEq(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" =`, v)
}

// Dhcp4BootFileNameNe filters for Dhcp4BootFileName being not equal to argument
func (qs HostsQS) Dhcp4BootFileNameNe(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" <>`, v)
}

// Dhcp4BootFileNameLt filters for Dhcp4BootFileName being less than argument
func (qs HostsQS) Dhcp4BootFileNameLt(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" <`, v)
}

// Dhcp4BootFileNameLe filters for Dhcp4BootFileName being less than or equal to argument
func (qs HostsQS) Dhcp4BootFileNameLe(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" <=`, v)
}

// Dhcp4BootFileNameGt filters for Dhcp4BootFileName being greater than argument
func (qs HostsQS) Dhcp4BootFileNameGt(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" >`, v)
}

// Dhcp4BootFileNameGe filters for Dhcp4BootFileName being greater than or equal to argument
func (qs HostsQS) Dhcp4BootFileNameGe(v string) HostsQS {
	return qs.filter(`"dhcp4_boot_file_name" >=`, v)
}

type inHostsDhcp4BootFileName []interface{}

func (in inHostsDhcp4BootFileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_boot_file_name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4BootFileNameIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsDhcp4BootFileName(vals),
	)

	return qs
}

type notinHostsDhcp4BootFileName []interface{}

func (in notinHostsDhcp4BootFileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"dhcp4_boot_file_name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) Dhcp4BootFileNameNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsDhcp4BootFileName(vals),
	)

	return qs
}

// OrderByDhcp4BootFileName sorts result by Dhcp4BootFileName in ascending order
func (qs HostsQS) OrderByDhcp4BootFileName() HostsQS {
	qs.order = append(qs.order, `"dhcp4_boot_file_name"`)

	return qs
}

// OrderByDhcp4BootFileNameDesc sorts result by Dhcp4BootFileName in descending order
func (qs HostsQS) OrderByDhcp4BootFileNameDesc() HostsQS {
	qs.order = append(qs.order, `"dhcp4_boot_file_name" DESC`)

	return qs
}

// DistinctOnDhcp4BootFileName marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnDhcp4BootFileName() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"dhcp4_boot_file_name"`)

	return qs
}

// END - django_kea.Hosts.dhcp4_boot_file_name

// BEGIN - django_kea.Hosts.user_context

// UserContextIsNull filters for UserContext being null
func (qs HostsQS) UserContextIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs HostsQS) UserContextIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs HostsQS) UserContextEq(v string) HostsQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs HostsQS) UserContextNe(v string) HostsQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs HostsQS) UserContextLt(v string) HostsQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs HostsQS) UserContextLe(v string) HostsQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs HostsQS) UserContextGt(v string) HostsQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs HostsQS) UserContextGe(v string) HostsQS {
	return qs.filter(`"user_context" >=`, v)
}

type inHostsUserContext []interface{}

func (in inHostsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) UserContextIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsUserContext(vals),
	)

	return qs
}

type notinHostsUserContext []interface{}

func (in notinHostsUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) UserContextNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsUserContext(vals),
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs HostsQS) OrderByUserContext() HostsQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs HostsQS) OrderByUserContextDesc() HostsQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// DistinctOnUserContext marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnUserContext() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"user_context"`)

	return qs
}

// END - django_kea.Hosts.user_context

// BEGIN - django_kea.Hosts.auth_key

// AuthKeyIsNull filters for AuthKey being null
func (qs HostsQS) AuthKeyIsNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"auth_key" IS NULL`,
		},
	)
	return qs
}

// AuthKeyIsNotNull filters for AuthKey being not null
func (qs HostsQS) AuthKeyIsNotNull() HostsQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"auth_key" IS NOT NULL`,
		},
	)
	return qs
}

// AuthKeyEq filters for AuthKey being equal to argument
func (qs HostsQS) AuthKeyEq(v string) HostsQS {
	return qs.filter(`"auth_key" =`, v)
}

// AuthKeyNe filters for AuthKey being not equal to argument
func (qs HostsQS) AuthKeyNe(v string) HostsQS {
	return qs.filter(`"auth_key" <>`, v)
}

// AuthKeyLt filters for AuthKey being less than argument
func (qs HostsQS) AuthKeyLt(v string) HostsQS {
	return qs.filter(`"auth_key" <`, v)
}

// AuthKeyLe filters for AuthKey being less than or equal to argument
func (qs HostsQS) AuthKeyLe(v string) HostsQS {
	return qs.filter(`"auth_key" <=`, v)
}

// AuthKeyGt filters for AuthKey being greater than argument
func (qs HostsQS) AuthKeyGt(v string) HostsQS {
	return qs.filter(`"auth_key" >`, v)
}

// AuthKeyGe filters for AuthKey being greater than or equal to argument
func (qs HostsQS) AuthKeyGe(v string) HostsQS {
	return qs.filter(`"auth_key" >=`, v)
}

type inHostsAuthKey []interface{}

func (in inHostsAuthKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"auth_key" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) AuthKeyIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inHostsAuthKey(vals),
	)

	return qs
}

type notinHostsAuthKey []interface{}

func (in notinHostsAuthKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"auth_key" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs HostsQS) AuthKeyNotIn(values []string) HostsQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinHostsAuthKey(vals),
	)

	return qs
}

// OrderByAuthKey sorts result by AuthKey in ascending order
func (qs HostsQS) OrderByAuthKey() HostsQS {
	qs.order = append(qs.order, `"auth_key"`)

	return qs
}

// OrderByAuthKeyDesc sorts result by AuthKey in descending order
func (qs HostsQS) OrderByAuthKeyDesc() HostsQS {
	qs.order = append(qs.order, `"auth_key" DESC`)

	return qs
}

// DistinctOnAuthKey marks field in queries to add to DISTINCT ON clause
func (qs HostsQS) DistinctOnAuthKey() HostsQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"auth_key"`)

	return qs
}

// END - django_kea.Hosts.auth_key

// OrderByRandom randomizes result
func (qs HostsQS) OrderByRandom() HostsQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs HostsQS) ForUpdate() HostsQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs HostsQS) ForUpdateNowait() HostsQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs HostsQS) ForUpdateSkipLocked() HostsQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs HostsQS) ClearForUpdate() HostsQS {
	qs.forClause = ""

	return qs
}

func (qs HostsQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs HostsQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs HostsQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"host_id", "dhcp_identifier", "dhcp_identifier_type", "dhcp4_subnet_id", "dhcp6_subnet_id", "ipv4_address", "hostname", "dhcp4_client_classes", "dhcp6_client_classes", "dhcp4_next_server", "dhcp4_server_hostname", "dhcp4_boot_file_name", "user_context", "auth_key" FROM "hosts"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs HostsQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "host_id" FROM "hosts"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs HostsQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"host_id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "hosts"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs HostsQS) All(ctx context.Context, db models.DBInterface) (HostsList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret HostsList
	for rows.Next() {
		obj := Hosts{existsInDB: true}
		if err = rows.Scan(&obj.hostId, &obj.DhcpIdentifier, &obj.dhcpIdentifierType, &obj.Dhcp4SubnetId, &obj.Dhcp6SubnetId, &obj.Ipv4Address, &obj.Hostname, &obj.Dhcp4ClientClasses, &obj.Dhcp6ClientClasses, &obj.Dhcp4NextServer, &obj.Dhcp4ServerHostname, &obj.Dhcp4BootFileName, &obj.UserContext, &obj.AuthKey); err != nil {
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
func (qs HostsQS) First(ctx context.Context, db models.DBInterface) (*Hosts, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Hosts{existsInDB: true}
	err := row.Scan(&obj.hostId, &obj.DhcpIdentifier, &obj.dhcpIdentifierType, &obj.Dhcp4SubnetId, &obj.Dhcp6SubnetId, &obj.Ipv4Address, &obj.Hostname, &obj.Dhcp4ClientClasses, &obj.Dhcp6ClientClasses, &obj.Dhcp4NextServer, &obj.Dhcp4ServerHostname, &obj.Dhcp4BootFileName, &obj.UserContext, &obj.AuthKey)
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
func (qs HostsQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "hosts"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs HostsQS) Update() HostsUpdateQS {
	return HostsUpdateQS{condFragments: qs.condFragments}
}

// HostsUpdateQS represents an updated queryset for django_kea.Hosts
type HostsUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs HostsUpdateQS) update(c string, v interface{}) HostsUpdateQS {
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

// SetHostId sets HostId to the given value
func (uqs HostsUpdateQS) SetHostId(v int32) HostsUpdateQS {
	return uqs.update(`"host_id"`, v)
}

// SetDhcpIdentifier sets DhcpIdentifier to the given value
func (uqs HostsUpdateQS) SetDhcpIdentifier(v string) HostsUpdateQS {
	return uqs.update(`"dhcp_identifier"`, v)
}

// SetDhcpIdentifierType sets foreign key pointer to Hostidentifiertype
func (uqs HostsUpdateQS) SetDhcpIdentifierType(ptr *Hostidentifiertype) HostsUpdateQS {
	if ptr != nil {
		return uqs.update(`"dhcp_identifier_type"`, ptr.Type)
	}

	return uqs.update(`"dhcp_identifier_type"`, nil)
} // SetDhcp4SubnetId sets Dhcp4SubnetId to the given value
func (uqs HostsUpdateQS) SetDhcp4SubnetId(v sql.NullInt64) HostsUpdateQS {
	return uqs.update(`"dhcp4_subnet_id"`, v)
}

// SetDhcp6SubnetId sets Dhcp6SubnetId to the given value
func (uqs HostsUpdateQS) SetDhcp6SubnetId(v sql.NullInt64) HostsUpdateQS {
	return uqs.update(`"dhcp6_subnet_id"`, v)
}

// SetIpv4Address sets Ipv4Address to the given value
func (uqs HostsUpdateQS) SetIpv4Address(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"ipv4_address"`, v)
}

// SetHostname sets Hostname to the given value
func (uqs HostsUpdateQS) SetHostname(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"hostname"`, v)
}

// SetDhcp4ClientClasses sets Dhcp4ClientClasses to the given value
func (uqs HostsUpdateQS) SetDhcp4ClientClasses(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"dhcp4_client_classes"`, v)
}

// SetDhcp6ClientClasses sets Dhcp6ClientClasses to the given value
func (uqs HostsUpdateQS) SetDhcp6ClientClasses(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"dhcp6_client_classes"`, v)
}

// SetDhcp4NextServer sets Dhcp4NextServer to the given value
func (uqs HostsUpdateQS) SetDhcp4NextServer(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"dhcp4_next_server"`, v)
}

// SetDhcp4ServerHostname sets Dhcp4ServerHostname to the given value
func (uqs HostsUpdateQS) SetDhcp4ServerHostname(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"dhcp4_server_hostname"`, v)
}

// SetDhcp4BootFileName sets Dhcp4BootFileName to the given value
func (uqs HostsUpdateQS) SetDhcp4BootFileName(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"dhcp4_boot_file_name"`, v)
}

// SetUserContext sets UserContext to the given value
func (uqs HostsUpdateQS) SetUserContext(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// SetAuthKey sets AuthKey to the given value
func (uqs HostsUpdateQS) SetAuthKey(v sql.NullString) HostsUpdateQS {
	return uqs.update(`"auth_key"`, v)
}

// Exec executes the update operation
func (uqs HostsUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := HostsQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "hosts" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (h *Hosts) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "hosts" ("dhcp_identifier", "dhcp_identifier_type", "dhcp4_subnet_id", "dhcp6_subnet_id", "ipv4_address", "hostname", "dhcp4_client_classes", "dhcp6_client_classes", "dhcp4_next_server", "dhcp4_server_hostname", "dhcp4_boot_file_name", "user_context", "auth_key") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING "host_id"`, h.DhcpIdentifier, h.dhcpIdentifierType, h.Dhcp4SubnetId, h.Dhcp6SubnetId, h.Ipv4Address, h.Hostname, h.Dhcp4ClientClasses, h.Dhcp6ClientClasses, h.Dhcp4NextServer, h.Dhcp4ServerHostname, h.Dhcp4BootFileName, h.UserContext, h.AuthKey)

	if err := row.Scan(&h.hostId); err != nil {
		return err
	}

	h.existsInDB = true

	return nil
}

// update operation
func (h *Hosts) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "hosts" SET "dhcp_identifier" = $1, "dhcp_identifier_type" = $2, "dhcp4_subnet_id" = $3, "dhcp6_subnet_id" = $4, "ipv4_address" = $5, "hostname" = $6, "dhcp4_client_classes" = $7, "dhcp6_client_classes" = $8, "dhcp4_next_server" = $9, "dhcp4_server_hostname" = $10, "dhcp4_boot_file_name" = $11, "user_context" = $12, "auth_key" = $13 WHERE "host_id" = $14`, h.DhcpIdentifier, h.dhcpIdentifierType, h.Dhcp4SubnetId, h.Dhcp6SubnetId, h.Ipv4Address, h.Hostname, h.Dhcp4ClientClasses, h.Dhcp6ClientClasses, h.Dhcp4NextServer, h.Dhcp4ServerHostname, h.Dhcp4BootFileName, h.UserContext, h.AuthKey, h.hostId)

	return err
}

// Save inserts or updates record
func (h *Hosts) Save(ctx context.Context, db models.DBInterface) error {
	if h.existsInDB {
		return h.update(ctx, db)
	}

	return h.insert(ctx, db)
}

// Delete removes row from database
func (h *Hosts) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "hosts" WHERE "host_id" = $1`, h.hostId)

	h.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (hl HostsList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts HostsList

	for _, h := range hl {
		if h.existsInDB {
			if err := h.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, h)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 13*len(inserts))
	offs := 1
	for _, h := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4, offs+5, offs+6, offs+7, offs+8, offs+9, offs+10, offs+11, offs+12))
		vaa = append(vaa, h.DhcpIdentifier, h.dhcpIdentifierType, h.Dhcp4SubnetId, h.Dhcp6SubnetId, h.Ipv4Address, h.Hostname, h.Dhcp4ClientClasses, h.Dhcp6ClientClasses, h.Dhcp4NextServer, h.Dhcp4ServerHostname, h.Dhcp4BootFileName, h.UserContext, h.AuthKey)
		offs += 13
	}

	qs := `INSERT INTO "hosts" ("dhcp_identifier", "dhcp_identifier_type", "dhcp4_subnet_id", "dhcp6_subnet_id", "ipv4_address", "hostname", "dhcp4_client_classes", "dhcp6_client_classes", "dhcp4_next_server", "dhcp4_server_hostname", "dhcp4_boot_file_name", "user_context", "auth_key") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "host_id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, h := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&h.hostId); err != nil {
			return err
		}

		h.existsInDB = true
	}

	return rows.Err()
}

// Dhcp4options returns the set of Dhcp4options referencing this Hosts instance
func (h *Hosts) Dhcp4options() Dhcp4optionsQS {
	return Dhcp4optionsQS{}.HostEq(h)
}

// Dhcp6options returns the set of Dhcp6options referencing this Hosts instance
func (h *Hosts) Dhcp6options() Dhcp6optionsQS {
	return Dhcp6optionsQS{}.HostEq(h)
}

// Ipv6reservations returns the set of Ipv6reservations referencing this Hosts instance
func (h *Hosts) Ipv6reservations() Ipv6reservationsQS {
	return Ipv6reservationsQS{}.HostEq(h)
}
