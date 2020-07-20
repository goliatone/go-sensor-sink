import App from './App.svelte';

const app = new App({
	target: document.getElementById('app'),
	props: {
		version: 'VERSION',
		environment: 'ENVIRONMENT'
	}
});

export default app;