export default new Map<
	string,
	{ tag: string; title: string; name: string; description: string }
>([
	[
		'classic-casual',
		{
			tag: 'classic',
			title: 'Casual',
			name: 'classic-casual',
			description: 'Just a normal, casual multiplayer snake'
		}
	],
	[
		'classic-huuge',
		{
			tag: 'classic',
			title: 'Huuge',
			name: 'classic-huuge',
			description: 'Quite big map with many players and apples'
		}
	],
	[
		'battle-royale-eat-n-win',
		{
			tag: 'battle royale',
			title: 'Eat & Win',
			name: 'battle-royale-eat-n-win',
			description: 'Popular battle royale mode, with one life and shrinking zones'
		}
	]
]);
