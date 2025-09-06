interface LifelinePanelProps {
	onFiftyFifty: () => void
	onAudience: () => void
	onCall: () => void
	fiftyFiftyUsed: boolean
	audienceUsed: boolean
	callUsed: boolean
}

export default function LifelinePanel({
	onFiftyFifty,
	onAudience,
	onCall,
	fiftyFiftyUsed,
	audienceUsed,
	callUsed,
}: LifelinePanelProps) {
	return (
		<div className='bg-black/30 backdrop-blur-sm rounded-xl p-6'>
			<h3 className='text-xl font-bold text-white mb-4'>Lifelines</h3>

			<div className='grid grid-cols-1 md:grid-cols-3 gap-4'>
				<button
					onClick={onFiftyFifty}
					disabled={fiftyFiftyUsed}
					className={`
            p-4 rounded-xl transition-all duration-200 flex flex-col items-center
            ${
							fiftyFiftyUsed
								? 'bg-gray-700 text-gray-500 cursor-not-allowed'
								: 'bg-purple-700 hover:bg-purple-600 text-white hover:scale-[1.02]'
						}
          `}>
					<span className='text-2xl font-bold'>50:50</span>
					<span className='mt-2'>Remove 2 wrong answers</span>
				</button>

				<button
					onClick={onAudience}
					disabled={audienceUsed}
					className={`
            p-4 rounded-xl transition-all duration-200 flex flex-col items-center
            ${
							audienceUsed
								? 'bg-gray-700 text-gray-500 cursor-not-allowed'
								: 'bg-purple-700 hover:bg-purple-600 text-white hover:scale-[1.02]'
						}
          `}>
					<span className='text-2xl font-bold'>Audience</span>
					<span className='mt-2'>Get poll results</span>
				</button>

				<button
					onClick={onCall}
					disabled={callUsed}
					className={`
            p-4 rounded-xl transition-all duration-200 flex flex-col items-center
            ${
							callUsed
								? 'bg-gray-700 text-gray-500 cursor-not-allowed'
								: 'bg-purple-700 hover:bg-purple-600 text-white hover:scale-[1.02]'
						}
          `}>
					<span className='text-2xl font-bold'>Call</span>
					<span className='mt-2'>Get friend&apos;s advice</span>
				</button>
			</div>
		</div>
	)
}
