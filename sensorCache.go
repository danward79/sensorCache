package sensorCache

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

//Cache provides a data store for sensor device readings
type Cache struct {
	life        time.Duration
	lastReading time.Time
	readings    map[string]interface{}
	time        map[string]time.Time
	mutex       sync.RWMutex
	done        chan bool
}

//String provides information about the cache
func (c *Cache) String() string {
	return fmt.Sprintf("Cache Details - Last Value Inserted: %s, Life: %s", c.lastReading, c.life)
}

//New start a new cache
func New(t time.Duration) *Cache {
	c := &Cache{
		life:     t,
		readings: make(map[string]interface{}),
		time:     make(map[string]time.Time),
		done:     make(chan bool),
	}

	log.Println(c)
	return c
}

//Insert update value in cache, if device does not exist create the device
func (c *Cache) Insert(device string, reading interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.readings[device] = reading

	t := time.Now()
	c.time[device] = t
	c.lastReading = t

	return nil
}

//Delete device and its readings from the cache
func (c *Cache) Delete(device string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, OK := c.readings[device]; !OK {
		return errors.New("Error: Invalid device")
	}
	delete(c.readings, device)

	if _, OK := c.time[device]; !OK {
		return errors.New("Error: Invalid device")
	}
	delete(c.time, device)

	return nil
}

//Expire check for expired values and delete
func (c *Cache) Expire() error {
	for device, lastReading := range c.time {
		age := time.Since(lastReading)
		if age >= c.life {
			err := c.Delete(device)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//MonitorExpiry automatically manage cache for expired values
func (c *Cache) MonitorExpiry(d time.Duration) {

	t := time.NewTicker(d)

	for {
		select {
		case <-t.C:
			c.Expire()
		case <-c.done:
			log.Println("sensorCache: Cache monitoring stopped")
			return
		}
	}
}

//StopMonitoring signal that MonitorExpiry should exist
func (c *Cache) StopMonitoring() {
	c.done <- true
}

//Values returns current values from cache
func (c *Cache) Values() map[string]interface{} {
	return c.readings
}
