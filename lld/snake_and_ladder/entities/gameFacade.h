#include<vector>
#include<set>
#include<iostream>
#include<time.h>
#include <cstdlib>
#include "Cell.h"
#include "Player.h"
#include "Dice.h"
#include "boardobject.h"

using namespace std;


class  GameFacade
{
    private:
        vector<vector<Cell>> board;
        vector<Player> players;
        vector<BoardObject*> boardObj;   
        Dice* boardDice;
        const int BOARD_LENGTH = 10;
        const int BOARD_HEIGHT = 10;
        const int SNAKE_COUNT = 3;
        const int LADDER_COUNT = 5;

    public:
      
       GameFacade(vector<string> playersName){
          
          // initial board
          for(int i=0;i<BOARD_HEIGHT;i++)
          {
             vector<Cell> row;
             for(int j=0;j<BOARD_LENGTH;j++)
             {
                 Cell cellObj = Cell(i,j);
                 row.push_back(cellObj); 
             }
             board.push_back(row);
          }

          // initial dice
          this->boardDice = Dice::getDiceInstance();

          // initial player
          for(int i=0;i<playersName.size();i++){
            players.push_back(Player(playersName[i], 0, 0));
          }

         srand(static_cast<unsigned>(time(nullptr)));

         
        
            set<pair<int,int>> usedStartCoordinates, usedEndCoordinates;
            for(pair<int,int> x: usedStartCoordinates){
                cout<<x.first <<" "<<x.second<<endl;
            }

             cout<<"Ladder Coordinate ..................."<<endl;
            // ladder
            for (int i = 0; i < LADDER_COUNT; ++i) {
                int startX, startY, endX, endY;

                do {
                    startX = rand() % BOARD_HEIGHT ;
                    startY = rand() % BOARD_LENGTH ;



                } while ((startX == 0 && startY == 0) || ( usedStartCoordinates.find({startX, startY}) != usedStartCoordinates.end()));

                usedStartCoordinates.insert({startX, startY});

                do {
                    endX = rand() % BOARD_HEIGHT;
                    endY = rand() % BOARD_LENGTH;

                } while (endY <= startY || endY == BOARD_LENGTH || usedStartCoordinates.find({endX, endY}) != usedStartCoordinates.end());

                usedEndCoordinates.insert({endX, endY});

                boardObj.push_back(new Ladder(Coordinates(startX, startY), Coordinates(endX, endY)));
                cout<<startX<< " " <<startY <<" "<<endX <<" "<<endY<<endl;
            }

            cout<<"Snake Coordinate ..................."<<endl;

        for (int i = 0; i < SNAKE_COUNT ; ++i) {
            int startX, startY, endX, endY; 

            do {
                startX = rand() % BOARD_HEIGHT;
                startY = rand() % BOARD_LENGTH;
            } while (!startX || (startX == BOARD_HEIGHT-1  && startY ==BOARD_LENGTH-1)  || usedStartCoordinates.find({startX, startY}) != usedStartCoordinates.end() || usedEndCoordinates.find({startX, startY}) != usedEndCoordinates.end());
            
            usedStartCoordinates.insert({startX,startY});

        do {
                endX = rand() % BOARD_HEIGHT ;
                endY = rand() % BOARD_LENGTH;
            } while ((endX >= startX) || usedStartCoordinates.find({endX, endY}) != usedStartCoordinates.end());

            boardObj.push_back(new Snake(Coordinates(startX, startY), Coordinates(endX, endY)));

             cout<<startX<< " " <<startY <<" "<<endX <<" "<<endY<<endl;
        }

       }

       void GameRun(){
        while(1){
         for(int i=0;i<players.size();i++)
         {
               cout<<"Player "<<players[i].getName()<<" throw dice"<<endl;
               cout<<"Your current coordinate: "<<endl;
               players[i].printLocation();
               int dicevalue = boardDice->rollDice();
               cout<<"Got the value: "<<dicevalue<<endl;
               players[i].updateLocation(increaseLocation(dicevalue, players[i].getLocation()));
               for(int j=0;j<boardObj.size();j++)
               {
                    players[i].updateLocation(boardObj[i]->updateCoordinate(players[i].getLocation()));
               }

               int xx = players[i].getLocation().x;
               int yy = players[i].getLocation().y;

               if(xx== BOARD_HEIGHT-1 && yy == BOARD_LENGTH-1)
                {
                     cout<<"Player "<<players[i].getName()<<" won the game";
                     return ;
                }
                cout<<"Updated the location after dice throw: "<<endl;
                players[i].printLocation();

                cout<<"........................................"<<endl;
               
         }
         }
       }
       Coordinates increaseLocation(int diceValue, Coordinates curr)
       {
                int x = curr.x ;
                int y = curr.y  + diceValue;
                if(x == BOARD_HEIGHT-1){
                    if(y > BOARD_LENGTH-1)
                    return curr;

                    return Coordinates(x,y);
                }else{
                    if(y > BOARD_LENGTH-1){
                        y = diceValue - (BOARD_LENGTH-1-curr.y) -1; 
                        x++;
                        return Coordinates(x,y);
                        
                    }
                    

                    return Coordinates(x,y);
                }
                
       }
};

