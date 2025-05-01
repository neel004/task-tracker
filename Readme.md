Requirements
------------

The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:

*   Add, Update, and Delete tasks
*   Mark a task as in progress or done
*   List all tasks
*   List all tasks that are done
*   List all tasks that are not done
*   List all tasks that are in progress

Here are some constraints to guide the implementation:

*   You can use any programming language to build this project.
*   Use positional arguments in command line to accept user inputs.
*   Use a JSON file to store the tasks in the current directory.
*   The JSON file should be created if it does not exist.
*   Use the native file system module of your programming language to interact with the JSON file.
*   Do not use any external libraries or frameworks to build this project.
*   Ensure to handle errors and edge cases gracefully.

### Example Usage

The list of commands and their usage is given below:

    # Adding a new task
    ./task-tracker add "Buy groceries"
    # Output: Task added successfully (ID: 1)
    
    # Updating and deleting tasks
    ./task-tracker update 1 "Buy groceries and cook dinner"
    ./task-tracker delete 1
    
    # Moving a task to inprogress or done or todo
    ./task-tracker move 1 inprogress
    ./task-tracker move 1 done
    
    # Listing all tasks
    ./task-tracker list
    
    # Listing tasks by status
    ./task-tracker list done
    ./task-tracker list todo
    ./task-tracker list inprogress


#### This project is inspired from https://roadmap.sh/projects/task-tracker with slight modifications on original requirements.
