import type { PageServerLoad } from './$types';
import { fetchJson } from '$lib/fetchJson';

export const load = (async ({ params, fetch }) => {
	const rooms = await fetchJson<App.RoomPreview>(`/room/${params.id}`, {
		fetcher: fetch,
		onError: (error) => {
			throw error;
		}
	});

	return rooms;
}) satisfies PageServerLoad;
