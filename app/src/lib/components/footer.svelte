<script>
	import Icon from '@iconify/svelte';
</script>

<footer>
	<b>Snakeish</b> by<br />
	<a href="https://github.com/Oskar-Przyborski" target="_blank" rel="noreferrer noopener nofollow">
		Oskar Przyborski <Icon icon="mdi:github" inline />
	</a>
</footer>

<style lang="scss">
	footer {
		margin-top: 1.2rem;
		padding-bottom: 0.3rem;
		text-align: center;
		a {
			color: inherit;
			text-decoration: inherit;
			&:hover {
				text-decoration: underline;
			}
		}
	}
</style>
