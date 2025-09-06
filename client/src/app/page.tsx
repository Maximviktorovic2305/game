'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Button } from '@/components/ui/button'

export default function Home() {
	const [gameStarted, setGameStarted] = useState(false)

	const handleStartGame = () => {
		setGameStarted(true)
	}

	return (
		<div className='min-h-screen bg-gradient-to-b from-blue-900 to-purple-900 flex flex-col items-center justify-center p-4'>
			<div className='text-center max-w-2xl'>
				<h1 className='text-4xl md:text-6xl font-bold text-yellow-400 mb-6'>
					Who Wants to Be a Millionaire?
				</h1>

				<div className='bg-black/30 backdrop-blur-sm rounded-xl p-8 mb-8'>
					<p className='text-white text-lg mb-6'>
						Test your knowledge and climb the prize ladder to win 1 million!
					</p>

					{!gameStarted ? (
						<Button
							onClick={handleStartGame}
							className='bg-yellow-500 hover:bg-yellow-600 text-black font-bold py-3 px-8 text-xl rounded-full transition-all duration-300 transform hover:scale-105'>
							Start Game
						</Button>
					) : (
						<div className='space-y-4'>
							<Link href='/game' className='block'>
								<Button className='bg-green-500 hover:bg-green-600 text-white font-bold py-3 px-8 text-xl rounded-full transition-all duration-300 transform hover:scale-105'>
									Continue to Game
								</Button>
							</Link>
							<p className='text-gray-300'>
								Click above to continue to the game screen
							</p>
						</div>
					)}
				</div>

				<div className='mt-8'>
					<Link href='/leaderboard' className='block'>
						<Button
							variant='outline'
							className='border-white text-white hover:bg-white hover:text-black'>
							View Leaderboard
						</Button>
					</Link>
				</div>
			</div>
		</div>
	)
}
