# Who Wants to Be a Millionaire? (Кто хочет стать миллионером?)

A web-based implementation of the famous quiz game "Who Wants to Be a Millionaire?" built with Next.js, Go, and PostgreSQL. The entire application has been translated to Russian and populated with 150 Russian questions.

## Features

- 15 challenging questions to answer in Russian
- Three lifelines: 50/50, Ask the Audience, and Phone a Friend (all translated to Russian)
- Prize ladder with checkpoints at question 5 and 10
- Leaderboard to track top scores
- Responsive design that works on all devices
- Full backend integration with no mock data

## Technologies Used

### Frontend
- Next.js 15 with App Router
- TypeScript
- Tailwind CSS
- Shadcn UI Components
- TanStack Query (React Query)

### Backend
- Go 1.22+
- Echo Framework
- GORM ORM
- PostgreSQL Database

## Project Structure

```
/project-root
├── client/                  # Next.js frontend
├── server/                  # Go backend
├── docker-compose.yml       # PostgreSQL setup
└── README.md
```

## Setup Instructions

### Prerequisites
- Node.js (v18 or higher)
- Go (v1.22 or higher)
- Docker (for PostgreSQL)
- PostgreSQL (if not using Docker)

### 1. Start the Database

```bash
docker-compose up -d
```

### 2. Seed the Database with Russian Questions

The database has been populated with 150 Russian questions (10 for each level). You can use either Docker or manual seeding:

**Option 1: Using Docker (recommended)**
```bash
docker-compose up -d
```

**Option 2: Manual seeding**
```bash
cd server
go run seed_russian_questions.go
```

### 3. Start the Backend Server

```bash
cd server
go mod tidy
go run cmd/main.go
```

### 4. Start the Frontend Application

```bash
cd client
npm install
npm run dev
```

The application will be available at:
- Frontend: http://localhost:3000 (or another available port if 3000 is in use)
- Backend API: http://localhost:8080

## Game Rules

1. Answer 15 questions correctly to win $1,000,000
2. Use lifelines strategically:
   - **50/50**: Removes two incorrect answers
   - **Ask the Audience**: Shows poll results
   - **Phone a Friend**: Gets advice from a "friend"
3. Checkpoints at questions 5 ($1,000) and 10 ($32,000)
4. If you answer incorrectly after a checkpoint, you keep the checkpoint prize
5. You can quit at any time and keep your current winnings

## API Endpoints

- `POST /api/game/start` - Start a new game session
- `GET /api/questions/next?level=:level` - Get the next question
- `POST /api/game/answer` - Submit an answer
- `POST /api/game/lifeline/fifty-fifty` - Use 50/50 lifeline
- `POST /api/game/lifeline/audience` - Use Ask the Audience lifeline
- `POST /api/game/lifeline/call` - Use Phone a Friend lifeline
- `POST /api/game/quit/:sessionID` - Quit the game
- `GET /api/leaderboard` - Get top scores

## Development

### Backend Structure

```
/server
├── cmd/
│   └── main.go               # Entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Environment configuration
│   ├── database/
│   │   ├── db.go             # Database connection
│   │   ├── migrate.go        # Database migrations
│   │   └── seed.go           # Seed data
│   ├── models/
│   │   ├── question.go       # Question model
│   │   ├── option.go         # Option model
│   │   └── game_session.go   # Game session model
│   ├── repository/
│   │   ├── question_repository.go
│   │   └── game_session_repository.go
│   ├── services/
│   │   ├── question_service.go
│   │   └── game_service.go
│   ├── handlers/
│   │   ├── question_handler.go
│   │   └── game_handler.go
│   └── routes/
│       └── routes.go         # API routes
├── .env                      # Environment variables
├── go.mod
└── go.sum
```

### Frontend Structure

```
/client
├── app/
│   ├── page.tsx              # Home page
│   ├── game/
│   │   └── page.tsx          # Game screen
│   └── leaderboard/
│       └── page.tsx          # Leaderboard
├── components/
│   ├── QuestionCard.tsx
│   ├── LifelinePanel.tsx
│   ├── PrizeChart.tsx
│   ├── GameOverModal.tsx
│   └── LeaderboardTable.tsx
├── lib/
│   ├── api/
│   │   ├── client.ts         # Axios client
│   │   ├── gameApi.ts        # API functions
│   │   └── queryClient.ts    # TanStack Query client
│   └── hooks/
│       └── useGameSession.ts # Custom hooks
├── types/
│   └── index.ts              # TypeScript interfaces
└── .env.local                # Environment variables
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.