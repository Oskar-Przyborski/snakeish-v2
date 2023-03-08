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
	let gameLayer: Konva.Layer;
	let overlayLayer: Konva.Layer;

	let gameFrameRenderer: GameFramesRenderer<GameState> | null;

	onMount(() => {
		initKonva();

		const gameState = get(store).gameState!;
		gameFrameRenderer = new GameFramesRenderer<GameState>(
			50,
			gameState.frameTime,
			gameState,
			render
		);

		unsubscribe = store.subscribe(storeUpdate);
	});

	function initKonva() {
		stage = new Konva.Stage({
			container: canvasWrapper,
			width: 550,
			height: 550
		});

		gridLayer = new Konva.Layer();
		gameLayer = new Konva.Layer();
		overlayLayer = new Konva.Layer();
		stage.add(gridLayer, gameLayer, overlayLayer);
	}

	function storeUpdate(state: PageState) {
		if (state.gameState == null) return;

		gameFrameRenderer?.setGameData(state.gameState);
	}

	function render(state: GameState, frameCompletion: number) {
		gridLayer.hide();
		gameLayer.hide();
		gameLayer.destroyChildren();
		overlayLayer.destroyChildren();

		const cellSize = stage.width() / state.gridSize;
		switch (state.gameStatus) {
			case 'playing':
				gridLayer.show();
				gameLayer.show();
				overlayLayer.show();

				drawGrid(state.gridSize, cellSize);
				addApples(gameLayer, state.apples, cellSize);
				addPlayers(gameLayer, state.players, stage.width() / state.gridSize, frameCompletion);
				addZone(gameLayer, state.shrinkSize, cellSize, state.gridSize);
				addFreezeCountdown(overlayLayer, state.unfreezeUnix);

				break;
			case 'waiting-for-players':
				overlayLayer.show();
				addWaitingForPlayers(overlayLayer, state.players.length, state.minPlayers, state.startUnix);
				break;
			case 'finished':
				overlayLayer.show();
				addFinish(overlayLayer, state.winner);
				break;
			default:
				console.log('unhandled game status');
				break;
		}
		gameLayer.draw();
		overlayLayer.draw();
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

	function addApples(layer: Konva.Layer, apples: Vector2[], cellSize: number) {
		const group = new Konva.Group();
		layer.add(group);

		for (let i = 0; i < apples.length; i++) {
			group.add(renderApple(apples[i], cellSize));
		}
	}

	function addPlayers(layer: Konva.Layer, players: Player[], cellSize: number, moveValue: number) {
		const group = new Konva.Group();
		layer.add(group);

		for (let i = 0; i < players.length; i++) {
			const player = players[i];
			const snake = renderSnake(player.snakeTail, player.direction, cellSize, moveValue, {
				color: player.color,
				name: player.name
			});
			group.add(snake);
		}
	}

	function addZone(layer: Konva.Layer, shrink: number, cellSize: number, gridSize: number) {
		const group = new Konva.Group();
		layer.add(group);

		const opacity = 0.3;
		const fill = 'red';

		group.add(
			// left
			new Konva.Rect({
				y: 0,
				x: 0,
				width: shrink * cellSize,
				height: gridSize * cellSize,
				fill,
				opacity
			}),
			// right
			new Konva.Rect({
				x: (gridSize - shrink) * cellSize,
				y: 0,
				width: shrink * cellSize,
				height: gridSize * cellSize,
				fill,
				opacity
			}),
			// top
			new Konva.Rect({
				x: shrink * cellSize,
				y: 0,
				width: (gridSize - shrink * 2) * cellSize,
				height: shrink * cellSize,
				fill,
				opacity
			}),
			// bottom
			new Konva.Rect({
				x: shrink * cellSize,
				y: (gridSize - shrink) * cellSize,
				width: (gridSize - shrink * 2) * cellSize,
				height: shrink * cellSize,
				fill,
				opacity
			})
		);
	}

	function addFreezeCountdown(layer: Konva.Layer, unfreezeUnix: number) {
		if (unfreezeUnix <= 0) return;

		const timeLeft = Math.max(Math.round((unfreezeUnix - Date.now()) / 1000), 0);
		const text = new Konva.Text({
			text: timeLeft.toString(),
			align: 'center',
			y: stage.height() / 2,
			width: stage.width(),
			fontSize: 30,
			fontFamily: 'DM Sans',
			fill: getCssVar('--text'),
			opacity: 0.5
		});
		text.offsetY(text.height() / 2);

		layer.add(text);
	}

	function addWaitingForPlayers(
		layer: Konva.Layer,
		players: number,
		min: number,
		startUnix: number
	) {
		const group = new Konva.Group({
			width: stage.width(),
			height: 62,
			y: stage.height() / 2 - 50
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

		if (startUnix > 0) {
			const startingIn = Math.max(Math.round((startUnix - Date.now()) / 1000), 0);
			const starting = new Konva.Text({
				text: `Starting in ${startingIn}s`,
				fill: getCssVar('--text'),
				width: group.width(),
				align: 'center',
				fontSize: 25,
				y: 75
			});

			group.add(starting);
		}

		layer.add(group);
	}

	function addFinish(layer: Konva.Layer, winner: Player) {
		const group = new Konva.Group({
			width: stage.width(),
			height: 62
		});

		if (winner != null) {
			const title = new Konva.Text({
				text: `${winner.name}`,
				fill: getCssVar('--text'),
				width: group.width(),
				fontSize: 28,
				fontFamily: 'DM Sans',
				align: 'center'
			});
			const sub = new Konva.Text({
				text: `Won, with ${winner.score} points!`,
				fill: getCssVar('--text'),
				width: group.width(),
				fontSize: 24,
				fontFamily: 'DM Sans',
				align: 'center',
				y: 36
			});
			group.y(stage.height() / 2 - 30);
			group.add(title, sub);
		} else {
			const title = new Konva.Text({
				text: 'Draw!',
				fill: getCssVar('--text'),
				width: group.width(),
				fontSize: 28,
				fontFamily: 'DM Sans',
				align: 'center'
			});
			group.y(stage.height() / 2);
			group.add(title);
		}

		layer.add(group);
	}
</script>

<div bind:this={canvasWrapper} />
