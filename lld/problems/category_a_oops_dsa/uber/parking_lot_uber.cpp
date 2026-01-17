/*

Design a parking lot system with the following specifications:

- The parking lot consists of multiple levels, and each level contains multiple rows of parking spots.
- There are two types of vehicles: motorcycles and cars.
- A motorcycle can park in any available parking spot, regardless of the spot type.
- A car can only park in a designated car spot.

Also implement park, unpark and search functions.




.... 

vehicle 
    - motorchcle
    - car 
parking 
    - level
        - spots m*n
            - car spo -> bike/car
            - bike spot -> bike 
    
-park()
-unpark()
-searhc() -> ?? , parking veichile spot as of now. , can also be free space with bike 

NOTE: this code took > 50 minutes 
better to say interview it will be interface, but creating as with type right now, like sport vehicle etc.



*/



#include<iostream> 
#include<vector>
#include<map>

using namespace std; 

enum class VehicleType {
    MOTORCYCLE, 
    CAR
};
class Vehicle {
    public: 
    VehicleType type; 
        string vehicleNumber; 
        
    Vehicle(){}
    Vehicle(string vehicleNumber, VehicleType type): vehicleNumber(vehicleNumber), type(type){}
        
};
class Motorcycle: public Vehicle {
    public: 
    
        Motorcycle(){}
        Motorcycle(string vehicleNumber): Vehicle(vehicleNumber, VehicleType::MOTORCYCLE) {}
};
class Car: public Vehicle {
    public: 
     
        Car(){}
        Car(string number): Vehicle(number, VehicleType::CAR) {}
};
enum class ParkingSpotType {
    MOTORCYCLE_SPOT, 
    CAR_SPOT
};
class ParkingSpot {
    public: 
        ParkingSpotType type; 
        int id; 
    ParkingSpot(){}
    ParkingSpot(int id, ParkingSpotType type): id(id), type(type) {}
        
};
class CarParkingSpot : public ParkingSpot{
    public: 
    
    CarParkingSpot(){}
    CarParkingSpot(int id): ParkingSpot(id, ParkingSpotType::CAR_SPOT)  {}
};
class BikeParkingSpot : public ParkingSpot{
    public: 
    ParkingSpotType type; 
    BikeParkingSpot(){}
    BikeParkingSpot(int id): ParkingSpot(id, ParkingSpotType::MOTORCYCLE_SPOT) {}
};
class ParkingLevel {
    public: 
        int level; 
        int spotWidth;
        int spotHeight;
        vector<vector<pair<ParkingSpot*, bool>>> spots; 
        ParkingLevel() {}
        ParkingLevel(int id,int h, int w): spotHeight(h), spotWidth(w), level(id) {}
};
class ParkingLot {
    public: 
    
        map<int, ParkingLevel*> levels; 
        map<string, ParkingSpot*> vehiclesParking;
        
        void initialise(int level, int spotL, int spotR, int bikeParkingCount, int CarParkingCount) { // dummy
            for(int i=0;i<level;i++){
                ParkingLevel *pl = new ParkingLevel(i, spotL, spotR); 
                int bikec = bikeParkingCount;
                int carc = CarParkingCount;
                for(int r=0;r<spotL;r++){
                    vector<pair<ParkingSpot*, bool>> spotRows;
                    for(int c=0;c<spotR;c++) {
                        if(bikec>0) {
                            spotRows.push_back({new BikeParkingSpot((r+1)*(c+1)), true});
                            bikec--;
                        } else {
                            spotRows.push_back({new CarParkingSpot((r+1)*(c+1)), true});
                            carc--;   
                        }
                    }
                    pl->spots.push_back(spotRows);
                }
                levels[i] = pl;
            }
        }
        
        void park(Vehicle* vehicle){
            for(int i=0;i<levels.size();i++) {
                bool found = false;    
                for(auto &spotRows: levels[i]->spots) {
                    for(auto &spot: spotRows) {
                        if(spot.second) {
                            
                            if((vehicle->type == VehicleType::CAR && spot.first->type == ParkingSpotType::CAR_SPOT ) || vehicle->type == VehicleType::MOTORCYCLE){
                                spot.second = false;
                                vehiclesParking[vehicle->vehicleNumber] = spot.first;
                                cout<<"parked at "<<spot.first->id<<" "<<endl;
                                found = true;
                                break;    
                            }
                        }
                    }
                    if(found) break;
                }
                if(found) break;
            }
            
            
            for(int i=0;i<levels.size();i++) {
                for(auto &spotRows: levels[i]->spots) {
                    for(auto &spot: spotRows) {
                        cout<<"("<<spot.first->id<<" "<<spot.second<<") ";        
                    }
                    cout<<endl;
                    
                }
                cout<<endl;
                
            }
            
        }
        
