import Link from 'next/link'

interface GameOverModalProps {
	score: number
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	onClose: () => void
}

export default function GameOverModal({
	score,
	status,
	onClose,
}: GameOverModalProps) {
	const getGameOverMessage = () => {
		switch (status) {
			case 'won':
				return {
					title: 'Поздравляем!',
					message: `Вы выиграли $${score.toLocaleString()}!`,
					color: 'text-green-500',
				}
			case 'lost':
				return {
					title: 'Игра окончена',
					message: `Правильный ответ был... Ваш счет: $${score.toLocaleString()}`,
					color: 'text-red-500',
				}
			case 'quit':
				return {
					title: 'Игра завершена',
					message: `Вы завершили игру с $${score.toLocaleString()}`,
					color: 'text-yellow-500',
				}
			default:
				return {
					title: 'Игра окончена',
					message: `Игра завершена со счетом: $${score.toLocaleString()}`,
					color: 'text-white',
				}
		}
	}

	const { title, message, color } = getGameOverMessage()

	return (
		<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
			<div className='bg-gray-800 rounded-xl p-8 max-w-md w-full mx-4 border-2 border-blue-500'>
				<h2 className='text-3xl font-bold mb-4 text-center text-white'>
					{title}
				</h2>
				<p className={`text-2xl mb-6 text-center ${color}`}>{message}</p>

				<div className='space-y-4'>
					<button
						onClick={onClose}
						className='w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full transition-colors'>
						Играть снова
					</button>

					<Link
						href='/'
						className='block w-full bg-gray-700 hover:bg-gray-600 text-white font-bold py-3 px-6 rounded-full transition-colors text-center'>
						Назад на главную
					</Link>
				</div>
			</div>
		</div>
	)
}
