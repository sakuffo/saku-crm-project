This will be the backend for a CRM application, using a REST API to manage customers. It will be able to add, remove and query specific customers.

sakuffo


### Installation

  - Clone application to a system that has golang available
  - (Optional) run "go version" to ensure that go is available
  - enter "go run ." to run from sourcecode
  - (Optional) go build . to create a simple binary/executable

### Usage

  - Start the application and go to http://localhost:3300/
  - NB: App is on port 3300 because Airplay sometimes ran on port 3000
  - Usage instructions for the API (Basic API documentation)
  - Use curl OR Postman to interact with the API. Use http://localhost:3300/ as root address.
  - - All API paths located at "/" on the Application