package main

import "fmt"

// Observer defines the interface for updating observers.
type Observer interface {
	Update(temperature, humidity, pressure float64)
}

// Subject defines the interface for attaching, detaching, and notifying observers.
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	NotifyObservers()
}

// WeatherMonitoringSystem is a concrete implementation of Subject.
type WeatherMonitoringSystem struct {
	temperature float64
	humidity    float64
	pressure    float64
	observers   []Observer
}

func (w *WeatherMonitoringSystem) Attach(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherMonitoringSystem) Detach(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherMonitoringSystem) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherMonitoringSystem) Update(temperature, humidity, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.NotifyObservers()
}

// PhoneApp, Website, and TV are concrete implementations of Observer.
type PhoneApp struct{}

func (p *PhoneApp) Update(temperature, humidity, pressure float64) {
	fmt.Printf("Phone App: Temperature is %.2f, Humidity is %.2f, Pressure is %.2f\n", temperature, humidity, pressure)
}

type Website struct{}

func (w *Website) Update(temperature, humidity, pressure float64) {
	fmt.Printf("Website: Temperature is %.2f, Humidity is %.2f, Pressure is %.2f\n", temperature, humidity, pressure)
}

type TV struct{}

func (t *TV) Update(temperature, humidity, pressure float64) {
	fmt.Printf("TV: Temperature is %.2f, Humidity is %.2f, Pressure is %.2f\n", temperature, humidity, pressure)
}

func main() {
	weatherMonitoringSystem := &WeatherMonitoringSystem{}
	phoneApp := &PhoneApp{}
	website := &Website{}
	tv := &TV{}

	weatherMonitoringSystem.Attach(phoneApp)
	weatherMonitoringSystem.Attach(website)
	weatherMonitoringSystem.Attach(tv)

	weatherMonitoringSystem.Update(30, 40, 50)
}
