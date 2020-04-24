""" Kea fields """

import struct
import socket
from django.db import models


class IPv4Address(models.Field):
    description = "IPv4 address stored as a bigint"

    def db_type(self, connection):
        return "bigint"

    def from_db_value(self, value, expression, connection):
        if value is not None:
            return socket.inet_ntoa(struct.pack('!L', value))

        return ''

    def get_prep_value(self, value):
        if value != '':
            return struct.unpack('!L', socket.inet_aton(value))

        return None


class Binary(models.BinaryField):
    description = "Binary data represented"

    def __init__(self, *args, **kwargs):
        kwargs['editable'] = True
        super().__init__(*args, **kwargs)

    def deconstruct(self):
        name, path, args, kwargs = super().deconstruct()
        del kwargs['editable']
        return name, path, args, kwargs

    def from_db_value(self, value, expression, connection):
        if value is not None:
            return bytes(value).hex()

        return None

    def to_python(self, value):
        return value

    def get_prep_value(self, value):
        if value is not None:
            return bytes.fromhex(value)

        return None
