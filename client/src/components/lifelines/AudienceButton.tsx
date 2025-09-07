import { ButtonHTMLAttributes } from 'react'
import { UsersRound } from 'lucide-react'

interface AudienceButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	used: boolean
}

export function AudienceIcon() {
	return (
		<div className='w-6 h-6 flex items-center justify-center cursor-pointer'>
			<div className='relative'>
				<UsersRound className='relative w-5 h-5 text-white' />
			</div>
		</div>
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
