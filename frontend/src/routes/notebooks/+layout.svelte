<script lang="ts">
	import CommandSearch from '$components/command-search.svelte';
	import { Button } from '$ui/button';
	import { createSearch } from '$stores/search.svelte';
	import { Settings, Home } from 'lucide-svelte';
	import ThemeSwitcher from '$components/theme-switcher.svelte';

	let { children } = $props();

	const search = createSearch();

	const items = [
		{ label: 'New Notebook', value: 'new' },
		{ label: 'Open Notebook', value: 'open' },
		{ label: 'Save Notebook', value: 'save' },
		{ label: 'Save Notebook As', value: 'save-as' },
		{ label: 'Rename Notebook', value: 'rename' },
		{ label: 'Delete Notebook', value: 'delete' },
		{ label: 'Export Notebook', value: 'export' },
		{ label: 'Import Notebook', value: 'import' },
		{ label: 'Print Notebook', value: 'print' },
		{ label: 'Print Notebook Preview', value: 'print-preview' },
		{ label: 'Exit Notebook', value: 'exit' }
	];
</script>

<div class="grid w-full h-screen grid-cols-5 overflow-hidden">
	<aside class="col-span-1 bg-muted">Sidebar</aside>
	<main class="col-span-4">
		<nav class="flex items-center justify-around px-4 py-2">
			<CommandSearch {items} onChange={search.setValue} />

			<div class="flex items-center justify-center gap-4">
				<Button size="icon" variant="ghost">
					<Settings class="w-5 h-5 text-muted-foreground" />
				</Button>
				<Button onclick={() => (window.location.href = '/')} size="icon" variant="ghost">
					<Home class="w-5 h-5 text-muted-foreground" />
				</Button>
				<ThemeSwitcher />
			</div>
		</nav>
		{@render children()}
	</main>
</div>
