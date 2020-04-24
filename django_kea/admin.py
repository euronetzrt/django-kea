from django.contrib import admin
from django_kea import models


@admin.register(models.HostIdentifierType)
class HostIdentifierTypeAdmin(admin.ModelAdmin):
    list_display = ('type', 'name')


@admin.register(models.DhcpOptionScope)
class DhcpOptionScopeAdmin(admin.ModelAdmin):
    list_display = ('scope_id', 'scope_name')


@admin.register(models.Host)
class HostAdmin(admin.ModelAdmin):
    list_display = ('dhcp_identifier', 'dhcp_identifier_type',
                    'dhcp4_subnet_id', 'dhcp6_subnet_id',
                    'ipv4_address', 'hostname',
                    'dhcp4_client_classes', 'dhcp6_client_classes',
                    'dhcp4_next_server', 'dhcp4_server_hostname',
                    'dhcp4_boot_file_name', 'user_context')


@admin.register(models.Dhcp4Option)
class Dhcp4OptionAdmin(admin.ModelAdmin):
    list_display = ('code', 'value', 'formatted_value', 'space',
                    'persistent', 'dhcp_client_class', 'dhcp4_subnet_id',
                    'host', 'scope', 'user_context')


@admin.register(models.Dhcp6Option)
class Dhcp6OptionAdmin(admin.ModelAdmin):
    list_display = ('code', 'value', 'formatted_value', 'space',
                    'persistent', 'dhcp_client_class', 'dhcp6_subnet_id',
                    'host', 'scope', 'user_context')


@admin.register(models.IPv6Reservation)
class IPv6ReservationAdmin(admin.ModelAdmin):
    list_display = ('address', 'prefix_len', 'type','dhcp6_iaid', 'host')


@admin.register(models.LeaseState)
class LeaseStateAdmin(admin.ModelAdmin):
    list_display = ('state', 'name')


@admin.register(models.Lease4)
class Lease4Admin(admin.ModelAdmin):
    list_display = ('address', 'hwaddr', 'client_id',
                    'valid_lifetime', 'expire', 'subnet_id',
                    'fqdn_fwd', 'fqdn_rev', 'hostname',
                    'state', 'user_context')
    list_filter = ('subnet_id',)


@admin.register(models.Lease6Type)
class Lease6TypeAdmin(admin.ModelAdmin):
    list_display = ('lease_type', 'name')


@admin.register(models.LeaseHwaddrSource)
class LeaseHwaddrSourceAdmin(admin.ModelAdmin):
    list_display = ('hwaddr_source', 'name')


@admin.register(models.Lease6)
class Lease6Admin(admin.ModelAdmin):
    list_display = ('address', 'duid',
                    'valid_lifetime', 'expire', 'subnet_id',
                    'pref_lifetime', 'lease_type', 'iaid', 'prefix_len',
                    'fqdn_fwd', 'fqdn_rev', 'hostname',
                    'state', 'hwaddr', 'hwtype', 'hwaddr_source',
                    'user_context')
    list_filter = ('subnet_id',)
