# Carrier - Technical Test - Go

In this repo, you'll find:
- The [PDF including the questions](https://github.com/thdelmas/carrier-technical-test/blob/main/20232012_Technical%20Assessment%20Test%20-%20Golang.pdf) to solve
- The answer for each question
- The code examples
- The code analysis

## Part 1 - Multiple choice questions
### 1. What is Golang's primary focus in terms of development?
Go was creaed by Google, with primary focus on APIs and Team collaboration through code maintability. It's also well know to be close to the C language even if it doesn't compete with Rust on system level

None of the possible answer seem correct:

```
a. Web development
-> NO: Go isn't javascript so no front end
-> YES: Go is used to develop API and basic servers so yes for the backend

b. System programming
-> NO: Go isn't Rust neither C and would be a terrible and heavy choice for a system app
-> YES: however it has capabilities to do it

c. Mobile application development
-> NOOOO: Definitely no 
-> COMMENT: From Swift for IOS, to Kotlin / Java for android going through all cross platform framework like flutter and react native, go would be one of the last choice for such use case

d. Data analysis
-> NO: The data scheme in Go aren't so dynamic and you will loose more time to setup your models than by learning pyhton
-> COMMENT: Choose pyhton3 or R in such cases
```
### 2. Which of the following is a correct way to declare a slice in Golang?
```
a. `var s []int `
b. `s := make([]int, 0) `
c. `s := []int{} `
d. All of the above

-> Answer D
```

### 3. In Golang, what does the `defer` keyword do?
```
a. Defer execution until the end of the program
b. Defer execution until the end of the enclosing function
c. Execute immediately
d. None of the above

-> Answer B, defer is very useful to postpone operations while keeping the line in a logical place
```

### 4. What is the purpose of the `init` function in Golang?
```
a. It is called before the main function
b. It is called after the main function
c. It is called when a package is imported
d. It is used for garbage collection

-> Answer C
```

### 5. Which of the following is true about goroutines in Golang?
```
a. They are lightweight threads managed by the Go runtime
b. They are only suitable for CPU-bound tasks
c. They are not supported in Golang
d. They require explicit memory management

-> Answer A
```

### 6. What is the correct way to create a new instance of a struct in Golang?
```
a. `new(MyStruct) `
b. `MyStruct{} `
c. `&MyStruct{} `
d. Both b and c

-> Answer B
```

### 7. In Golang, what does the `make` function do when used with a channel?
```
a. Creates a new channel
b. Allocates and initializes a channel
c. Closes the channel
d. Sends a value on the channel

-> Answer B
```

### 8. How do you handle errors in Golang?
```
a. Use panic and recover
b. Check for errors and return them
c. Ignore errors for simplicity
d. Errors are automatically handled by the runtime

-> Answer B
```

### 9. What is the purpose of the `context` package in Golang?
```
a. It provides a way to cancel operations
b. It is used for defining constants
c. It handles HTTP requests
d. It manages database connections

-> Answer C
```

### 10. How do you declare and initialize a map in Golang?
```
a. `var m map[string]int `
b. `m := map[string]int{} `
c. `m := make(map[string]int) `
d. All of the above

-> Answer C
```

## Part 2 - Code examples & demo
### 11. Database Interaction
### 12. Implement a Golang REST API endpoint for creating a new user. The user data should be
sent as JSON in the request body, and the API should return the created user's details.
### Write a Golang function that checks if a given string is a palindrome.
Explain your approach.
### 14. Implement a Golang program that concurrently fetches data from multiple URLs and
aggregates the results. Explain how you handle concurrent operations.
## Part 3 - Code Analysis
### 15. Analyze the following code. What will be the output, and can you identify any potential
issues or improvements?
### 16. Consider the following code. Explain the purpose of the done channel and any potential
issues you may foresee.
### 17. In Golang, how would you prevent data races in concurrent programs? Provide an example
or explain your approach.
### 18. Describe the role of the `http.HandleFunc` function in Golang and how it contributes to
building a web server.
## Part 4 - Feedback
