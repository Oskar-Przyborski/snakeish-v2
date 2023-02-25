<script lang="ts">
	import type { PageServerData } from './$types';
	import { fetchJson } from '$lib/fetchJson';
	import Panel from '$lib/components/panel.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import Divider from '$lib/components/divider.svelte';
	import RefreshButton from '$lib/components/refresh_button.svelte';
	import SuggestedRooms from './suggested_rooms.svelte';
	import GameDescription from './game_description.svelte';

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

<svelte:head>
	<title>Snakeish - Online Multiplayer Snake Game</title>
	<meta name="title" content="Snakeish - Online Multiplayer Snake Game" />
	<meta
		name="description"
		content="Snakeish - Best free online multiplayer snake game with many gamemodes. Play with friends!"
	/>
</svelte:head>

<div class="header">
	<h1>Home</h1>
	<RefreshButton callback={refresh} />
</div>
<SuggestedRooms {rooms} {remainingRooms} />
<Divider />

<Panel margin>
	<div class="own-room">
		<span>Want to have your own room?</span>
		<Button href="/create-room" outline>Create</Button>
	</div>
</Panel>

<div style="margin: 1rem;">
	<GameDescription />
</div>

<style lang="scss">
	.header {
		margin: 1rem;
		display: flex;
		justify-content: space-between;
		h1 {
			margin: 0;
		}
	}

	.own-room {
		display: flex;
		flex-flow: row wrap;
		justify-content: space-between;
		align-items: center;
		gap: 1rem;
		span {
			font-size: 1.3rem;
			font-weight: 500;
		}
	}
</style>
