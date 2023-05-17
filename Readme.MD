The  `Promise` Struct has three Functions:
- `onFulfilled` : This method is called when the promise is fulfilled.
- `onRejected` : This method is called when the promise is rejected.
- `onFinally` : This method is called when the promise is fulfilled or rejected.

The `NewPromise` Function creates a new Promise and returns it.

1. The `Then` method is used to specify the callback functions to be executed when the Promise is fulfilled or rejected. It takes two arguments: onFulfilled and onRejected. It sets the respective fields in the Promise struct and returns the same Promise instance, allowing for method chaining.
2. The `Catch` method is used to specify the callback function to be executed when the Promise is rejected. It takes one argument: onRejected. It sets the respective field in the Promise struct and returns the same Promise instance, allowing for method chaining.
3. The `Finally` method is used to specify the callback function to be executed when the Promise is fulfilled or rejected. It takes one argument: onFinally. It sets the respective field in the Promise struct and returns the same Promise instance, allowing for method chaining.
4. The `Resolve` method is used to resolve the Promise. It takes one argument: value. It sets the respective field in the Promise struct and returns the same Promise instance, allowing for method chaining.
5. The `Reject` method is used to reject the Promise. It takes one argument: value. It sets the respective field in the Promise struct and returns the same Promise instance, allowing for method chaining.


