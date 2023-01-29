<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/button.svelte';
	import GamemodeCard from '$lib/components/gamemode_card.svelte';
	const dispatch = createEventDispatcher();

	export let gameModes: { title: string; tag: string; description: string; name: string }[];
	export let selectedConfig: string;

	const onContinue = () => {
		dispatch('continue');
	};
</script>

<div class="wrapper">
	<h1>Which mode do you want to play on?</h1>
	<div class="modes">
		{#each gameModes as gm}
			<GamemodeCard
				title={gm.title}
				tag={gm.tag}
				description={gm.description}
				selected={gm.name == selectedConfig}
				on:select={() => (selectedConfig = gm.name)}
			/>
		{/each}
	</div>
	<div class="btn-wrapper">
		<Button on:click={onContinue}>Continue</Button>
	</div>
</div>

<style lang="scss">
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
