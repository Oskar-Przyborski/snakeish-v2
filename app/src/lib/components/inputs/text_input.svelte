<script lang="ts">
	import Icon from '@iconify/svelte';

	export let value = '';
	export let placeholder = '';
	export let width: any = null;
	export let fontSize: any = null;
	export let altText: string = '';
	export let error: boolean = false;

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
	<div class="border" class:focused />
	{#if altText}
		<div class="alt-text"><Icon icon="mdi:information" inline /> {altText}</div>
	{/if}
</div>

<style lang="scss">
	.text-input {
		margin: 1rem 0;
		.name {
			color: #fffa;
		}
		input {
			all: unset;
			width: 100%;
			font-size: inherit;
			padding: 0.5rem 0.1rem;
			text-align: left;
		}
		.border {
			border-bottom: 3px solid white;
			transition: border-bottom 0.1s, box-shadow 0.1s;
			&.focused {
				border-bottom: 3px solid #d62246;
				box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.384);
			}
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
			.border {
				border-bottom: 3px solid #f43737;
			}
		}
	}
</style>
