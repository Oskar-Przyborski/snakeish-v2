<script lang="ts">
	import Button from '$lib/components/buttons/button.svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import { requestJoin } from './classic_game';
	import PinInput from '$lib/components/inputs/pin_input.svelte';
	import SkinSelector from '$lib/components/inputs/color_input.svelte';

	export let pinEnabled: boolean = false;
	let name: string;
	let selectedColor: number;
	let pin: (number | null)[] = [];

	const skinColors = [
		'#deb135',
		'#80e356',
		'#e37e56',
		'#56e37e',
		'#56dee3',
		'#5672e3',
		'#8a56e3',
		'#e356bd',
		'#e3566b'
	];

	const submit = () => {
		if (!canSubmit()) return;
		requestJoin(name, skinColors[selectedColor], pin);
	};
	const canSubmit = () => !(pinEnabled && pin.some((n) => n == null));
</script>

<div class="join-form">
	<h1>Join Game</h1>
	<div class="form">
		<TextInput bind:value={name}>Name</TextInput>
		<SkinSelector colors={skinColors} bind:selected={selectedColor}>Skin color</SkinSelector>
		{#if pinEnabled}
			<PinInput bind:pin>Room's PIN Code</PinInput>
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
