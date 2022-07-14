package elon

import (
    "fmt"
    "strconv"
)

// TODO: define the 'Drive()' method

func (c *Car) Drive() {
    if c.battery > c.batteryDrain{
        c.distance += c.speed
        c.battery -= c.batteryDrain
    }
    
}

// TODO: define the 'DisplayDistance() string' method

func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method

func (c *Car) DisplayBattery() string {
    bat := strconv.Itoa(c.battery)
    return "Battery at " + bat + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method

func (c *Car) CanFinish(trackDistance int) bool {
    return c.battery / c.batteryDrain * c.speed >= trackDistance
}
