import { ButtonHTMLAttributes } from 'react'
import { DiamondPercent } from 'lucide-react'

interface FiftyFiftyButtonProps
	extends ButtonHTMLAttributes<HTMLButtonElement> {
	used: boolean
}

export function FiftyFiftyIcon() {
	return (
		<div className='w-6 h-6 flex items-center justify-center cursor-pointer'>
			<div className='relative'>
				<DiamondPercent className='relative w-5 h-5 text-white' />
			</div>
		</div>
	)
}

export default function FiftyFiftyButton({
	used,
	...props
}: FiftyFiftyButtonProps) {
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
