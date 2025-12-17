#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <unordered_map>
#include <vector>
#include <string>
#include <queue>
#include <unordered_set>
using namespace std;

/* ---------------- Mock HtmlParser ----------------
   This simulates web pages and links.
   In real LeetCode, HtmlParser is provided.
--------------------------------------------------*/
class HtmlParser {
public:
    unordered_map<string, vector<string>> graph;

    HtmlParser() {
        graph["http://site.com"] = {
            "http://site.com/a",
            "http://site.com/b",
            "http://othersite.com/x"
        };
        graph["http://site.com/a"] = {
            "http://site.com",
            "http://site.com/c"
        };
        graph["http://site.com/b"] = {
            "http://site.com/c"
        };
        graph["http://site.com/c"] = {};
    }

    vector<string> getUrls(const string& url) {
        // simulate slow IO
        this_thread::sleep_for(100ms);
        return graph[url];
    }
};

/* ---------------- Solution ---------------- */
class Solution {
public:
    mutex mx;
    condition_variable cv;
    vector<string> crawl(string startUrl, HtmlParser& htmlParser) {
        vector<string> result;

        queue<string> q;
        q.push(startUrl);
        unordered_set<string> seen;
        seen.insert(startUrl);
        int activeWorkers = 0;
        bool done = false;



        auto worker = [&]() {
            while (true) {
                string url;
                {
                    unique_lock<mutex> lock(mx);
                    cv.wait(lock, [&](){ return done || !q.empty(); });

                    if (done) {
                        return;
                    }

                     url = q.front();
                     cout << "Processing URL: " << url << endl;
                     q.pop();
                     activeWorkers++;
                }

                vector<string> urls = htmlParser.getUrls(url);

                for (auto &url : urls) {
                    

                    if (seen.find(url) == seen.end() && getHostname(url) == getHostname(startUrl)) {
                        lock_guard<mutex> lock(mx);
                        seen.insert(url);
                        q.push(url);
                        cout << "Added URL to queue: " << url << endl;
                        cv.notify_one();
                    }
                }


                cout << "Notifying all workers" << endl;    
                cv.notify_all();

                {
                    lock_guard<mutex> lock(mx);
                    activeWorkers--;
                    if (activeWorkers == 0 && q.empty()) {
                        done = true;
                        cv.notify_all();
                    }
                }
            }
        };

        vector<thread> threads;
        int cnt = max(1u, thread::hardware_concurrency()/2);

        for (int i = 0; i < cnt; i++) {
            threads.push_back(thread(worker));
        }

        for (auto &t : threads) {
            t.join();
        }

        for (auto url: seen) {
            result.push_back(url);
        }
        return result;
    }

private:
    string getHostname(const string& url) {
        return "http://site.com";
    }
};

/* ---------------- main() ---------------- */
int main() {
    HtmlParser parser;
    Solution sol;

    auto result = sol.crawl("http://site.com", parser);

    cout << "Crawled URLs:\n";
    for (const auto& url : result) {
        cout << url << "\n";
    }

    return 0;
}
