import { ButtonHTMLAttributes } from 'react'

interface AudienceButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	used: boolean
}

export function AudienceIcon() {
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
			{/* Improved audience visualization with more people */}
			<g transform='translate(10, 12)'>
				{/* First person */}
				<circle cx='4' cy='4' r='2.5' fill='white' />
				<path d='M4 7c-1.5 0-3 1-3 2.5v1.5h6v-1.5c0-1.5-1.5-2.5-3-2.5z' fill='white' />
				
				{/* Second person */}
				<circle cx='12' cy='4' r='2.5' fill='white' />
				<path d='M12 7c-1.5 0-3 1-3 2.5v1.5h6v-1.5c0-1.5-1.5-2.5-3-2.5z' fill='white' />
				
				{/* Third person */}
				<circle cx='20' cy='4' r='2.5' fill='white' />
				<path d='M20 7c-1.5 0-3 1-3 2.5v1.5h6v-1.5c0-1.5-1.5-2.5-3-2.5z' fill='white' />
				
				{/* Fourth person */}
				<circle cx='28' cy='4' r='2.5' fill='white' />
				<path d='M28 7c-1.5 0-3 1-3 2.5v1.5h6v-1.5c0-1.5-1.5-2.5-3-2.5z' fill='white' />
			</g>
			{/* Sound waves to indicate audience noise */}
			<path d='M15 35c0 1 1 2 2 2s2-1 2-2m4 0c0 2 2 3 4 3s4-1 4-3m8 0c0 1 1 2 2 2s2-1 2-2' stroke='white' strokeWidth='1.5' fill='none' />
		</svg>
	)
}

export default function AudienceButton({ used, ...props }: AudienceButtonProps) {
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
			title='Помощь зала'>
			<AudienceIcon />
		</button>
	)
}