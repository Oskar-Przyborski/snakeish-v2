<script lang="ts">
	import { store } from './room_creation_state';
	import { createEventDispatcher, onDestroy } from 'svelte';
	const dispatch = createEventDispatcher();
	import ToggleInput from '$lib/components/inputs/toggle_input.svelte';
	import PinInput from '$lib/components/inputs/pin_input.svelte';
	import TextButton from '$lib/components/buttons/text_button.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import Icon from '@iconify/svelte';

	const onBack = () => dispatch('back');
	const onCreateRoom = () => dispatch('create-room');

	let btnDisabled = false;

	const updateBtnState = async (state: App.RoomCreationState) => {
		if (state.pinEnabled) {
			if (state.pin.some((s) => s == null)) btnDisabled = true;
			else btnDisabled = false;
			return;
		}

		btnDisabled = false;
	};

	const unsubscibe = store.subscribe(updateBtnState);
	onDestroy(unsubscibe);
</script>

<div class="choose-pin">
	<div class="use-pin">
		<h2>Use PIN code protection</h2>
		<ToggleInput bind:value={$store.pinEnabled} />
	</div>
	<PinInput
		pin={$store.pin}
		on:change={(e) => ($store.pin = e.detail.pin)}
		disabled={!$store.pinEnabled}
	>
		Enter the pin
	</PinInput>
</div>
<div class="continue-section">
	<TextButton on:click={onBack}>Back</TextButton>
	<Button on:click={onCreateRoom} disabled={btnDisabled}
		>Create Room
		<Icon icon="material-symbols:rocket-launch-rounded" inline />
	</Button>
</div>

<style lang="scss">
	.use-pin {
		margin: 1rem 0;

		display: flex;
		flex-flow: row nowrap;
		align-items: center;
		gap: 1rem;
		h2 {
			margin: 0;
		}
	}

	.continue-section {
		margin-top: 1.5rem;
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		align-items: center;
	}
</style>
