<script lang="ts">
	import { createEventDispatcher, onDestroy } from 'svelte';
	const dispatch = createEventDispatcher();
	import { store } from '$lib/room_creation_state';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import { fetchJson } from '$lib/fetchJson';

	const onContinue = () => dispatch('continue');

	let nameError = '';
	let btnDisabled = false;

	const updateBtnState = async (state: App.RoomCreationState) => {
		if (state.roomName == '') {
			nameError = '';
			btnDisabled = true;
			return;
		}

		const { available } = await fetchJson<{ available: boolean }>(
			`/rooms/name/${state.roomName}`
		);

		if (available) {
			nameError = '';
			btnDisabled = false;
		} else {
			nameError = 'This name is already used.';
			btnDisabled = true;
		}
	};

	const unsubscibe = store.subscribe(updateBtnState);
	onDestroy(unsubscibe);
</script>

<Panel>
	<div class="input-wrapper">
		<TextInput bind:value={$store.roomName} error={nameError != ''} altText={nameError}>
			Room Name
		</TextInput>
	</div>
</Panel>

<div class="continue-section">
	<Button on:click={onContinue} disabled={btnDisabled}>Continue</Button>
</div>

<style lang="scss">
	.input-wrapper {
		font-size: 1.2rem;
	}
	.continue-section {
		margin-top: 1.5rem;
		display: flex;
		justify-content: flex-end;
	}
</style>
