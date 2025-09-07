import { GameSession, Question, Option } from '@/types'
import PrizeChart from '@/components/PrizeChart'
import QuestionCard from '@/components/QuestionCard'

interface GameAreaProps {
	gameSession: GameSession
	currentQuestion: Question | null
	removedOptions: number[]
	onAnswer: (option: Option) => void
	onQuit: () => void
}

export default function GameArea({
	gameSession,
	currentQuestion,
	removedOptions,
	onAnswer,
	onQuit,
}: GameAreaProps) {
	return (
		<div className='h-screen flex flex-col'>
			{/* Основная игровая область с логотипом в центре */}
			<div className='flex-grow flex items-center justify-center relative'>
				{/* Логотип в центре */}
				<div className='absolute inset-0 flex items-center justify-center z-0'>
					<img
						src='/logo.webp'
						alt='Логотип Кто хочет стать миллионером?'
						className='w-120 h-120 object-contain'
					/>
				</div>

				{/* Таблица призов справа */}
				<div className='absolute right-0 top-0 h-full w-fit lg:p-4 z-10'>
					<PrizeChart currentLevel={gameSession.current_level} />
				</div>
			</div>

			{/* Область вопроса внизу */}
			<div className='bg-gray-800 rounded-t-xl p-4'>
				<div className='flex justify-between items-center mb-4'>
					<h2 className='text-xl font-bold text-blue-400'>
						Вопрос {gameSession.current_level}
					</h2>
					<div className='flex items-center space-x-4'>
						<div className='text-lg font-bold text-white'>
							Счет: ${gameSession.score.toLocaleString()}
						</div>
						<button
							onClick={onQuit}
							className='bg-red-600 hover:bg-red-700 text-white font-bold py-1 px-3 rounded-full text-sm'>
							Завершить игру
						</button>
					</div>
				</div>

				{currentQuestion ? (
					<QuestionCard
						question={currentQuestion}
						onAnswer={onAnswer}
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
							{/* Пустой div для поддержания согласованности макета */}
						</div>
					</div>
				)}
			</div>
		</div>
	)
}