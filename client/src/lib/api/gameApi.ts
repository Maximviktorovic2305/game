import apiClient from './client'
import {
	StartGameResponse,
	AnswerRequest,
	AnswerResponse,
	FiftyFiftyResponse,
	AudienceResponse,
	CallResponse,
	GameSession,
	LeaderboardEntry,
} from '@/types'

// Game API functions
export const startGame = async (): Promise<StartGameResponse> => {
	const response = await apiClient.post<StartGameResponse>('/game/start')
	return response.data
}

export const getGameSession = async (
	sessionId: number,
): Promise<GameSession> => {
	const response = await apiClient.get<GameSession>(
		`/game/session/${sessionId}`,
	)
	return response.data
}

export const getNextQuestion = async (level: number) => {
	const response = await apiClient.get(`/questions/next?level=${level}`)
	return response.data
}

export const answerQuestion = async (
	data: AnswerRequest,
): Promise<AnswerResponse> => {
	const response = await apiClient.post<AnswerResponse>('/game/answer', data)
	return response.data
}

export const useFiftyFifty = async (
	sessionID: number,
	questionID: number,
): Promise<FiftyFiftyResponse> => {
	const response = await apiClient.post<FiftyFiftyResponse>(
		'/game/lifeline/fifty-fifty',
		{
			session_id: sessionID,
			question_id: questionID,
		},
	)
	return response.data
}

export const useAudience = async (
	sessionID: number,
	questionID: number,
): Promise<AudienceResponse> => {
	const response = await apiClient.post<AudienceResponse>(
		'/game/lifeline/audience',
		{
			session_id: sessionID,
			question_id: questionID,
		},
	)
	return response.data
}

export const useCall = async (
	sessionID: number,
	questionID: number,
): Promise<CallResponse> => {
	const response = await apiClient.post<CallResponse>('/game/lifeline/call', {
		session_id: sessionID,
		question_id: questionID,
	})
	return response.data
}

export const quitGame = async (sessionID: number): Promise<GameSession> => {
	const response = await apiClient.post<GameSession>(`/game/quit/${sessionID}`)
	return response.data
}

export const getLeaderboard = async (): Promise<LeaderboardEntry[]> => {
	const response = await apiClient.get<LeaderboardEntry[]>('/leaderboard')
	return response.data
}
