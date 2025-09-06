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
					title: 'Congratulations!',
					message: `You've won $${score.toLocaleString()}!`,
					color: 'text-green-500',
				}
			case 'lost':
				return {
					title: 'Game Over',
					message: `The correct answer was... Your score: $${score.toLocaleString()}`,
					color: 'text-red-500',
				}
			case 'quit':
				return {
					title: 'Game Quit',
					message: `You quit the game with $${score.toLocaleString()}`,
					color: 'text-yellow-500',
				}
			default:
				return {
					title: 'Game Over',
					message: `Game finished with score: $${score.toLocaleString()}`,
					color: 'text-white',
				}
		}
	}

	const { title, message, color } = getGameOverMessage()

	return (
		<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
			<div className='bg-white rounded-xl p-8 max-w-md w-full mx-4'>
				<h2 className='text-3xl font-bold mb-4 text-center'>{title}</h2>
				<p className={`text-2xl mb-6 text-center ${color}`}>{message}</p>

				<div className='space-y-4'>
					<button
						onClick={() => window.location.reload()}
						className='w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full transition-colors'>
						Play Again
					</button>

					<Link
						href='/'
						className='block w-full bg-gray-600 hover:bg-gray-700 text-white font-bold py-3 px-6 rounded-full transition-colors text-center'>
						Back to Home
					</Link>
				</div>
			</div>
		</div>
	)
}
