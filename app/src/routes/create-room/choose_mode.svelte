<script lang="ts">
	import GamemodeCard from './gamemode_card.svelte';
	import modes from '$lib/modes';
	import { state } from '$lib/room_creation_state';
</script>

<div class="modes">
	{#each Array.from(modes.keys()) as mode}
		<GamemodeCard
			configName={mode}
			selected={$state.configName == mode}
			on:select={() => ($state.configName = mode)}
		/>
	{/each}
</div>

<style lang="scss">
	.modes {
		display: flex;
		flex-flow: column;
		gap: 1.5rem;
	}
</style>
