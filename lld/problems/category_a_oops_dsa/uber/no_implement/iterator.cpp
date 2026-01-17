/*

https://leetcode.com/discuss/post/5838801/uber-sde-2-phone-screen-by-anonymous_use-99jw/

Question: Implement a Counter class that has the following methods:

1. `put(number)`: put the number to the data structure
2. `count(number)`: count the number of times `number` was put during the last `window=5 minutes`
3. `countAll()`: count the number of times any number was put during the last `window=5 minutes`.

Example:

At t = 10PM, `put(2)`

At t = 10:02PM, `put(2)`

At t = 10:03PM, `put(3)`

At t = 10:04PM, `count(2)` should return 2

At t = 10:04PM, `countAll()` should return 3

At t = 10:06PM, `count(2)` should return 1 (the one that was put at 10:02PM)

At t = 10:06PM, `countAll()` should return 2

Follow-ups:

1. If you were to write unit tests, what would they be
2. If your code was to be in production, what issues may it cause?
*/