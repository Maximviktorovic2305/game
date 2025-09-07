import FiftyFiftyButton from '@/components/lifelines/FiftyFiftyButton'
import AudienceButton from '@/components/lifelines/AudienceButton'
import CallButton from '@/components/lifelines/CallButton'

interface LifelinesAreaProps {
	fiftyFiftyUsed: boolean
	audienceUsed: boolean
	callUsed: boolean
	onFiftyFifty: () => void
	onAudience: () => void
	onCall: () => void
}

export default function LifelinesArea({
	fiftyFiftyUsed,
	audienceUsed,
	callUsed,
	onFiftyFifty,
	onAudience,
	onCall,
}: LifelinesAreaProps) {
	return (
		<div className='absolute top-4 left-4 z-20 flex space-x-4'>
			<FiftyFiftyButton used={fiftyFiftyUsed} onClick={onFiftyFifty} />
			<AudienceButton used={audienceUsed} onClick={onAudience} />
			<CallButton used={callUsed} onClick={onCall} />
		</div>
	)
}