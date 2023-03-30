export default new Map<string, { tag: string; title: string; name: string; description: string }>([
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
		'battleroyale-eat&win',
		{
			tag: 'battle-royale',
			title: 'Eat&Win',
			name: 'battleroyale-eat&win',
			description: 'Popular Battle Royale! Compete with friends, escape from zone, eat apples and win!'
		}
	],
	[
		'battleroyale-zoneshark',
		{
			tag: 'battle-royale',
			title: 'Zone Shark',
			name: 'battleroyale-zoneshark',
			description: "Don't touch the zone! It eats you instantly."
		}
	],
]);
