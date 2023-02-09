<script lang="ts">
	import { store } from './room_creation_state';
	import { createEventDispatcher, onDestroy } from 'svelte';
	const dispatch = createEventDispatcher();
	import GamemodeCard from './gamemode_card.svelte';
	import modes from '$lib/modes';
	import Button from '$lib/components/buttons/button.svelte';
	import TextButton from '$lib/components/buttons/text_button.svelte';

	const onContinue = () => dispatch('continue');
	const onBack = () => dispatch('back');

	let btnDisabled = false;

	const updateBtnState = async (state: App.RoomCreationState) => {
		if (state.configName == null) {
			btnDisabled = true;
			return;
		}
		btnDisabled = false;
	};

	const unsubscibe = store.subscribe(updateBtnState);
	onDestroy(unsubscibe);
</script>

<div class="modes">
	{#each Array.from(modes.keys()) as mode}
		<GamemodeCard
			configName={mode}
			selected={$store.configName == mode}
			on:select={() => ($store.configName = $store.configName == mode ? null : mode)}
		/>
	{/each}
</div>
<div class="continue-section">
	<TextButton on:click={onBack}>Back</TextButton>
	<Button on:click={onContinue} disabled={btnDisabled}>Continue</Button>
</div>

<style lang="scss">
	.modes {
		display: flex;
		flex-flow: column;
		gap: 1.5rem;
	}
	.continue-section {
		margin-top: 1.5rem;
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		align-items: center;
	}
</style>
