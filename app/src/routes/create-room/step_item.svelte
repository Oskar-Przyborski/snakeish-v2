<script lang="ts">
	import Icon from '@iconify/svelte';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	const onClick = () => dispatch('click');

	export let step: { name: string; description: string; iconEmpty: string; iconFull: string };
	export let prevActive: boolean;
	export let active: boolean;
</script>

<svelte:element
	this={prevActive ? 'button' : 'div'}
	class="step-item"
	class:active
	class:prev-active={prevActive}
	on:click={onClick}
	on:keypress={onClick}
>
	<div class="icon">
		<Icon icon={prevActive || active ? step.iconFull : step.iconEmpty} />
	</div>
	<div class="name">
		<div class="title">{step.name}</div>
		<div class="description">{step.description}</div>
	</div>
</svelte:element>

<style lang="scss">
	.step-item {
		all: unset;
		margin: 2rem 0;

		display: flex;
		gap: 1rem;
		align-items: center;

		color: #fff7;

		.name {
			.title {
				font-weight: bold;
				font-size: 1.3rem;
			}
			.description {
				font-size: 0.9rem;
			}
		}
		.icon {
			color: #fff7;
			font-size: 1.8rem;
			display: grid;
			place-items: center;
			background-color: #fff2;
			border-radius: 50%;
			padding: 0.6rem;
		}

		&.active {
			color: #fff;
			.icon {
				color: #fff;
				background-color: var(--red);
			}
		}

		&.prev-active {
			cursor: pointer;
			color: #fff;
			.icon {
				color: #fff;
				background-color: #fff7;
			}

			&:hover {
				color: #fffc;
			}
			&:focus-visible {
				outline-color: rgb(255, 255, 255);
				outline-width: 0.8px;
				outline-style: auto;
				outline-offset: 0px;
			}
		}
	}
</style>
