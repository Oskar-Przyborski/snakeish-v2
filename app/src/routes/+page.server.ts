import { fetchJson } from '$lib/fetchJson';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	const { rooms, remainingRooms } = await fetchJson<{
		rooms: App.RoomPreview[];
		remainingRooms: number;
	}>('/rooms/suggested', {
		fetcher: fetch,
		default: { remainingRooms: 0, rooms: [] }
	});

	return { rooms: rooms ?? [], remainingRooms: remainingRooms };
}) satisfies PageServerLoad;
