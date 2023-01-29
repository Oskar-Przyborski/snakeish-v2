<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import Button from '$lib/components/button.svelte';
	import TextInput from '$lib/components/text_input.svelte';
	const dispatch = createEventDispatcher()

	export let value: string;
	
	const onContinue = () => {
		dispatch("continue")
	}
</script>

<div class="wrapper">
	<div class="choose-name">
		<h1>What would you like to name your room?</h1>
		<div class="input-wrapper"><TextInput placeholder="Room Name" bind:value /></div>
		<Button on:click={onContinue}>Continue</Button>
	</div>
</div>

<style lang="scss">
	.wrapper {
		height: 100%;
		display: grid;
		place-items: center;
	}

	.choose-name {
		max-width: 500px;
		text-align: center;
		.input-wrapper {
			padding: 0 2rem;
			font-size: 1.4rem;
		}
	}
</style>
