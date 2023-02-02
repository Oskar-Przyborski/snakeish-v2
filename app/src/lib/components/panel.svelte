<script lang="ts">
	export let full = false;
	export let padding: string = "1.3rem"
</script>

<div class="panel" class:full style:padding>
	<slot />
</div>

<style lang="scss">
	.panel {
		background-color: #ffffff1a;
		border-radius: 10px;

		&.full {
			width: 100%;
			height: 100%;
		}
	}
</style>
