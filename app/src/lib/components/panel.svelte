<script lang="ts">
	export let full = false;
	export let padding = '1.3rem';
	export let margin = false;
</script>

<div class="panel" class:full style:padding class:margin>
	<slot />
</div>

<style lang="scss">
	.panel {
		border: 2px solid var(--outline);
		background-color: var(--surface);
		border-radius: 0.4rem;
		&.margin {
			margin: 1rem;
		}

		&.full {
			width: 100%;
			height: 100%;
		}
	}
</style>
