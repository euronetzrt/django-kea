/*
  AUTO-GENERATED file for Django model django_kea.Host

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

// Host mirrors model django_kea.Host
type Host struct {
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

// HostQS represents a queryset for django_kea.Host
type HostQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs HostQS) filter(c string, p interface{}) HostQS {
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
func (qs HostQS) Or(exprs ...HostQS) HostQS {
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

// GetHostId returns Host.HostId
func (h *Host) GetHostId() int32 {
	return h.hostId
}

// HostIdEq filters for hostId being equal to argument
func (qs HostQS) HostIdEq(v int32) HostQS {
	return qs.filter(`"host_id" =`, v)
}

// HostIdNe filters for hostId being not equal to argument
func (qs HostQS) HostIdNe(v int32) HostQS {
	return qs.filter(`"host_id" <>`, v)
}

// HostIdLt filters for hostId being less than argument
func (qs HostQS) HostIdLt(v int32) HostQS {
	return qs.filter(`"host_id" <`, v)
}

// HostIdLe filters for hostId being less than or equal to argument
func (qs HostQS) HostIdLe(v int32) HostQS {
	return qs.filter(`"host_id" <=`, v)
}

// HostIdGt filters for hostId being greater than argument
func (qs HostQS) HostIdGt(v int32) HostQS {
	return qs.filter(`"host_id" >`, v)
}

// HostIdGe filters for hostId being greater than or equal to argument
func (qs HostQS) HostIdGe(v int32) HostQS {
	return qs.filter(`"host_id" >=`, v)
}

type inHosthostId struct {
	values []interface{}
}

func (in *inHosthostId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"host_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) HostIdIn(values []int32) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHosthostId{
			values: vals,
		},
	)

	return qs
}

type notinHosthostId struct {
	values []interface{}
}

func (in *notinHosthostId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"host_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) HostIdNotIn(values []int32) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHosthostId{
			values: vals,
		},
	)

	return qs
}

// OrderByHostId sorts result by HostId in ascending order
func (qs HostQS) OrderByHostId() HostQS {
	qs.order = append(qs.order, `"host_id"`)

	return qs
}

// OrderByHostIdDesc sorts result by HostId in descending order
func (qs HostQS) OrderByHostIdDesc() HostQS {
	qs.order = append(qs.order, `"host_id" DESC`)

	return qs
}

// DhcpIdentifierEq filters for DhcpIdentifier being equal to argument
func (qs HostQS) DhcpIdentifierEq(v string) HostQS {
	return qs.filter(`"dhcp_identifier" =`, v)
}

// DhcpIdentifierNe filters for DhcpIdentifier being not equal to argument
func (qs HostQS) DhcpIdentifierNe(v string) HostQS {
	return qs.filter(`"dhcp_identifier" <>`, v)
}

// DhcpIdentifierLt filters for DhcpIdentifier being less than argument
func (qs HostQS) DhcpIdentifierLt(v string) HostQS {
	return qs.filter(`"dhcp_identifier" <`, v)
}

// DhcpIdentifierLe filters for DhcpIdentifier being less than or equal to argument
func (qs HostQS) DhcpIdentifierLe(v string) HostQS {
	return qs.filter(`"dhcp_identifier" <=`, v)
}

// DhcpIdentifierGt filters for DhcpIdentifier being greater than argument
func (qs HostQS) DhcpIdentifierGt(v string) HostQS {
	return qs.filter(`"dhcp_identifier" >`, v)
}

// DhcpIdentifierGe filters for DhcpIdentifier being greater than or equal to argument
func (qs HostQS) DhcpIdentifierGe(v string) HostQS {
	return qs.filter(`"dhcp_identifier" >=`, v)
}

type inHostDhcpIdentifier struct {
	values []interface{}
}

func (in *inHostDhcpIdentifier) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_identifier" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) DhcpIdentifierIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcpIdentifier{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcpIdentifier struct {
	values []interface{}
}

func (in *notinHostDhcpIdentifier) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp_identifier" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) DhcpIdentifierNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcpIdentifier{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcpIdentifier sorts result by DhcpIdentifier in ascending order
func (qs HostQS) OrderByDhcpIdentifier() HostQS {
	qs.order = append(qs.order, `"dhcp_identifier"`)

	return qs
}

// OrderByDhcpIdentifierDesc sorts result by DhcpIdentifier in descending order
func (qs HostQS) OrderByDhcpIdentifierDesc() HostQS {
	qs.order = append(qs.order, `"dhcp_identifier" DESC`)

	return qs
}

// GetDhcpIdentifierType returns Hostidentifiertype
func (h *Host) GetDhcpIdentifierType(db models.DBInterface) (*Hostidentifiertype, error) {
	return HostidentifiertypeQS{}.TypeEq(h.dhcpIdentifierType).First(db)
}

// SetDhcpIdentifierType sets foreign key pointer to Hostidentifiertype
func (h *Host) SetDhcpIdentifierType(ptr *Hostidentifiertype) error {
	if ptr != nil {
		h.dhcpIdentifierType = ptr.Type
	} else {
		return fmt.Errorf("Host.SetDhcpIdentifierType: non-null field received null value")
	}

	return nil
}

// GetDhcpIdentifierTypeRaw returns Host.DhcpIdentifierType
func (h *Host) GetDhcpIdentifierTypeRaw() int32 {
	return h.dhcpIdentifierType
}

// DhcpIdentifierTypeEq filters for dhcpIdentifierType being equal to argument
func (qs HostQS) DhcpIdentifierTypeEq(v *Hostidentifiertype) HostQS {
	return qs.filter(`"dhcp_identifier_type" =`, v.Type)
}

// DhcpIdentifierTypeRawEq filters for dhcpIdentifierType being equal to raw argument
func (qs HostQS) DhcpIdentifierTypeRawEq(v int32) HostQS {
	return qs.filter(`"dhcp_identifier_type" =`, v)
}

type inHostdhcpIdentifierTypeHostidentifiertype struct {
	qs HostidentifiertypeQS
}

func (in *inHostdhcpIdentifierTypeHostidentifiertype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"dhcp_identifier_type" IN (` + s + `)`, p
}

func (qs HostQS) DhcpIdentifierTypeIn(oqs HostidentifiertypeQS) HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&inHostdhcpIdentifierTypeHostidentifiertype{
			qs: oqs,
		},
	)

	return qs
}

// OrderByDhcpIdentifierType sorts result by DhcpIdentifierType in ascending order
func (qs HostQS) OrderByDhcpIdentifierType() HostQS {
	qs.order = append(qs.order, `"dhcp_identifier_type"`)

	return qs
}

// OrderByDhcpIdentifierTypeDesc sorts result by DhcpIdentifierType in descending order
func (qs HostQS) OrderByDhcpIdentifierTypeDesc() HostQS {
	qs.order = append(qs.order, `"dhcp_identifier_type" DESC`)

	return qs
}

// Dhcp4SubnetIdIsNull filters for Dhcp4SubnetId being null
func (qs HostQS) Dhcp4SubnetIdIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdIsNotNull filters for Dhcp4SubnetId being not null
func (qs HostQS) Dhcp4SubnetIdIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4SubnetIdEq filters for Dhcp4SubnetId being equal to argument
func (qs HostQS) Dhcp4SubnetIdEq(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" =`, v)
}

// Dhcp4SubnetIdNe filters for Dhcp4SubnetId being not equal to argument
func (qs HostQS) Dhcp4SubnetIdNe(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" <>`, v)
}

// Dhcp4SubnetIdLt filters for Dhcp4SubnetId being less than argument
func (qs HostQS) Dhcp4SubnetIdLt(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" <`, v)
}

// Dhcp4SubnetIdLe filters for Dhcp4SubnetId being less than or equal to argument
func (qs HostQS) Dhcp4SubnetIdLe(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" <=`, v)
}

// Dhcp4SubnetIdGt filters for Dhcp4SubnetId being greater than argument
func (qs HostQS) Dhcp4SubnetIdGt(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" >`, v)
}

// Dhcp4SubnetIdGe filters for Dhcp4SubnetId being greater than or equal to argument
func (qs HostQS) Dhcp4SubnetIdGe(v int64) HostQS {
	return qs.filter(`"dhcp4_subnet_id" >=`, v)
}

type inHostDhcp4SubnetId struct {
	values []interface{}
}

func (in *inHostDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4SubnetIdIn(values []int64) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp4SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp4SubnetId struct {
	values []interface{}
}

func (in *notinHostDhcp4SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4SubnetIdNotIn(values []int64) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp4SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4SubnetId sorts result by Dhcp4SubnetId in ascending order
func (qs HostQS) OrderByDhcp4SubnetId() HostQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id"`)

	return qs
}

// OrderByDhcp4SubnetIdDesc sorts result by Dhcp4SubnetId in descending order
func (qs HostQS) OrderByDhcp4SubnetIdDesc() HostQS {
	qs.order = append(qs.order, `"dhcp4_subnet_id" DESC`)

	return qs
}

// Dhcp6SubnetIdIsNull filters for Dhcp6SubnetId being null
func (qs HostQS) Dhcp6SubnetIdIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdIsNotNull filters for Dhcp6SubnetId being not null
func (qs HostQS) Dhcp6SubnetIdIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_subnet_id" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6SubnetIdEq filters for Dhcp6SubnetId being equal to argument
func (qs HostQS) Dhcp6SubnetIdEq(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" =`, v)
}

// Dhcp6SubnetIdNe filters for Dhcp6SubnetId being not equal to argument
func (qs HostQS) Dhcp6SubnetIdNe(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" <>`, v)
}

// Dhcp6SubnetIdLt filters for Dhcp6SubnetId being less than argument
func (qs HostQS) Dhcp6SubnetIdLt(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" <`, v)
}

// Dhcp6SubnetIdLe filters for Dhcp6SubnetId being less than or equal to argument
func (qs HostQS) Dhcp6SubnetIdLe(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" <=`, v)
}

// Dhcp6SubnetIdGt filters for Dhcp6SubnetId being greater than argument
func (qs HostQS) Dhcp6SubnetIdGt(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" >`, v)
}

// Dhcp6SubnetIdGe filters for Dhcp6SubnetId being greater than or equal to argument
func (qs HostQS) Dhcp6SubnetIdGe(v int64) HostQS {
	return qs.filter(`"dhcp6_subnet_id" >=`, v)
}

type inHostDhcp6SubnetId struct {
	values []interface{}
}

func (in *inHostDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp6SubnetIdIn(values []int64) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp6SubnetId{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp6SubnetId struct {
	values []interface{}
}

func (in *notinHostDhcp6SubnetId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_subnet_id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp6SubnetIdNotIn(values []int64) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp6SubnetId{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp6SubnetId sorts result by Dhcp6SubnetId in ascending order
func (qs HostQS) OrderByDhcp6SubnetId() HostQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id"`)

	return qs
}

// OrderByDhcp6SubnetIdDesc sorts result by Dhcp6SubnetId in descending order
func (qs HostQS) OrderByDhcp6SubnetIdDesc() HostQS {
	qs.order = append(qs.order, `"dhcp6_subnet_id" DESC`)

	return qs
}

// Ipv4AddressIsNull filters for Ipv4Address being null
func (qs HostQS) Ipv4AddressIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"ipv4_address" IS NULL`,
		},
	)
	return qs
}

// Ipv4AddressIsNotNull filters for Ipv4Address being not null
func (qs HostQS) Ipv4AddressIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"ipv4_address" IS NOT NULL`,
		},
	)
	return qs
}

// Ipv4AddressEq filters for Ipv4Address being equal to argument
func (qs HostQS) Ipv4AddressEq(v string) HostQS {
	return qs.filter(`"ipv4_address" =`, v)
}

// Ipv4AddressNe filters for Ipv4Address being not equal to argument
func (qs HostQS) Ipv4AddressNe(v string) HostQS {
	return qs.filter(`"ipv4_address" <>`, v)
}

// Ipv4AddressLt filters for Ipv4Address being less than argument
func (qs HostQS) Ipv4AddressLt(v string) HostQS {
	return qs.filter(`"ipv4_address" <`, v)
}

// Ipv4AddressLe filters for Ipv4Address being less than or equal to argument
func (qs HostQS) Ipv4AddressLe(v string) HostQS {
	return qs.filter(`"ipv4_address" <=`, v)
}

// Ipv4AddressGt filters for Ipv4Address being greater than argument
func (qs HostQS) Ipv4AddressGt(v string) HostQS {
	return qs.filter(`"ipv4_address" >`, v)
}

// Ipv4AddressGe filters for Ipv4Address being greater than or equal to argument
func (qs HostQS) Ipv4AddressGe(v string) HostQS {
	return qs.filter(`"ipv4_address" >=`, v)
}

type inHostIpv4Address struct {
	values []interface{}
}

func (in *inHostIpv4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ipv4_address" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Ipv4AddressIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostIpv4Address{
			values: vals,
		},
	)

	return qs
}

type notinHostIpv4Address struct {
	values []interface{}
}

func (in *notinHostIpv4Address) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ipv4_address" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Ipv4AddressNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostIpv4Address{
			values: vals,
		},
	)

	return qs
}

// OrderByIpv4Address sorts result by Ipv4Address in ascending order
func (qs HostQS) OrderByIpv4Address() HostQS {
	qs.order = append(qs.order, `"ipv4_address"`)

	return qs
}

// OrderByIpv4AddressDesc sorts result by Ipv4Address in descending order
func (qs HostQS) OrderByIpv4AddressDesc() HostQS {
	qs.order = append(qs.order, `"ipv4_address" DESC`)

	return qs
}

// HostnameIsNull filters for Hostname being null
func (qs HostQS) HostnameIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NULL`,
		},
	)
	return qs
}

// HostnameIsNotNull filters for Hostname being not null
func (qs HostQS) HostnameIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"hostname" IS NOT NULL`,
		},
	)
	return qs
}

// HostnameEq filters for Hostname being equal to argument
func (qs HostQS) HostnameEq(v string) HostQS {
	return qs.filter(`"hostname" =`, v)
}

// HostnameNe filters for Hostname being not equal to argument
func (qs HostQS) HostnameNe(v string) HostQS {
	return qs.filter(`"hostname" <>`, v)
}

// HostnameLt filters for Hostname being less than argument
func (qs HostQS) HostnameLt(v string) HostQS {
	return qs.filter(`"hostname" <`, v)
}

// HostnameLe filters for Hostname being less than or equal to argument
func (qs HostQS) HostnameLe(v string) HostQS {
	return qs.filter(`"hostname" <=`, v)
}

// HostnameGt filters for Hostname being greater than argument
func (qs HostQS) HostnameGt(v string) HostQS {
	return qs.filter(`"hostname" >`, v)
}

// HostnameGe filters for Hostname being greater than or equal to argument
func (qs HostQS) HostnameGe(v string) HostQS {
	return qs.filter(`"hostname" >=`, v)
}

type inHostHostname struct {
	values []interface{}
}

func (in *inHostHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) HostnameIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostHostname{
			values: vals,
		},
	)

	return qs
}

type notinHostHostname struct {
	values []interface{}
}

func (in *notinHostHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) HostnameNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostHostname{
			values: vals,
		},
	)

	return qs
}

// OrderByHostname sorts result by Hostname in ascending order
func (qs HostQS) OrderByHostname() HostQS {
	qs.order = append(qs.order, `"hostname"`)

	return qs
}

// OrderByHostnameDesc sorts result by Hostname in descending order
func (qs HostQS) OrderByHostnameDesc() HostQS {
	qs.order = append(qs.order, `"hostname" DESC`)

	return qs
}

// Dhcp4ClientClassesIsNull filters for Dhcp4ClientClasses being null
func (qs HostQS) Dhcp4ClientClassesIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_client_classes" IS NULL`,
		},
	)
	return qs
}

// Dhcp4ClientClassesIsNotNull filters for Dhcp4ClientClasses being not null
func (qs HostQS) Dhcp4ClientClassesIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_client_classes" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4ClientClassesEq filters for Dhcp4ClientClasses being equal to argument
func (qs HostQS) Dhcp4ClientClassesEq(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" =`, v)
}

// Dhcp4ClientClassesNe filters for Dhcp4ClientClasses being not equal to argument
func (qs HostQS) Dhcp4ClientClassesNe(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" <>`, v)
}

// Dhcp4ClientClassesLt filters for Dhcp4ClientClasses being less than argument
func (qs HostQS) Dhcp4ClientClassesLt(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" <`, v)
}

// Dhcp4ClientClassesLe filters for Dhcp4ClientClasses being less than or equal to argument
func (qs HostQS) Dhcp4ClientClassesLe(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" <=`, v)
}

// Dhcp4ClientClassesGt filters for Dhcp4ClientClasses being greater than argument
func (qs HostQS) Dhcp4ClientClassesGt(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" >`, v)
}

// Dhcp4ClientClassesGe filters for Dhcp4ClientClasses being greater than or equal to argument
func (qs HostQS) Dhcp4ClientClassesGe(v string) HostQS {
	return qs.filter(`"dhcp4_client_classes" >=`, v)
}

type inHostDhcp4ClientClasses struct {
	values []interface{}
}

func (in *inHostDhcp4ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_client_classes" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4ClientClassesIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp4ClientClasses{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp4ClientClasses struct {
	values []interface{}
}

func (in *notinHostDhcp4ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_client_classes" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4ClientClassesNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp4ClientClasses{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4ClientClasses sorts result by Dhcp4ClientClasses in ascending order
func (qs HostQS) OrderByDhcp4ClientClasses() HostQS {
	qs.order = append(qs.order, `"dhcp4_client_classes"`)

	return qs
}

// OrderByDhcp4ClientClassesDesc sorts result by Dhcp4ClientClasses in descending order
func (qs HostQS) OrderByDhcp4ClientClassesDesc() HostQS {
	qs.order = append(qs.order, `"dhcp4_client_classes" DESC`)

	return qs
}

// Dhcp6ClientClassesIsNull filters for Dhcp6ClientClasses being null
func (qs HostQS) Dhcp6ClientClassesIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_client_classes" IS NULL`,
		},
	)
	return qs
}

// Dhcp6ClientClassesIsNotNull filters for Dhcp6ClientClasses being not null
func (qs HostQS) Dhcp6ClientClassesIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp6_client_classes" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp6ClientClassesEq filters for Dhcp6ClientClasses being equal to argument
func (qs HostQS) Dhcp6ClientClassesEq(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" =`, v)
}

// Dhcp6ClientClassesNe filters for Dhcp6ClientClasses being not equal to argument
func (qs HostQS) Dhcp6ClientClassesNe(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" <>`, v)
}

// Dhcp6ClientClassesLt filters for Dhcp6ClientClasses being less than argument
func (qs HostQS) Dhcp6ClientClassesLt(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" <`, v)
}

// Dhcp6ClientClassesLe filters for Dhcp6ClientClasses being less than or equal to argument
func (qs HostQS) Dhcp6ClientClassesLe(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" <=`, v)
}

// Dhcp6ClientClassesGt filters for Dhcp6ClientClasses being greater than argument
func (qs HostQS) Dhcp6ClientClassesGt(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" >`, v)
}

// Dhcp6ClientClassesGe filters for Dhcp6ClientClasses being greater than or equal to argument
func (qs HostQS) Dhcp6ClientClassesGe(v string) HostQS {
	return qs.filter(`"dhcp6_client_classes" >=`, v)
}

type inHostDhcp6ClientClasses struct {
	values []interface{}
}

func (in *inHostDhcp6ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_client_classes" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp6ClientClassesIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp6ClientClasses{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp6ClientClasses struct {
	values []interface{}
}

func (in *notinHostDhcp6ClientClasses) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp6_client_classes" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp6ClientClassesNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp6ClientClasses{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp6ClientClasses sorts result by Dhcp6ClientClasses in ascending order
func (qs HostQS) OrderByDhcp6ClientClasses() HostQS {
	qs.order = append(qs.order, `"dhcp6_client_classes"`)

	return qs
}

// OrderByDhcp6ClientClassesDesc sorts result by Dhcp6ClientClasses in descending order
func (qs HostQS) OrderByDhcp6ClientClassesDesc() HostQS {
	qs.order = append(qs.order, `"dhcp6_client_classes" DESC`)

	return qs
}

// Dhcp4NextServerIsNull filters for Dhcp4NextServer being null
func (qs HostQS) Dhcp4NextServerIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_next_server" IS NULL`,
		},
	)
	return qs
}

// Dhcp4NextServerIsNotNull filters for Dhcp4NextServer being not null
func (qs HostQS) Dhcp4NextServerIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_next_server" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4NextServerEq filters for Dhcp4NextServer being equal to argument
func (qs HostQS) Dhcp4NextServerEq(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" =`, v)
}

// Dhcp4NextServerNe filters for Dhcp4NextServer being not equal to argument
func (qs HostQS) Dhcp4NextServerNe(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" <>`, v)
}

// Dhcp4NextServerLt filters for Dhcp4NextServer being less than argument
func (qs HostQS) Dhcp4NextServerLt(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" <`, v)
}

// Dhcp4NextServerLe filters for Dhcp4NextServer being less than or equal to argument
func (qs HostQS) Dhcp4NextServerLe(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" <=`, v)
}

// Dhcp4NextServerGt filters for Dhcp4NextServer being greater than argument
func (qs HostQS) Dhcp4NextServerGt(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" >`, v)
}

// Dhcp4NextServerGe filters for Dhcp4NextServer being greater than or equal to argument
func (qs HostQS) Dhcp4NextServerGe(v string) HostQS {
	return qs.filter(`"dhcp4_next_server" >=`, v)
}

type inHostDhcp4NextServer struct {
	values []interface{}
}

func (in *inHostDhcp4NextServer) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_next_server" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4NextServerIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp4NextServer{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp4NextServer struct {
	values []interface{}
}

func (in *notinHostDhcp4NextServer) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_next_server" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4NextServerNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp4NextServer{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4NextServer sorts result by Dhcp4NextServer in ascending order
func (qs HostQS) OrderByDhcp4NextServer() HostQS {
	qs.order = append(qs.order, `"dhcp4_next_server"`)

	return qs
}

// OrderByDhcp4NextServerDesc sorts result by Dhcp4NextServer in descending order
func (qs HostQS) OrderByDhcp4NextServerDesc() HostQS {
	qs.order = append(qs.order, `"dhcp4_next_server" DESC`)

	return qs
}

// Dhcp4ServerHostnameIsNull filters for Dhcp4ServerHostname being null
func (qs HostQS) Dhcp4ServerHostnameIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_server_hostname" IS NULL`,
		},
	)
	return qs
}

// Dhcp4ServerHostnameIsNotNull filters for Dhcp4ServerHostname being not null
func (qs HostQS) Dhcp4ServerHostnameIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_server_hostname" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4ServerHostnameEq filters for Dhcp4ServerHostname being equal to argument
func (qs HostQS) Dhcp4ServerHostnameEq(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" =`, v)
}

// Dhcp4ServerHostnameNe filters for Dhcp4ServerHostname being not equal to argument
func (qs HostQS) Dhcp4ServerHostnameNe(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" <>`, v)
}

// Dhcp4ServerHostnameLt filters for Dhcp4ServerHostname being less than argument
func (qs HostQS) Dhcp4ServerHostnameLt(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" <`, v)
}

// Dhcp4ServerHostnameLe filters for Dhcp4ServerHostname being less than or equal to argument
func (qs HostQS) Dhcp4ServerHostnameLe(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" <=`, v)
}

// Dhcp4ServerHostnameGt filters for Dhcp4ServerHostname being greater than argument
func (qs HostQS) Dhcp4ServerHostnameGt(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" >`, v)
}

// Dhcp4ServerHostnameGe filters for Dhcp4ServerHostname being greater than or equal to argument
func (qs HostQS) Dhcp4ServerHostnameGe(v string) HostQS {
	return qs.filter(`"dhcp4_server_hostname" >=`, v)
}

type inHostDhcp4ServerHostname struct {
	values []interface{}
}

func (in *inHostDhcp4ServerHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_server_hostname" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4ServerHostnameIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp4ServerHostname{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp4ServerHostname struct {
	values []interface{}
}

func (in *notinHostDhcp4ServerHostname) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_server_hostname" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4ServerHostnameNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp4ServerHostname{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4ServerHostname sorts result by Dhcp4ServerHostname in ascending order
func (qs HostQS) OrderByDhcp4ServerHostname() HostQS {
	qs.order = append(qs.order, `"dhcp4_server_hostname"`)

	return qs
}

// OrderByDhcp4ServerHostnameDesc sorts result by Dhcp4ServerHostname in descending order
func (qs HostQS) OrderByDhcp4ServerHostnameDesc() HostQS {
	qs.order = append(qs.order, `"dhcp4_server_hostname" DESC`)

	return qs
}

// Dhcp4BootFileNameIsNull filters for Dhcp4BootFileName being null
func (qs HostQS) Dhcp4BootFileNameIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_boot_file_name" IS NULL`,
		},
	)
	return qs
}

// Dhcp4BootFileNameIsNotNull filters for Dhcp4BootFileName being not null
func (qs HostQS) Dhcp4BootFileNameIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"dhcp4_boot_file_name" IS NOT NULL`,
		},
	)
	return qs
}

// Dhcp4BootFileNameEq filters for Dhcp4BootFileName being equal to argument
func (qs HostQS) Dhcp4BootFileNameEq(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" =`, v)
}

// Dhcp4BootFileNameNe filters for Dhcp4BootFileName being not equal to argument
func (qs HostQS) Dhcp4BootFileNameNe(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" <>`, v)
}

// Dhcp4BootFileNameLt filters for Dhcp4BootFileName being less than argument
func (qs HostQS) Dhcp4BootFileNameLt(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" <`, v)
}

// Dhcp4BootFileNameLe filters for Dhcp4BootFileName being less than or equal to argument
func (qs HostQS) Dhcp4BootFileNameLe(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" <=`, v)
}

// Dhcp4BootFileNameGt filters for Dhcp4BootFileName being greater than argument
func (qs HostQS) Dhcp4BootFileNameGt(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" >`, v)
}

// Dhcp4BootFileNameGe filters for Dhcp4BootFileName being greater than or equal to argument
func (qs HostQS) Dhcp4BootFileNameGe(v string) HostQS {
	return qs.filter(`"dhcp4_boot_file_name" >=`, v)
}

type inHostDhcp4BootFileName struct {
	values []interface{}
}

func (in *inHostDhcp4BootFileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_boot_file_name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4BootFileNameIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostDhcp4BootFileName{
			values: vals,
		},
	)

	return qs
}

type notinHostDhcp4BootFileName struct {
	values []interface{}
}

func (in *notinHostDhcp4BootFileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"dhcp4_boot_file_name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) Dhcp4BootFileNameNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostDhcp4BootFileName{
			values: vals,
		},
	)

	return qs
}

// OrderByDhcp4BootFileName sorts result by Dhcp4BootFileName in ascending order
func (qs HostQS) OrderByDhcp4BootFileName() HostQS {
	qs.order = append(qs.order, `"dhcp4_boot_file_name"`)

	return qs
}

// OrderByDhcp4BootFileNameDesc sorts result by Dhcp4BootFileName in descending order
func (qs HostQS) OrderByDhcp4BootFileNameDesc() HostQS {
	qs.order = append(qs.order, `"dhcp4_boot_file_name" DESC`)

	return qs
}

// UserContextIsNull filters for UserContext being null
func (qs HostQS) UserContextIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NULL`,
		},
	)
	return qs
}

// UserContextIsNotNull filters for UserContext being not null
func (qs HostQS) UserContextIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"user_context" IS NOT NULL`,
		},
	)
	return qs
}

// UserContextEq filters for UserContext being equal to argument
func (qs HostQS) UserContextEq(v string) HostQS {
	return qs.filter(`"user_context" =`, v)
}

// UserContextNe filters for UserContext being not equal to argument
func (qs HostQS) UserContextNe(v string) HostQS {
	return qs.filter(`"user_context" <>`, v)
}

// UserContextLt filters for UserContext being less than argument
func (qs HostQS) UserContextLt(v string) HostQS {
	return qs.filter(`"user_context" <`, v)
}

// UserContextLe filters for UserContext being less than or equal to argument
func (qs HostQS) UserContextLe(v string) HostQS {
	return qs.filter(`"user_context" <=`, v)
}

// UserContextGt filters for UserContext being greater than argument
func (qs HostQS) UserContextGt(v string) HostQS {
	return qs.filter(`"user_context" >`, v)
}

// UserContextGe filters for UserContext being greater than or equal to argument
func (qs HostQS) UserContextGe(v string) HostQS {
	return qs.filter(`"user_context" >=`, v)
}

type inHostUserContext struct {
	values []interface{}
}

func (in *inHostUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) UserContextIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostUserContext{
			values: vals,
		},
	)

	return qs
}

type notinHostUserContext struct {
	values []interface{}
}

func (in *notinHostUserContext) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"user_context" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) UserContextNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostUserContext{
			values: vals,
		},
	)

	return qs
}

// OrderByUserContext sorts result by UserContext in ascending order
func (qs HostQS) OrderByUserContext() HostQS {
	qs.order = append(qs.order, `"user_context"`)

	return qs
}

// OrderByUserContextDesc sorts result by UserContext in descending order
func (qs HostQS) OrderByUserContextDesc() HostQS {
	qs.order = append(qs.order, `"user_context" DESC`)

	return qs
}

// AuthKeyIsNull filters for AuthKey being null
func (qs HostQS) AuthKeyIsNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"auth_key" IS NULL`,
		},
	)
	return qs
}

// AuthKeyIsNotNull filters for AuthKey being not null
func (qs HostQS) AuthKeyIsNotNull() HostQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"auth_key" IS NOT NULL`,
		},
	)
	return qs
}

// AuthKeyEq filters for AuthKey being equal to argument
func (qs HostQS) AuthKeyEq(v string) HostQS {
	return qs.filter(`"auth_key" =`, v)
}

// AuthKeyNe filters for AuthKey being not equal to argument
func (qs HostQS) AuthKeyNe(v string) HostQS {
	return qs.filter(`"auth_key" <>`, v)
}

// AuthKeyLt filters for AuthKey being less than argument
func (qs HostQS) AuthKeyLt(v string) HostQS {
	return qs.filter(`"auth_key" <`, v)
}

// AuthKeyLe filters for AuthKey being less than or equal to argument
func (qs HostQS) AuthKeyLe(v string) HostQS {
	return qs.filter(`"auth_key" <=`, v)
}

// AuthKeyGt filters for AuthKey being greater than argument
func (qs HostQS) AuthKeyGt(v string) HostQS {
	return qs.filter(`"auth_key" >`, v)
}

// AuthKeyGe filters for AuthKey being greater than or equal to argument
func (qs HostQS) AuthKeyGe(v string) HostQS {
	return qs.filter(`"auth_key" >=`, v)
}

type inHostAuthKey struct {
	values []interface{}
}

func (in *inHostAuthKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"auth_key" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) AuthKeyIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inHostAuthKey{
			values: vals,
		},
	)

	return qs
}

type notinHostAuthKey struct {
	values []interface{}
}

func (in *notinHostAuthKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"auth_key" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs HostQS) AuthKeyNotIn(values []string) HostQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinHostAuthKey{
			values: vals,
		},
	)

	return qs
}

// OrderByAuthKey sorts result by AuthKey in ascending order
func (qs HostQS) OrderByAuthKey() HostQS {
	qs.order = append(qs.order, `"auth_key"`)

	return qs
}

// OrderByAuthKeyDesc sorts result by AuthKey in descending order
func (qs HostQS) OrderByAuthKeyDesc() HostQS {
	qs.order = append(qs.order, `"auth_key" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs HostQS) ForUpdate() HostQS {
	qs.forUpdate = true

	return qs
}

func (qs HostQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs HostQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs HostQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "host_id", "dhcp_identifier", "dhcp_identifier_type", "dhcp4_subnet_id", "dhcp6_subnet_id", "ipv4_address", "hostname", "dhcp4_client_classes", "dhcp6_client_classes", "dhcp4_next_server", "dhcp4_server_hostname", "dhcp4_boot_file_name", "user_context", "auth_key" FROM "hosts"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs HostQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "host_id" FROM "hosts"` + s, p
}

// All returns all rows matching queryset filters
func (qs HostQS) All(db models.DBInterface) ([]*Host, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Host
	for rows.Next() {
		obj := Host{existsInDB: true}
		if err = rows.Scan(&obj.hostId, &obj.DhcpIdentifier, &obj.dhcpIdentifierType, &obj.Dhcp4SubnetId, &obj.Dhcp6SubnetId, &obj.Ipv4Address, &obj.Hostname, &obj.Dhcp4ClientClasses, &obj.Dhcp6ClientClasses, &obj.Dhcp4NextServer, &obj.Dhcp4ServerHostname, &obj.Dhcp4BootFileName, &obj.UserContext, &obj.AuthKey); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs HostQS) First(db models.DBInterface) (*Host, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Host{existsInDB: true}
	err := row.Scan(&obj.hostId, &obj.DhcpIdentifier, &obj.dhcpIdentifierType, &obj.Dhcp4SubnetId, &obj.Dhcp6SubnetId, &obj.Ipv4Address, &obj.Hostname, &obj.Dhcp4ClientClasses, &obj.Dhcp6ClientClasses, &obj.Dhcp4NextServer, &obj.Dhcp4ServerHostname, &obj.Dhcp4BootFileName, &obj.UserContext, &obj.AuthKey)
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
func (qs HostQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "hosts"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs HostQS) Update() HostUpdateQS {
	return HostUpdateQS{condFragments: qs.condFragments}
}

// HostUpdateQS represents an updated queryset for django_kea.Host
type HostUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs HostUpdateQS) update(c string, v interface{}) HostUpdateQS {
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
func (uqs HostUpdateQS) SetHostId(v int32) HostUpdateQS {
	return uqs.update(`"host_id"`, v)
}

// SetDhcpIdentifier sets DhcpIdentifier to the given value
func (uqs HostUpdateQS) SetDhcpIdentifier(v string) HostUpdateQS {
	return uqs.update(`"dhcp_identifier"`, v)
}

// SetDhcpIdentifierType sets foreign key pointer to Hostidentifiertype
func (uqs HostUpdateQS) SetDhcpIdentifierType(ptr *Hostidentifiertype) HostUpdateQS {
	if ptr != nil {
		return uqs.update(`"dhcp_identifier_type"`, ptr.Type)
	}

	return uqs.update(`"dhcp_identifier_type"`, nil)
} // SetDhcp4SubnetId sets Dhcp4SubnetId to the given value
func (uqs HostUpdateQS) SetDhcp4SubnetId(v sql.NullInt64) HostUpdateQS {
	return uqs.update(`"dhcp4_subnet_id"`, v)
}

// SetDhcp6SubnetId sets Dhcp6SubnetId to the given value
func (uqs HostUpdateQS) SetDhcp6SubnetId(v sql.NullInt64) HostUpdateQS {
	return uqs.update(`"dhcp6_subnet_id"`, v)
}

// SetIpv4Address sets Ipv4Address to the given value
func (uqs HostUpdateQS) SetIpv4Address(v sql.NullString) HostUpdateQS {
	return uqs.update(`"ipv4_address"`, v)
}

// SetHostname sets Hostname to the given value
func (uqs HostUpdateQS) SetHostname(v sql.NullString) HostUpdateQS {
	return uqs.update(`"hostname"`, v)
}

// SetDhcp4ClientClasses sets Dhcp4ClientClasses to the given value
func (uqs HostUpdateQS) SetDhcp4ClientClasses(v sql.NullString) HostUpdateQS {
	return uqs.update(`"dhcp4_client_classes"`, v)
}

// SetDhcp6ClientClasses sets Dhcp6ClientClasses to the given value
func (uqs HostUpdateQS) SetDhcp6ClientClasses(v sql.NullString) HostUpdateQS {
	return uqs.update(`"dhcp6_client_classes"`, v)
}

// SetDhcp4NextServer sets Dhcp4NextServer to the given value
func (uqs HostUpdateQS) SetDhcp4NextServer(v sql.NullString) HostUpdateQS {
	return uqs.update(`"dhcp4_next_server"`, v)
}

// SetDhcp4ServerHostname sets Dhcp4ServerHostname to the given value
func (uqs HostUpdateQS) SetDhcp4ServerHostname(v sql.NullString) HostUpdateQS {
	return uqs.update(`"dhcp4_server_hostname"`, v)
}

// SetDhcp4BootFileName sets Dhcp4BootFileName to the given value
func (uqs HostUpdateQS) SetDhcp4BootFileName(v sql.NullString) HostUpdateQS {
	return uqs.update(`"dhcp4_boot_file_name"`, v)
}

// SetUserContext sets UserContext to the given value
func (uqs HostUpdateQS) SetUserContext(v sql.NullString) HostUpdateQS {
	return uqs.update(`"user_context"`, v)
}

// SetAuthKey sets AuthKey to the given value
func (uqs HostUpdateQS) SetAuthKey(v sql.NullString) HostUpdateQS {
	return uqs.update(`"auth_key"`, v)
}

// Exec executes the update operation
func (uqs HostUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := HostQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "hosts" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (h *Host) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "hosts" ("dhcp_identifier", "dhcp_identifier_type", "dhcp4_subnet_id", "dhcp6_subnet_id", "ipv4_address", "hostname", "dhcp4_client_classes", "dhcp6_client_classes", "dhcp4_next_server", "dhcp4_server_hostname", "dhcp4_boot_file_name", "user_context", "auth_key") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING "host_id"`, h.DhcpIdentifier, h.dhcpIdentifierType, h.Dhcp4SubnetId, h.Dhcp6SubnetId, h.Ipv4Address, h.Hostname, h.Dhcp4ClientClasses, h.Dhcp6ClientClasses, h.Dhcp4NextServer, h.Dhcp4ServerHostname, h.Dhcp4BootFileName, h.UserContext, h.AuthKey)

	if err := row.Scan(&h.hostId); err != nil {
		return err
	}

	h.existsInDB = true

	return nil
}

// update operation
func (h *Host) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "hosts" SET "dhcp_identifier" = $1, "dhcp_identifier_type" = $2, "dhcp4_subnet_id" = $3, "dhcp6_subnet_id" = $4, "ipv4_address" = $5, "hostname" = $6, "dhcp4_client_classes" = $7, "dhcp6_client_classes" = $8, "dhcp4_next_server" = $9, "dhcp4_server_hostname" = $10, "dhcp4_boot_file_name" = $11, "user_context" = $12, "auth_key" = $13 WHERE "host_id" = $14`, h.DhcpIdentifier, h.dhcpIdentifierType, h.Dhcp4SubnetId, h.Dhcp6SubnetId, h.Ipv4Address, h.Hostname, h.Dhcp4ClientClasses, h.Dhcp6ClientClasses, h.Dhcp4NextServer, h.Dhcp4ServerHostname, h.Dhcp4BootFileName, h.UserContext, h.AuthKey, h.hostId)

	return err
}

// Save inserts or updates record
func (h *Host) Save(db models.DBInterface) error {
	if h.existsInDB {
		return h.update(db)
	}

	return h.insert(db)
}

// Delete removes row from database
func (h *Host) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "hosts" WHERE "host_id" = $1`, h.hostId)

	h.existsInDB = false

	return err
}

// Dhcp4option returns the set of Dhcp4option referencing this Host instance
func (h *Host) Dhcp4option() Dhcp4optionQS {
	return Dhcp4optionQS{}.HostEq(h)
}

// Dhcp6option returns the set of Dhcp6option referencing this Host instance
func (h *Host) Dhcp6option() Dhcp6optionQS {
	return Dhcp6optionQS{}.HostEq(h)
}

// Ipv6reservation returns the set of Ipv6reservation referencing this Host instance
func (h *Host) Ipv6reservation() Ipv6reservationQS {
	return Ipv6reservationQS{}.HostEq(h)
}
