#include <iostream>
#include <iomanip>
#include <limits>
#include <cmath>
#include <unordered_set>
#include <set>
#include <list>
#include <map>
using namespace std;

class User {
    public:
    int userId;
    list<int> recentPlayList;
    map<int, list<int>::iterator> recentPlayListIterator;
    User(int userId) {
        this->userId = userId;
        this->recentPlayList = list<int>();
        this->recentPlayListIterator = map<int, list<int>::iterator>();
    }

    void updateUserSongList(int songId) {
        if(recentPlayListIterator.find(songId) == recentPlayListIterator.end()) {
            recentPlayList.push_front(songId);
            recentPlayListIterator[songId] = recentPlayList.begin();
        } else {
            recentPlayList.erase(recentPlayListIterator[songId]);
            recentPlayList.push_front(songId);
            recentPlayListIterator[songId] = recentPlayList.begin();
        }
    }
};
class Song {
    public:
    int songId;
    string title;
    bool isStarred;
    unordered_set<int> uniqueUsers;
    
    Song() = default;
    Song(int songId, string title) {
        this->songId = songId;
        this->title = title;
        this->isStarred = false;
        this->uniqueUsers = unordered_set<int>();
    }
    
};
class MusicPlayerAnalytics {
    public: 
       static int songSequenceId;
       map<int , Song*> songMap;
       map<int, User*> userMap;
       set<pair<int, int>> songPlayCountSet;



    int addSong(string title) {  // constant
        songSequenceId++;
        Song* newSong = new Song(songSequenceId, title);
        songMap[songSequenceId] = newSong;
        return songSequenceId;
    }
    void addUser(int userId) {
        User* newUser = new User(userId);
        userMap[userId] = newUser;
    }

    void playSong(int songId, int userId) {  
        
        if(songMap[songId]->uniqueUsers.find(userId) == songMap[songId]->uniqueUsers.end()) {
            songMap[songId]->uniqueUsers.insert(userId);
            if (songMap[songId]->uniqueUsers.size() > 1) {
                songPlayCountSet.erase(make_pair(-(songMap[songId]->uniqueUsers.size()-1), songId));
            }
            songPlayCountSet.insert(make_pair(-(songMap[songId]->uniqueUsers.size()), songId));
        }

        userMap[userId]->updateUserSongList(songId);
    }

    void printAnalytics() { // o(n)
        for(auto it = songPlayCountSet.begin(); it != songPlayCountSet.end(); it++) {
            cout << songMap[it->second]->title << " " << abs(it->first)<< endl;
        }
        
        
    }

    vector<int> getLastThreeSongs(int userId) {
        list<int> songList = userMap[userId]->recentPlayList;
        
        vector<int> lastThreeSongs;
        for(auto it = songList.begin(); it != songList.end(); it++) {
            lastThreeSongs.push_back(*it);
            if(lastThreeSongs.size() == 3) {
                break;
            }
        }
        
        return lastThreeSongs;
    }

    vector<int> getLastNSongs(int userId, int N) {
        if(userMap.find(userId) == userMap.end()) {
            return vector<int>();
        }
        list<int> songList = userMap[userId]->recentPlayList;
        vector<int> lastNSongs;
        for(auto it = songList.begin(); it != songList.end(); it++) {
            lastNSongs.push_back(*it);
            if(lastNSongs.size() == N) {
                break;
            }
        }
        return lastNSongs;
    }

    void starSong(int userId, int songId) {
        songMap[songId]->isStarred = true;
    }

    void unstarSong(int userId, int songId){ 
        songMap[songId]->isStarred = false;
    }

    vector<int> getLastNStarredSongsPlayed(int userId, int N) {
        list<int> songList = userMap[userId]->recentPlayList;
        vector<int> lastNStarredSongs;
        for(auto it = songList.begin(); it != songList.end(); it++) {
            if(songMap[*it]->isStarred) {
                lastNStarredSongs.push_back(*it);
            }
            if(lastNStarredSongs.size() == N) {
                break;
            }
        }
        return lastNStarredSongs;
    }   
};
int MusicPlayerAnalytics::songSequenceId = 0;

int main() {
    MusicPlayerAnalytics obj;
    cout << obj.addSong("Song 1") << endl;
    cout << obj.addSong("Song 2") << endl;
    cout << obj.addSong("Song 3") << endl;
    cout << obj.addSong("Song 4") << endl;
    cout << obj.addSong("Song 5") << endl;
    cout << obj.addSong("Song 6") << endl;
    obj.addUser(1);
    obj.addUser(2);
    obj.playSong(1, 1);
    obj.playSong(2, 1);
    obj.playSong(2, 2);
    obj.playSong(3, 1);
    obj.playSong(4, 1);
    obj.playSong(5, 1);
    obj.playSong(6, 1);
    obj.playSong(4, 1);
    obj.printAnalytics();
    vector<int> lastThreeSongs = obj.getLastThreeSongs(1);
    for(auto x: lastThreeSongs) {
        cout << x << " ";
    }
    cout << endl;
    vector<int> lastNSongs = obj.getLastNSongs(1, 5);
    for(auto it = lastNSongs.begin(); it != lastNSongs.end(); it++) {
        cout << *it << " ";
    }
    cout << endl;
    obj.starSong(1, 1);
    vector<int> lastNStarredSongs = obj.getLastNStarredSongsPlayed(1, 2);
    for(auto x: lastNStarredSongs) {
        cout << x << " ";
    }
    cout << endl;
    return 0;
}