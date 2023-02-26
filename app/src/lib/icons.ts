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

const materialSymbolsIcons = [
	'home-rounded',
	'home-outline-rounded',
	'search-rounded',
	'coffee',
	'coffee-outline',
	'add-rounded',
	'lock',
	'lock-outline',
	'lock-open-outline-rounded',
	'refresh-rounded',
	'nightlight-rounded',
	'sunny-rounded',
	'visibility-outline-rounded',
	'visibility-off-outline-rounded',
	'rocket-launch-rounded',
	'edit-outline-rounded',
	'edit-rounded',
	'gamepad-rounded',
	'gamepad-outline-rounded',
	'chevron-right-rounded',
	'chevron-left-rounded',
	'arrow-drop-down-rounded',
	'arrow-drop-up-rounded',
	'star-outline-rounded',
	'person-outline-rounded',
	'group-outline-rounded',
	'info-outline-rounded'
];

const mdiIcons = ['github', 'check-bold'];

loadIcons('mdi', mdiIcons);
loadIcons('material-symbols', materialSymbolsIcons);
