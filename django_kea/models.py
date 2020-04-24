from django.db import models
from django_kea import fields


class HostIdentifierType(models.Model):
    type = models.SmallIntegerField(primary_key=True)
    name = models.CharField(max_length=32, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'host_identifier_type'

    def __str__(self):
        return self.name or str(self.typ)


class DhcpOptionScope(models.Model):
    scope_id = models.SmallIntegerField(primary_key=True)
    scope_name = models.CharField(max_length=32, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'dhcp_option_scope'

    def __str__(self):
        return self.scope_name or str(self.scope_id)


class Host(models.Model):
    host_id = models.AutoField(primary_key=True)
    dhcp_identifier = fields.Binary()
    dhcp_identifier_type = models.ForeignKey(HostIdentifierType, db_column='dhcp_identifier_type', on_delete=models.CASCADE)
    dhcp4_subnet_id = models.BigIntegerField(null=True, blank=True)
    dhcp6_subnet_id = models.BigIntegerField(null=True, blank=True)
    ipv4_address = fields.IPv4Address(null=True, blank=True)
    hostname = models.CharField(max_length=255, null=True, blank=True)
    dhcp4_client_classes = models.CharField(max_length=255, null=True, blank=True)
    dhcp6_client_classes = models.CharField(max_length=255, null=True, blank=True)
    dhcp4_next_server = fields.IPv4Address(null=True, blank=True)
    dhcp4_server_hostname = models.CharField(max_length=64, null=True, blank=True)
    dhcp4_boot_file_name = models.CharField(max_length=128, null=True, blank=True)
    user_context = models.TextField(null=True, blank=True)
    auth_key = models.CharField(max_length=32, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'hosts'


class Dhcp4Option(models.Model):
    option_id = models.AutoField(primary_key=True)
    code = models.SmallIntegerField()
    value = fields.Binary(null=True, blank=True)
    formatted_value = models.TextField(null=True, blank=True)
    space = models.CharField(max_length=128, null=True, blank=True)
    persistent = models.BooleanField(default=False)
    dhcp_client_class = models.CharField(max_length=128, null=True, blank=True)
    dhcp4_subnet_id = models.BigIntegerField(null=True, blank=True)
    host = models.ForeignKey(Host, on_delete=models.CASCADE, null=True, blank=True)
    scope = models.ForeignKey(DhcpOptionScope, on_delete=models.CASCADE, null=True, blank=True)
    user_context = models.TextField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'dhcp4_options'


class Dhcp6Option(models.Model):
    option_id = models.AutoField(primary_key=True)
    code = models.SmallIntegerField()
    value = fields.Binary(null=True, blank=True)
    formatted_value = models.TextField(null=True, blank=True)
    space = models.CharField(max_length=128, null=True, blank=True)
    persistent = models.BooleanField(default=False)
    dhcp_client_class = models.CharField(max_length=128, null=True, blank=True)
    dhcp6_subnet_id = models.BigIntegerField(null=True, blank=True)
    host = models.ForeignKey(Host, on_delete=models.CASCADE, null=True, blank=True)
    scope = models.ForeignKey(DhcpOptionScope, on_delete=models.CASCADE, null=True, blank=True)
    user_context = models.TextField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'dhcp6_options'


class IPv6Reservation(models.Model):
    reservation_id = models.AutoField(primary_key=True)
    address = models.CharField(max_length=39)
    prefix_len = models.SmallIntegerField(default=128)
    type = models.SmallIntegerField(default=0)
    dhcp6_iaid = models.IntegerField(null=True, blank=True)
    host = models.ForeignKey(Host, on_delete=models.CASCADE)

    class Meta:
        managed = False
        db_table = 'ipv6_reservations'
        verbose_name = 'IPv6 reservation'


class LeaseState(models.Model):
    state = models.BigIntegerField(primary_key=True)
    name = models.CharField(max_length=64)

    class Meta:
        managed = False
        db_table = 'lease_state'

    def __str__(self):
        return self.name


class Lease4(models.Model):
    address = fields.IPv4Address(primary_key=True)
    hwaddr = fields.Binary(null=True, blank=True)
    client_id = fields.Binary(null=True, blank=True)
    valid_lifetime = models.BigIntegerField(null=True, blank=True)
    expire = models.DateTimeField(null=True, blank=True)
    subnet_id = models.BigIntegerField(null=True, blank=True)
    fqdn_fwd = models.NullBooleanField()
    fqdn_rev = models.NullBooleanField()
    hostname = models.CharField(max_length=255, null=True, blank=True)
    state = models.ForeignKey(LeaseState, db_column='state', on_delete=models.DO_NOTHING, null=True, blank=True)
    user_context = models.TextField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'lease4'


class Lease6Type(models.Model):
    lease_type = models.SmallIntegerField(primary_key=True)
    name = models.CharField(max_length=5, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'lease6_types'


class LeaseHwaddrSource(models.Model):
    hwaddr_source = models.IntegerField(primary_key=True)
    name = models.CharField(max_length=40, null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'lease_hwaddr_source'


class Lease6(models.Model):
    address = models.CharField(max_length=39, primary_key=True)
    duid = fields.Binary(null=True, blank=True)
    valid_lifetime = models.BigIntegerField(null=True, blank=True)
    expire = models.DateTimeField(null=True, blank=True)
    subnet_id = models.BigIntegerField(null=True, blank=True)
    pref_lifetime = models.BigIntegerField(null=True, blank=True)
    lease_type = models.SmallIntegerField(null=True, blank=True)
    iaid = models.IntegerField(null=True, blank=True)
    prefix_len = models.SmallIntegerField(null=True, blank=True)
    fqdn_fwd = models.NullBooleanField()
    fqdn_rev = models.NullBooleanField()
    hostname = models.CharField(max_length=255, null=True, blank=True)
    state = models.ForeignKey(LeaseState, db_column='state', on_delete=models.DO_NOTHING, null=True, blank=True)
    hwaddr = fields.Binary(null=True, blank=True)
    hwtype = models.SmallIntegerField(null=True, blank=True)
    hwaddr_source = models.SmallIntegerField(null=True, blank=True)
    user_context = models.TextField(null=True, blank=True)

    class Meta:
        managed = False
        db_table = 'lease6'
