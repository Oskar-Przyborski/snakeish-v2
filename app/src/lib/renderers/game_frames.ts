export class FrameManager {
	lastDeltaTimes: number[];
	lastFrameTime: number;
	constructor() {
		this.lastDeltaTimes = [];
		this.lastFrameTime = Date.now();
	}

	registerFrame() {
		const now = Date.now();
		const deltaTime = now - this.lastFrameTime;
		this.lastFrameTime = now;
		this.lastDeltaTimes.push(deltaTime);
		if (this.lastDeltaTimes.length > 10) {
			this.lastDeltaTimes.shift();
		}
	}

	avgDelta() {
		return this.lastDeltaTimes.reduce((prev, curr) => prev + curr) / this.lastDeltaTimes.length;
	}
}
export class GameFramesRenderer<T> {
	frameManager: FrameManager;
	fps: number;
	gameData: T;
	updateCallback: (arg0: T, arg1: number) => void;
	estimatedMainFrameTime: number;
	currentMainFrameTime = 0;

	constructor(
		fps: number,
		estimatedMainFrameTime: number,
		gameData: T,
		updateCallback: (arg0: T, arg1: number) => void
	) {
		this.estimatedMainFrameTime = estimatedMainFrameTime;
		this.fps = fps;
		this.gameData = gameData;
		this.updateCallback = updateCallback;
		this.frameManager = new FrameManager();
		this.frameManager.registerFrame();
		this.requestNextFrame(this.getFrameTime());
	}

	getFrameTime() {
		return 1000 / this.fps;
	}
	getMainUpdateFramesCount() {
		return this.estimatedMainFrameTime / this.frameManager.avgDelta();
	}

	setGameData(data: T) {
		this.gameData = data;
		this.currentMainFrameTime = 0;
	}

	private requestNextFrame(inMilliseconds: number) {
		this.currentMainFrameTime += this.frameManager.avgDelta();
		setTimeout(() => {
			const realFrameTime =
				this.getFrameTime() * (this.getFrameTime() / this.frameManager.avgDelta());
			this.updateCallback(this.gameData, this.currentMainFrameTime / this.estimatedMainFrameTime);
			this.frameManager.registerFrame();
			this.requestNextFrame(realFrameTime);
		}, inMilliseconds);
	}
}
