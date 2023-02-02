<button on:click>
	<slot />
</button>

<style lang="scss">
	button {
		cursor: pointer;
		background: none;
		border: none;

		color: inherit;
		font-family: inherit;
		font-size: inherit;

		&:hover {
			text-decoration: underline;
		}
	}
</style>
