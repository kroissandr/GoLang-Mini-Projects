ToDo-List SPEC (in my own words)

Firstly we need to import new package. And it is bufio. This is classical package input/output but with buferization. Then we need to create a slice to our tasks as the global variable.

Then func main. Where we declare our "scanner", object that helps us to read data from command line. We declare it like bufio.NewScanner(), inside brackets we should write os.Stdin (from Operation System package choose Standart Input).

Next step will be blank for loop where we put our text menu, callings of our object and switch-case with choosing menu options.

Let's start from menu. It is header (that points this is the menu), view tasks, add a task, delete a task, exit and choose an option string.

After we need call our object methods. Firstly Scan() that tries to read next token. Secondly Text() that tries to return read token. We need to declare new variable for using second method.

Then we need to create switch-case with our options.

Next will be three func that let us manipulate our app.

viewTasks provide us whether all the tasks or message that we don't have any tasks.

In addTask we use pointer to our scanner as the input parameter. Then we use our familiar callings Scan() and Text() and do an append to add our task to the tasks slice.

In deleteTask, we also use the same pattern as in addTask, except we need to convert our taskNum to an integer (use strconv.Atoi). Then we need to use append to the our slice again, but to delete our specified task. Here we have sophisticated way to do that. We need to use append and as the arguments first slice with values before our target task number and second slice with values after our target task number. Then we need to convert second slice to numbers because append func only works in the pattern <slice>, <element1>, <element2> etc. We do it by the help of the "..." operator, that unpackage our second slice to bunch of numbers.