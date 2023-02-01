<script lang="ts">
	import Icon from '@iconify/svelte';

	export let disabled: boolean = false;
	export let pin = ['', '', '', ''];
	let char1: HTMLInputElement;
	let char2: HTMLInputElement;
	let char3: HTMLInputElement;
	let char4: HTMLInputElement;

	let pinHidden: boolean = true;

	const focusPrevElement = (currEl: any) => {
		if (currEl == char1) char1.blur();
		else if (currEl == char2) char1.select();
		else if (currEl == char3) char2.select();
		else if (currEl == char4) char3.select();
	};
	const focusNextElement = (currEl: any) => {
		if (currEl == char1) char2.select();
		else if (currEl == char2) char3.select();
		else if (currEl == char3) char4.select();
		else if (currEl == char4) char4.blur();
	};

	const registerCharFromEl = (el: any) => {
		if (el == char1) pin[0] = el.value;
		else if (el == char2) pin[1] = el.value;
		else if (el == char3) pin[2] = el.value;
		else if (el == char4) pin[3] = el.value;
	};

	const onKeypress = (e: any) => {
		e.preventDefault();
		const isBackspace = e.key == 'Backspace';
		const isNumber = !isNaN(e.key);
		const charEl = e.target;

		if (!isBackspace && !isNumber) return;

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

<div class="input-wrapper" class:disabled>
	<input
		class="pass-char-input"
		type={pinHidden ? 'password' : 'text'}
		maxlength="1"
		{disabled}
		bind:this={char1}
		on:keydown={onKeypress}
		value={pin[0]}
	/>
	<input
		class="pass-char-input"
		type={pinHidden ? 'password' : 'text'}
		maxlength="1"
		{disabled}
		bind:this={char2}
		on:keydown={onKeypress}
		value={pin[1]}
	/>
	<input
		class="pass-char-input"
		type={pinHidden ? 'password' : 'text'}
		maxlength="1"
		{disabled}
		bind:this={char3}
		on:keydown={onKeypress}
		value={pin[2]}
	/>
	<input
		class="pass-char-input"
		type={pinHidden ? 'password' : 'text'}
		maxlength="1"
		{disabled}
		bind:this={char4}
		on:keydown={onKeypress}
		value={pin[3]}
	/>

	<button class="show-hide-eye" on:click={() => (pinHidden = !pinHidden)} {disabled}>
		<Icon icon={pinHidden ? 'mdi:eye' : 'mdi-eye-off'} inline />
	</button>
</div>

<style lang="scss">
	.input-wrapper {
		padding-top: 0.5rem;
		display: flex;
		flex-flow: row nowrap;
		gap: 0.5rem;

		&.disabled {
			color: #fff8;
			.pass-char-input {
				border-bottom-color: #fff8;
			}
		}

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
		}
	}
</style>
