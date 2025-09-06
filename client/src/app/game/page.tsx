'use client'

import { useState, useEffect } from 'react'
import {
	useStartGame,
	useNextQuestion,
	useAnswerQuestion,
	useFiftyFiftyLifeline,
	useAudienceLifeline,
	useCallLifeline,
} from '@/lib/hooks/useGameSession'
import { GameSession, Question, Option } from '@/types'
import QuestionCard from '@/components/QuestionCard'
import LifelinePanel from '@/components/LifelinePanel'
import PrizeChart from '@/components/PrizeChart'
import GameOverModal from '@/components/GameOverModal'

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

	// Start a new game when the component mounts
	useEffect(() => {
		startGame(undefined, {
			onSuccess: (data) => {
				setGameSession(data.session)
			},
		})
	}, [startGame])

	// Fetch the next question when the level changes
	useEffect(() => {
		if (gameSession?.current_level) {
			fetchQuestion()
		}
	}, [gameSession?.current_level, fetchQuestion])

	// Update the current question when data is fetched
	useEffect(() => {
		if (questionData) {
			setCurrentQuestion(questionData)
		}
	}, [questionData])

	// Handle fifty-fifty response
	useEffect(() => {
		if (fiftyFiftyData) {
			const removedIds = fiftyFiftyData.removed_options.map(
				(option) => option.id,
			)
			setRemovedOptions(removedIds)
			setFiftyFiftyUsed(true)
		}
	}, [fiftyFiftyData])

	// Handle audience response
	useEffect(() => {
		if (audienceData) {
			setAudiencePercentages(audienceData.percentages)
			setAudienceUsed(true)
		}
	}, [audienceData])

	// Handle call response
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
						// Game is over
						setShowGameOver(true)
					} else {
						// Move to the next question
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
		// In a real implementation, you would call the quit API
		setShowGameOver(true)
	}

	if (!gameSession) {
		return (
			<div className='min-h-screen bg-gradient-to-b from-blue-900 to-purple-900 flex items-center justify-center'>
				<p className='text-white text-xl'>Starting game...</p>
			</div>
		)
	}

	return (
		<div className='min-h-screen bg-gradient-to-b from-blue-900 to-purple-900 p-4'>
			<div className='max-w-6xl mx-auto'>
				<div className='flex flex-col lg:flex-row gap-6'>
					{/* Prize Chart */}
					<div className='w-full lg:w-1/4'>
						<PrizeChart currentLevel={gameSession.current_level} />
					</div>

					{/* Main Game Area */}
					<div className='w-full lg:w-3/4'>
						<div className='bg-black/30 backdrop-blur-sm rounded-xl p-6 mb-6'>
							<div className='flex justify-between items-center mb-6'>
								<h2 className='text-2xl font-bold text-yellow-400'>
									Question {gameSession.current_level}
								</h2>
								<div className='text-xl font-bold text-white'>
									Score: ${gameSession.score.toLocaleString()}
								</div>
							</div>

							{currentQuestion ? (
								<QuestionCard
									question={currentQuestion}
									onAnswer={handleAnswer}
									removedOptions={removedOptions}
								/>
							) : (
								<div className='text-center py-12'>
									<p className='text-white text-xl'>Loading next question...</p>
								</div>
							)}
						</div>

						{/* Lifeline Panel */}
						<LifelinePanel
							onFiftyFifty={handleFiftyFifty}
							onAudience={handleAudience}
							onCall={handleCall}
							fiftyFiftyUsed={fiftyFiftyUsed}
							audienceUsed={audienceUsed}
							callUsed={callUsed}
						/>

						{/* Quit Button */}
						<div className='mt-6 text-center'>
							<button
								onClick={handleQuit}
								className='bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-6 rounded-full transition-colors'>
								Quit Game
							</button>
						</div>
					</div>
				</div>
			</div>

			{/* Modals */}
			{audiencePercentages && (
				<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
					<div className='bg-white rounded-xl p-6 max-w-md w-full mx-4'>
						<h3 className='text-2xl font-bold mb-4 text-center'>
							Audience Poll
						</h3>
						<div className='space-y-4'>
							{Object.entries(audiencePercentages).map(
								([letter, percentage]) => (
									<div key={letter}>
										<div className='flex justify-between mb-1'>
											<span className='font-bold'>{letter}</span>
											<span>{percentage}%</span>
										</div>
										<div className='w-full bg-gray-200 rounded-full h-4'>
											<div
												className='bg-blue-600 h-4 rounded-full'
												style={{ width: `${percentage}%` }}></div>
										</div>
									</div>
								),
							)}
						</div>
						<div className='mt-6 text-center'>
							<button
								onClick={handleCloseModals}
								className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-full'>
								Close
							</button>
						</div>
					</div>
				</div>
			)}

			{callAdvice && (
				<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
					<div className='bg-white rounded-xl p-6 max-w-md w-full mx-4'>
						<h3 className='text-2xl font-bold mb-4 text-center'>
							Friend&apos;s Advice
						</h3>
						<p className='text-lg mb-6 text-center'>{callAdvice}</p>
						<div className='text-center'>
							<button
								onClick={handleCloseModals}
								className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-full'>
								Close
							</button>
						</div>
					</div>
				</div>
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
