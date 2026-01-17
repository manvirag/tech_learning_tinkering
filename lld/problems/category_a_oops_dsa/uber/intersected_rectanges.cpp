/*

https://leetcode.com/discuss/post/2126058/uber-sde-2-l4-bangalore-india-offer-by-a-9ctr/

Given a 2-D plane, there are many rectangles(whose edges are parallel to X and Y axis) 
and there are many infinite lines(vertical and horizontal). 
The ask was to find total number of intersection points of the lines with the rectangles.

I created classes for all the models(line,rectangle,separate class for 
binary search methods) and proposed a binary search solution.
 Here, the focus is on writing modular code with proper naming
  conventions and handling all the edge cases. 
  Also, the expectation was to decide how to take the 
  input and write a working code passing against few test cases.

Also, i clarified with the interviewer about the edge cases when a line passes through one
 of the edges of the rectangle and if we count that as 2 points or 1.
This round also went really well.
*/

// this is more like coding quesion.




#include <iostream>
using namespace std;

// Count elements in sorted vector v within range [L, R]
long long countInRange(const vector<int>& v, int L, int R) {
    if (L > R) return 0;
    auto it1 = lower_bound(v.begin(), v.end(), L);
    auto it2 = upper_bound(v.begin(), v.end(), R);
    return it2 - it1;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int H, V;
    cin >> H >> V;

    vector<int> Y(H), X(V);
    for (int i = 0; i < H; i++) cin >> Y[i];
    for (int i = 0; i < V; i++) cin >> X[i];

    sort(Y.begin(), Y.end());
    sort(X.begin(), X.end());

    int R;
    cin >> R;

    long long totalIntersections = 0;

    for (int i = 0; i < R; i++) {
        int x1, y1, x2, y2;
        cin >> x1 >> y1 >> x2 >> y2;

        // normalize coordinates
        if (x1 > x2) swap(x1, x2);
        if (y1 > y2) swap(y1, y2);

        long long verticalHits   = countInRange(X, x1, x2);
        long long horizontalHits = countInRange(Y, y1, y2);

        totalIntersections += 2 * verticalHits;
        totalIntersections += 2 * horizontalHits;
    }

    cout << totalIntersections << "\n";
    return 0;
}





