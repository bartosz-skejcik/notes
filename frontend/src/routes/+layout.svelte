<script>
	import { onMount } from 'svelte';
	import '../app.css';
	import '@fontsource-variable/inter';
	import { theme, setTheme } from '$stores/theme';
	import { onNavigate } from '$app/navigation';

	onNavigate((navigation) => {
		// @ts-ignore
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			// @ts-ignore
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});

	let { children } = $props();

	onMount(() => {
		setTheme($theme);
	});
</script>

{@render children()}
