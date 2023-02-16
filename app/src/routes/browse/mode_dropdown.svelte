<script lang="ts">
	import Panel from '$lib/components/panel.svelte';
	import Icon from '@iconify/svelte';
	let expanded = false;
</script>

<div class="mode-tag">
	<button class="name" on:click={() => (expanded = !expanded)}>
		Mode Tag
		<Icon icon={`mdi:arrow-${expanded ? 'up' : 'down'}-drop-circle-outline`} inline />
	</button>
	<div class="content" class:expanded>
		<div class="list">
			<div>
				<input type="checkbox" /> Casual
			</div>
			<div>
				<input type="checkbox" /> Huuge
			</div>
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
			padding: 1rem;
			margin-top: 0.4rem;

			.list {
				display: flex;
				flex-flow: column nowrap;
				align-items: flex-start;
			}
		}
	}
</style>
