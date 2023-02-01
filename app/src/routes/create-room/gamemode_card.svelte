<script lang="ts">
	import modes from '$lib/modes';
	import Icon from '@iconify/svelte';
	import { createEventDispatcher } from 'svelte';
	const dispatcher = createEventDispatcher();

	export let configName: string;
	export let selected: boolean = false;

	$: mode = modes.get(configName)!;
</script>

<button class="gamemode-card" class:selected on:click={() => dispatcher('select')}>
	<div class="top">
		<div class="name">
			<h2>{mode.title}</h2>
			<div class="tag">{mode.tag.toUpperCase()}</div>
		</div>
		{#if selected}
			<div class="checkbox">
				<Icon icon="mdi:check-bold" inline />
			</div>
		{/if}
	</div>
	<div class="description">{mode.description}</div>
</button>

<style lang="scss">
	.gamemode-card {
		all: unset;
		border: 3px solid #ffffff;
		border-radius: 0.8rem;
		padding: 1rem;
		transition: border 0.1s ease-in-out, box-shadow 0.1s ease-in-out;
		cursor: pointer;
		.top {
			display: flex;
			flex-flow: row nowrap;
			gap: 1rem;
			justify-content: space-between;
			align-items: flex-start;
			.name {
				display: flex;
				flex-flow: row nowrap;
				gap: 0.5rem;
				align-items: center;
				.tag {
					font-size: 0.75rem;
					font-weight: bold;
					background-color: #0000002b;
					width: max-content;
					padding: 0.3rem;
				}
				h2 {
					margin: 0.3rem 0;
				}
			}

			.checkbox {
				font-size: 1.2rem;
				background-color: var(--red);
				width: 1.6rem;
				height: 1.6rem;
				display: grid;
				place-items: center;
				border-radius: 0.3rem;
			}
		}

		&:hover {
			box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.2);
		}
		&:focus-visible {
			outline-color: rgb(255, 255, 255);
			outline-width: 0.8px;
			outline-style: auto;
			outline-offset: 0px;
		}

		&.selected {
			border: 3px solid #d62246;
			// box-shadow: 0px 3px 10px rgba(0, 0, 0, 0.384);
		}
	}
</style>
