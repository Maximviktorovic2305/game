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
		// Запуск анимации появления при монтировании компонента
		setIsVisible(true)

		// Очистка для следующего перехода
		return () => setIsVisible(false)
	}, [])

	const handleOptionClick = (option: Option) => {
		// Запрет нажатия на удаленные варианты
		if (removedOptions.includes(option.id)) return

		// Запуск анимации исчезновения перед ответом
		setIsVisible(false)
		setTimeout(() => {
			onAnswer(option)
		}, 150) // Соответствие продолжительности анимации
	}

	// Сортировка вариантов по букве (A, B, C, D)
	const sortedOptions = [...question.options].sort((a, b) =>
		a.letter.localeCompare(b.letter),
	)

	return (
		<div
			className={`flex flex-col transition-opacity duration-300 ease-in-out ${
				isVisible ? 'opacity-100' : 'opacity-0'
			}`}>
			<h3 className='text-lg md:text-xl font-bold text-white mb-4 flex-shrink-0 overflow-hidden'>
				{question.text}
			</h3>

			<div className='grid grid-cols-1 md:grid-cols-2 gap-3 flex-shrink-0'>
				{sortedOptions.map((option) => (
					<button
						key={option.id}
						onClick={() => handleOptionClick(option)}
						disabled={removedOptions.includes(option.id)}
						className={`
              p-3 text-left rounded-lg transition-all duration-200 h-full text-sm
              ${
								removedOptions.includes(option.id)
									? 'bg-gray-700 text-gray-500 cursor-not-allowed'
									: 'bg-blue-700 hover:bg-blue-600 text-white hover:scale-[1.02]'
							}
            `}
						style={{ minHeight: '50px' }}>
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