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
	// Sort sessions by score descending
	const sortedSessions = [...sessions].sort((a, b) => b.score - a.score)

	return (
		<div className='overflow-x-auto'>
			<table className='w-full'>
				<thead>
					<tr className='border-b border-gray-700'>
						<th className='text-left py-3 px-4 text-yellow-400'>Rank</th>
						<th className='text-left py-3 px-4 text-yellow-400'>Score</th>
						<th className='text-left py-3 px-4 text-yellow-400'>Status</th>
						<th className='text-left py-3 px-4 text-yellow-400'>Date</th>
					</tr>
				</thead>
				<tbody>
					{sortedSessions.map((session, index) => (
						<tr
							key={session.id}
							className='border-b border-gray-800 hover:bg-gray-800/50 transition-colors'>
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
											? 'bg-green-500 text-white'
											: session.status === 'lost'
											? 'bg-red-500 text-white'
											: 'bg-yellow-500 text-black'
									}
                `}>
									{session.status.charAt(0).toUpperCase() +
										session.status.slice(1)}
								</span>
							</td>
							<td className='py-3 px-4 text-gray-300'>
								{new Date(session.start_time).toLocaleDateString()}
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	)
}
