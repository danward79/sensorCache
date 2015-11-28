# sensorcache

Provides a cache for sensor or device values.

Each device is stored with the last time and reading, devices have a defined life, which can be checked for expiry. If a device reading has expired it is removed, from the cache.

sensorcache provides methods to allow updating, getting, deleting, checking expiry and automatic management of expiry.

### An Example

A working example of this code can be found [here](https://github.com/danward79/wuMQTTAgregate) in wuMQTTAgregate which is an example of how to publish dispersed sensor readings using the cache and MQTT
