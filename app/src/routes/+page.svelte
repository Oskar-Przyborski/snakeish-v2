<script lang="ts">
	import backendRequest, { responseToMap } from '$lib/backend_request';
	import Button from '$lib/components/button.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import type { PageServerData } from './$types';
	import SuggestedRooms from './suggested_rooms.svelte';
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

{#if data.isOnline}
	<div class="top">
		<SuggestedRooms {rooms} onRefresh={refresh} />
		<div class="right">
			<Panel>
				<div class="centered-panel">
					<h1>Wanna have own room?</h1>
					<div class="btn-wrapper">
						<Button href="/create-room">
							<b><Icon icon="ic:sharp-plus" inline /> Create</b>
						</Button>
					</div>
				</div>
			</Panel>
			<Panel>
				<div class="centered-panel">
					<h2>Do you like this game?</h2>
					<Button><Icon icon="ic:baseline-coffee" inline/> Buy me a coffee!</Button>
				</div>
			</Panel>
		</div>
	</div>
	<!-- <GameDescription />s -->
{:else}
	<h1>Server offline</h1>
	<p>Server is currently offline.</p>
	<ol>
		<li>Wait a few seconds and try reloading the page, maybe the server is still waking up.</li>
		<li>If it doesn't work, the server is probably completely down</li>
	</ol>
{/if}

<style lang="scss">
	.top {
		display: grid;
		grid-template-columns: 75% 25%;
		gap: 1rem;
		align-items: flex-start;
	}

	.right {
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
