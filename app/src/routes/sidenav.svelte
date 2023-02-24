<script lang="ts">
	import { onMount } from 'svelte';
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

	let theme = 'dark';

	onMount(() => {
		const systemDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
		theme = localStorage.getItem('theme') ?? systemDark ? 'dark' : 'light'
		document.documentElement?.setAttribute('data-theme', theme);
	});

	const toggleTheme = () => {
		const newTheme = theme == 'dark' ? 'light' : 'dark';

		theme = newTheme;
		document.documentElement?.setAttribute('data-theme', newTheme);
		localStorage.setItem('theme', newTheme);
	};
</script>

<nav class="sidenav">
	<div class="top">
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
	</div>
	<div class="bottom">
		<ul>
			<li on:click={toggleTheme} on:keypress={toggleTheme}>
				{#if theme == 'light'}
					<Icon icon="mdi:weather-sunny" inline />
					Light
				{:else}
					<Icon icon="mdi:moon-waning-crescent" inline />
					Dark
				{/if}
			</li>
		</ul>
	</div>
</nav>

<style lang="scss">
	.sidenav {
		height: 100%;
		z-index: 1;
		top: 0;
		left: 0;
		overflow-x: hidden;

		font-size: 1.3rem;
		padding: 0.8rem 1rem;
		border-right: 2px solid var(--outline);
		background-color: var(--surface);
		display: flex;
		flex-flow: column nowrap;
		justify-content: space-between;

		.logo {
			height: auto;
			width: 100%;
		}
		ul {
			margin: 0;
			padding: 0;
			font-weight: 400;
			li {
				list-style-type: none;
				cursor: pointer;
				margin: 0.5rem 0;
				padding: 0.6rem;
				border-radius: 0.3rem;
				color: var(--text-light);

				&:hover,
				&.active {
					color: var(--text-light);
					background-color: var(--outline);
				}
			}
			a {
				color: inherit;
				text-decoration: none;
			}
		}
	}
</style>
