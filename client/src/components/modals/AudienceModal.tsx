interface AudienceModalProps {
	percentages: Record<string, number>
	onClose: () => void
}

export default function AudienceModal({ percentages, onClose }: AudienceModalProps) {
	return (
		<div className='fixed inset-0 bg-black/70 flex items-center justify-center z-50'>
			<div className='bg-gray-800 rounded-xl p-6 max-w-md w-full mx-4 border-2 border-blue-500'>
				<h3 className='text-2xl font-bold mb-4 text-center text-blue-400'>
					Опрос зала
				</h3>
				<div className='space-y-4'>
					{Object.entries(percentages).map(([letter, percentage]) => (
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
					))}
				</div>
				<div className='mt-6 text-center'>
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