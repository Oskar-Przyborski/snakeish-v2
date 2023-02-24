<script lang="ts">
	export let vertical = false;
</script>

<div class="divider" class:vertical />

<style lang="scss">
	.divider {
		height: 0px;
		border-bottom: 2px solid var(--outline);
		margin: 1rem;

		&.vertical {
			height: initial;
			border-bottom: none;

			width: 0px;
			border-right: 2px solid var(--outline);
		}
	}
</style>
