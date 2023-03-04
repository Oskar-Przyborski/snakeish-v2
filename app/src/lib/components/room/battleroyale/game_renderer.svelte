<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { get, type Writable } from 'svelte/store';
	import Konva from 'konva';
	import type { Vector2 } from '$lib/vector';
	import type { GameState, PageState, Player } from './types';
	import { GameFramesRenderer } from '$lib/renderers/game_frames';
	import { renderSnake } from '$lib/renderers/snake';
	import { renderApple } from '$lib/renderers/apple';

	export let store: Writable<PageState>;
	let unsubscribe = () => {};
	onDestroy(unsubscribe);

	let canvasWrapper: HTMLDivElement;
	let stage: Konva.Stage;
	let gridLayer: Konva.Layer;
	let applesLayer: Konva.Layer;
	let snakesLayer: Konva.Layer;
	let waitingLayer: Konva.Layer;
	let startingLayer: Konva.Layer;
	let gameFrameRenderer: GameFramesRenderer<GameState> | null;

	onMount(() => {
		initKonva();

		gameFrameRenderer = new GameFramesRenderer<GameState>(
			50,
			get(store).gameState!.frameTime,
			get(store).gameState!,
			render
		);

		unsubscribe = store.subscribe((state) => {
			if (state.gameState == null) return;

			gameFrameRenderer?.setGameData(state.gameState);
		});
	});

	function render(state: GameState, frameCompletion: number) {
		gridLayer.hide();
		applesLayer.hide();
		snakesLayer.hide();
		waitingLayer.hide();
		startingLayer.hide();

		const cellSize = stage.width() / state.gridSize;
		switch (state.gameStatus) {
			case 'playing':
				gridLayer.show();
				applesLayer.show();
				snakesLayer.show();
				drawGrid(state.gridSize, cellSize);
				drawApples(state.apples, cellSize);
				drawPlayers(state.players, stage.width() / state.gridSize, frameCompletion);
				break;
			case 'waiting-for-players':
				waitingLayer.show();
				drawWaitingForPlayers(state.players.length, state.minPlayers);
				break;
			case 'starting':
				startingLayer.show();
				drawStarting(state.startUnix);
				break;
			case 'finished':
				// console.log("finished")
				break;
			default:
				console.log('unhandled game status');
				break;
		}
	}

	function initKonva() {
		stage = new Konva.Stage({
			container: canvasWrapper,
			width: 550,
			height: 550
		});

		gridLayer = new Konva.Layer();
		applesLayer = new Konva.Layer();
		snakesLayer = new Konva.Layer();
		waitingLayer = new Konva.Layer();
		startingLayer = new Konva.Layer();
		stage.add(gridLayer, applesLayer, snakesLayer, waitingLayer, startingLayer);
	}

	const getCssVar = (varName: string) =>
		getComputedStyle(document.documentElement).getPropertyValue(varName);

	function drawGrid(gridSize: number, cellSize: number) {
		if (gridLayer.hasChildren()) return;

		for (let y = 0; y < gridSize; y++) {
			for (let x = 0; x < gridSize; x++) {
				const darked = y % 2 == 0 ? x % 2 == 0 : x % 2 == 1;
				const color = darked ? getCssVar('--grid-1') : getCssVar('--grid-2');
				gridLayer.add(
					new Konva.Rect({
						width: cellSize,
						height: cellSize,
						x: x * cellSize,
						y: y * cellSize,
						fill: color
					})
				);
			}
		}

		gridLayer.draw();
	}

	function drawApples(apples: Vector2[], cellSize: number) {
		applesLayer.removeChildren();
		for (let i = 0; i < apples.length; i++) {
			const apple = apples[i];
			applesLayer.add(renderApple(apple, cellSize));
		}
		applesLayer.draw();
	}

	function drawPlayers(players: Player[], cellSize: number, moveValue: number) {
		snakesLayer.removeChildren();
		for (let i = 0; i < players.length; i++) {
			const player = players[i];
			const snake = renderSnake(player.snakeTail, player.direction, cellSize, moveValue, {
				color: player.color,
				name: player.name
			});
			snakesLayer.add(snake);
		}
		snakesLayer.draw();
	}

	function drawWaitingForPlayers(players: number, min: number) {
		waitingLayer.removeChildren();

		const group = new Konva.Group({
			width: stage.width(),
			height: 62,
			y: stage.height() / 2 - 31
		});

		const title = new Konva.Text({
			text: 'Waiting for players ...',
			fill: getCssVar('--text'),
			width: group.width(),
			fontSize: 28,
			fontFamily: 'DM Sans',
			align: 'center'
		});
		const subtitle = new Konva.Text({
			text: `${players} / ${min}`,
			fill: getCssVar('--text'),
			width: group.width(),
			fontSize: 26,
			fontFamily: 'DM Sans',
			align: 'center',
			y: 36
		});

		group.add(title, subtitle);
		waitingLayer.add(group);
	}

	function drawStarting(startUnix: number) {
		startingLayer.removeChildren();
		const startingIn = Math.max(Math.round((startUnix - Date.now()) / 1000), 0);

		const group = new Konva.Group({
			width: stage.width(),
			y: stage.height() / 2 - 13
		});
		const title = new Konva.Text({
			text: `Starting in ${startingIn}s`,
			fill: getCssVar('--text'),
			width: group.width(),
			align: 'center',
			fontSize: 26
		});

		group.add(title);
		startingLayer.add(group);
	}
</script>

<div bind:this={canvasWrapper} />
