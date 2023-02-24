<script lang="ts">
	import Icon from '@iconify/svelte';

	export let step: { name: string; description: string; iconEmpty: string; iconFull: string };
	export let prevActive: boolean;
	export let active: boolean;
</script>

<div class="step-item" class:active class:prev-active={prevActive}>
	<div class="icon">
		<Icon icon={prevActive || active ? step.iconFull : step.iconEmpty} />
	</div>
	<div class="name">
		<div class="title">{step.name}</div>
		<div class="description">{step.description}</div>
	</div>
</div>

<style lang="scss">
	.step-item {
		padding: 1rem 0;
		display: flex;
		gap: 1rem;
		align-items: center;

		color: var(--text-light);

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
			color: white;
			font-size: 1.8rem;
			display: grid;
			place-items: center;
			background-color: var(--text-disabled);
			border-radius: 50%;
			padding: 0.6rem;
		}

		&.active {
			color: var(--text);
			.icon {
				color: white;
				background-color: var(--red);
			}
		}

		&.prev-active {
			color: var(--text);
			.icon {
				color: white;
				background-color: #aaa;
			}
		}
	}
</style>
