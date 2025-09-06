import { Question, Option } from '@/types'

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
	const handleOptionClick = (option: Option) => {
		// Don't allow clicking on removed options
		if (removedOptions.includes(option.id)) return

		onAnswer(option)
	}

	// Sort options by letter (A, B, C, D)
	const sortedOptions = [...question.options].sort((a, b) =>
		a.letter.localeCompare(b.letter),
	)

	return (
		<div className='space-y-6'>
			<h3 className='text-xl md:text-2xl font-bold text-white'>
				{question.text}
			</h3>

			<div className='grid grid-cols-1 md:grid-cols-2 gap-4'>
				{sortedOptions.map((option) => (
					<button
						key={option.id}
						onClick={() => handleOptionClick(option)}
						disabled={removedOptions.includes(option.id)}
						className={`
              p-4 text-left rounded-xl transition-all duration-200
              ${
								removedOptions.includes(option.id)
									? 'bg-gray-700 text-gray-500 cursor-not-allowed'
									: 'bg-blue-700 hover:bg-blue-600 text-white hover:scale-[1.02]'
							}
            `}>
						<div className='flex items-center'>
							<span className='font-bold mr-3'>{option.letter}.</span>
							<span>{option.text}</span>
						</div>
					</button>
				))}
			</div>
		</div>
	)
}
