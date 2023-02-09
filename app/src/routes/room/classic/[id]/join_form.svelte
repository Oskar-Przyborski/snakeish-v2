<script lang="ts">
	import Button from '$lib/components/buttons/button.svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import { requestJoin } from './classic_game';
	import PinInput from '$lib/components/inputs/pin_input.svelte';

	export let pinEnabled: boolean = false;
	let name: string;
	let color: string;
	let pin: (number | null)[] = []

	const submit = () => {
		if (!canSubmit()) return;
		requestJoin(name, color, pin);
	};
	const canSubmit = () => !(pinEnabled && pin.some((n) => n == null));
</script>

<div class="join-form">
	<h1>Join Game</h1>
	<div class="form">
		<TextInput bind:value={name}>Name</TextInput>
		<TextInput bind:value={color}>Color</TextInput>
		{#if pinEnabled}
			<div>
				Room's PIN Code
				<PinInput bind:pin />
			</div>
		{/if}
	</div>
	<div class="join-btn">
		<Button on:click={submit} disabled={!canSubmit()}>Join</Button>
	</div>
</div>

<style lang="scss">
	.join-form {
		h1 {
			margin: 0;
		}
		.form {
			margin: 2rem 0;
		}
		.join-btn {
			float: right;
		}
	}
</style>
