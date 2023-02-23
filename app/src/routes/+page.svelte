<script lang="ts">
	import type { PageServerData } from './$types';
	import GameDescription from './game_description.svelte';
	import { fetchJson } from '$lib/fetchJson';
	import RefreshButton from './refresh_button.svelte';
	import RoomPreview from '$lib/components/room_preview.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import Divider from '$lib/components/divider.svelte';
	import Icon from '@iconify/svelte';

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

<div class="header">
	<h1>Home</h1>
	<RefreshButton callback={refresh} />
</div>
{#if rooms.length != 0}
	<div class="grid">
		{#each rooms as room}
			<RoomPreview {room} />
		{/each}
	</div>
	{#if remainingRooms != 0}
		<div class="remaining-rooms">
			...and {remainingRooms} more
			<Button href="/browse" outline>
				<Icon icon="mdi:magnify" inline /> Browse All
			</Button>
		</div>
	{/if}
{:else}
	<Panel margin>
		<div class="no-rooms-error">
			<h2>There is no any room here 😬</h2>
			<p>Don't worry, you can <a href="/create-room">create one</a>!</p>
		</div>
	</Panel>
{/if}
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
	.no-rooms-error {
		text-align: center;
		color: #333;
	}
	.remaining-rooms {
		margin: 1rem;
		font-size: 1.1rem;

		display: flex;
		justify-content: flex-end;
		align-items: center;
		gap: 1rem;
	}
</style>
