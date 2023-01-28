<script lang="ts">
	import backendRequest, { responseToMap } from '$lib/backend_request';
	import type { PageServerData } from './$types';
	import AvailableRooms from './available_rooms.svelte';
	import GameDescription from './game_description.svelte';

	export let data: PageServerData;
	let rooms = data.data ? data.data : [];

	const refresh = async () => {
		const { data } = await responseToMap<App.RoomPreview>(backendRequest(fetch, '/get-rooms'));
		rooms = [];
		if (data != null) {
			data.forEach((v) => rooms.push(v));
		}
	};
</script>

<div class="snakeish-logo">
	<img alt="Snakeish Logo" src="/logo-long-white.png" width="260" />
</div>
<AvailableRooms {rooms} onRefresh={refresh} />
<GameDescription />

<style lang="scss">
	.snakeish-logo {
		display: grid;
		place-items: center;
		padding: 0.5em 0;
	}
</style>
