<script lang="ts">
	import Icon from '@iconify/svelte';

	export let colors: string[];
	export let selected: number = 0;
	const nextColor = () => (selected = (selected + 1) % colors.length);
	const prevColor = () => (selected = selected == 0 ? colors.length - 1 : selected - 1);
</script>

<div class="skin-selector">
	<span class="name"><slot /></span>
	<div class="row">
		<button on:click={prevColor} class="arrow-btn">
			<Icon inline icon="mdi:chevron-left" />
		</button>
		<div class="color-indicator" style={`background-color: ${colors[selected]};`} />
		<button on:click={nextColor} class="arrow-btn">
			<Icon inline icon="mdi:chevron-right" />
		</button>
	</div>
</div>

<style lang="scss">
	.name {
		color: #222;
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
		color: #000;
		background-color: transparent;
		border: none;
		&:hover {
			color: #222;
		}
	}
</style>
