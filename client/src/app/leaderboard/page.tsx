'use client'

import LeaderboardTable from '@/components/LeaderboardTable'
import { useLeaderboard } from '@/lib/hooks/useGameSession'
import { LeaderboardEntry } from '@/types'
import Link from 'next/link'

export default function LeaderboardPage() {
	const { data: leaderboard, isLoading, error } = useLeaderboard()

	// Transform the leaderboard data to only include serializable properties
	const serializableLeaderboard =
		leaderboard?.map((session: LeaderboardEntry) => ({
			id: session.id,
			score: session.score,
			status: session.status,
			start_time: session.start_time,
		})) || []

	return (
		<div className='min-h-screen bg-gradient-to-b from-blue-900 to-purple-900 p-4'>
			<div className='max-w-4xl mx-auto'>
				<div className='text-center mb-8'>
					<h1 className='text-3xl md:text-4xl font-bold text-yellow-400 mb-2'>
						Leaderboard
					</h1>
					<p className='text-white'>
						Top players who reached the highest scores
					</p>
				</div>

				<div className='bg-black/30 backdrop-blur-sm rounded-xl p-6'>
					{isLoading ? (
						<div className='text-center py-12'>
							<p className='text-white text-xl'>Loading leaderboard...</p>
						</div>
					) : error ? (
						<div className='text-center py-12'>
							<p className='text-red-400 text-xl'>Error loading leaderboard</p>
						</div>
					) : serializableLeaderboard && serializableLeaderboard.length > 0 ? (
						<LeaderboardTable sessions={serializableLeaderboard} />
					) : (
						<div className='text-center py-12'>
							<p className='text-white text-xl'>
								No scores yet. Be the first to play!
							</p>
						</div>
					)}
				</div>

				<div className='mt-8 text-center'>
					<Link
						href='/'
						className='inline-block bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full transition-colors'>
						Back to Home
					</Link>
				</div>
			</div>
		</div>
	)
}
