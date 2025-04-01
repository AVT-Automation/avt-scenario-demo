# avt-scenario-demo
 
This project is a Go-based representation of the backend service API described in the recruitment scenario.

---

## Features

- **Set Contribution**: Allows users to set a fixed contribution or calculate a percentage-based contribution based on salary.
- **Get Contribution**: Retrieves the current contribution value.
- **Unit Tests**: Includes a unit test for the fixed contribution scenario.

---

## Endpoints

1. **Set Contribution**
   - **URL**: `/contribution`
   - **Method**: `POST`
   - **Request Body**:
     {
       "contribution": 100.0, // Fixed contribution (optional)
       "percentage": 10.0,    // Percentage for salary-based contribution (optional)
       "salary": 120000.0     // Annual salary for percentage-based calculation (optional)
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
   git clone https://github.com/your-username/avt-scenario-demo.git
   cd avt-scenario-demo

2. Start Docker container:
   docker-compose up

4. Access the application at:
   http://localhost:8080