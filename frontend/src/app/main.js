import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		name: 'mundo',
		version: 'VERSION',
		environment: 'ENVIRONMENT'
	}
});

export default app;