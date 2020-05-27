# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey has `on_delete` set to the desired behavior.
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models
from django_kea import fields


class Dhcp4Option(models.Model):
    option_id = models.AutoField(primary_key=True)
    code = models.SmallIntegerField()
    value = fields.Binary(blank=True, null=True)
    formatted_value = models.TextField(blank=True, null=True)
    space = models.CharField(max_length=128, blank=True, null=True)
    persistent = models.BooleanField()
    dhcp_client_class = models.CharField(max_length=128, blank=True, null=True)
    dhcp4_subnet_id = models.BigIntegerField(blank=True, null=True)
    host = models.ForeignKey('Host', models.DO_NOTHING, blank=True, null=True)
    scope = models.ForeignKey('DhcpOptionScope', models.DO_NOTHING)
    user_context = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'dhcp4_options'


class Dhcp6Option(models.Model):
    option_id = models.AutoField(primary_key=True)
    code = models.IntegerField()
    value = fields.Binary(blank=True, null=True)
    formatted_value = models.TextField(blank=True, null=True)
    space = models.CharField(max_length=128, blank=True, null=True)
    persistent = models.BooleanField()
    dhcp_client_class = models.CharField(max_length=128, blank=True, null=True)
    dhcp6_subnet_id = models.BigIntegerField(blank=True, null=True)
    host = models.ForeignKey('Host', models.DO_NOTHING, blank=True, null=True)
    scope = models.ForeignKey('DhcpOptionScope', models.DO_NOTHING)
    user_context = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'dhcp6_options'


