<script lang="ts">
	import PinInput from '$lib/components/inputs/pin_input.svelte';

	export let pin: string[];
</script>

<div class="choose-pin">
	<div>Enter the pin</div>
	<PinInput {pin} />
</div>

<style lang="scss">
	.choose-pin {
		// max-width: 500px;
		text-align: center;
	}
</style>
