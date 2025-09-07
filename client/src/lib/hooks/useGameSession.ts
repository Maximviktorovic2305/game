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

// Хук для начала новой игры
export const useStartGame = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: startGame,
		onSuccess: (data) => {
			// Сохранение сессии в кэше
			queryClient.setQueryData(['gameSession', data.session.id], data.session)
		},
	})
}

// Хук для получения сессии игры
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

// Хук для получения следующего вопроса
export const useNextQuestion = (level: number) => {
	return useQuery({
		queryKey: ['question', level],
		queryFn: () => getNextQuestion(level),
		enabled: level > 0,
	})
}

// Хук для ответа на вопрос
export const useAnswerQuestion = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: answerQuestion,
		onSuccess: (data) => {
			// Обновление сессии в кэше
			queryClient.setQueryData(['gameSession', data.session.id], data.session)
		},
	})
}

// Хук для использования подсказки 50:50
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

// Хук для использования подсказки "Помощь зала"
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

// Хук для использования подсказки "Звонок другу"
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

// Хук для завершения игры
export const useQuitGame = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: quitGame,
		onSuccess: (data) => {
			// Обновление сессии в кэше
			queryClient.setQueryData(['gameSession', data.id], data)
		},
	})
}

// Хук для получения таблицы лидеров
export const useLeaderboard = () => {
	return useQuery({
		queryKey: ['leaderboard'],
		queryFn: getLeaderboard,
	})
}
