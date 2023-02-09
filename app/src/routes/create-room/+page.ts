import type { PageLoad } from './$types';
export const load = (() => {
	return {
		steps: [
			{
				name: 'Name',
				description: 'Enhance your room with your own name!',
				iconEmpty: 'mdi:pencil-outline',
				iconFull: 'mdi:pencil',
				stepView: {
					title: "Let's start with name",
					description: 'What would you like to name your room? Let it be short and descriptive'
				}
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!',
				iconEmpty: 'mdi:gamepad-variant-outline',
				iconFull: 'mdi:gamepad-variant',
				stepView: {
					title: 'What mode fits you the best?',
					description: 'There are many modes to choose from. Choose your favorite!'
				}
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players',
				iconEmpty: 'mdi:lock-outline',
				iconFull: 'mdi:lock',
				stepView: {
					title: 'Want to protect your room?',
					description: 'Players will have to enter the PIN code before entering the game.'
				}
			}
		]
	};
}) satisfies PageLoad;
