import type { PageLoad } from './$types';
import backendRequest from '$lib/backend_request';

export const load = (({ fetch }) => {
	async function createRoom(data: { roomName: string; configName: string }) {
		await backendRequest(fetch, '/create-classic-room', data);
	}

	return {
		createRoom,
		modes: [
			{
				tag: 'classic',
				title: 'Casual',
				name: 'classic-casual',
				description: 'Just a normal, casual multiplayer snake'
			},
			{
				tag: 'classic',
				title: 'Huuge',
				name: 'classic-huuge',
				description: 'Quite big map with many players and apples'
			},
			{
				tag: 'battle royale',
				title: 'Eat & Win',
				name: 'battle-royale',
				description: 'Popular battle royale mode, with one life and shrinking zones'
			}
		],
		steps: [
			{
				name: 'Name',
				description: 'Enhance your room with your own name!'
			},
			{
				name: 'Mode',
				description: 'Select game mode, you most like!'
			},
			{
				name: 'PIN Code',
				description: 'Protect your room against uninvited players'
			},
			{
				name: 'Summary',
				description: 'Make sure of everything and jump right into the game!'
			}
		]
	};
}) satisfies PageLoad;
