<script lang="ts">
	import ToggleInput from '$lib/components/inputs/toggle_input.svelte';
	import PinInput from '$lib/components/inputs/pin_input.svelte';
	import { state } from '$lib/room_creation_state';
	import Panel from '$lib/components/panel.svelte';
</script>

<div class="choose-pin">
	<div class="use-pin">
		<h2>Use PIN code protection</h2>
		<ToggleInput bind:value={$state.pinEnabled} />
	</div>
	<Panel>
		<div class="pin-wrapper" class:disabled={!$state.pinEnabled}>
			Enter the pin
			<PinInput pin={$state.pin} disabled={!$state.pinEnabled} />
		</div>
	</Panel>
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

	.pin-wrapper.disabled {
		color: #fff8;
	}
</style>
