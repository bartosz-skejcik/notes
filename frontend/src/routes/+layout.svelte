<script>
	import { onMount } from 'svelte';
	import '../app.css';
	import '@fontsource-variable/inter';
	import { theme, setTheme } from '$stores/theme';

	let { children } = $props();

	onMount(() => {
		setTheme($theme);
	});
</script>

{@render children()}
