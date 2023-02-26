import type { PageLoad } from './$types';
export const load = (() => {
	return {
		steps: [
			{
				name: 'Name',
				description: 'Enhance your room with your own name!',
				iconEmpty: 'material-symbols:edit-outline-rounded',
				iconFull: 'material-symbols:edit-rounded',
				stepView: {
					title: "Let's start with name",
					description: 'What would you like to name your room? Let it be short and descriptive'
				}
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!',
				iconEmpty: 'material-symbols:gamepad-outline-rounded',
				iconFull: 'material-symbols:gamepad-rounded',
				stepView: {
					title: 'What mode fits you the best?',
					description: 'There are many modes to choose from. Choose your favorite!'
				}
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players',
				iconEmpty: 'material-symbols:lock-outline',
				iconFull: 'material-symbols:lock',
				stepView: {
					title: 'Want to protect your room?',
					description: 'Players will have to enter the PIN code before entering the game.'
				}
			}
		]
	};
}) satisfies PageLoad;
