<script lang="ts">
	import Button from '$lib/components/buttons/button.svelte';
	import TextInput from '$lib/components/inputs/text_input.svelte';
	import { requestJoin, store } from './battleroyale';
	import PinInput from '$lib/components/inputs/pin_input.svelte';
	import SkinSelector from '$lib/components/inputs/color_input.svelte';
	import { onMount } from 'svelte';

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

	const onNameInputBlur = () => {
		if (name.length < 3 || name.length > 10) {
			$store.errors.name = 'Name should be 3-10 chars long!';
		} else {
			$store.errors.name = undefined;
		}
	};

	onMount(() => {
		name = localStorage.getItem('player-name') ?? '';
		const color = localStorage.getItem('player-color');
		if (color) {
			selectedColor = skinColors.indexOf(color);
		}
	});
</script>

<div class="join-form">
	<h1>Join Game</h1>
	<div class="form">
		<TextInput
			bind:value={name}
			error={$store.errors.name != undefined}
			altText={$store.errors.name}
			on:blur={onNameInputBlur}
		>
			Name
		</TextInput>
		<SkinSelector colors={skinColors} bind:selected={selectedColor}>Skin color</SkinSelector>
		{#if pinEnabled}
			<PinInput
				bind:pin
				error={$store.errors.pin != undefined}
				altText={$store.errors.pin}
				on:blur={($store.errors.pin = undefined)}
			>
				Room's PIN Code
			</PinInput>
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
