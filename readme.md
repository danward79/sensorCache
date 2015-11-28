# sensorcache

Provides a cache for sensor or device values.

Each device is stored with last time and reading, devices have a defined life, which can be checked for expiry. If a device reading has expired it is removed, from the cache.

sensorcache provides methods to allow updating, getting, deleting, checking expiry and automatic management of expiry.