        void unpark(string vId) {
            ParkingSpot* ps=vehiclesParking[vId];
            vehiclesParking.erase(vId);
            for(int i=0;i<levels.size();i++) {
                for(auto &spotRows: levels[i]->spots) {
                    for(auto &spot: spotRows) {
                        if(spot.first->id == ps->id) {
                            spot.second = true;
                        }
                    }
                    cout<<endl;
                    
                }
                cout<<endl;   
            }
            
            for(int i=0;i<levels.size();i++) {
                for(auto &spotRows: levels[i]->spots) {
                    for(auto &spot: spotRows) {
                        cout<<"("<<spot.first->id<<" "<<spot.second<<") ";        
                    }
                    cout<<endl;
                    
                }
                cout<<endl;
                
            }
        }
    
    
};
int main() {
    ParkingLot *pl = new ParkingLot();
    pl->initialise(3, 5, 4, 10, 10);
    pl -> park( new Car("2"));
    pl -> park( new Motorcycle("1"));
    pl -> unpark( "2");
    return 0; 
}


/*
BY gpt there are few issues in this. 

1. unpark in o(n), what's the benefit of storing then ? 
    -> better to keep vehicle pointer ( see below)
2. no unique id in parking spot -> us global count at start of initilisation.
3. No sense of inheritance -> just wasting type -> if require tell to recruiter this can be interface and have multiple type. 
    -> because no polymorphis, you have same data in both side and also have same type of fucntion its not like area func which has different implements

    class Shape {
    public:
        virtual double area() = 0;
    };

    class Circle : public Shape {
    public:
        double area() override { ... }
    };


4. optimize park -> logn/1 -> maintain empty spot -> park rmeove -> unpark push again. 
    -> if not writing code atleast tell them. 
5. never delete pointer in c++ specially -> Memory Leaks -> redflag. -> if so tell before
    -> solution use normal type and reference instead pointer, or unique pointer.  see gpt solution
    -> destructor -> not need its default -> but inhertance -> virtual desctructor, or explicit delete like spot when exist 




a/c gpt -> if has to done in 40mins 



#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

enum class VehicleType { MOTORCYCLE, CAR };
enum class SpotType { MOTORCYCLE, CAR };

class Vehicle {
public:
    string number;
    VehicleType type;

    Vehicle(string n, VehicleType t) : number(n), type(t) {}
};

class ParkingSpot {
public:
    int id;
    SpotType type;
    Vehicle* vehicle;

    ParkingSpot(int id_, SpotType t)
        : id(id_), type(t), vehicle(nullptr) {}

    bool canPark(Vehicle* v) {
        if (vehicle != nullptr) return false;
        if (v->type == VehicleType::MOTORCYCLE) return true;
        return v->type == VehicleType::CAR && type == SpotType::CAR;
    }
};

class ParkingLevel {
public:
    int levelId;
    vector<vector<ParkingSpot>> spots;

    ParkingLevel(int id, int rows, int cols, int& globalId, int bikeSpots) 
        : levelId(id) {
        for (int r = 0; r < rows; r++) {
            vector<ParkingSpot> row;
            for (int c = 0; c < cols; c++) {
                if (bikeSpots-- > 0)
                    row.emplace_back(globalId++, SpotType::MOTORCYCLE);
                else
                    row.emplace_back(globalId++, SpotType::CAR);
            }
            spots.push_back(row);
        }
    }
};

class ParkingLot {
    vector<ParkingLevel> levels;
    unordered_map<string, ParkingSpot*> parked; // vehicle → spot

public:
    ParkingLot(int levelCount, int rows, int cols, int bikeSpotsPerLevel) {
        int globalId = 1;
        for (int i = 0; i < levelCount; i++) {
            levels.emplace_back(i, rows, cols, globalId, bikeSpotsPerLevel);
        }
    }

    bool park(Vehicle* v) {
        for (auto& level : levels) {
            for (auto& row : level.spots) {
                for (auto& spot : row) {
                    if (spot.canPark(v)) {
                        spot.vehicle = v;
                        parked[v->number] = &spot;
                        cout << "Parked " << v->number
                             << " at level " << level.levelId
                             << ", spot " << spot.id << endl;
                        return true;
                    }
                }
            }
        }
        return false;
    }

    bool unpark(string vehicleNo) {
        if (!parked.count(vehicleNo)) return false;
        ParkingSpot* spot = parked[vehicleNo];
        spot->vehicle = nullptr;
        parked.erase(vehicleNo);
        cout << "Unparked " << vehicleNo << endl;
        return true;
    }

    ParkingSpot* search(string vehicleNo) {
        if (parked.count(vehicleNo))
            return parked[vehicleNo];
        return nullptr;
    }
};

int main() {
    ParkingLot lot(2, 3, 4, 4);

    Vehicle* car = new Vehicle("CAR1", VehicleType::CAR);
    Vehicle* bike = new Vehicle("BIKE1", VehicleType::MOTORCYCLE);

    lot.park(car);
    lot.park(bike);

    auto s = lot.search("CAR1");
    if (s) cout << "CAR1 found at spot " << s->id << endl;

    lot.unpark("CAR1");
    return 0;
}



*/





// ..........

// code after fixing memory leak by gpt 


// #include <iostream>
// #include <vector>
// #include <map>

// using namespace std;

// /* ================= VEHICLES ================= */

// enum class VehicleType {
//     MOTORCYCLE,
//     CAR
// };

// class Vehicle {
// public:
//     VehicleType type;
//     string vehicleNumber;

