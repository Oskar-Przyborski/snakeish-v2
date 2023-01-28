<script lang="ts">
	import Icon from '@iconify/svelte';

	export let callback: () => void;
	let icon: HTMLDivElement;

	function rotate() {
		callback();

		const keyframes: Keyframe[] = [
			{
				transform: 'scale(0.95) rotate(0deg)',
				offset: 0
			},
			{
				transform: 'scale(1) rotate(18deg)',
				offset: 0.05
			},
			{
				transform: 'scale(1) rotate(360deg)',
				offset: 1
			}
		];
        
		icon.animate(keyframes, { duration: 500, easing: 'ease-in-out' });
	}
</script>

<div class="refresh-button" bind:this={icon} on:click={rotate} on:keypress={rotate}>
	<Icon icon="material-symbols:refresh-rounded" />
</div>

<style lang="scss">
	.refresh-button {
		display: grid;
		place-items: center;
		font-size: 2rem;
		&:hover {
			cursor: pointer;
		}
	}
</style>
