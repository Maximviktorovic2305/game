import { ButtonHTMLAttributes } from 'react'

interface CallButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	used: boolean
}

export function CallIcon() {
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
			{/* Improved phone icon with more detail */}
			<rect x='18' y='8' width='12' height='20' rx='2' fill='white' />
			{/* Screen */}
			<rect x='20' y='10' width='8' height='12' fill='#93c5fd' />
			{/* Keypad */}
			<circle cx='22' cy='24' r='1' fill='#1e40af' />
			<circle cx='26' cy='24' r='1' fill='#1e40af' />
			<circle cx='22' cy='27' r='1' fill='#1e40af' />
			<circle cx='26' cy='27' r='1' fill='#1e40af' />
			<circle cx='22' cy='30' r='1' fill='#1e40af' />
			<circle cx='26' cy='30' r='1' fill='#1e40af' />
			{/* Antenna */}
			<line x1='24' y1='8' x2='24' y2='4' stroke='white' strokeWidth='2' />
			<circle cx='24' cy='3' r='1' fill='white' />
		</svg>
	)
}

export default function CallButton({ used, ...props }: CallButtonProps) {
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
			title='Звонок другу'>
			<CallIcon />
		</button>
	)
}