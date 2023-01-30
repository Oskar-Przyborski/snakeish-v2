<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/button.svelte';
	import PinInput from '$lib/components/pin_input.svelte';
	const dispatch = createEventDispatcher();

	export let pin: string[];

	const onContinue = () => {
		dispatch('continue');
		console.log(pin);
	};
</script>

<div class="wrapper">
	<div class="choose-password">
		<h1>Do you want to add PIN?</h1>
		<div class="under-header">(Optional)</div>
		<PinInput {pin} />
		<Button on:click={onContinue}>
			{#if pin.every((s) => (s = ''))}
				Skip
			{:else}
				Continue
			{/if}
		</Button>
	</div>
</div>

<style lang="scss">
	.wrapper {
		height: 100%;
		display: grid;
		place-items: center;
	}

	.choose-password {
		h1 {
			margin-bottom: 0.2rem;
		}
		max-width: 500px;
		text-align: center;
	}
</style>
