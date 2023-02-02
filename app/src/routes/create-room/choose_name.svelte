<script lang="ts">
	import { createEventDispatcher, onDestroy } from 'svelte';
	const dispatch = createEventDispatcher();
	import { store } from '$lib/room_creation_state';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import Panel from '$lib/components/panel.svelte';
	import Button from '$lib/components/buttons/button.svelte';
	import backendRequest from '$lib/backend_request';

	const onContinue = () => dispatch('continue');

	let btnDisabled = false;

	const updateBtnState = async (state: App.RoomCreationState) => {
		if (state.roomName == '') {
			btnDisabled = true;
			return;
		}

		const { data } = await backendRequest<{ available: boolean }>(
			fetch,
			`/room-name-available?name=${state.roomName}`
		);

		if (data?.available) {
			btnDisabled = false;
			return;
		}
		btnDisabled = true;
	};

	const unsubscibe = store.subscribe(updateBtnState);
	onDestroy(unsubscibe);
</script>

<Panel>
	<div class="input-wrapper">
		<TextInput placeholder="Room name" bind:value={$store.roomName} />
	</div>
</Panel>

<div class="continue-section">
	<Button on:click={onContinue} disabled={btnDisabled}>
		Continue
	</Button>
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
