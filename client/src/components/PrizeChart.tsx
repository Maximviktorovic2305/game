const PRIZES = [
	{ level: 1, amount: 100 },
	{ level: 2, amount: 200 },
	{ level: 3, amount: 300 },
	{ level: 4, amount: 500 },
	{ level: 5, amount: 1000 },
	{ level: 6, amount: 2000 },
	{ level: 7, amount: 4000 },
	{ level: 8, amount: 8000 },
	{ level: 9, amount: 16000 },
	{ level: 10, amount: 32000 },
	{ level: 11, amount: 64000 },
	{ level: 12, amount: 125000 },
	{ level: 13, amount: 250000 },
	{ level: 14, amount: 500000 },
	{ level: 15, amount: 1000000 },
]

interface PrizeChartProps {
	currentLevel: number
}

export default function PrizeChart({ currentLevel }: PrizeChartProps) {
	return (
		<div className='bg-black/30 backdrop-blur-sm rounded-xl p-6'>
			<h3 className='text-xl font-bold text-white mb-4 text-center'>
				Prize Chart
			</h3>

			<div className='space-y-2'>
				{PRIZES.map((prize) => {
					const isCurrent = prize.level === currentLevel
					const isSafe = prize.level <= 5 || prize.level === 10

					return (
						<div
							key={prize.level}
							className={`
                p-3 rounded-lg flex justify-between items-center
                ${
									isCurrent
										? 'bg-yellow-500 text-black font-bold'
										: isSafe
										? 'bg-green-700 text-white'
										: 'bg-blue-900 text-white'
								}
              `}>
							<span>Question {prize.level}</span>
							<span>${prize.amount.toLocaleString()}</span>
						</div>
					)
				})}
			</div>
		</div>
	)
}
