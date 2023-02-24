import { fetchJson } from '$lib/fetchJson';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch }) => {
	const rooms = await fetchJson<App.RoomPreview[]>('/rooms', { fetcher: fetch });

	return {
		rooms
	};
}) satisfies PageServerLoad;
