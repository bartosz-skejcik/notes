<!-- file: src/routes/notebooks/[slug]/+page.svelte -->

<script lang="ts">
	import { createSearch } from '$stores/search.svelte';

	import CommandSearch from '$components/command-search.svelte';
	import ThemeSwitcher from '$components/theme-switcher.svelte';
	import { Button } from '$ui/button';

	import { flip } from 'svelte/animate';

	import { Settings, Home } from 'lucide-svelte';
	import { fetchEntries } from '$lib/entry';
	import type { PageData } from './$types';
	import type { Entry as EntryType } from '$lib/entry';
	import Entry from '$components/entry.svelte';
	import EntryForm from '$components/entry-form.svelte';
	import { quintOut } from 'svelte/easing';

	let { data }: { data: PageData } = $props();

	const search = createSearch();

	// let entryStore = $state(createEntryStore(data.sessionId, data.slug));

	let fetchedEntries = $state<EntryType[]>([]);

	$effect(() => {
		(async () => {
			if (data.slug) {
				const e = await fetchEntries({ sessionId: data.sessionId, slug: data.slug });
				fetchedEntries = e as EntryType[];
			}
		})();
	});

	$effect(() => {
		if (fetchedEntries.length > 0) {
			entries = fetchedEntries;
		}
	});

	let entries = $state<EntryType[]>([]);

	function convertEntriesToItems(entries: EntryType[]): {
		label: string;
		value: string;
		id: number;
	}[] {
		const result: {
			label: string;
			value: string;
			id: number;
		}[] = [];

		function processEntry(entry: EntryType) {
			result.push({
				label: entry.title,
				value: entry.title,
				id: entry.id
			});

			// Process children recursively
			if (entry.children && entry.children.length > 0) {
				entry.children.forEach(processEntry);
			}
		}

		entries.forEach(processEntry);

		return result;
	}
</script>

<nav class="sticky top-0 flex items-center justify-between px-3 py-2 z-4 bg-background">
	<div class="flex items-center gap-2 text-[0.97rem] text-muted-foreground">
		<p>{data.notebook.name}</p>
		<span>•</span>
		<p>
			{new Date().toLocaleString('en-PL', {
				weekday: 'long',
				year: 'numeric',
				month: 'long',
				day: 'numeric'
			})}
		</p>
	</div>
	<div class="flex items-center justify-end w-1/2 gap-4">
		<CommandSearch items={convertEntriesToItems(entries)} onChange={search.setValue} />
		<div class="flex items-center justify-center gap-3">
			<ThemeSwitcher />
			<Button size="icon" variant="ghost" class={'relative text-muted-foreground w-8 h-8'}>
				<Settings class="w-5 h-5" />
			</Button>
			<Button
				onclick={() => (window.location.href = '/')}
				size="icon"
				variant="ghost"
				class="w-8 h-8"
			>
				<Home class="w-5 h-5 text-muted-foreground" />
			</Button>
		</div>
	</div>
</nav>
<div class="p-3 pb-20 space-y-1">
	<EntryForm {...data} bind:entries />
	{#each entries as entry, i (i)}
		<div animate:flip={{ duration: 300, easing: quintOut }} class="w-full">
			<Entry bind:entries {entry} {...data} tags={data.tags} />
		</div>
	{/each}
</div>
