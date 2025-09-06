import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import {
	startGame,
	getGameSession,
	getNextQuestion,
	answerQuestion,
	useFiftyFifty as apiUseFiftyFifty,
	useAudience as apiUseAudience,
	useCall as apiUseCall,
	quitGame,
	getLeaderboard,
} from '@/lib/api/gameApi'

export const useStartGame = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: startGame,
		onSuccess: (data) => {
			// Store session in cache
			queryClient.setQueryData(['gameSession', data.session.id], data.session)
		},
	})
}

export const useGameSession = (sessionId: number | null) => {
	return useQuery({
		queryKey: ['gameSession', sessionId],
		queryFn: async () => {
			if (!sessionId) return null
			return getGameSession(sessionId)
		},
		enabled: !!sessionId,
	})
}

export const useNextQuestion = (level: number) => {
	return useQuery({
		queryKey: ['question', level],
		queryFn: () => getNextQuestion(level),
		enabled: level > 0,
	})
}

export const useAnswerQuestion = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: answerQuestion,
		onSuccess: (data) => {
			// Update session in cache
			queryClient.setQueryData(['gameSession', data.session.id], data.session)
		},
	})
}

export const useFiftyFiftyLifeline = () => {
	return useMutation({
		mutationFn: ({
			sessionID,
			questionID,
		}: {
			sessionID: number
			questionID: number
		}) => apiUseFiftyFifty(sessionID, questionID),
	})
}

export const useAudienceLifeline = () => {
	return useMutation({
		mutationFn: ({
			sessionID,
			questionID,
		}: {
			sessionID: number
			questionID: number
		}) => apiUseAudience(sessionID, questionID),
	})
}

export const useCallLifeline = () => {
	return useMutation({
		mutationFn: ({
			sessionID,
			questionID,
		}: {
			sessionID: number
			questionID: number
		}) => apiUseCall(sessionID, questionID),
	})
}

export const useQuitGame = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: quitGame,
		onSuccess: (data) => {
			// Update session in cache
			queryClient.setQueryData(['gameSession', data.id], data)
		},
	})
}

export const useLeaderboard = () => {
	return useQuery({
		queryKey: ['leaderboard'],
		queryFn: getLeaderboard,
	})
}
