# Real Image Challenge 2016 (Distribution System)

This project is a distribution system designed to handle the creation of distributors, sub-distributors, checking permissions, and viewing distributor information. The system utilizes a CSV file containing data about cities, states, and countries for region validation.

## How to Run the code base
Ensure you have Go installed on your system. Navigate to the root directory of the project and run the following command:

```bash
go run main.go
```
Follow the on-screen prompts to interact with the program.

## How to Run the .exe file
To run the executable file (distibuter.exe), follow these steps:

1. Ensure that both the executable file and the cities.csv file are in the same folder.
2. Double-click on the executable file (distibuter.exe).
3. The program will start running, and you can interact with it through the command-line interface.
4. Follow the on-screen prompts to perform various actions like creating distributors, sub-distributors, checking permissions, and viewing distributor information.

## Usage

To run the program, execute the `main.go` file. Upon execution, the program will prompt the user with a menu to select various options:

1. Create a new distributor: Allows the user to create a new distributor.
2. Create a sub-distributor: Allows the user to create a sub-distributor under an existing distributor.
3. Check permission for a distributor: Allows the user to check permissions for a distributor.
4. View Distributors information: Displays information about existing distributors.
5. Exit the program: Exits the program.

### Sample Inputs/Outputs:
> **_NOTE:_**  List of choices is prompted to perform actions based on user input and continues to prompt the user until the program is exited.
- List of choices
```
Use the arrow keys to navigate: ↓ ↑ → ←
Select one of the below choices:
1. Create a new distributor
2. Create a sub distributor
3. Check permission for a distributor
4. Add distributors for the producer
5. View Distributors information
6. Create a new movie 
7. Exit the program

```
Enter distributor name: annapoorna
Enter the regions you want to include for this distributor: india
Enter the regions you want to exclude for this distributor: karnataka-india
```

- Choosing "Create a new distributor"

```
Enter distributor name: annapoorna
Enter the regions you want to include for this distributor: india
Enter the regions you want to exclude for this distributor: karnataka-india
```

- Choosing "Create a sub distributor"

```
Enter distributor name: manoranjan
Enter the regions you want to include for this distributor: chennai-tamil nadu-india
Enter the regions you want to exclude for this distributor: salem-tamil nadu-india
Enter the name of the parent distributor: annapoorna
```

- Choosing "Check permission for a distributor"

```
Enter distributor name that needs to be checked: manoranjan
Enter regions that need to be checked: india
Check Permission Result: [manoranjan has access to INDIA]
```


- Choosing "Add distributor for a producer"

```
Enter producer name: nolan
Enter distributors that need to be added: annapoorna, manoranjan
```

- Choosing "View Distributors information"

```
Distributor Information:
Name: annapoorna, Include: [INDIA], Exclude: [KARNATAKA-INDIA], Parent:
Name: manoranjan, Include: [INDIA], Exclude: [SALEM-TAMIL NADU-INDIA], Parent: annapoorna
```

- Choosing "View Distributors information", it will exit the program.

### Author
mukund (**https://github.com/FlagellumDiabolus**)

