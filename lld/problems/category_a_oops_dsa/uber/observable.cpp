/* https://leetcode.com/discuss/post/6474653/problem-simple-observable-implementation-eycc/

Create a simple Observable class that implements the observer pattern. The class should:

- Allow subscribing to data changes via a subscribe method
- Notify all subscribers when data changes via a notify method
- Allow unsubscribing from updates
- Maintain a list of subscriber callbacks

```jsx
const observable=newObservable();

// Subscribe to changes
const subscription= observable.subscribe(data=>{
console.log('Received:', data);
});

// Notify subscribers
observable.notify('Hello!');// logs: "Received: Hello!"

// Unsubscribe
subscription.unsubscribe();

// No longer logs anything
observable.notify('Hello again!');
```

Found similar question on PrepareFrontend

*/