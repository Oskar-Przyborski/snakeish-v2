<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	import Button from '$lib/components/buttons/button.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	import type { Player } from './classic_game';

	export let players: Player[];
</script>

<div class="leaderboard">
	<h1>Leaderboard</h1>
	<div class="list">
		{#each players.sort((a, b) => b.snakeTail.length - a.snakeTail.length).slice(0, 5) as player}
			<Panel padding="0.8rem">
				<div class="player-item">
					<div class="name">
						{player.name}
					</div>
					<div class="score"><Icon icon="mdi:star-outline" inline /> {player.score}</div>
				</div>
			</Panel>
		{/each}
	</div>
	<div class="leave-btn">
		<Button on:click={() => dispatch('leave')} outline>Leave game</Button>
	</div>
</div>

<style lang="scss">
	h1 {
		margin: 0;
	}
	.list {
		margin: 1.5rem 0;
		display: flex;
		flex-flow: column nowrap;
		gap: 0.5rem;
		.player-item {
			display: flex;
			justify-content: space-between;
			align-items: center;
			gap: 1rem;
			.name {
				font-weight: 500;
				font-size: 1.3rem;
			}
			.score {
				font-size: 1.2rem;
			}
		}
	}
	.leave-btn {
		text-align: center;
	}
</style>
