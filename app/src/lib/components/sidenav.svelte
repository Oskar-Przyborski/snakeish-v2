<script lang="ts">
	import { onMount } from 'svelte';
	import Icon from '@iconify/svelte';
	export let route: string;

	const elements = [
		{
			title: 'Home',
			href: '/',
			iconFull: 'material-symbols:home-rounded',
			iconEmpty: 'material-symbols:home-outline-rounded'
		},
		{
			title: 'Browse',
			href: '/browse',
			iconFull: 'material-symbols:search-rounded',
			iconEmpty: 'material-symbols:search-rounded'
		},
		{
			title: 'Create',
			href: '/create-room',
			iconFull: 'material-symbols:add-rounded',
			iconEmpty: 'material-symbols:add-rounded'
		},
		{
			title: 'Donate',
			href: '#',
			iconFull: 'material-symbols:coffee',
			iconEmpty: 'material-symbols:coffee-outline'
		}
	];

	let theme = 'dark';

	onMount(() => {
		theme = document.documentElement?.getAttribute('data-theme') ?? 'light';
	});

	const toggleTheme = () => {
		theme = theme == 'dark' ? 'light' : 'dark';
		document.documentElement?.setAttribute('data-theme', theme);
		localStorage.setItem('data-theme', theme);
	};
</script>

<nav class="sidenav">
	<div class="top">
		<a href="/">
			{#if theme == 'light'}
				<img class="logo" alt="Snakeish Logo" src="/logo-long-black.png" width="260" />
			{:else}
				<img class="logo" alt="Snakeish Logo" src="/logo-long-white.png" width="260" />
			{/if}
		</a>
		<ul style="margin: 1rem 0;">
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
					<Icon icon="material-symbols:sunny-rounded" inline />
					Light
				{:else}
					<Icon icon="material-symbols:nightlight-rounded" inline />
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
