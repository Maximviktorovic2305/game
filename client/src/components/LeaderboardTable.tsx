interface LeaderboardEntry {
	id: number
	score: number
	status: 'in_progress' | 'won' | 'lost' | 'quit'
	start_time: string
}

interface LeaderboardTableProps {
	sessions: LeaderboardEntry[]
}

export default function LeaderboardTable({ sessions }: LeaderboardTableProps) {
	// Сортировка сессий по убыванию счета
	const sortedSessions = [...sessions].sort((a, b) => b.score - a.score)

	// Функция для получения текста статуса на русском языке
	const getStatusText = (status: string) => {
		switch (status) {
			case 'won':
				return 'Выиграл'
			case 'lost':
				return 'Проиграл'
			case 'quit':
				return 'Завершил'
			default:
				return 'В процессе'
		}
	}

	return (
		<div className='overflow-x-auto'>
			<table className='w-full'>
				<thead>
					<tr className='border-b border-gray-700'>
						<th className='text-left py-3 px-4 text-blue-400'>Место</th>
						<th className='text-left py-3 px-4 text-blue-400'>Счет</th>
						<th className='text-left py-3 px-4 text-blue-400'>Статус</th>
						<th className='text-left py-3 px-4 text-blue-400'>Дата</th>
					</tr>
				</thead>
				<tbody>
					{sortedSessions.map((session, index) => (
						<tr
							key={session.id}
							className='border-b border-gray-700 hover:bg-gray-700/50 transition-colors'>
							<td className='py-3 px-4 text-white font-bold'>
								{index === 0
									? '🥇'
									: index === 1
									? '🥈'
									: index === 2
									? '🥉'
									: `#${index + 1}`}
							</td>
							<td className='py-3 px-4 text-white'>
								${session.score.toLocaleString()}
							</td>
							<td className='py-3 px-4'>
								<span
									className={`
                  px-2 py-1 rounded-full text-xs font-bold
                  ${
										session.status === 'won'
											? 'bg-green-700 text-white'
											: session.status === 'lost'
											? 'bg-red-700 text-white'
											: 'bg-blue-700 text-white'
									}
                `}>
									{getStatusText(session.status)}
								</span>
							</td>
							<td className='py-3 px-4 text-gray-300'>
								{new Date(session.start_time).toLocaleDateString('ru-RU')}
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	)
}