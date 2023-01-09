<script lang="ts">
	import type { PageData } from './$types';
	import { onMount, onDestroy } from 'svelte';
	import { connectRoomWebsocket, leaveRoom, joinPlayer, appStateStore } from '$lib/app_state';
	import JoinForm from './join_form.svelte';
	import Leaderboard from './leaderboard.svelte';
	import Button from '$lib/components/button.svelte';
	import GameRenderer from './game_renderer.svelte';

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

{#if data.id != null && appState.gameState != null}
	<div class="page">
		<h1>{data.name}</h1>
		<div class="game">
			<div class="left-side">
				<GameRenderer />
			</div>
			<div class="right-side">
				{#if !appState.isPlaying}
					<JoinForm on:join={onJoin} />
				{:else}
					<Leaderboard players={appState.gameState.players} />
					<Button>Leave game</Button>
				{/if}
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	.page {
		display: grid;
		place-items: center;

		.game {
			width: 100%;
			border-radius: 1em;
			overflow: hidden;

			display: grid;
			grid-template-columns: 2fr 1fr;

			.left-side {
				display:grid;
				place-items: center;

				padding: 1em;
				background-color: #4a525a;
			}

			.right-side {
				background-color: #d62246;
				padding: 1em;
				display: flex;
				flex-direction: column;
				justify-content: space-between;
			}
		}
	}
</style>
