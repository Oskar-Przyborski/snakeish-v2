<script lang="ts">
	import Icon from '@iconify/svelte';

	export let value = '';
	export let placeholder = '';
	export let width: any = null;
	export let fontSize: any = null;
	export let altText = '';
	export let error = false;

	let focused = false;
</script>

<div class="text-input" class:error style:width style:font-size={fontSize}>
	<div class="name">
		<slot />
	</div>
	<input
		type="text"
		{placeholder}
		bind:value
		on:focus={() => (focused = true)}
		on:blur={() => (focused = false)}
		on:change
		on:focus
		on:blur
	/>
	{#if altText}
		<div class="alt-text"><Icon icon="mdi:information" inline /> {altText}</div>
	{/if}
</div>

<style lang="scss">
	.text-input {
		margin: 1rem 0;
		.name {
			color: #222;
			font-size: 0.92rem;
			margin-bottom: 0.4rem;
		}
		input {
			all: unset;
			width: calc(100% - 2rem);
			font-size: 1rem;
			padding: 0.5rem 1rem;
			text-align: left;
			border: 2px solid #eee;
			border-radius: 0.4rem;
			background-color: white;
		}
		.alt-text {
			margin-top: 0.3rem;
			font-size: 0.9rem;
			color: #fffa;
		}
		&.error {
			.name,
			.alt-text {
				color: #f43737;
			}
			input {
				border-color: #f43737;
			}
		}
	}
</style>