//     Vehicle() {}
//     Vehicle(string vehicleNumber, VehicleType type)
//         : vehicleNumber(vehicleNumber), type(type) {}

//     virtual ~Vehicle() {}   // IMPORTANT
// };

// class Motorcycle : public Vehicle {
// public:
//     Motorcycle(string vehicleNumber)
//         : Vehicle(vehicleNumber, VehicleType::MOTORCYCLE) {}
// };

// class Car : public Vehicle {
// public:
//     Car(string number)
//         : Vehicle(number, VehicleType::CAR) {}
// };

// /* ================= PARKING SPOTS ================= */

// enum class ParkingSpotType {
//     MOTORCYCLE_SPOT,
//     CAR_SPOT
// };

// class ParkingSpot {
// public:
//     ParkingSpotType type;
//     int id;

//     ParkingSpot() {}
//     ParkingSpot(int id, ParkingSpotType type)
//         : id(id), type(type) {}

//     virtual ~ParkingSpot() {}   // IMPORTANT
// };

// class CarParkingSpot : public ParkingSpot {
// public:
//     CarParkingSpot(int id)
//         : ParkingSpot(id, ParkingSpotType::CAR_SPOT) {}
// };

// class BikeParkingSpot : public ParkingSpot {
// public:
//     BikeParkingSpot(int id)
//         : ParkingSpot(id, ParkingSpotType::MOTORCYCLE_SPOT) {}
// };

// /* ================= PARKING LEVEL ================= */

// class ParkingLevel {
// public:
//     int level;
//     int spotWidth;
//     int spotHeight;
//     vector<vector<pair<ParkingSpot*, bool>>> spots;

//     ParkingLevel() {}
//     ParkingLevel(int id, int h, int w)
//         : level(id), spotHeight(h), spotWidth(w) {}
// };

// /* ================= PARKING LOT ================= */

// class ParkingLot {
// public:
//     map<int, ParkingLevel*> levels;
//     map<string, ParkingSpot*> vehiclesParking;

//     ~ParkingLot() {
//         // delete all parking spots and levels
//         for (auto &lvl : levels) {
//             ParkingLevel* level = lvl.second;

//             for (auto &row : level->spots) {
//                 for (auto &spot : row) {
//                     delete spot.first;   // ParkingSpot*
//                 }
//             }
//             delete level;   // ParkingLevel*
//         }
//         levels.clear();
//         vehiclesParking.clear();
//     }

//     void initialise(int level, int spotL, int spotR,
//                     int bikeParkingCount, int carParkingCount) {

//         for (int i = 0; i < level; i++) {
//             ParkingLevel *pl = new ParkingLevel(i, spotL, spotR);

//             int bikec = bikeParkingCount;
//             int carc = carParkingCount;

//             for (int r = 0; r < spotL; r++) {
//                 vector<pair<ParkingSpot*, bool>> spotRows;

//                 for (int c = 0; c < spotR; c++) {
//                     if (bikec > 0) {
//                         spotRows.push_back(
//                             { new BikeParkingSpot((r+1)*(c+1)), true }
//                         );
//                         bikec--;
//                     } else {
//                         spotRows.push_back(
//                             { new CarParkingSpot((r+1)*(c+1)), true }
//                         );
//                         carc--;
//                     }
//                 }
//                 pl->spots.push_back(spotRows);
//             }
//             levels[i] = pl;
//         }
//     }

//     void park(Vehicle* vehicle) {
//         for (auto &lvl : levels) {
//             for (auto &row : lvl.second->spots) {
//                 for (auto &spot : row) {
//                     if (!spot.second) continue;

//                     if ((vehicle->type == VehicleType::CAR &&
//                          spot.first->type == ParkingSpotType::CAR_SPOT) ||
//                         (vehicle->type == VehicleType::MOTORCYCLE &&
//                          spot.first->type == ParkingSpotType::MOTORCYCLE_SPOT)) {

//                         spot.second = false;
//                         vehiclesParking[vehicle->vehicleNumber] = spot.first;
//                         cout << "Parked at " << spot.first->id << endl;
//                         return;
//                     }
//                 }
//             }
//         }
//     }

//     void unpark(const string &vId) {
//         if (!vehiclesParking.count(vId)) return;

//         ParkingSpot* ps = vehiclesParking[vId];
//         vehiclesParking.erase(vId);

//         for (auto &lvl : levels) {
//             for (auto &row : lvl.second->spots) {
//                 for (auto &spot : row) {
//                     if (spot.first == ps) {
//                         spot.second = true;
//                         cout << "Unparked from " << ps->id << endl;
//                         return;
//                     }
//                 }
//             }
//         }
//     }
// };

// /* ================= MAIN ================= */

// int main() {
//     ParkingLot *pl = new ParkingLot();

//     Vehicle* car = new Car("2");
//     Vehicle* bike = new Motorcycle("1");

//     pl->initialise(3, 5, 4, 10, 10);
//     pl->park(car);
//     pl->park(bike);
//     pl->unpark("2");

//     delete car;
//     delete bike;
//     delete pl;   // triggers full cleanup

//     return 0;
// }
