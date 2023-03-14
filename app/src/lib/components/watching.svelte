<script lang="ts">
	import Icon from '@iconify/svelte';
	import Panel from './panel.svelte';

	export let count: number;
</script>

<Panel>
	<div class="watching">
		<h2>Watching</h2>
		<div class="count">
			<Icon icon="material-symbols:visibility-outline-rounded" inline />
			{count}
		</div>
	</div>
</Panel>

<style lang="scss">
	.watching {
		display: flex;
		flex-flow: row wrap;
		gap: 1rem;
		align-items: center;
		justify-content: space-between;
		h2 {
			margin: 0;
		}
		.count {
			font-size: 1.3rem;
		}
	}
</style>
