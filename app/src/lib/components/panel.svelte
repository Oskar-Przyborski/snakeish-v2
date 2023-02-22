<script lang="ts">
	export let full = false;
	export let padding: string = '1.3rem';
	export let margin = false;
</script>

<div class="panel" class:full style:padding class:margin>
	<slot />
</div>

<style lang="scss">
	.panel {
		border: 2px solid #eee;
		background-color: #fff;
		border-radius: 10px;
		&.margin {
			margin: 1rem;
		}
		

		&.full {
			width: 100%;
			height: 100%;
		}
	}
</style>
