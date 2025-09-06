import { Question, Option } from '@/types'
import { useState, useEffect } from 'react'

interface QuestionCardProps {
	question: Question
	onAnswer: (option: Option) => void
	removedOptions: number[]
}

export default function QuestionCard({
	question,
	onAnswer,
	removedOptions,
}: QuestionCardProps) {
	const [isVisible, setIsVisible] = useState(false)

	useEffect(() => {
		// Trigger fade-in animation when component mounts
		setIsVisible(true)

		// Clean up for next transition
		return () => setIsVisible(false)
	}, [])

	const handleOptionClick = (option: Option) => {
		// Don't allow clicking on removed options
		if (removedOptions.includes(option.id)) return

		// Trigger fade-out animation before answering
		setIsVisible(false)
		setTimeout(() => {
			onAnswer(option)
		}, 150) // Match the animation duration
	}

	// Sort options by letter (A, B, C, D)
	const sortedOptions = [...question.options].sort((a, b) =>
		a.letter.localeCompare(b.letter),
	)

	return (
		<div
			className={`flex flex-col transition-opacity duration-300 ease-in-out ${
				isVisible ? 'opacity-100' : 'opacity-0'
			}`}>
			<h3 className='text-xl md:text-2xl font-bold text-white mb-4 flex-shrink-0 overflow-hidden'>
				{question.text}
			</h3>

			<div className='grid grid-cols-1 md:grid-cols-2 gap-3 flex-shrink-0'>
				{sortedOptions.map((option) => (
					<button
						key={option.id}
						onClick={() => handleOptionClick(option)}
						disabled={removedOptions.includes(option.id)}
						className={`
              p-3 text-left rounded-lg transition-all duration-200 h-full
              ${
								removedOptions.includes(option.id)
									? 'bg-gray-700 text-gray-500 cursor-not-allowed'
									: 'bg-blue-700 hover:bg-blue-600 text-white hover:scale-[1.02]'
							}
            `}
						style={{ minHeight: '60px' }}>
						<div className='flex items-center h-full'>
							<span className='font-bold mr-2 flex-shrink-0'>
								{option.letter}.
							</span>
							<span className='overflow-hidden text-ellipsis'>
								{option.text}
							</span>
						</div>
					</button>
				))}
			</div>
		</div>
	)
}
