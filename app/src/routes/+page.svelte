<script lang="ts">
	import type { PageServerData } from './$types';
	import SuggestedRooms from './suggested_rooms.svelte';
	import GameDescription from './game_description.svelte';
	import { fetchJson } from '$lib/fetchJson';

	export let data: PageServerData;
	let rooms = data.rooms;
	let remainingRooms = data.remainingRooms;

	const refresh = async () => {
		const newData = await fetchJson<{
			rooms: App.RoomPreview[];
			remainingRooms: number;
		}>('/rooms/suggested');
		rooms = newData.rooms ?? [];
		remainingRooms = newData.remainingRooms ?? 0;
	};
</script>

<SuggestedRooms {rooms} {remainingRooms} onRefresh={refresh} />
<GameDescription />