class DhcpOptionScope(models.Model):
    scope_id = models.SmallIntegerField(primary_key=True)
    scope_name = models.CharField(max_length=32, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'dhcp_option_scope'

    def __str__(self):
        return self.scope_name or str(self.scope_id)


class HostIdentifierType(models.Model):
    type = models.SmallIntegerField(primary_key=True)
    name = models.CharField(max_length=32, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'host_identifier_type'

    def __str__(self):
        return self.name or str(self.type)


class Host(models.Model):
    host_id = models.AutoField(primary_key=True)
    dhcp_identifier = fields.Binary()
    dhcp_identifier_type = models.ForeignKey(HostIdentifierType, models.DO_NOTHING, db_column='dhcp_identifier_type')
    dhcp4_subnet_id = models.BigIntegerField(blank=True, null=True)
    dhcp6_subnet_id = models.BigIntegerField(blank=True, null=True)
    ipv4_address = fields.IPv4Address(blank=True, null=True)
    hostname = models.CharField(max_length=255, blank=True, null=True)
    dhcp4_client_classes = models.CharField(max_length=255, blank=True, null=True)
    dhcp6_client_classes = models.CharField(max_length=255, blank=True, null=True)
    dhcp4_next_server = fields.IPv4Address(blank=True, null=True)
    dhcp4_server_hostname = models.CharField(max_length=64, blank=True, null=True)
    dhcp4_boot_file_name = models.CharField(max_length=128, blank=True, null=True)
    user_context = models.TextField(blank=True, null=True)
    auth_key = models.CharField(max_length=32, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'hosts'
        unique_together = (('dhcp_identifier', 'dhcp_identifier_type', 'dhcp4_subnet_id'), ('ipv4_address', 'dhcp4_subnet_id'), ('dhcp_identifier', 'dhcp_identifier_type', 'dhcp6_subnet_id'),)


class Ipv6Reservation(models.Model):
    reservation_id = models.AutoField(primary_key=True)
    address = models.CharField(max_length=39)
    prefix_len = models.SmallIntegerField()
    type = models.SmallIntegerField()
    dhcp6_iaid = models.IntegerField(blank=True, null=True)
    host = models.ForeignKey(Host, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'ipv6_reservations'
        unique_together = (('address', 'prefix_len'),)
        verbose_name = 'IPv6 reservation'


class Lease4(models.Model):
    address = fields.IPv4Address(primary_key=True)
    hwaddr = fields.Binary(blank=True, null=True)
    client_id = fields.Binary(blank=True, null=True)
    valid_lifetime = models.BigIntegerField(blank=True, null=True)
    expire = models.DateTimeField(blank=True, null=True)
    subnet_id = models.BigIntegerField(blank=True, null=True)
    fqdn_fwd = models.BooleanField(blank=True, null=True)
    fqdn_rev = models.BooleanField(blank=True, null=True)
    hostname = models.CharField(max_length=255, blank=True, null=True)
    state = models.ForeignKey('LeaseState', models.DO_NOTHING, db_column='state', blank=True, null=True)
    user_context = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'lease4'


# # Lease4Stat has a composite primary key
# class Lease4Stat(models.Model):
#     subnet_id = models.BigIntegerField(primary_key=True)
#     state = models.BigIntegerField()
#     leases = models.BigIntegerField(blank=True, null=True)

#     class Meta:
#         managed = False
#         db_table = 'lease4_stat'
#         unique_together = (('subnet_id', 'state'),)


class Lease6(models.Model):
    address = models.CharField(primary_key=True, max_length=39)
    duid = fields.Binary(blank=True, null=True)
    valid_lifetime = models.BigIntegerField(blank=True, null=True)
    expire = models.DateTimeField(blank=True, null=True)
    subnet_id = models.BigIntegerField(blank=True, null=True)
    pref_lifetime = models.BigIntegerField(blank=True, null=True)
    lease_type = models.ForeignKey('Lease6Type', models.DO_NOTHING, db_column='lease_type', blank=True, null=True)
    iaid = models.IntegerField(blank=True, null=True)
    prefix_len = models.SmallIntegerField(blank=True, null=True)
    fqdn_fwd = models.BooleanField(blank=True, null=True)
    fqdn_rev = models.BooleanField(blank=True, null=True)
    hostname = models.CharField(max_length=255, blank=True, null=True)
    state = models.ForeignKey('LeaseState', models.DO_NOTHING, db_column='state', blank=True, null=True)
    hwaddr = fields.Binary(blank=True, null=True)
    hwtype = models.SmallIntegerField(blank=True, null=True)
    hwaddr_source = models.SmallIntegerField(blank=True, null=True)
    user_context = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'lease6'


# # Lease6Stat has a composite primary key
# class Lease6Stat(models.Model):
#     subnet_id = models.BigIntegerField(primary_key=True)
#     lease_type = models.SmallIntegerField()
#     state = models.BigIntegerField()
#     leases = models.BigIntegerField(blank=True, null=True)

#     class Meta:
#         managed = False
#         db_table = 'lease6_stat'
#         unique_together = (('subnet_id', 'lease_type', 'state'),)


class Lease6Type(models.Model):
    lease_type = models.SmallIntegerField(primary_key=True)
    name = models.CharField(max_length=5, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'lease6_types'

    def __str__(self):
        return self.name or str(self.lease_type)


class LeaseHwaddrSource(models.Model):
    hwaddr_source = models.IntegerField(primary_key=True)
    name = models.CharField(max_length=40, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'lease_hwaddr_source'

    def __str__(self):
        return self.name or str(self.hwaddr_source)


class LeaseState(models.Model):
    state = models.BigIntegerField(primary_key=True)
    name = models.CharField(max_length=64)

    class Meta:
        managed = False
        db_table = 'lease_state'

    def __str__(self):
        return self.name


# # Log does not have a Primary Key
# class Log(models.Model):
#     timestamp = models.DateTimeField(blank=True, null=True)
#     address = models.CharField(max_length=43, blank=True, null=True)
#     log = models.TextField()

#     class Meta:
#         managed = False
#         db_table = 'logs'


class SchemaVersion(models.Model):
    version = models.IntegerField(primary_key=True)
    minor = models.IntegerField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'schema_version'
