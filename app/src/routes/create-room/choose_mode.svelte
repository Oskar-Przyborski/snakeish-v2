<script lang="ts">
	import GamemodeCard from './gamemode_card.svelte';
	import modes from '$lib/modes';
	import { store } from '$lib/room_creation_state';
</script>

<div class="modes">
	{#each Array.from(modes.keys()) as mode}
		<GamemodeCard
			configName={mode}
			selected={$store.configName == mode}
			on:select={() => ($store.configName = mode)}
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
