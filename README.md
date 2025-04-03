# avt-scenario-demo
 
This project is a Go-based representation of the backend service API described in the recruitment scenario.

---

## Features

- **Set Contribution**: Allows users to set a fixed contribution or calculate a percentage-based contribution based on salary.
- **Get Contribution**: Retrieves the current contribution value.
- **Unit Tests**: Includes a unit test for the fixed contribution scenario.
- **Postman Tests**: A set of functional API tests running within a Newman container.

---

## Endpoints

1. **Set Contribution**
   - **URL**: `/contribution`
   - **Method**: `POST`
   - **Request Body**:
     {
       "contribution": 100, // Fixed contribution (optional)
       "percentage": 10,    // Percentage for salary-based contribution (optional)
       "salary": 120000   // Annual salary for percentage-based calculation (optional)
     }
   - **Response**:
     - 200 OK: Returns the current contribution value.
     - 404 Not Found: Returns an error if no contribution is set.

2. **Get Contribution**
   - **URL**: `/contribution`
   - **Method**: `GET`
   - **Response**:
     - 200 OK: Returns the current contribution value.
     - 404 Not Found: Returns an error if no contribution is set.

---

## Running Locally

1. Clone the repository:
   - git clone https://github.com/AVT-Automation/avt-scenario-demo
   - cd avt-scenario-demo

2. Start Docker containers

   **To start containers without running tests**
   - docker-compose up

   **To run tests and shut down containers**
   - docker-compose up --build --abort-on-container-exit

3. Access the application at:
   - http://localhost:8080