import { fetchJson } from '$lib/fetchJson';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	const rooms = await fetchJson<App.RoomPreview[]>('/rooms', { fetcher: fetch });

	return {
		rooms
		// rooms: [
		// 	{
		// 		id: '1',
		// 		maxPlayers: 4,
		// 		modeName: 'Casual',
		// 		modeTag: 'classic',
		// 		name: 'room1',
		// 		pinEnabled: false,
		// 		players: 2
		// 	} satisfies App.RoomPreview,
		// 	{
		// 		id: '2',
		// 		maxPlayers: 4,
		// 		modeName: 'Casual',
		// 		modeTag: 'classic',
		// 		name: 'room2',
		// 		pinEnabled: false,
		// 		players: 2
		// 	} satisfies App.RoomPreview,
		// 	{
		// 		id: '3',
		// 		maxPlayers: 4,
		// 		modeName: 'Casual',
		// 		modeTag: 'classic',
		// 		name: 'room3',
		// 		pinEnabled: false,
		// 		players: 2
		// 	} satisfies App.RoomPreview,
		// 	{
		// 		id: '4',
		// 		maxPlayers: 4,
		// 		modeName: 'Casual',
		// 		modeTag: 'classic',
		// 		name: 'room4',
		// 		pinEnabled: false,
		// 		players: 2
		// 	} satisfies App.RoomPreview,
		// ]
	};
}) satisfies PageServerLoad;
