'use client'

import { useState } from 'react'
import Link from 'next/link'
import { Button } from '@/components/ui/button'

export default function Home() {
	const [gameStarted, setGameStarted] = useState(false)

	const handleStartGame = () => {
		// Установка состояния начала игры
		setGameStarted(true)
	}

	return (
		<div className='min-h-screen bg-gray-900 flex flex-col items-center justify-center p-4'>
			<div className='text-center max-w-2xl'>
				<h1 className='text-4xl md:text-6xl font-bold text-blue-400 mb-6'>
					Кто хочет стать миллионером?
				</h1>

				<div className='bg-gray-800 rounded-xl p-8 mb-8'>
					<p className='text-white text-lg mb-6'>
						Проверьте свои знания и поднимайтесь по лестнице призов, чтобы
						выиграть 1 миллион!
					</p>

					{!gameStarted ? (
						<Button
							onClick={handleStartGame}
							className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-8 text-xl rounded-full transition-all duration-300'>
							Начать игру
						</Button>
					) : (
						<div className='space-y-4'>
							<Link href='/game' className='block'>
								<Button className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-8 text-xl rounded-full transition-all duration-300'>
									Продолжить игру
								</Button>
							</Link>
							<p className='text-gray-300'>
								Нажмите выше, чтобы перейти к экрану игры
							</p>
						</div>
					)}
				</div>

				<div className='mt-8'>
					<Link href='/leaderboard' className='block'>
						<Button
							variant='outline'
							className='border-blue-500 text-blue-400 hover:bg-blue-500 hover:text-white'>
							Посмотреть таблицу лидеров
						</Button>
					</Link>
				</div>
			</div>
		</div>
	)
}
