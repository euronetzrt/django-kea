# Django-kea

Package for ISC KEA models. Models are registered to Django Admin site. Only tested with PostgreSQL.

## Usage

Add `django_kea.ISCKea` to `INSTALLED_APPS`.

## Separate database for KEA

If Kea's database resides in a different database, then create a database configuration named `kea` in `DATABASES`, and use `django_kea` dbrouter in settings:

```python
DATABASE_ROUTERS = ['django_kea.dbrouter.Router']
```
