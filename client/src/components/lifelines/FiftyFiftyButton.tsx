import { ButtonHTMLAttributes } from 'react'

interface FiftyFiftyButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	used: boolean
}

export function FiftyFiftyIcon() {
	return (
		<svg
			xmlns='http://www.w3.org/2000/svg'
			viewBox='0 0 48 48'
			className='w-12 h-12'>
			{/* Background circle */}
			<circle
				cx='24'
				cy='24'
				r='22'
				fill='#3b82f6'
			/>
			<circle
				cx='24'
				cy='24'
				r='18'
				fill='#1d4ed8'
			/>
			{/* 50/50 text with better styling */}
			<text
				x='24'
				y='29'
				textAnchor='middle'
				fill='white'
				fontSize='18'
				fontWeight='bold'
				fontFamily='Arial, sans-serif'>
				50
			</text>
			<text
				x='24'
				y='35'
				textAnchor='middle'
				fill='white'
				fontSize='10'
				fontWeight='bold'
				fontFamily='Arial, sans-serif'>
				50
			</text>
		</svg>
	)
}

export default function FiftyFiftyButton({ used, ...props }: FiftyFiftyButtonProps) {
	return (
		<button
			{...props}
			disabled={used}
			className={`
				p-2 rounded-full transition-all duration-200 flex items-center justify-center transform hover:scale-105
				${
					used
						? 'bg-gray-600 text-gray-400 cursor-not-allowed'
						: 'bg-blue-600 text-white hover:bg-blue-500'
				}
			`}
			title='Подсказка 50:50'>
			<FiftyFiftyIcon />
		</button>
	)
}