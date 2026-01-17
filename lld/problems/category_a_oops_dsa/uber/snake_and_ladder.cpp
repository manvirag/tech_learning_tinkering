/*
Was asked Snake and Ladder. Be proficient with SOLID design principals.

board -> 0 - 100 -> x,y 10*10 
    - enityt
        - snake
            -> start location -> end location , location -> x,y
        - ladder
            -> start -> end 
    - user -> 2 player 
        location
    - dice -> [1-6]

player -> 0 - 100 with dice 
one by one 
6 -> again chance to throw dice ->?? do we have imit here ? or unilited 5 limit for safer side . 

user reach snake/latter -( start point) -> move it o end point. 

whose event reach 100 -> first -> win 

if move go out > 100 -> no move 

flow 
    - one by one user run until one reach to 100 
    - get dice value 
        -> 6 -> run again -> at max 5 , 6 
    - move their location 
    - change location if snake/ladder 
    


NOTE took > 45 min


*/

/*
91................100
11................20
1 2 3 4 5 6 7 8 9 10
*/
#include<iostream> 
#include<map> 
#include<random> 
using namespace std; 

random_device rd;
mt19937 rng(rd());
class Location {
    public: 
        int y; // 1 - 10 
        int x; // 1 - 10
    Location(){}
    Location(int x, int y): x(x), y(y) {}
    bool operator<(const Location &a) const {
        return true;
    }
    bool operator==(const Location &a) const {
        return a.x == x && a.y==y;
    }
};
enum class BoardEntityType{
    SNAKE, LADDER
};
class BoardEntity {
    public: 
    BoardEntityType type; 
    Location start; 
    Location end; 
    BoardEntity(){}
    BoardEntity(Location s, Location e, BoardEntityType t): type(t), start(s), end(e) {}
};
class Player {
    public:    
        Location pos; 
        string name;
        Player(Location pos, string name): pos(pos), name(name) {}
};
class BoardSpot {
    public: 
        int x;
        int y;  // location -> x*(y-1) + y -> 1 - 100
    BoardSpot(){}
    BoardSpot(int x, int y): x(x), y(y) {}
};
class Board {
    public: 
        int length; 
        int width; 
        int maxSix;
        vector<vector<BoardSpot>> board; 
        map<Location, BoardEntity> entites; 
    Board(int l, int w, vector<BoardEntity> entityList, int maxSix): length(l), width(w) , maxSix(maxSix) {
        board.resize(l, vector<BoardSpot>(w));
        for(int i=0;i<l;i++) {
            for(int j=0;j<w;j++) {
                board[i][j] = BoardSpot(i+1,j+1);
            }
        }
        for(auto entity: entityList) {
            entites[entity.start] = entity; 
        }

    }
    int runDice() {
        uniform_int_distribution<int> a(1,6);
        return a(rng);
    }
    void updateLocation(Player &p, int diceValue) {
            Location lc = p.pos;
            int score = (lc.x-1) * width + lc.y;
            cout<<score<<endl;
            score += diceValue;
            cout<<score<<endl; 
            if(score > length*width) {
                return; 
            } 
            lc.x = score/width + (score < length*width);
            lc.y = score-(lc.x-1)*width;
            cout<<lc.y<<endl;
            p.pos = lc;

            if(entites.find(p.pos)!=entites.end()) {
                p.pos = entites[p.pos].end;
            }
    }
    bool runPlayer(Player &a) {
        for(int i=1;i<=maxSix;i++) {
            cout<<a.name<<" "<<"throwing dice"<<endl;
            int dc = runDice();
            cout<<"Got "<<dc<<endl;
            updateLocation(a, dc); 
            cout<<"initial location "<<a.pos.x<<" "<<a.pos.y<<endl;
            cout<<"update location "<<a.pos.x<<" "<<a.pos.y<<endl;
            if(a.pos.x == length && a.pos.y == width) {
                cout<<a.name<<" won the game"<<endl;
                return true;
            }
            if(dc != 6 ) break;
        }
        return false;
            
    }
    void RunFacade(Player a, Player b) { 
        cout<<a.name <<" "<<b.name<<endl;

        while(1){
            if(runPlayer(a))break;
            if(runPlayer(b)) break; 
        }

    }

};
int main() {
    vector<BoardEntity> entities;
    entities.emplace_back(Location(2,4),  Location(4,5) ,BoardEntityType::LADDER);
    entities.emplace_back(Location(1,5),  Location(4,5) ,BoardEntityType::LADDER);
    entities.emplace_back(Location(6,7),  Location(3,5) ,BoardEntityType::SNAKE);
    entities.emplace_back(Location(4,8),  Location(1,5) ,BoardEntityType::SNAKE);
    Board board = Board(10,10, entities, 5);
    Player a = Player(Location(1,1), "a");
    Player b = Player(Location(1,1), "b");
    board.RunFacade(a, b);

    return 0;
}


/*
feedback
- complecated with x,y -> linear score. 
- memory leak is fine
- why do we need this -> BoardEntityType - >readability
- < operation incorrect  ( knownlingy )

lc.x = 10 / 10 + (10 < 100)
     = 1 + 1
     = 2   ❌





#include <iostream>
#include <vector>
#include <map>
#include <random>
using namespace std;

// Random dice generator
random_device rd;
mt19937 rng(rd());

class BoardEntity {
public:
    int start; // 1..100
    int end;
    enum class Type { SNAKE, LADDER } type;

    BoardEntity(int s, int e, Type t) : start(s), end(e), type(t) {}
};

class Player {
public:
    int score; // 1..100
    string name;

    Player(string n) : score(1), name(n) {}
};

class Board {
    int length; // number of cells (10x10)
    int maxSix; // max consecutive rolls if 6
    map<int, BoardEntity> entities; // start position → entity
public:
    Board(vector<BoardEntity> entityList, int maxSixRolls)
        : length(100), maxSix(maxSixRolls)
    {
        for (auto &e : entityList)
            entities[e.start] = e;
    }

    int rollDice() {
        uniform_int_distribution<int> dist(1, 6);
        return dist(rng);
    }

    void movePlayer(Player &p, int dice) {
        int newScore = p.score + dice;
        if (newScore > length)
            return; // can't go past last cell
        p.score = newScore;

        // Check snakes or ladders
        if (entities.find(p.score) != entities.end()) {
            p.score = entities[p.score].end;
        }
    }

    bool runTurn(Player &p) {
        for (int i = 0; i < maxSix; i++) {
            cout << p.name << " rolling dice..." << endl;
            int dice = rollDice();
            cout << "Got " << dice << endl;

            movePlayer(p, dice);
            cout << p.name << " at position " << p.score << endl;

            if (p.score == length) {
                cout << p.name << " won the game!" << endl;
                return true;
            }

            if (dice != 6) break; // only extra turn if dice == 6
        }
        return false;
    }

    void playGame(Player &p1, Player &p2) {
        while (true) {
            if (runTurn(p1)) break;
            if (runTurn(p2)) break;
        }
    }
};

int main() {
    vector<BoardEntity> entities = {
        BoardEntity(8, 22, BoardEntity::Type::LADDER),
        BoardEntity(5, 25, BoardEntity::Type::LADDER),
        BoardEntity(67, 35, BoardEntity::Type::SNAKE),
        BoardEntity(78, 15, BoardEntity::Type::SNAKE)
    };

    Board board(entities, 5);

    Player a("Alice");
    Player b("Bob");

    board.playGame(a, b);

    return 0;
}

*/