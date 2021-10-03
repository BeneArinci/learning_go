## Channels and Go routines

### Routines

When we launch a Go program we automatically create a Go routine. The main function is a built in Go routine.<br>

The main in routine takes all the lines in the program run/execute all of them one by one. This means that, if something gets stuck and takes some time to be resolved, all the rest of the program would wait for it to be resolved. <br>

Go routines would allow us to solve this problem. By creating a new routine every time we find a blocking function (a function that needs a longer amount of time to be resolved), we are basically asking to our program to initiate that routine but also to move to the next one while waiting for its result!<br>

To initiate a new routine you just add the go keyword just before the function name:

```
    go checkLink(link)
```

##### What happens when we run many routines?

Behind the scene there is something called **Go scheduler** that is using 1 CPU at a time. <br>
Even if more routines are created, only one of them will be executed at a time. <br>
The scheduler is checking what is happening inside the routines and orders them.

### Channels

The main function is what is establishing what it's needed to be done and what code needs to be executed. When we initialise multiple Go routines inside the main routine, we basically tell the main routine to run other routines, any time it finds a blocker function, and to move forward. <br>

These child routines are detached from the main one. This means that the main routine, after initialising them, automatically forgets that they are running. If the main execution ends before the child routines are executed, the program stops and the child routines will never be executed. <br>

Go **Channels** help us resolve this issue. A channel intermediates between different routines.

To create a channel we can

```
    c := make(chan string)
```

Then we can inject the channel inside the routine we are creating (as an argument) and send messages into that channel using the following sintax:

```
    c <- "Might be down I think"
```

After this, it is necessary to block the code and make it wait for the channel to receive a message. More specifically, we want it to receive a message from all the routines we have opened. In the exercise, the opened routines were as many as the links we wanted to reach. For this reason, we blocked the code with this for loop:

```
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
```
