<script lang="ts">
	import Icon from '@iconify/svelte';

	export let colors: string[];
	export let selected = 0;
	const nextColor = () => (selected = (selected + 1) % colors.length);
	const prevColor = () => (selected = selected == 0 ? colors.length - 1 : selected - 1);
</script>

<div class="skin-selector">
	<span class="name"><slot /></span>
	<div class="row">
		<button on:click={prevColor} class="arrow-btn">
			<Icon inline icon="material-symbols:chevron-left-rounded" />
		</button>
		<div class="color-indicator" style={`background-color: ${colors[selected]};`} />
		<button on:click={nextColor} class="arrow-btn">
			<Icon inline icon="material-symbols:chevron-right-rounded" />
		</button>
	</div>
</div>

<style lang="scss">
	.name {
		color: var(--text-light);
	}
	.row {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin: 0.5rem 0;
	}
	.color-indicator {
		width: 2rem;
		height: 2rem;
		border-radius: 0.4rem;
	}
	.arrow-btn {
		cursor: pointer;
		font-size: 2rem;
		color: var(--text);
		background-color: transparent;
		border: none;
		&:hover {
			color: var(--text-light);
		}
	}
</style>
