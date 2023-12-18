This project is a demonstration of CRUD (Create, Read, Update, Delete) operations implemented in Go using the GoFr framework. It provides a simple API for managing employee records.


<h2>Getting Started:- </h2>
Follow these steps to get the project up and running:

<h3>1. Installation</h3>
Clone the repository:

<ul>
  
git clone https://github.com/ritikk112/GoAssignment.git
  
cd GoAssignment

</ul>


Install project dependencies:
<ul>
go mod tidy
  
</ul>

<h3>2. Configuration </h3>
   
The project may require configuration depending on your environment, such as database settings or port numbers. You can find the configuration options in the config directory.

<h3>3. Running the Application </h3>
 
To start the application, run the following command from the project's root directory:
<ul>
go run main.go
  
</ul>

The application will start, and you will see output indicating the server is running on a specific port.

<h4>API Endpoints</h4>
The following API endpoints are available:-
<li>
POST /api/employees: Create a new employee record.
</li>
<li>
GET /api/employees/{id} Retrieve an employee record by ID.
</li>
<li>
PUT /api/employees/{id} Update an existing employee record by ID.
  
</li>
<li>
DELETE /api/employees/{id} Delete an employee record by ID.
  
</li>
<br>

You can use Postman to test various API Calls 

<h3>4.Testing </h3>

To run unit tests for the project, use the following command:
<ul>
  
cd handler

go test -cover
</ul>

