// Призы для каждого уровня игры
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
		<div className='h-full flex w-fit flex-col justify-center'>
			{[...PRIZES].reverse().map((prize) => {
				// Текущий уровень
				const isCurrent = prize.level === currentLevel
				// Уровни безопасности
				const isSafe = prize.level === 5 || prize.level === 10 || prize.level === 15

				return (
					<div
						key={prize.level}
						className={`
							py-1 px-4
							${isCurrent ? 'text-yellow-500 font-bold' : ''}
							${isSafe ? 'text-green-400' : 'text-white'}
						`}>
						<span>${prize.amount.toLocaleString()}</span>
					</div>
				)
			})}
		</div>
	)
}