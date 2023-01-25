<script>
	import Icon from '@iconify/svelte';
</script>

<footer>
	<b>Snakeish</b> by<br>
	<a href="https://github.com/Oskar-Przyborski">
		Oskar Przyborski <Icon icon="mdi:github" inline={true} />
	</a>
</footer>

<style lang="scss">
	footer {
		margin-top: 1.5em;
		text-align: center;
		a {
			color: inherit;
			text-decoration: inherit;
            &:hover{
                text-decoration: underline;
            }
		}
	}
</style>
