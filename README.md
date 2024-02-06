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
-> Noo please: Definitely no 
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
```
1. Clone the repo
2. cd carrier-technical-test/part2
```
### 11. Database Interaction
```bash
cd 11_database_interaction && ./install.sh
```

### 12. Implement a Golang REST API endpoint for creating a new user. The user data should be sent as JSON in the request body, and the API should return the created user's details.
```bash
cd 12_REST_user_creation && ./install.sh
```

### 13. Write a Golang function that checks if a given string is a palindrome.
```bash
cd 13_PALINDROM && ./install.sh
```

### 14. Implement a Golang program that concurrently fetches data from multiple URLs and aggregates the results. Explain how you handle concurrent operations.
```bash
cd 14_data_fetching && ./install.sh
```

## Part 3 - Code Analysis
### 15. Analyze the following code. What will be the output, and can you identify any potential issues or improvements?
```go
package main
import "fmt"
func main() {
s1 := []int{1, 2, 3}
s2 := s1[:2]
s2[0] = 4
fmt.Println(s1[0])
}
```
> Here the varaiable `s2` is a pointer array that copy the two first address of `s1`
>
> this is not a recommended behavior because it's not explicit that by modifying one you modify the other one
>
> in this case `s2` follows the order fo `s1` b ut we could imagine cases where `s2[0]` is a pointer to `s1[1]` instead of `s1[0]`, in such case it becomes really easy to lost track and get confused
>
> In this example the output is 4, `s2[0]` and `s1[0]` being the same place in memory

### 16. Consider the following code. Explain the purpose of the done channel and any potential issues you may foresee.
```go
func main() {
ch := make(chan int)
done := make(chan bool)
go func() {
// Some time-consuming task
ch <- 1
done <- true
}()
// Your explanation and code here
<-done
close(ch)
}
```
> In this case we have one channel with time consuming operation and another one with a synchronous and short operations
> 
> We place the synchronous and short after the other one, so it will finish after the previous task is done
>
> it permits avoiding to read something uncomplete in the important channel
>
> To do so we read the non important and atomical operation channel, when that one finish it means the other finished as well


### 17. In Golang, how would you prevent data races in concurrent programs? Provide an example or explain your approach.
> In golang it's possible to prevent data race by using channels or locks
>
> Channel will suspend the workflow until the writing on the asked resources is done, it allow us to avoid multiple editions at the same time

### 18. Describe the role of the `http.HandleFunc` function in Golang and how it contributes to building a web server.
> `http.HandleFunc` allow us to call functions on specifc endpoints calls, it's useful to build a web server as long as you need to handle request on specific endpoints

### 19. Review the code snippet below. Explain the purpose of the `defer` statement in the `someFunction` function.
```go
func someFunction() {
defer fmt.Println("World")
fmt.Print("Hello, ")
}
```
> `defer` will allow us to execute the line at the end of the function , so it will print `Hello` and then print `World`

### 20. In Golang, what is the purpose of the `json` package, and how would you use it to handle JSON data?
> `json` package is used to decode json data format and extract it in our data model, cf. part2

## Part 4 - Beyond

### 21. Explain the concept of interfaces in Golang and provide an example of how you would use them in a program.
> An interface is like a wrapper of function, you use it like class methods in Object oriented langage to perform standard actions on yopur predefined datatypes

### 22. Discuss the significance of the `context` package in the context of handling timeouts and cancellations in Golang applications.
> `context` package contain all information you need about the request including time and status, so by passing it along you can refer to those value to check if your contyext is still valid depending on your use case, you could use it for too old request to be canceled

### 23. Elaborate on the principles of error handling in Golang. How do you design a robust error-handling strategy in a large-scale application?

> Error handling is necessary in any function as it permits to check that the current behavior is the expected one, and to stop if it's not the case
>
> All function should return an error variable that contain nil if no error show up
>
> All big system should have an external log & error monitor to save this outside of the runtime and come back to it later or with a specific team (sentry, syslog, etc...)

### 24. Describe the characteristics and use cases of the `sync.WaitGroup` type in Golang. Provide an example of its usage in concurrent programming.

> No idea

### 25. Discuss the benefits and potential challenges of using goroutines in Golang for concurrent programming. How would you manage synchronization between goroutines?

> go routines allow us to perform background operations, however it's challenging to ensure data safety access to avoid writing / reading conflict between process
