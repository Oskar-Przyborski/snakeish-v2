<script lang="ts">
	import { browser } from '$app/environment';
	import { onDestroy, onMount } from 'svelte';
	import Icon from '@iconify/svelte';
	import CheckboxInput from '$lib/components/inputs/checkbox_input.svelte';

	let expanded = false;
	let element: HTMLDivElement;

	function onClickOutside(e: MouseEvent) {
		if (e.target != null && !element.contains(e.target as Node)) {
			expanded = false;
		}
	}

	onMount(() => {
		if (browser) window.addEventListener('click', onClickOutside);
	});

	onDestroy(() => {
		if (browser) window.removeEventListener('click', onClickOutside);
	});
</script>

<div class="mode-tag" bind:this={element}>
	<button class="name" on:click={() => (expanded = !expanded)}>
		Mode Tag
		<Icon icon={`mdi:arrow-${expanded ? 'up' : 'down'}-drop-circle-outline`} inline />
	</button>
	<div class="content" class:expanded>
		<div class="list">
			<CheckboxInput>Casual</CheckboxInput>
			<CheckboxInput>Huuge</CheckboxInput>
		</div>
	</div>
</div>

<style lang="scss">
	.mode-tag {
		position: relative;

		.name {
			all: unset;
			cursor: pointer;
			user-select: none;
			font-size: 1.3rem;
			padding: 0.8rem 1rem;
			background-color: #fff1;
			border-radius: 0.5rem;
		}

		.content {
			display: none;
			&.expanded {
				display: block;
			}

			position: absolute;
			right: 0;

			background: #474a4e;
			border-radius: 0.5rem;
			padding: 0.9rem;
			margin-top: 0.4rem;

			.list {
				display: flex;
				flex-flow: column nowrap;
				align-items: flex-start;
				gap: 0.5rem;
			}
		}
	}
</style>
