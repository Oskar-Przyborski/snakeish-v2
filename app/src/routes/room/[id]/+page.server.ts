import type { PageServerLoad } from './$types';
import { fetchJson } from '$lib/fetchJson';

export const load = (async ({ params, fetch }) => {
	return await fetchJson<App.RoomPreview>(`/get-room?id=${params.id}`, { fetcher: fetch });
}) satisfies PageServerLoad;
