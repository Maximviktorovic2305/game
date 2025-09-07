'use client'

import { useState, useEffect } from 'react'
import {
	useStartGame,
	useNextQuestion,
	useAnswerQuestion,
	useFiftyFiftyLifeline,
	useAudienceLifeline,
	useCallLifeline,
	useQuitGame,
} from '@/lib/hooks/useGameSession'
import { GameSession, Question, Option } from '@/types'
import GameOverModal from '@/components/GameOverModal'
import GameArea from '@/components/game/GameArea'
import LifelinesArea from '@/components/game/LifelinesArea'
import AudienceModal from '@/components/modals/AudienceModal'
import CallModal from '@/components/modals/CallModal'

export default function GamePage() {
	const [gameSession, setGameSession] = useState<GameSession | null>(null)
	const [currentQuestion, setCurrentQuestion] = useState<Question | null>(null)
	const [fiftyFiftyUsed, setFiftyFiftyUsed] = useState(false)
	const [audienceUsed, setAudienceUsed] = useState(false)
	const [callUsed, setCallUsed] = useState(false)
	const [removedOptions, setRemovedOptions] = useState<number[]>([])
	const [audiencePercentages, setAudiencePercentages] = useState<Record<
		string,
		number
	> | null>(null)
	const [callAdvice, setCallAdvice] = useState<string | null>(null)
	const [showGameOver, setShowGameOver] = useState(false)

	const { mutate: startGame } = useStartGame()
	const { data: questionData, refetch: fetchQuestion } = useNextQuestion(
		gameSession?.current_level || 0,
	)
	const { mutate: answerQuestion } = useAnswerQuestion()
	const { mutate: triggerFiftyFifty, data: fiftyFiftyData } =
		useFiftyFiftyLifeline()
	const { mutate: triggerAudience, data: audienceData } = useAudienceLifeline()
	const { mutate: triggerCall, data: callData } = useCallLifeline()
	const { mutate: quitGame } = useQuitGame()

	// Начало новой игры при монтировании компонента
	useEffect(() => {
		startGame(undefined, {
			onSuccess: (data) => {
				setGameSession(data.session)
			},
		})
	}, [startGame])

	// Получение следующего вопроса при изменении уровня
	useEffect(() => {
		if (gameSession?.current_level) {
			fetchQuestion()
		}
	}, [gameSession?.current_level, fetchQuestion])

	// Обновление текущего вопроса при получении данных
	useEffect(() => {
		if (questionData) {
			setCurrentQuestion(questionData)
		}
	}, [questionData])

	// Обработка ответа 50:50
	useEffect(() => {
		if (fiftyFiftyData) {
			const removedIds = fiftyFiftyData.removed_options.map(
				(option) => option.id,
			)
			setRemovedOptions(removedIds)
			setFiftyFiftyUsed(true)
		}
	}, [fiftyFiftyData])

	// Обработка ответа аудитории
	useEffect(() => {
		if (audienceData) {
			setAudiencePercentages(audienceData.percentages)
			setAudienceUsed(true)
		}
	}, [audienceData])

	// Обработка ответа звонка другу
	useEffect(() => {
		if (callData) {
			setCallAdvice(callData.advice)
			setCallUsed(true)
		}
	}, [callData])

	const handleAnswer = (option: Option) => {
		if (!gameSession || !currentQuestion) return

		answerQuestion(
			{
				session_id: gameSession.id,
				question_id: currentQuestion.id,
				option_id: option.id,
			},
			{
				onSuccess: (data) => {
					setGameSession(data.session)

					if (data.session.status !== 'in_progress') {
						// Игра окончена
						setShowGameOver(true)
					} else {
						// Переход к следующему вопросу
						setCurrentQuestion(null)
						setRemovedOptions([])
						setAudiencePercentages(null)
						setCallAdvice(null)
					}
				},
			},
		)
	}

	const handleFiftyFifty = () => {
		if (!gameSession || !currentQuestion || fiftyFiftyUsed) return

		triggerFiftyFifty({
			sessionID: gameSession.id,
			questionID: currentQuestion.id,
		})
	}

	const handleAudience = () => {
		if (!gameSession || !currentQuestion || audienceUsed) return

		triggerAudience({
			sessionID: gameSession.id,
			questionID: currentQuestion.id,
		})
	}

	const handleCall = () => {
		if (!gameSession || !currentQuestion || callUsed) return

		triggerCall({ sessionID: gameSession.id, questionID: currentQuestion.id })
	}

	const handleCloseModals = () => {
		setAudiencePercentages(null)
		setCallAdvice(null)
	}

	const handleQuit = () => {
		if (!gameSession) return

		quitGame(gameSession.id, {
			onSuccess: (data) => {
				setGameSession(data)
				setShowGameOver(true)
			},
		})
	}

	if (!gameSession) {
		return (
			<div className='min-h-screen bg-gray-900 flex items-center justify-center'>
				<p className='text-white text-xl'>Начало игры...</p>
			</div>
		)
	}

	return (
		<div className='min-h-screen bg-gray-900 p-0'>
			<LifelinesArea
				fiftyFiftyUsed={fiftyFiftyUsed}
				audienceUsed={audienceUsed}
				callUsed={callUsed}
				onFiftyFifty={handleFiftyFifty}
				onAudience={handleAudience}
				onCall={handleCall}
			/>
			<GameArea
				gameSession={gameSession}
				currentQuestion={currentQuestion}
				removedOptions={removedOptions}
				onAnswer={handleAnswer}
				onQuit={handleQuit}
			/>

			{/* Модальные окна */}
			{audiencePercentages && (
				<AudienceModal
					percentages={audiencePercentages}
					onClose={handleCloseModals}
				/>
			)}

			{callAdvice && (
				<CallModal advice={callAdvice} onClose={handleCloseModals} />
			)}

			{showGameOver && gameSession && (
				<GameOverModal
					score={gameSession.score}
					status={gameSession.status}
					onClose={() => setShowGameOver(false)}
				/>
			)}
		</div>
	)
}
