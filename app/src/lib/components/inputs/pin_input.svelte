<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	import Icon from '@iconify/svelte';

	export let disabled = false;
	export let altText: string = '';
	export let error: boolean = false;
	export let pin: (number | null)[] = [null, null, null, null];
	let char1: HTMLInputElement;
	let char2: HTMLInputElement;
	let char3: HTMLInputElement;
	let char4: HTMLInputElement;

	let pinHidden = true;

	const focusPrevElement = (currEl: HTMLInputElement) => {
		if (currEl == char1) char1.blur();
		else if (currEl == char2) char1.select();
		else if (currEl == char3) char2.select();
		else if (currEl == char4) char3.select();
	};
	const focusNextElement = (currEl: HTMLInputElement) => {
		if (currEl == char1) char2.select();
		else if (currEl == char2) char3.select();
		else if (currEl == char3) char4.select();
		else if (currEl == char4) char4.blur();
	};

	const registerCharFromEl = (el: HTMLInputElement) => {
		const value = el.value == '' ? null : parseInt(el.value);
		console.log(value);
		if (el == char1) pin[0] = value;
		else if (el == char2) pin[1] = value;
		else if (el == char3) pin[2] = value;
		else if (el == char4) pin[3] = value;
		dispatch('change', { pin });
	};

	const onKeypress = (e: any) => {
		e.preventDefault();
		const isTabForward = e.key == 'Tab' && !e.shiftKey;
		const isTabBackward = e.key == 'Tab' && e.shiftKey;
		const isBackspace = e.key == 'Backspace';
		const isNumber = !isNaN(e.key);
		const charEl = e.target;

		if (isTabForward) focusNextElement(charEl);
		else if (isTabBackward) focusPrevElement(charEl);
		else if (!isBackspace && !isNumber) return;

		if (isBackspace) {
			if (charEl.value == '') {
				focusPrevElement(charEl);
			} else {
				charEl.value = '';
			}
		} else if (isNumber) {
			charEl.value = e.key;
			focusNextElement(charEl);
		}
		registerCharFromEl(charEl);
	};
</script>

<div class="pin-input" class:error class:disabled>
	<div class="name">
		<slot />
	</div>
	<div class="input-wrapper">
		<input
			class="pass-char-input"
			type={pinHidden ? 'password' : 'text'}
			maxlength="1"
			{disabled}
			bind:this={char1}
			on:keydown={onKeypress}
			value={pin[0] == null ? '' : pin[0]}
		/>
		<input
			class="pass-char-input"
			type={pinHidden ? 'password' : 'text'}
			maxlength="1"
			{disabled}
			bind:this={char2}
			on:keydown={onKeypress}
			value={pin[1] == null ? '' : pin[1]}
		/>
		<input
			class="pass-char-input"
			type={pinHidden ? 'password' : 'text'}
			maxlength="1"
			{disabled}
			bind:this={char3}
			on:keydown={onKeypress}
			value={pin[2] == null ? '' : pin[2]}
		/>
		<input
			class="pass-char-input"
			type={pinHidden ? 'password' : 'text'}
			maxlength="1"
			{disabled}
			bind:this={char4}
			on:keydown={onKeypress}
			value={pin[3] == null ? '' : pin[3]}
		/>

		<button class="show-hide-eye" on:click={() => (pinHidden = !pinHidden)} {disabled}>
			<Icon icon={pinHidden ? 'mdi:eye' : 'mdi:eye-off'} inline />
		</button>
	</div>
	{#if altText}
		<div class="alt-text"><Icon icon="mdi:information" inline /> {altText}</div>
	{/if}
</div>

<style lang="scss">
	.pin-input {
		margin: 1rem 0;
		.name {
			color: #fffa;
		}
		.input-wrapper {
			padding-top: 0.5rem;
			display: flex;
			flex-flow: row nowrap;
			gap: 0.5rem;

			.pass-char-input {
				all: unset;
				padding-bottom: 0.5rem;
				border-bottom: 3px solid white;
				width: 1em;
				font-size: 1.2rem;
				text-align: center;
				&:focus {
					border-bottom: 3px solid var(--red);
				}
			}
			.show-hide-eye {
				all: unset;
				font-size: 1.5rem;
				margin-left: 1rem;
				cursor: pointer;

				&:focus-visible {
					outline: auto;
				}
			}
		}
		.alt-text {
			margin-top: 0.5rem;
			font-size: 0.9rem;
			color: #fffa;
		}
		&.error {
			.name,
			.alt-text {
				color: #f43737;
			}
			.pass-char-input {
				border-bottom-color: #f43737;
			}
		}

		&.disabled {
			.name{
				color: #fff8;
			}
			.alt-text {
				color: #fff8;
			}
			.pass-char-input {
				color: #fff8;
				border-bottom-color: #fff8;
			}
			.show-hide-eye{
				color: #fff8;
			}
		}
	}
</style>
