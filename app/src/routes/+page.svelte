<script lang="ts">
	import type { PageServerData } from './$types';
	import Button from '$lib/components/buttons/button.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import SuggestedRooms from './suggested_rooms.svelte';
	import GameDescription from './game_description.svelte';
	import { fetchJson } from '$lib/fetchJson';

	export let data: PageServerData;
	let rooms = data.rooms
	let remainingRooms = data.remainingRooms

	const refresh = async () => {
		const newData = await fetchJson<{
			rooms: App.RoomPreview[];
			remainingRooms: number;
		}>('/rooms/suggested');
		rooms = newData.rooms
		remainingRooms = newData.remainingRooms
	};
</script>

<div class="top">
	<div class="panels-col">
		<SuggestedRooms rooms={data.rooms} remainingRooms={data.remainingRooms} onRefresh={refresh} />
		<GameDescription />
	</div>
	<div class="panels-col">
		<Panel>
			<div class="centered-panel">
				<h1>Wanna have own room?</h1>
				<div class="btn-wrapper">
					<Button href="/create-room">
						<b><Icon icon="mdi:plus" inline /> Create</b>
					</Button>
				</div>
			</div>
		</Panel>
		<Panel>
			<div class="centered-panel">
				<h2>Do you like this game?</h2>
				<Button><Icon icon="mdi:coffee" inline /> Buy me a coffee!</Button>
			</div>
		</Panel>
	</div>
</div>

<style lang="scss">
	.top {
		display: grid;
		grid-template-columns: 75% 25%;
		gap: 1rem;
		align-items: flex-start;
	}

	.panels-col {
		display: flex;
		flex-flow: column;
		gap: 1rem;
	}
	.centered-panel {
		text-align: center;
		h1,
		h2 {
			margin: 0;
			margin-bottom: 1rem;
		}
		.btn-wrapper {
			font-size: 1.2rem;
		}
	}
</style>
