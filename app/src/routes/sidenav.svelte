<script lang="ts">
	import Icon from '@iconify/svelte';
	export let route: string;

	const elements = [
		{
			title: 'Home',
			href: '/',
			iconFull: 'mdi:home',
			iconEmpty: 'mdi:home-outline'
		},
		{
			title: 'Browse',
			href: '/browse',
			iconFull: 'mdi:magnify',
			iconEmpty: 'mdi:magnify-outline'
		},
		{
			title: 'Create',
			href: '/create-room',
			iconFull: 'mdi:plus-bold',
			iconEmpty: 'mdi:plus-outline'
		},
		{
			title: 'Donate',
			href: '#',
			iconFull: 'mdi:coffee',
			iconEmpty: 'mdi:coffee-outline'
		}
	];
</script>

<nav class="sidenav">
	<a href="/">
		<img class="logo" alt="Snakeish Logo" src="/logo-long-white.png" width="260" />
	</a>
	<ul>
		{#each elements as el}
			<a href={el.href}>
				<li class:active={route === el.href}>
					<Icon icon={route === el.href ? el.iconFull : el.iconEmpty} inline />
					{el.title}
				</li>
			</a>
		{/each}
	</ul>
</nav>

<style lang="scss">
	.sidenav {
		height: 100%;
		z-index: 1;
		top: 0;
		left: 0;
		overflow-x: hidden;

		padding: 0.8rem 1rem;
		border-right: 2px solid #eee;
		background-color: #fff;
		font-size: 1.3rem;

		.logo {
			height: auto;
			width: 100%;
		}
		ul {
			padding: 0;
			font-weight: 400;
			li {
				list-style-type: none;
				cursor: pointer;
				margin: 0.5rem 0;
				padding: 0.6rem;
				border-radius: 0.3rem;
				color: #555;

				&:hover,
				&.active {
					color: #333;
					background-color: #eee;
				}
			}
			a {
				color: inherit;
				text-decoration: none;
			}
		}
	}
</style>
