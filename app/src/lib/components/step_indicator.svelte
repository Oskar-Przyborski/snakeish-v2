<script lang="ts">
	export let current: number;
	export let steps: number;
</script>

<div class="steps">
	{#each Array(steps) as _, idx}
		{#if idx == current}
			<svg viewBox={`0 0 32 20`} width="30" height="20">
				<line
					x1="6"
					y1="10"
					x2="26"
					y2="10"
					stroke="#d62246"
					stroke-width="10"
					stroke-linecap="round"
				/>
			</svg>
		{:else}
			<svg viewBox={`0 0 12 20`} width="20" height="20">
				<circle cx="6" cy="10" r="5" fill="#fff5" />
			</svg>
		{/if}
	{/each}
</div>

<style lang="scss">
	.steps {
		display: flex;
		justify-content: center;
	}
</style>
