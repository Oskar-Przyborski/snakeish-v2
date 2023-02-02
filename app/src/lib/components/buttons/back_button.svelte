<script>
	import Icon from '@iconify/svelte';
</script>

<button on:click>
	<Icon icon="mdi:chevron-left" inline />
</button>

<style lang="scss">
	button {
		all: unset;
		cursor: pointer;
		font-size: 2rem;
		
		&:focus-visible{
			outline: auto;
		}
	}
</style>
