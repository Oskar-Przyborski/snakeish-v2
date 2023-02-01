<script lang="ts">
	export let value: boolean;
</script>

<button class="switch" class:checked={value} on:click={() => (value = !value)}>
	<div class="slider" />
</button>

<style lang="scss">
	.switch {
		all: unset;
		margin: 0 0.25rem;
		cursor: pointer;
		vertical-align: baseline;
		width: 2em;
		border-radius: 2em;
		padding: 0.25em;
		background-color: #fff8;
		transition: background-color 0.1s ease-in-out;

		&:hover {
			background-color: #fff6;
		}

		.slider {
			width: 1em;
			height: 1em;
			background-color: white;
			border-radius: 50%;
			transition: margin-left 0.1s ease-in-out;
		}

		&.checked {
			background-color: var(--red);
			&:hover {
				background-color: var(--hover-red);
			}
			.slider {
				margin-left: 1em;
			}
		}
	}
</style>
