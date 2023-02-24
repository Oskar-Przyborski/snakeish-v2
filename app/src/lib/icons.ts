import { addCollection } from '@iconify/svelte';

const loadIcons = async (prefix: string, icons: string[]) => {
	const resp = await fetch(
		`https://api.iconify.design/${prefix.trim()}.json?icons=${icons.join(',')}`
	);
	if (resp.status == 200) {
		addCollection(await resp.json());
		console.log(`Downloaded ${icons.length} icons with prefix: ${prefix}`);
	} else {
		console.error(`Error while downloading icons with prefix: ${prefix}`);
	}
};

const mdiIcons = [
	'github',
	'eye',
	'eye-off',
	'rocket-launch',
	'check-bold',
	'pencil',
	'pencil-outline',
	'gamepad-variant-outline',
	'gamepad-variant',
	'lock',
	'lock-outline',
	'lock-open-variant-outline',
	'magnify',
	'plus-bold',
	'plus-outline',
	'coffee',
	'coffee-outline',
	'refresh',
	'chevron-left',
	'chevron-right',
	'information',
	'account',
	'account-supervisor',
	'menu-up',
	'menu-down',
	'home',
	'home-outline',
	'weather-sunny',
	'moon-waning-crescent',
	'star-outline'
];

loadIcons('mdi', mdiIcons);
