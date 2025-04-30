from django.contrib import admin
from django_kea import models


@admin.register(models.Dhcp4Options)
class Dhcp4OptionsAdmin(admin.ModelAdmin):
    list_display = ('code', 'value', 'formatted_value', 'space',
                    'persistent', 'dhcp_client_class', 'dhcp4_subnet_id',
                    'host', 'scope', 'user_context', 'modification_ts', 'cancelled')


@admin.register(models.Dhcp6Options)
class Dhcp6OptionsAdmin(admin.ModelAdmin):
    list_display = ('code', 'value', 'formatted_value', 'space',
                    'persistent', 'dhcp_client_class', 'dhcp6_subnet_id',
                    'host', 'scope', 'user_context', 'cancelled')


@admin.register(models.Hosts)
class HostAdmin(admin.ModelAdmin):
    list_display = ('dhcp_identifier', 'dhcp_identifier_type',
                    'dhcp4_subnet_id', 'dhcp6_subnet_id',
                    'ipv4_address', 'hostname',
                    'dhcp4_client_classes', 'dhcp6_client_classes',
                    'dhcp4_next_server', 'dhcp4_server_hostname',
                    'dhcp4_boot_file_name', 'user_context',
                    'auth_key')


@admin.register(models.Ipv6Reservations)
class Ipv6ReservationsAdmin(admin.ModelAdmin):
    list_display = ('address', 'prefix_len', 'type','dhcp6_iaid', 'host')


@admin.register(models.Lease4)
class Lease4Admin(admin.ModelAdmin):
    list_display = ('address', 'hwaddr', 'client_id',
                    'valid_lifetime', 'expire', 'subnet_id',
                    'fqdn_fwd', 'fqdn_rev', 'hostname',
                    'state', 'user_context',
                    'relay_id', 'remote_id', 'pool_id')
    list_filter = ('subnet_id',)


@admin.register(models.Lease6)
class Lease6Admin(admin.ModelAdmin):
    list_display = ('address', 'duid',
                    'valid_lifetime', 'expire', 'subnet_id',
                    'pref_lifetime', 'lease_type', 'iaid', 'prefix_len',
                    'fqdn_fwd', 'fqdn_rev', 'hostname',
                    'state', 'hwaddr', 'hwtype', 'hwaddr_source',
                    'user_context', 'pool_id')
    list_filter = ('subnet_id',)
