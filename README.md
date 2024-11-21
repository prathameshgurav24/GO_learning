 ![image](https://github.com/user-attachments/assets/b085262a-26d0-45ed-bc25-8b9d6665b00f)

Go Essentials
	Values, Basic Types & Core Language Features
	Understanding the Key Components of a Go Program
	Working with Values & Types
	Creating & Executing Functions
	Controlling Execution with Control Structures
Go standard Library : https://pkg.go.dev/std

 ![image](https://github.com/user-attachments/assets/014ebac7-9dba-45e2-a228-9d4452be12ca)

The entry point for the application is the main function in the main package.
A complete program is created by linking a single, unimported package called the main package with all the packages it imports, transitively. The main package must have package name main and declare a function main that takes no arguments and returns no value.
func main() { … }
Program execution begins by initializing the main package and then invoking the function main. When that function invocation returns, the program exits. It does not wait for other (non-main) goroutines to complete.
The language specification does not give special meaning to the name main outside of this context. The name main is not a reserved name.
It's OK to declare a main function in non-main packages. In such cases, it's just a function named main.

Go Commands
o	go run app.go
o	go mod init example/first-app
o	go build
variable declarartion:
 ![image](https://github.com/user-attachments/assets/62bccefc-c9ad-478d-80bb-33e3e08ff637)
 ![image](https://github.com/user-attachments/assets/90113192-14d4-491c-ba96-472fcc61279b)
![image](https://github.com/user-attachments/assets/56decd53-994e-4777-b123-68caac8e24b7)


 
 


The fmt.Scan() function is a great function for easily fetching & using user input through the command line.
But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.
For the moment, we only need single words / digits as input, so that's no problem.
Later in the course, when we work on projects where more complex input values are required, you'll therefore learn about an alternative to fmt.Scan()
 ![image](https://github.com/user-attachments/assets/7ec95e56-8c11-4048-8b12-480380805267)
![image](https://github.com/user-attachments/assets/8156ac8a-ccfe-4212-aa2f-1b2d33c7c4a9)

 
Sprint for saving print statements:
 ![image](https://github.com/user-attachments/assets/d57baccb-fe7e-43b0-8612-045d169ace21)
![image](https://github.com/user-attachments/assets/21e89624-f659-4e40-a178-f242f9133992)

 
Func returning one value:
 ![image](https://github.com/user-attachments/assets/103f90c0-a950-44e0-996e-d8c80d2b094c)

Func returning multiple values:
![image](https://github.com/user-attachments/assets/a3514bbe-afd4-4a59-9347-2644e19c188d)
![image](https://github.com/user-attachments/assets/244551a5-e8dd-433d-a5f3-482a9113bc29)

 
 
An Alternative Return Value Syntax  

Read and write to the file:
 ![image](https://github.com/user-attachments/assets/df10ad58-862e-4197-948a-105d67ace00f)

Loops:
Conditional For Loops
Besides the for variations introduced in the last lectures, there also is another common variation (which will also be shown again later in the course):
1.	for someCondition {
2.	  // do something ...
3.	}
someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).
In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
![image](https://github.com/user-attachments/assets/b07c30c0-39aa-436d-ac3b-b4013762763d)


 

Error Handling:
![image](https://github.com/user-attachments/assets/fe270080-5630-4369-9fcf-e2621a9e1fa3)
 ![image](https://github.com/user-attachments/assets/0989b7b1-cbbf-458c-9dd9-96fc71e53df0)

 

Packages - A Closer Look
	Working with Go Packages
	Splitting Code Across Multiple Files
	Splitting Files Across Multiple Packages
	Importing & Using Custom Packages

