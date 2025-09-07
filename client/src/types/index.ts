// Вариант ответа на вопрос
export interface Option {
	id: number
	question_id: number
	letter: string
	text: string
	is_correct: boolean
}

// Вопрос игры
export interface Question {
	id: number
	text: string
	level: number
	prize: number
	options: Option[]
}

// Подсказки
export interface Lifelines {
	fifty_fifty: boolean
	audience: boolean
	call: boolean
}

// Сессия игры
export interface GameSession {
	id: number
	current_level: number
	score: number
	lifelines: Lifelines
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	start_time: string
	end_time?: string
}

// Запись в таблице лидеров
export interface LeaderboardEntry {
	id: number
	score: number
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	start_time: string
}

// Ответ на начало игры
export interface StartGameResponse {
	session: GameSession
}

// Запрос на ответ
export interface AnswerRequest {
	session_id: number
	question_id: number
	option_id: number
}

// Ответ на запрос
export interface AnswerResponse {
	correct: boolean
	session: GameSession
}

// Ответ 50:50
export interface FiftyFiftyResponse {
	removed_options: Option[]
}

// Ответ аудитории
export interface AudienceResponse {
	percentages: Record<string, number>
}

// Ответ звонка другу
export interface CallResponse {
	advice: string
}
