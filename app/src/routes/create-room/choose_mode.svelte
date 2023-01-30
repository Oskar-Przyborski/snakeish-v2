<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/button.svelte';
	import GamemodeCard from '$lib/components/gamemode_card.svelte';
	import modes from '$lib/modes';
	const dispatch = createEventDispatcher();

	export let selectedConfig: string;

	const onContinue = () => {
		dispatch('continue');
	};
</script>

<div class="wrapper">
	<h1>Which mode do you want to play on?</h1>
	<div>
		<div class="modes">
			{#each Array.from(modes.keys()) as mode}
				<GamemodeCard
					configName={mode}
					selected={selectedConfig == mode}
					on:select={() => (selectedConfig = mode)}
				/>
			{/each}
		</div>
	</div>
	<div class="btn-wrapper">
		<Button on:click={onContinue}>Continue</Button>
	</div>
</div>

<style lang="scss">
	.wrapper {
		display: grid;
		grid-template-rows: max-content 1fr max-content;
	}
	h1 {
		text-align: center;
		margin: 0;
		margin-bottom: 2rem;
	}
	.modes {
		display: grid;
		grid-template-columns: 1fr;
		gap: 1.5rem;
	}
	@media (min-width: 768px) {
		.modes {
			grid-template-columns: 1fr 1fr;
		}
	}

	.btn-wrapper {
		text-align: center;
		margin-top: 2rem;
	}
</style>
