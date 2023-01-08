<script lang="ts">
	import type { PageData } from './$types';
	import { onMount, onDestroy } from 'svelte';
	import { connectRoomWebsocket, leaveRoom, joinPlayer, appStateStore } from '$lib/app_state';
	import JoinForm from './join_form.svelte';

	export let data: PageData;
	
	let appState: App.AppState;
	const unsubscribe = appStateStore.subscribe((newState) => (appState = newState));

	onMount(() => {
		if (data.id == null) return;
		connectRoomWebsocket(data.id);
	});

	onDestroy(() => {
		unsubscribe();
		leaveRoom();
	});


	function onJoin(event: CustomEvent<{ name: string; color: string }>) {
		joinPlayer(event.detail.name, event.detail.color);
	}
</script>

<div>
	{#if data.id != null && appState.gameState != null}
		<div class="app">
			<h1>{data.name}</h1>
			<h2>Players</h2>
			<table>
				<thead>
					<tr>
						<th scope="col">Name</th>
						<th scope="col">Color</th>
					</tr>
				</thead>
				<tbody>
					{#each appState.gameState.players as player}
						<tr>
							<td>{player.name}</td>
							<td>{player.color}</td>
						</tr>
					{/each}
				</tbody>
			</table>
			<h2>Apples</h2>
			<table>
				<thead>
					<tr>
						<th scope="col">X</th>
						<th scope="col">Y</th>
					</tr>
				</thead>
				<tbody>
					{#each appState.gameState.apples as apple}
						<tr>
							<td>{apple.x}</td>
							<td>{apple.y}</td>
						</tr>
					{/each}
				</tbody>
			</table>
			{#if !appState.isPlaying}
				<JoinForm on:join={onJoin} />
			{/if}
		</div>
	{/if}
</div>

<style lang="scss">
	.app {
		text-align: center;
		display: grid;
		place-items: center;
		table,
		tr,
		td {
			width: max-content;
			padding: 0.5em;
			border: 1px white solid;
			border-collapse: collapse;
		}
	}
</style>
