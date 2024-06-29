# Recruitment System API

## Overview

The Resume Builder API is designed to handle various operations related to resume management, including uploading and parsing resumes for applicants. The API uses a third-party resume parser to extract relevant information from resumes and store it in the database.


### Project Structure

```
resume-builder/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── db/
│   └── postgres.go
├── handlers/
│   ├── job_handler.go
│   ├── resume_handler.go
│   ├── auth_handler.go
│   └── user_handler.go
├── middleware/
│   ├── auth_middleware.go
│   └── role_middleware.go
├── models/
│   ├── profile.go
│   ├── user.go
│   └── job.go
├── utils/
│   ├── jwt.go
│   └── resume_parser.go
├── .env
└── README.md

```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/prashantrewar/recruitment-system.git
cd recruitment-system
```

2. Install dependencies:

```bash
go mod tidy

```

3. Set up the .env file:

```bash
DSN=host=localhost user=yourusername password=yourpassword dbname=yourdatabasename port=5432 sslmode=disable JWT_SECRET_KEY=your_secret_key
RESUME_PARSER_API_KEY=your_resume_parser_api_key

```

4. Run the server:

```bash
go run cmd/main.go

```

#### Configuration

Make sure you have a PostgreSQL database running and configured correctly. The connection string should be placed in the .env file.



## Compiling and running the server

### For Backend

The curl commands to test the API manually. These commands assume that the backend server is running on http://localhost:8080.


- Sign Up

To create a new user profile:

```bash
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "userType": "Applicant",
    "profileHeadline": "Software Developer",
    "address": "123 Main St, City, Country"
}'


```

-  Log In

To authenticate a user and get a JWT token:

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
    "email": "john@example.com",
    "password": "password123"
}'

```

- Upload Resume

To upload a resume (replace YOUR_JWT_TOKEN with the token from the login response):

```bash
curl --location --request POST 'https://api.apilayer.com/resume_parser/upload' \
--header 'Content-Type: application/octet-stream' \
--header 'apikey: your_api_key' \
--data-binary '@/home/yourusername/Path to/Resume.pdf'


```

- Create Job (Admin only)

To create a job opening (replace YOUR_JWT_TOKEN with the token from an admin login response):

```bash
curl -X POST http://localhost:8080/admin/job -H "Authorization: Bearer YOUR_JWT_TOKEN" -H "Content-Type: application/json" -d '{
    "title": "Software Engineer",
    "description": "Job description here",
    "companyName": "Tech Company"
}'

```

- List Job Openings

To fetch job openings:

```bash
curl -X GET http://localhost:8080/jobs -H "Authorization: Bearer YOUR_JWT_TOKEN"

```

- Apply for Job

To apply for a job (replace YOUR_JWT_TOKEN with the token from the login response and JOB_ID with the job ID you want to apply for):

```bash
curl -X POST http://localhost:8080/jobs/apply?job_id=JOB_ID -H "Authorization: Bearer YOUR_JWT_TOKEN"

```

- Get Job Details (Admin only)

To fetch details about a job and the list of applicants (replace YOUR_JWT_TOKEN with the token from an admin login response and JOB_ID with the job ID):

```bash
curl -X GET http://localhost:8080/admin/job/JOB_ID -H "Authorization: Bearer YOUR_JWT_TOKEN"

```

- List All Applicants (Admin only)

To fetch a list of all users in the system (replace YOUR_JWT_TOKEN with the token from an admin login response):

```bash
curl -X GET http://localhost:8080/admin/applicants -H "Authorization: Bearer YOUR_JWT_TOKEN"

```

- Get Applicant Details (Admin only)

To fetch extracted data of an applicant (replace YOUR_JWT_TOKEN with the token from an admin login response and APPLICANT_ID with the applicant ID):

```bash
curl -X GET http://localhost:8080/admin/applicant/APPLICANT_ID -H "Authorization: Bearer YOUR_JWT_TOKEN"

```


These commands should help you verify that your Recruitment Management System backend is functioning correctly.