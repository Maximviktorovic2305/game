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

	// SVG Icons for Lifelines
	const FiftyFiftyIcon = () => (
		<svg
			xmlns='http://www.w3.org/2000/svg'
			viewBox='0 0 40 40'
			className='w-10 h-10'>
			<circle
				cx='20'
				cy='20'
				r='18'
				fill={fiftyFiftyUsed ? '#4b5563' : '#3b82f6'}
			/>
			<text
				x='20'
				y='26'
				textAnchor='middle'
				fill='white'
				fontSize='16'
				fontWeight='bold'
				fontFamily='Arial, sans-serif'>
				50%
			</text>
		</svg>
	)

	const AudienceIcon = () => (
		<svg
			xmlns='http://www.w3.org/2000/svg'
			viewBox='0 0 40 40'
			className='w-10 h-10'>
			<circle
				cx='20'
				cy='20'
				r='18'
				fill={audienceUsed ? '#4b5563' : '#3b82f6'}
			/>
			<g transform='translate(8, 10)'>
				{/* First person */}
				<circle cx='5' cy='5' r='3' fill='white' />
				<path d='M5 9c-2 0-4 1-4 3v2h8v-2c0-2-2-3-4-3z' fill='white' />

				{/* Second person */}
				<circle cx='15' cy='5' r='3' fill='white' />
				<path d='M15 9c-2 0-4 1-4 3v2h8v-2c0-2-2-3-4-3z' fill='white' />

				{/* Third person */}
				<circle cx='25' cy='5' r='3' fill='white' />
				<path d='M25 9c-2 0-4 1-4 3v2h8v-2c0-2-2-3-4-3z' fill='white' />
			</g>
		</svg>
	)

	const CallIcon = () => (
		<svg
			xmlns='http://www.w3.org/2000/svg'
			viewBox='0 0 24 24'
			fill='currentColor'
			className='w-8 h-8'>
			<path d='M20 15.5c-1.25 0-2.45-.2-3.57-.57-.35-.11-.74-.03-1.02.24l-2.2 2.2c-2.83-1.44-5.15-3.75-6.59-6.59l2.2-2.21c.28-.26.36-.65.25-1C8.7 6.45 8.5 5.25 8.5 4c0-.55-.45-1-1-1H4c-.55 0-1 .45-1 1 0 9.39 7.61 17 17 17 .55 0 1-.45 1-1v-3.5c0-.55-.45-1-1-1z' />
		</svg>
	)

	if (!gameSession) {
		return (
			<div className='min-h-screen bg-gray-900 flex items-center justify-center'>
				<p className='text-white text-xl'>Starting game...</p>
			</div>
		)
	}

	return (
		<div className='min-h-screen bg-gray-900 p-0'>
			<div className='h-screen flex flex-col'>
				{/* Main game area with logo in center */}
				<div className='flex-grow flex items-center justify-center relative'>
					{/* Lifelines at top left */}
					<div className='absolute top-4 left-4 z-20 flex space-x-4'>
						<button
							onClick={handleFiftyFifty}
							disabled={fiftyFiftyUsed}
							className={`
								p-2 rounded-full transition-all duration-200 flex items-center justify-center
								${
									fiftyFiftyUsed
										? 'bg-gray-600 text-gray-400 cursor-not-allowed'
										: 'bg-blue-600 text-white hover:bg-blue-500'
								}
							`}
							title='50:50 Lifeline'>
							<FiftyFiftyIcon />
						</button>

						<button
							onClick={handleAudience}
							disabled={audienceUsed}
							className={`
								p-2 rounded-full transition-all duration-200 flex items-center justify-center
								${
									audienceUsed
										? 'bg-gray-600 text-gray-400 cursor-not-allowed'
										: 'bg-blue-600 text-white hover:bg-blue-500'
								}
							`}
							title='Ask the Audience'>
							<AudienceIcon />
						</button>

						<button
							onClick={handleCall}
							disabled={callUsed}
							className={`
								p-2 rounded-full transition-all duration-200
								${
									callUsed
										? 'bg-gray-600 text-gray-400 cursor-not-allowed'
										: 'bg-blue-600 text-white hover:bg-blue-500'
								}
							`}
							title='Phone a Friend'>
							<CallIcon />
						</button>
					</div>

					{/* Logo in the center */}
					<div className='absolute inset-0 flex items-center justify-center z-0'>
						<img
							src='/logo.webp'
							alt='Who Wants to Be a Millionaire Logo'
							className='w-120 h-120 object-contain'
						/>
					</div>

					{/* Prize Chart on the right */}
					<div className='absolute right-0 top-0 h-full w-fit lg:p-4 z-10'>
						<PrizeChart currentLevel={gameSession.current_level} />
					</div>
				</div>

				{/* Question area at the bottom */}
				<div
					className='bg-gray-800 rounded-t-xl p-4'
					>
					<div className='flex justify-between items-center mb-4'>
						<h2 className='text-xl font-bold text-blue-400'>
							Question {gameSession.current_level}
						</h2>
						<div className='text-lg font-bold text-white'>
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
						<div
							className='flex flex-col transition-opacity duration-300 ease-in-out opacity-0'
							style={{
								height: '200px',
								display: 'flex',
								flexDirection: 'column',
							}}>
							<div className='flex-grow'></div>
							<div className='space-y-4'>
								{/* Empty div to maintain layout consistency */}
							</div>
						</div>
					)}
				</div>
			</div>

			{/* Modals */}
			{audiencePercentages && (
				<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
					<div className='bg-gray-800 rounded-xl p-6 max-w-md w-full mx-4 border-2 border-blue-500'>
						<h3 className='text-2xl font-bold mb-4 text-center text-blue-400'>
							Audience Poll
						</h3>
						<div className='space-y-4'>
							{Object.entries(audiencePercentages).map(
								([letter, percentage]) => (
									<div key={letter}>
										<div className='flex justify-between mb-1'>
											<span className='font-bold text-white'>{letter}</span>
											<span className='text-white'>{percentage}%</span>
										</div>
										<div className='w-full bg-gray-700 rounded-full h-4'>
											<div
												className='bg-blue-500 h-4 rounded-full'
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
					<div className='bg-gray-800 rounded-xl p-6 max-w-md w-full mx-4 border-2 border-blue-500'>
						<h3 className='text-2xl font-bold mb-4 text-center text-blue-400'>
							Friend&apos;s Advice
						</h3>
						<p className='text-lg mb-6 text-center text-white'>{callAdvice}</p>
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
