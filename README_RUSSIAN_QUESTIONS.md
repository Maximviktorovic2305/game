# Russian Questions Seeding Script

This script allows you to manually seed the database with Russian questions without using Docker.

## Prerequisites

1. Make sure PostgreSQL is running on your system
2. Ensure the database settings in the script match your PostgreSQL configuration:
   - Host: localhost
   - Port: 5432
   - Database name: mill
   - Username: postgres
   - Password: admin

## How to Run

1. Make sure your PostgreSQL database is running
2. Navigate to the server directory:
   ```bash
   cd server
   ```
3. Run the seeding script:
   ```bash
   go run seed_russian_questions.go
   ```

## What the Script Does

- Connects to the PostgreSQL database
- Runs migrations to ensure tables exist
- Clears any existing questions and options
- Inserts 150 Russian questions (10 for each of the 15 levels)
- Each question has 4 options with one correct answer

## Database Structure

The script will populate two tables:
- `questions`: Contains the questions with their level and prize amount
- `options`: Contains the answer options for each question

## Troubleshooting

If you encounter connection issues:
1. Verify PostgreSQL is running
2. Check that the database credentials in the script match your setup
3. Ensure the "mill" database exists in PostgreSQL