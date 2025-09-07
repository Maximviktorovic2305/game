interface CallModalProps {
	advice: string
	onClose: () => void
}

export default function CallModal({ advice, onClose }: CallModalProps) {
	return (
		<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
			<div className='bg-gray-800 rounded-xl p-6 max-w-md w-full mx-4 border-2 border-blue-500'>
				<h3 className='text-2xl font-bold mb-4 text-center text-blue-400'>
					Совет друга
				</h3>
				<p className='text-lg mb-6 text-center text-white'>{advice}</p>
				<div className='text-center'>
					<button
						onClick={onClose}
						className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-full'>
						Закрыть
					</button>
				</div>
			</div>
		</div>
	)
}