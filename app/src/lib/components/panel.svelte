<script lang="ts">
	export let full: boolean = false;
</script>

<div class="panel" class:full>
	<slot />
</div>

<style lang="scss">
	.panel {
		background-color: #ffffff1a;
		border-radius: 10px;
		padding: 1.3rem;

		&.full {
			width: 100%;
			height: 100%;
		}
	}
</style>
