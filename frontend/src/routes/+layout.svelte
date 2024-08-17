<script>
	import { ModeWatcher } from 'mode-watcher';
	import '../app.css';
	import '@fontsource-variable/inter';

	let { children } = $props();
</script>

<ModeWatcher />
{@render children()}
