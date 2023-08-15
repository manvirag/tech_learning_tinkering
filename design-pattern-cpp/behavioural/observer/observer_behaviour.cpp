#include <iostream>
#include <vector>

class Observer {
public:
    virtual void update(float temperature, float humidity, float pressure) = 0;
};


class Subject {
public:
    virtual void attach(Observer* observer) = 0;
    virtual void detach(Observer* observer) = 0;
    virtual void notifyObservers() = 0;
};

class WeatherMonitoringSystem : public Subject {
private:
    float temperature;
    float humidity;
    float pressure;
    std::vector<Observer*> observers;

public:
    WeatherMonitoringSystem() {}

    void attach(Observer* observer) override {
        observers.push_back(observer);
    }

    void detach(Observer* observer) override {
        for (auto it = observers.begin(); it != observers.end(); ++it) {
            if (*it == observer) {
                observers.erase(it);
                break;
            }
        }
    }

    void notifyObservers() override {
        for (Observer* observer : observers) {
            observer->update(temperature, humidity, pressure);
        }
    }

    void update(float newTemperature, float newHumidity, float newPressure) {
        temperature = newTemperature;
        humidity = newHumidity;
        pressure = newPressure;
        notifyObservers();
    }
};


class PhoneApp : public Observer {
public:
    void update(float temperature, float humidity, float pressure) override {
        std::cout << "Phone App: Temperature is " << temperature
                  << ", Humidity is " << humidity << ", Pressure is " << pressure << std::endl;
    }
};

class Website : public Observer {
public:
    void update(float temperature, float humidity, float pressure) override {
        std::cout << "Website: Temperature is " << temperature
                  << ", Humidity is " << humidity << ", Pressure is " << pressure << std::endl;
    }
};

class TV : public Observer {
public:
    void update(float temperature, float humidity, float pressure) override {
        std::cout << "TV: Temperature is " << temperature
                  << ", Humidity is " << humidity << ", Pressure is " << pressure << std::endl;
    }
};

int main() {
    WeatherMonitoringSystem weatherMonitoringSystem;
    PhoneApp phoneApp;
    Website website;
    TV tv;

    weatherMonitoringSystem.attach(&phoneApp);
    weatherMonitoringSystem.attach(&website);
    weatherMonitoringSystem.attach(&tv);

    weatherMonitoringSystem.update(30, 40, 50);

    return 0;
}
