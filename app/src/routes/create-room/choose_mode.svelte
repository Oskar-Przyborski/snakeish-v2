<script lang="ts">
	import GamemodeCard from './gamemode_card.svelte';
	import modes from '$lib/modes';

	export let selectedConfig: string;
</script>

<div class="modes">
	{#each Array.from(modes.keys()) as mode}
		<GamemodeCard
			configName={mode}
			selected={selectedConfig == mode}
			on:select={() => (selectedConfig = mode)}
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
