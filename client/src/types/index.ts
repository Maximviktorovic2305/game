export interface Option {
	id: number
	question_id: number
	letter: string
	text: string
	is_correct: boolean
}

export interface Question {
	id: number
	text: string
	level: number
	prize: number
	options: Option[]
}

export interface Lifelines {
	fifty_fifty: boolean
	audience: boolean
	call: boolean
}

export interface GameSession {
	id: number
	current_level: number
	score: number
	lifelines: Lifelines
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	start_time: string
	end_time?: string
}

export interface LeaderboardEntry {
	id: number
	score: number
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	start_time: string
}

export interface StartGameResponse {
	session: GameSession
}

export interface AnswerRequest {
	session_id: number
	question_id: number
	option_id: number
}

export interface AnswerResponse {
	correct: boolean
	session: GameSession
}

export interface FiftyFiftyResponse {
	removed_options: Option[]
}

export interface AudienceResponse {
	percentages: Record<string, number>
}

export interface CallResponse {
	advice: string
}
